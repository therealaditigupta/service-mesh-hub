package istio

import (
	"context"
	"strings"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/rotisserie/eris"
	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	skv1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/mesh/detector"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/utils"
	istiov1alpha1 "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/util/gogoprotomarshal"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	legacyPilotDeploymentName = "istio-pilot"
	istiodDeploymentName      = "istiod"
	istioContainerKeyword     = "istio"
	pilotContainerKeyword     = "pilot"
	istioConfigMapName        = "istio"
	istioConfigMapMeshDataKey = "mesh"

	// https://istio.io/docs/ops/deployment/requirements/#ports-used-by-istio
	defaultGatewayPortName = "tls"
)

var (
	// "istio": "ingressgateway" is a known string pair to Istio- it's semantically meaningful but unfortunately not exported from anywhere
	// their ingress gateway is hardcoded in their own implementation to have this label
	// https://github.com/istio/istio/blob/4e27ddc64f6a12e622c4cd5c836f5d7edf94e971/istioctl/cmd/describe.go#L1138
	defaultGatewayWorkloadLabels = map[string]string{
		"istio": "ingressgateway",
	}
)

// detects Istio if a deployment contains the istiod container.
type meshDetector struct {
	ctx        context.Context
	configMaps corev1sets.ConfigMapSet
	services   corev1sets.ServiceSet
	pods       corev1sets.PodSet
	nodes      corev1sets.NodeSet
}

func NewMeshDetector(
	ctx context.Context,
	configMaps corev1sets.ConfigMapSet,
	services corev1sets.ServiceSet,
	pods corev1sets.PodSet,
	nodes corev1sets.NodeSet,
) detector.MeshDetector {
	return &meshDetector{
		ctx:        contextutils.WithLogger(ctx, "detector"),
		configMaps: configMaps,
		services:   services,
		pods:       pods,
		nodes:      nodes,
	}
}

// returns nil, nil of the deployment does not contain the istiod image
func (d *meshDetector) DetectMesh(deployment *appsv1.Deployment) (*v1alpha1.Mesh, error) {
	version, err := d.getIstiodVersion(deployment)
	if err != nil {
		return nil, err
	}

	if version == "" {
		return nil, nil
	}

	trustDomain, err := getTrustDomain(d.configMaps, deployment.ClusterName, deployment.Namespace)
	if err != nil {
		return nil, err
	}

	// TODO(ilackarms): allow configuring ingress gateway workload labels
	ingressGateways := getIngressGateways(
		d.ctx,
		deployment.Namespace,
		defaultGatewayWorkloadLabels,
		d.services,
		d.pods,
		d.nodes,
	)

	mesh := &v1alpha1.Mesh{
		ObjectMeta: utils.DiscoveredObjectMeta(deployment),
		Spec: v1alpha1.MeshSpec{
			MeshType: &v1alpha1.MeshSpec_Istio_{
				Istio: &v1alpha1.MeshSpec_Istio{
					Installation: &v1alpha1.MeshSpec_MeshInstallation{
						Namespace: deployment.Namespace,
						Cluster:   deployment.ClusterName,
						Version:   version,
					},
					CitadelInfo: &v1alpha1.MeshSpec_Istio_CitadelInfo{
						TrustDomain: trustDomain,
						// This assumes that the istiod deployment is the cert provider
						CitadelServiceAccount: deployment.Spec.Template.Spec.ServiceAccountName,
					},
					IngressGateways: ingressGateways,
				},
			},
		},
	}

	return mesh, nil
}

func getIngressGateways(
	ctx context.Context,
	namespace string,
	workloadLabels map[string]string,
	allServices corev1sets.ServiceSet,
	allPods corev1sets.PodSet,
	allNodes corev1sets.NodeSet,
) []*v1alpha1.MeshSpec_Istio_IngressGatewayInfo {
	ingressSvcs := getIngressServices(allServices, namespace, workloadLabels)
	var ingressGateways []*v1alpha1.MeshSpec_Istio_IngressGatewayInfo
	for _, svc := range ingressSvcs {
		gateway, err := getIngressGateway(svc, workloadLabels, allPods, allNodes)
		if err != nil {
			contextutils.LoggerFrom(ctx).Warnw("detection failed for matching istio ingress service", "error", err, "service", sets.Key(svc))
		}
		ingressGateways = append(ingressGateways, gateway)
	}
	return ingressGateways
}

func getIngressGateway(
	svc *corev1.Service,
	workloadLabels map[string]string,
	allPods corev1sets.PodSet,
	allNodes corev1sets.NodeSet,
) (*v1alpha1.MeshSpec_Istio_IngressGatewayInfo, error) {
	var (
		tlsPort *corev1.ServicePort
	)
	for _, port := range svc.Spec.Ports {
		port := port // pike
		if port.Name == defaultGatewayPortName {
			tlsPort = &port
			break
		}
	}
	if tlsPort == nil {
		return nil, eris.Errorf("no TLS port found on ingress gateway")
	}

	var (
		// TODO(ilackarms): currently we only use one address to connect to the gateway.
		// We can support multiple addresses per gateway for load balancing purposes in the future

		externalAddress string
		externalPort    uint32
	)
	switch svc.Spec.Type {
	case corev1.ServiceTypeNodePort:
		addr, err := getNodeIp(
			svc.ClusterName,
			svc.Namespace,
			workloadLabels,
			allPods,
			allNodes,
		)
		if err != nil {
			return nil, err
		}
		externalAddress = addr
		externalPort = uint32(tlsPort.NodePort)

	case corev1.ServiceTypeLoadBalancer:
		ingress := svc.Status.LoadBalancer.Ingress
		if len(ingress) == 0 {
			return nil, eris.Errorf("no loadBalancer.ingress status reported for service")
		}

		externalAddress = ingress[0].Hostname
		if externalAddress == "" {
			externalAddress = ingress[0].IP
		}
		externalPort = uint32(tlsPort.Port)

	default:
		return nil, eris.Errorf("unsupported service type %v for ingress gateway", svc.Spec.Type)
	}

	if tlsPort.TargetPort.StrVal != "" {
		// TODO(ilackarms): for the sake of simplicity, we only support number target ports
		// if we come across the need to support named ports, we can add the lookup on the pod container itself here
		return nil, eris.Errorf("named target ports are not currently supported on ingress gateway")
	}
	containerPort := tlsPort.TargetPort.IntVal
	if containerPort == 0 {
		containerPort = tlsPort.Port
	}

	return &v1alpha1.MeshSpec_Istio_IngressGatewayInfo{
		WorkloadLabels:   workloadLabels,
		ExternalAddress:  externalAddress,
		ExternalTlsPort:  externalPort,
		TlsContainerPort: uint32(containerPort),
	}, nil
}

func getIngressServices(
	allServices corev1sets.ServiceSet,
	namespace string,
	workloadLabels map[string]string,
) []*corev1.Service {
	return allServices.List(func(svc *corev1.Service) bool {
		return svc.Namespace != namespace ||
			!labels.SelectorFromSet(svc.Spec.Selector).Matches(labels.Set(workloadLabels))
	})
}

func getNodeIp(
	cluster,
	namespace string,
	workloadLabels map[string]string,
	pods corev1sets.PodSet,
	nodes corev1sets.NodeSet,
) (string, error) {
	ingressPods := pods.List(func(pod *corev1.Pod) bool {
		return pod.ClusterName != cluster ||
			pod.Namespace != namespace ||
			!labels.SelectorFromSet(workloadLabels).Matches(labels.Set(pod.Labels))
	})
	if len(ingressPods) < 1 {
		return "", eris.Errorf("no pods found backing ingress workload %v in namespace %v", workloadLabels, namespace)
	}
	// TODO(ilackarms): currently we just grab the node ip of the first available pod
	// Eventually we may want to consider supporting multiple nodes/IPs for load balancing.
	ingressPod := ingressPods[0]
	ingressNode, err := nodes.Find(&skv1.ClusterObjectRef{
		ClusterName: cluster,
		Name:        ingressPod.Spec.NodeName,
	})
	if err != nil {
		return "", eris.Wrapf(err, "failed to find ingress node for pod %v", sets.Key(ingressPod))
	}

	for _, addr := range ingressNode.Status.Addresses {
		if addr.Type == corev1.NodeInternalIP || addr.Type == corev1.NodeInternalDNS {
			// skip internal addresses
			continue
		}
		// select the first external address of the node
		return addr.Address, nil
	}
	return "", eris.Errorf("no external addresses reported for ingress node %v", sets.Key(ingressNode))
}

func (d *meshDetector) getIstiodVersion(deployment *appsv1.Deployment) (string, error) {
	for _, container := range deployment.Spec.Template.Spec.Containers {
		if isIstiod(deployment, &container) {
			parsedImage, err := docker.ParseImageName(container.Image)
			if err != nil {
				return "", eris.Wrapf(err, "failed to parse istiod image tag: %s", container.Image)
			}
			version := parsedImage.Tag
			if parsedImage.Digest != "" {
				version = parsedImage.Digest
			}
			return version, nil
		}
	}
	return "", nil
}

// Return true if deployment is inferred to be an Istiod deployment
func isIstiod(deployment *appsv1.Deployment, container *corev1.Container) bool {
	return (deployment.GetName() == istiodDeploymentName || deployment.GetName() == legacyPilotDeploymentName) &&
		strings.Contains(container.Image, istioContainerKeyword) &&
		strings.Contains(container.Image, pilotContainerKeyword)
}

func getTrustDomain(
	configMaps corev1sets.ConfigMapSet,
	cluster,
	namespace string,
) (string, error) {
	istioConfigMap, err := configMaps.Find(&skv1.ClusterObjectRef{
		Name:        istioConfigMapName,
		Namespace:   namespace,
		ClusterName: cluster,
	})
	if err != nil {
		return "", err
	}

	meshConfigString, ok := istioConfigMap.Data[istioConfigMapMeshDataKey]
	if !ok {
		return "", eris.Errorf("Failed to find 'mesh' entry in ConfigMap with name/namespace/cluster %s/%s/%s", istioConfigMapName, namespace, cluster)
	}
	var meshConfig istiov1alpha1.MeshConfig
	err = gogoprotomarshal.ApplyYAML(meshConfigString, &meshConfig)
	if err != nil {
		return "", eris.Errorf("Failed to find 'mesh' entry in ConfigMap with name/namespace/cluster %s/%s/%s", istioConfigMapName, namespace, cluster)
	}
	return meshConfig.TrustDomain, nil
}