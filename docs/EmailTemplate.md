# \EmailTemplate

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailCustomLocalizationTemplateSyncPost**](EmailTemplate.md#EmailCustomLocalizationTemplateSyncPost) | **Post** /merchant/email/custom_localization_template_sync | Customize Localization Template Sync
[**EmailTemplateActivateLocalizationVersionPost**](EmailTemplate.md#EmailTemplateActivateLocalizationVersionPost) | **Post** /merchant/email/template_activate_localization_version | Activate Email Template Localization Version
[**EmailTemplateAddLocalizationVersionPost**](EmailTemplate.md#EmailTemplateAddLocalizationVersionPost) | **Post** /merchant/email/template_add_localization_version | Add Email Template Localization Version
[**EmailTemplateDeleteLocalizationVersionPost**](EmailTemplate.md#EmailTemplateDeleteLocalizationVersionPost) | **Post** /merchant/email/template_delete_localization_version | Delete Email Template Localization Version
[**EmailTemplateEditLocalizationVersionPost**](EmailTemplate.md#EmailTemplateEditLocalizationVersionPost) | **Post** /merchant/email/template_edit_localization_version | Edit Email Template Localization Version
[**EmailTemplateListGet**](EmailTemplate.md#EmailTemplateListGet) | **Get** /merchant/email/template_list | Get Email Template List



## EmailCustomLocalizationTemplateSyncPost

> MerchantAuthSsoClearTotpPost200Response EmailCustomLocalizationTemplateSyncPost(ctx).UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq(unibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq).Execute()

Customize Localization Template Sync



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
	unibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq := *openapiclient.NewUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq([]openapiclient.UnibeeApiBeanMerchantLocalizationEmailTemplate{*openapiclient.NewUnibeeApiBeanMerchantLocalizationEmailTemplate()}) // UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailCustomLocalizationTemplateSyncPost(context.Background()).UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq(unibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailCustomLocalizationTemplateSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailCustomLocalizationTemplateSyncPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailCustomLocalizationTemplateSyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailCustomLocalizationTemplateSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq** | [**UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq**](UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailTemplateActivateLocalizationVersionPost

> MerchantAuthSsoClearTotpPost200Response EmailTemplateActivateLocalizationVersionPost(ctx).UnibeeApiMerchantEmailActivateLocalizationVersionReq(unibeeApiMerchantEmailActivateLocalizationVersionReq).Execute()

Activate Email Template Localization Version

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
	unibeeApiMerchantEmailActivateLocalizationVersionReq := *openapiclient.NewUnibeeApiMerchantEmailActivateLocalizationVersionReq() // UnibeeApiMerchantEmailActivateLocalizationVersionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateActivateLocalizationVersionPost(context.Background()).UnibeeApiMerchantEmailActivateLocalizationVersionReq(unibeeApiMerchantEmailActivateLocalizationVersionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateActivateLocalizationVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateActivateLocalizationVersionPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateActivateLocalizationVersionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateActivateLocalizationVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailActivateLocalizationVersionReq** | [**UnibeeApiMerchantEmailActivateLocalizationVersionReq**](UnibeeApiMerchantEmailActivateLocalizationVersionReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailTemplateAddLocalizationVersionPost

> MerchantEmailTemplateAddLocalizationVersionPost200Response EmailTemplateAddLocalizationVersionPost(ctx).UnibeeApiMerchantEmailAddLocalizationVersionReq(unibeeApiMerchantEmailAddLocalizationVersionReq).Execute()

Add Email Template Localization Version

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
	unibeeApiMerchantEmailAddLocalizationVersionReq := *openapiclient.NewUnibeeApiMerchantEmailAddLocalizationVersionReq() // UnibeeApiMerchantEmailAddLocalizationVersionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateAddLocalizationVersionPost(context.Background()).UnibeeApiMerchantEmailAddLocalizationVersionReq(unibeeApiMerchantEmailAddLocalizationVersionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateAddLocalizationVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateAddLocalizationVersionPost`: MerchantEmailTemplateAddLocalizationVersionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateAddLocalizationVersionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateAddLocalizationVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailAddLocalizationVersionReq** | [**UnibeeApiMerchantEmailAddLocalizationVersionReq**](UnibeeApiMerchantEmailAddLocalizationVersionReq.md) |  | 

### Return type

[**MerchantEmailTemplateAddLocalizationVersionPost200Response**](MerchantEmailTemplateAddLocalizationVersionPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailTemplateDeleteLocalizationVersionPost

> MerchantAuthSsoClearTotpPost200Response EmailTemplateDeleteLocalizationVersionPost(ctx).UnibeeApiMerchantEmailDeleteLocalizationVersionReq(unibeeApiMerchantEmailDeleteLocalizationVersionReq).Execute()

Delete Email Template Localization Version

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
	unibeeApiMerchantEmailDeleteLocalizationVersionReq := *openapiclient.NewUnibeeApiMerchantEmailDeleteLocalizationVersionReq() // UnibeeApiMerchantEmailDeleteLocalizationVersionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateDeleteLocalizationVersionPost(context.Background()).UnibeeApiMerchantEmailDeleteLocalizationVersionReq(unibeeApiMerchantEmailDeleteLocalizationVersionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateDeleteLocalizationVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateDeleteLocalizationVersionPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateDeleteLocalizationVersionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateDeleteLocalizationVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailDeleteLocalizationVersionReq** | [**UnibeeApiMerchantEmailDeleteLocalizationVersionReq**](UnibeeApiMerchantEmailDeleteLocalizationVersionReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailTemplateEditLocalizationVersionPost

> MerchantEmailTemplateAddLocalizationVersionPost200Response EmailTemplateEditLocalizationVersionPost(ctx).UnibeeApiMerchantEmailEditLocalizationVersionReq(unibeeApiMerchantEmailEditLocalizationVersionReq).Execute()

Edit Email Template Localization Version

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
	unibeeApiMerchantEmailEditLocalizationVersionReq := *openapiclient.NewUnibeeApiMerchantEmailEditLocalizationVersionReq() // UnibeeApiMerchantEmailEditLocalizationVersionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplate.EmailTemplateEditLocalizationVersionPost(context.Background()).UnibeeApiMerchantEmailEditLocalizationVersionReq(unibeeApiMerchantEmailEditLocalizationVersionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplate.EmailTemplateEditLocalizationVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailTemplateEditLocalizationVersionPost`: MerchantEmailTemplateAddLocalizationVersionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplate.EmailTemplateEditLocalizationVersionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailTemplateEditLocalizationVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailEditLocalizationVersionReq** | [**UnibeeApiMerchantEmailEditLocalizationVersionReq**](UnibeeApiMerchantEmailEditLocalizationVersionReq.md) |  | 

### Return type

[**MerchantEmailTemplateAddLocalizationVersionPost200Response**](MerchantEmailTemplateAddLocalizationVersionPost200Response.md)

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

Get Email Template List

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

