# \Profile

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountryConfigListPost**](Profile.md#CountryConfigListPost) | **Post** /merchant/country_config_list | Merchant Edit Country Config
[**EditCountryConfigPost**](Profile.md#EditCountryConfigPost) | **Post** /merchant/edit_country_config | Merchant Country Config List
[**GetGet**](Profile.md#GetGet) | **Get** /merchant/get | Get Merchant Info
[**NewApikeyPost**](Profile.md#NewApikeyPost) | **Post** /merchant/new_apikey | Merchant New ApiKey, The old key will expire in one hour
[**UpdatePost**](Profile.md#UpdatePost) | **Post** /merchant/update | Update Merchant Info



## CountryConfigListPost

> MerchantCountryConfigListPost200Response CountryConfigListPost(ctx).Body(body).Execute()

Merchant Edit Country Config

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Profile.CountryConfigListPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Profile.CountryConfigListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountryConfigListPost`: MerchantCountryConfigListPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Profile.CountryConfigListPost`: %v\n", resp)
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

Merchant Country Config List

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
	unibeeApiMerchantProfileEditCountryConfigReq := *openapiclient.NewUnibeeApiMerchantProfileEditCountryConfigReq("CountryCode_example") // UnibeeApiMerchantProfileEditCountryConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Profile.EditCountryConfigPost(context.Background()).UnibeeApiMerchantProfileEditCountryConfigReq(unibeeApiMerchantProfileEditCountryConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Profile.EditCountryConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCountryConfigPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Profile.EditCountryConfigPost`: %v\n", resp)
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

Get Merchant Info

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Profile.GetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Profile.GetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGet`: MerchantGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Profile.GetGet`: %v\n", resp)
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


## NewApikeyPost

> MerchantNewApikeyPost200Response NewApikeyPost(ctx).Body(body).Execute()

Merchant New ApiKey, The old key will expire in one hour

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Profile.NewApikeyPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Profile.NewApikeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApikeyPost`: MerchantNewApikeyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Profile.NewApikeyPost`: %v\n", resp)
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

Update Merchant Info

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
	unibeeApiMerchantProfileUpdateReq := *openapiclient.NewUnibeeApiMerchantProfileUpdateReq() // UnibeeApiMerchantProfileUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Profile.UpdatePost(context.Background()).UnibeeApiMerchantProfileUpdateReq(unibeeApiMerchantProfileUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Profile.UpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePost`: MerchantUpdatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Profile.UpdatePost`: %v\n", resp)
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

