# \Email

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailEmailSenderSetupPost**](Email.md#EmailEmailSenderSetupPost) | **Post** /merchant/email/email_sender_setup | Email Sender Setup
[**EmailGatewaySetDefaultPost**](Email.md#EmailGatewaySetDefaultPost) | **Post** /merchant/email/gateway_set_default | Email Gateway Set Default
[**EmailGatewaySetupPost**](Email.md#EmailGatewaySetupPost) | **Post** /merchant/email/gateway_setup | Email Gateway Setup(Sendgrid)
[**EmailGatewaySetupV2Post**](Email.md#EmailGatewaySetupV2Post) | **Post** /merchant/email/gateway_setup_v2 | Email Gateway Setup V2(Sendgrid|smtp)
[**EmailHistoryListGet**](Email.md#EmailHistoryListGet) | **Get** /merchant/email/history_list | Get Email History List
[**EmailSendEmailToUserPost**](Email.md#EmailSendEmailToUserPost) | **Post** /merchant/email/send_email_to_user | Send Email To User
[**EmailSendEmailViaApiCredentialPost**](Email.md#EmailSendEmailViaApiCredentialPost) | **Post** /merchant/email/send_email_via_api_credential | Send Email Via API Credential
[**EmailSendTemplateEmailToUserPost**](Email.md#EmailSendTemplateEmailToUserPost) | **Post** /merchant/email/send_template_email_to_user | Send Template Email To User



## EmailEmailSenderSetupPost

> MerchantAuthSsoClearTotpPost200Response EmailEmailSenderSetupPost(ctx).UnibeeApiMerchantEmailSenderSetupReq(unibeeApiMerchantEmailSenderSetupReq).Execute()

Email Sender Setup

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
	unibeeApiMerchantEmailSenderSetupReq := *openapiclient.NewUnibeeApiMerchantEmailSenderSetupReq("Address_example", "Name_example") // UnibeeApiMerchantEmailSenderSetupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailEmailSenderSetupPost(context.Background()).UnibeeApiMerchantEmailSenderSetupReq(unibeeApiMerchantEmailSenderSetupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailEmailSenderSetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailEmailSenderSetupPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailEmailSenderSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailEmailSenderSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailSenderSetupReq** | [**UnibeeApiMerchantEmailSenderSetupReq**](UnibeeApiMerchantEmailSenderSetupReq.md) |  | 

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


## EmailGatewaySetDefaultPost

> MerchantAuthSsoClearTotpPost200Response EmailGatewaySetDefaultPost(ctx).UnibeeApiMerchantEmailGatewaySetDefaultReq(unibeeApiMerchantEmailGatewaySetDefaultReq).Execute()

Email Gateway Set Default

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
	unibeeApiMerchantEmailGatewaySetDefaultReq := *openapiclient.NewUnibeeApiMerchantEmailGatewaySetDefaultReq("GatewayName_example") // UnibeeApiMerchantEmailGatewaySetDefaultReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailGatewaySetDefaultPost(context.Background()).UnibeeApiMerchantEmailGatewaySetDefaultReq(unibeeApiMerchantEmailGatewaySetDefaultReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailGatewaySetDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailGatewaySetDefaultPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailGatewaySetDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailGatewaySetDefaultPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailGatewaySetDefaultReq** | [**UnibeeApiMerchantEmailGatewaySetDefaultReq**](UnibeeApiMerchantEmailGatewaySetDefaultReq.md) |  | 

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


## EmailGatewaySetupPost

> MerchantEmailGatewaySetupPost200Response EmailGatewaySetupPost(ctx).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()

Email Gateway Setup(Sendgrid)

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
	// response from `EmailGatewaySetupPost`: MerchantEmailGatewaySetupPost200Response
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

[**MerchantEmailGatewaySetupPost200Response**](MerchantEmailGatewaySetupPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailGatewaySetupV2Post

> MerchantAuthSsoClearTotpPost200Response EmailGatewaySetupV2Post(ctx).UnibeeApiMerchantEmailGatewaySetupV2Req(unibeeApiMerchantEmailGatewaySetupV2Req).Execute()

Email Gateway Setup V2(Sendgrid|smtp)

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
	unibeeApiMerchantEmailGatewaySetupV2Req := *openapiclient.NewUnibeeApiMerchantEmailGatewaySetupV2Req("GatewayName_example") // UnibeeApiMerchantEmailGatewaySetupV2Req | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailGatewaySetupV2Post(context.Background()).UnibeeApiMerchantEmailGatewaySetupV2Req(unibeeApiMerchantEmailGatewaySetupV2Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailGatewaySetupV2Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailGatewaySetupV2Post`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailGatewaySetupV2Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailGatewaySetupV2PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailGatewaySetupV2Req** | [**UnibeeApiMerchantEmailGatewaySetupV2Req**](UnibeeApiMerchantEmailGatewaySetupV2Req.md) |  | 

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


## EmailHistoryListGet

> MerchantEmailHistoryListGet200Response EmailHistoryListGet(ctx).SearchKey(searchKey).Email(email).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Email History List



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
	searchKey := "searchKey_example" // string | Search Key, email or title (optional)
	email := "email_example" // string | Filter Email (optional)
	status := []int32{int32(123)} // []int32 | status, 0-pending, 1-success, 2-failure (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailHistoryListGet(context.Background()).SearchKey(searchKey).Email(email).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailHistoryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailHistoryListGet`: MerchantEmailHistoryListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailHistoryListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailHistoryListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | Search Key, email or title | 
 **email** | **string** | Filter Email | 
 **status** | **[]int32** | status, 0-pending, 1-success, 2-failure | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantEmailHistoryListGet200Response**](MerchantEmailHistoryListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSendEmailToUserPost

> MerchantAuthSsoClearTotpPost200Response EmailSendEmailToUserPost(ctx).UnibeeApiMerchantEmailSendEmailToUserReq(unibeeApiMerchantEmailSendEmailToUserReq).Execute()

Send Email To User

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
	unibeeApiMerchantEmailSendEmailToUserReq := *openapiclient.NewUnibeeApiMerchantEmailSendEmailToUserReq("Email_example") // UnibeeApiMerchantEmailSendEmailToUserReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailSendEmailToUserPost(context.Background()).UnibeeApiMerchantEmailSendEmailToUserReq(unibeeApiMerchantEmailSendEmailToUserReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailSendEmailToUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSendEmailToUserPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailSendEmailToUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSendEmailToUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailSendEmailToUserReq** | [**UnibeeApiMerchantEmailSendEmailToUserReq**](UnibeeApiMerchantEmailSendEmailToUserReq.md) |  | 

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


## EmailSendEmailViaApiCredentialPost

> MerchantAuthSsoClearTotpPost200Response EmailSendEmailViaApiCredentialPost(ctx).UnibeeApiMerchantEmailSendEmailViaAPICredentialReq(unibeeApiMerchantEmailSendEmailViaAPICredentialReq).Execute()

Send Email Via API Credential

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
	unibeeApiMerchantEmailSendEmailViaAPICredentialReq := *openapiclient.NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReq("Email_example", "GatewayName_example") // UnibeeApiMerchantEmailSendEmailViaAPICredentialReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailSendEmailViaApiCredentialPost(context.Background()).UnibeeApiMerchantEmailSendEmailViaAPICredentialReq(unibeeApiMerchantEmailSendEmailViaAPICredentialReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Email.EmailSendEmailViaApiCredentialPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSendEmailViaApiCredentialPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Email.EmailSendEmailViaApiCredentialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSendEmailViaApiCredentialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantEmailSendEmailViaAPICredentialReq** | [**UnibeeApiMerchantEmailSendEmailViaAPICredentialReq**](UnibeeApiMerchantEmailSendEmailViaAPICredentialReq.md) |  | 

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


## EmailSendTemplateEmailToUserPost

> MerchantAuthSsoClearTotpPost200Response EmailSendTemplateEmailToUserPost(ctx).UnibeeApiMerchantEmailSendTemplateEmailToUserReq(unibeeApiMerchantEmailSendTemplateEmailToUserReq).Execute()

Send Template Email To User

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
	// response from `EmailSendTemplateEmailToUserPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

