# \Invoice

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceCancelPost**](Invoice.md#InvoiceCancelPost) | **Post** /merchant/invoice/cancel | Admin Cancel Invoice Of Processing Status
[**InvoiceDeletePost**](Invoice.md#InvoiceDeletePost) | **Post** /merchant/invoice/delete | DeletePendingInvoice
[**InvoiceDetailGet**](Invoice.md#InvoiceDetailGet) | **Get** /merchant/invoice/detail | InvoiceDetail
[**InvoiceDetailPost**](Invoice.md#InvoiceDetailPost) | **Post** /merchant/invoice/detail | InvoiceDetail
[**InvoiceEditPost**](Invoice.md#InvoiceEditPost) | **Post** /merchant/invoice/edit | InvoiceEdit
[**InvoiceFinishPost**](Invoice.md#InvoiceFinishPost) | **Post** /merchant/invoice/finish | FinishInvoice
[**InvoiceListGet**](Invoice.md#InvoiceListGet) | **Get** /merchant/invoice/list | InvoiceList
[**InvoiceListPost**](Invoice.md#InvoiceListPost) | **Post** /merchant/invoice/list | InvoiceList
[**InvoiceMarkRefundPost**](Invoice.md#InvoiceMarkRefundPost) | **Post** /merchant/invoice/mark_refund | MarkInvoiceRefund
[**InvoiceNewPost**](Invoice.md#InvoiceNewPost) | **Post** /merchant/invoice/new | NewInvoice
[**InvoicePdfGeneratePost**](Invoice.md#InvoicePdfGeneratePost) | **Post** /merchant/invoice/pdf_generate | GenerateInvoicePDF
[**InvoiceReconvertCryptoAndSendEmailPost**](Invoice.md#InvoiceReconvertCryptoAndSendEmailPost) | **Post** /merchant/invoice/reconvert_crypto_and_send_email | Admin Reconvert Crypto Data and Send Invoice Email to User
[**InvoiceRefundPost**](Invoice.md#InvoiceRefundPost) | **Post** /merchant/invoice/refund | CreateInvoiceRefund
[**InvoiceSendEmailPost**](Invoice.md#InvoiceSendEmailPost) | **Post** /merchant/invoice/send_email | SendInvoiceEmail



## InvoiceCancelPost

> MerchantAuthSsoLoginOTPPost200Response InvoiceCancelPost(ctx).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()

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
	unibeeApiMerchantInvoiceCancelReq := *openapiclient.NewUnibeeApiMerchantInvoiceCancelReq("InvoiceId_example") // UnibeeApiMerchantInvoiceCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceCancelPost(context.Background()).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceCancelPostRequest struct via the builder pattern


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


## InvoiceDeletePost

> MerchantAuthSsoLoginOTPPost200Response InvoiceDeletePost(ctx).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()

DeletePendingInvoice



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
	resp, r, err := apiClient.Invoice.InvoiceDeletePost(context.Background()).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceDeletePostRequest struct via the builder pattern


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


## InvoiceDetailGet

> MerchantInvoiceDetailGet200Response InvoiceDetailGet(ctx).InvoiceId(invoiceId).Execute()

InvoiceDetail



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
	invoiceId := "invoiceId_example" // string | The unique id of invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceDetailGet(context.Background()).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceDetailGet`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string** | The unique id of invoice | 

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


## InvoiceDetailPost

> MerchantInvoiceDetailGet200Response InvoiceDetailPost(ctx).UnibeeApiMerchantInvoiceDetailReq(unibeeApiMerchantInvoiceDetailReq).Execute()

InvoiceDetail



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
	resp, r, err := apiClient.Invoice.InvoiceDetailPost(context.Background()).UnibeeApiMerchantInvoiceDetailReq(unibeeApiMerchantInvoiceDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceDetailPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceDetailPostRequest struct via the builder pattern


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


## InvoiceEditPost

> MerchantInvoiceDetailGet200Response InvoiceEditPost(ctx).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()

InvoiceEdit



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
	resp, r, err := apiClient.Invoice.InvoiceEditPost(context.Background()).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceEditPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceEditPostRequest struct via the builder pattern


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


## InvoiceFinishPost

> MerchantInvoiceFinishPost200Response InvoiceFinishPost(ctx).UnibeeApiMerchantInvoiceFinishReq(unibeeApiMerchantInvoiceFinishReq).Execute()

FinishInvoice



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
	unibeeApiMerchantInvoiceFinishReq := *openapiclient.NewUnibeeApiMerchantInvoiceFinishReq(int32(123), "InvoiceId_example") // UnibeeApiMerchantInvoiceFinishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceFinishPost(context.Background()).UnibeeApiMerchantInvoiceFinishReq(unibeeApiMerchantInvoiceFinishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceFinishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceFinishPost`: MerchantInvoiceFinishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceFinishPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceFinishPostRequest struct via the builder pattern


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


## InvoiceListGet

> MerchantInvoiceListGet200Response InvoiceListGet(ctx).FirstName(firstName).LastName(lastName).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Page(page).Count(count).Execute()

InvoiceList



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
	firstName := "firstName_example" // string | The firstName of invoice (optional)
	lastName := "lastName_example" // string | The lastName of invoice (optional)
	currency := "currency_example" // string | The currency of invoice (optional)
	status := []int32{int32(123)} // []int32 | The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled (optional)
	amountStart := int64(789) // int64 | The filter start amount of invoice (optional)
	amountEnd := int64(789) // int64 | The filter end amount of invoice (optional)
	userId := int64(789) // int64 | The filter userid of invoice (optional)
	sendEmail := "sendEmail_example" // string | The filter email of invoice (optional)
	sortField := "sortField_example" // string | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort，asc|desc，Default desc (optional)
	deleteInclude := true // bool | Deleted Involved，Need Admin Permission (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count By Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceListGet(context.Background()).FirstName(firstName).LastName(lastName).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceListGet`: MerchantInvoiceListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstName** | **string** | The firstName of invoice | 
 **lastName** | **string** | The lastName of invoice | 
 **currency** | **string** | The currency of invoice | 
 **status** | **[]int32** | The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | 
 **amountStart** | **int64** | The filter start amount of invoice | 
 **amountEnd** | **int64** | The filter end amount of invoice | 
 **userId** | **int64** | The filter userid of invoice | 
 **sendEmail** | **string** | The filter email of invoice | 
 **sortField** | **string** | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | 
 **sortType** | **string** | Sort，asc|desc，Default desc | 
 **deleteInclude** | **bool** | Deleted Involved，Need Admin Permission | 
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


## InvoiceListPost

> MerchantInvoiceListGet200Response InvoiceListPost(ctx).UnibeeApiMerchantInvoiceListReq(unibeeApiMerchantInvoiceListReq).Execute()

InvoiceList



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
	resp, r, err := apiClient.Invoice.InvoiceListPost(context.Background()).UnibeeApiMerchantInvoiceListReq(unibeeApiMerchantInvoiceListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceListPost`: MerchantInvoiceListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceListPostRequest struct via the builder pattern


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


## InvoiceMarkRefundPost

> MerchantInvoiceMarkRefundPost200Response InvoiceMarkRefundPost(ctx).UnibeeApiMerchantInvoiceMarkRefundReq(unibeeApiMerchantInvoiceMarkRefundReq).Execute()

MarkInvoiceRefund



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
	unibeeApiMerchantInvoiceMarkRefundReq := *openapiclient.NewUnibeeApiMerchantInvoiceMarkRefundReq("InvoiceId_example", "Reason_example", int64(123)) // UnibeeApiMerchantInvoiceMarkRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceMarkRefundPost(context.Background()).UnibeeApiMerchantInvoiceMarkRefundReq(unibeeApiMerchantInvoiceMarkRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceMarkRefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceMarkRefundPost`: MerchantInvoiceMarkRefundPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceMarkRefundPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceMarkRefundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceMarkRefundReq** | [**UnibeeApiMerchantInvoiceMarkRefundReq**](UnibeeApiMerchantInvoiceMarkRefundReq.md) |  | 

### Return type

[**MerchantInvoiceMarkRefundPost200Response**](MerchantInvoiceMarkRefundPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceNewPost

> MerchantInvoiceDetailGet200Response InvoiceNewPost(ctx).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()

NewInvoice

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
	resp, r, err := apiClient.Invoice.InvoiceNewPost(context.Background()).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceNewPost`: MerchantInvoiceDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceNewPostRequest struct via the builder pattern


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


## InvoicePdfGeneratePost

> MerchantAuthSsoLoginOTPPost200Response InvoicePdfGeneratePost(ctx).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()

GenerateInvoicePDF

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
	resp, r, err := apiClient.Invoice.InvoicePdfGeneratePost(context.Background()).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoicePdfGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicePdfGeneratePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoicePdfGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicePdfGeneratePostRequest struct via the builder pattern


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


## InvoiceReconvertCryptoAndSendEmailPost

> MerchantAuthSsoLoginOTPPost200Response InvoiceReconvertCryptoAndSendEmailPost(ctx).UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(unibeeApiMerchantInvoiceReconvertCryptoAndSendReq).Execute()

Admin Reconvert Crypto Data and Send Invoice Email to User

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
	unibeeApiMerchantInvoiceReconvertCryptoAndSendReq := *openapiclient.NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq("InvoiceId_example") // UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceReconvertCryptoAndSendEmailPost(context.Background()).UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(unibeeApiMerchantInvoiceReconvertCryptoAndSendReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceReconvertCryptoAndSendEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceReconvertCryptoAndSendEmailPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceReconvertCryptoAndSendEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceReconvertCryptoAndSendEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceReconvertCryptoAndSendReq** | [**UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq**](UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq.md) |  | 

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


## InvoiceRefundPost

> MerchantInvoiceMarkRefundPost200Response InvoiceRefundPost(ctx).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()

CreateInvoiceRefund



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
	unibeeApiMerchantInvoiceRefundReq := *openapiclient.NewUnibeeApiMerchantInvoiceRefundReq("InvoiceId_example", "Reason_example", int64(123)) // UnibeeApiMerchantInvoiceRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceRefundPost(context.Background()).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceRefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceRefundPost`: MerchantInvoiceMarkRefundPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceRefundPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceRefundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceRefundReq** | [**UnibeeApiMerchantInvoiceRefundReq**](UnibeeApiMerchantInvoiceRefundReq.md) |  | 

### Return type

[**MerchantInvoiceMarkRefundPost200Response**](MerchantInvoiceMarkRefundPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSendEmailPost

> MerchantAuthSsoLoginOTPPost200Response InvoiceSendEmailPost(ctx).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()

SendInvoiceEmail

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
	resp, r, err := apiClient.Invoice.InvoiceSendEmailPost(context.Background()).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceSendEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceSendEmailPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceSendEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceSendEmailPostRequest struct via the builder pattern


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

