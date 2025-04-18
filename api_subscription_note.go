/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
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


// SubscriptionNoteService SubscriptionNote service
type SubscriptionNoteService service

type SubscriptionNoteSubscriptionAdminNoteListGetRequest struct {
	ctx context.Context
	ApiService *SubscriptionNoteService
	subscriptionId *string
	page *int32
	count *int32
}

// SubscriptionId
func (r SubscriptionNoteSubscriptionAdminNoteListGetRequest) SubscriptionId(subscriptionId string) SubscriptionNoteSubscriptionAdminNoteListGetRequest {
	r.subscriptionId = &subscriptionId
	return r
}

// Page, Start With 0
func (r SubscriptionNoteSubscriptionAdminNoteListGetRequest) Page(page int32) SubscriptionNoteSubscriptionAdminNoteListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r SubscriptionNoteSubscriptionAdminNoteListGetRequest) Count(count int32) SubscriptionNoteSubscriptionAdminNoteListGetRequest {
	r.count = &count
	return r
}

func (r SubscriptionNoteSubscriptionAdminNoteListGetRequest) Execute() (*MerchantSubscriptionAdminNoteListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionAdminNoteListGetExecute(r)
}

/*
SubscriptionAdminNoteListGet Get Subscription Note List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionNoteSubscriptionAdminNoteListGetRequest
*/
func (a *SubscriptionNoteService) SubscriptionAdminNoteListGet(ctx context.Context) SubscriptionNoteSubscriptionAdminNoteListGetRequest {
	return SubscriptionNoteSubscriptionAdminNoteListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionAdminNoteListGet200Response
func (a *SubscriptionNoteService) SubscriptionAdminNoteListGetExecute(r SubscriptionNoteSubscriptionAdminNoteListGetRequest) (*MerchantSubscriptionAdminNoteListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionAdminNoteListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionNoteService.SubscriptionAdminNoteListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/admin_note_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionId == nil {
		return localVarReturnValue, nil, reportError("subscriptionId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionId", r.subscriptionId, "")
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
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

type SubscriptionNoteSubscriptionAdminNoteListPostRequest struct {
	ctx context.Context
	ApiService *SubscriptionNoteService
	unibeeApiMerchantSubscriptionAdminNoteListReq *UnibeeApiMerchantSubscriptionAdminNoteListReq
}

func (r SubscriptionNoteSubscriptionAdminNoteListPostRequest) UnibeeApiMerchantSubscriptionAdminNoteListReq(unibeeApiMerchantSubscriptionAdminNoteListReq UnibeeApiMerchantSubscriptionAdminNoteListReq) SubscriptionNoteSubscriptionAdminNoteListPostRequest {
	r.unibeeApiMerchantSubscriptionAdminNoteListReq = &unibeeApiMerchantSubscriptionAdminNoteListReq
	return r
}

func (r SubscriptionNoteSubscriptionAdminNoteListPostRequest) Execute() (*MerchantSubscriptionAdminNoteListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionAdminNoteListPostExecute(r)
}

/*
SubscriptionAdminNoteListPost Get Subscription Note List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionNoteSubscriptionAdminNoteListPostRequest
*/
func (a *SubscriptionNoteService) SubscriptionAdminNoteListPost(ctx context.Context) SubscriptionNoteSubscriptionAdminNoteListPostRequest {
	return SubscriptionNoteSubscriptionAdminNoteListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionAdminNoteListGet200Response
func (a *SubscriptionNoteService) SubscriptionAdminNoteListPostExecute(r SubscriptionNoteSubscriptionAdminNoteListPostRequest) (*MerchantSubscriptionAdminNoteListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionAdminNoteListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionNoteService.SubscriptionAdminNoteListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/admin_note_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSubscriptionAdminNoteListReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSubscriptionAdminNoteListReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSubscriptionAdminNoteListReq
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

type SubscriptionNoteSubscriptionNewAdminNotePostRequest struct {
	ctx context.Context
	ApiService *SubscriptionNoteService
	unibeeApiMerchantSubscriptionNewAdminNoteReq *UnibeeApiMerchantSubscriptionNewAdminNoteReq
}

func (r SubscriptionNoteSubscriptionNewAdminNotePostRequest) UnibeeApiMerchantSubscriptionNewAdminNoteReq(unibeeApiMerchantSubscriptionNewAdminNoteReq UnibeeApiMerchantSubscriptionNewAdminNoteReq) SubscriptionNoteSubscriptionNewAdminNotePostRequest {
	r.unibeeApiMerchantSubscriptionNewAdminNoteReq = &unibeeApiMerchantSubscriptionNewAdminNoteReq
	return r
}

func (r SubscriptionNoteSubscriptionNewAdminNotePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.SubscriptionNewAdminNotePostExecute(r)
}

/*
SubscriptionNewAdminNotePost New Subscription Note

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionNoteSubscriptionNewAdminNotePostRequest
*/
func (a *SubscriptionNoteService) SubscriptionNewAdminNotePost(ctx context.Context) SubscriptionNoteSubscriptionNewAdminNotePostRequest {
	return SubscriptionNoteSubscriptionNewAdminNotePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *SubscriptionNoteService) SubscriptionNewAdminNotePostExecute(r SubscriptionNoteSubscriptionNewAdminNotePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionNoteService.SubscriptionNewAdminNotePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/new_admin_note"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSubscriptionNewAdminNoteReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSubscriptionNewAdminNoteReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSubscriptionNewAdminNoteReq
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
