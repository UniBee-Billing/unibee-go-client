# \Invoice

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceCancelPost**](Invoice.md#InvoiceCancelPost) | **Post** /merchant/invoice/cancel | Admin Cancel Invoice
[**InvoiceClearPaymentPost**](Invoice.md#InvoiceClearPaymentPost) | **Post** /merchant/invoice/clearPayment | Admin Clear Invoice Current Pending Payment
[**InvoiceCreateMarkRefundPost**](Invoice.md#InvoiceCreateMarkRefundPost) | **Post** /merchant/invoice/create_mark_refund | Create And Mark Manual Refund (Crypto/WireTransfer/External)
[**InvoiceCreditNoteListPost**](Invoice.md#InvoiceCreditNoteListPost) | **Post** /merchant/invoice/credit_note/list | Bulk CreditNote(Refund Invoice) Invoice List
[**InvoiceDeletePost**](Invoice.md#InvoiceDeletePost) | **Post** /merchant/invoice/delete | Delete Pending Invoice
[**InvoiceDetailGet**](Invoice.md#InvoiceDetailGet) | **Get** /merchant/invoice/detail | Invoice Detail
[**InvoiceDetailPost**](Invoice.md#InvoiceDetailPost) | **Post** /merchant/invoice/detail | Invoice Detail
[**InvoiceEditPost**](Invoice.md#InvoiceEditPost) | **Post** /merchant/invoice/edit | Invoice Edit
[**InvoiceExternalGatewayInvoiceMarkPaidPost**](Invoice.md#InvoiceExternalGatewayInvoiceMarkPaidPost) | **Post** /merchant/invoice/external_gateway_invoice/mark_paid | Mark External Gateway Invoice As Paid
[**InvoiceExternalGatewayInvoiceMarkRefundedPost**](Invoice.md#InvoiceExternalGatewayInvoiceMarkRefundedPost) | **Post** /merchant/invoice/external_gateway_invoice/mark_refunded | Mark External Gateway Invoice As Refunded
[**InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost**](Invoice.md#InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost) | **Post** /merchant/invoice/external_gateway_invoice/update_payment_link | Update External Gateway Invoice Payment Link
[**InvoiceFinishPost**](Invoice.md#InvoiceFinishPost) | **Post** /merchant/invoice/finish | Finish Invoice
[**InvoiceListGet**](Invoice.md#InvoiceListGet) | **Get** /merchant/invoice/list | Get Invoice List
[**InvoiceListPost**](Invoice.md#InvoiceListPost) | **Post** /merchant/invoice/list | Get Invoice List
[**InvoiceMarkRefundSuccessPost**](Invoice.md#InvoiceMarkRefundSuccessPost) | **Post** /merchant/invoice/mark_refund_success | Mark Manual Invoice Refund As Success
[**InvoiceMarkWireTransferSuccessPost**](Invoice.md#InvoiceMarkWireTransferSuccessPost) | **Post** /merchant/invoice/mark_wire_transfer_success | Mark Wire Transfer Invoice As Paid (Manual)
[**InvoiceNewPost**](Invoice.md#InvoiceNewPost) | **Post** /merchant/invoice/new | New Invoice
[**InvoicePdfGeneratePost**](Invoice.md#InvoicePdfGeneratePost) | **Post** /merchant/invoice/pdf_generate | Generate Invoice PDF
[**InvoicePdfUpdatePost**](Invoice.md#InvoicePdfUpdatePost) | **Post** /merchant/invoice/pdf_update | Update Invoice PDF
[**InvoiceReconvertCryptoAndSendEmailPost**](Invoice.md#InvoiceReconvertCryptoAndSendEmailPost) | **Post** /merchant/invoice/reconvert_crypto_and_send_email | Admin Reconvert Crypto Data and Send Invoice Email
[**InvoiceRefundPost**](Invoice.md#InvoiceRefundPost) | **Post** /merchant/invoice/refund | Create InvoiceRefund
[**InvoiceSendEmailPost**](Invoice.md#InvoiceSendEmailPost) | **Post** /merchant/invoice/send_email | Send Invoice Email
[**InvoiceSplitPaymentsGet**](Invoice.md#InvoiceSplitPaymentsGet) | **Get** /merchant/invoice/split_payments | Get Split Payments
[**InvoiceSplitPaymentsPost**](Invoice.md#InvoiceSplitPaymentsPost) | **Post** /merchant/invoice/split_payments | Get Split Payments
[**InvoiceSyncUserSnapshotPost**](Invoice.md#InvoiceSyncUserSnapshotPost) | **Post** /merchant/invoice/sync_user_snapshot | Sync Invoice User Snapshot With Latest User Info



## InvoiceCancelPost

> MerchantAuthSsoClearTotpPost200Response InvoiceCancelPost(ctx).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()

Admin Cancel Invoice

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
	unibeeApiMerchantInvoiceCancelReq := *openapiclient.NewUnibeeApiMerchantInvoiceCancelReq("InvoiceId_example") // UnibeeApiMerchantInvoiceCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceCancelPost(context.Background()).UnibeeApiMerchantInvoiceCancelReq(unibeeApiMerchantInvoiceCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceCancelPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceClearPaymentPost

> MerchantAuthSsoClearTotpPost200Response InvoiceClearPaymentPost(ctx).UnibeeApiMerchantInvoiceClearPaymentReq(unibeeApiMerchantInvoiceClearPaymentReq).Execute()

Admin Clear Invoice Current Pending Payment

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
	unibeeApiMerchantInvoiceClearPaymentReq := *openapiclient.NewUnibeeApiMerchantInvoiceClearPaymentReq("InvoiceId_example") // UnibeeApiMerchantInvoiceClearPaymentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceClearPaymentPost(context.Background()).UnibeeApiMerchantInvoiceClearPaymentReq(unibeeApiMerchantInvoiceClearPaymentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceClearPaymentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceClearPaymentPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceClearPaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceClearPaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceClearPaymentReq** | [**UnibeeApiMerchantInvoiceClearPaymentReq**](UnibeeApiMerchantInvoiceClearPaymentReq.md) |  | 

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


## InvoiceCreateMarkRefundPost

> MerchantInvoiceCreateMarkRefundPost200Response InvoiceCreateMarkRefundPost(ctx).UnibeeApiMerchantInvoiceCreateMarkRefundReq(unibeeApiMerchantInvoiceCreateMarkRefundReq).Execute()

Create And Mark Manual Refund (Crypto/WireTransfer/External)



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
	unibeeApiMerchantInvoiceCreateMarkRefundReq := *openapiclient.NewUnibeeApiMerchantInvoiceCreateMarkRefundReq("InvoiceId_example", "Reason_example", int64(123)) // UnibeeApiMerchantInvoiceCreateMarkRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceCreateMarkRefundPost(context.Background()).UnibeeApiMerchantInvoiceCreateMarkRefundReq(unibeeApiMerchantInvoiceCreateMarkRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceCreateMarkRefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceCreateMarkRefundPost`: MerchantInvoiceCreateMarkRefundPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceCreateMarkRefundPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceCreateMarkRefundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceCreateMarkRefundReq** | [**UnibeeApiMerchantInvoiceCreateMarkRefundReq**](UnibeeApiMerchantInvoiceCreateMarkRefundReq.md) |  | 

### Return type

[**MerchantInvoiceCreateMarkRefundPost200Response**](MerchantInvoiceCreateMarkRefundPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceCreditNoteListPost

> MerchantInvoiceCreditNoteListPost200Response InvoiceCreditNoteListPost(ctx).UnibeeApiMerchantInvoiceCreditNoteListReq(unibeeApiMerchantInvoiceCreditNoteListReq).Execute()

Bulk CreditNote(Refund Invoice) Invoice List



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
	unibeeApiMerchantInvoiceCreditNoteListReq := *openapiclient.NewUnibeeApiMerchantInvoiceCreditNoteListReq() // UnibeeApiMerchantInvoiceCreditNoteListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceCreditNoteListPost(context.Background()).UnibeeApiMerchantInvoiceCreditNoteListReq(unibeeApiMerchantInvoiceCreditNoteListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceCreditNoteListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceCreditNoteListPost`: MerchantInvoiceCreditNoteListPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceCreditNoteListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceCreditNoteListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceCreditNoteListReq** | [**UnibeeApiMerchantInvoiceCreditNoteListReq**](UnibeeApiMerchantInvoiceCreditNoteListReq.md) |  | 

### Return type

[**MerchantInvoiceCreditNoteListPost200Response**](MerchantInvoiceCreditNoteListPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceDeletePost

> MerchantAuthSsoClearTotpPost200Response InvoiceDeletePost(ctx).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()

Delete Pending Invoice



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
	unibeeApiMerchantInvoiceDeleteReq := *openapiclient.NewUnibeeApiMerchantInvoiceDeleteReq("InvoiceId_example") // UnibeeApiMerchantInvoiceDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceDeletePost(context.Background()).UnibeeApiMerchantInvoiceDeleteReq(unibeeApiMerchantInvoiceDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceDeletePost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

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

Invoice Detail



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

Invoice Detail



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

> MerchantInvoiceEditPost200Response InvoiceEditPost(ctx).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()

Invoice Edit



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
	unibeeApiMerchantInvoiceEditReq := *openapiclient.NewUnibeeApiMerchantInvoiceEditReq("InvoiceId_example") // UnibeeApiMerchantInvoiceEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceEditPost(context.Background()).UnibeeApiMerchantInvoiceEditReq(unibeeApiMerchantInvoiceEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceEditPost`: MerchantInvoiceEditPost200Response
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

[**MerchantInvoiceEditPost200Response**](MerchantInvoiceEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceExternalGatewayInvoiceMarkPaidPost

> MerchantInvoiceExternalGatewayInvoiceMarkPaidPost200Response InvoiceExternalGatewayInvoiceMarkPaidPost(ctx).UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq(unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq).Execute()

Mark External Gateway Invoice As Paid



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
	unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq := *openapiclient.NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq("ExternalTransactionId_example", "InvoiceId_example", "Signature_example", int64(123)) // UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceExternalGatewayInvoiceMarkPaidPost(context.Background()).UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq(unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceExternalGatewayInvoiceMarkPaidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceExternalGatewayInvoiceMarkPaidPost`: MerchantInvoiceExternalGatewayInvoiceMarkPaidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceExternalGatewayInvoiceMarkPaidPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceExternalGatewayInvoiceMarkPaidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq** | [**UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq**](UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq.md) |  | 

### Return type

[**MerchantInvoiceExternalGatewayInvoiceMarkPaidPost200Response**](MerchantInvoiceExternalGatewayInvoiceMarkPaidPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceExternalGatewayInvoiceMarkRefundedPost

> MerchantInvoiceExternalGatewayInvoiceMarkRefundedPost200Response InvoiceExternalGatewayInvoiceMarkRefundedPost(ctx).UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq(unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq).Execute()

Mark External Gateway Invoice As Refunded



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
	unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq := *openapiclient.NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq("ExternalTransactionId_example", "InvoiceId_example", "Signature_example", int64(123)) // UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceExternalGatewayInvoiceMarkRefundedPost(context.Background()).UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq(unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceExternalGatewayInvoiceMarkRefundedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceExternalGatewayInvoiceMarkRefundedPost`: MerchantInvoiceExternalGatewayInvoiceMarkRefundedPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceExternalGatewayInvoiceMarkRefundedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceExternalGatewayInvoiceMarkRefundedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq** | [**UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq**](UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq.md) |  | 

### Return type

[**MerchantInvoiceExternalGatewayInvoiceMarkRefundedPost200Response**](MerchantInvoiceExternalGatewayInvoiceMarkRefundedPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost

> MerchantInvoiceExternalGatewayInvoiceUpdatePaymentLinkPost200Response InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost(ctx).UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq(unibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq).Execute()

Update External Gateway Invoice Payment Link



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
	unibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq := *openapiclient.NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq("ExternalTransactionId_example", "InvoiceId_example", "PaymentLink_example", "Signature_example", int64(123)) // UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost(context.Background()).UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq(unibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost`: MerchantInvoiceExternalGatewayInvoiceUpdatePaymentLinkPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceExternalGatewayInvoiceUpdatePaymentLinkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceExternalGatewayInvoiceUpdatePaymentLinkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq** | [**UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq**](UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq.md) |  | 

### Return type

[**MerchantInvoiceExternalGatewayInvoiceUpdatePaymentLinkPost200Response**](MerchantInvoiceExternalGatewayInvoiceUpdatePaymentLinkPost200Response.md)

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

Finish Invoice



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

> MerchantInvoiceListGet200Response InvoiceListGet(ctx).FirstName(firstName).LastName(lastName).SearchKey(searchKey).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).GatewayIds(gatewayIds).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Type_(type_).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).ReportTimeStart(reportTimeStart).ReportTimeEnd(reportTimeEnd).Execute()

Get Invoice List



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
	firstName := "firstName_example" // string | The firstName of invoice (optional)
	lastName := "lastName_example" // string | The lastName of invoice (optional)
	searchKey := "searchKey_example" // string | Search InvoiceName|Email|ProductName (optional)
	currency := "currency_example" // string | The currency of invoice (optional)
	status := []int32{int32(123)} // []int32 | The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled (optional)
	amountStart := int32(56) // int32 | The filter start amount of invoice (optional)
	amountEnd := int32(56) // int32 | The filter end amount of invoice (optional)
	userId := int64(789) // int64 | The filter userid of invoice (optional)
	gatewayIds := []int64{int64(123)} // []int64 | GatewayIds, Search Filter GatewayIds (optional)
	sendEmail := "sendEmail_example" // string | The filter email of invoice (optional)
	sortField := "sortField_example" // string | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort，asc|desc，Default desc (optional)
	deleteInclude := true // bool | Deleted Involved，Need Admin Permission (optional)
	type_ := int32(56) // int32 | invoice Type, 0-payment, 1-refund (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count By Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)
	reportTimeStart := int64(789) // int64 | ReportTimeStart，UTC timestamp，seconds (optional)
	reportTimeEnd := int64(789) // int64 | ReportTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceListGet(context.Background()).FirstName(firstName).LastName(lastName).SearchKey(searchKey).Currency(currency).Status(status).AmountStart(amountStart).AmountEnd(amountEnd).UserId(userId).GatewayIds(gatewayIds).SendEmail(sendEmail).SortField(sortField).SortType(sortType).DeleteInclude(deleteInclude).Type_(type_).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).ReportTimeStart(reportTimeStart).ReportTimeEnd(reportTimeEnd).Execute()
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
 **searchKey** | **string** | Search InvoiceName|Email|ProductName | 
 **currency** | **string** | The currency of invoice | 
 **status** | **[]int32** | The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | 
 **amountStart** | **int32** | The filter start amount of invoice | 
 **amountEnd** | **int32** | The filter end amount of invoice | 
 **userId** | **int64** | The filter userid of invoice | 
 **gatewayIds** | **[]int64** | GatewayIds, Search Filter GatewayIds | 
 **sendEmail** | **string** | The filter email of invoice | 
 **sortField** | **string** | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | 
 **sortType** | **string** | Sort，asc|desc，Default desc | 
 **deleteInclude** | **bool** | Deleted Involved，Need Admin Permission | 
 **type_** | **int32** | invoice Type, 0-payment, 1-refund | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count By Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 
 **reportTimeStart** | **int64** | ReportTimeStart，UTC timestamp，seconds | 
 **reportTimeEnd** | **int64** | ReportTimeEnd，UTC timestamp，seconds | 

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

Get Invoice List



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


## InvoiceMarkRefundSuccessPost

> MerchantAuthSsoClearTotpPost200Response InvoiceMarkRefundSuccessPost(ctx).UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq(unibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq).Execute()

Mark Manual Invoice Refund As Success



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
	unibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq := *openapiclient.NewUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq("InvoiceId_example") // UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceMarkRefundSuccessPost(context.Background()).UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq(unibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceMarkRefundSuccessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceMarkRefundSuccessPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceMarkRefundSuccessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceMarkRefundSuccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq** | [**UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq**](UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq.md) |  | 

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


## InvoiceMarkWireTransferSuccessPost

> MerchantAuthSsoClearTotpPost200Response InvoiceMarkWireTransferSuccessPost(ctx).UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(unibeeApiMerchantInvoiceMarkWireTransferSuccessReq).Execute()

Mark Wire Transfer Invoice As Paid (Manual)



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
	unibeeApiMerchantInvoiceMarkWireTransferSuccessReq := *openapiclient.NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq("InvoiceId_example", "TransferNumber_example") // UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceMarkWireTransferSuccessPost(context.Background()).UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(unibeeApiMerchantInvoiceMarkWireTransferSuccessReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceMarkWireTransferSuccessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceMarkWireTransferSuccessPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceMarkWireTransferSuccessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceMarkWireTransferSuccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceMarkWireTransferSuccessReq** | [**UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq**](UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq.md) |  | 

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


## InvoiceNewPost

> MerchantInvoiceEditPost200Response InvoiceNewPost(ctx).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()

New Invoice

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
	unibeeApiMerchantInvoiceNewReq := *openapiclient.NewUnibeeApiMerchantInvoiceNewReq("Currency_example", int64(123), int64(123)) // UnibeeApiMerchantInvoiceNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceNewPost(context.Background()).UnibeeApiMerchantInvoiceNewReq(unibeeApiMerchantInvoiceNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceNewPost`: MerchantInvoiceEditPost200Response
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

[**MerchantInvoiceEditPost200Response**](MerchantInvoiceEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicePdfGeneratePost

> MerchantAuthSsoClearTotpPost200Response InvoicePdfGeneratePost(ctx).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()

Generate Invoice PDF

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
	unibeeApiMerchantInvoicePdfGenerateReq := *openapiclient.NewUnibeeApiMerchantInvoicePdfGenerateReq("InvoiceId_example") // UnibeeApiMerchantInvoicePdfGenerateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoicePdfGeneratePost(context.Background()).UnibeeApiMerchantInvoicePdfGenerateReq(unibeeApiMerchantInvoicePdfGenerateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoicePdfGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicePdfGeneratePost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicePdfUpdatePost

> MerchantAuthSsoClearTotpPost200Response InvoicePdfUpdatePost(ctx).UnibeeApiMerchantInvoicePdfUpdateReq(unibeeApiMerchantInvoicePdfUpdateReq).Execute()

Update Invoice PDF

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
	unibeeApiMerchantInvoicePdfUpdateReq := *openapiclient.NewUnibeeApiMerchantInvoicePdfUpdateReq("InvoiceId_example") // UnibeeApiMerchantInvoicePdfUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoicePdfUpdatePost(context.Background()).UnibeeApiMerchantInvoicePdfUpdateReq(unibeeApiMerchantInvoicePdfUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoicePdfUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicePdfUpdatePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoicePdfUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicePdfUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoicePdfUpdateReq** | [**UnibeeApiMerchantInvoicePdfUpdateReq**](UnibeeApiMerchantInvoicePdfUpdateReq.md) |  | 

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


## InvoiceReconvertCryptoAndSendEmailPost

> MerchantAuthSsoClearTotpPost200Response InvoiceReconvertCryptoAndSendEmailPost(ctx).UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(unibeeApiMerchantInvoiceReconvertCryptoAndSendReq).Execute()

Admin Reconvert Crypto Data and Send Invoice Email

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
	unibeeApiMerchantInvoiceReconvertCryptoAndSendReq := *openapiclient.NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq("InvoiceId_example") // UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceReconvertCryptoAndSendEmailPost(context.Background()).UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(unibeeApiMerchantInvoiceReconvertCryptoAndSendReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceReconvertCryptoAndSendEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceReconvertCryptoAndSendEmailPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceRefundPost

> MerchantInvoiceCreateMarkRefundPost200Response InvoiceRefundPost(ctx).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()

Create InvoiceRefund



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
	unibeeApiMerchantInvoiceRefundReq := *openapiclient.NewUnibeeApiMerchantInvoiceRefundReq("InvoiceId_example", "Reason_example", int64(123)) // UnibeeApiMerchantInvoiceRefundReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceRefundPost(context.Background()).UnibeeApiMerchantInvoiceRefundReq(unibeeApiMerchantInvoiceRefundReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceRefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceRefundPost`: MerchantInvoiceCreateMarkRefundPost200Response
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

[**MerchantInvoiceCreateMarkRefundPost200Response**](MerchantInvoiceCreateMarkRefundPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSendEmailPost

> MerchantAuthSsoClearTotpPost200Response InvoiceSendEmailPost(ctx).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()

Send Invoice Email

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
	unibeeApiMerchantInvoiceSendEmailReq := *openapiclient.NewUnibeeApiMerchantInvoiceSendEmailReq("InvoiceId_example") // UnibeeApiMerchantInvoiceSendEmailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceSendEmailPost(context.Background()).UnibeeApiMerchantInvoiceSendEmailReq(unibeeApiMerchantInvoiceSendEmailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceSendEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceSendEmailPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSplitPaymentsGet

> MerchantInvoiceSplitPaymentsGet200Response InvoiceSplitPaymentsGet(ctx).InvoiceId(invoiceId).Execute()

Get Split Payments



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
	invoiceId := "invoiceId_example" // string | The invoice id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceSplitPaymentsGet(context.Background()).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceSplitPaymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceSplitPaymentsGet`: MerchantInvoiceSplitPaymentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceSplitPaymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceSplitPaymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string** | The invoice id | 

### Return type

[**MerchantInvoiceSplitPaymentsGet200Response**](MerchantInvoiceSplitPaymentsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSplitPaymentsPost

> MerchantInvoiceSplitPaymentsGet200Response InvoiceSplitPaymentsPost(ctx).UnibeeApiMerchantInvoiceGetSplitPaymentsReq(unibeeApiMerchantInvoiceGetSplitPaymentsReq).Execute()

Get Split Payments



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
	unibeeApiMerchantInvoiceGetSplitPaymentsReq := *openapiclient.NewUnibeeApiMerchantInvoiceGetSplitPaymentsReq("InvoiceId_example") // UnibeeApiMerchantInvoiceGetSplitPaymentsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceSplitPaymentsPost(context.Background()).UnibeeApiMerchantInvoiceGetSplitPaymentsReq(unibeeApiMerchantInvoiceGetSplitPaymentsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceSplitPaymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceSplitPaymentsPost`: MerchantInvoiceSplitPaymentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceSplitPaymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceSplitPaymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceGetSplitPaymentsReq** | [**UnibeeApiMerchantInvoiceGetSplitPaymentsReq**](UnibeeApiMerchantInvoiceGetSplitPaymentsReq.md) |  | 

### Return type

[**MerchantInvoiceSplitPaymentsGet200Response**](MerchantInvoiceSplitPaymentsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSyncUserSnapshotPost

> MerchantInvoiceEditPost200Response InvoiceSyncUserSnapshotPost(ctx).UnibeeApiMerchantInvoiceSyncUserSnapshotReq(unibeeApiMerchantInvoiceSyncUserSnapshotReq).Execute()

Sync Invoice User Snapshot With Latest User Info



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
	unibeeApiMerchantInvoiceSyncUserSnapshotReq := *openapiclient.NewUnibeeApiMerchantInvoiceSyncUserSnapshotReq("InvoiceId_example") // UnibeeApiMerchantInvoiceSyncUserSnapshotReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Invoice.InvoiceSyncUserSnapshotPost(context.Background()).UnibeeApiMerchantInvoiceSyncUserSnapshotReq(unibeeApiMerchantInvoiceSyncUserSnapshotReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Invoice.InvoiceSyncUserSnapshotPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceSyncUserSnapshotPost`: MerchantInvoiceEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Invoice.InvoiceSyncUserSnapshotPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceSyncUserSnapshotPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceSyncUserSnapshotReq** | [**UnibeeApiMerchantInvoiceSyncUserSnapshotReq**](UnibeeApiMerchantInvoiceSyncUserSnapshotReq.md) |  | 

### Return type

[**MerchantInvoiceEditPost200Response**](MerchantInvoiceEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

