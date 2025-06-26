# \Integrations

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationQuickbooksGetAuthorizationUrlGet**](Integrations.md#IntegrationQuickbooksGetAuthorizationUrlGet) | **Get** /merchant/integration/quickbooks/get_authorization_url | Get Quickbooks Authorization Connection URL



## IntegrationQuickbooksGetAuthorizationUrlGet

> MerchantIntegrationQuickbooksGetAuthorizationUrlGet200Response IntegrationQuickbooksGetAuthorizationUrlGet(ctx).ReturnUrl(returnUrl).Execute()

Get Quickbooks Authorization Connection URL

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
	returnUrl := "returnUrl_example" // string | ReturnUrl (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Integrations.IntegrationQuickbooksGetAuthorizationUrlGet(context.Background()).ReturnUrl(returnUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Integrations.IntegrationQuickbooksGetAuthorizationUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationQuickbooksGetAuthorizationUrlGet`: MerchantIntegrationQuickbooksGetAuthorizationUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Integrations.IntegrationQuickbooksGetAuthorizationUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationQuickbooksGetAuthorizationUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnUrl** | **string** | ReturnUrl | 

### Return type

[**MerchantIntegrationQuickbooksGetAuthorizationUrlGet200Response**](MerchantIntegrationQuickbooksGetAuthorizationUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

