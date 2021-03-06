// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package appconfiguration

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/appconfiguration/mgmt/2019-02-01-preview/appconfiguration"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type APIKey = original.APIKey
type APIKeyListResult = original.APIKeyListResult
type APIKeyListResultIterator = original.APIKeyListResultIterator
type APIKeyListResultPage = original.APIKeyListResultPage
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type ConfigurationStore = original.ConfigurationStore
type ConfigurationStoreListResult = original.ConfigurationStoreListResult
type ConfigurationStoreListResultIterator = original.ConfigurationStoreListResultIterator
type ConfigurationStoreListResultPage = original.ConfigurationStoreListResultPage
type ConfigurationStoreProperties = original.ConfigurationStoreProperties
type ConfigurationStoreUpdateParameters = original.ConfigurationStoreUpdateParameters
type ConfigurationStoresClient = original.ConfigurationStoresClient
type ConfigurationStoresCreateFuture = original.ConfigurationStoresCreateFuture
type ConfigurationStoresDeleteFuture = original.ConfigurationStoresDeleteFuture
type ConfigurationStoresUpdateFuture = original.ConfigurationStoresUpdateFuture
type Error = original.Error
type KeyValue = original.KeyValue
type ListKeyValueParameters = original.ListKeyValueParameters
type NameAvailabilityStatus = original.NameAvailabilityStatus
type OperationDefinition = original.OperationDefinition
type OperationDefinitionDisplay = original.OperationDefinitionDisplay
type OperationDefinitionListResult = original.OperationDefinitionListResult
type OperationDefinitionListResultIterator = original.OperationDefinitionListResultIterator
type OperationDefinitionListResultPage = original.OperationDefinitionListResultPage
type OperationsClient = original.OperationsClient
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIKeyListResultIterator(page APIKeyListResultPage) APIKeyListResultIterator {
	return original.NewAPIKeyListResultIterator(page)
}
func NewAPIKeyListResultPage(getNextPage func(context.Context, APIKeyListResult) (APIKeyListResult, error)) APIKeyListResultPage {
	return original.NewAPIKeyListResultPage(getNextPage)
}
func NewConfigurationStoreListResultIterator(page ConfigurationStoreListResultPage) ConfigurationStoreListResultIterator {
	return original.NewConfigurationStoreListResultIterator(page)
}
func NewConfigurationStoreListResultPage(getNextPage func(context.Context, ConfigurationStoreListResult) (ConfigurationStoreListResult, error)) ConfigurationStoreListResultPage {
	return original.NewConfigurationStoreListResultPage(getNextPage)
}
func NewConfigurationStoresClient(subscriptionID string) ConfigurationStoresClient {
	return original.NewConfigurationStoresClient(subscriptionID)
}
func NewConfigurationStoresClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationStoresClient {
	return original.NewConfigurationStoresClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationDefinitionListResultIterator(page OperationDefinitionListResultPage) OperationDefinitionListResultIterator {
	return original.NewOperationDefinitionListResultIterator(page)
}
func NewOperationDefinitionListResultPage(getNextPage func(context.Context, OperationDefinitionListResult) (OperationDefinitionListResult, error)) OperationDefinitionListResultPage {
	return original.NewOperationDefinitionListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
