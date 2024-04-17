# \Gateway

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayEditCountryConfigPost**](Gateway.md#GatewayEditCountryConfigPost) | **Post** /merchant/gateway/edit_country_config | PaymentGatewayCountryConfigEdit
[**GatewayEditPost**](Gateway.md#GatewayEditPost) | **Post** /merchant/gateway/edit | PaymentGatewayEdit
[**GatewayListGet**](Gateway.md#GatewayListGet) | **Get** /merchant/gateway/list | PaymentGatewayList
[**GatewaySetupPost**](Gateway.md#GatewaySetupPost) | **Post** /merchant/gateway/setup | PaymentGatewaySetup
[**GatewaySetupWebhookPost**](Gateway.md#GatewaySetupWebhookPost) | **Post** /merchant/gateway/setup_webhook | PaymentGatewayWebhookSetup



## GatewayEditCountryConfigPost

> MerchantAuthSsoLoginOTPPost200Response GatewayEditCountryConfigPost(ctx).UnibeeApiMerchantGatewayEditCountryConfigReq(unibeeApiMerchantGatewayEditCountryConfigReq).Execute()

PaymentGatewayCountryConfigEdit



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
	unibeeApiMerchantGatewayEditCountryConfigReq := *openapiclient.NewUnibeeApiMerchantGatewayEditCountryConfigReq(int64(123)) // UnibeeApiMerchantGatewayEditCountryConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayEditCountryConfigPost(context.Background()).UnibeeApiMerchantGatewayEditCountryConfigReq(unibeeApiMerchantGatewayEditCountryConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayEditCountryConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayEditCountryConfigPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayEditCountryConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayEditCountryConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayEditCountryConfigReq** | [**UnibeeApiMerchantGatewayEditCountryConfigReq**](UnibeeApiMerchantGatewayEditCountryConfigReq.md) |  | 

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


## GatewayEditPost

> MerchantAuthSsoLoginOTPPost200Response GatewayEditPost(ctx).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()

PaymentGatewayEdit



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
	unibeeApiMerchantGatewayEditReq := *openapiclient.NewUnibeeApiMerchantGatewayEditReq(int64(123)) // UnibeeApiMerchantGatewayEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayEditPost(context.Background()).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayEditPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayEditPostRequest struct via the builder pattern


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


## GatewayListGet

> MerchantGatewayListGet200Response GatewayListGet(ctx).Execute()

PaymentGatewayList

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
	resp, r, err := apiClient.Gateway.GatewayListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListGet`: MerchantGatewayListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListGetRequest struct via the builder pattern


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


## GatewaySetupPost

> MerchantAuthSsoLoginOTPPost200Response GatewaySetupPost(ctx).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()

PaymentGatewaySetup



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
	unibeeApiMerchantGatewaySetupReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupReq("GatewayName_example") // UnibeeApiMerchantGatewaySetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetupPost(context.Background()).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupPostRequest struct via the builder pattern


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


## GatewaySetupWebhookPost

> MerchantGatewaySetupWebhookPost200Response GatewaySetupWebhookPost(ctx).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()

PaymentGatewayWebhookSetup

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
	unibeeApiMerchantGatewaySetupWebhookReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupWebhookReq(int64(123)) // UnibeeApiMerchantGatewaySetupWebhookReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetupWebhookPost(context.Background()).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupWebhookPost`: MerchantGatewaySetupWebhookPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupWebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupWebhookReq** | [**UnibeeApiMerchantGatewaySetupWebhookReq**](UnibeeApiMerchantGatewaySetupWebhookReq.md) |  | 

### Return type

[**MerchantGatewaySetupWebhookPost200Response**](MerchantGatewaySetupWebhookPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

