# \Webhook

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDeleteEndpointPost**](Webhook.md#WebhookDeleteEndpointPost) | **Post** /merchant/webhook/delete_endpoint | DeleteWebhookEndpoint
[**WebhookEndpointListGet**](Webhook.md#WebhookEndpointListGet) | **Get** /merchant/webhook/endpoint_list | WebhookEndpointlist
[**WebhookEndpointLogListGet**](Webhook.md#WebhookEndpointLogListGet) | **Get** /merchant/webhook/endpoint_log_list | WebhookEndpointLogList
[**WebhookEventListGet**](Webhook.md#WebhookEventListGet) | **Get** /merchant/webhook/event_list | WebhookEventList
[**WebhookNewEndpointPost**](Webhook.md#WebhookNewEndpointPost) | **Post** /merchant/webhook/new_endpoint | NewWebhookEndpoint
[**WebhookResendPost**](Webhook.md#WebhookResendPost) | **Post** /merchant/webhook/resend | ResentWebhook
[**WebhookUpdateEndpointPost**](Webhook.md#WebhookUpdateEndpointPost) | **Post** /merchant/webhook/update_endpoint | UpdateWebhookEndpoint



## WebhookDeleteEndpointPost

> MerchantAuthSsoLoginOTPPost200Response WebhookDeleteEndpointPost(ctx).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()

DeleteWebhookEndpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantWebhookDeleteEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookDeleteEndpointReq(int64(123)) // UnibeeApiMerchantWebhookDeleteEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookDeleteEndpointPost(context.Background()).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookDeleteEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeleteEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookDeleteEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeleteEndpointPostRequest struct via the builder pattern


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


## WebhookEndpointListGet

> MerchantWebhookEndpointListGet200Response WebhookEndpointListGet(ctx).Execute()

WebhookEndpointlist

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookEndpointListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEndpointListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointListGet`: MerchantWebhookEndpointListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEndpointListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointListGetRequest struct via the builder pattern


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


## WebhookEndpointLogListGet

> MerchantWebhookEndpointLogListGet200Response WebhookEndpointLogListGet(ctx).EndpointId(endpointId).Page(page).Count(count).Execute()

WebhookEndpointLogList

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	endpointId := int64(789) // int64 | EndpointId
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookEndpointLogListGet(context.Background()).EndpointId(endpointId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEndpointLogListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointLogListGet`: MerchantWebhookEndpointLogListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEndpointLogListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointLogListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointId** | **int64** | EndpointId | 
 **page** | **int32** | Page, Start With 0 | 
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


## WebhookEventListGet

> MerchantWebhookEventListGet200Response WebhookEventListGet(ctx).Execute()

WebhookEventList

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookEventListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEventListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEventListGet`: MerchantWebhookEventListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEventListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEventListGetRequest struct via the builder pattern


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


## WebhookNewEndpointPost

> MerchantAuthSsoLoginOTPPost200Response WebhookNewEndpointPost(ctx).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()

NewWebhookEndpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantWebhookNewEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookNewEndpointReq("Url_example") // UnibeeApiMerchantWebhookNewEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookNewEndpointPost(context.Background()).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookNewEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookNewEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookNewEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookNewEndpointPostRequest struct via the builder pattern


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


## WebhookResendPost

> MerchantWebhookResendPost200Response WebhookResendPost(ctx).UnibeeApiMerchantWebhookResendWebhookReq(unibeeApiMerchantWebhookResendWebhookReq).Execute()

ResentWebhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantWebhookResendWebhookReq := *openapiclient.NewUnibeeApiMerchantWebhookResendWebhookReq(int64(123)) // UnibeeApiMerchantWebhookResendWebhookReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookResendPost(context.Background()).UnibeeApiMerchantWebhookResendWebhookReq(unibeeApiMerchantWebhookResendWebhookReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookResendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookResendPost`: MerchantWebhookResendPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookResendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookResendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookResendWebhookReq** | [**UnibeeApiMerchantWebhookResendWebhookReq**](UnibeeApiMerchantWebhookResendWebhookReq.md) |  | 

### Return type

[**MerchantWebhookResendPost200Response**](MerchantWebhookResendPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookUpdateEndpointPost

> MerchantAuthSsoLoginOTPPost200Response WebhookUpdateEndpointPost(ctx).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()

UpdateWebhookEndpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantWebhookUpdateEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookUpdateEndpointReq(int64(123), "Url_example") // UnibeeApiMerchantWebhookUpdateEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookUpdateEndpointPost(context.Background()).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookUpdateEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookUpdateEndpointPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookUpdateEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookUpdateEndpointPostRequest struct via the builder pattern


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

