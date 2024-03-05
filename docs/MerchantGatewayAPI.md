# \MerchantGatewayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantGatewayEditPost**](MerchantGatewayAPI.md#MerchantGatewayEditPost) | **Post** /merchant/gateway/edit | Gateway Webhook Edit
[**MerchantGatewayListGet**](MerchantGatewayAPI.md#MerchantGatewayListGet) | **Get** /merchant/gateway/list | Gateway List
[**MerchantGatewaySetupPost**](MerchantGatewayAPI.md#MerchantGatewaySetupPost) | **Post** /merchant/gateway/setup | Gateway Setup
[**MerchantGatewaySetupWebhookPost**](MerchantGatewayAPI.md#MerchantGatewaySetupWebhookPost) | **Post** /merchant/gateway/setup_webhook | Gateway Webhook Setup



## MerchantGatewayEditPost

> MerchantAuthSsoLoginOTPPost200Response MerchantGatewayEditPost(ctx).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()

Gateway Webhook Edit

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
	unibeeApiMerchantGatewayEditReq := *openapiclient.NewUnibeeApiMerchantGatewayEditReq(int32(123)) // UnibeeApiMerchantGatewayEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantGatewayAPI.MerchantGatewayEditPost(context.Background()).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantGatewayAPI.MerchantGatewayEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantGatewayEditPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantGatewayAPI.MerchantGatewayEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantGatewayEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayEditReq** | [**UnibeeApiMerchantGatewayEditReq**](UnibeeApiMerchantGatewayEditReq.md) |  | 

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


## MerchantGatewayListGet

> MerchantGatewayListGet200Response MerchantGatewayListGet(ctx).Execute()

Gateway List

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
	resp, r, err := apiClient.MerchantGatewayAPI.MerchantGatewayListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantGatewayAPI.MerchantGatewayListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantGatewayListGet`: MerchantGatewayListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantGatewayAPI.MerchantGatewayListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantGatewayListGetRequest struct via the builder pattern


### Return type

[**MerchantGatewayListGet200Response**](MerchantGatewayListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantGatewaySetupPost

> MerchantAuthSsoLoginOTPPost200Response MerchantGatewaySetupPost(ctx).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()

Gateway Setup

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
	unibeeApiMerchantGatewaySetupReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupReq("GatewayName_example") // UnibeeApiMerchantGatewaySetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantGatewayAPI.MerchantGatewaySetupPost(context.Background()).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantGatewayAPI.MerchantGatewaySetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantGatewaySetupPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantGatewayAPI.MerchantGatewaySetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantGatewaySetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupReq** | [**UnibeeApiMerchantGatewaySetupReq**](UnibeeApiMerchantGatewaySetupReq.md) |  | 

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


## MerchantGatewaySetupWebhookPost

> MerchantAuthSsoLoginOTPPost200Response MerchantGatewaySetupWebhookPost(ctx).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()

Gateway Webhook Setup

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
	unibeeApiMerchantGatewaySetupWebhookReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupWebhookReq(int32(123)) // UnibeeApiMerchantGatewaySetupWebhookReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantGatewayAPI.MerchantGatewaySetupWebhookPost(context.Background()).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantGatewayAPI.MerchantGatewaySetupWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantGatewaySetupWebhookPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantGatewayAPI.MerchantGatewaySetupWebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantGatewaySetupWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupWebhookReq** | [**UnibeeApiMerchantGatewaySetupWebhookReq**](UnibeeApiMerchantGatewaySetupWebhookReq.md) |  | 

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

