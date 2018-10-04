package initialize

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/log"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
)

type Initializer interface {
	InitializeCluster(ctx context.Context, cs *api.OpenShiftManagedCluster) error
}

type simpleInitializer struct {
	pluginConfig   api.PluginConfig
	accountsClient azureclient.AccountsClient
	storageClient  azureclient.StorageClient
}

var _ Initializer = &simpleInitializer{}

// NewSimpleInitializer create a new Initilizer
func NewSimpleInitializer(entry *logrus.Entry, pluginConfig api.PluginConfig) Initializer {
	log.New(entry)
	return &simpleInitializer{pluginConfig: pluginConfig}
}

func (si *simpleInitializer) InitializeCluster(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	var err error
	if si.accountsClient == nil {
		authorizer, err := azureclient.NewAuthorizerFromCtx(ctx)
		if err != nil {
			return err
		}

		si.accountsClient = azureclient.NewAccountsClient(cs.Properties.AzProfile.SubscriptionID, authorizer, si.pluginConfig)
	}
	if si.storageClient == nil {
		keys, err := si.accountsClient.ListKeys(ctx, cs.Properties.AzProfile.ResourceGroup, cs.Config.ConfigStorageAccount)
		if err != nil {
			return err
		}

		si.storageClient, err = azureclient.NewStorageClient(cs.Config.ConfigStorageAccount, *(*keys.Keys)[0].Value)
		if err != nil {
			return err
		}
	}
	// etcd data container
	c := si.storageClient.GetContainerReference("etcd")
	_, err = c.CreateIfNotExists(nil)
	if err != nil {
		return err
	}

	// cluster config container
	c = si.storageClient.GetContainerReference("config")
	_, err = c.CreateIfNotExists(nil)
	if err != nil {
		return err
	}

	b := c.GetBlobReference("config")

	csj, err := json.Marshal(cs)
	if err != nil {
		return err
	}

	return b.CreateBlockBlobFromReader(bytes.NewReader(csj), nil)
}
