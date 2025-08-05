# \SubscriptionImport

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionActiveSubscriptionImportPost**](SubscriptionImport.md#SubscriptionActiveSubscriptionImportPost) | **Post** /merchant/subscription/active_subscription_import | Active Subscription Import(allows repetition imports)
[**SubscriptionHistorySubscriptionImportPost**](SubscriptionImport.md#SubscriptionHistorySubscriptionImportPost) | **Post** /merchant/subscription/history_subscription_import | History Subscription Import(Allows repetition imports)



## SubscriptionActiveSubscriptionImportPost

> MerchantSubscriptionActiveSubscriptionImportPost200Response SubscriptionActiveSubscriptionImportPost(ctx).UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq(unibeeApiMerchantSubscriptionActiveSubscriptionImportReq).Execute()

Active Subscription Import(allows repetition imports)



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
	unibeeApiMerchantSubscriptionActiveSubscriptionImportReq := *openapiclient.NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReq() // UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImport.SubscriptionActiveSubscriptionImportPost(context.Background()).UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq(unibeeApiMerchantSubscriptionActiveSubscriptionImportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImport.SubscriptionActiveSubscriptionImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionActiveSubscriptionImportPost`: MerchantSubscriptionActiveSubscriptionImportPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImport.SubscriptionActiveSubscriptionImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionActiveSubscriptionImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionActiveSubscriptionImportReq** | [**UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq**](UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq.md) |  | 

### Return type

[**MerchantSubscriptionActiveSubscriptionImportPost200Response**](MerchantSubscriptionActiveSubscriptionImportPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionHistorySubscriptionImportPost

> MerchantSubscriptionActiveSubscriptionImportPost200Response SubscriptionHistorySubscriptionImportPost(ctx).UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq(unibeeApiMerchantSubscriptionHistorySubscriptionImportReq).Execute()

History Subscription Import(Allows repetition imports)



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
	unibeeApiMerchantSubscriptionHistorySubscriptionImportReq := *openapiclient.NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReq() // UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImport.SubscriptionHistorySubscriptionImportPost(context.Background()).UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq(unibeeApiMerchantSubscriptionHistorySubscriptionImportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImport.SubscriptionHistorySubscriptionImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionHistorySubscriptionImportPost`: MerchantSubscriptionActiveSubscriptionImportPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImport.SubscriptionHistorySubscriptionImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionHistorySubscriptionImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionHistorySubscriptionImportReq** | [**UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq**](UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq.md) |  | 

### Return type

[**MerchantSubscriptionActiveSubscriptionImportPost200Response**](MerchantSubscriptionActiveSubscriptionImportPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

