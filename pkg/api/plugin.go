// Package api defines the external API for the plugin.
package api

import (
	"context"
)

// ContextKey is a type for context property bag payload keys
type ContextKey string

const (
	ContextKeyClientID     ContextKey = "ClientID"
	ContextKeyClientSecret ContextKey = "ClientSecret"
	ContextKeyTenantID     ContextKey = "TenantID"
)

// Deployer interface describes the functions to call (and in the order below) to
// deploy a cluster for the first time.
type Deployer interface {
	DefaultDeployment(ctx context.Context, cs *OpenShiftManagedCluster, azuredeploy []byte) error
	Initialize(ctx context.Context, cs *OpenShiftManagedCluster) error
	WaitForCompletion(ctx context.Context, cs *OpenShiftManagedCluster) error
}

// Updater interface describes the functions to call (and in the order below) to
// update a cluster.
type Updater interface {
	GetNodesPreUpdate(ctx context.Context, cs *OpenShiftManagedCluster) (map[string]struct{}, error)
	DefaultDeployment(ctx context.Context, cs *OpenShiftManagedCluster, azuredeploy []byte) error
	Initialize(ctx context.Context, cs *OpenShiftManagedCluster) error
	WaitForCompletion(ctx context.Context, cs *OpenShiftManagedCluster, nodes map[string]struct{}) error
	UpdateInPlace(ctx context.Context, cs *OpenShiftManagedCluster) error
}

// PluginConfig is passed into NewPlugin
type PluginConfig struct {
	SyncImage       string
	AcceptLanguages []string
}

// Plugin is the main interface to openshift-azure
type Plugin interface {
	// MergeConfig merges new and old config so that no unnecessary config
	// is going to get regenerated during generation. It also handles merging
	// partial user requests to allow reusing the same validation code during
	// upgrades. This method should be the first one called by the RP, before
	// validation and generation.
	MergeConfig(ctx context.Context, new, old *OpenShiftManagedCluster)

	// Validate exists (a) to be able to place validation logic in a
	// single place in the event of multiple external API versions, and (b) to
	// be able to compare a new API manifest against a pre-existing API manifest
	// (for update, upgrade, etc.)
	// externalOnly indicates that fields set by the RP (FQDN and routerProfile.FQDN)
	// should be excluded.
	Validate(ctx context.Context, new, old *OpenShiftManagedCluster, externalOnly bool) []error

	// GenerateConfig ensures all the necessary in-cluster config is generated
	// for an Openshift cluster.
	GenerateConfig(ctx context.Context, cs *OpenShiftManagedCluster) error

	// GenerateARM creates an ARM based off of the config.
	GenerateARM(ctx context.Context, cs *OpenShiftManagedCluster, isUpdate bool) ([]byte, error)

	Deployer() Deployer
	Updater() Updater

	// CreateOrUpdate either deploys or runs the update depending on the isUpdate argument
	// this will call the deployer.
	CreateOrUpdate(ctx context.Context, cs *OpenShiftManagedCluster, azuredeploy []byte, isUpdate bool) error
}
