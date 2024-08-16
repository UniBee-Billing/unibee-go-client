/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
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


// SubscriptionTimelineService SubscriptionTimeline service
type SubscriptionTimelineService service

type SubscriptionTimelineSubscriptionTimelineListGetRequest struct {
	ctx context.Context
	ApiService *SubscriptionTimelineService
	userId *int64
	sortField *string
	sortType *string
	page *int32
	count *int32
}

// Filter UserId, Default All 
func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) UserId(userId int64) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	r.userId = &userId
	return r
}

// Sort Field，gmt_create|gmt_modify，Default gmt_modify
func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) SortField(sortField string) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	r.sortField = &sortField
	return r
}

// Sort Type，asc|desc，Default desc
func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) SortType(sortType string) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	r.sortType = &sortType
	return r
}

// Page, Start With 0
func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) Page(page int32) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) Count(count int32) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	r.count = &count
	return r
}

func (r SubscriptionTimelineSubscriptionTimelineListGetRequest) Execute() (*MerchantSubscriptionTimelineListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionTimelineListGetExecute(r)
}

/*
SubscriptionTimelineListGet SubscriptionTimeLineList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionTimelineSubscriptionTimelineListGetRequest
*/
func (a *SubscriptionTimelineService) SubscriptionTimelineListGet(ctx context.Context) SubscriptionTimelineSubscriptionTimelineListGetRequest {
	return SubscriptionTimelineSubscriptionTimelineListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionTimelineListGet200Response
func (a *SubscriptionTimelineService) SubscriptionTimelineListGetExecute(r SubscriptionTimelineSubscriptionTimelineListGetRequest) (*MerchantSubscriptionTimelineListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionTimelineListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionTimelineService.SubscriptionTimelineListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/timeline_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortField", r.sortField, "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortType", r.sortType, "")
	}
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

type SubscriptionTimelineSubscriptionTimelineListPostRequest struct {
	ctx context.Context
	ApiService *SubscriptionTimelineService
	unibeeApiMerchantSubscriptionTimeLineListReq *UnibeeApiMerchantSubscriptionTimeLineListReq
}

func (r SubscriptionTimelineSubscriptionTimelineListPostRequest) UnibeeApiMerchantSubscriptionTimeLineListReq(unibeeApiMerchantSubscriptionTimeLineListReq UnibeeApiMerchantSubscriptionTimeLineListReq) SubscriptionTimelineSubscriptionTimelineListPostRequest {
	r.unibeeApiMerchantSubscriptionTimeLineListReq = &unibeeApiMerchantSubscriptionTimeLineListReq
	return r
}

func (r SubscriptionTimelineSubscriptionTimelineListPostRequest) Execute() (*MerchantSubscriptionTimelineListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionTimelineListPostExecute(r)
}

/*
SubscriptionTimelineListPost SubscriptionTimeLineList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionTimelineSubscriptionTimelineListPostRequest
*/
func (a *SubscriptionTimelineService) SubscriptionTimelineListPost(ctx context.Context) SubscriptionTimelineSubscriptionTimelineListPostRequest {
	return SubscriptionTimelineSubscriptionTimelineListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionTimelineListGet200Response
func (a *SubscriptionTimelineService) SubscriptionTimelineListPostExecute(r SubscriptionTimelineSubscriptionTimelineListPostRequest) (*MerchantSubscriptionTimelineListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionTimelineListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionTimelineService.SubscriptionTimelineListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/timeline_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSubscriptionTimeLineListReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSubscriptionTimeLineListReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSubscriptionTimeLineListReq
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
