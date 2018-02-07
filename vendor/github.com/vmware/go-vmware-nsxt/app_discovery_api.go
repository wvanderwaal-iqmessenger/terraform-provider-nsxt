/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package nsxt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/go-vmware-nsxt/appdiscovery"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AppDiscoveryApiService service

/* AppDiscoveryApiService Adds a new app profile
Adds a new app profile
* @param ctx context.Context Authentication Context
@param appProfile
@return AppProfile*/
func (a *AppDiscoveryApiService) AddAppProfile(ctx context.Context, appProfile appdiscovery.AppProfile) (appdiscovery.AppProfile, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppProfile
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/app-profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &appProfile
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Cancel and delete the application discovery session
Cancel and delete the application discovery session
* @param ctx context.Context Authentication Context
@param sessionId
@return */
func (a *AppDiscoveryApiService) DeleteAppDiscoverySession(ctx context.Context, sessionId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* AppDiscoveryApiService Delete App Profile
Deletes the specified AppProfile.
* @param ctx context.Context Authentication Context
@param appProfileId
@param optional (nil or map[string]interface{}) with one or more of:
    @param "force" (bool) Force delete the resource even if it is being used somewhere
@return */
func (a *AppDiscoveryApiService) DeleteAppProfile(ctx context.Context, appProfileId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/app-profiles/{app-profile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app-profile-id"+"}", fmt.Sprintf("%v", appProfileId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["force"], "bool", "force"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["force"].(bool); localVarOk {
		localVarQueryParams.Add("force", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns the status of the application discovery session and other details
Returns the status of the application discovery session and other details
* @param ctx context.Context Authentication Context
@param sessionId
@return AppDiscoverySession*/
func (a *AppDiscoveryApiService) GetAppDiscoverySession(ctx context.Context, sessionId string) (appdiscovery.AppDiscoverySession, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppDiscoverySession
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns the details of the installed apps for the app profile ID in that session
Returns the details of the installed apps for the app profile ID in that session
* @param ctx context.Context Authentication Context
@param sessionId
@param optional (nil or map[string]interface{}) with one or more of:
    @param "appProfileId" (string)
    @param "cursor" (string) Opaque cursor to be used for getting next page of records (supplied by current result page)
    @param "includedFields" (string) Comma separated list of fields that should be included to result of query
    @param "pageSize" (int64) Maximum number of results to return in this page (server may return fewer)
    @param "sortAscending" (bool)
    @param "sortBy" (string) Field by which records are sorted
    @param "vmId" (string)
@return AppInfoListResult*/
func (a *AppDiscoveryApiService) GetAppDiscoverySessionInstalledApps(ctx context.Context, sessionId string, localVarOptionals map[string]interface{}) (appdiscovery.AppInfoListResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppInfoListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}/installed-apps"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["appProfileId"], "string", "appProfileId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["cursor"], "string", "cursor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includedFields"], "string", "includedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int64", "pageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortAscending"], "bool", "sortAscending"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortBy"], "string", "sortBy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["vmId"], "string", "vmId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["appProfileId"].(string); localVarOk {
		localVarQueryParams.Add("app_profile_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["cursor"].(string); localVarOk {
		localVarQueryParams.Add("cursor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includedFields"].(string); localVarOk {
		localVarQueryParams.Add("included_fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int64); localVarOk {
		localVarQueryParams.Add("page_size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortAscending"].(bool); localVarOk {
		localVarQueryParams.Add("sort_ascending", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortBy"].(string); localVarOk {
		localVarQueryParams.Add("sort_by", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["vmId"].(string); localVarOk {
		localVarQueryParams.Add("vm_id", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService vms in the ns-group of the application discovery session
Returns the vms in the ns-group of the application discovery session
* @param ctx context.Context Authentication Context
@param sessionId
@param nsGroupId
@return AppDiscoveryVmInfoListResult*/
func (a *AppDiscoveryApiService) GetAppDiscoverySessionNsGroupMembers(ctx context.Context, sessionId string, nsGroupId string) (appdiscovery.AppDiscoveryVmInfoListResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppDiscoveryVmInfoListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}/ns-groups/{ns-group-id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ns-group-id"+"}", fmt.Sprintf("%v", nsGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService ns-groups in this application discovery session
Returns the ns groups that was part of the application discovery session | while it was started
* @param ctx context.Context Authentication Context
@param sessionId
@return NsGroupMetaInfoListResult*/
func (a *AppDiscoveryApiService) GetAppDiscoverySessionNsGroups(ctx context.Context, sessionId string) (appdiscovery.NsGroupMetaInfoListResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.NsGroupMetaInfoListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}/ns-groups"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns the summary of the application discovery session
Returns the summary of the application discovery session
* @param ctx context.Context Authentication Context
@param sessionId
@return AppDiscoverySessionResultSummary*/
func (a *AppDiscoveryApiService) GetAppDiscoverySessionSummary(ctx context.Context, sessionId string) (appdiscovery.AppDiscoverySessionResultSummary, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppDiscoverySessionResultSummary
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions/{session-id}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"session-id"+"}", fmt.Sprintf("%v", sessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns the list of the application discovery sessions available
Returns the list of the application discovery sessions available
* @param ctx context.Context Authentication Context
@param optional (nil or map[string]interface{}) with one or more of:
    @param "cursor" (string) Opaque cursor to be used for getting next page of records (supplied by current result page)
    @param "groupId" (string) NSGroup id, helps user query sessions related to one nsgroup
    @param "includedFields" (string) Comma separated list of fields that should be included to result of query
    @param "pageSize" (int64) Maximum number of results to return in this page (server may return fewer)
    @param "sortAscending" (bool)
    @param "sortBy" (string) Field by which records are sorted
    @param "status" (string) Session Status, e.g. get all running sessions
@return AppDiscoverySessionsListResult*/
func (a *AppDiscoveryApiService) GetAppDiscoverySessions(ctx context.Context, localVarOptionals map[string]interface{}) (appdiscovery.AppDiscoverySessionsListResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppDiscoverySessionsListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["cursor"], "string", "cursor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["groupId"], "string", "groupId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includedFields"], "string", "includedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int64", "pageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortAscending"], "bool", "sortAscending"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortBy"], "string", "sortBy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["status"], "string", "status"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["cursor"].(string); localVarOk {
		localVarQueryParams.Add("cursor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["groupId"].(string); localVarOk {
		localVarQueryParams.Add("group_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includedFields"].(string); localVarOk {
		localVarQueryParams.Add("included_fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int64); localVarOk {
		localVarQueryParams.Add("page_size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortAscending"].(bool); localVarOk {
		localVarQueryParams.Add("sort_ascending", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortBy"].(string); localVarOk {
		localVarQueryParams.Add("sort_by", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["status"].(string); localVarOk {
		localVarQueryParams.Add("status", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns detail of the app profile
Returns detail of the app profile
* @param ctx context.Context Authentication Context
@param appProfileId
@return AppProfile*/
func (a *AppDiscoveryApiService) GetAppProfileDetails(ctx context.Context, appProfileId string) (appdiscovery.AppProfile, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppProfile
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/app-profiles/{app-profile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app-profile-id"+"}", fmt.Sprintf("%v", appProfileId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Returns list of app profile IDs created
Returns list of app profile IDs created
* @param ctx context.Context Authentication Context
@return AppProfileListResult*/
func (a *AppDiscoveryApiService) GetAppProfiles(ctx context.Context) (appdiscovery.AppProfileListResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppProfileListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/app-profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Starts the discovery of application discovery session
Starts the discovery of application discovery session
* @param ctx context.Context Authentication Context
@param startAppDiscoverySessionParameters
@return AppDiscoverySession*/
func (a *AppDiscoveryApiService) StartAppDiscoverySession(ctx context.Context, startAppDiscoverySessionParameters appdiscovery.StartAppDiscoverySessionParameters) (appdiscovery.AppDiscoverySession, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppDiscoverySession
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &startAppDiscoverySessionParameters
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* AppDiscoveryApiService Update AppProfile
Update AppProfile
* @param ctx context.Context Authentication Context
@param appProfileId
@param appProfile
@return AppProfile*/
func (a *AppDiscoveryApiService) UpdateAppProfile(ctx context.Context, appProfileId string, appProfile appdiscovery.AppProfile) (appdiscovery.AppProfile, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     appdiscovery.AppProfile
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/app-discovery/app-profiles/{app-profile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app-profile-id"+"}", fmt.Sprintf("%v", appProfileId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &appProfile
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}