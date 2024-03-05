# \MerchantUserMetricAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantMetricUserMetricGet**](MerchantUserMetricAPI.md#MerchantMetricUserMetricGet) | **Get** /merchant/metric/user/metric | Query User Metric



## MerchantMetricUserMetricGet

> MerchantMetricUserMetricGet200Response MerchantMetricUserMetricGet(ctx).UserId(userId).ExternalUserId(externalUserId).Execute()

Query User Metric

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
	userId := int64(789) // int64 | UserId, One Of UserId|ExternalUserId Needed (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, One Of UserId|ExternalUserId Needed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantUserMetricAPI.MerchantMetricUserMetricGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserMetricAPI.MerchantMetricUserMetricGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricUserMetricGet`: MerchantMetricUserMetricGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserMetricAPI.MerchantMetricUserMetricGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricUserMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId, One Of UserId|ExternalUserId Needed | 
 **externalUserId** | **string** | ExternalUserId, One Of UserId|ExternalUserId Needed | 

### Return type

[**MerchantMetricUserMetricGet200Response**](MerchantMetricUserMetricGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

