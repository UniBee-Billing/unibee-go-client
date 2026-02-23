# \Subscription

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionAddNewTrialStartPost**](Subscription.md#SubscriptionAddNewTrialStartPost) | **Post** /merchant/subscription/add_new_trial_start | Append Subscription TrialEnd
[**SubscriptionApplySubscriptionNextInvoicePost**](Subscription.md#SubscriptionApplySubscriptionNextInvoicePost) | **Post** /merchant/subscription/apply_subscription_next_invoice | Apply Discount Or Premo Credit To Next Invoice
[**SubscriptionCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_at_period_end | Cancel Subscription At Period End
[**SubscriptionCancelLastCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelLastCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_last_cancel_at_period_end | Cancel Last Cancel Subscription At Period End
[**SubscriptionCancelPost**](Subscription.md#SubscriptionCancelPost) | **Post** /merchant/subscription/cancel | Cancel Subscription Immediately
[**SubscriptionChangeCurrencyPost**](Subscription.md#SubscriptionChangeCurrencyPost) | **Post** /merchant/subscription/change_currency | Change Subscription Currency
[**SubscriptionChangeDueDayPost**](Subscription.md#SubscriptionChangeDueDayPost) | **Post** /merchant/subscription/change_due_day | Change Subscription Due Day
[**SubscriptionChangeGatewayPost**](Subscription.md#SubscriptionChangeGatewayPost) | **Post** /merchant/subscription/change_gateway | Change Subscription Gateway
[**SubscriptionCreatePreviewPost**](Subscription.md#SubscriptionCreatePreviewPost) | **Post** /merchant/subscription/create_preview | Create Subscription Preview
[**SubscriptionCreateSubmitPost**](Subscription.md#SubscriptionCreateSubmitPost) | **Post** /merchant/subscription/create_submit | Create Subscription
[**SubscriptionDetailGet**](Subscription.md#SubscriptionDetailGet) | **Get** /merchant/subscription/detail | Subscription Detail
[**SubscriptionDetailPost**](Subscription.md#SubscriptionDetailPost) | **Post** /merchant/subscription/detail | Subscription Detail
[**SubscriptionListGet**](Subscription.md#SubscriptionListGet) | **Get** /merchant/subscription/list | Get Subscription List
[**SubscriptionListPost**](Subscription.md#SubscriptionListPost) | **Post** /merchant/subscription/list | Get Subscription List
[**SubscriptionPreviewSubscriptionNextInvoiceGet**](Subscription.md#SubscriptionPreviewSubscriptionNextInvoiceGet) | **Get** /merchant/subscription/preview_subscription_next_invoice | Subscription Next Invoice Preview
[**SubscriptionUserPendingCryptoSubscriptionDetailGet**](Subscription.md#SubscriptionUserPendingCryptoSubscriptionDetailGet) | **Get** /merchant/subscription/user_pending_crypto_subscription_detail | User Pending Crypto Subscription Detail
[**SubscriptionUserPendingCryptoSubscriptionDetailPost**](Subscription.md#SubscriptionUserPendingCryptoSubscriptionDetailPost) | **Post** /merchant/subscription/user_pending_crypto_subscription_detail | User Pending Crypto Subscription Detail
[**SubscriptionUserSubscriptionDetailGet**](Subscription.md#SubscriptionUserSubscriptionDetailGet) | **Get** /merchant/subscription/user_subscription_detail | User Subscription Detail
[**SubscriptionUserSubscriptionDetailPost**](Subscription.md#SubscriptionUserSubscriptionDetailPost) | **Post** /merchant/subscription/user_subscription_detail | User Subscription Detail



## SubscriptionAddNewTrialStartPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionAddNewTrialStartPost(ctx).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()

Append Subscription TrialEnd

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
	unibeeApiMerchantSubscriptionAddNewTrialStartReq := *openapiclient.NewUnibeeApiMerchantSubscriptionAddNewTrialStartReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionAddNewTrialStartReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionAddNewTrialStartPost(context.Background()).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionAddNewTrialStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAddNewTrialStartPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionAddNewTrialStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAddNewTrialStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionAddNewTrialStartReq** | [**UnibeeApiMerchantSubscriptionAddNewTrialStartReq**](UnibeeApiMerchantSubscriptionAddNewTrialStartReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionApplySubscriptionNextInvoicePost

> MerchantSubscriptionApplySubscriptionNextInvoicePost200Response SubscriptionApplySubscriptionNextInvoicePost(ctx).UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq(unibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq).Execute()

Apply Discount Or Premo Credit To Next Invoice

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
	unibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq := *openapiclient.NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq() // UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionApplySubscriptionNextInvoicePost(context.Background()).UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq(unibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionApplySubscriptionNextInvoicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionApplySubscriptionNextInvoicePost`: MerchantSubscriptionApplySubscriptionNextInvoicePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionApplySubscriptionNextInvoicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionApplySubscriptionNextInvoicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq** | [**UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq**](UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq.md) |  | 

### Return type

[**MerchantSubscriptionApplySubscriptionNextInvoicePost200Response**](MerchantSubscriptionApplySubscriptionNextInvoicePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelAtPeriodEndPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()

Cancel Subscription At Period End



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
	unibeeApiMerchantSubscriptionCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq() // UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelAtPeriodEndPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelLastCancelAtPeriodEndPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionCancelLastCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()

Cancel Last Cancel Subscription At Period End



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
	unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq() // UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelLastCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelLastCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelLastCancelAtPeriodEndPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelLastCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelLastCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionCancelPost(ctx).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()

Cancel Subscription Immediately



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
	unibeeApiMerchantSubscriptionCancelReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelReq() // UnibeeApiMerchantSubscriptionCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelPost(context.Background()).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelReq** | [**UnibeeApiMerchantSubscriptionCancelReq**](UnibeeApiMerchantSubscriptionCancelReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionChangeCurrencyPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionChangeCurrencyPost(ctx).UnibeeApiMerchantSubscriptionChangeCurrencyReq(unibeeApiMerchantSubscriptionChangeCurrencyReq).Execute()

Change Subscription Currency



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
	unibeeApiMerchantSubscriptionChangeCurrencyReq := *openapiclient.NewUnibeeApiMerchantSubscriptionChangeCurrencyReq("Currency_example", "SubscriptionId_example") // UnibeeApiMerchantSubscriptionChangeCurrencyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionChangeCurrencyPost(context.Background()).UnibeeApiMerchantSubscriptionChangeCurrencyReq(unibeeApiMerchantSubscriptionChangeCurrencyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionChangeCurrencyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionChangeCurrencyPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionChangeCurrencyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionChangeCurrencyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionChangeCurrencyReq** | [**UnibeeApiMerchantSubscriptionChangeCurrencyReq**](UnibeeApiMerchantSubscriptionChangeCurrencyReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionChangeDueDayPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionChangeDueDayPost(ctx).UnibeeApiMerchantSubscriptionChangeDueDayReq(unibeeApiMerchantSubscriptionChangeDueDayReq).Execute()

Change Subscription Due Day



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
	unibeeApiMerchantSubscriptionChangeDueDayReq := *openapiclient.NewUnibeeApiMerchantSubscriptionChangeDueDayReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionChangeDueDayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionChangeDueDayPost(context.Background()).UnibeeApiMerchantSubscriptionChangeDueDayReq(unibeeApiMerchantSubscriptionChangeDueDayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionChangeDueDayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionChangeDueDayPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionChangeDueDayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionChangeDueDayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionChangeDueDayReq** | [**UnibeeApiMerchantSubscriptionChangeDueDayReq**](UnibeeApiMerchantSubscriptionChangeDueDayReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionChangeGatewayPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionChangeGatewayPost(ctx).UnibeeApiMerchantSubscriptionChangeGatewayReq(unibeeApiMerchantSubscriptionChangeGatewayReq).Execute()

Change Subscription Gateway

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
	unibeeApiMerchantSubscriptionChangeGatewayReq := *openapiclient.NewUnibeeApiMerchantSubscriptionChangeGatewayReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionChangeGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionChangeGatewayPost(context.Background()).UnibeeApiMerchantSubscriptionChangeGatewayReq(unibeeApiMerchantSubscriptionChangeGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionChangeGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionChangeGatewayPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionChangeGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionChangeGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionChangeGatewayReq** | [**UnibeeApiMerchantSubscriptionChangeGatewayReq**](UnibeeApiMerchantSubscriptionChangeGatewayReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreatePreviewPost

> MerchantSubscriptionCreatePreviewPost200Response SubscriptionCreatePreviewPost(ctx).UnibeeApiMerchantSubscriptionCreatePreviewReq(unibeeApiMerchantSubscriptionCreatePreviewReq).Execute()

Create Subscription Preview

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
	unibeeApiMerchantSubscriptionCreatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCreatePreviewReq(int64(123)) // UnibeeApiMerchantSubscriptionCreatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCreatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionCreatePreviewReq(unibeeApiMerchantSubscriptionCreatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCreatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreatePreviewPost`: MerchantSubscriptionCreatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCreatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreatePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCreatePreviewReq** | [**UnibeeApiMerchantSubscriptionCreatePreviewReq**](UnibeeApiMerchantSubscriptionCreatePreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionCreatePreviewPost200Response**](MerchantSubscriptionCreatePreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreateSubmitPost

> MerchantSubscriptionCreateSubmitPost200Response SubscriptionCreateSubmitPost(ctx).UnibeeApiMerchantSubscriptionCreateReq(unibeeApiMerchantSubscriptionCreateReq).Execute()

Create Subscription

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
	unibeeApiMerchantSubscriptionCreateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCreateReq(int64(123)) // UnibeeApiMerchantSubscriptionCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCreateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionCreateReq(unibeeApiMerchantSubscriptionCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCreateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreateSubmitPost`: MerchantSubscriptionCreateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCreateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreateSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCreateReq** | [**UnibeeApiMerchantSubscriptionCreateReq**](UnibeeApiMerchantSubscriptionCreateReq.md) |  | 

### Return type

[**MerchantSubscriptionCreateSubmitPost200Response**](MerchantSubscriptionCreateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDetailGet

> MerchantSubscriptionDetailGet200Response SubscriptionDetailGet(ctx).SubscriptionId(subscriptionId).Execute()

Subscription Detail

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
	subscriptionId := "subscriptionId_example" // string | SubscriptionId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionDetailGet(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionDetailGet`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDetailPost

> MerchantSubscriptionDetailGet200Response SubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()

Subscription Detail

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
	unibeeApiMerchantSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionDetailReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionDetailPost`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionListGet

> MerchantSubscriptionListGet200Response SubscriptionListGet(ctx).UserId(userId).ExternalUserId(externalUserId).SearchKey(searchKey).Email(email).Status(status).Currency(currency).PlanIds(planIds).ProductIds(productIds).AmountStart(amountStart).AmountEnd(amountEnd).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Subscription List

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
	userId := int64(789) // int64 | UserId (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId (optional)
	searchKey := "searchKey_example" // string | Search SubscriptionId|Email (optional)
	email := "email_example" // string | The filter email of subscription user (optional)
	status := []int32{int32(123)} // []int32 | Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed (optional)
	currency := "currency_example" // string | The currency of subscription (optional)
	planIds := []int64{int64(123)} // []int64 | The filter ids of plan (optional)
	productIds := []int64{int64(123)} // []int64 | The filter ids of product, invalid if planIds is provided (optional)
	amountStart := int32(56) // int32 | The filter start amount of subscription (optional)
	amountEnd := int32(56) // int32 | The filter end amount of subscription (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionListGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).SearchKey(searchKey).Email(email).Status(status).Currency(currency).PlanIds(planIds).ProductIds(productIds).AmountStart(amountStart).AmountEnd(amountEnd).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionListGet`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **externalUserId** | **string** | ExternalUserId | 
 **searchKey** | **string** | Search SubscriptionId|Email | 
 **email** | **string** | The filter email of subscription user | 
 **status** | **[]int32** | Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | 
 **currency** | **string** | The currency of subscription | 
 **planIds** | **[]int64** | The filter ids of plan | 
 **productIds** | **[]int64** | The filter ids of product, invalid if planIds is provided | 
 **amountStart** | **int32** | The filter start amount of subscription | 
 **amountEnd** | **int32** | The filter end amount of subscription | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionListPost

> MerchantSubscriptionListGet200Response SubscriptionListPost(ctx).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()

Get Subscription List

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
	unibeeApiMerchantSubscriptionListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionListReq() // UnibeeApiMerchantSubscriptionListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionListPost(context.Background()).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionListPost`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionListReq** | [**UnibeeApiMerchantSubscriptionListReq**](UnibeeApiMerchantSubscriptionListReq.md) |  | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPreviewSubscriptionNextInvoiceGet

> MerchantSubscriptionApplySubscriptionNextInvoicePost200Response SubscriptionPreviewSubscriptionNextInvoiceGet(ctx).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).SubscriptionId(subscriptionId).Execute()

Subscription Next Invoice Preview

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
	userId := int64(789) // int64 | UserId (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, unique, either ExternalUserId&Email or UserId needed if subscriptionId not specified (optional)
	productId := int64(789) // int64 | default product will use if productId not specified and subscriptionId is blank (optional)
	subscriptionId := "subscriptionId_example" // string | SubscriptionId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionPreviewSubscriptionNextInvoiceGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionPreviewSubscriptionNextInvoiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPreviewSubscriptionNextInvoiceGet`: MerchantSubscriptionApplySubscriptionNextInvoicePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionPreviewSubscriptionNextInvoiceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPreviewSubscriptionNextInvoiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **externalUserId** | **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed if subscriptionId not specified | 
 **productId** | **int64** | default product will use if productId not specified and subscriptionId is blank | 
 **subscriptionId** | **string** | SubscriptionId | 

### Return type

[**MerchantSubscriptionApplySubscriptionNextInvoicePost200Response**](MerchantSubscriptionApplySubscriptionNextInvoicePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserPendingCryptoSubscriptionDetailGet

> MerchantSubscriptionActiveSubscriptionImportPost200Response SubscriptionUserPendingCryptoSubscriptionDetailGet(ctx).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()

User Pending Crypto Subscription Detail

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
	userId := int64(789) // int64 | UserId (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, unique, either ExternalUserId&Email or UserId needed (optional)
	productId := int64(789) // int64 | default product will use if productId not specified and subscriptionId is blank (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserPendingCryptoSubscriptionDetailGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserPendingCryptoSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserPendingCryptoSubscriptionDetailGet`: MerchantSubscriptionActiveSubscriptionImportPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserPendingCryptoSubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserPendingCryptoSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **externalUserId** | **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **productId** | **int64** | default product will use if productId not specified and subscriptionId is blank | 

### Return type

[**MerchantSubscriptionActiveSubscriptionImportPost200Response**](MerchantSubscriptionActiveSubscriptionImportPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserPendingCryptoSubscriptionDetailPost

> MerchantSubscriptionActiveSubscriptionImportPost200Response SubscriptionUserPendingCryptoSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq).Execute()

User Pending Crypto Subscription Detail

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
	unibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq() // UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserPendingCryptoSubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserPendingCryptoSubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserPendingCryptoSubscriptionDetailPost`: MerchantSubscriptionActiveSubscriptionImportPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserPendingCryptoSubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserPendingCryptoSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq.md) |  | 

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


## SubscriptionUserSubscriptionDetailGet

> MerchantSubscriptionUserSubscriptionDetailGet200Response SubscriptionUserSubscriptionDetailGet(ctx).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()

User Subscription Detail

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
	userId := int64(789) // int64 | UserId (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, unique, either ExternalUserId&Email or UserId needed (optional)
	productId := int64(789) // int64 | default product will use if productId not specified and subscriptionId is blank (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserSubscriptionDetailGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserSubscriptionDetailGet`: MerchantSubscriptionUserSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserSubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **externalUserId** | **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **productId** | **int64** | default product will use if productId not specified and subscriptionId is blank | 

### Return type

[**MerchantSubscriptionUserSubscriptionDetailGet200Response**](MerchantSubscriptionUserSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserSubscriptionDetailPost

> MerchantSubscriptionUserSubscriptionDetailGet200Response SubscriptionUserSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()

User Subscription Detail

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
	unibeeApiMerchantSubscriptionUserSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq() // UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserSubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserSubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserSubscriptionDetailPost`: MerchantSubscriptionUserSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserSubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUserSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionUserSubscriptionDetailGet200Response**](MerchantSubscriptionUserSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

