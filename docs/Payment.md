# \Payment

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentCancelPost**](Payment.md#PaymentCancelPost) | **Post** /merchant/payment/cancel | Cancel Payment
[**PaymentCapturePost**](Payment.md#PaymentCapturePost) | **Post** /merchant/payment/capture | Capture Payment
[**PaymentDetailGet**](Payment.md#PaymentDetailGet) | **Get** /merchant/payment/detail | Query Payment Detail
[**PaymentListGet**](Payment.md#PaymentListGet) | **Get** /merchant/payment/list | Query Payment List
[**PaymentMethodGetGet**](Payment.md#PaymentMethodGetGet) | **Get** /merchant/payment/method_get | Query Payment Method
[**PaymentMethodListGet**](Payment.md#PaymentMethodListGet) | **Get** /merchant/payment/method_list | Query Payment Method List
[**PaymentMethodNewPost**](Payment.md#PaymentMethodNewPost) | **Post** /merchant/payment/method_new | Create New Payment Method And Attach To User
[**PaymentNewPost**](Payment.md#PaymentNewPost) | **Post** /merchant/payment/new | New Payment
[**PaymentRefundCancelPost**](Payment.md#PaymentRefundCancelPost) | **Post** /merchant/payment/refund/cancel | Cancel Payment Refund
[**PaymentRefundDetailGet**](Payment.md#PaymentRefundDetailGet) | **Get** /merchant/payment/refund/detail | Query Payment Refund Detail
[**PaymentRefundListGet**](Payment.md#PaymentRefundListGet) | **Get** /merchant/payment/refund/list | Query Payment Refund List
[**PaymentRefundNewPost**](Payment.md#PaymentRefundNewPost) | **Post** /merchant/payment/refund/new | New Payment Refund
[**PaymentTimelineListGet**](Payment.md#PaymentTimelineListGet) | **Get** /merchant/payment/timeline/list | Payment TimeLine List



## PaymentCancelPost

> MerchantAuthSsoLoginOTPPost200Response PaymentCancelPost(ctx).UnibeeApiMerchantPaymentCancelReq(unibeeApiMerchantPaymentCancelReq).Execute()

Cancel Payment

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
	unibeeApiMerchantPaymentCancelReq := *openapiclient.NewUnibeeApiMerchantPaymentCancelReq("PaymentId_example") // UnibeeApiMerchantPaymentCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentCancelPost(context.Background()).UnibeeApiMerchantPaymentCancelReq(unibeeApiMerchantPaymentCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentCancelReq** | [**UnibeeApiMerchantPaymentCancelReq**](UnibeeApiMerchantPaymentCancelReq.md) |  | 

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


## PaymentCapturePost

> MerchantAuthSsoLoginOTPPost200Response PaymentCapturePost(ctx).UnibeeApiMerchantPaymentCaptureReq(unibeeApiMerchantPaymentCaptureReq).Execute()

Capture Payment

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
	unibeeApiMerchantPaymentCaptureReq := *openapiclient.NewUnibeeApiMerchantPaymentCaptureReq(int64(123), "Currency_example", "ExternalCaptureId_example", "PaymentId_example") // UnibeeApiMerchantPaymentCaptureReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentCapturePost(context.Background()).UnibeeApiMerchantPaymentCaptureReq(unibeeApiMerchantPaymentCaptureReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentCapturePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentCapturePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentCapturePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCapturePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentCaptureReq** | [**UnibeeApiMerchantPaymentCaptureReq**](UnibeeApiMerchantPaymentCaptureReq.md) |  | 

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


## PaymentDetailGet

> MerchantPaymentDetailGet200Response PaymentDetailGet(ctx).PaymentId(paymentId).Execute()

Query Payment Detail

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
	paymentId := "paymentId_example" // string | PaymentId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentDetailGet(context.Background()).PaymentId(paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentDetailGet`: MerchantPaymentDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string** | PaymentId | 

### Return type

[**MerchantPaymentDetailGet200Response**](MerchantPaymentDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentListGet

> MerchantPaymentListGet200Response PaymentListGet(ctx).GatewayId(gatewayId).UserId(userId).Email(email).Status(status).Currency(currency).CountryCode(countryCode).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Query Payment List

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
	gatewayId := int64(789) // int64 | GatewayId (optional)
	userId := int64(789) // int64 | UserId  (optional)
	email := "email_example" // string | Email (optional)
	status := int32(56) // int32 | Status, 10-Created|20-Success|30-Failed|40-Cancelled (optional)
	currency := "currency_example" // string | Currency (optional)
	countryCode := "countryCode_example" // string | CountryCode (optional)
	sortField := "sortField_example" // string | Sort Field，user_id|create_time|status (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentListGet(context.Background()).GatewayId(gatewayId).UserId(userId).Email(email).Status(status).Currency(currency).CountryCode(countryCode).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentListGet`: MerchantPaymentListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64** | GatewayId | 
 **userId** | **int64** | UserId  | 
 **email** | **string** | Email | 
 **status** | **int32** | Status, 10-Created|20-Success|30-Failed|40-Cancelled | 
 **currency** | **string** | Currency | 
 **countryCode** | **string** | CountryCode | 
 **sortField** | **string** | Sort Field，user_id|create_time|status | 
 **sortType** | **string** | Sort Type，asc|desc | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantPaymentListGet200Response**](MerchantPaymentListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodGetGet

> MerchantPaymentMethodGetGet200Response PaymentMethodGetGet(ctx).GatewayId(gatewayId).UserId(userId).PaymentMethodId(paymentMethodId).Execute()

Query Payment Method

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
	gatewayId := int64(789) // int64 | GatewayId
	userId := int64(789) // int64 | UserId
	paymentMethodId := "paymentMethodId_example" // string | PaymentMethodId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentMethodGetGet(context.Background()).GatewayId(gatewayId).UserId(userId).PaymentMethodId(paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentMethodGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodGetGet`: MerchantPaymentMethodGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentMethodGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64** | GatewayId | 
 **userId** | **int64** | UserId | 
 **paymentMethodId** | **string** | PaymentMethodId | 

### Return type

[**MerchantPaymentMethodGetGet200Response**](MerchantPaymentMethodGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodListGet

> MerchantPaymentMethodListGet200Response PaymentMethodListGet(ctx).GatewayId(gatewayId).UserId(userId).PaymentId(paymentId).Execute()

Query Payment Method List

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
	gatewayId := int64(789) // int64 | GatewayId
	userId := int64(789) // int64 | UserId (optional)
	paymentId := "paymentId_example" // string | PaymentId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentMethodListGet(context.Background()).GatewayId(gatewayId).UserId(userId).PaymentId(paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentMethodListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodListGet`: MerchantPaymentMethodListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentMethodListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64** | GatewayId | 
 **userId** | **int64** | UserId | 
 **paymentId** | **string** | PaymentId | 

### Return type

[**MerchantPaymentMethodListGet200Response**](MerchantPaymentMethodListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodNewPost

> MerchantPaymentMethodNewPost200Response PaymentMethodNewPost(ctx).UnibeeApiMerchantPaymentMethodNewReq(unibeeApiMerchantPaymentMethodNewReq).Execute()

Create New Payment Method And Attach To User

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
	unibeeApiMerchantPaymentMethodNewReq := *openapiclient.NewUnibeeApiMerchantPaymentMethodNewReq("Currency_example", int64(123), int64(123)) // UnibeeApiMerchantPaymentMethodNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentMethodNewPost(context.Background()).UnibeeApiMerchantPaymentMethodNewReq(unibeeApiMerchantPaymentMethodNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentMethodNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodNewPost`: MerchantPaymentMethodNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentMethodNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentMethodNewReq** | [**UnibeeApiMerchantPaymentMethodNewReq**](UnibeeApiMerchantPaymentMethodNewReq.md) |  | 

### Return type

[**MerchantPaymentMethodNewPost200Response**](MerchantPaymentMethodNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentNewPost

> MerchantPaymentNewPost200Response PaymentNewPost(ctx).UnibeeApiMerchantPaymentNewReq(unibeeApiMerchantPaymentNewReq).Execute()

New Payment

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
	unibeeApiMerchantPaymentNewReq := *openapiclient.NewUnibeeApiMerchantPaymentNewReq(int64(123)) // UnibeeApiMerchantPaymentNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentNewPost(context.Background()).UnibeeApiMerchantPaymentNewReq(unibeeApiMerchantPaymentNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentNewPost`: MerchantPaymentNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentNewReq** | [**UnibeeApiMerchantPaymentNewReq**](UnibeeApiMerchantPaymentNewReq.md) |  | 

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


## PaymentRefundCancelPost

> MerchantAuthSsoLoginOTPPost200Response PaymentRefundCancelPost(ctx).UnibeeApiMerchantPaymentRefundCancelReq(unibeeApiMerchantPaymentRefundCancelReq).Execute()

Cancel Payment Refund

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
	unibeeApiMerchantPaymentRefundCancelReq := *openapiclient.NewUnibeeApiMerchantPaymentRefundCancelReq("RefundId_example") // UnibeeApiMerchantPaymentRefundCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentRefundCancelPost(context.Background()).UnibeeApiMerchantPaymentRefundCancelReq(unibeeApiMerchantPaymentRefundCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentRefundCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRefundCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentRefundCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRefundCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentRefundCancelReq** | [**UnibeeApiMerchantPaymentRefundCancelReq**](UnibeeApiMerchantPaymentRefundCancelReq.md) |  | 

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


## PaymentRefundDetailGet

> MerchantPaymentRefundDetailGet200Response PaymentRefundDetailGet(ctx).RefundId(refundId).Execute()

Query Payment Refund Detail

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
	refundId := "refundId_example" // string | RefundId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentRefundDetailGet(context.Background()).RefundId(refundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentRefundDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRefundDetailGet`: MerchantPaymentRefundDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentRefundDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRefundDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundId** | **string** | RefundId | 

### Return type

[**MerchantPaymentRefundDetailGet200Response**](MerchantPaymentRefundDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRefundListGet

> MerchantPaymentRefundListGet200Response PaymentRefundListGet(ctx).PaymentId(paymentId).Status(status).GatewayId(gatewayId).UserId(userId).Email(email).Currency(currency).Execute()

Query Payment Refund List

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
	paymentId := "paymentId_example" // string | PaymentId
	status := int32(56) // int32 | Status,10-create|20-success|30-Failed|40-Reverse (optional)
	gatewayId := int64(789) // int64 | GatewayId (optional)
	userId := int64(789) // int64 | UserId (optional)
	email := "email_example" // string | Email (optional)
	currency := "currency_example" // string | Currency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentRefundListGet(context.Background()).PaymentId(paymentId).Status(status).GatewayId(gatewayId).UserId(userId).Email(email).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentRefundListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRefundListGet`: MerchantPaymentRefundListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentRefundListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRefundListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string** | PaymentId | 
 **status** | **int32** | Status,10-create|20-success|30-Failed|40-Reverse | 
 **gatewayId** | **int64** | GatewayId | 
 **userId** | **int64** | UserId | 
 **email** | **string** | Email | 
 **currency** | **string** | Currency | 

### Return type

[**MerchantPaymentRefundListGet200Response**](MerchantPaymentRefundListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRefundNewPost

> MerchantPaymentRefundNewPost200Response PaymentRefundNewPost(ctx).UnibeeApiMerchantPaymentNewPaymentRefundReq(unibeeApiMerchantPaymentNewPaymentRefundReq).Execute()

New Payment Refund

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
	unibeeApiMerchantPaymentNewPaymentRefundReq := *openapiclient.NewUnibeeApiMerchantPaymentNewPaymentRefundReq("Currency_example", "ExternalRefundId_example", "PaymentId_example", int64(123)) // UnibeeApiMerchantPaymentNewPaymentRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentRefundNewPost(context.Background()).UnibeeApiMerchantPaymentNewPaymentRefundReq(unibeeApiMerchantPaymentNewPaymentRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentRefundNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRefundNewPost`: MerchantPaymentRefundNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentRefundNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRefundNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentNewPaymentRefundReq** | [**UnibeeApiMerchantPaymentNewPaymentRefundReq**](UnibeeApiMerchantPaymentNewPaymentRefundReq.md) |  | 

### Return type

[**MerchantPaymentRefundNewPost200Response**](MerchantPaymentRefundNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentTimelineListGet

> MerchantPaymentTimelineListGet200Response PaymentTimelineListGet(ctx).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Payment TimeLine List

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
	userId := int64(789) // int64 | Filter UserId, Default All (optional)
	sortField := "sortField_example" // string | Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Payment.PaymentTimelineListGet(context.Background()).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Payment.PaymentTimelineListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentTimelineListGet`: MerchantPaymentTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Payment.PaymentTimelineListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentTimelineListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | Filter UserId, Default All | 
 **sortField** | **string** | Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantPaymentTimelineListGet200Response**](MerchantPaymentTimelineListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

