# \Merchant

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountryConfigListPost**](Merchant.md#CountryConfigListPost) | **Post** /merchant/country_config_list | Edit Country Config
[**EditCountryConfigPost**](Merchant.md#EditCountryConfigPost) | **Post** /merchant/edit_country_config | Get Country Config List
[**GetGet**](Merchant.md#GetGet) | **Get** /merchant/get | Get Profile
[**GetLicenseGet**](Merchant.md#GetLicenseGet) | **Get** /merchant/get_license | Get License
[**GetLicenseUpdateUrlGet**](Merchant.md#GetLicenseUpdateUrlGet) | **Get** /merchant/get_license_update_url | Get License Update Url
[**NewApikeyPost**](Merchant.md#NewApikeyPost) | **Post** /merchant/new_apikey | Generate New APIKey
[**UpdatePost**](Merchant.md#UpdatePost) | **Post** /merchant/update | Update Profile



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

> MerchantAuthSsoLoginOTPPost200Response EditCountryConfigPost(ctx).UnibeeApiMerchantProfileEditCountryConfigReq(unibeeApiMerchantProfileEditCountryConfigReq).Execute()

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
	// response from `EditCountryConfigPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

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

