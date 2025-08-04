# \Email

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailEmailSenderSetupPost**](Email.md#EmailEmailSenderSetupPost) | **Post** /merchant/email/email_sender_setup | Email Sender Setup
[**EmailGatewaySetupPost**](Email.md#EmailGatewaySetupPost) | **Post** /merchant/email/gateway_setup | Email Gateway Setup
[**EmailHistoryListGet**](Email.md#EmailHistoryListGet) | **Get** /merchant/email/history_list | Get Email History List
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


## EmailGatewaySetupPost

> MerchantEmailGatewaySetupPost200Response EmailGatewaySetupPost(ctx).UnibeeApiMerchantEmailGatewaySetupReq(unibeeApiMerchantEmailGatewaySetupReq).Execute()

Email Gateway Setup

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


## EmailHistoryListGet

> MerchantEmailHistoryListGet200Response EmailHistoryListGet(ctx).SearchKey(searchKey).Email(email).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

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
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Email.EmailHistoryListGet(context.Background()).SearchKey(searchKey).Email(email).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
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
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart | 
 **createTimeEnd** | **int64** | CreateTimeEnd | 

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

