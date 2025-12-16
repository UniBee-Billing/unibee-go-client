# \MetricLimitQuota

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricCarryoverReprocessPost**](MetricLimitQuota.md#MetricCarryoverReprocessPost) | **Post** /merchant/metric/carryover/reprocess | Reprocess Metric Carry Over (Admin Tool)
[**MetricLimitAdjustPost**](MetricLimitQuota.md#MetricLimitAdjustPost) | **Post** /merchant/metric/limit_adjust | Manual Adjust Metric Limit



## MetricCarryoverReprocessPost

> MerchantMetricCarryoverReprocessPost200Response MetricCarryoverReprocessPost(ctx).UnibeeApiMerchantMetricMetricCarryoverReprocessReq(unibeeApiMerchantMetricMetricCarryoverReprocessReq).Execute()

Reprocess Metric Carry Over (Admin Tool)

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
	unibeeApiMerchantMetricMetricCarryoverReprocessReq := *openapiclient.NewUnibeeApiMerchantMetricMetricCarryoverReprocessReq("PreviousInvoiceId_example") // UnibeeApiMerchantMetricMetricCarryoverReprocessReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricLimitQuota.MetricCarryoverReprocessPost(context.Background()).UnibeeApiMerchantMetricMetricCarryoverReprocessReq(unibeeApiMerchantMetricMetricCarryoverReprocessReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricLimitQuota.MetricCarryoverReprocessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricCarryoverReprocessPost`: MerchantMetricCarryoverReprocessPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricLimitQuota.MetricCarryoverReprocessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricCarryoverReprocessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricMetricCarryoverReprocessReq** | [**UnibeeApiMerchantMetricMetricCarryoverReprocessReq**](UnibeeApiMerchantMetricMetricCarryoverReprocessReq.md) |  | 

### Return type

[**MerchantMetricCarryoverReprocessPost200Response**](MerchantMetricCarryoverReprocessPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricLimitAdjustPost

> MerchantMetricLimitAdjustPost200Response MetricLimitAdjustPost(ctx).UnibeeApiMerchantMetricMetricLimitAdjustReq(unibeeApiMerchantMetricMetricLimitAdjustReq).Execute()

Manual Adjust Metric Limit

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
	unibeeApiMerchantMetricMetricLimitAdjustReq := *openapiclient.NewUnibeeApiMerchantMetricMetricLimitAdjustReq(int64(123), "MetricCode_example", "Reason_example") // UnibeeApiMerchantMetricMetricLimitAdjustReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricLimitQuota.MetricLimitAdjustPost(context.Background()).UnibeeApiMerchantMetricMetricLimitAdjustReq(unibeeApiMerchantMetricMetricLimitAdjustReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricLimitQuota.MetricLimitAdjustPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricLimitAdjustPost`: MerchantMetricLimitAdjustPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricLimitQuota.MetricLimitAdjustPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricLimitAdjustPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricMetricLimitAdjustReq** | [**UnibeeApiMerchantMetricMetricLimitAdjustReq**](UnibeeApiMerchantMetricMetricLimitAdjustReq.md) |  | 

### Return type

[**MerchantMetricLimitAdjustPost200Response**](MerchantMetricLimitAdjustPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

