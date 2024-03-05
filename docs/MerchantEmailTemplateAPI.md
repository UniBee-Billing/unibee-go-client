# \MerchantEmailTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantEmailTemplateActivatePost**](MerchantEmailTemplateAPI.md#MerchantEmailTemplateActivatePost) | **Post** /merchant/email/template_activate | Merchant Email Template Activate
[**MerchantEmailTemplateDeactivatePost**](MerchantEmailTemplateAPI.md#MerchantEmailTemplateDeactivatePost) | **Post** /merchant/email/template_deactivate | Merchant Email Template Deactivate
[**MerchantEmailTemplateListGet**](MerchantEmailTemplateAPI.md#MerchantEmailTemplateListGet) | **Get** /merchant/email/template_list | Merchant Email Template List
[**MerchantEmailTemplateSetDefaultPost**](MerchantEmailTemplateAPI.md#MerchantEmailTemplateSetDefaultPost) | **Post** /merchant/email/template_set_default | Merchant Email Template Set Default
[**MerchantEmailTemplateUpdatePost**](MerchantEmailTemplateAPI.md#MerchantEmailTemplateUpdatePost) | **Post** /merchant/email/template_update | Merchant Email Template Update



## MerchantEmailTemplateActivatePost

> MerchantAuthSsoLoginOTPPost200Response MerchantEmailTemplateActivatePost(ctx).UnibeeApiMerchantEmailTemplateActivateReq(unibeeApiMerchantEmailTemplateActivateReq).Execute()

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
	resp, r, err := apiClient.MerchantEmailTemplateAPI.MerchantEmailTemplateActivatePost(context.Background()).UnibeeApiMerchantEmailTemplateActivateReq(unibeeApiMerchantEmailTemplateActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailTemplateAPI.MerchantEmailTemplateActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailTemplateActivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailTemplateAPI.MerchantEmailTemplateActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailTemplateActivatePostRequest struct via the builder pattern


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


## MerchantEmailTemplateDeactivatePost

> MerchantAuthSsoLoginOTPPost200Response MerchantEmailTemplateDeactivatePost(ctx).UnibeeApiMerchantEmailTemplateDeactivateReq(unibeeApiMerchantEmailTemplateDeactivateReq).Execute()

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
	resp, r, err := apiClient.MerchantEmailTemplateAPI.MerchantEmailTemplateDeactivatePost(context.Background()).UnibeeApiMerchantEmailTemplateDeactivateReq(unibeeApiMerchantEmailTemplateDeactivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailTemplateAPI.MerchantEmailTemplateDeactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailTemplateDeactivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailTemplateAPI.MerchantEmailTemplateDeactivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailTemplateDeactivatePostRequest struct via the builder pattern


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


## MerchantEmailTemplateListGet

> MerchantEmailTemplateListGet200Response MerchantEmailTemplateListGet(ctx).Execute()

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
	resp, r, err := apiClient.MerchantEmailTemplateAPI.MerchantEmailTemplateListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailTemplateAPI.MerchantEmailTemplateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailTemplateListGet`: MerchantEmailTemplateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailTemplateAPI.MerchantEmailTemplateListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailTemplateListGetRequest struct via the builder pattern


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


## MerchantEmailTemplateSetDefaultPost

> MerchantAuthSsoLoginOTPPost200Response MerchantEmailTemplateSetDefaultPost(ctx).UnibeeApiMerchantEmailTemplateSetDefaultReq(unibeeApiMerchantEmailTemplateSetDefaultReq).Execute()

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
	resp, r, err := apiClient.MerchantEmailTemplateAPI.MerchantEmailTemplateSetDefaultPost(context.Background()).UnibeeApiMerchantEmailTemplateSetDefaultReq(unibeeApiMerchantEmailTemplateSetDefaultReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailTemplateAPI.MerchantEmailTemplateSetDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailTemplateSetDefaultPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailTemplateAPI.MerchantEmailTemplateSetDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailTemplateSetDefaultPostRequest struct via the builder pattern


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


## MerchantEmailTemplateUpdatePost

> MerchantAuthSsoLoginOTPPost200Response MerchantEmailTemplateUpdatePost(ctx).UnibeeApiMerchantEmailTemplateUpdateReq(unibeeApiMerchantEmailTemplateUpdateReq).Execute()

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
	resp, r, err := apiClient.MerchantEmailTemplateAPI.MerchantEmailTemplateUpdatePost(context.Background()).UnibeeApiMerchantEmailTemplateUpdateReq(unibeeApiMerchantEmailTemplateUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantEmailTemplateAPI.MerchantEmailTemplateUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantEmailTemplateUpdatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantEmailTemplateAPI.MerchantEmailTemplateUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantEmailTemplateUpdatePostRequest struct via the builder pattern


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

