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
	pluginConfig api.PluginConfig
}

var _ Initializer = &simpleInitializer{}

// NewSimpleInitializer create a new Initilizer
func NewSimpleInitializer(entry *logrus.Entry, pluginConfig api.PluginConfig) Initializer {
	log.New(entry)
	return &simpleInitializer{pluginConfig: pluginConfig}
}

func (si *simpleInitializer) InitializeCluster(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	authorizer, err := azureclient.NewAuthorizerFromCtx(ctx)
	if err != nil {
		return err
	}

	accountClient := azureclient.NewAccountsClient(cs.Properties.AzProfile.SubscriptionID, authorizer, si.pluginConfig)
	keys, err := accountClient.ListKeys(ctx, cs.Properties.AzProfile.ResourceGroup, cs.Config.ConfigStorageAccount)
	if err != nil {
		return err
	}

	storageClient, err := azureclient.NewStorageClient(cs.Config.ConfigStorageAccount, *(*keys.Keys)[0].Value)
	if err != nil {
		return err
	}

	bsc := storageClient.GetBlobService()

	// etcd data container
	c := bsc.GetContainerReference("etcd")
	_, err = c.CreateIfNotExists(nil)
	if err != nil {
		return err
	}

	// cluster config container
	c = bsc.GetContainerReference("config")
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
