package automation

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
	"github.com/satori/go.uuid"
	"net/http"
)

// SourceControlSyncJobClient is the automation Client
type SourceControlSyncJobClient struct {
	BaseClient
}

// NewSourceControlSyncJobClient creates an instance of the SourceControlSyncJobClient client.
func NewSourceControlSyncJobClient(subscriptionID string) SourceControlSyncJobClient {
	return NewSourceControlSyncJobClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSourceControlSyncJobClientWithBaseURI creates an instance of the SourceControlSyncJobClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSourceControlSyncJobClientWithBaseURI(baseURI string, subscriptionID string) SourceControlSyncJobClient {
	return SourceControlSyncJobClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates the sync job for a source control.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// sourceControlName - the source control name.
// sourceControlSyncJobID - the source control sync job id.
// parameters - the parameters supplied to the create source control sync job operation.
func (client SourceControlSyncJobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, parameters SourceControlSyncJobCreateParameters) (result SourceControlSyncJob, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SourceControlSyncJobCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.SourceControlSyncJobCreateProperties.CommitID", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.SourceControlSyncJobCreateProperties.CommitID", Name: validation.MinLength, Rule: 0, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlSyncJobClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SourceControlSyncJobClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID, parameters SourceControlSyncJobCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":  autorest.Encode("path", automationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sourceControlName":      autorest.Encode("path", sourceControlName),
		"sourceControlSyncJobId": autorest.Encode("path", sourceControlSyncJobID),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlSyncJobClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SourceControlSyncJobClient) CreateResponder(resp *http.Response) (result SourceControlSyncJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve the source control sync job identified by job id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// sourceControlName - the source control name.
// sourceControlSyncJobID - the source control sync job id.
func (client SourceControlSyncJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID) (result SourceControlSyncJobByID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlSyncJobClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SourceControlSyncJobClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":  autorest.Encode("path", automationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sourceControlName":      autorest.Encode("path", sourceControlName),
		"sourceControlSyncJobId": autorest.Encode("path", sourceControlSyncJobID),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlSyncJobClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SourceControlSyncJobClient) GetResponder(resp *http.Response) (result SourceControlSyncJobByID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of source control sync jobs.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// sourceControlName - the source control name.
// filter - the filter to apply on the operation.
func (client SourceControlSyncJobClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, filter string) (result SourceControlSyncJobListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.scsjlr.Response.Response != nil {
				sc = result.scsjlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlSyncJobClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, sourceControlName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.scsjlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.scsjlr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}
	if result.scsjlr.hasNextLink() && result.scsjlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client SourceControlSyncJobClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sourceControlName":     autorest.Encode("path", sourceControlName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlSyncJobClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client SourceControlSyncJobClient) ListByAutomationAccountResponder(resp *http.Response) (result SourceControlSyncJobListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client SourceControlSyncJobClient) listByAutomationAccountNextResults(ctx context.Context, lastResults SourceControlSyncJobListResult) (result SourceControlSyncJobListResult, err error) {
	req, err := lastResults.sourceControlSyncJobListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlSyncJobClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client SourceControlSyncJobClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, filter string) (result SourceControlSyncJobListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlSyncJobClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, sourceControlName, filter)
	return
}
