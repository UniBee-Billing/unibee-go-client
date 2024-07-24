# \Subscription

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionActiveTemporarilyPost**](Subscription.md#SubscriptionActiveTemporarilyPost) | **Post** /merchant/subscription/active_temporarily | SubscriptionActiveTemporarily
[**SubscriptionAddNewTrialStartPost**](Subscription.md#SubscriptionAddNewTrialStartPost) | **Post** /merchant/subscription/add_new_trial_start | AppendSubscriptionTrialEnd
[**SubscriptionCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_at_period_end | CancelSubscriptionAtPeriodEnd
[**SubscriptionCancelLastCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelLastCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_last_cancel_at_period_end | CancelLastCancelSubscriptionAtPeriodEnd
[**SubscriptionCancelPost**](Subscription.md#SubscriptionCancelPost) | **Post** /merchant/subscription/cancel | CancelSubscriptionImmediately
[**SubscriptionChangeGatewayPost**](Subscription.md#SubscriptionChangeGatewayPost) | **Post** /merchant/subscription/change_gateway | ChangeSubscriptionGateway
[**SubscriptionConfigGet**](Subscription.md#SubscriptionConfigGet) | **Get** /merchant/subscription/config | SubscriptionConfig
[**SubscriptionConfigUpdatePost**](Subscription.md#SubscriptionConfigUpdatePost) | **Post** /merchant/subscription/config/update | Update Merchant Subscription Config
[**SubscriptionCreatePreviewPost**](Subscription.md#SubscriptionCreatePreviewPost) | **Post** /merchant/subscription/create_preview | CreateSubscriptionPreview
[**SubscriptionCreateSubmitPost**](Subscription.md#SubscriptionCreateSubmitPost) | **Post** /merchant/subscription/create_submit | CreateSubscription
[**SubscriptionDetailGet**](Subscription.md#SubscriptionDetailGet) | **Get** /merchant/subscription/detail | SubscriptionDetail
[**SubscriptionDetailPost**](Subscription.md#SubscriptionDetailPost) | **Post** /merchant/subscription/detail | SubscriptionDetail
[**SubscriptionListGet**](Subscription.md#SubscriptionListGet) | **Get** /merchant/subscription/list | SubscriptionList
[**SubscriptionListPost**](Subscription.md#SubscriptionListPost) | **Post** /merchant/subscription/list | SubscriptionList
[**SubscriptionNewOnetimeAddonPaymentPost**](Subscription.md#SubscriptionNewOnetimeAddonPaymentPost) | **Post** /merchant/subscription/new_onetime_addon_payment | NewSubscriptionOnetimeAddonPayment
[**SubscriptionOnetimeAddonListGet**](Subscription.md#SubscriptionOnetimeAddonListGet) | **Get** /merchant/subscription/onetime_addon_list | SubscriptionOnetimeAddonList
[**SubscriptionPaymentNewPost**](Subscription.md#SubscriptionPaymentNewPost) | **Post** /merchant/subscription/payment/new | NewSubscriptionPayment
[**SubscriptionRenewPost**](Subscription.md#SubscriptionRenewPost) | **Post** /merchant/subscription/renew | RenewSubscription
[**SubscriptionResumePost**](Subscription.md#SubscriptionResumePost) | **Post** /merchant/subscription/resume | Merchant Edit Subscription-Resume
[**SubscriptionSuspendPost**](Subscription.md#SubscriptionSuspendPost) | **Post** /merchant/subscription/suspend | Merchant Edit Subscription-Stop
[**SubscriptionUpdatePreviewPost**](Subscription.md#SubscriptionUpdatePreviewPost) | **Post** /merchant/subscription/update_preview | UpdateSubscriptionPreview
[**SubscriptionUpdateSubmitPost**](Subscription.md#SubscriptionUpdateSubmitPost) | **Post** /merchant/subscription/update_submit | UpdateSubscription
[**SubscriptionUserPendingCryptoSubscriptionDetailGet**](Subscription.md#SubscriptionUserPendingCryptoSubscriptionDetailGet) | **Get** /merchant/subscription/user_pending_crypto_subscription_detail | UserPendingCryptoSubscriptionDetail
[**SubscriptionUserPendingCryptoSubscriptionDetailPost**](Subscription.md#SubscriptionUserPendingCryptoSubscriptionDetailPost) | **Post** /merchant/subscription/user_pending_crypto_subscription_detail | UserPendingCryptoSubscriptionDetail
[**SubscriptionUserSubscriptionDetailGet**](Subscription.md#SubscriptionUserSubscriptionDetailGet) | **Get** /merchant/subscription/user_subscription_detail | UserSubscriptionDetail
[**SubscriptionUserSubscriptionDetailPost**](Subscription.md#SubscriptionUserSubscriptionDetailPost) | **Post** /merchant/subscription/user_subscription_detail | UserSubscriptionDetail



## SubscriptionActiveTemporarilyPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionActiveTemporarilyPost(ctx).UnibeeApiMerchantSubscriptionActiveTemporarilyReq(unibeeApiMerchantSubscriptionActiveTemporarilyReq).Execute()

SubscriptionActiveTemporarily



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
	unibeeApiMerchantSubscriptionActiveTemporarilyReq := *openapiclient.NewUnibeeApiMerchantSubscriptionActiveTemporarilyReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionActiveTemporarilyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionActiveTemporarilyPost(context.Background()).UnibeeApiMerchantSubscriptionActiveTemporarilyReq(unibeeApiMerchantSubscriptionActiveTemporarilyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionActiveTemporarilyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionActiveTemporarilyPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionActiveTemporarilyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionActiveTemporarilyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionActiveTemporarilyReq** | [**UnibeeApiMerchantSubscriptionActiveTemporarilyReq**](UnibeeApiMerchantSubscriptionActiveTemporarilyReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionAddNewTrialStartPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionAddNewTrialStartPost(ctx).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()

AppendSubscriptionTrialEnd

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
	// response from `SubscriptionAddNewTrialStartPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()

CancelSubscriptionAtPeriodEnd



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
	// response from `SubscriptionCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelLastCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelLastCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()

CancelLastCancelSubscriptionAtPeriodEnd



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
	// response from `SubscriptionCancelLastCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelPost(ctx).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()

CancelSubscriptionImmediately



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
	// response from `SubscriptionCancelPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionChangeGatewayPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionChangeGatewayPost(ctx).UnibeeApiMerchantSubscriptionChangeGatewayReq(unibeeApiMerchantSubscriptionChangeGatewayReq).Execute()

ChangeSubscriptionGateway

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
	// response from `SubscriptionChangeGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionConfigGet

> MerchantSubscriptionConfigGet200Response SubscriptionConfigGet(ctx).Execute()

SubscriptionConfig

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigGet`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigGetRequest struct via the builder pattern


### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionConfigUpdatePost

> MerchantSubscriptionConfigGet200Response SubscriptionConfigUpdatePost(ctx).UnibeeApiMerchantSubscriptionConfigUpdateReq(unibeeApiMerchantSubscriptionConfigUpdateReq).Execute()

Update Merchant Subscription Config

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
	unibeeApiMerchantSubscriptionConfigUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionConfigUpdateReq() // UnibeeApiMerchantSubscriptionConfigUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionConfigUpdatePost(context.Background()).UnibeeApiMerchantSubscriptionConfigUpdateReq(unibeeApiMerchantSubscriptionConfigUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigUpdatePost`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionConfigUpdateReq** | [**UnibeeApiMerchantSubscriptionConfigUpdateReq**](UnibeeApiMerchantSubscriptionConfigUpdateReq.md) |  | 

### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

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

CreateSubscriptionPreview

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

CreateSubscription

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

SubscriptionDetail

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

SubscriptionDetail

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

> MerchantSubscriptionListGet200Response SubscriptionListGet(ctx).UserId(userId).Status(status).Currency(currency).PlanIds(planIds).AmountStart(amountStart).AmountEnd(amountEnd).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

SubscriptionList

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
	status := []int32{int32(123)} // []int32 | Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed (optional)
	currency := "currency_example" // string | The currency of subscription (optional)
	planIds := []int32{int32(123)} // []int32 | The filter ids of plan (optional)
	amountStart := int32(56) // int32 | The filter start amount of subscription (optional)
	amountEnd := int32(56) // int32 | The filter end amount of subscription (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionListGet(context.Background()).UserId(userId).Status(status).Currency(currency).PlanIds(planIds).AmountStart(amountStart).AmountEnd(amountEnd).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
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
 **status** | **[]int32** | Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | 
 **currency** | **string** | The currency of subscription | 
 **planIds** | **[]int32** | The filter ids of plan | 
 **amountStart** | **int32** | The filter start amount of subscription | 
 **amountEnd** | **int32** | The filter end amount of subscription | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 
 **createTimeStart** | **int64** | CreateTimeStart | 
 **createTimeEnd** | **int64** | CreateTimeEnd | 

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

SubscriptionList

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


## SubscriptionNewOnetimeAddonPaymentPost

> MerchantSubscriptionNewOnetimeAddonPaymentPost200Response SubscriptionNewOnetimeAddonPaymentPost(ctx).UnibeeApiMerchantSubscriptionOnetimeAddonNewReq(unibeeApiMerchantSubscriptionOnetimeAddonNewReq).Execute()

NewSubscriptionOnetimeAddonPayment



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
	unibeeApiMerchantSubscriptionOnetimeAddonNewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq(int64(123), int64(123)) // UnibeeApiMerchantSubscriptionOnetimeAddonNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionNewOnetimeAddonPaymentPost(context.Background()).UnibeeApiMerchantSubscriptionOnetimeAddonNewReq(unibeeApiMerchantSubscriptionOnetimeAddonNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionNewOnetimeAddonPaymentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionNewOnetimeAddonPaymentPost`: MerchantSubscriptionNewOnetimeAddonPaymentPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionNewOnetimeAddonPaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionNewOnetimeAddonPaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionOnetimeAddonNewReq** | [**UnibeeApiMerchantSubscriptionOnetimeAddonNewReq**](UnibeeApiMerchantSubscriptionOnetimeAddonNewReq.md) |  | 

### Return type

[**MerchantSubscriptionNewOnetimeAddonPaymentPost200Response**](MerchantSubscriptionNewOnetimeAddonPaymentPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOnetimeAddonListGet

> MerchantSubscriptionOnetimeAddonListGet200Response SubscriptionOnetimeAddonListGet(ctx).UserId(userId).Page(page).Count(count).Execute()

SubscriptionOnetimeAddonList

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
	userId := int64(789) // int64 | UserId
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionOnetimeAddonListGet(context.Background()).UserId(userId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionOnetimeAddonListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOnetimeAddonListGet`: MerchantSubscriptionOnetimeAddonListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionOnetimeAddonListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOnetimeAddonListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionOnetimeAddonListGet200Response**](MerchantSubscriptionOnetimeAddonListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPaymentNewPost

> MerchantPaymentNewPost200Response SubscriptionPaymentNewPost(ctx).UnibeeApiMerchantSubscriptionNewPaymentReq(unibeeApiMerchantSubscriptionNewPaymentReq).Execute()

NewSubscriptionPayment

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
	unibeeApiMerchantSubscriptionNewPaymentReq := *openapiclient.NewUnibeeApiMerchantSubscriptionNewPaymentReq(int64(123)) // UnibeeApiMerchantSubscriptionNewPaymentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionPaymentNewPost(context.Background()).UnibeeApiMerchantSubscriptionNewPaymentReq(unibeeApiMerchantSubscriptionNewPaymentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionPaymentNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPaymentNewPost`: MerchantPaymentNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionPaymentNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPaymentNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionNewPaymentReq** | [**UnibeeApiMerchantSubscriptionNewPaymentReq**](UnibeeApiMerchantSubscriptionNewPaymentReq.md) |  | 

### Return type

[**MerchantPaymentNewPost200Response**](MerchantPaymentNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionRenewPost

> MerchantSubscriptionCreateSubmitPost200Response SubscriptionRenewPost(ctx).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()

RenewSubscription



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
	unibeeApiMerchantSubscriptionRenewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionRenewReq() // UnibeeApiMerchantSubscriptionRenewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionRenewPost(context.Background()).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionRenewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRenewPost`: MerchantSubscriptionCreateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionRenewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionRenewReq** | [**UnibeeApiMerchantSubscriptionRenewReq**](UnibeeApiMerchantSubscriptionRenewReq.md) |  | 

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


## SubscriptionResumePost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionResumePost(ctx).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()

Merchant Edit Subscription-Resume

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
	unibeeApiMerchantSubscriptionResumeReq := *openapiclient.NewUnibeeApiMerchantSubscriptionResumeReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionResumeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionResumePost(context.Background()).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionResumePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionResumePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionResumeReq** | [**UnibeeApiMerchantSubscriptionResumeReq**](UnibeeApiMerchantSubscriptionResumeReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionSuspendPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionSuspendPost(ctx).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()

Merchant Edit Subscription-Stop

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
	unibeeApiMerchantSubscriptionSuspendReq := *openapiclient.NewUnibeeApiMerchantSubscriptionSuspendReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionSuspendReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionSuspendPost(context.Background()).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionSuspendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionSuspendPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionSuspendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionSuspendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionSuspendReq** | [**UnibeeApiMerchantSubscriptionSuspendReq**](UnibeeApiMerchantSubscriptionSuspendReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdatePreviewPost

> MerchantSubscriptionUpdatePreviewPost200Response SubscriptionUpdatePreviewPost(ctx).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()

UpdateSubscriptionPreview

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
	unibeeApiMerchantSubscriptionUpdatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUpdatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUpdatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdatePreviewPost`: MerchantSubscriptionUpdatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUpdatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdatePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdatePreviewReq** | [**UnibeeApiMerchantSubscriptionUpdatePreviewReq**](UnibeeApiMerchantSubscriptionUpdatePreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdatePreviewPost200Response**](MerchantSubscriptionUpdatePreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdateSubmitPost

> MerchantSubscriptionUpdateSubmitPost200Response SubscriptionUpdateSubmitPost(ctx).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()

UpdateSubscription

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
	unibeeApiMerchantSubscriptionUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdateReq(int64(123), int64(123)) // UnibeeApiMerchantSubscriptionUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUpdateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUpdateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdateSubmitPost`: MerchantSubscriptionUpdateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUpdateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdateReq** | [**UnibeeApiMerchantSubscriptionUpdateReq**](UnibeeApiMerchantSubscriptionUpdateReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdateSubmitPost200Response**](MerchantSubscriptionUpdateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserPendingCryptoSubscriptionDetailGet

> MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response SubscriptionUserPendingCryptoSubscriptionDetailGet(ctx).UserId(userId).ExternalUserId(externalUserId).Execute()

UserPendingCryptoSubscriptionDetail

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserPendingCryptoSubscriptionDetailGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserPendingCryptoSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserPendingCryptoSubscriptionDetailGet`: MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response
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

### Return type

[**MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response**](MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserPendingCryptoSubscriptionDetailPost

> MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response SubscriptionUserPendingCryptoSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq).Execute()

UserPendingCryptoSubscriptionDetail

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
	// response from `SubscriptionUserPendingCryptoSubscriptionDetailPost`: MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response
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

[**MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response**](MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserSubscriptionDetailGet

> MerchantSubscriptionUserSubscriptionDetailGet200Response SubscriptionUserSubscriptionDetailGet(ctx).UserId(userId).ExternalUserId(externalUserId).Execute()

UserSubscriptionDetail

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserSubscriptionDetailGet(context.Background()).UserId(userId).ExternalUserId(externalUserId).Execute()
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

UserSubscriptionDetail

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

