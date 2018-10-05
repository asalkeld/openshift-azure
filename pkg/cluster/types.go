package cluster

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/log"
)

type Upgrader interface {
	Initialize(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	DefaultDeploy(ctx context.Context, cs *api.OpenShiftManagedCluster, azuredeploy []byte) error
	GetNodesPreUpdate(ctx context.Context, cs *api.OpenShiftManagedCluster) (map[string]struct{}, error)
	WaitForInfraServices(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	WaitForNodes(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	WaitForNewNodes(ctx context.Context, cs *api.OpenShiftManagedCluster, nodes map[string]struct{}) error
	UpdateInPlace(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	WaitForConsole(ctx context.Context, cs *api.OpenShiftManagedCluster) error
}

type simpleUpgrader struct {
	pluginConfig api.PluginConfig
}

var _ Upgrader = &simpleUpgrader{}

func NewSimpleUpgrader(entry *logrus.Entry, pluginConfig *api.PluginConfig) Upgrader {
	log.New(entry)
	return &simpleUpgrader{
		pluginConfig: *pluginConfig,
	}
}
