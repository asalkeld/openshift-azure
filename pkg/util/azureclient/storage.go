package azureclient

import (
	"context"
	"errors"
	"io"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"

	"github.com/openshift/openshift-azure/pkg/api"
)

// NewAccountsClient return AccountStorageClient implementation
func NewAccountsClient(subscriptionID string, authorizer autorest.Authorizer, pluginConfig api.PluginConfig) AccountsClient {
	client := storage.NewAccountsClient(subscriptionID)
	client.Authorizer = authorizer
	client.RequestInspector = addAcceptLanguages(pluginConfig.AcceptLanguages)
	return &azAccountsClient{
		client: client,
	}
}

// ListKeys returns all keys of the storage account
func (az azAccountsClient) ListKeys(ctx context.Context, resourceGroup, accountName string) (result storage.AccountListKeysResult, err error) {
	return az.client.ListKeys(ctx, resourceGroup, accountName)
}

// NewStorageClient return StorageClient implementation
func NewStorageClient(accountName, key string) (StorageClient, error) {
	// get az storage client
	azs, err := azstorage.NewClient(accountName, key, azstorage.DefaultBaseURL, azstorage.DefaultAPIVersion, true)
	if err != nil {
		return nil, err
	}
	// get blob service client
	bs := azs.GetBlobService()
	return &azStorageClient{
		azs: azs,
		bs:  bs,
	}, nil
}

func (az azStorageClient) GetBlobService() azstorage.BlobStorageClient {
	return az.azs.GetBlobService()
}

// NewStorageContainerReference purely for unit tests
func NewStorageContainerReference(name string) StorageContainerReference {
	return azStorageContainerReference{}
}

func (az azStorageClient) GetContainerReference(name string) StorageContainerReference {
	return azStorageContainerReference{container: az.bs.GetContainerReference(name)}
}

func (az azStorageContainerReference) CreateIfNotExists(options *azstorage.CreateContainerOptions) (bool, error) {
	return az.container.CreateIfNotExists(options)
}

func (az azStorageContainerReference) GetBlobReference(name string) StorageBlob {
	return azStorageBlob{blob: az.container.GetBlobReference(name)}
}

// NewStorageBlob purely for unit tests
func NewStorageBlob() StorageBlob {
	return azStorageBlob{}
}

func (az azStorageBlob) CreateBlockBlobFromReader(blob io.Reader, options *azstorage.PutBlobOptions) error {
	return az.blob.CreateBlockBlobFromReader(blob, options)
}

func (az azStorageBlob) Get(options *azstorage.GetBlobOptions) (io.ReadCloser, error) {
	return az.blob.Get(options)
}

func (az azAccountsClient) GetStorageAccount(ctx context.Context, resourceGroup, typeTag string) (map[string]string, error) {
	accts, err := az.client.ListByResourceGroup(context.Background(), resourceGroup)
	if err != nil {
		return nil, err
	}
	var acct storage.Account
	var found bool
	for _, acct = range *accts.Value {
		found = acct.Tags["type"] != nil && *acct.Tags["type"] == typeTag
		if found {
			break
		}
	}
	if !found {
		return nil, errors.New("storage account not found")
	}
	keys, err := az.client.ListKeys(context.Background(), resourceGroup, *acct.Name)
	if err != nil {
		return nil, err
	}
	// TODO: Allow choosing between the two storage account keys to
	//// enable more convenient key rotation.
	result := map[string]string{
		"name": *acct.Name,
		"key":  *(*keys.Keys)[0].Value,
	}
	return result, nil
}

func (az azAccountsClient) GetStorageAccountKey(ctx context.Context, resourceGroup, accountName string) (string, error) {
	response, err := az.client.ListKeys(context.Background(), resourceGroup, accountName)
	if err != nil {
		return "", err
	}
	// TODO: Allow choosing between the two storage account keys to
	// enable more convenient key rotation.
	return *(((*response.Keys)[0]).Value), nil
}

func (az azAccountsClient) ListByResourceGroup(ctx context.Context, resourceGroup string) (storage.AccountListResult, error) {
	return az.client.ListByResourceGroup(ctx, resourceGroup)
}
