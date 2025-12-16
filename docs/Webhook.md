# \Webhook

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VatUsVatUploadHistoryListGet**](Webhook.md#VatUsVatUploadHistoryListGet) | **Get** /merchant/vat/us_vat_upload_history_list | Get US VAT Upload History List
[**WebhookDeleteEndpointPost**](Webhook.md#WebhookDeleteEndpointPost) | **Post** /merchant/webhook/delete_endpoint | Delete Webhook Endpoint
[**WebhookEndpointListGet**](Webhook.md#WebhookEndpointListGet) | **Get** /merchant/webhook/endpoint_list | Get Webhook Endpoint list
[**WebhookEndpointLogListGet**](Webhook.md#WebhookEndpointLogListGet) | **Get** /merchant/webhook/endpoint_log_list | Get Webhook Endpoint Log List
[**WebhookEventListGet**](Webhook.md#WebhookEventListGet) | **Get** /merchant/webhook/event_list | Webhook Event List
[**WebhookGetWebhookSecretGet**](Webhook.md#WebhookGetWebhookSecretGet) | **Get** /merchant/webhook/get_webhook_secret | Get Webhook Secret
[**WebhookNewEndpointPost**](Webhook.md#WebhookNewEndpointPost) | **Post** /merchant/webhook/new_endpoint | New Webhook Endpoint
[**WebhookRemoveWebhookZapierConnectPost**](Webhook.md#WebhookRemoveWebhookZapierConnectPost) | **Post** /merchant/webhook/remove_webhook_zapier_connect | Remove Webhook Zapier Connnect
[**WebhookResendPost**](Webhook.md#WebhookResendPost) | **Post** /merchant/webhook/resend | Resent Webhook
[**WebhookUpdateEndpointPost**](Webhook.md#WebhookUpdateEndpointPost) | **Post** /merchant/webhook/update_endpoint | Update Webhook Endpoint
[**WebhookWebhookZapierConnectUrlPost**](Webhook.md#WebhookWebhookZapierConnectUrlPost) | **Post** /merchant/webhook/webhook_zapier_connect_url | Get Webhook Zapier Connect Url



## VatUsVatUploadHistoryListGet

> MerchantVatUsVatUploadHistoryListGet200Response VatUsVatUploadHistoryListGet(ctx).InvoiceId(invoiceId).GatewayName(gatewayName).UploadType(uploadType).Page(page).Count(count).Execute()

Get US VAT Upload History List

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
	invoiceId := "invoiceId_example" // string | Invoice Id (optional)
	gatewayName := "gatewayName_example" // string | GatewayName, em. TaxJar (optional)
	uploadType := []string{"Inner_example"} // []string | UploadType, PaymentUpload or RefundUpload, default both (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.VatUsVatUploadHistoryListGet(context.Background()).InvoiceId(invoiceId).GatewayName(gatewayName).UploadType(uploadType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.VatUsVatUploadHistoryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatUsVatUploadHistoryListGet`: MerchantVatUsVatUploadHistoryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.VatUsVatUploadHistoryListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatUsVatUploadHistoryListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string** | Invoice Id | 
 **gatewayName** | **string** | GatewayName, em. TaxJar | 
 **uploadType** | **[]string** | UploadType, PaymentUpload or RefundUpload, default both | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantVatUsVatUploadHistoryListGet200Response**](MerchantVatUsVatUploadHistoryListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeleteEndpointPost

> MerchantAuthSsoClearTotpPost200Response WebhookDeleteEndpointPost(ctx).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()

Delete Webhook Endpoint

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
	unibeeApiMerchantWebhookDeleteEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookDeleteEndpointReq(int64(123)) // UnibeeApiMerchantWebhookDeleteEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookDeleteEndpointPost(context.Background()).UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookDeleteEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeleteEndpointPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookDeleteEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeleteEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookDeleteEndpointReq** | [**UnibeeApiMerchantWebhookDeleteEndpointReq**](UnibeeApiMerchantWebhookDeleteEndpointReq.md) |  | 

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


## WebhookEndpointListGet

> MerchantWebhookEndpointListGet200Response WebhookEndpointListGet(ctx).Execute()

Get Webhook Endpoint list

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
	resp, r, err := apiClient.Webhook.WebhookEndpointListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEndpointListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointListGet`: MerchantWebhookEndpointListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEndpointListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointListGetRequest struct via the builder pattern


### Return type

[**MerchantWebhookEndpointListGet200Response**](MerchantWebhookEndpointListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointLogListGet

> MerchantWebhookEndpointLogListGet200Response WebhookEndpointLogListGet(ctx).EndpointId(endpointId).Page(page).Count(count).Execute()

Get Webhook Endpoint Log List

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
	endpointId := int64(789) // int64 | EndpointId
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookEndpointLogListGet(context.Background()).EndpointId(endpointId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEndpointLogListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointLogListGet`: MerchantWebhookEndpointLogListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEndpointLogListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointLogListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointId** | **int64** | EndpointId | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantWebhookEndpointLogListGet200Response**](MerchantWebhookEndpointLogListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEventListGet

> MerchantWebhookEventListGet200Response WebhookEventListGet(ctx).Execute()

Webhook Event List

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
	resp, r, err := apiClient.Webhook.WebhookEventListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookEventListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEventListGet`: MerchantWebhookEventListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookEventListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEventListGetRequest struct via the builder pattern


### Return type

[**MerchantWebhookEventListGet200Response**](MerchantWebhookEventListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookGetWebhookSecretGet

> MerchantWebhookGetWebhookSecretGet200Response WebhookGetWebhookSecretGet(ctx).Execute()

Get Webhook Secret

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
	resp, r, err := apiClient.Webhook.WebhookGetWebhookSecretGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookGetWebhookSecretGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGetWebhookSecretGet`: MerchantWebhookGetWebhookSecretGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookGetWebhookSecretGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetWebhookSecretGetRequest struct via the builder pattern


### Return type

[**MerchantWebhookGetWebhookSecretGet200Response**](MerchantWebhookGetWebhookSecretGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookNewEndpointPost

> MerchantAuthSsoClearTotpPost200Response WebhookNewEndpointPost(ctx).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()

New Webhook Endpoint

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
	unibeeApiMerchantWebhookNewEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookNewEndpointReq("Url_example") // UnibeeApiMerchantWebhookNewEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookNewEndpointPost(context.Background()).UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookNewEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookNewEndpointPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookNewEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookNewEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookNewEndpointReq** | [**UnibeeApiMerchantWebhookNewEndpointReq**](UnibeeApiMerchantWebhookNewEndpointReq.md) |  | 

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


## WebhookRemoveWebhookZapierConnectPost

> MerchantAuthSsoClearTotpPost200Response WebhookRemoveWebhookZapierConnectPost(ctx).Body(body).Execute()

Remove Webhook Zapier Connnect

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookRemoveWebhookZapierConnectPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookRemoveWebhookZapierConnectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookRemoveWebhookZapierConnectPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookRemoveWebhookZapierConnectPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookRemoveWebhookZapierConnectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## WebhookResendPost

> MerchantWebhookResendPost200Response WebhookResendPost(ctx).UnibeeApiMerchantWebhookResendWebhookReq(unibeeApiMerchantWebhookResendWebhookReq).Execute()

Resent Webhook

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
	unibeeApiMerchantWebhookResendWebhookReq := *openapiclient.NewUnibeeApiMerchantWebhookResendWebhookReq(int64(123)) // UnibeeApiMerchantWebhookResendWebhookReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookResendPost(context.Background()).UnibeeApiMerchantWebhookResendWebhookReq(unibeeApiMerchantWebhookResendWebhookReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookResendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookResendPost`: MerchantWebhookResendPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookResendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookResendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookResendWebhookReq** | [**UnibeeApiMerchantWebhookResendWebhookReq**](UnibeeApiMerchantWebhookResendWebhookReq.md) |  | 

### Return type

[**MerchantWebhookResendPost200Response**](MerchantWebhookResendPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookUpdateEndpointPost

> MerchantAuthSsoClearTotpPost200Response WebhookUpdateEndpointPost(ctx).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()

Update Webhook Endpoint

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
	unibeeApiMerchantWebhookUpdateEndpointReq := *openapiclient.NewUnibeeApiMerchantWebhookUpdateEndpointReq(int64(123), "Url_example") // UnibeeApiMerchantWebhookUpdateEndpointReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookUpdateEndpointPost(context.Background()).UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookUpdateEndpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookUpdateEndpointPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookUpdateEndpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookUpdateEndpointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookUpdateEndpointReq** | [**UnibeeApiMerchantWebhookUpdateEndpointReq**](UnibeeApiMerchantWebhookUpdateEndpointReq.md) |  | 

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


## WebhookWebhookZapierConnectUrlPost

> MerchantWebhookWebhookZapierConnectUrlPost200Response WebhookWebhookZapierConnectUrlPost(ctx).UnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq(unibeeApiMerchantWebhookGetWebhookZapierConnectURLReq).Execute()

Get Webhook Zapier Connect Url

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
	unibeeApiMerchantWebhookGetWebhookZapierConnectURLReq := *openapiclient.NewUnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq() // UnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Webhook.WebhookWebhookZapierConnectUrlPost(context.Background()).UnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq(unibeeApiMerchantWebhookGetWebhookZapierConnectURLReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Webhook.WebhookWebhookZapierConnectUrlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookWebhookZapierConnectUrlPost`: MerchantWebhookWebhookZapierConnectUrlPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Webhook.WebhookWebhookZapierConnectUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookWebhookZapierConnectUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantWebhookGetWebhookZapierConnectURLReq** | [**UnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq**](UnibeeApiMerchantWebhookGetWebhookZapierConnectURLReq.md) |  | 

### Return type

[**MerchantWebhookWebhookZapierConnectUrlPost200Response**](MerchantWebhookWebhookZapierConnectUrlPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

