// Package plugin holds the implementation of a plugin.
package plugin

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/cluster"
	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/log"
)

type plugin struct {
	entry        *logrus.Entry
	config       api.PluginConfig
	updater      api.Updater
	deployer     api.Deployer
	armGenerator arm.Generator
}

var _ api.Plugin = &plugin{}

// NewPlugin creates a new plugin instance
func NewPlugin(entry *logrus.Entry, pluginConfig *api.PluginConfig) api.Plugin {
	log.New(entry)
	return &plugin{
		entry:        entry,
		config:       *pluginConfig,
		deployer:     NewSimpleDeployer(entry, pluginConfig),
		updater:      NewSimpleUpdater(entry, pluginConfig),
		armGenerator: arm.NewSimpleGenerator(entry),
	}
}

func (p *plugin) MergeConfig(ctx context.Context, cs, oldCs *api.OpenShiftManagedCluster) {
	if oldCs == nil {
		return
	}
	log.Info("merging internal data models")

	// generated config should be copied as is
	old := oldCs.DeepCopy()
	cs.Config = old.Config

	// user request data
	// need to merge partial requests
	if len(cs.Properties.AgentPoolProfiles) == 0 {
		cs.Properties.AgentPoolProfiles = oldCs.Properties.AgentPoolProfiles
	}
	if len(cs.Properties.OpenShiftVersion) == 0 {
		cs.Properties.OpenShiftVersion = oldCs.Properties.OpenShiftVersion
	}
	if len(cs.Properties.PublicHostname) == 0 {
		cs.Properties.PublicHostname = oldCs.Properties.PublicHostname
	}
	if cs.Properties.NetworkProfile == nil {
		cs.Properties.NetworkProfile = oldCs.Properties.NetworkProfile
	}
	if len(cs.Properties.RouterProfiles) == 0 {
		cs.Properties.RouterProfiles = oldCs.Properties.RouterProfiles
	}
	if cs.Properties.ServicePrincipalProfile == nil {
		cs.Properties.ServicePrincipalProfile = oldCs.Properties.ServicePrincipalProfile
	}
	if cs.Properties.AzProfile == nil {
		cs.Properties.AzProfile = oldCs.Properties.AzProfile
	}
	if cs.Properties.AuthProfile == nil {
		cs.Properties.AuthProfile = oldCs.Properties.AuthProfile
	}
	if len(cs.Properties.FQDN) == 0 {
		cs.Properties.FQDN = oldCs.Properties.FQDN
	}
}

func (p *plugin) Validate(ctx context.Context, new, old *api.OpenShiftManagedCluster, externalOnly bool) []error {
	log.Info("validating internal data models")
	return api.Validate(new, old, externalOnly)
}

func (p *plugin) GenerateConfig(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	log.Info("generating configs")
	// TODO should we save off the original config here and if there are any errors we can restore it?
	if cs.Config == nil {
		cs.Config = &api.Config{}
	}

	err := config.Generate(cs, p.config)
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GenerateARM(ctx context.Context, cs *api.OpenShiftManagedCluster, isUpdate bool) ([]byte, error) {
	log.Info("generating arm templates")
	return p.armGenerator.Generate(ctx, cs, isUpdate)
}

func (p *plugin) Deployer() api.Deployer {
	return p.deployer
}

func (p *plugin) Updater() api.Updater {
	return p.updater
}

func (p *plugin) CreateOrUpdate(ctx context.Context, cs *api.OpenShiftManagedCluster, azuredeploy []byte, isUpdate bool) error {
	if isUpdate {
		log.Info("starting update")
		nodes, err := p.Updater().GetNodesPreUpdate(ctx, cs)
		if err != nil {
			return err
		}
		err = p.Updater().DefaultDeployment(ctx, cs, azuredeploy)
		if err != nil {
			return err
		}
		err = p.Updater().Initialize(ctx, cs)
		if err != nil {
			return err
		}
		err = p.Updater().WaitForCompletion(ctx, cs, nodes)
		if err != nil {
			return err
		}
		return p.Updater().UpdateInPlace(ctx, cs)
	}

	log.Info("starting deploy")
	err := p.Deployer().DefaultDeployment(ctx, cs, azuredeploy)
	if err != nil {
		return err
	}
	err = p.Deployer().Initialize(ctx, cs)
	if err != nil {
		return err
	}
	return p.Deployer().WaitForCompletion(ctx, cs)
}

type simpleDeployer struct {
	pluginConfig    api.PluginConfig
	clusterUpgrader cluster.Upgrader
}

var _ api.Deployer = &simpleDeployer{}

// NewSimpleDeployer create a new Upgrade instance.
func NewSimpleDeployer(entry *logrus.Entry, pluginConfig *api.PluginConfig) api.Deployer {
	log.New(entry)
	return &simpleDeployer{
		pluginConfig:    *pluginConfig,
		clusterUpgrader: cluster.NewSimpleUpgrader(entry, pluginConfig),
	}
}

func (d *simpleDeployer) DefaultDeployment(ctx context.Context, cs *api.OpenShiftManagedCluster, azuredeploy []byte) error {
	return d.clusterUpgrader.DefaultDeploy(ctx, cs, azuredeploy)
}

func (d *simpleDeployer) Initialize(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	return d.clusterUpgrader.Initialize(ctx, cs)
}

func (d *simpleDeployer) WaitForCompletion(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	// 1.
	err := d.clusterUpgrader.WaitForNodes(ctx, cs)
	if err != nil {
		return err
	}

	// 2.
	err = d.clusterUpgrader.WaitForInfraServices(ctx, cs)
	if err != nil {
		return err
	}
	// 3
	return d.clusterUpgrader.WaitForConsole(ctx, cs)
}

type simpleUpgrader struct {
	pluginConfig    api.PluginConfig
	clusterUpgrader cluster.Upgrader
}

var _ api.Updater = &simpleUpgrader{}

// NewSimpleUpdater create a new Upgrade instance.
func NewSimpleUpdater(entry *logrus.Entry, pluginConfig *api.PluginConfig) api.Updater {
	log.New(entry)
	return &simpleUpgrader{
		pluginConfig:    *pluginConfig,
		clusterUpgrader: cluster.NewSimpleUpgrader(entry, pluginConfig),
	}
}

func (d *simpleUpgrader) GetNodesPreUpdate(ctx context.Context, cs *api.OpenShiftManagedCluster) (map[string]struct{}, error) {
	return d.clusterUpgrader.GetNodesPreUpdate(ctx, cs)
}

func (d *simpleUpgrader) DefaultDeployment(ctx context.Context, cs *api.OpenShiftManagedCluster, azuredeploy []byte) error {
	return d.clusterUpgrader.DefaultDeploy(ctx, cs, azuredeploy)
}

func (d *simpleUpgrader) Initialize(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	return d.clusterUpgrader.Initialize(ctx, cs)
}

func (d *simpleUpgrader) UpdateInPlace(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	return d.clusterUpgrader.UpdateInPlace(ctx, cs)
}

func (d *simpleUpgrader) WaitForCompletion(ctx context.Context, cs *api.OpenShiftManagedCluster, nodes map[string]struct{}) error {
	// 1.
	err := d.clusterUpgrader.WaitForNewNodes(ctx, cs, nodes)
	if err != nil {
		return err
	}

	// 2.
	err = d.clusterUpgrader.WaitForInfraServices(ctx, cs)
	if err != nil {
		return err
	}
	// 3
	return d.clusterUpgrader.WaitForConsole(ctx, cs)
}
