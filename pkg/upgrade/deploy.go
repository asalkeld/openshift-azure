package upgrade

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/initialize"
	"github.com/openshift/openshift-azure/pkg/log"
)

func Deploy(ctx context.Context, cs *api.OpenShiftManagedCluster, i initialize.Initializer, azuredeploy []byte, config api.PluginConfig) error {
	var t map[string]interface{}
	err := json.Unmarshal(azuredeploy, &t)
	if err != nil {
		return err
	}

	authorizer, err := azureclient.NewAuthorizerFromCtx(ctx)
	if err != nil {
		return err
	}

	deploymentClient := azureclient.NewDeploymentsClient(cs.Properties.AzProfile.SubscriptionID, authorizer, config)

	log.Info("applying arm template deployment")
	_, err = deploymentClient.CreateOrUpdate(ctx, cs.Properties.AzProfile.ResourceGroup, "azuredeploy", resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template: t,
			Mode:     resources.Incremental,
		},
	})
	if err != nil {
		return err
	}

	return i.InitializeCluster(ctx, cs)
}
