package consul

import (
	"github.com/solo-io/smh/pkg/mesh-discovery/snapshot/translation/mesh/detector/deployment"
	"github.com/solo-io/smh/pkg/mesh-discovery/snapshot/translation/mesh/detector/deployment/utils"
	"regexp"
	"strings"

	consulconfig "github.com/hashicorp/consul/agent/config"
	"github.com/hashicorp/hcl"
	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

var (
	ErrorDetectingDeployment = func(err error) error {
		return eris.Wrap(err, "Error while detecting consul deployment")
	}
	InvalidImageFormatError = func(err error, imageName string) error {
		return eris.Wrapf(err, "invalid or unexpected image format for image name: %s", imageName)
	}
	HclParseError = func(err error, invalidHcl string) error {
		return eris.Wrapf(err, "error parsing HCL in consul invocation: %s", invalidHcl)
	}
)

var (
	// a consul's invocation can include a line like:
	// -hcl="connect { enabled = true }"
	// hcl is HashiCorp configuration Language
	// https://github.com/hashicorp/hcl
	hclRegex = regexp.MustCompile("-hcl=\"([^\"]*)\"")
)

const (
	consulServerArg           = "-server"
	normalizedConsulImagePath = "library/consul"
)

// detects Consul Connect if a deployment contains the istiod container.
type meshDetector struct{}

func NewMeshDetector() deployment.MeshDetector {
	return &meshDetector{}
}

func (c *meshDetector) DetectMesh(deployment *appsv1.Deployment) (*v1alpha1.Mesh, error) {
	for _, container := range deployment.Spec.Template.Spec.Containers {
		isConsulInstallation, err := isConsulConnect(container)
		if err != nil {
			return nil, ErrorDetectingDeployment(err)
		}

		if !isConsulInstallation {
			continue
		}

		return c.buildConsulMeshObject(deployment, container)
	}

	return nil, nil
}

// returns an error if the image name is un-parsable
func (c *meshDetector) buildConsulMeshObject(
	deployment *appsv1.Deployment,
	container corev1.Container,
) (*v1alpha1.Mesh, error) {

	parsedImage, err := docker.ParseImageName(container.Image)
	if err != nil {
		return nil, err
	}

	imageVersion := parsedImage.Tag
	if parsedImage.Digest != "" {
		imageVersion = parsedImage.Digest
	}

	return &v1alpha1.Mesh{
		ObjectMeta: utils.MeshObjectMeta(deployment),
		Spec: v1alpha1.MeshSpec{
			MeshType: &v1alpha1.MeshSpec_ConsulConnect{
				ConsulConnect: &v1alpha1.MeshSpec_ConsulConnectMesh{
					Installation: &v1alpha1.MeshSpec_MeshInstallation{
						Namespace: deployment.Namespace,
						Cluster:   deployment.ClusterName,
						Version:   imageVersion,
					},
				},
			},
		},
	}, nil
}

func isConsulConnect(container corev1.Container) (bool, error) {
	parsedImage, err := docker.ParseImageName(container.Image)
	if err != nil {
		return false, InvalidImageFormatError(err, container.Image)
	}

	// if the image appears to be a consul image, and
	// the container is starting up with a "-server" arg,
	// then declare that we've found consul
	if parsedImage.Path != normalizedConsulImagePath {
		return false, nil
	}

	cmd := strings.Join(container.Command, " ")

	isServerMode := strings.Contains(cmd, consulServerArg)
	if !isServerMode {
		return false, nil
	}

	hclMatches := hclRegex.FindStringSubmatch(cmd)
	if len(hclMatches) < 2 {
		return false, nil
	}

	config := &consulconfig.Config{}
	err = hcl.Decode(config, hclMatches[1])
	if err != nil {
		return false, HclParseError(err, hclMatches[1])
	}

	return config.Connect.Enabled != nil && *config.Connect.Enabled, nil
}
