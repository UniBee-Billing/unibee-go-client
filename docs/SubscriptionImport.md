# \SubscriptionImport

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionActiveSubscriptionImportGet**](SubscriptionImport.md#SubscriptionActiveSubscriptionImportGet) | **Get** /merchant/subscription/active_subscription_import | Active Subscription Import(Allows repetition imports)
[**SubscriptionHistorySubscriptionImportGet**](SubscriptionImport.md#SubscriptionHistorySubscriptionImportGet) | **Get** /merchant/subscription/history_subscription_import | History Subscription Import(Allows repetition imports)



## SubscriptionActiveSubscriptionImportGet

> MerchantSubscriptionActiveSubscriptionImportGet200Response SubscriptionActiveSubscriptionImportGet(ctx).ExternalSubscriptionId(externalSubscriptionId).ExternalPlanId(externalPlanId).Quantity(quantity).CountryCode(countryCode).Gateway(gateway).CurrentPeriodStart(currentPeriodStart).CurrentPeriodEnd(currentPeriodEnd).BillingCycleAnchor(billingCycleAnchor).CreateTime(createTime).Email(email).ExternalUserId(externalUserId).VatNumber(vatNumber).TaxPercentage(taxPercentage).FirstPaidTime(firstPaidTime).Features(features).ExpectedTotalAmount(expectedTotalAmount).Execute()

Active Subscription Import(Allows repetition imports)

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
	externalSubscriptionId := "externalSubscriptionId_example" // string | Required, The external id of subscription
	externalPlanId := "externalPlanId_example" // string | The external id of plan, plan should created at UniBee at first
	quantity := int64(789) // int64 | the quantity of plan, default 1 if not provided 
	countryCode := "countryCode_example" // string | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription.
	gateway := "gateway_example" // string | Required, should one of stripe|paypal|wire_transfer|changelly 
	currentPeriodStart := "currentPeriodStart_example" // string | Required, UTC time, the current period start time of subscription, format '2006-01-02 15:04:05'
	currentPeriodEnd := "currentPeriodEnd_example" // string | Required, UTC time, the current period end time of subscription, format '2006-01-02 15:04:05'
	billingCycleAnchor := "billingCycleAnchor_example" // string | Required, UTC time, The reference point that aligns future billing cycle dates. It sets the day of week for week intervals, the day of month for month and year intervals, and the month of year for year intervals, format '2006-01-02 15:04:05'
	createTime := "createTime_example" // string | Required, UTC time, the creation time of subscription, format '2006-01-02 15:04:05'
	email := "email_example" // string | The email of user, one of Email or ExternalUserId is required (optional)
	externalUserId := "externalUserId_example" // string | The external id of user, one of Email or ExternalUserId is required  (optional)
	vatNumber := "vatNumber_example" // string | The Vat Number of user (optional)
	taxPercentage := int64(789) // int64 | The TaxPercentage of user, valid when system vat gateway enabled, ，1000 = 10% (optional)
	firstPaidTime := "firstPaidTime_example" // string | UTC time, the first payment success time of subscription, format '2006-01-02 15:04:05' (optional)
	features := "features_example" // string | In json format, additional features data of subscription, will join user's metric data in user api if provided' (optional)
	expectedTotalAmount := int64(789) // int64 | Optional. Unit: cents. If greater than 0, the system will verify the calculated total amount against this value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImport.SubscriptionActiveSubscriptionImportGet(context.Background()).ExternalSubscriptionId(externalSubscriptionId).ExternalPlanId(externalPlanId).Quantity(quantity).CountryCode(countryCode).Gateway(gateway).CurrentPeriodStart(currentPeriodStart).CurrentPeriodEnd(currentPeriodEnd).BillingCycleAnchor(billingCycleAnchor).CreateTime(createTime).Email(email).ExternalUserId(externalUserId).VatNumber(vatNumber).TaxPercentage(taxPercentage).FirstPaidTime(firstPaidTime).Features(features).ExpectedTotalAmount(expectedTotalAmount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImport.SubscriptionActiveSubscriptionImportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionActiveSubscriptionImportGet`: MerchantSubscriptionActiveSubscriptionImportGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImport.SubscriptionActiveSubscriptionImportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionActiveSubscriptionImportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalSubscriptionId** | **string** | Required, The external id of subscription | 
 **externalPlanId** | **string** | The external id of plan, plan should created at UniBee at first | 
 **quantity** | **int64** | the quantity of plan, default 1 if not provided  | 
 **countryCode** | **string** | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription. | 
 **gateway** | **string** | Required, should one of stripe|paypal|wire_transfer|changelly  | 
 **currentPeriodStart** | **string** | Required, UTC time, the current period start time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **currentPeriodEnd** | **string** | Required, UTC time, the current period end time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **billingCycleAnchor** | **string** | Required, UTC time, The reference point that aligns future billing cycle dates. It sets the day of week for week intervals, the day of month for month and year intervals, and the month of year for year intervals, format &#39;2006-01-02 15:04:05&#39; | 
 **createTime** | **string** | Required, UTC time, the creation time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **email** | **string** | The email of user, one of Email or ExternalUserId is required | 
 **externalUserId** | **string** | The external id of user, one of Email or ExternalUserId is required  | 
 **vatNumber** | **string** | The Vat Number of user | 
 **taxPercentage** | **int64** | The TaxPercentage of user, valid when system vat gateway enabled, ，1000 &#x3D; 10% | 
 **firstPaidTime** | **string** | UTC time, the first payment success time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **features** | **string** | In json format, additional features data of subscription, will join user&#39;s metric data in user api if provided&#39; | 
 **expectedTotalAmount** | **int64** | Optional. Unit: cents. If greater than 0, the system will verify the calculated total amount against this value | 

### Return type

[**MerchantSubscriptionActiveSubscriptionImportGet200Response**](MerchantSubscriptionActiveSubscriptionImportGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionHistorySubscriptionImportGet

> MerchantSubscriptionActiveSubscriptionImportGet200Response SubscriptionHistorySubscriptionImportGet(ctx).ExternalSubscriptionId(externalSubscriptionId).ExternalPlanId(externalPlanId).Quantity(quantity).CountryCode(countryCode).Gateway(gateway).CurrentPeriodStart(currentPeriodStart).CurrentPeriodEnd(currentPeriodEnd).TotalAmount(totalAmount).Email(email).ExternalUserId(externalUserId).TaxPercentage(taxPercentage).Execute()

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
	externalSubscriptionId := "externalSubscriptionId_example" // string | Required, The external id of subscription
	externalPlanId := "externalPlanId_example" // string | The external id of plan, plan should created at first
	quantity := int64(789) // int64 | the quantity of plan, default 1 if not provided 
	countryCode := "countryCode_example" // string | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription.
	gateway := "gateway_example" // string | Required, should one of stripe|paypal|wire_transfer|changelly 
	currentPeriodStart := "currentPeriodStart_example" // string | Required, UTC time, the current period start time of subscription, format '2006-01-02 15:04:05'
	currentPeriodEnd := "currentPeriodEnd_example" // string | Required, UTC time, the current period end time of subscription, format '2006-01-02 15:04:05'
	totalAmount := int64(789) // int64 | Required. Unit: cents.
	email := "email_example" // string | The email of user, one of Email or ExternalUserId is required (optional)
	externalUserId := "externalUserId_example" // string | The external id of user, one of Email or ExternalUserId is required  (optional)
	taxPercentage := int64(789) // int64 | The TaxPercentage of subscription, valid when system vat gateway enabled, 1000 = 10% (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImport.SubscriptionHistorySubscriptionImportGet(context.Background()).ExternalSubscriptionId(externalSubscriptionId).ExternalPlanId(externalPlanId).Quantity(quantity).CountryCode(countryCode).Gateway(gateway).CurrentPeriodStart(currentPeriodStart).CurrentPeriodEnd(currentPeriodEnd).TotalAmount(totalAmount).Email(email).ExternalUserId(externalUserId).TaxPercentage(taxPercentage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImport.SubscriptionHistorySubscriptionImportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionHistorySubscriptionImportGet`: MerchantSubscriptionActiveSubscriptionImportGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImport.SubscriptionHistorySubscriptionImportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionHistorySubscriptionImportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalSubscriptionId** | **string** | Required, The external id of subscription | 
 **externalPlanId** | **string** | The external id of plan, plan should created at first | 
 **quantity** | **int64** | the quantity of plan, default 1 if not provided  | 
 **countryCode** | **string** | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription. | 
 **gateway** | **string** | Required, should one of stripe|paypal|wire_transfer|changelly  | 
 **currentPeriodStart** | **string** | Required, UTC time, the current period start time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **currentPeriodEnd** | **string** | Required, UTC time, the current period end time of subscription, format &#39;2006-01-02 15:04:05&#39; | 
 **totalAmount** | **int64** | Required. Unit: cents. | 
 **email** | **string** | The email of user, one of Email or ExternalUserId is required | 
 **externalUserId** | **string** | The external id of user, one of Email or ExternalUserId is required  | 
 **taxPercentage** | **int64** | The TaxPercentage of subscription, valid when system vat gateway enabled, 1000 &#x3D; 10% | 

### Return type

[**MerchantSubscriptionActiveSubscriptionImportGet200Response**](MerchantSubscriptionActiveSubscriptionImportGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

