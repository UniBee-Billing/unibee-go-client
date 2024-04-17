# \Balance

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalanceMerchantBalanceQueryGet**](Balance.md#BalanceMerchantBalanceQueryGet) | **Get** /merchant/balance/merchant_balance_query | Query Merchant Gateway Balance
[**BalanceUserBalanceQueryGet**](Balance.md#BalanceUserBalanceQueryGet) | **Get** /merchant/balance/user_balance_query | Query User Balance



## BalanceMerchantBalanceQueryGet

> MerchantBalanceMerchantBalanceQueryGet200Response BalanceMerchantBalanceQueryGet(ctx).GatewayId(gatewayId).Execute()

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
	resp, r, err := apiClient.Balance.BalanceMerchantBalanceQueryGet(context.Background()).GatewayId(gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Balance.BalanceMerchantBalanceQueryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceMerchantBalanceQueryGet`: MerchantBalanceMerchantBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Balance.BalanceMerchantBalanceQueryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceMerchantBalanceQueryGetRequest struct via the builder pattern


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


## BalanceUserBalanceQueryGet

> MerchantBalanceUserBalanceQueryGet200Response BalanceUserBalanceQueryGet(ctx).UserId(userId).GatewayId(gatewayId).Execute()

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
	resp, r, err := apiClient.Balance.BalanceUserBalanceQueryGet(context.Background()).UserId(userId).GatewayId(gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Balance.BalanceUserBalanceQueryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceUserBalanceQueryGet`: MerchantBalanceUserBalanceQueryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Balance.BalanceUserBalanceQueryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceUserBalanceQueryGetRequest struct via the builder pattern


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

