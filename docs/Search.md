# \Search

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchKeySearchGet**](Search.md#SearchKeySearchGet) | **Get** /merchant/search/key_search | Search
[**SearchKeySearchPost**](Search.md#SearchKeySearchPost) | **Post** /merchant/search/key_search | Search



## SearchKeySearchGet

> MerchantSearchKeySearchGet200Response SearchKeySearchGet(ctx).SearchKey(searchKey).Execute()

Search

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
	searchKey := "searchKey_example" // string | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Search.SearchKeySearchGet(context.Background()).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchKeySearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchKeySearchGet`: MerchantSearchKeySearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Search.SearchKeySearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKeySearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId | 

### Return type

[**MerchantSearchKeySearchGet200Response**](MerchantSearchKeySearchGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKeySearchPost

> MerchantSearchKeySearchGet200Response SearchKeySearchPost(ctx).UnibeeApiMerchantSearchSearchReq(unibeeApiMerchantSearchSearchReq).Execute()

Search

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
	unibeeApiMerchantSearchSearchReq := *openapiclient.NewUnibeeApiMerchantSearchSearchReq() // UnibeeApiMerchantSearchSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Search.SearchKeySearchPost(context.Background()).UnibeeApiMerchantSearchSearchReq(unibeeApiMerchantSearchSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchKeySearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchKeySearchPost`: MerchantSearchKeySearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Search.SearchKeySearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKeySearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSearchSearchReq** | [**UnibeeApiMerchantSearchSearchReq**](UnibeeApiMerchantSearchSearchReq.md) |  | 

### Return type

[**MerchantSearchKeySearchGet200Response**](MerchantSearchKeySearchGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

