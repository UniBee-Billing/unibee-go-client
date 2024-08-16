# \Email

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailGatewaySetupPost**](Email.md#EmailGatewaySetupPost) | **Post** /merchant/email/gateway_setup | EmailGatewaySetup
[**EmailSendTemplateEmailToUserPost**](Email.md#EmailSendTemplateEmailToUserPost) | **Post** /merchant/email/send_template_email_to_user | SendTemplateEmailToUser



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
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
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


## EmailSendTemplateEmailToUserPost

> MerchantAuthSsoLoginOTPPost200Response EmailSendTemplateEmailToUserPost(ctx).UnibeeApiMerchantEmailSendTemplateEmailToUserReq(unibeeApiMerchantEmailSendTemplateEmailToUserReq).Execute()

SendTemplateEmailToUser

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
	unibeeApiMerchantEmailSendTemplateEmailToUserReq := *openapiclient.NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq("TemplateName_example", int64(123)) // UnibeeApiMerchantEmailSendTemplateEmailToUserReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailSendTemplateEmailToUserPost(context.Background()).UnibeeApiMerchantEmailSendTemplateEmailToUserReq(unibeeApiMerchantEmailSendTemplateEmailToUserReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailSendTemplateEmailToUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSendTemplateEmailToUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailSendTemplateEmailToUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSendTemplateEmailToUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailSendTemplateEmailToUserReq** | [**UnibeeApiMerchantEmailSendTemplateEmailToUserReq**](UnibeeApiMerchantEmailSendTemplateEmailToUserReq.md) |  | 

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

