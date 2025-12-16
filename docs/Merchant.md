# \Merchant

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmountMultiCurrenciesExchangePost**](Merchant.md#AmountMultiCurrenciesExchangePost) | **Post** /merchant/amount_multi_currencies_exchange | Amount Multi Currencies Exchange
[**CountryConfigListPost**](Merchant.md#CountryConfigListPost) | **Post** /merchant/country_config_list | Edit Country Config
[**EditCountryConfigPost**](Merchant.md#EditCountryConfigPost) | **Post** /merchant/edit_country_config | Get Country Config List
[**EditTotpConfigPost**](Merchant.md#EditTotpConfigPost) | **Post** /merchant/edit_totp_config | Admin Edit 2FA Config
[**GetGet**](Merchant.md#GetGet) | **Get** /merchant/get | Get Profile
[**GetLicenseGet**](Merchant.md#GetLicenseGet) | **Get** /merchant/get_license | Get License
[**GetLicenseUpdateUrlGet**](Merchant.md#GetLicenseUpdateUrlGet) | **Get** /merchant/get_license_update_url | Get License Update Url
[**GetLicenseUpdateUrlPost**](Merchant.md#GetLicenseUpdateUrlPost) | **Post** /merchant/get_license_update_url | Get License Update Url
[**NewApikeyPost**](Merchant.md#NewApikeyPost) | **Post** /merchant/new_apikey | Generate New APIKey
[**SetupMultiCurrenciesPost**](Merchant.md#SetupMultiCurrenciesPost) | **Post** /merchant/setup_multi_currencies | Multi Currencies Setup
[**UpdateCnameDomainPost**](Merchant.md#UpdateCnameDomainPost) | **Post** /merchant/update_cname_domain | Update Merchant CNAME Domain
[**UpdatePortalHostDomainPost**](Merchant.md#UpdatePortalHostDomainPost) | **Post** /merchant/update_portal_host_domain | Update Profile Portal Host Domain
[**UpdatePost**](Merchant.md#UpdatePost) | **Post** /merchant/update | Update Profile



## AmountMultiCurrenciesExchangePost

> MerchantAmountMultiCurrenciesExchangePost200Response AmountMultiCurrenciesExchangePost(ctx).UnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq(unibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq).Execute()

Amount Multi Currencies Exchange

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
	unibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq := *openapiclient.NewUnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq(int64(123), "Currency_example") // UnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.AmountMultiCurrenciesExchangePost(context.Background()).UnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq(unibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.AmountMultiCurrenciesExchangePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AmountMultiCurrenciesExchangePost`: MerchantAmountMultiCurrenciesExchangePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.AmountMultiCurrenciesExchangePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAmountMultiCurrenciesExchangePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq** | [**UnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq**](UnibeeApiMerchantProfileAmountMultiCurrenciesExchangeReq.md) |  | 

### Return type

[**MerchantAmountMultiCurrenciesExchangePost200Response**](MerchantAmountMultiCurrenciesExchangePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountryConfigListPost

> MerchantCountryConfigListPost200Response CountryConfigListPost(ctx).Body(body).Execute()

Edit Country Config

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
	resp, r, err := apiClient.Merchant.CountryConfigListPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.CountryConfigListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountryConfigListPost`: MerchantCountryConfigListPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.CountryConfigListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountryConfigListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MerchantCountryConfigListPost200Response**](MerchantCountryConfigListPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCountryConfigPost

> MerchantAuthSsoClearTotpPost200Response EditCountryConfigPost(ctx).UnibeeApiMerchantProfileEditCountryConfigReq(unibeeApiMerchantProfileEditCountryConfigReq).Execute()

Get Country Config List

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
	unibeeApiMerchantProfileEditCountryConfigReq := *openapiclient.NewUnibeeApiMerchantProfileEditCountryConfigReq("CountryCode_example") // UnibeeApiMerchantProfileEditCountryConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.EditCountryConfigPost(context.Background()).UnibeeApiMerchantProfileEditCountryConfigReq(unibeeApiMerchantProfileEditCountryConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.EditCountryConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCountryConfigPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.EditCountryConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditCountryConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileEditCountryConfigReq** | [**UnibeeApiMerchantProfileEditCountryConfigReq**](UnibeeApiMerchantProfileEditCountryConfigReq.md) |  | 

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


## EditTotpConfigPost

> MerchantAuthSsoClearTotpPost200Response EditTotpConfigPost(ctx).UnibeeApiMerchantProfileEditTotpConfigReq(unibeeApiMerchantProfileEditTotpConfigReq).Execute()

Admin Edit 2FA Config

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
	unibeeApiMerchantProfileEditTotpConfigReq := *openapiclient.NewUnibeeApiMerchantProfileEditTotpConfigReq() // UnibeeApiMerchantProfileEditTotpConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.EditTotpConfigPost(context.Background()).UnibeeApiMerchantProfileEditTotpConfigReq(unibeeApiMerchantProfileEditTotpConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.EditTotpConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTotpConfigPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.EditTotpConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditTotpConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileEditTotpConfigReq** | [**UnibeeApiMerchantProfileEditTotpConfigReq**](UnibeeApiMerchantProfileEditTotpConfigReq.md) |  | 

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


## GetGet

> MerchantGetGet200Response GetGet(ctx).Execute()

Get Profile

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
	resp, r, err := apiClient.Merchant.GetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.GetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGet`: MerchantGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.GetGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGetRequest struct via the builder pattern


### Return type

[**MerchantGetGet200Response**](MerchantGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseGet

> MerchantGetLicenseGet200Response GetLicenseGet(ctx).Execute()

Get License

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
	resp, r, err := apiClient.Merchant.GetLicenseGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.GetLicenseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseGet`: MerchantGetLicenseGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.GetLicenseGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseGetRequest struct via the builder pattern


### Return type

[**MerchantGetLicenseGet200Response**](MerchantGetLicenseGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseUpdateUrlGet

> MerchantGetLicenseUpdateUrlGet200Response GetLicenseUpdateUrlGet(ctx).PlanId(planId).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()

Get License Update Url

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
	planId := int64(789) // int64 | Id of plan to update (optional)
	returnUrl := "returnUrl_example" // string | ReturnUrl (optional)
	cancelUrl := "cancelUrl_example" // string | CancelUrl (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.GetLicenseUpdateUrlGet(context.Background()).PlanId(planId).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.GetLicenseUpdateUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseUpdateUrlGet`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.GetLicenseUpdateUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseUpdateUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planId** | **int64** | Id of plan to update | 
 **returnUrl** | **string** | ReturnUrl | 
 **cancelUrl** | **string** | CancelUrl | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseUpdateUrlPost

> MerchantGetLicenseUpdateUrlGet200Response GetLicenseUpdateUrlPost(ctx).UnibeeApiMerchantProfileGetLicenseUpdateUrlReq(unibeeApiMerchantProfileGetLicenseUpdateUrlReq).Execute()

Get License Update Url

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
	unibeeApiMerchantProfileGetLicenseUpdateUrlReq := *openapiclient.NewUnibeeApiMerchantProfileGetLicenseUpdateUrlReq() // UnibeeApiMerchantProfileGetLicenseUpdateUrlReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.GetLicenseUpdateUrlPost(context.Background()).UnibeeApiMerchantProfileGetLicenseUpdateUrlReq(unibeeApiMerchantProfileGetLicenseUpdateUrlReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.GetLicenseUpdateUrlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseUpdateUrlPost`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.GetLicenseUpdateUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseUpdateUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileGetLicenseUpdateUrlReq** | [**UnibeeApiMerchantProfileGetLicenseUpdateUrlReq**](UnibeeApiMerchantProfileGetLicenseUpdateUrlReq.md) |  | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewApikeyPost

> MerchantNewApikeyPost200Response NewApikeyPost(ctx).Body(body).Execute()

Generate New APIKey



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
	resp, r, err := apiClient.Merchant.NewApikeyPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.NewApikeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApikeyPost`: MerchantNewApikeyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.NewApikeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewApikeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**MerchantNewApikeyPost200Response**](MerchantNewApikeyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupMultiCurrenciesPost

> MerchantSetupMultiCurrenciesPost200Response SetupMultiCurrenciesPost(ctx).UnibeeApiMerchantProfileSetupMultiCurrenciesReq(unibeeApiMerchantProfileSetupMultiCurrenciesReq).Execute()

Multi Currencies Setup

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
	unibeeApiMerchantProfileSetupMultiCurrenciesReq := *openapiclient.NewUnibeeApiMerchantProfileSetupMultiCurrenciesReq() // UnibeeApiMerchantProfileSetupMultiCurrenciesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.SetupMultiCurrenciesPost(context.Background()).UnibeeApiMerchantProfileSetupMultiCurrenciesReq(unibeeApiMerchantProfileSetupMultiCurrenciesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.SetupMultiCurrenciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupMultiCurrenciesPost`: MerchantSetupMultiCurrenciesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.SetupMultiCurrenciesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetupMultiCurrenciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileSetupMultiCurrenciesReq** | [**UnibeeApiMerchantProfileSetupMultiCurrenciesReq**](UnibeeApiMerchantProfileSetupMultiCurrenciesReq.md) |  | 

### Return type

[**MerchantSetupMultiCurrenciesPost200Response**](MerchantSetupMultiCurrenciesPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCnameDomainPost

> MerchantAuthSsoClearTotpPost200Response UpdateCnameDomainPost(ctx).UnibeeApiMerchantProfileUpdateCnameDomainReq(unibeeApiMerchantProfileUpdateCnameDomainReq).Execute()

Update Merchant CNAME Domain

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
	unibeeApiMerchantProfileUpdateCnameDomainReq := *openapiclient.NewUnibeeApiMerchantProfileUpdateCnameDomainReq("Domain_example") // UnibeeApiMerchantProfileUpdateCnameDomainReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.UpdateCnameDomainPost(context.Background()).UnibeeApiMerchantProfileUpdateCnameDomainReq(unibeeApiMerchantProfileUpdateCnameDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.UpdateCnameDomainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCnameDomainPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.UpdateCnameDomainPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCnameDomainPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileUpdateCnameDomainReq** | [**UnibeeApiMerchantProfileUpdateCnameDomainReq**](UnibeeApiMerchantProfileUpdateCnameDomainReq.md) |  | 

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


## UpdatePortalHostDomainPost

> MerchantAuthSsoClearTotpPost200Response UpdatePortalHostDomainPost(ctx).UnibeeApiMerchantProfileUpdatePortalHostReq(unibeeApiMerchantProfileUpdatePortalHostReq).Execute()

Update Profile Portal Host Domain

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
	unibeeApiMerchantProfileUpdatePortalHostReq := *openapiclient.NewUnibeeApiMerchantProfileUpdatePortalHostReq("Domain_example") // UnibeeApiMerchantProfileUpdatePortalHostReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.UpdatePortalHostDomainPost(context.Background()).UnibeeApiMerchantProfileUpdatePortalHostReq(unibeeApiMerchantProfileUpdatePortalHostReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.UpdatePortalHostDomainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortalHostDomainPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.UpdatePortalHostDomainPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortalHostDomainPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileUpdatePortalHostReq** | [**UnibeeApiMerchantProfileUpdatePortalHostReq**](UnibeeApiMerchantProfileUpdatePortalHostReq.md) |  | 

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


## UpdatePost

> MerchantUpdatePost200Response UpdatePost(ctx).UnibeeApiMerchantProfileUpdateReq(unibeeApiMerchantProfileUpdateReq).Execute()

Update Profile

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
	unibeeApiMerchantProfileUpdateReq := *openapiclient.NewUnibeeApiMerchantProfileUpdateReq() // UnibeeApiMerchantProfileUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Merchant.UpdatePost(context.Background()).UnibeeApiMerchantProfileUpdateReq(unibeeApiMerchantProfileUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Merchant.UpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePost`: MerchantUpdatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Merchant.UpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProfileUpdateReq** | [**UnibeeApiMerchantProfileUpdateReq**](UnibeeApiMerchantProfileUpdateReq.md) |  | 

### Return type

[**MerchantUpdatePost200Response**](MerchantUpdatePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

