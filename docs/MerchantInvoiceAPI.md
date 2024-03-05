# \MerchantInvoiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantInvoiceCancelPost**](MerchantInvoiceAPI.md#MerchantInvoiceCancelPost) | **Post** /merchant/invoice/cancel | Admin Cancel Invoice Of Processing Status
[**MerchantInvoiceDeletePost**](MerchantInvoiceAPI.md#MerchantInvoiceDeletePost) | **Post** /merchant/invoice/delete | Admin Delete Invoice Of Pending Status
[**MerchantInvoiceDetailGet**](MerchantInvoiceAPI.md#MerchantInvoiceDetailGet) | **Get** /merchant/invoice/detail | Invoice Detail
[**MerchantInvoiceDetailPost**](MerchantInvoiceAPI.md#MerchantInvoiceDetailPost) | **Post** /merchant/invoice/detail | Invoice Detail
[**MerchantInvoiceEditPost**](MerchantInvoiceAPI.md#MerchantInvoiceEditPost) | **Post** /merchant/invoice/edit | Admin Edit Invoice
[**MerchantInvoiceFinishPost**](MerchantInvoiceAPI.md#MerchantInvoiceFinishPost) | **Post** /merchant/invoice/finish | Admin Finish Invoice，Generate Pay Link
[**MerchantInvoiceListGet**](MerchantInvoiceAPI.md#MerchantInvoiceListGet) | **Get** /merchant/invoice/list | Invoice List
[**MerchantInvoiceListPost**](MerchantInvoiceAPI.md#MerchantInvoiceListPost) | **Post** /merchant/invoice/list | Invoice List
[**MerchantInvoiceNewPost**](MerchantInvoiceAPI.md#MerchantInvoiceNewPost) | **Post** /merchant/invoice/new | Admin Create New Invoice
[**MerchantInvoicePdfGeneratePost**](MerchantInvoiceAPI.md#MerchantInvoicePdfGeneratePost) | **Post** /merchant/invoice/pdf_generate | Admin Generate Merchant Invoice pdf
[**MerchantInvoiceRefundPost**](MerchantInvoiceAPI.md#MerchantInvoiceRefundPost) | **Post** /merchant/invoice/refund | Admin Create Refund From Invoice
[**MerchantInvoiceSendEmailPost**](MerchantInvoiceAPI.md#MerchantInvoiceSendEmailPost) | **Post** /merchant/invoice/send_email | Admin Send Merchant Invoice Email to User



## MerchantInvoiceCancelPost

> MerchantAuthSsoLoginOTPPost200Response MerchantInvoiceCancelPost(ctx).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()

Admin Cancel Invoice Of Processing Status

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
	unibeeApiMerchantInvoiceCancelReq := *openapiclient.NewUnibeeApiMerchantInvoiceCancelReq() // UnibeeApiMerchantInvoiceCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceCancelPost(context.Background()).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceCancelReq** | [**UnibeeApiMerchantInvoiceCancelReq**](UnibeeApiMerchantInvoiceCancelReq.md) |  | 

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


## MerchantInvoiceDeletePost

> MerchantAuthSsoLoginOTPPost200Response MerchantInvoiceDeletePost(ctx).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()

Admin Delete Invoice Of Pending Status

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
	unibeeApiMerchantInvoiceDeleteReq := *openapiclient.NewUnibeeApiMerchantInvoiceDeleteReq("InvoiceId_example") // UnibeeApiMerchantInvoiceDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceDeletePost(context.Background()).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceDeleteReq** | [**UnibeeApiMerchantInvoiceDeleteReq**](UnibeeApiMerchantInvoiceDeleteReq.md) |  | 

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


## MerchantInvoiceDetailGet

> MerchantInvoiceDetailGet200Response MerchantInvoiceDetailGet(ctx).InvoiceId(invoiceId).Execute()

Invoice Detail

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
	invoiceId := "invoiceId_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceDetailGet(context.Background()).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceDetailGet`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string** | Invoice ID | 

### Return type

[**MerchantInvoiceDetailGet200Response**](MerchantInvoiceDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceDetailPost

> MerchantInvoiceDetailGet200Response MerchantInvoiceDetailPost(ctx).UnibeeApiMerchantInvoiceDetailReq(unibeeApiMerchantInvoiceDetailReq).Execute()

Invoice Detail

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
	unibeeApiMerchantInvoiceDetailReq := *openapiclient.NewUnibeeApiMerchantInvoiceDetailReq("InvoiceId_example") // UnibeeApiMerchantInvoiceDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceDetailPost(context.Background()).UnibeeApiMerchantInvoiceDetailReq(unibeeApiMerchantInvoiceDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceDetailPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceDetailReq** | [**UnibeeApiMerchantInvoiceDetailReq**](UnibeeApiMerchantInvoiceDetailReq.md) |  | 

### Return type

[**MerchantInvoiceDetailGet200Response**](MerchantInvoiceDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceEditPost

> MerchantInvoiceDetailGet200Response MerchantInvoiceEditPost(ctx).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()

Admin Edit Invoice

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
	unibeeApiMerchantInvoiceEditReq := *openapiclient.NewUnibeeApiMerchantInvoiceEditReq("InvoiceId_example") // UnibeeApiMerchantInvoiceEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceEditPost(context.Background()).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceEditPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceEditReq** | [**UnibeeApiMerchantInvoiceEditReq**](UnibeeApiMerchantInvoiceEditReq.md) |  | 

### Return type

[**MerchantInvoiceDetailGet200Response**](MerchantInvoiceDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceFinishPost

> MerchantInvoiceFinishPost200Response MerchantInvoiceFinishPost(ctx).UnibeeApiMerchantInvoiceFinishReq(unibeeApiMerchantInvoiceFinishReq).Execute()

Admin Finish Invoice，Generate Pay Link

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
	unibeeApiMerchantInvoiceFinishReq := *openapiclient.NewUnibeeApiMerchantInvoiceFinishReq(int32(123), "InvoiceId_example", int32(123)) // UnibeeApiMerchantInvoiceFinishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceFinishPost(context.Background()).UnibeeApiMerchantInvoiceFinishReq(unibeeApiMerchantInvoiceFinishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceFinishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceFinishPost`: MerchantInvoiceFinishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceFinishPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceFinishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceFinishReq** | [**UnibeeApiMerchantInvoiceFinishReq**](UnibeeApiMerchantInvoiceFinishReq.md) |  | 

### Return type

[**MerchantInvoiceFinishPost200Response**](MerchantInvoiceFinishPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceListGet

> MerchantInvoiceListGet200Response MerchantInvoiceListGet(ctx).FirstName(firstName).LastName(lastName).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Page(page).Count(count).Execute()

Invoice List

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
	firstName := "firstName_example" // string | FirstName (optional)
	lastName := "lastName_example" // string | LastName (optional)
	currency := "currency_example" // string | Currency (optional)
	status := []int32{int32(123)} // []int32 | Status, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled (optional)
	amountStart := int64(789) // int64 | AmountStart (optional)
	amountEnd := int64(789) // int64 | AmountEnd (optional)
	userId := int32(56) // int32 | UserId Filter, Default Filter All (optional)
	sendEmail := "sendEmail_example" // string | SendEmail Filter , Default Filter All (optional)
	sortField := "sortField_example" // string | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort，asc|desc，Default desc (optional)
	deleteInclude := true // bool | Deleted Involved，Need Admin (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count By Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceListGet(context.Background()).FirstName(firstName).LastName(lastName).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceListGet`: MerchantInvoiceListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstName** | **string** | FirstName | 
 **lastName** | **string** | LastName | 
 **currency** | **string** | Currency | 
 **status** | **[]int32** | Status, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | 
 **amountStart** | **int64** | AmountStart | 
 **amountEnd** | **int64** | AmountEnd | 
 **userId** | **int32** | UserId Filter, Default Filter All | 
 **sendEmail** | **string** | SendEmail Filter , Default Filter All | 
 **sortField** | **string** | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | 
 **sortType** | **string** | Sort，asc|desc，Default desc | 
 **deleteInclude** | **bool** | Deleted Involved，Need Admin | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count By Page | 

### Return type

[**MerchantInvoiceListGet200Response**](MerchantInvoiceListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceListPost

> MerchantInvoiceListGet200Response MerchantInvoiceListPost(ctx).UnibeeApiMerchantInvoiceListReq(unibeeApiMerchantInvoiceListReq).Execute()

Invoice List

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
	unibeeApiMerchantInvoiceListReq := *openapiclient.NewUnibeeApiMerchantInvoiceListReq() // UnibeeApiMerchantInvoiceListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceListPost(context.Background()).UnibeeApiMerchantInvoiceListReq(unibeeApiMerchantInvoiceListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceListPost`: MerchantInvoiceListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceListReq** | [**UnibeeApiMerchantInvoiceListReq**](UnibeeApiMerchantInvoiceListReq.md) |  | 

### Return type

[**MerchantInvoiceListGet200Response**](MerchantInvoiceListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceNewPost

> MerchantInvoiceDetailGet200Response MerchantInvoiceNewPost(ctx).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()

Admin Create New Invoice

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
	unibeeApiMerchantInvoiceNewReq := *openapiclient.NewUnibeeApiMerchantInvoiceNewReq("Currency_example", int64(123), int64(123), int64(123)) // UnibeeApiMerchantInvoiceNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceNewPost(context.Background()).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceNewPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceNewReq** | [**UnibeeApiMerchantInvoiceNewReq**](UnibeeApiMerchantInvoiceNewReq.md) |  | 

### Return type

[**MerchantInvoiceDetailGet200Response**](MerchantInvoiceDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoicePdfGeneratePost

> MerchantAuthSsoLoginOTPPost200Response MerchantInvoicePdfGeneratePost(ctx).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()

Admin Generate Merchant Invoice pdf

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
	unibeeApiMerchantInvoicePdfGenerateReq := *openapiclient.NewUnibeeApiMerchantInvoicePdfGenerateReq("InvoiceId_example") // UnibeeApiMerchantInvoicePdfGenerateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoicePdfGeneratePost(context.Background()).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoicePdfGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoicePdfGeneratePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoicePdfGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoicePdfGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoicePdfGenerateReq** | [**UnibeeApiMerchantInvoicePdfGenerateReq**](UnibeeApiMerchantInvoicePdfGenerateReq.md) |  | 

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


## MerchantInvoiceRefundPost

> MerchantInvoiceRefundPost200Response MerchantInvoiceRefundPost(ctx).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()

Admin Create Refund From Invoice

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
	unibeeApiMerchantInvoiceRefundReq := *openapiclient.NewUnibeeApiMerchantInvoiceRefundReq("InvoiceId_example", "Reason_example", int64(123), "RefundNo_example") // UnibeeApiMerchantInvoiceRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceRefundPost(context.Background()).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceRefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceRefundPost`: MerchantInvoiceRefundPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceRefundPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceRefundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceRefundReq** | [**UnibeeApiMerchantInvoiceRefundReq**](UnibeeApiMerchantInvoiceRefundReq.md) |  | 

### Return type

[**MerchantInvoiceRefundPost200Response**](MerchantInvoiceRefundPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantInvoiceSendEmailPost

> MerchantAuthSsoLoginOTPPost200Response MerchantInvoiceSendEmailPost(ctx).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()

Admin Send Merchant Invoice Email to User

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
	unibeeApiMerchantInvoiceSendEmailReq := *openapiclient.NewUnibeeApiMerchantInvoiceSendEmailReq("InvoiceId_example") // UnibeeApiMerchantInvoiceSendEmailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInvoiceAPI.MerchantInvoiceSendEmailPost(context.Background()).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInvoiceAPI.MerchantInvoiceSendEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantInvoiceSendEmailPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInvoiceAPI.MerchantInvoiceSendEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantInvoiceSendEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceSendEmailReq** | [**UnibeeApiMerchantInvoiceSendEmailReq**](UnibeeApiMerchantInvoiceSendEmailReq.md) |  | 

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

