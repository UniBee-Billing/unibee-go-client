# \VatGateway

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VatCountryListGet**](VatGateway.md#VatCountryListGet) | **Get** /merchant/vat/country_list | Get Vat Country List
[**VatCountryListPost**](VatGateway.md#VatCountryListPost) | **Post** /merchant/vat/country_list | Get Vat Country List
[**VatInitDefaultGatewayPost**](VatGateway.md#VatInitDefaultGatewayPost) | **Post** /merchant/vat/init_default_gateway | Init Default Vat Gateway
[**VatSetupGatewayPost**](VatGateway.md#VatSetupGatewayPost) | **Post** /merchant/vat/setup_gateway | Vat Gateway Setup
[**VatSetupGlobalUsVatConfigPost**](VatGateway.md#VatSetupGlobalUsVatConfigPost) | **Post** /merchant/vat/setup_global_us_vat_config | US Vat Config Setup
[**VatSetupUsVatGatewayPost**](VatGateway.md#VatSetupUsVatGatewayPost) | **Post** /merchant/vat/setup_us_vat_gateway | US Vat Gateway Setup
[**VatUsVatCategoryListGet**](VatGateway.md#VatUsVatCategoryListGet) | **Get** /merchant/vat/us_vat_category_list | Get US Vat Category List
[**VatUsVatValidateAddressPost**](VatGateway.md#VatUsVatValidateAddressPost) | **Post** /merchant/vat/us_vat_validate_address | Validate US Vat Address
[**VatVatNumberValidateHistoryActivatePost**](VatGateway.md#VatVatNumberValidateHistoryActivatePost) | **Post** /merchant/vat/vat_number_validate_history_activate | Vat Number Validation History Activate
[**VatVatNumberValidateHistoryDeactivatePost**](VatGateway.md#VatVatNumberValidateHistoryDeactivatePost) | **Post** /merchant/vat/vat_number_validate_history_deactivate | Vat Number Validation History Deactivate
[**VatVatNumberValidateHistoryPost**](VatGateway.md#VatVatNumberValidateHistoryPost) | **Post** /merchant/vat/vat_number_validate_history | Vat Number Validation History
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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatInitDefaultGatewayPost

> MerchantAuthSsoClearTotpPost200Response VatInitDefaultGatewayPost(ctx).Body(body).Execute()

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
	// response from `VatInitDefaultGatewayPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatSetupGlobalUsVatConfigPost

> MerchantAuthSsoClearTotpPost200Response VatSetupGlobalUsVatConfigPost(ctx).UnibeeApiMerchantVatSetupGlobalUSVATConfigReq(unibeeApiMerchantVatSetupGlobalUSVATConfigReq).Execute()

US Vat Config Setup

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
	unibeeApiMerchantVatSetupGlobalUSVATConfigReq := *openapiclient.NewUnibeeApiMerchantVatSetupGlobalUSVATConfigReq() // UnibeeApiMerchantVatSetupGlobalUSVATConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatSetupGlobalUsVatConfigPost(context.Background()).UnibeeApiMerchantVatSetupGlobalUSVATConfigReq(unibeeApiMerchantVatSetupGlobalUSVATConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatSetupGlobalUsVatConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatSetupGlobalUsVatConfigPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatSetupGlobalUsVatConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatSetupGlobalUsVatConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatSetupGlobalUSVATConfigReq** | [**UnibeeApiMerchantVatSetupGlobalUSVATConfigReq**](UnibeeApiMerchantVatSetupGlobalUSVATConfigReq.md) |  | 

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


## VatSetupUsVatGatewayPost

> MerchantAuthSsoClearTotpPost200Response VatSetupUsVatGatewayPost(ctx).UnibeeApiMerchantVatSetupUSVATGatewayReq(unibeeApiMerchantVatSetupUSVATGatewayReq).Execute()

US Vat Gateway Setup

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
	unibeeApiMerchantVatSetupUSVATGatewayReq := *openapiclient.NewUnibeeApiMerchantVatSetupUSVATGatewayReq("GatewayName_example") // UnibeeApiMerchantVatSetupUSVATGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatSetupUsVatGatewayPost(context.Background()).UnibeeApiMerchantVatSetupUSVATGatewayReq(unibeeApiMerchantVatSetupUSVATGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatSetupUsVatGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatSetupUsVatGatewayPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatSetupUsVatGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatSetupUsVatGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatSetupUSVATGatewayReq** | [**UnibeeApiMerchantVatSetupUSVATGatewayReq**](UnibeeApiMerchantVatSetupUSVATGatewayReq.md) |  | 

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


## VatUsVatCategoryListGet

> MerchantVatUsVatCategoryListGet200Response VatUsVatCategoryListGet(ctx).Execute()

Get US Vat Category List

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
	resp, r, err := apiClient.VatGateway.VatUsVatCategoryListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatUsVatCategoryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatUsVatCategoryListGet`: MerchantVatUsVatCategoryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatUsVatCategoryListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVatUsVatCategoryListGetRequest struct via the builder pattern


### Return type

[**MerchantVatUsVatCategoryListGet200Response**](MerchantVatUsVatCategoryListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatUsVatValidateAddressPost

> MerchantVatUsVatValidateAddressPost200Response VatUsVatValidateAddressPost(ctx).UnibeeApiMerchantVatUSVATValidateAddressReq(unibeeApiMerchantVatUSVATValidateAddressReq).Execute()

Validate US Vat Address

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
	unibeeApiMerchantVatUSVATValidateAddressReq := *openapiclient.NewUnibeeApiMerchantVatUSVATValidateAddressReq() // UnibeeApiMerchantVatUSVATValidateAddressReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatUsVatValidateAddressPost(context.Background()).UnibeeApiMerchantVatUSVATValidateAddressReq(unibeeApiMerchantVatUSVATValidateAddressReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatUsVatValidateAddressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatUsVatValidateAddressPost`: MerchantVatUsVatValidateAddressPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatUsVatValidateAddressPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatUsVatValidateAddressPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatUSVATValidateAddressReq** | [**UnibeeApiMerchantVatUSVATValidateAddressReq**](UnibeeApiMerchantVatUSVATValidateAddressReq.md) |  | 

### Return type

[**MerchantVatUsVatValidateAddressPost200Response**](MerchantVatUsVatValidateAddressPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VatVatNumberValidateHistoryActivatePost

> MerchantAuthSsoClearTotpPost200Response VatVatNumberValidateHistoryActivatePost(ctx).UnibeeApiMerchantVatNumberValidateHistoryActivateReq(unibeeApiMerchantVatNumberValidateHistoryActivateReq).Execute()

Vat Number Validation History Activate

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
	unibeeApiMerchantVatNumberValidateHistoryActivateReq := *openapiclient.NewUnibeeApiMerchantVatNumberValidateHistoryActivateReq() // UnibeeApiMerchantVatNumberValidateHistoryActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatVatNumberValidateHistoryActivatePost(context.Background()).UnibeeApiMerchantVatNumberValidateHistoryActivateReq(unibeeApiMerchantVatNumberValidateHistoryActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatVatNumberValidateHistoryActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatVatNumberValidateHistoryActivatePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatVatNumberValidateHistoryActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatVatNumberValidateHistoryActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatNumberValidateHistoryActivateReq** | [**UnibeeApiMerchantVatNumberValidateHistoryActivateReq**](UnibeeApiMerchantVatNumberValidateHistoryActivateReq.md) |  | 

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


## VatVatNumberValidateHistoryDeactivatePost

> MerchantAuthSsoClearTotpPost200Response VatVatNumberValidateHistoryDeactivatePost(ctx).UnibeeApiMerchantVatNumberValidateHistoryDeactivateReq(unibeeApiMerchantVatNumberValidateHistoryDeactivateReq).Execute()

Vat Number Validation History Deactivate

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
	unibeeApiMerchantVatNumberValidateHistoryDeactivateReq := *openapiclient.NewUnibeeApiMerchantVatNumberValidateHistoryDeactivateReq() // UnibeeApiMerchantVatNumberValidateHistoryDeactivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatVatNumberValidateHistoryDeactivatePost(context.Background()).UnibeeApiMerchantVatNumberValidateHistoryDeactivateReq(unibeeApiMerchantVatNumberValidateHistoryDeactivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatVatNumberValidateHistoryDeactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatVatNumberValidateHistoryDeactivatePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatVatNumberValidateHistoryDeactivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatVatNumberValidateHistoryDeactivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatNumberValidateHistoryDeactivateReq** | [**UnibeeApiMerchantVatNumberValidateHistoryDeactivateReq**](UnibeeApiMerchantVatNumberValidateHistoryDeactivateReq.md) |  | 

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


## VatVatNumberValidateHistoryPost

> MerchantVatVatNumberValidateHistoryPost200Response VatVatNumberValidateHistoryPost(ctx).UnibeeApiMerchantVatNumberValidateHistoryReq(unibeeApiMerchantVatNumberValidateHistoryReq).Execute()

Vat Number Validation History

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
	unibeeApiMerchantVatNumberValidateHistoryReq := *openapiclient.NewUnibeeApiMerchantVatNumberValidateHistoryReq() // UnibeeApiMerchantVatNumberValidateHistoryReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VatGateway.VatVatNumberValidateHistoryPost(context.Background()).UnibeeApiMerchantVatNumberValidateHistoryReq(unibeeApiMerchantVatNumberValidateHistoryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VatGateway.VatVatNumberValidateHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VatVatNumberValidateHistoryPost`: MerchantVatVatNumberValidateHistoryPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VatGateway.VatVatNumberValidateHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVatVatNumberValidateHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantVatNumberValidateHistoryReq** | [**UnibeeApiMerchantVatNumberValidateHistoryReq**](UnibeeApiMerchantVatNumberValidateHistoryReq.md) |  | 

### Return type

[**MerchantVatVatNumberValidateHistoryPost200Response**](MerchantVatVatNumberValidateHistoryPost200Response.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

