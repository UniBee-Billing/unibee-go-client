# \MerchantSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSearchKeySearchGet**](MerchantSearchAPI.md#MerchantSearchKeySearchGet) | **Get** /merchant/search/key_search | Merchant Search
[**MerchantSearchKeySearchPost**](MerchantSearchAPI.md#MerchantSearchKeySearchPost) | **Post** /merchant/search/key_search | Merchant Search



## MerchantSearchKeySearchGet

> MerchantSearchKeySearchGet200Response MerchantSearchKeySearchGet(ctx).SearchKey(searchKey).Execute()

Merchant Search

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
	searchKey := "searchKey_example" // string | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSearchAPI.MerchantSearchKeySearchGet(context.Background()).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSearchAPI.MerchantSearchKeySearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSearchKeySearchGet`: MerchantSearchKeySearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSearchAPI.MerchantSearchKeySearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSearchKeySearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId | 

### Return type

[**MerchantSearchKeySearchGet200Response**](MerchantSearchKeySearchGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSearchKeySearchPost

> MerchantSearchKeySearchGet200Response MerchantSearchKeySearchPost(ctx).UnibeeApiMerchantSearchSearchReq(unibeeApiMerchantSearchSearchReq).Execute()

Merchant Search

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
	unibeeApiMerchantSearchSearchReq := *openapiclient.NewUnibeeApiMerchantSearchSearchReq() // UnibeeApiMerchantSearchSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSearchAPI.MerchantSearchKeySearchPost(context.Background()).UnibeeApiMerchantSearchSearchReq(unibeeApiMerchantSearchSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSearchAPI.MerchantSearchKeySearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSearchKeySearchPost`: MerchantSearchKeySearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSearchAPI.MerchantSearchKeySearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSearchKeySearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSearchSearchReq** | [**UnibeeApiMerchantSearchSearchReq**](UnibeeApiMerchantSearchSearchReq.md) |  | 

### Return type

[**MerchantSearchKeySearchGet200Response**](MerchantSearchKeySearchGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

