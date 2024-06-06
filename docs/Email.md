# \Email

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailGatewaySetupPost**](Email.md#EmailGatewaySetupPost) | **Post** /merchant/email/gateway_setup | EmailGatewaySetup



## EmailGatewaySetupPost

> MerchantAuthSsoLoginOTPPost200Response EmailGatewaySetupPost(ctx).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()

EmailGatewaySetup

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
	unibeeApiMerchantEmailGatewaySetupReq := *openapiclient.NewUnibeeApiMerchantEmailGatewaySetupReq("Data_example", "GatewayName_example") // UnibeeApiMerchantEmailGatewaySetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailGatewaySetupPost(context.Background()).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailGatewaySetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailGatewaySetupPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailGatewaySetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailGatewaySetupPostRequest struct via the builder pattern


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

