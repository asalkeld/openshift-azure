package initialize

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/log"
	"github.com/openshift/openshift-azure/pkg/util/fixtures"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_azureclient"
)

func TestInitializeCluster(t *testing.T) {
	gmc := gomock.NewController(t)
	accountsClient := mock_azureclient.NewMockAccountsClient(gmc)
	storageClient := mock_azureclient.NewMockStorageClient(gmc)
	si := &simpleInitializer{
		pluginConfig:   api.PluginConfig{},
		accountsClient: accountsClient,
		storageClient:  storageClient,
	}

	etcdCr := mock_azureclient.NewMockStorageContainerReference(gmc)
	storageClient.EXPECT().GetContainerReference("etcd").Return(etcdCr)
	etcdCr.EXPECT().CreateIfNotExists(nil).Return(true, nil)

	configCr := mock_azureclient.NewMockStorageContainerReference(gmc)
	storageClient.EXPECT().GetContainerReference("config").Return(configCr)
	configCr.EXPECT().CreateIfNotExists(nil).Return(true, nil)
	configBlob := mock_azureclient.NewMockStorageBlob(gmc)
	configCr.EXPECT().GetBlobReference("config").Return(configBlob)

	cs := fixtures.NewTestOpenShiftCluster()
	csj, err := json.Marshal(cs)
	if err != nil {
		t.Fatal(err)
	}
	var finalError error
	configBlob.EXPECT().CreateBlockBlobFromReader(bytes.NewReader(csj), nil).Return(finalError)

	log.New(logrus.NewEntry(logrus.New()))
	if err := si.InitializeCluster(context.Background(), cs); err != finalError {
		t.Errorf("simpleInitializer.InitializeCluster() error = %v", err)
	}
}
