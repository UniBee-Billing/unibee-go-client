# \Gateway

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayArchivePost**](Gateway.md#GatewayArchivePost) | **Post** /merchant/gateway/archive | Payment Gateway Archive
[**GatewayDetailGet**](Gateway.md#GatewayDetailGet) | **Get** /merchant/gateway/detail | Payment Gateway
[**GatewayDetailPost**](Gateway.md#GatewayDetailPost) | **Post** /merchant/gateway/detail | Payment Gateway
[**GatewayEditCountryConfigPost**](Gateway.md#GatewayEditCountryConfigPost) | **Post** /merchant/gateway/edit_country_config | Payment Gateway Country Config Edit
[**GatewayEditPost**](Gateway.md#GatewayEditPost) | **Post** /merchant/gateway/edit | Payment Gateway Edit
[**GatewayEditSortPost**](Gateway.md#GatewayEditSortPost) | **Post** /merchant/gateway/edit_sort | Edit Payment Gateway Sort
[**GatewayListGet**](Gateway.md#GatewayListGet) | **Get** /merchant/gateway/list | Get Payment Gateway List
[**GatewayRestorePost**](Gateway.md#GatewayRestorePost) | **Post** /merchant/gateway/restore | Payment Gateway Restore
[**GatewaySetDefaultPost**](Gateway.md#GatewaySetDefaultPost) | **Post** /merchant/gateway/set_default | Payment Gateway Set Default
[**GatewaySetupExchangeRateApiPost**](Gateway.md#GatewaySetupExchangeRateApiPost) | **Post** /merchant/gateway/setup_exchange_rate_api | Exchange Rate Api Setup
[**GatewaySetupListGet**](Gateway.md#GatewaySetupListGet) | **Get** /merchant/gateway/setup_list | Get Payment Gateway Setup List
[**GatewaySetupPost**](Gateway.md#GatewaySetupPost) | **Post** /merchant/gateway/setup | Payment Gateway Setup
[**GatewaySetupWebhookPost**](Gateway.md#GatewaySetupWebhookPost) | **Post** /merchant/gateway/setup_webhook | Payment Gateway Webhook Setup
[**GatewayWireTransferEditPost**](Gateway.md#GatewayWireTransferEditPost) | **Post** /merchant/gateway/wire_transfer_edit | Wire Transfer Edit
[**GatewayWireTransferSetupPost**](Gateway.md#GatewayWireTransferSetupPost) | **Post** /merchant/gateway/wire_transfer_setup | Wire Transfer Setup
[**TrackSetupSegmentPost**](Gateway.md#TrackSetupSegmentPost) | **Post** /merchant/track/setup_segment | Segment Setup



## GatewayArchivePost

> MerchantGatewayArchivePost200Response GatewayArchivePost(ctx).UnibeeApiMerchantGatewayArchiveReq(unibeeApiMerchantGatewayArchiveReq).Execute()

Payment Gateway Archive



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
	unibeeApiMerchantGatewayArchiveReq := *openapiclient.NewUnibeeApiMerchantGatewayArchiveReq(int64(123)) // UnibeeApiMerchantGatewayArchiveReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayArchivePost(context.Background()).UnibeeApiMerchantGatewayArchiveReq(unibeeApiMerchantGatewayArchiveReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayArchivePost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayArchiveReq** | [**UnibeeApiMerchantGatewayArchiveReq**](UnibeeApiMerchantGatewayArchiveReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDetailGet

> MerchantGatewayArchivePost200Response GatewayDetailGet(ctx).GatewayId(gatewayId).GatewayName(gatewayName).Execute()

Payment Gateway



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
	gatewayId := int32(56) // int32 | The id of payment gateway, either gatewayId or gatewayName should provided (optional)
	gatewayName := "gatewayName_example" // string | The specified name of payment gateway, either gatewayId or gatewayName should provided, stripe|paypal|changelly|unitpay|payssion|cryptadium, return default gateway if provided (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayDetailGet(context.Background()).GatewayId(gatewayId).GatewayName(gatewayName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayDetailGet`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int32** | The id of payment gateway, either gatewayId or gatewayName should provided | 
 **gatewayName** | **string** | The specified name of payment gateway, either gatewayId or gatewayName should provided, stripe|paypal|changelly|unitpay|payssion|cryptadium, return default gateway if provided | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDetailPost

> MerchantGatewayArchivePost200Response GatewayDetailPost(ctx).UnibeeApiMerchantGatewayDetailReq(unibeeApiMerchantGatewayDetailReq).Execute()

Payment Gateway



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
	unibeeApiMerchantGatewayDetailReq := *openapiclient.NewUnibeeApiMerchantGatewayDetailReq() // UnibeeApiMerchantGatewayDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayDetailPost(context.Background()).UnibeeApiMerchantGatewayDetailReq(unibeeApiMerchantGatewayDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayDetailPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayDetailReq** | [**UnibeeApiMerchantGatewayDetailReq**](UnibeeApiMerchantGatewayDetailReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayEditCountryConfigPost

> MerchantAuthSsoClearTotpPost200Response GatewayEditCountryConfigPost(ctx).UnibeeApiMerchantGatewayEditCountryConfigReq(unibeeApiMerchantGatewayEditCountryConfigReq).Execute()

Payment Gateway Country Config Edit



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
	unibeeApiMerchantGatewayEditCountryConfigReq := *openapiclient.NewUnibeeApiMerchantGatewayEditCountryConfigReq(int64(123)) // UnibeeApiMerchantGatewayEditCountryConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayEditCountryConfigPost(context.Background()).UnibeeApiMerchantGatewayEditCountryConfigReq(unibeeApiMerchantGatewayEditCountryConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayEditCountryConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayEditCountryConfigPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayEditCountryConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayEditCountryConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayEditCountryConfigReq** | [**UnibeeApiMerchantGatewayEditCountryConfigReq**](UnibeeApiMerchantGatewayEditCountryConfigReq.md) |  | 

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


## GatewayEditPost

> MerchantGatewayArchivePost200Response GatewayEditPost(ctx).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()

Payment Gateway Edit



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
	unibeeApiMerchantGatewayEditReq := *openapiclient.NewUnibeeApiMerchantGatewayEditReq(int64(123)) // UnibeeApiMerchantGatewayEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayEditPost(context.Background()).UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayEditPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayEditReq** | [**UnibeeApiMerchantGatewayEditReq**](UnibeeApiMerchantGatewayEditReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayEditSortPost

> MerchantGatewayEditSortPost200Response GatewayEditSortPost(ctx).UnibeeApiMerchantGatewayEditSortReq(unibeeApiMerchantGatewayEditSortReq).Execute()

Edit Payment Gateway Sort

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
	unibeeApiMerchantGatewayEditSortReq := *openapiclient.NewUnibeeApiMerchantGatewayEditSortReq() // UnibeeApiMerchantGatewayEditSortReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayEditSortPost(context.Background()).UnibeeApiMerchantGatewayEditSortReq(unibeeApiMerchantGatewayEditSortReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayEditSortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayEditSortPost`: MerchantGatewayEditSortPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayEditSortPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayEditSortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayEditSortReq** | [**UnibeeApiMerchantGatewayEditSortReq**](UnibeeApiMerchantGatewayEditSortReq.md) |  | 

### Return type

[**MerchantGatewayEditSortPost200Response**](MerchantGatewayEditSortPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListGet

> MerchantGatewayListGet200Response GatewayListGet(ctx).Archive(archive).Execute()

Get Payment Gateway List

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
	archive := true // bool | Filter archive gateway or not, default all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayListGet(context.Background()).Archive(archive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListGet`: MerchantGatewayListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archive** | **bool** | Filter archive gateway or not, default all | 

### Return type

[**MerchantGatewayListGet200Response**](MerchantGatewayListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayRestorePost

> MerchantGatewayArchivePost200Response GatewayRestorePost(ctx).UnibeeApiMerchantGatewayRestoreReq(unibeeApiMerchantGatewayRestoreReq).Execute()

Payment Gateway Restore



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
	unibeeApiMerchantGatewayRestoreReq := *openapiclient.NewUnibeeApiMerchantGatewayRestoreReq(int64(123)) // UnibeeApiMerchantGatewayRestoreReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayRestorePost(context.Background()).UnibeeApiMerchantGatewayRestoreReq(unibeeApiMerchantGatewayRestoreReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayRestorePost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayRestorePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayRestoreReq** | [**UnibeeApiMerchantGatewayRestoreReq**](UnibeeApiMerchantGatewayRestoreReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaySetDefaultPost

> MerchantGatewayArchivePost200Response GatewaySetDefaultPost(ctx).UnibeeApiMerchantGatewaySetDefaultReq(unibeeApiMerchantGatewaySetDefaultReq).Execute()

Payment Gateway Set Default



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
	unibeeApiMerchantGatewaySetDefaultReq := *openapiclient.NewUnibeeApiMerchantGatewaySetDefaultReq(int64(123)) // UnibeeApiMerchantGatewaySetDefaultReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetDefaultPost(context.Background()).UnibeeApiMerchantGatewaySetDefaultReq(unibeeApiMerchantGatewaySetDefaultReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetDefaultPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetDefaultPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetDefaultReq** | [**UnibeeApiMerchantGatewaySetDefaultReq**](UnibeeApiMerchantGatewaySetDefaultReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaySetupExchangeRateApiPost

> MerchantGatewaySetupExchangeRateApiPost200Response GatewaySetupExchangeRateApiPost(ctx).UnibeeApiMerchantGatewaySetupExchangeApiReq(unibeeApiMerchantGatewaySetupExchangeApiReq).Execute()

Exchange Rate Api Setup

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
	unibeeApiMerchantGatewaySetupExchangeApiReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupExchangeApiReq() // UnibeeApiMerchantGatewaySetupExchangeApiReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetupExchangeRateApiPost(context.Background()).UnibeeApiMerchantGatewaySetupExchangeApiReq(unibeeApiMerchantGatewaySetupExchangeApiReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupExchangeRateApiPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupExchangeRateApiPost`: MerchantGatewaySetupExchangeRateApiPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupExchangeRateApiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupExchangeRateApiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupExchangeApiReq** | [**UnibeeApiMerchantGatewaySetupExchangeApiReq**](UnibeeApiMerchantGatewaySetupExchangeApiReq.md) |  | 

### Return type

[**MerchantGatewaySetupExchangeRateApiPost200Response**](MerchantGatewaySetupExchangeRateApiPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaySetupListGet

> MerchantGatewayEditSortPost200Response GatewaySetupListGet(ctx).Execute()

Get Payment Gateway Setup List

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
	resp, r, err := apiClient.Gateway.GatewaySetupListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupListGet`: MerchantGatewayEditSortPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupListGetRequest struct via the builder pattern


### Return type

[**MerchantGatewayEditSortPost200Response**](MerchantGatewayEditSortPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaySetupPost

> MerchantGatewayArchivePost200Response GatewaySetupPost(ctx).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()

Payment Gateway Setup



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
	unibeeApiMerchantGatewaySetupReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupReq("GatewayName_example") // UnibeeApiMerchantGatewaySetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetupPost(context.Background()).UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupReq** | [**UnibeeApiMerchantGatewaySetupReq**](UnibeeApiMerchantGatewaySetupReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaySetupWebhookPost

> MerchantGatewaySetupWebhookPost200Response GatewaySetupWebhookPost(ctx).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()

Payment Gateway Webhook Setup

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
	unibeeApiMerchantGatewaySetupWebhookReq := *openapiclient.NewUnibeeApiMerchantGatewaySetupWebhookReq(int64(123)) // UnibeeApiMerchantGatewaySetupWebhookReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewaySetupWebhookPost(context.Background()).UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewaySetupWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewaySetupWebhookPost`: MerchantGatewaySetupWebhookPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewaySetupWebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewaySetupWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewaySetupWebhookReq** | [**UnibeeApiMerchantGatewaySetupWebhookReq**](UnibeeApiMerchantGatewaySetupWebhookReq.md) |  | 

### Return type

[**MerchantGatewaySetupWebhookPost200Response**](MerchantGatewaySetupWebhookPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayWireTransferEditPost

> MerchantGatewayArchivePost200Response GatewayWireTransferEditPost(ctx).UnibeeApiMerchantGatewayWireTransferEditReq(unibeeApiMerchantGatewayWireTransferEditReq).Execute()

Wire Transfer Edit



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
	unibeeApiMerchantGatewayWireTransferEditReq := *openapiclient.NewUnibeeApiMerchantGatewayWireTransferEditReq("Currency_example", int64(123), int64(123)) // UnibeeApiMerchantGatewayWireTransferEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayWireTransferEditPost(context.Background()).UnibeeApiMerchantGatewayWireTransferEditReq(unibeeApiMerchantGatewayWireTransferEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayWireTransferEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayWireTransferEditPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayWireTransferEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayWireTransferEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayWireTransferEditReq** | [**UnibeeApiMerchantGatewayWireTransferEditReq**](UnibeeApiMerchantGatewayWireTransferEditReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayWireTransferSetupPost

> MerchantGatewayArchivePost200Response GatewayWireTransferSetupPost(ctx).UnibeeApiMerchantGatewayWireTransferSetupReq(unibeeApiMerchantGatewayWireTransferSetupReq).Execute()

Wire Transfer Setup



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
	unibeeApiMerchantGatewayWireTransferSetupReq := *openapiclient.NewUnibeeApiMerchantGatewayWireTransferSetupReq("Currency_example", int64(123)) // UnibeeApiMerchantGatewayWireTransferSetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.GatewayWireTransferSetupPost(context.Background()).UnibeeApiMerchantGatewayWireTransferSetupReq(unibeeApiMerchantGatewayWireTransferSetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.GatewayWireTransferSetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayWireTransferSetupPost`: MerchantGatewayArchivePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.GatewayWireTransferSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayWireTransferSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantGatewayWireTransferSetupReq** | [**UnibeeApiMerchantGatewayWireTransferSetupReq**](UnibeeApiMerchantGatewayWireTransferSetupReq.md) |  | 

### Return type

[**MerchantGatewayArchivePost200Response**](MerchantGatewayArchivePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackSetupSegmentPost

> MerchantAuthSsoClearTotpPost200Response TrackSetupSegmentPost(ctx).UnibeeApiMerchantTrackSetupSegmentReq(unibeeApiMerchantTrackSetupSegmentReq).Execute()

Segment Setup

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
	unibeeApiMerchantTrackSetupSegmentReq := *openapiclient.NewUnibeeApiMerchantTrackSetupSegmentReq("ServerSideSecret_example", "UserPortalSecret_example") // UnibeeApiMerchantTrackSetupSegmentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Gateway.TrackSetupSegmentPost(context.Background()).UnibeeApiMerchantTrackSetupSegmentReq(unibeeApiMerchantTrackSetupSegmentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Gateway.TrackSetupSegmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackSetupSegmentPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Gateway.TrackSetupSegmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackSetupSegmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTrackSetupSegmentReq** | [**UnibeeApiMerchantTrackSetupSegmentReq**](UnibeeApiMerchantTrackSetupSegmentReq.md) |  | 

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

