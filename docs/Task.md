# \Task

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskExportColumnListPost**](Task.md#TaskExportColumnListPost) | **Post** /merchant/task/export_column_list | ExportColumnList
[**TaskListGet**](Task.md#TaskListGet) | **Get** /merchant/task/list | GetTaskList
[**TaskListPost**](Task.md#TaskListPost) | **Post** /merchant/task/list | GetTaskList
[**TaskNewExportPost**](Task.md#TaskNewExportPost) | **Post** /merchant/task/new_export | NewExport
[**TaskNewImportPost**](Task.md#TaskNewImportPost) | **Post** /merchant/task/new_import | NewImport



## TaskExportColumnListPost

> MerchantTaskExportColumnListPost200Response TaskExportColumnListPost(ctx).UnibeeApiMerchantTaskExportColumnListReq(unibeeApiMerchantTaskExportColumnListReq).Execute()

ExportColumnList

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
	unibeeApiMerchantTaskExportColumnListReq := *openapiclient.NewUnibeeApiMerchantTaskExportColumnListReq() // UnibeeApiMerchantTaskExportColumnListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Task.TaskExportColumnListPost(context.Background()).UnibeeApiMerchantTaskExportColumnListReq(unibeeApiMerchantTaskExportColumnListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Task.TaskExportColumnListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskExportColumnListPost`: MerchantTaskExportColumnListPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Task.TaskExportColumnListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskExportColumnListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTaskExportColumnListReq** | [**UnibeeApiMerchantTaskExportColumnListReq**](UnibeeApiMerchantTaskExportColumnListReq.md) |  | 

### Return type

[**MerchantTaskExportColumnListPost200Response**](MerchantTaskExportColumnListPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskListGet

> MerchantTaskListGet200Response TaskListGet(ctx).Page(page).Count(count).Execute()

GetTaskList

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
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Task.TaskListGet(context.Background()).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Task.TaskListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskListGet`: MerchantTaskListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Task.TaskListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantTaskListGet200Response**](MerchantTaskListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskListPost

> MerchantTaskListGet200Response TaskListPost(ctx).UnibeeApiMerchantTaskListReq(unibeeApiMerchantTaskListReq).Execute()

GetTaskList

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
	unibeeApiMerchantTaskListReq := *openapiclient.NewUnibeeApiMerchantTaskListReq() // UnibeeApiMerchantTaskListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Task.TaskListPost(context.Background()).UnibeeApiMerchantTaskListReq(unibeeApiMerchantTaskListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Task.TaskListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskListPost`: MerchantTaskListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Task.TaskListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTaskListReq** | [**UnibeeApiMerchantTaskListReq**](UnibeeApiMerchantTaskListReq.md) |  | 

### Return type

[**MerchantTaskListGet200Response**](MerchantTaskListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskNewExportPost

> MerchantAuthSsoLoginOTPPost200Response TaskNewExportPost(ctx).UnibeeApiMerchantTaskNewReq(unibeeApiMerchantTaskNewReq).Execute()

NewExport

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
	unibeeApiMerchantTaskNewReq := *openapiclient.NewUnibeeApiMerchantTaskNewReq() // UnibeeApiMerchantTaskNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Task.TaskNewExportPost(context.Background()).UnibeeApiMerchantTaskNewReq(unibeeApiMerchantTaskNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Task.TaskNewExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskNewExportPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Task.TaskNewExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskNewExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTaskNewReq** | [**UnibeeApiMerchantTaskNewReq**](UnibeeApiMerchantTaskNewReq.md) |  | 

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


## TaskNewImportPost

> MerchantAuthSsoLoginOTPPost200Response TaskNewImportPost(ctx).File(file).Task(task).Execute()

NewImport

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
	file := TODO // *os.File | File To Upload (optional)
	task := "task_example" // string | Task,UserImport|ActiveSubscriptionImport (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Task.TaskNewImportPost(context.Background()).File(file).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Task.TaskNewImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskNewImportPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Task.TaskNewImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskNewImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [***os.File**](*os.File.md) | File To Upload | 
 **task** | **string** | Task,UserImport|ActiveSubscriptionImport | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

