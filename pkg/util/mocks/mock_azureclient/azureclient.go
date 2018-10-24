// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient (interfaces: Client,VirtualMachineScaleSetsClient,VirtualMachineScaleSetVMsClient,VirtualMachineScaleSetExtensionsClient,ApplicationsClient,MarketPlaceAgreementsClient,DeploymentsClient,AccountsClient)

// Package mock_azureclient is a generated GoMock package.
package mock_azureclient

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	marketplaceordering "github.com/Azure/azure-sdk-for-go/services/marketplaceordering/mgmt/2015-06-01/marketplaceordering"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	managedapplications "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/managedapplications"
	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockClient)(nil).Client))
}

// MockVirtualMachineScaleSetsClient is a mock of VirtualMachineScaleSetsClient interface
type MockVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetsClientMockRecorder
}

// MockVirtualMachineScaleSetsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetsClient
type MockVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetsClient
}

// NewMockVirtualMachineScaleSetsClient creates a new mock instance
func NewMockVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetsClient {
	mock := &MockVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetsClient) EXPECT() *MockVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualMachineScaleSetsClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Client))
}

// List mocks base method
func (m *MockVirtualMachineScaleSetsClient) List(arg0 context.Context, arg1 string) (compute.VirtualMachineScaleSetListResultPage, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockVirtualMachineScaleSetsClient) Update(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetUpdate) (compute.VirtualMachineScaleSetsUpdateFuture, error) {
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetsUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Update), arg0, arg1, arg2, arg3)
}

// UpdateInstances mocks base method
func (m *MockVirtualMachineScaleSetsClient) UpdateInstances(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (compute.VirtualMachineScaleSetsUpdateInstancesFuture, error) {
	ret := m.ctrl.Call(m, "UpdateInstances", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetsUpdateInstancesFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInstances indicates an expected call of UpdateInstances
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) UpdateInstances(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstances", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).UpdateInstances), arg0, arg1, arg2, arg3)
}

// MockVirtualMachineScaleSetVMsClient is a mock of VirtualMachineScaleSetVMsClient interface
type MockVirtualMachineScaleSetVMsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetVMsClientMockRecorder
}

// MockVirtualMachineScaleSetVMsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetVMsClient
type MockVirtualMachineScaleSetVMsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetVMsClient
}

// NewMockVirtualMachineScaleSetVMsClient creates a new mock instance
func NewMockVirtualMachineScaleSetVMsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetVMsClient {
	mock := &MockVirtualMachineScaleSetVMsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetVMsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetVMsClient) EXPECT() *MockVirtualMachineScaleSetVMsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Client))
}

// Deallocate mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Deallocate(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsDeallocateFuture, error) {
	ret := m.ctrl.Call(m, "Deallocate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsDeallocateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deallocate indicates an expected call of Deallocate
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Deallocate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deallocate", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Deallocate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Delete(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsDeleteFuture, error) {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Delete), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) List(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (compute.VirtualMachineScaleSetVMListResultPage, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Reimage mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Reimage(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsReimageFuture, error) {
	ret := m.ctrl.Call(m, "Reimage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsReimageFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reimage indicates an expected call of Reimage
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Reimage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reimage", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Reimage), arg0, arg1, arg2, arg3)
}

// Restart mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Restart(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsRestartFuture, error) {
	ret := m.ctrl.Call(m, "Restart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsRestartFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Restart indicates an expected call of Restart
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Restart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Restart), arg0, arg1, arg2, arg3)
}

// Start mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Start(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsStartFuture, error) {
	ret := m.ctrl.Call(m, "Start", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsStartFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Start(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Start), arg0, arg1, arg2, arg3)
}

// MockVirtualMachineScaleSetExtensionsClient is a mock of VirtualMachineScaleSetExtensionsClient interface
type MockVirtualMachineScaleSetExtensionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetExtensionsClientMockRecorder
}

// MockVirtualMachineScaleSetExtensionsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetExtensionsClient
type MockVirtualMachineScaleSetExtensionsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetExtensionsClient
}

// NewMockVirtualMachineScaleSetExtensionsClient creates a new mock instance
func NewMockVirtualMachineScaleSetExtensionsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetExtensionsClient {
	mock := &MockVirtualMachineScaleSetExtensionsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetExtensionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetExtensionsClient) EXPECT() *MockVirtualMachineScaleSetExtensionsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 compute.VirtualMachineScaleSetExtension) (compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture, error) {
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (compute.VirtualMachineScaleSetExtension, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// List mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) List(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineScaleSetExtensionListResultPage, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).List), arg0, arg1, arg2)
}

// MockApplicationsClient is a mock of ApplicationsClient interface
type MockApplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationsClientMockRecorder
}

// MockApplicationsClientMockRecorder is the mock recorder for MockApplicationsClient
type MockApplicationsClientMockRecorder struct {
	mock *MockApplicationsClient
}

// NewMockApplicationsClient creates a new mock instance
func NewMockApplicationsClient(ctrl *gomock.Controller) *MockApplicationsClient {
	mock := &MockApplicationsClient{ctrl: ctrl}
	mock.recorder = &MockApplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationsClient) EXPECT() *MockApplicationsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockApplicationsClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockApplicationsClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockApplicationsClient)(nil).Client))
}

// Get mocks base method
func (m *MockApplicationsClient) Get(arg0 context.Context, arg1, arg2 string) (managedapplications.Application, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockApplicationsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationsClient)(nil).Get), arg0, arg1, arg2)
}

// GetByID mocks base method
func (m *MockApplicationsClient) GetByID(arg0 context.Context, arg1 string) (managedapplications.Application, error) {
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockApplicationsClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockApplicationsClient)(nil).GetByID), arg0, arg1)
}

// ListByResourceGroup mocks base method
func (m *MockApplicationsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (managedapplications.ApplicationListResultPage, error) {
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.ApplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockApplicationsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockApplicationsClient)(nil).ListByResourceGroup), arg0, arg1)
}

// MockMarketPlaceAgreementsClient is a mock of MarketPlaceAgreementsClient interface
type MockMarketPlaceAgreementsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMarketPlaceAgreementsClientMockRecorder
}

// MockMarketPlaceAgreementsClientMockRecorder is the mock recorder for MockMarketPlaceAgreementsClient
type MockMarketPlaceAgreementsClientMockRecorder struct {
	mock *MockMarketPlaceAgreementsClient
}

// NewMockMarketPlaceAgreementsClient creates a new mock instance
func NewMockMarketPlaceAgreementsClient(ctrl *gomock.Controller) *MockMarketPlaceAgreementsClient {
	mock := &MockMarketPlaceAgreementsClient{ctrl: ctrl}
	mock.recorder = &MockMarketPlaceAgreementsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketPlaceAgreementsClient) EXPECT() *MockMarketPlaceAgreementsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMarketPlaceAgreementsClient) Create(arg0 context.Context, arg1, arg2, arg3 string, arg4 marketplaceordering.AgreementTerms) (marketplaceordering.AgreementTerms, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Create(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Create), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockMarketPlaceAgreementsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (marketplaceordering.AgreementTerms, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockDeploymentsClient is a mock of DeploymentsClient interface
type MockDeploymentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentsClientMockRecorder
}

// MockDeploymentsClientMockRecorder is the mock recorder for MockDeploymentsClient
type MockDeploymentsClientMockRecorder struct {
	mock *MockDeploymentsClient
}

// NewMockDeploymentsClient creates a new mock instance
func NewMockDeploymentsClient(ctrl *gomock.Controller) *MockDeploymentsClient {
	mock := &MockDeploymentsClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeploymentsClient) EXPECT() *MockDeploymentsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockDeploymentsClient) Client() autorest.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockDeploymentsClientMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDeploymentsClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockDeploymentsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 resources.Deployment) (resources.DeploymentsCreateOrUpdateFuture, error) {
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(resources.DeploymentsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockDeploymentsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockDeploymentsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// MockAccountsClient is a mock of AccountsClient interface
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// ListByResourceGroup mocks base method
func (m *MockAccountsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (storage.AccountListResult, error) {
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(storage.AccountListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockAccountsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockAccountsClient)(nil).ListByResourceGroup), arg0, arg1)
}

// ListKeys mocks base method
func (m *MockAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string) (storage.AccountListKeysResult, error) {
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(storage.AccountListKeysResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockAccountsClient)(nil).ListKeys), arg0, arg1, arg2)
}