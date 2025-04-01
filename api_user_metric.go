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


// UserMetricService UserMetric service
type UserMetricService service

type UserMetricMetricUserMetricGetRequest struct {
	ctx context.Context
	ApiService *UserMetricService
	userId *int64
	externalUserId *string
	productId *int64
}

// UserId, One Of UserId|ExternalUserId Needed
func (r UserMetricMetricUserMetricGetRequest) UserId(userId int64) UserMetricMetricUserMetricGetRequest {
	r.userId = &userId
	return r
}

// ExternalUserId, One Of UserId|ExternalUserId Needed
func (r UserMetricMetricUserMetricGetRequest) ExternalUserId(externalUserId string) UserMetricMetricUserMetricGetRequest {
	r.externalUserId = &externalUserId
	return r
}

// default product will use if productId not specified and subscriptionId is blank
func (r UserMetricMetricUserMetricGetRequest) ProductId(productId int64) UserMetricMetricUserMetricGetRequest {
	r.productId = &productId
	return r
}

func (r UserMetricMetricUserMetricGetRequest) Execute() (*MerchantMetricUserMetricGet200Response, *http.Response, error) {
	return r.ApiService.MetricUserMetricGetExecute(r)
}

/*
MetricUserMetricGet Query User Metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserMetricMetricUserMetricGetRequest
*/
func (a *UserMetricService) MetricUserMetricGet(ctx context.Context) UserMetricMetricUserMetricGetRequest {
	return UserMetricMetricUserMetricGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricUserMetricGet200Response
func (a *UserMetricService) MetricUserMetricGetExecute(r UserMetricMetricUserMetricGetRequest) (*MerchantMetricUserMetricGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricUserMetricGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserMetricService.MetricUserMetricGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/user/metric"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.externalUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "externalUserId", r.externalUserId, "")
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "")
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

type UserMetricMetricUserSubMetricGetRequest struct {
	ctx context.Context
	ApiService *UserMetricService
	subscriptionId *string
}

// SubscriptionId
func (r UserMetricMetricUserSubMetricGetRequest) SubscriptionId(subscriptionId string) UserMetricMetricUserSubMetricGetRequest {
	r.subscriptionId = &subscriptionId
	return r
}

func (r UserMetricMetricUserSubMetricGetRequest) Execute() (*MerchantMetricUserMetricGet200Response, *http.Response, error) {
	return r.ApiService.MetricUserSubMetricGetExecute(r)
}

/*
MetricUserSubMetricGet Query User Metric By Subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserMetricMetricUserSubMetricGetRequest
*/
func (a *UserMetricService) MetricUserSubMetricGet(ctx context.Context) UserMetricMetricUserSubMetricGetRequest {
	return UserMetricMetricUserSubMetricGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantMetricUserMetricGet200Response
func (a *UserMetricService) MetricUserSubMetricGetExecute(r UserMetricMetricUserSubMetricGetRequest) (*MerchantMetricUserMetricGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantMetricUserMetricGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserMetricService.MetricUserSubMetricGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/metric/user/sub/metric"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.subscriptionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionId", r.subscriptionId, "")
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
