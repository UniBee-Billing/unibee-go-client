# \MerchantVatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantVatCountryListGet**](MerchantVatAPI.md#MerchantVatCountryListGet) | **Get** /merchant/vat/country_list | Vat Country List
[**MerchantVatCountryListPost**](MerchantVatAPI.md#MerchantVatCountryListPost) | **Post** /merchant/vat/country_list | Vat Country List
[**MerchantVatSetupGatewayPost**](MerchantVatAPI.md#MerchantVatSetupGatewayPost) | **Post** /merchant/vat/setup_gateway | Vat Gateway Setup



## MerchantVatCountryListGet

> MerchantVatCountryListGet200Response MerchantVatCountryListGet(ctx).MerchantId(merchantId).Execute()

Vat Country List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	merchantId := int32(56) // int32 | MerchantId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantVatAPI.MerchantVatCountryListGet(context.Background()).MerchantId(merchantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantVatAPI.MerchantVatCountryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantVatCountryListGet`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantVatAPI.MerchantVatCountryListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantVatCountryListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantId** | **int32** | MerchantId | 

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


## MerchantVatCountryListPost

> MerchantVatCountryListGet200Response MerchantVatCountryListPost(ctx).UnibeeApiMerchantVatCountryListReq(unibeeApiMerchantVatCountryListReq).Execute()

Vat Country List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantVatCountryListReq := *openapiclient.NewUnibeeApiMerchantVatCountryListReq(int32(123)) // UnibeeApiMerchantVatCountryListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantVatAPI.MerchantVatCountryListPost(context.Background()).UnibeeApiMerchantVatCountryListReq(unibeeApiMerchantVatCountryListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantVatAPI.MerchantVatCountryListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantVatCountryListPost`: MerchantVatCountryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantVatAPI.MerchantVatCountryListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantVatCountryListPostRequest struct via the builder pattern


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


## MerchantVatSetupGatewayPost

> MerchantAuthSsoLoginOTPPost200Response MerchantVatSetupGatewayPost(ctx).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()

Vat Gateway Setup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantVatSetupGatewayReq := *openapiclient.NewUnibeeApiMerchantVatSetupGatewayReq("Data_example", "GatewayName_example") // UnibeeApiMerchantVatSetupGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantVatAPI.MerchantVatSetupGatewayPost(context.Background()).UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantVatAPI.MerchantVatSetupGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantVatSetupGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantVatAPI.MerchantVatSetupGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantVatSetupGatewayPostRequest struct via the builder pattern


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

