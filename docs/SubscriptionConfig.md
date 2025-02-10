# \SubscriptionConfig

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionConfigGet**](SubscriptionConfig.md#SubscriptionConfigGet) | **Get** /merchant/subscription/config | Subscription Config
[**SubscriptionConfigUpdatePost**](SubscriptionConfig.md#SubscriptionConfigUpdatePost) | **Post** /merchant/subscription/config/update | Update Merchant Subscription Config



## SubscriptionConfigGet

> MerchantSubscriptionConfigGet200Response SubscriptionConfigGet(ctx).Execute()

Subscription Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionConfig.SubscriptionConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionConfig.SubscriptionConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigGet`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionConfig.SubscriptionConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigGetRequest struct via the builder pattern


### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionConfigUpdatePost

> MerchantSubscriptionConfigGet200Response SubscriptionConfigUpdatePost(ctx).UnibeeApiMerchantSubscriptionConfigUpdateReq(unibeeApiMerchantSubscriptionConfigUpdateReq).Execute()

Update Merchant Subscription Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionConfigUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionConfigUpdateReq() // UnibeeApiMerchantSubscriptionConfigUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionConfig.SubscriptionConfigUpdatePost(context.Background()).UnibeeApiMerchantSubscriptionConfigUpdateReq(unibeeApiMerchantSubscriptionConfigUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionConfig.SubscriptionConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigUpdatePost`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionConfig.SubscriptionConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionConfigUpdateReq** | [**UnibeeApiMerchantSubscriptionConfigUpdateReq**](UnibeeApiMerchantSubscriptionConfigUpdateReq.md) |  | 

### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

