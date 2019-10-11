package kubeclient

import (
	security "github.com/openshift/client-go/security/clientset/versioned"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	v1 "k8s.io/client-go/tools/clientcmd/api/v1"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/util/managedcluster"
	"github.com/openshift/openshift-azure/pkg/util/roundtrippers"
)

// NewKubeclient creates a new kubeclient
// If PrivateEndpointIp is not nil - kubeclient roundtripper will point PrivateEndpoint IP address
func NewKubeclient(log *logrus.Entry, config *v1.Config, disableKeepAlives bool, location string, privateEndpointIP *string, testConfig api.TestConfig) (Interface, error) {
	restconfig, err := managedcluster.RestConfigFromV1Config(config)
	if err != nil {
		return nil, err
	}

	if privateEndpointIP == nil {
		restconfig.WrapTransport = roundtrippers.NewRetryingRoundTripper(log, disableKeepAlives)
	} else {
		tlsConfig, err := restclient.TLSConfigFor(restconfig)
		if err != nil {
			return nil, err
		}

		restconfig.WrapTransport = roundtrippers.NewPrivateEndpoint(log, location, *privateEndpointIP, disableKeepAlives, testConfig, tlsConfig)
		if testConfig.RunningUnderTest {
			//override dialer in manual mode
			restconfig.Host = *privateEndpointIP
		}
	}

	cli, err := kubernetes.NewForConfig(restconfig)
	if err != nil {
		return nil, err
	}

	seccli, err := security.NewForConfig(restconfig)
	if err != nil {
		return nil, err
	}

	return &Kubeclientset{
		Log:               log,
		Client:            cli,
		Seccli:            seccli,
		disableKeepAlives: disableKeepAlives,
		restconfig:        restconfig,
		testConfig:        testConfig,
	}, nil
}
