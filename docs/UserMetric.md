# \UserMetric

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricEventDeletePost**](UserMetric.md#MetricEventDeletePost) | **Post** /merchant/metric/event/delete | Del Merchant Metric Event
[**MetricEventNewPost**](UserMetric.md#MetricEventNewPost) | **Post** /merchant/metric/event/new | New Merchant Metric Event
[**MetricUserMetricGet**](UserMetric.md#MetricUserMetricGet) | **Get** /merchant/metric/user/metric | Query User Metric



## MetricEventDeletePost

> MerchantAuthSsoLoginOTPPost200Response MetricEventDeletePost(ctx).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()

Del Merchant Metric Event

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
	unibeeApiMerchantMetricDeleteEventReq := *openapiclient.NewUnibeeApiMerchantMetricDeleteEventReq("ExternalEventId_example", "ExternalUserId_example", "MetricCode_example") // UnibeeApiMerchantMetricDeleteEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricEventDeletePost(context.Background()).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricEventDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricEventDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDeleteEventReq** | [**UnibeeApiMerchantMetricDeleteEventReq**](UnibeeApiMerchantMetricDeleteEventReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricEventNewPost

> MerchantMetricEventNewPost200Response MetricEventNewPost(ctx).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()

New Merchant Metric Event

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
	unibeeApiMerchantMetricNewEventReq := *openapiclient.NewUnibeeApiMerchantMetricNewEventReq("ExternalEventId_example", "ExternalUserId_example", "MetricCode_example") // UnibeeApiMerchantMetricNewEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricEventNewPost(context.Background()).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricEventNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventNewPost`: MerchantMetricEventNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricEventNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricNewEventReq** | [**UnibeeApiMerchantMetricNewEventReq**](UnibeeApiMerchantMetricNewEventReq.md) |  | 

### Return type

[**MerchantMetricEventNewPost200Response**](MerchantMetricEventNewPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricUserMetricGet

> MerchantMetricUserMetricGet200Response MetricUserMetricGet(ctx).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()

Query User Metric

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
	userId := int64(789) // int64 | UserId, One Of UserId|ExternalUserId Needed (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, One Of UserId|ExternalUserId Needed (optional)
	productId := int64(789) // int64 | default product will use if productId not specified and subscriptionId is blank (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricUserMetricGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()
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
 **productId** | **int64** | default product will use if productId not specified and subscriptionId is blank | 

### Return type

[**MerchantMetricUserMetricGet200Response**](MerchantMetricUserMetricGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

