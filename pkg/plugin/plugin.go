// Package plugin holds the implementation of a plugin.
package plugin

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	pluginapi "github.com/openshift/openshift-azure/pkg/api/plugin/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/cluster"
	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/util/resourceid"
)

type plugin struct {
	log                     *logrus.Entry
	config                  api.PluginConfig
	clusterUpgrader         cluster.Upgrader
	armGenerator            arm.Generator
	configGenerator         config.Generator
	apiValidator            *api.Validator
	adminValidator          *api.Validator
	pluginTemplateValidator *api.Validator
}

var _ api.Plugin = &plugin{}

// NewPlugin creates a new plugin instance
func NewPlugin(log *logrus.Entry, pluginConfig *api.PluginConfig) (api.Plugin, []error) {
	return &plugin{
		log:                     log,
		config:                  *pluginConfig,
		clusterUpgrader:         cluster.NewSimpleUpgrader(log, pluginConfig),
		armGenerator:            arm.NewSimpleGenerator(pluginConfig),
		configGenerator:         config.NewSimpleGenerator(pluginConfig),
		apiValidator:            api.NewValidator(pluginConfig.TestConfig.RunningUnderTest),
		pluginTemplateValidator: api.NewValidator(pluginConfig.TestConfig.RunningUnderTest),
		adminValidator:          api.NewAdminValidator(pluginConfig.TestConfig.RunningUnderTest),
	}, nil
}

func (p *plugin) Validate(ctx context.Context, new, old *api.OpenShiftManagedCluster, externalOnly bool) []error {
	p.log.Info("validating internal data models")
	return p.apiValidator.Validate(new, old, externalOnly)
}

func (p *plugin) ValidateAdmin(ctx context.Context, new, old *api.OpenShiftManagedCluster) []error {
	p.log.Info("validating internal admin data models")
	return p.adminValidator.Validate(new, old, false)
}

func (p *plugin) ValidatePluginTemplate(ctx context.Context, template *pluginapi.Config) []error {
	p.log.Info("validating external plugin api data models")
	return p.pluginTemplateValidator.ValidatePluginTemplate(template)
}

func (p *plugin) GenerateConfig(ctx context.Context, cs *api.OpenShiftManagedCluster, template *pluginapi.Config) error {
	p.log.Info("generating configs")
	// TODO should we save off the original config here and if there are any errors we can restore it?
	err := p.configGenerator.Generate(cs, template)
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) RecoverEtcdCluster(ctx context.Context, cs *api.OpenShiftManagedCluster, deployFn api.DeployFn, backupBlob string) *api.PluginError {
	suffix := fmt.Sprintf("%d", time.Now().Unix())

	p.log.Info("generating arm templates")
	azuretemplate, err := p.armGenerator.Generate(ctx, cs, backupBlob, true, suffix)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepGenerateARM}
	}

	err = p.clusterUpgrader.CreateClients(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepClientCreation}
	}
	p.log.Info("restoring the cluster")
	if err := p.clusterUpgrader.EtcdRestoreDeleteMasterScaleSet(ctx, cs); err != nil {
		return err
	}
	err = deployFn(ctx, azuretemplate)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepDeploy}
	}
	err = p.clusterUpgrader.Initialize(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepInitialize}
	}
	if err := p.clusterUpgrader.EtcdRestoreDeleteMasterScaleSetHashes(ctx, cs); err != nil {
		return err
	}
	err = p.clusterUpgrader.WaitForHealthzStatusOk(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepWaitForWaitForOpenShiftAPI}
	}
	for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleMaster) {
		err := p.clusterUpgrader.WaitForNodesInAgentPoolProfile(ctx, cs, &app, "")
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepWaitForNodes}
		}
	}

	p.log.Info("running CreateOrUpdate")
	if err := p.CreateOrUpdate(ctx, cs, true, deployFn); err != nil {
		return err
	}
	return nil
}

func (p *plugin) CreateOrUpdate(ctx context.Context, cs *api.OpenShiftManagedCluster, isUpdate bool, deployFn api.DeployFn) *api.PluginError {
	suffix := fmt.Sprintf("%d", time.Now().Unix())

	p.log.Info("generating arm templates")
	azuretemplate, err := p.armGenerator.Generate(ctx, cs, "", isUpdate, suffix)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepGenerateARM}
	}
	err = p.clusterUpgrader.CreateClients(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepClientCreation}
	}
	if isUpdate {
		p.log.Info("starting update")
	} else {
		p.log.Info("starting deploy")
	}
	err = deployFn(ctx, azuretemplate)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepDeploy}
	}
	err = p.clusterUpgrader.Initialize(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepInitialize}
	}
	if isUpdate {
		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleMaster) {
			if perr := p.clusterUpgrader.UpdateMasterAgentPool(ctx, cs, &app); perr != nil {
				return perr
			}
		}
		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleInfra) {
			if perr := p.clusterUpgrader.UpdateWorkerAgentPool(ctx, cs, &app, suffix); perr != nil {
				return perr
			}
		}
		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleCompute) {
			if perr := p.clusterUpgrader.UpdateWorkerAgentPool(ctx, cs, &app, suffix); perr != nil {
				return perr
			}
		}
	} else {
		err = p.clusterUpgrader.InitializeUpdateBlob(cs, suffix)
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepInitializeUpdateBlob}
		}
		err = p.clusterUpgrader.WaitForHealthzStatusOk(ctx, cs)
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepWaitForWaitForOpenShiftAPI}
		}

		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleMaster) {
			err := p.clusterUpgrader.WaitForNodesInAgentPoolProfile(ctx, cs, &app, "")
			if err != nil {
				return &api.PluginError{Err: err, Step: api.PluginStepWaitForNodes}
			}
		}
		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleInfra) {
			err := p.clusterUpgrader.WaitForNodesInAgentPoolProfile(ctx, cs, &app, suffix)
			if err != nil {
				return &api.PluginError{Err: err, Step: api.PluginStepWaitForNodes}
			}
		}
		for _, app := range p.clusterUpgrader.SortedAgentPoolProfilesForRole(cs, api.AgentPoolProfileRoleCompute) {
			err := p.clusterUpgrader.WaitForNodesInAgentPoolProfile(ctx, cs, &app, suffix)
			if err != nil {
				return &api.PluginError{Err: err, Step: api.PluginStepWaitForNodes}
			}
		}
	}

	// Wait for infrastructure services to be healthy
	if err := p.clusterUpgrader.HealthCheck(ctx, cs); err != nil {
		return err
	}

	if cs != nil {
		// setting VnetID based on VnetName
		cs.Properties.NetworkProfile.VnetID = resourceid.ResourceID(cs.Properties.AzProfile.SubscriptionID, cs.Properties.AzProfile.ResourceGroup, "Microsoft.Network/virtualNetworks", arm.VnetName)
	}

	// explicitly return nil if all went well
	return nil
}
