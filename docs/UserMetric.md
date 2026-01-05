# \UserMetric

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricUserHistoryMetricByInvoiceGet**](UserMetric.md#MetricUserHistoryMetricByInvoiceGet) | **Get** /merchant/metric/user/history/metric_by_invoice | Query User History Limit Metric By Invoice (History Query Mode, limit_metered and limit_recurring only)
[**MetricUserHistoryMetricBySubscriptionGet**](UserMetric.md#MetricUserHistoryMetricBySubscriptionGet) | **Get** /merchant/metric/user/history/metric_by_subscription | Query User History Limit Metric By Subscription (History Query Mode, limit_metered and limit_recurring only, non-active subscription only)
[**MetricUserMetricGet**](UserMetric.md#MetricUserMetricGet) | **Get** /merchant/metric/user/metric | Query User Metric
[**MetricUserSubMetricGet**](UserMetric.md#MetricUserSubMetricGet) | **Get** /merchant/metric/user/sub/metric | Query User Metric By Subscription



## MetricUserHistoryMetricByInvoiceGet

> MerchantMetricUserHistoryMetricByInvoiceGet200Response MetricUserHistoryMetricByInvoiceGet(ctx).InvoiceId(invoiceId).Execute()

Query User History Limit Metric By Invoice (History Query Mode, limit_metered and limit_recurring only)

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
	invoiceId := "invoiceId_example" // string | The unique id of invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricUserHistoryMetricByInvoiceGet(context.Background()).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricUserHistoryMetricByInvoiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricUserHistoryMetricByInvoiceGet`: MerchantMetricUserHistoryMetricByInvoiceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricUserHistoryMetricByInvoiceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricUserHistoryMetricByInvoiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string** | The unique id of invoice | 

### Return type

[**MerchantMetricUserHistoryMetricByInvoiceGet200Response**](MerchantMetricUserHistoryMetricByInvoiceGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricUserHistoryMetricBySubscriptionGet

> MerchantMetricUserHistoryMetricByInvoiceGet200Response MetricUserHistoryMetricBySubscriptionGet(ctx).SubscriptionId(subscriptionId).Execute()

Query User History Limit Metric By Subscription (History Query Mode, limit_metered and limit_recurring only, non-active subscription only)

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
	subscriptionId := "subscriptionId_example" // string | The unique id of subscription (must be non-active subscription or non-incomplete subscription)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricUserHistoryMetricBySubscriptionGet(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricUserHistoryMetricBySubscriptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricUserHistoryMetricBySubscriptionGet`: MerchantMetricUserHistoryMetricByInvoiceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricUserHistoryMetricBySubscriptionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricUserHistoryMetricBySubscriptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | The unique id of subscription (must be non-active subscription or non-incomplete subscription) | 

### Return type

[**MerchantMetricUserHistoryMetricByInvoiceGet200Response**](MerchantMetricUserHistoryMetricByInvoiceGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricUserMetricGet

> MerchantMetricUserMetricGet200Response MetricUserMetricGet(ctx).UserId(userId).Email(email).ExternalUserId(externalUserId).ProductId(productId).ReloadCache(reloadCache).Execute()

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
	userId := int64(789) // int64 | UserId, One Of UserId|Email|ExternalUserId Needed (optional)
	email := "email_example" // string | Email, One Of UserId|Email|ExternalUserId Needed (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, One Of UserId|Email|ExternalUserId Needed (optional)
	productId := int64(789) // int64 | Id of product, default product will use if productId not specified and subscriptionId is blank (optional)
	reloadCache := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricUserMetricGet(context.Background()).UserId(userId).Email(email).ExternalUserId(externalUserId).ProductId(productId).ReloadCache(reloadCache).Execute()
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
 **userId** | **int64** | UserId, One Of UserId|Email|ExternalUserId Needed | 
 **email** | **string** | Email, One Of UserId|Email|ExternalUserId Needed | 
 **externalUserId** | **string** | ExternalUserId, One Of UserId|Email|ExternalUserId Needed | 
 **productId** | **int64** | Id of product, default product will use if productId not specified and subscriptionId is blank | 
 **reloadCache** | **bool** |  | 

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


## MetricUserSubMetricGet

> MerchantMetricUserMetricGet200Response MetricUserSubMetricGet(ctx).SubscriptionId(subscriptionId).ReloadCache(reloadCache).Execute()

Query User Metric By Subscription

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
	subscriptionId := "subscriptionId_example" // string | SubscriptionId (optional)
	reloadCache := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserMetric.MetricUserSubMetricGet(context.Background()).SubscriptionId(subscriptionId).ReloadCache(reloadCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserMetric.MetricUserSubMetricGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricUserSubMetricGet`: MerchantMetricUserMetricGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserMetric.MetricUserSubMetricGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricUserSubMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 
 **reloadCache** | **bool** |  | 

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

