# \InvoiceImport

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceHistoryInvoiceImportDeletePost**](InvoiceImport.md#InvoiceHistoryInvoiceImportDeletePost) | **Post** /merchant/invoice/history_invoice_import/delete | Delete imported history invoice
[**InvoiceHistoryInvoiceImportPost**](InvoiceImport.md#InvoiceHistoryInvoiceImportPost) | **Post** /merchant/invoice/history_invoice_import | History Invoice Import (supports override)



## InvoiceHistoryInvoiceImportDeletePost

> MerchantAuthSsoClearTotpPost200Response InvoiceHistoryInvoiceImportDeletePost(ctx).UnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq(unibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq).Execute()

Delete imported history invoice



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
	unibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq := *openapiclient.NewUnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq("InvoiceId_example") // UnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceImport.InvoiceHistoryInvoiceImportDeletePost(context.Background()).UnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq(unibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceImport.InvoiceHistoryInvoiceImportDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceHistoryInvoiceImportDeletePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoiceImport.InvoiceHistoryInvoiceImportDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceHistoryInvoiceImportDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq** | [**UnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq**](UnibeeApiMerchantInvoiceHistoryInvoiceImportDeleteReq.md) |  | 

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


## InvoiceHistoryInvoiceImportPost

> MerchantInvoiceEditPost200Response InvoiceHistoryInvoiceImportPost(ctx).UnibeeApiMerchantInvoiceHistoryInvoiceImportReq(unibeeApiMerchantInvoiceHistoryInvoiceImportReq).Execute()

History Invoice Import (supports override)



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
	unibeeApiMerchantInvoiceHistoryInvoiceImportReq := *openapiclient.NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReq("Currency_example", "InvoiceId_example", int64(123)) // UnibeeApiMerchantInvoiceHistoryInvoiceImportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceImport.InvoiceHistoryInvoiceImportPost(context.Background()).UnibeeApiMerchantInvoiceHistoryInvoiceImportReq(unibeeApiMerchantInvoiceHistoryInvoiceImportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceImport.InvoiceHistoryInvoiceImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceHistoryInvoiceImportPost`: MerchantInvoiceEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoiceImport.InvoiceHistoryInvoiceImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceHistoryInvoiceImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInvoiceHistoryInvoiceImportReq** | [**UnibeeApiMerchantInvoiceHistoryInvoiceImportReq**](UnibeeApiMerchantInvoiceHistoryInvoiceImportReq.md) |  | 

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

