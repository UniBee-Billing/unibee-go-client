# \MerchantEmailAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantEmailGatewaySetupPost**](MerchantEmailAPI.md#MerchantEmailGatewaySetupPost) | **Post** /merchant/email/gateway_setup | Merchant Email Gateway Setup



## MerchantEmailGatewaySetupPost

> MerchantAuthSsoLoginOTPPost200Response MerchantEmailGatewaySetupPost(ctx).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()

Merchant Email Gateway Setup

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
	unibeeApiMerchantEmailGatewaySetupReq := *openapiclient.NewUnibeeApiMerchantEmailGatewaySetupReq("Data_example", "GatewayName_example") // UnibeeApiMerchantEmailGatewaySetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantEmailAPI.MerchantEmailGatewaySetupPost(context.Background()).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailAPI.MerchantEmailGatewaySetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailGatewaySetupPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailAPI.MerchantEmailGatewaySetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailGatewaySetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailGatewaySetupReq** | [**UnibeeApiMerchantEmailGatewaySetupReq**](UnibeeApiMerchantEmailGatewaySetupReq.md) |  | 

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

