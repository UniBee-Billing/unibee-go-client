# \Vat

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VatCountryListGet**](Vat.md#VatCountryListGet) | **Get** /merchant/vat/country_list | Vat Country List
[**VatCountryListPost**](Vat.md#VatCountryListPost) | **Post** /merchant/vat/country_list | Vat Country List
[**VatSetupGatewayPost**](Vat.md#VatSetupGatewayPost) | **Post** /merchant/vat/setup_gateway | Vat Gateway Setup



## VatCountryListGet

> MerchantVatCountryListGet200Response VatCountryListGet(ctx).MerchantId(merchantId).Execute()

Vat Country List

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
	merchantId := int64(789) // int64 | MerchantId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Vat.VatCountryListGet(context.Background()).MerchantId(merchantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Vat.VatCountryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatCountryListGet`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Vat.VatCountryListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatCountryListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantId** | **int64** | MerchantId | 

### Return type

[**MerchantVatCountryListGet200Response**](MerchantVatCountryListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatCountryListPost

> MerchantVatCountryListGet200Response VatCountryListPost(ctx).UnibeeApiMerchantVatCountryListReq(unibeeApiMerchantVatCountryListReq).Execute()

Vat Country List

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
	unibeeApiMerchantVatCountryListReq := *openapiclient.NewUnibeeApiMerchantVatCountryListReq(int64(123)) // UnibeeApiMerchantVatCountryListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Vat.VatCountryListPost(context.Background()).UnibeeApiMerchantVatCountryListReq(unibeeApiMerchantVatCountryListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Vat.VatCountryListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatCountryListPost`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Vat.VatCountryListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatCountryListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatCountryListReq** | [**UnibeeApiMerchantVatCountryListReq**](UnibeeApiMerchantVatCountryListReq.md) |  | 

### Return type

[**MerchantVatCountryListGet200Response**](MerchantVatCountryListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatSetupGatewayPost

> MerchantAuthSsoLoginOTPPost200Response VatSetupGatewayPost(ctx).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()

Vat Gateway Setup

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
	unibeeApiMerchantVatSetupGatewayReq := *openapiclient.NewUnibeeApiMerchantVatSetupGatewayReq("Data_example", "GatewayName_example") // UnibeeApiMerchantVatSetupGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Vat.VatSetupGatewayPost(context.Background()).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Vat.VatSetupGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatSetupGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Vat.VatSetupGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatSetupGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatSetupGatewayReq** | [**UnibeeApiMerchantVatSetupGatewayReq**](UnibeeApiMerchantVatSetupGatewayReq.md) |  | 

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

