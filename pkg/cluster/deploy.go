package cluster

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/log"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
)

func (u *simpleUpgrader) DefaultDeploy(ctx context.Context, cs *api.OpenShiftManagedCluster, azuredeploy []byte) error {
	var azuretemplate map[string]interface{}
	err := json.Unmarshal(azuredeploy, &azuretemplate)
	if err != nil {
		return err
	}
	log.Info("applying arm template deployment")
	authorizer, err := azureclient.NewAuthorizerFromContext(ctx)
	if err != nil {
		return err
	}

	deployments := azureclient.NewDeploymentsClient(cs.Properties.AzProfile.SubscriptionID, authorizer, u.pluginConfig.AcceptLanguages)
	future, err := deployments.CreateOrUpdate(ctx, cs.Properties.AzProfile.ResourceGroup, "azuredeploy", resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template: azuretemplate,
			Mode:     resources.Incremental,
		},
	})
	if err != nil {
		return err
	}

	log.Info("waiting for arm template deployment to complete")
	return future.WaitForCompletionRef(ctx, deployments.Client())
}
