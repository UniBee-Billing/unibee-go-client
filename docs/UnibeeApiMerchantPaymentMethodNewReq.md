# UnibeeApiMerchantPaymentMethodNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of payment method | [optional] 
**GatewayId** | **int64** | The unique id of gateway | 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**RedirectUrl** | Pointer to **string** | The redirect url when method created return back | [optional] 
**SubscriptionId** | Pointer to **string** | The id of subscription that want to attach | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | **int64** | The customer&#39;s unique id | 

## Methods

### NewUnibeeApiMerchantPaymentMethodNewReq

`func NewUnibeeApiMerchantPaymentMethodNewReq(gatewayId int64, userId int64, ) *UnibeeApiMerchantPaymentMethodNewReq`

NewUnibeeApiMerchantPaymentMethodNewReq instantiates a new UnibeeApiMerchantPaymentMethodNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentMethodNewReqWithDefaults

`func NewUnibeeApiMerchantPaymentMethodNewReqWithDefaults() *UnibeeApiMerchantPaymentMethodNewReq`

NewUnibeeApiMerchantPaymentMethodNewReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPaymentMethodNewReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantPaymentMethodNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UnibeeApiMerchantPaymentMethodNewReq) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantPaymentMethodNewReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentMethodNewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentMethodNewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


