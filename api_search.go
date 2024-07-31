/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
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


// SearchService Search service
type SearchService service

type SearchSearchKeySearchGetRequest struct {
	ctx context.Context
	ApiService *SearchService
	searchKey *string
}

// SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId
func (r SearchSearchKeySearchGetRequest) SearchKey(searchKey string) SearchSearchKeySearchGetRequest {
	r.searchKey = &searchKey
	return r
}

func (r SearchSearchKeySearchGetRequest) Execute() (*MerchantSearchKeySearchGet200Response, *http.Response, error) {
	return r.ApiService.SearchKeySearchGetExecute(r)
}

/*
SearchKeySearchGet Search

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SearchSearchKeySearchGetRequest
*/
func (a *SearchService) SearchKeySearchGet(ctx context.Context) SearchSearchKeySearchGetRequest {
	return SearchSearchKeySearchGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSearchKeySearchGet200Response
func (a *SearchService) SearchKeySearchGetExecute(r SearchSearchKeySearchGetRequest) (*MerchantSearchKeySearchGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSearchKeySearchGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchService.SearchKeySearchGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/search/key_search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchKey", r.searchKey, "")
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

type SearchSearchKeySearchPostRequest struct {
	ctx context.Context
	ApiService *SearchService
	unibeeApiMerchantSearchSearchReq *UnibeeApiMerchantSearchSearchReq
}

func (r SearchSearchKeySearchPostRequest) UnibeeApiMerchantSearchSearchReq(unibeeApiMerchantSearchSearchReq UnibeeApiMerchantSearchSearchReq) SearchSearchKeySearchPostRequest {
	r.unibeeApiMerchantSearchSearchReq = &unibeeApiMerchantSearchSearchReq
	return r
}

func (r SearchSearchKeySearchPostRequest) Execute() (*MerchantSearchKeySearchGet200Response, *http.Response, error) {
	return r.ApiService.SearchKeySearchPostExecute(r)
}

/*
SearchKeySearchPost Search

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SearchSearchKeySearchPostRequest
*/
func (a *SearchService) SearchKeySearchPost(ctx context.Context) SearchSearchKeySearchPostRequest {
	return SearchSearchKeySearchPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSearchKeySearchGet200Response
func (a *SearchService) SearchKeySearchPostExecute(r SearchSearchKeySearchPostRequest) (*MerchantSearchKeySearchGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSearchKeySearchGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchService.SearchKeySearchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/search/key_search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSearchSearchReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSearchSearchReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSearchSearchReq
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
