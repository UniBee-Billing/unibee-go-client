/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MetricService Metric service
type MetricService service

type MetricMetricDeletePostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricDeleteReq *UnibeeApiMerchantMetricDeleteReq
}

func (r MetricMetricDeletePostRequest) UnibeeApiMerchantMetricDeleteReq(unibeeApiMerchantMetricDeleteReq UnibeeApiMerchantMetricDeleteReq) MetricMetricDeletePostRequest {
	r.unibeeApiMerchantMetricDeleteReq = &unibeeApiMerchantMetricDeleteReq
	return r
}

func (r MetricMetricDeletePostRequest) Execute() (*MerchantMetricDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricDeletePostExecute(r)
}

/*
MetricDeletePost Delete Merchant Metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricDeletePostRequest
*/
func (a *MetricService) MetricDeletePost(ctx context.Context) MetricMetricDeletePostRequest {
	return MetricMetricDeletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricDeletePost200Response
func (a *MetricService) MetricDeletePostExecute(r MetricMetricDeletePostRequest) (*MerchantMetricDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricDeletePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricDeleteReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricDeleteReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricDeleteReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricDetailPostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricDetailReq *UnibeeApiMerchantMetricDetailReq
}

func (r MetricMetricDetailPostRequest) UnibeeApiMerchantMetricDetailReq(unibeeApiMerchantMetricDetailReq UnibeeApiMerchantMetricDetailReq) MetricMetricDetailPostRequest {
	r.unibeeApiMerchantMetricDetailReq = &unibeeApiMerchantMetricDetailReq
	return r
}

func (r MetricMetricDetailPostRequest) Execute() (*MerchantMetricDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricDetailPostExecute(r)
}

/*
MetricDetailPost Merchant Metric Detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricDetailPostRequest
*/
func (a *MetricService) MetricDetailPost(ctx context.Context) MetricMetricDetailPostRequest {
	return MetricMetricDetailPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricDeletePost200Response
func (a *MetricService) MetricDetailPostExecute(r MetricMetricDetailPostRequest) (*MerchantMetricDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricDetailPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricDetailReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricDetailReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricDetailReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricEditPostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricEditReq *UnibeeApiMerchantMetricEditReq
}

func (r MetricMetricEditPostRequest) UnibeeApiMerchantMetricEditReq(unibeeApiMerchantMetricEditReq UnibeeApiMerchantMetricEditReq) MetricMetricEditPostRequest {
	r.unibeeApiMerchantMetricEditReq = &unibeeApiMerchantMetricEditReq
	return r
}

func (r MetricMetricEditPostRequest) Execute() (*MerchantMetricDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricEditPostExecute(r)
}

/*
MetricEditPost Edit Merchant Metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricEditPostRequest
*/
func (a *MetricService) MetricEditPost(ctx context.Context) MetricMetricEditPostRequest {
	return MetricMetricEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricDeletePost200Response
func (a *MetricService) MetricEditPostExecute(r MetricMetricEditPostRequest) (*MerchantMetricDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricEditReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricEditReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricEditReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricListGetRequest struct {
	ctx context.Context
	ApiService *MetricService
}

func (r MetricMetricListGetRequest) Execute() (*MerchantMetricListGet200Response, *http.Response, error) {
	return r.ApiService.MetricListGetExecute(r)
}

/*
MetricListGet Get Merchant Metric list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricListGetRequest
*/
func (a *MetricService) MetricListGet(ctx context.Context) MetricMetricListGetRequest {
	return MetricMetricListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricListGet200Response
func (a *MetricService) MetricListGetExecute(r MetricMetricListGetRequest) (*MerchantMetricListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricNewPostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricNewReq *UnibeeApiMerchantMetricNewReq
}

func (r MetricMetricNewPostRequest) UnibeeApiMerchantMetricNewReq(unibeeApiMerchantMetricNewReq UnibeeApiMerchantMetricNewReq) MetricMetricNewPostRequest {
	r.unibeeApiMerchantMetricNewReq = &unibeeApiMerchantMetricNewReq
	return r
}

func (r MetricMetricNewPostRequest) Execute() (*MerchantMetricDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricNewPostExecute(r)
}

/*
MetricNewPost New Merchant Metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricNewPostRequest
*/
func (a *MetricService) MetricNewPost(ctx context.Context) MetricMetricNewPostRequest {
	return MetricMetricNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricDeletePost200Response
func (a *MetricService) MetricNewPostExecute(r MetricMetricNewPostRequest) (*MerchantMetricDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricNewReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricNewReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricNewReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricPlanLimitDeletePostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricDeletePlanLimitReq *UnibeeApiMerchantMetricDeletePlanLimitReq
}

func (r MetricMetricPlanLimitDeletePostRequest) UnibeeApiMerchantMetricDeletePlanLimitReq(unibeeApiMerchantMetricDeletePlanLimitReq UnibeeApiMerchantMetricDeletePlanLimitReq) MetricMetricPlanLimitDeletePostRequest {
	r.unibeeApiMerchantMetricDeletePlanLimitReq = &unibeeApiMerchantMetricDeletePlanLimitReq
	return r
}

func (r MetricMetricPlanLimitDeletePostRequest) Execute() (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricPlanLimitDeletePostExecute(r)
}

/*
MetricPlanLimitDeletePost Delete Merchant Metric Plan TotalLimit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricPlanLimitDeletePostRequest
*/
func (a *MetricService) MetricPlanLimitDeletePost(ctx context.Context) MetricMetricPlanLimitDeletePostRequest {
	return MetricMetricPlanLimitDeletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricPlanLimitDeletePost200Response
func (a *MetricService) MetricPlanLimitDeletePostExecute(r MetricMetricPlanLimitDeletePostRequest) (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricPlanLimitDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricPlanLimitDeletePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/plan/limit/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricDeletePlanLimitReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricDeletePlanLimitReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricDeletePlanLimitReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricPlanLimitEditPostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricEditPlanLimitReq *UnibeeApiMerchantMetricEditPlanLimitReq
}

func (r MetricMetricPlanLimitEditPostRequest) UnibeeApiMerchantMetricEditPlanLimitReq(unibeeApiMerchantMetricEditPlanLimitReq UnibeeApiMerchantMetricEditPlanLimitReq) MetricMetricPlanLimitEditPostRequest {
	r.unibeeApiMerchantMetricEditPlanLimitReq = &unibeeApiMerchantMetricEditPlanLimitReq
	return r
}

func (r MetricMetricPlanLimitEditPostRequest) Execute() (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricPlanLimitEditPostExecute(r)
}

/*
MetricPlanLimitEditPost Edit Merchant Metric Plan TotalLimit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricPlanLimitEditPostRequest
*/
func (a *MetricService) MetricPlanLimitEditPost(ctx context.Context) MetricMetricPlanLimitEditPostRequest {
	return MetricMetricPlanLimitEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricPlanLimitDeletePost200Response
func (a *MetricService) MetricPlanLimitEditPostExecute(r MetricMetricPlanLimitEditPostRequest) (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricPlanLimitDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricPlanLimitEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/plan/limit/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricEditPlanLimitReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricEditPlanLimitReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricEditPlanLimitReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MetricMetricPlanLimitNewPostRequest struct {
	ctx context.Context
	ApiService *MetricService
	unibeeApiMerchantMetricNewPlanLimitReq *UnibeeApiMerchantMetricNewPlanLimitReq
}

func (r MetricMetricPlanLimitNewPostRequest) UnibeeApiMerchantMetricNewPlanLimitReq(unibeeApiMerchantMetricNewPlanLimitReq UnibeeApiMerchantMetricNewPlanLimitReq) MetricMetricPlanLimitNewPostRequest {
	r.unibeeApiMerchantMetricNewPlanLimitReq = &unibeeApiMerchantMetricNewPlanLimitReq
	return r
}

func (r MetricMetricPlanLimitNewPostRequest) Execute() (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	return r.ApiService.MetricPlanLimitNewPostExecute(r)
}

/*
MetricPlanLimitNewPost New Merchant Metric Plan TotalLimit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MetricMetricPlanLimitNewPostRequest
*/
func (a *MetricService) MetricPlanLimitNewPost(ctx context.Context) MetricMetricPlanLimitNewPostRequest {
	return MetricMetricPlanLimitNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricPlanLimitDeletePost200Response
func (a *MetricService) MetricPlanLimitNewPostExecute(r MetricMetricPlanLimitNewPostRequest) (*MerchantMetricPlanLimitDeletePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricPlanLimitDeletePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricService.MetricPlanLimitNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/plan/limit/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantMetricNewPlanLimitReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantMetricNewPlanLimitReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantMetricNewPlanLimitReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
