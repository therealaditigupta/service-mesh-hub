package translator

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/solo-io/mesh-projects/pkg/api/external/istio/networking/v1alpha3"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
	zephyr_core "github.com/solo-io/mesh-projects/pkg/api/v1/core"
	"github.com/solo-io/mesh-projects/services/internal/kube"
	"github.com/solo-io/mesh-projects/services/internal/networking"
	"github.com/solo-io/mesh-projects/services/mesh-bridge/pkg/setup/config"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type Translator interface {
	Translate(ctx context.Context, meshBridgesByNamespace MeshBridgesByNamespace) (v1alpha3.ServiceEntryList, error)
}

func NewMeshBridgeTranslator(clientset config.ClientSet) Translator {
	return &translator{
		clientset: clientset,
	}
}

type translator struct {
	clientset config.ClientSet
}

type MeshBridgesByNamespace map[string]v1.MeshBridgeList
type ServiceEntriesByNamespace map[string]v1alpha3.ServiceEntryList

func (t *translator) Translate(ctx context.Context, meshBridgesByNamespace MeshBridgesByNamespace) (v1alpha3.ServiceEntryList, error) {
	var result v1alpha3.ServiceEntryList
	ipGen := newIpGenerator()
	for namespace, meshBridges := range meshBridgesByNamespace {
		serviceEntries, err := t.meshBridgesToServiceEntry(ctx, namespace, meshBridges, ipGen)
		if err != nil {
			return nil, err
		}
		result = append(result, serviceEntries...)
	}
	return result, nil
}

func (t *translator) meshBridgesToServiceEntry(ctx context.Context, namespace string, meshBridges v1.MeshBridgeList,
	generator *ipGenerator) ([]*v1alpha3.ServiceEntry, error) {

	var result []*v1alpha3.ServiceEntry

	uniqueAddresses := make(map[string]*BridgeDetails)
	for _, meshBridge := range meshBridges {
		meshRef := meshBridge.GetTarget().GetMeshService().GetMesh()
		mesh, err := t.clientset.Mesh().Read(meshRef.GetNamespace(),
			meshRef.GetName(), clients.ReadOpts{})
		if err != nil {
			return nil, err
		}

		entryPoint, err := t.clientset.MeshIngress().Read(mesh.GetEntryPoint().GetResource().GetNamespace(),
			mesh.GetEntryPoint().GetResource().GetName(), clients.ReadOpts{})
		if err != nil {
			return nil, err
		}

		var (
			address string
			port    uint32
		)
		switch typedIngress := entryPoint.GetIngressType().(type) {
		case *v1.MeshIngress_Gloo:
			clusterRestCfg, err := kube.GetKubeConfigForCluster(ctx, t.clientset.LocalClientGo(),
				mesh.GetDiscoveryMetadata().GetCluster())
			if err != nil {
				return nil, err
			}
			address, port, err = networking.GetIngressHostAndPort(clusterRestCfg, &zephyr_core.ClusterResourceRef{
				Resource: &core.ResourceRef{
					Name:      typedIngress.Gloo.GetServiceName(),
					Namespace: typedIngress.Gloo.GetNamespace(),
				},
				Cluster: mesh.GetDiscoveryMetadata().GetCluster(),
			}, typedIngress.Gloo.GetPort())
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.Errorf("current only gloo ingress types are supported")
		}

		if bridgeWithExit, ok := uniqueAddresses[address]; ok {
			bridgeWithExit.meshBridges[meshBridge] = mesh.GetDiscoveryMetadata().GetCluster()
			bridgeWithExit.ports = append(bridgeWithExit.ports, port)
		} else {
			uniqueAddresses[address] = &BridgeDetails{
				meshBridges: map[*v1.MeshBridge]string{meshBridge: mesh.GetDiscoveryMetadata().GetCluster()},
				ports:       []uint32{port},
			}
		}
	}

	for address, bridges := range uniqueAddresses {
		info, err := t.infoFromTargetServices(bridges.meshBridges)
		if err != nil {
			return nil, err
		}
		ports := make(map[string]uint32)
		for i, v := range bridges.ports {
			ports[fmt.Sprintf("http%d", i+1)] = v
		}
		serviceEntry := &v1alpha3.ServiceEntry{
			Metadata: core.Metadata{
				Name:      addressToServiceEntryName(address, namespace),
				Namespace: namespace,
			},
			Hosts:      info.hosts,
			Addresses:  []string{generator.nextIp()},
			Ports:      info.ports,
			Location:   v1alpha3.ServiceEntry_MESH_INTERNAL,
			Resolution: v1alpha3.ServiceEntry_DNS,
			Endpoints: []*v1alpha3.ServiceEntry_Endpoint{
				{
					Address: address,
					Ports:   ports,
				},
			},
			// Only apply the service entry to the local namespace
			ExportTo: []string{"."},
		}
		result = append(result, serviceEntry)
	}
	return result, nil
}

type BridgeDetails struct {
	meshBridges map[*v1.MeshBridge]string
	ports       []uint32
}

type UpstreamInfo struct {
	hosts []string
	ports []*v1alpha3.Port
}

func (t *translator) infoFromTargetServices(list map[*v1.MeshBridge]string) (*UpstreamInfo, error) {
	hosts := make(map[string]bool)
	var endpoints []*v1alpha3.Port
	i := 0
	for bridge, cluster := range list {
		i += 1
		upstreamRef := bridge.GetTarget().GetMeshService().GetUpstream()
		upstream, err := t.clientset.Upstreams().Read(upstreamRef.GetNamespace(),
			upstreamRef.GetName(), clients.ReadOpts{
				Cluster: cluster,
			})
		if err != nil {
			return nil, err
		}
		kubeUpstream := upstream.GetUpstreamSpec().GetKube()
		if kubeUpstream == nil {
			return nil, fmt.Errorf("currently only kube upstreams are supported, %s supplied",
				bridge.GetTarget().String())
		}
		hostName := fmt.Sprintf("%s.%s.global", kubeUpstream.GetServiceName(), kubeUpstream.GetServiceNamespace())
		hosts[hostName] = true
		endpoints = append(endpoints, &v1alpha3.Port{
			Number:   kubeUpstream.GetServicePort(),
			Protocol: "http",
			Name:     fmt.Sprintf("http%d", i),
		})
	}
	var result []string
	for service, _ := range hosts {
		result = append(result, service)
	}
	return &UpstreamInfo{
		hosts: result,
		ports: endpoints,
	}, nil
}

func addressToServiceEntryName(address string, namespace string) string {
	return fmt.Sprintf("%s-mesh-bridge-%s", namespace, address)
}
