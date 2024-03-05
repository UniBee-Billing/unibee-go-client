# \MerchantBalanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantBalanceMerchantBalanceQueryGet**](MerchantBalanceAPI.md#MerchantBalanceMerchantBalanceQueryGet) | **Get** /merchant/balance/merchant_balance_query | Query Merchant Gateway Balance
[**MerchantBalanceMerchantBalanceQueryPost**](MerchantBalanceAPI.md#MerchantBalanceMerchantBalanceQueryPost) | **Post** /merchant/balance/merchant_balance_query | Query Merchant Gateway Balance
[**MerchantBalanceUserBalanceQueryGet**](MerchantBalanceAPI.md#MerchantBalanceUserBalanceQueryGet) | **Get** /merchant/balance/user_balance_query | Query User Balance
[**MerchantBalanceUserBalanceQueryPost**](MerchantBalanceAPI.md#MerchantBalanceUserBalanceQueryPost) | **Post** /merchant/balance/user_balance_query | Query User Balance



## MerchantBalanceMerchantBalanceQueryGet

> MerchantBalanceMerchantBalanceQueryGet200Response MerchantBalanceMerchantBalanceQueryGet(ctx).GatewayId(gatewayId).Execute()

Query Merchant Gateway Balance

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
	gatewayId := int64(789) // int64 | gatewayId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryGet(context.Background()).GatewayId(gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBalanceMerchantBalanceQueryGet`: MerchantBalanceMerchantBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBalanceMerchantBalanceQueryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64** | gatewayId | 

### Return type

[**MerchantBalanceMerchantBalanceQueryGet200Response**](MerchantBalanceMerchantBalanceQueryGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantBalanceMerchantBalanceQueryPost

> MerchantBalanceMerchantBalanceQueryGet200Response MerchantBalanceMerchantBalanceQueryPost(ctx).UnibeeApiMerchantBalanceDetailQueryReq(unibeeApiMerchantBalanceDetailQueryReq).Execute()

Query Merchant Gateway Balance

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
	unibeeApiMerchantBalanceDetailQueryReq := *openapiclient.NewUnibeeApiMerchantBalanceDetailQueryReq(int64(123)) // UnibeeApiMerchantBalanceDetailQueryReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryPost(context.Background()).UnibeeApiMerchantBalanceDetailQueryReq(unibeeApiMerchantBalanceDetailQueryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBalanceMerchantBalanceQueryPost`: MerchantBalanceMerchantBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantBalanceAPI.MerchantBalanceMerchantBalanceQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBalanceMerchantBalanceQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantBalanceDetailQueryReq** | [**UnibeeApiMerchantBalanceDetailQueryReq**](UnibeeApiMerchantBalanceDetailQueryReq.md) |  | 

### Return type

[**MerchantBalanceMerchantBalanceQueryGet200Response**](MerchantBalanceMerchantBalanceQueryGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantBalanceUserBalanceQueryGet

> MerchantBalanceUserBalanceQueryGet200Response MerchantBalanceUserBalanceQueryGet(ctx).UserId(userId).GatewayId(gatewayId).Execute()

Query User Balance

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
	userId := int64(789) // int64 | userId
	gatewayId := int64(789) // int64 | gatewayId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantBalanceAPI.MerchantBalanceUserBalanceQueryGet(context.Background()).UserId(userId).GatewayId(gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantBalanceAPI.MerchantBalanceUserBalanceQueryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBalanceUserBalanceQueryGet`: MerchantBalanceUserBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantBalanceAPI.MerchantBalanceUserBalanceQueryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBalanceUserBalanceQueryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | userId | 
 **gatewayId** | **int64** | gatewayId | 

### Return type

[**MerchantBalanceUserBalanceQueryGet200Response**](MerchantBalanceUserBalanceQueryGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantBalanceUserBalanceQueryPost

> MerchantBalanceUserBalanceQueryGet200Response MerchantBalanceUserBalanceQueryPost(ctx).UnibeeApiMerchantBalanceUserDetailQueryReq(unibeeApiMerchantBalanceUserDetailQueryReq).Execute()

Query User Balance

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
	unibeeApiMerchantBalanceUserDetailQueryReq := *openapiclient.NewUnibeeApiMerchantBalanceUserDetailQueryReq(int64(123), int64(123)) // UnibeeApiMerchantBalanceUserDetailQueryReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantBalanceAPI.MerchantBalanceUserBalanceQueryPost(context.Background()).UnibeeApiMerchantBalanceUserDetailQueryReq(unibeeApiMerchantBalanceUserDetailQueryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantBalanceAPI.MerchantBalanceUserBalanceQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBalanceUserBalanceQueryPost`: MerchantBalanceUserBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantBalanceAPI.MerchantBalanceUserBalanceQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBalanceUserBalanceQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantBalanceUserDetailQueryReq** | [**UnibeeApiMerchantBalanceUserDetailQueryReq**](UnibeeApiMerchantBalanceUserDetailQueryReq.md) |  | 

### Return type

[**MerchantBalanceUserBalanceQueryGet200Response**](MerchantBalanceUserBalanceQueryGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

