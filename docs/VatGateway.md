# \VatGateway

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VatCountryListGet**](VatGateway.md#VatCountryListGet) | **Get** /merchant/vat/country_list | Get Vat Country List
[**VatCountryListPost**](VatGateway.md#VatCountryListPost) | **Post** /merchant/vat/country_list | Get Vat Country List
[**VatInitDefaultGatewayPost**](VatGateway.md#VatInitDefaultGatewayPost) | **Post** /merchant/vat/init_default_gateway | Init Default Vat Gateway
[**VatSetupGatewayPost**](VatGateway.md#VatSetupGatewayPost) | **Post** /merchant/vat/setup_gateway | Vat Gateway Setup
[**VatVatNumberValidatePost**](VatGateway.md#VatVatNumberValidatePost) | **Post** /merchant/vat/vat_number_validate | Vat Number Validation



## VatCountryListGet

> MerchantVatCountryListGet200Response VatCountryListGet(ctx).Execute()

Get Vat Country List

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
	resp, r, err := apiClient.VatGateway.VatCountryListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatCountryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatCountryListGet`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatCountryListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVatCountryListGetRequest struct via the builder pattern


### Return type

[**MerchantVatCountryListGet200Response**](MerchantVatCountryListGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatCountryListPost

> MerchantVatCountryListGet200Response VatCountryListPost(ctx).Body(body).Execute()

Get Vat Country List

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
	resp, r, err := apiClient.VatGateway.VatCountryListPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatCountryListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatCountryListPost`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatCountryListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatCountryListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MerchantVatCountryListGet200Response**](MerchantVatCountryListGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatInitDefaultGatewayPost

> MerchantAuthSsoLoginOTPPost200Response VatInitDefaultGatewayPost(ctx).Body(body).Execute()

Init Default Vat Gateway

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
	resp, r, err := apiClient.VatGateway.VatInitDefaultGatewayPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatInitDefaultGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatInitDefaultGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatInitDefaultGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatInitDefaultGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatSetupGatewayPost

> MerchantEmailGatewaySetupPost200Response VatSetupGatewayPost(ctx).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()

Vat Gateway Setup

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
	unibeeApiMerchantVatSetupGatewayReq := *openapiclient.NewUnibeeApiMerchantVatSetupGatewayReq("Data_example", "GatewayName_example") // UnibeeApiMerchantVatSetupGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatSetupGatewayPost(context.Background()).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatSetupGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatSetupGatewayPost`: MerchantEmailGatewaySetupPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatSetupGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatSetupGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatSetupGatewayReq** | [**UnibeeApiMerchantVatSetupGatewayReq**](UnibeeApiMerchantVatSetupGatewayReq.md) |  | 

### Return type

[**MerchantEmailGatewaySetupPost200Response**](MerchantEmailGatewaySetupPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatVatNumberValidatePost

> MerchantVatVatNumberValidatePost200Response VatVatNumberValidatePost(ctx).UnibeeApiMerchantVatNumberValidateReq(unibeeApiMerchantVatNumberValidateReq).Execute()

Vat Number Validation

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
	unibeeApiMerchantVatNumberValidateReq := *openapiclient.NewUnibeeApiMerchantVatNumberValidateReq("VatNumber_example") // UnibeeApiMerchantVatNumberValidateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatVatNumberValidatePost(context.Background()).UnibeeApiMerchantVatNumberValidateReq(unibeeApiMerchantVatNumberValidateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatVatNumberValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatVatNumberValidatePost`: MerchantVatVatNumberValidatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatVatNumberValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatVatNumberValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatNumberValidateReq** | [**UnibeeApiMerchantVatNumberValidateReq**](UnibeeApiMerchantVatNumberValidateReq.md) |  | 

### Return type

[**MerchantVatVatNumberValidatePost200Response**](MerchantVatVatNumberValidatePost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

