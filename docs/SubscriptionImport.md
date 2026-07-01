# \SubscriptionImport

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionActiveSubscriptionImportPost**](SubscriptionImport.md#SubscriptionActiveSubscriptionImportPost) | **Post** /merchant/subscription/active_subscription_import | Active Subscription Import(allows repetition imports)
[**SubscriptionHistorySubscriptionImportPost**](SubscriptionImport.md#SubscriptionHistorySubscriptionImportPost) | **Post** /merchant/subscription/history_subscription_import | History Subscription Import(Allows repetition imports)
[**SubscriptionImportProrationBaseSetupPost**](SubscriptionImport.md#SubscriptionImportProrationBaseSetupPost) | **Post** /merchant/subscription/import_proration_base_setup | Setup historical proration base on a zero-amount subscription invoice



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


## SubscriptionImportProrationBaseSetupPost

> MerchantSubscriptionImportProrationBaseSetupPost200Response SubscriptionImportProrationBaseSetupPost(ctx).UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq(unibeeApiMerchantSubscriptionImportProrationBaseSetupReq).Execute()

Setup historical proration base on a zero-amount subscription invoice



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
	unibeeApiMerchantSubscriptionImportProrationBaseSetupReq := *openapiclient.NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReq() // UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImport.SubscriptionImportProrationBaseSetupPost(context.Background()).UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq(unibeeApiMerchantSubscriptionImportProrationBaseSetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImport.SubscriptionImportProrationBaseSetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionImportProrationBaseSetupPost`: MerchantSubscriptionImportProrationBaseSetupPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImport.SubscriptionImportProrationBaseSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionImportProrationBaseSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionImportProrationBaseSetupReq** | [**UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq**](UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq.md) |  | 

### Return type

[**MerchantSubscriptionImportProrationBaseSetupPost200Response**](MerchantSubscriptionImportProrationBaseSetupPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

