# \MerchantInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantGetGet**](MerchantInfoAPI.md#MerchantGetGet) | **Get** /merchant/get | Get Merchant Info
[**MerchantUpdatePost**](MerchantInfoAPI.md#MerchantUpdatePost) | **Post** /merchant/update | Update Merchant Info



## MerchantGetGet

> MerchantGetGet200Response MerchantGetGet(ctx).Execute()

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
	resp, r, err := apiClient.MerchantInfoAPI.MerchantGetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInfoAPI.MerchantGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantGetGet`: MerchantGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInfoAPI.MerchantGetGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantGetGetRequest struct via the builder pattern


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


## MerchantUpdatePost

> MerchantGetGet200Response MerchantUpdatePost(ctx).UnibeeApiMerchantInfoUpdateReq(unibeeApiMerchantInfoUpdateReq).Execute()

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
	unibeeApiMerchantInfoUpdateReq := *openapiclient.NewUnibeeApiMerchantInfoUpdateReq() // UnibeeApiMerchantInfoUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantInfoAPI.MerchantUpdatePost(context.Background()).UnibeeApiMerchantInfoUpdateReq(unibeeApiMerchantInfoUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantInfoAPI.MerchantUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUpdatePost`: MerchantGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantInfoAPI.MerchantUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantInfoUpdateReq** | [**UnibeeApiMerchantInfoUpdateReq**](UnibeeApiMerchantInfoUpdateReq.md) |  | 

### Return type

[**MerchantGetGet200Response**](MerchantGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

