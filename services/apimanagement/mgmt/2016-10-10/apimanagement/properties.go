package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PropertiesClient is the apiManagement Client
type PropertiesClient struct {
	BaseClient
}

// NewPropertiesClient creates an instance of the PropertiesClient client.
func NewPropertiesClient(subscriptionID string) PropertiesClient {
	return NewPropertiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPropertiesClientWithBaseURI creates an instance of the PropertiesClient client.
func NewPropertiesClientWithBaseURI(baseURI string, subscriptionID string) PropertiesClient {
	return PropertiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByService lists a collection of properties defined within a service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// filter - | Field | Supported operators    | Supported functions                                   |
// |-------|------------------------|-------------------------------------------------------|
// | tags  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith, any, all |
// | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith           |
// top - number of records to return.
// skip - number of records to skip.
func (client PropertiesClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result PropertyCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PropertiesClient.ListByService")
		defer func() {
			sc := -1
			if result.pc.Response.Response != nil {
				sc = result.pc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.PropertiesClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.pc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.pc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client PropertiesClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client PropertiesClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client PropertiesClient) ListByServiceResponder(resp *http.Response) (result PropertyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client PropertiesClient) listByServiceNextResults(ctx context.Context, lastResults PropertyCollection) (result PropertyCollection, err error) {
	req, err := lastResults.propertyCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PropertiesClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PropertiesClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result PropertyCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PropertiesClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter, top, skip)
	return
}
