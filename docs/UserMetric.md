# \UserMetric

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricUserMetricGet**](UserMetric.md#MetricUserMetricGet) | **Get** /merchant/metric/user/metric | Query User Metric



## MetricUserMetricGet

> MerchantMetricUserMetricGet200Response MetricUserMetricGet(ctx).UserId(userId).ExternalUserId(externalUserId).Execute()

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
	resp, r, err := apiClient.UserMetric.MetricUserMetricGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricUserMetricGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricUserMetricGet`: MerchantMetricUserMetricGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricUserMetricGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricUserMetricGetRequest struct via the builder pattern


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

