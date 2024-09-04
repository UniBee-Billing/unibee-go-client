/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// ProductService Product service
type ProductService service

type ProductProductActivatePostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductActivateReq *UnibeeApiMerchantProductActivateReq
}

func (r ProductProductActivatePostRequest) UnibeeApiMerchantProductActivateReq(unibeeApiMerchantProductActivateReq UnibeeApiMerchantProductActivateReq) ProductProductActivatePostRequest {
	r.unibeeApiMerchantProductActivateReq = &unibeeApiMerchantProductActivateReq
	return r
}

func (r ProductProductActivatePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.ProductActivatePostExecute(r)
}

/*
ProductActivatePost ActivateProduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductActivatePostRequest
*/
func (a *ProductService) ProductActivatePost(ctx context.Context) ProductProductActivatePostRequest {
	return ProductProductActivatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *ProductService) ProductActivatePostExecute(r ProductProductActivatePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductActivatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/activate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductActivateReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductActivateReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductActivateReq
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

type ProductProductCopyPostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductCopyReq *UnibeeApiMerchantProductCopyReq
}

func (r ProductProductCopyPostRequest) UnibeeApiMerchantProductCopyReq(unibeeApiMerchantProductCopyReq UnibeeApiMerchantProductCopyReq) ProductProductCopyPostRequest {
	r.unibeeApiMerchantProductCopyReq = &unibeeApiMerchantProductCopyReq
	return r
}

func (r ProductProductCopyPostRequest) Execute() (*MerchantProductCopyPost200Response, *http.Response, error) {
	return r.ApiService.ProductCopyPostExecute(r)
}

/*
ProductCopyPost CopyProduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductCopyPostRequest
*/
func (a *ProductService) ProductCopyPost(ctx context.Context) ProductProductCopyPostRequest {
	return ProductProductCopyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductCopyPost200Response
func (a *ProductService) ProductCopyPostExecute(r ProductProductCopyPostRequest) (*MerchantProductCopyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductCopyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductCopyPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/copy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductCopyReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductCopyReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductCopyReq
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

type ProductProductDeletePostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductDeleteReq *UnibeeApiMerchantProductDeleteReq
}

func (r ProductProductDeletePostRequest) UnibeeApiMerchantProductDeleteReq(unibeeApiMerchantProductDeleteReq UnibeeApiMerchantProductDeleteReq) ProductProductDeletePostRequest {
	r.unibeeApiMerchantProductDeleteReq = &unibeeApiMerchantProductDeleteReq
	return r
}

func (r ProductProductDeletePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.ProductDeletePostExecute(r)
}

/*
ProductDeletePost DeleteProduct

Product can being deleted when has no plan linked

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductDeletePostRequest
*/
func (a *ProductService) ProductDeletePost(ctx context.Context) ProductProductDeletePostRequest {
	return ProductProductDeletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *ProductService) ProductDeletePostExecute(r ProductProductDeletePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductDeletePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductDeleteReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductDeleteReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductDeleteReq
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

type ProductProductDetailGetRequest struct {
	ctx context.Context
	ApiService *ProductService
	productId *int64
}

// ProductId
func (r ProductProductDetailGetRequest) ProductId(productId int64) ProductProductDetailGetRequest {
	r.productId = &productId
	return r
}

func (r ProductProductDetailGetRequest) Execute() (*MerchantProductCopyPost200Response, *http.Response, error) {
	return r.ApiService.ProductDetailGetExecute(r)
}

/*
ProductDetailGet ProductDetail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductDetailGetRequest
*/
func (a *ProductService) ProductDetailGet(ctx context.Context) ProductProductDetailGetRequest {
	return ProductProductDetailGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductCopyPost200Response
func (a *ProductService) ProductDetailGetExecute(r ProductProductDetailGetRequest) (*MerchantProductCopyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductCopyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductDetailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productId == nil {
		return localVarReturnValue, nil, reportError("productId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "")
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

type ProductProductDetailPostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductDetailReq *UnibeeApiMerchantProductDetailReq
}

func (r ProductProductDetailPostRequest) UnibeeApiMerchantProductDetailReq(unibeeApiMerchantProductDetailReq UnibeeApiMerchantProductDetailReq) ProductProductDetailPostRequest {
	r.unibeeApiMerchantProductDetailReq = &unibeeApiMerchantProductDetailReq
	return r
}

func (r ProductProductDetailPostRequest) Execute() (*MerchantProductCopyPost200Response, *http.Response, error) {
	return r.ApiService.ProductDetailPostExecute(r)
}

/*
ProductDetailPost ProductDetail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductDetailPostRequest
*/
func (a *ProductService) ProductDetailPost(ctx context.Context) ProductProductDetailPostRequest {
	return ProductProductDetailPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductCopyPost200Response
func (a *ProductService) ProductDetailPostExecute(r ProductProductDetailPostRequest) (*MerchantProductCopyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductCopyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductDetailPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductDetailReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductDetailReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductDetailReq
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

type ProductProductEditPostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductEditReq *UnibeeApiMerchantProductEditReq
}

func (r ProductProductEditPostRequest) UnibeeApiMerchantProductEditReq(unibeeApiMerchantProductEditReq UnibeeApiMerchantProductEditReq) ProductProductEditPostRequest {
	r.unibeeApiMerchantProductEditReq = &unibeeApiMerchantProductEditReq
	return r
}

func (r ProductProductEditPostRequest) Execute() (*MerchantProductCopyPost200Response, *http.Response, error) {
	return r.ApiService.ProductEditPostExecute(r)
}

/*
ProductEditPost EditProduct

Edit exist product, product is editable for both active and inactive status 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductEditPostRequest
*/
func (a *ProductService) ProductEditPost(ctx context.Context) ProductProductEditPostRequest {
	return ProductProductEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductCopyPost200Response
func (a *ProductService) ProductEditPostExecute(r ProductProductEditPostRequest) (*MerchantProductCopyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductCopyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductEditReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductEditReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductEditReq
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

type ProductProductInactivatePostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductInactiveReq *UnibeeApiMerchantProductInactiveReq
}

func (r ProductProductInactivatePostRequest) UnibeeApiMerchantProductInactiveReq(unibeeApiMerchantProductInactiveReq UnibeeApiMerchantProductInactiveReq) ProductProductInactivatePostRequest {
	r.unibeeApiMerchantProductInactiveReq = &unibeeApiMerchantProductInactiveReq
	return r
}

func (r ProductProductInactivatePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.ProductInactivatePostExecute(r)
}

/*
ProductInactivatePost InactivateProduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductInactivatePostRequest
*/
func (a *ProductService) ProductInactivatePost(ctx context.Context) ProductProductInactivatePostRequest {
	return ProductProductInactivatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *ProductService) ProductInactivatePostExecute(r ProductProductInactivatePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductInactivatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/inactivate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductInactiveReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductInactiveReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductInactiveReq
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

type ProductProductListGetRequest struct {
	ctx context.Context
	ApiService *ProductService
	status *[]int32
	sortField *string
	sortType *string
	page *int32
	count *int32
}

// Filter, Default All，,Status，1-active，2-inactive
func (r ProductProductListGetRequest) Status(status []int32) ProductProductListGetRequest {
	r.status = &status
	return r
}

// Sort Field，id|create_time|gmt_modify，Default id
func (r ProductProductListGetRequest) SortField(sortField string) ProductProductListGetRequest {
	r.sortField = &sortField
	return r
}

// Sort Type，asc|desc，Default desc
func (r ProductProductListGetRequest) SortType(sortType string) ProductProductListGetRequest {
	r.sortType = &sortType
	return r
}

// Page, Start 0
func (r ProductProductListGetRequest) Page(page int32) ProductProductListGetRequest {
	r.page = &page
	return r
}

// Count Of Per Page
func (r ProductProductListGetRequest) Count(count int32) ProductProductListGetRequest {
	r.count = &count
	return r
}

func (r ProductProductListGetRequest) Execute() (*MerchantProductListGet200Response, *http.Response, error) {
	return r.ApiService.ProductListGetExecute(r)
}

/*
ProductListGet ProductList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductListGetRequest
*/
func (a *ProductService) ProductListGet(ctx context.Context) ProductProductListGetRequest {
	return ProductProductListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductListGet200Response
func (a *ProductService) ProductListGetExecute(r ProductProductListGetRequest) (*MerchantProductListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
		}
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

type ProductProductListPostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductListReq *UnibeeApiMerchantProductListReq
}

func (r ProductProductListPostRequest) UnibeeApiMerchantProductListReq(unibeeApiMerchantProductListReq UnibeeApiMerchantProductListReq) ProductProductListPostRequest {
	r.unibeeApiMerchantProductListReq = &unibeeApiMerchantProductListReq
	return r
}

func (r ProductProductListPostRequest) Execute() (*MerchantProductListGet200Response, *http.Response, error) {
	return r.ApiService.ProductListPostExecute(r)
}

/*
ProductListPost ProductList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductListPostRequest
*/
func (a *ProductService) ProductListPost(ctx context.Context) ProductProductListPostRequest {
	return ProductProductListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductListGet200Response
func (a *ProductService) ProductListPostExecute(r ProductProductListPostRequest) (*MerchantProductListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductListReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductListReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductListReq
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

type ProductProductNewPostRequest struct {
	ctx context.Context
	ApiService *ProductService
	unibeeApiMerchantProductNewReq *UnibeeApiMerchantProductNewReq
}

func (r ProductProductNewPostRequest) UnibeeApiMerchantProductNewReq(unibeeApiMerchantProductNewReq UnibeeApiMerchantProductNewReq) ProductProductNewPostRequest {
	r.unibeeApiMerchantProductNewReq = &unibeeApiMerchantProductNewReq
	return r
}

func (r ProductProductNewPostRequest) Execute() (*MerchantProductCopyPost200Response, *http.Response, error) {
	return r.ApiService.ProductNewPostExecute(r)
}

/*
ProductNewPost CreateProduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductProductNewPostRequest
*/
func (a *ProductService) ProductNewPost(ctx context.Context) ProductProductNewPostRequest {
	return ProductProductNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantProductCopyPost200Response
func (a *ProductService) ProductNewPostExecute(r ProductProductNewPostRequest) (*MerchantProductCopyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantProductCopyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductService.ProductNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/product/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantProductNewReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantProductNewReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantProductNewReq
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
