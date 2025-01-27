//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewMigrationRecoveryPointsClient() *MigrationRecoveryPointsClient {
	subClient, _ := NewMigrationRecoveryPointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRecoveryPointsClient() *RecoveryPointsClient {
	subClient, _ := NewRecoveryPointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationAlertSettingsClient() *ReplicationAlertSettingsClient {
	subClient, _ := NewReplicationAlertSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationAppliancesClient() *ReplicationAppliancesClient {
	subClient, _ := NewReplicationAppliancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationEligibilityResultsClient() *ReplicationEligibilityResultsClient {
	subClient, _ := NewReplicationEligibilityResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationEventsClient() *ReplicationEventsClient {
	subClient, _ := NewReplicationEventsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationFabricsClient() *ReplicationFabricsClient {
	subClient, _ := NewReplicationFabricsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationJobsClient() *ReplicationJobsClient {
	subClient, _ := NewReplicationJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationLogicalNetworksClient() *ReplicationLogicalNetworksClient {
	subClient, _ := NewReplicationLogicalNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationMigrationItemsClient() *ReplicationMigrationItemsClient {
	subClient, _ := NewReplicationMigrationItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationNetworkMappingsClient() *ReplicationNetworkMappingsClient {
	subClient, _ := NewReplicationNetworkMappingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationNetworksClient() *ReplicationNetworksClient {
	subClient, _ := NewReplicationNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationPoliciesClient() *ReplicationPoliciesClient {
	subClient, _ := NewReplicationPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationProtectableItemsClient() *ReplicationProtectableItemsClient {
	subClient, _ := NewReplicationProtectableItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationProtectedItemsClient() *ReplicationProtectedItemsClient {
	subClient, _ := NewReplicationProtectedItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationProtectionContainerMappingsClient() *ReplicationProtectionContainerMappingsClient {
	subClient, _ := NewReplicationProtectionContainerMappingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationProtectionContainersClient() *ReplicationProtectionContainersClient {
	subClient, _ := NewReplicationProtectionContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationProtectionIntentsClient() *ReplicationProtectionIntentsClient {
	subClient, _ := NewReplicationProtectionIntentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationRecoveryPlansClient() *ReplicationRecoveryPlansClient {
	subClient, _ := NewReplicationRecoveryPlansClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationRecoveryServicesProvidersClient() *ReplicationRecoveryServicesProvidersClient {
	subClient, _ := NewReplicationRecoveryServicesProvidersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationStorageClassificationMappingsClient() *ReplicationStorageClassificationMappingsClient {
	subClient, _ := NewReplicationStorageClassificationMappingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationStorageClassificationsClient() *ReplicationStorageClassificationsClient {
	subClient, _ := NewReplicationStorageClassificationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationVaultHealthClient() *ReplicationVaultHealthClient {
	subClient, _ := NewReplicationVaultHealthClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationVaultSettingClient() *ReplicationVaultSettingClient {
	subClient, _ := NewReplicationVaultSettingClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationvCentersClient() *ReplicationvCentersClient {
	subClient, _ := NewReplicationvCentersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSupportedOperatingSystemsClient() *SupportedOperatingSystemsClient {
	subClient, _ := NewSupportedOperatingSystemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTargetComputeSizesClient() *TargetComputeSizesClient {
	subClient, _ := NewTargetComputeSizesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
