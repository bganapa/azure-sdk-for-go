package resources

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
    "github.com/Azure/go-autorest/autorest/validation"
)

// DeploymentOperationsClient is the client for the DeploymentOperations methods of the Resources service.
type DeploymentOperationsClient struct {
    ManagementClient
}
// NewDeploymentOperationsClient creates an instance of the DeploymentOperationsClient client.
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
        return NewDeploymentOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
        }

// NewDeploymentOperationsClientWithBaseURI creates an instance of the DeploymentOperationsClient client.
    func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
        return DeploymentOperationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case insensitive. deploymentName is the name of the
// deployment. operationID is operation Id.
func (client DeploymentOperationsClient) Get(resourceGroupName string, deploymentName string, operationID string) (result DeploymentOperation, err error) {
    if err := validation.Validate([]validation.Validation{
    { TargetValue: resourceGroupName,
     Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
    	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
    	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
    return result, validation.NewErrorWithValidationError(err, "resources.DeploymentOperationsClient","Get")
}

    req, err := client.GetPreparer(resourceGroupName, deploymentName, operationID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client DeploymentOperationsClient) GetPreparer(resourceGroupName string, deploymentName string, operationID string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "deploymentName": autorest.Encode("path",deploymentName),
    "operationId": autorest.Encode("path",operationID),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-02-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsClient) GetResponder(resp *http.Response) (result DeploymentOperation, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List gets a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case insensitive. deploymentName is the name of the
// deployment. top is query parameters.
func (client DeploymentOperationsClient) List(resourceGroupName string, deploymentName string, top *int32) (result DeploymentOperationsListResult, err error) {
    if err := validation.Validate([]validation.Validation{
    { TargetValue: resourceGroupName,
     Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
    	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
    	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
    return result, validation.NewErrorWithValidationError(err, "resources.DeploymentOperationsClient","List")
}

    req, err := client.ListPreparer(resourceGroupName, deploymentName, top)
    if err != nil {
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client DeploymentOperationsClient) ListPreparer(resourceGroupName string, deploymentName string, top *int32) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "deploymentName": autorest.Encode("path",deploymentName),
    "resourceGroupName": autorest.Encode("path",resourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2016-02-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if top != nil {
        queryParameters["$top"] = autorest.Encode("query",*top)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsClient) ListResponder(resp *http.Response) (result DeploymentOperationsListResult, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// ListNextResults retrieves the next set of results, if any.
func (client DeploymentOperationsClient) ListNextResults(lastResults DeploymentOperationsListResult) (result DeploymentOperationsListResult, err error) {
    req, err := lastResults.DeploymentOperationsListResultPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", nil , "Failure preparing next results request")
    }
    if req == nil {
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", resp, "Failure sending next results request")
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", resp, "Failure responding to next results request")
    }

    return
}

// ListComplete gets all elements from the list without paging.
func (client DeploymentOperationsClient) ListComplete (resourceGroupName string, deploymentName string, top *int32, cancel <-chan struct{}) (<-chan DeploymentOperation, <-chan error) {
    resultChan := make(chan DeploymentOperation)
    errChan := make(chan error, 1)
    go func() {
        defer func() {
			close(resultChan)
			close(errChan)
		}()
        list, err := client.List(resourceGroupName, deploymentName, top)
        if err != nil {
            errChan <- err
            return
        }
        if list.Value != nil {
            for _, item := range *list.Value {
                select {
                case <- cancel:
                    return
                case resultChan <- item:
                     // Intentionally left blank
                }
            }
        }
         for list.NextLink != nil {
                    list, err = client.ListNextResults(list)
                    if err != nil {
                        errChan <- err
                        return
                    }
                    if list.Value != nil {
                        for _, item := range *list.Value {
                            select {
                            case <- cancel:
                                return
                            case resultChan <- item:
                                 // Intentionally left blank
                            }
                        }
                    }
                }
    }()
    return resultChan, errChan
}

