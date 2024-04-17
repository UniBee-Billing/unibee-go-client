# \File

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OssFilePost**](File.md#OssFilePost) | **Post** /merchant/oss/file | Upload File



## OssFilePost

> MerchantOssFilePost200Response OssFilePost(ctx).File(file).Execute()

Upload File

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
	file := TODO // *os.File | File To Upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.File.OssFilePost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `File.OssFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OssFilePost`: MerchantOssFilePost200Response
	fmt.Fprintf(os.Stdout, "Response from `File.OssFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOssFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [***os.File**](*os.File.md) | File To Upload | 

### Return type

[**MerchantOssFilePost200Response**](MerchantOssFilePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

