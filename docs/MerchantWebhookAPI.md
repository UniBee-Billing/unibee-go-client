# \MerchantWebhookAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantWebhookDeleteEndpointPost**](MerchantWebhookAPI.md#MerchantWebhookDeleteEndpointPost) | **Post** /merchant/webhook/delete_endpoint | Merchant Delete Webhook Endpoint
[**MerchantWebhookEndpointListGet**](MerchantWebhookAPI.md#MerchantWebhookEndpointListGet) | **Get** /merchant/webhook/endpoint_list | Merchant Webhook Endpoint list
[**MerchantWebhookEndpointLogListGet**](MerchantWebhookAPI.md#MerchantWebhookEndpointLogListGet) | **Get** /merchant/webhook/endpoint_log_list | Merchant Webhook Endpoint Log list
[**MerchantWebhookEventListGet**](MerchantWebhookAPI.md#MerchantWebhookEventListGet) | **Get** /merchant/webhook/event_list | Webhook Event list
[**MerchantWebhookNewEndpointPost**](MerchantWebhookAPI.md#MerchantWebhookNewEndpointPost) | **Post** /merchant/webhook/new_endpoint | Merchant New Webhook Endpoint
[**MerchantWebhookUpdateEndpointPost**](MerchantWebhookAPI.md#MerchantWebhookUpdateEndpointPost) | **Post** /merchant/webhook/update_endpoint | Merchant Update Webhook Endpoint



## MerchantWebhookDeleteEndpointPost

> MerchantAuthSsoLoginOTPPost200Response MerchantWebhookDeleteEndpointPost(ctx).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()

Merchant Delete Webhook Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantWebhookDeleteEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookDeleteEndpointReq(int64(123)) // UnibeeApiMerchantWebhookDeleteEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookDeleteEndpointPost(context.Background()).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookDeleteEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookDeleteEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookDeleteEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookDeleteEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookDeleteEndpointReq** | [**UnibeeApiMerchantWebhookDeleteEndpointReq**](UnibeeApiMerchantWebhookDeleteEndpointReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantWebhookEndpointListGet

> MerchantWebhookEndpointListGet200Response MerchantWebhookEndpointListGet(ctx).Execute()

Merchant Webhook Endpoint list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookEndpointListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookEndpointListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookEndpointListGet`: MerchantWebhookEndpointListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookEndpointListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookEndpointListGetRequest struct via the builder pattern


### Return type

[**MerchantWebhookEndpointListGet200Response**](MerchantWebhookEndpointListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantWebhookEndpointLogListGet

> MerchantWebhookEndpointLogListGet200Response MerchantWebhookEndpointLogListGet(ctx).EndpointId(endpointId).Page(page).Count(count).Execute()

Merchant Webhook Endpoint Log list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	endpointId := int64(789) // int64 | EndpointId
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookEndpointLogListGet(context.Background()).EndpointId(endpointId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookEndpointLogListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookEndpointLogListGet`: MerchantWebhookEndpointLogListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookEndpointLogListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookEndpointLogListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointId** | **int64** | EndpointId | 
 **page** | **int32** | Page, Start WIth 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantWebhookEndpointLogListGet200Response**](MerchantWebhookEndpointLogListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantWebhookEventListGet

> MerchantWebhookEventListGet200Response MerchantWebhookEventListGet(ctx).Execute()

Webhook Event list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookEventListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookEventListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookEventListGet`: MerchantWebhookEventListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookEventListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookEventListGetRequest struct via the builder pattern


### Return type

[**MerchantWebhookEventListGet200Response**](MerchantWebhookEventListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantWebhookNewEndpointPost

> MerchantAuthSsoLoginOTPPost200Response MerchantWebhookNewEndpointPost(ctx).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()

Merchant New Webhook Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantWebhookNewEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookNewEndpointReq("Url_example") // UnibeeApiMerchantWebhookNewEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookNewEndpointPost(context.Background()).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookNewEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookNewEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookNewEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookNewEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookNewEndpointReq** | [**UnibeeApiMerchantWebhookNewEndpointReq**](UnibeeApiMerchantWebhookNewEndpointReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantWebhookUpdateEndpointPost

> MerchantAuthSsoLoginOTPPost200Response MerchantWebhookUpdateEndpointPost(ctx).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()

Merchant Update Webhook Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantWebhookUpdateEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookUpdateEndpointReq(int64(123), "Url_example") // UnibeeApiMerchantWebhookUpdateEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWebhookAPI.MerchantWebhookUpdateEndpointPost(context.Background()).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWebhookAPI.MerchantWebhookUpdateEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantWebhookUpdateEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantWebhookAPI.MerchantWebhookUpdateEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantWebhookUpdateEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookUpdateEndpointReq** | [**UnibeeApiMerchantWebhookUpdateEndpointReq**](UnibeeApiMerchantWebhookUpdateEndpointReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

