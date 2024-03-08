# \EmailTemplate

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailTemplateActivatePost**](EmailTemplate.md#EmailTemplateActivatePost) | **Post** /merchant/email/template_activate | Merchant Email Template Activate
[**EmailTemplateDeactivatePost**](EmailTemplate.md#EmailTemplateDeactivatePost) | **Post** /merchant/email/template_deactivate | Merchant Email Template Deactivate
[**EmailTemplateListGet**](EmailTemplate.md#EmailTemplateListGet) | **Get** /merchant/email/template_list | Merchant Email Template List
[**EmailTemplateSetDefaultPost**](EmailTemplate.md#EmailTemplateSetDefaultPost) | **Post** /merchant/email/template_set_default | Merchant Email Template Set Default
[**EmailTemplateUpdatePost**](EmailTemplate.md#EmailTemplateUpdatePost) | **Post** /merchant/email/template_update | Merchant Email Template Update



## EmailTemplateActivatePost

> MerchantAuthSsoLoginOTPPost200Response EmailTemplateActivatePost(ctx).UnibeeApiMerchantEmailTemplateActivateReq(unibeeApiMerchantEmailTemplateActivateReq).Execute()

Merchant Email Template Activate

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
	unibeeApiMerchantEmailTemplateActivateReq := *openapiclient.NewUnibeeApiMerchantEmailTemplateActivateReq("TemplateName_example") // UnibeeApiMerchantEmailTemplateActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateActivatePost(context.Background()).UnibeeApiMerchantEmailTemplateActivateReq(unibeeApiMerchantEmailTemplateActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateActivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailTemplateActivateReq** | [**UnibeeApiMerchantEmailTemplateActivateReq**](UnibeeApiMerchantEmailTemplateActivateReq.md) |  | 

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


## EmailTemplateDeactivatePost

> MerchantAuthSsoLoginOTPPost200Response EmailTemplateDeactivatePost(ctx).UnibeeApiMerchantEmailTemplateDeactivateReq(unibeeApiMerchantEmailTemplateDeactivateReq).Execute()

Merchant Email Template Deactivate

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
	unibeeApiMerchantEmailTemplateDeactivateReq := *openapiclient.NewUnibeeApiMerchantEmailTemplateDeactivateReq("TemplateName_example") // UnibeeApiMerchantEmailTemplateDeactivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateDeactivatePost(context.Background()).UnibeeApiMerchantEmailTemplateDeactivateReq(unibeeApiMerchantEmailTemplateDeactivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateDeactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateDeactivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateDeactivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateDeactivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailTemplateDeactivateReq** | [**UnibeeApiMerchantEmailTemplateDeactivateReq**](UnibeeApiMerchantEmailTemplateDeactivateReq.md) |  | 

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


## EmailTemplateListGet

> MerchantEmailTemplateListGet200Response EmailTemplateListGet(ctx).Execute()

Merchant Email Template List

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
	resp, r, err := apiClient.EmailTemplate.EmailTemplateListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateListGet`: MerchantEmailTemplateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateListGetRequest struct via the builder pattern


### Return type

[**MerchantEmailTemplateListGet200Response**](MerchantEmailTemplateListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailTemplateSetDefaultPost

> MerchantAuthSsoLoginOTPPost200Response EmailTemplateSetDefaultPost(ctx).UnibeeApiMerchantEmailTemplateSetDefaultReq(unibeeApiMerchantEmailTemplateSetDefaultReq).Execute()

Merchant Email Template Set Default

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
	unibeeApiMerchantEmailTemplateSetDefaultReq := *openapiclient.NewUnibeeApiMerchantEmailTemplateSetDefaultReq("TemplateName_example") // UnibeeApiMerchantEmailTemplateSetDefaultReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateSetDefaultPost(context.Background()).UnibeeApiMerchantEmailTemplateSetDefaultReq(unibeeApiMerchantEmailTemplateSetDefaultReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateSetDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateSetDefaultPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateSetDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateSetDefaultPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailTemplateSetDefaultReq** | [**UnibeeApiMerchantEmailTemplateSetDefaultReq**](UnibeeApiMerchantEmailTemplateSetDefaultReq.md) |  | 

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


## EmailTemplateUpdatePost

> MerchantAuthSsoLoginOTPPost200Response EmailTemplateUpdatePost(ctx).UnibeeApiMerchantEmailTemplateUpdateReq(unibeeApiMerchantEmailTemplateUpdateReq).Execute()

Merchant Email Template Update

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
	unibeeApiMerchantEmailTemplateUpdateReq := *openapiclient.NewUnibeeApiMerchantEmailTemplateUpdateReq("TemplateContent_example", "TemplateName_example", "TemplateTitle_example") // UnibeeApiMerchantEmailTemplateUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateUpdatePost(context.Background()).UnibeeApiMerchantEmailTemplateUpdateReq(unibeeApiMerchantEmailTemplateUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateUpdatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailTemplateUpdateReq** | [**UnibeeApiMerchantEmailTemplateUpdateReq**](UnibeeApiMerchantEmailTemplateUpdateReq.md) |  | 

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

