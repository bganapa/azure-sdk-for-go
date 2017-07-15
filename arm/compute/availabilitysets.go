package compute

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.2.0.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// AvailabilitySetsClient is the the Compute Management Client.
type AvailabilitySetsClient struct {
    ManagementClient
}
// NewAvailabilitySetsClient creates an instance of the AvailabilitySetsClient client.
func NewAvailabilitySetsClient(subscriptionID string) AvailabilitySetsClient {
        return NewAvailabilitySetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
        }

// NewAvailabilitySetsClientWithBaseURI creates an instance of the AvailabilitySetsClient client.
    func NewAvailabilitySetsClientWithBaseURI(baseURI string, subscriptionID string) AvailabilitySetsClient {
        return AvailabilitySetsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update an availability set.
//
// resourceGroupName is the name of the resource group. name is the name of the availability set. parameters is
// parameters supplied to the Create Availability Set operation.
func (client AvailabilitySetsClient) CreateOrUpdate(resourceGroupName string, name string, parameters AvailabilitySet) (result AvailabilitySet, err error) {
    req, err := client.CreateOrUpdatePreparer(resourceGroupName, name, parameters)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "CreateOrUpdate", nil , "Failure preparing request")
        return
    }

    resp, err := client.CreateOrUpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
    }

    result, err = client.CreateOrUpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "CreateOrUpdate", resp, "Failure responding to request")
    }

    return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AvailabilitySetsClient) CreateOrUpdatePreparer(resourceGroupName string, name string, parameters AvailabilitySet) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "name": autorest.Encode("path",name),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-03-30"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{name}",pathParameters),
                        autorest.WithJSON(parameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsClient) CreateOrUpdateResponder(resp *http.Response) (result AvailabilitySet, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Delete delete an availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is the name of the availability set.
func (client AvailabilitySetsClient) Delete(resourceGroupName string, availabilitySetName string) (result OperationStatusResponse, err error) {
    req, err := client.DeletePreparer(resourceGroupName, availabilitySetName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Delete", nil , "Failure preparing request")
        return
    }

    resp, err := client.DeleteSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Delete", resp, "Failure sending request")
        return
    }

    result, err = client.DeleteResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Delete", resp, "Failure responding to request")
    }

    return
}

// DeletePreparer prepares the Delete request.
func (client AvailabilitySetsClient) DeletePreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "availabilitySetName": autorest.Encode("path",availabilitySetName),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-03-30"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsClient) DeleteResponder(resp *http.Response) (result OperationStatusResponse, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Get retrieves information about an availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is the name of the availability set.
func (client AvailabilitySetsClient) Get(resourceGroupName string, availabilitySetName string) (result AvailabilitySet, err error) {
    req, err := client.GetPreparer(resourceGroupName, availabilitySetName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client AvailabilitySetsClient) GetPreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "availabilitySetName": autorest.Encode("path",availabilitySetName),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-03-30"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsClient) GetResponder(resp *http.Response) (result AvailabilitySet, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List lists all availability sets in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client AvailabilitySetsClient) List(resourceGroupName string) (result AvailabilitySetListResult, err error) {
    req, err := client.ListPreparer(resourceGroupName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client AvailabilitySetsClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-03-30"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsClient) ListResponder(resp *http.Response) (result AvailabilitySetListResult, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// ListAvailableSizes lists all available virtual machine sizes that can be used to create a new virtual machine in an
// existing availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is the name of the availability set.
func (client AvailabilitySetsClient) ListAvailableSizes(resourceGroupName string, availabilitySetName string) (result VirtualMachineSizeListResult, err error) {
    req, err := client.ListAvailableSizesPreparer(resourceGroupName, availabilitySetName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "ListAvailableSizes", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListAvailableSizesSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "ListAvailableSizes", resp, "Failure sending request")
        return
    }

    result, err = client.ListAvailableSizesResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "compute.AvailabilitySetsClient", "ListAvailableSizes", resp, "Failure responding to request")
    }

    return
}

// ListAvailableSizesPreparer prepares the ListAvailableSizes request.
func (client AvailabilitySetsClient) ListAvailableSizesPreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "availabilitySetName": autorest.Encode("path",availabilitySetName),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-03-30"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ListAvailableSizesSender sends the ListAvailableSizes request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsClient) ListAvailableSizesSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListAvailableSizesResponder handles the response to the ListAvailableSizes request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsClient) ListAvailableSizesResponder(resp *http.Response) (result VirtualMachineSizeListResult, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

