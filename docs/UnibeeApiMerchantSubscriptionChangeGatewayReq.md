# UnibeeApiMerchantSubscriptionChangeGatewayReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **int64** | GatewayId | 
**PaymentMethodId** | Pointer to **string** | PaymentMethodId | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 

## Methods

### NewUnibeeApiMerchantSubscriptionChangeGatewayReq

`func NewUnibeeApiMerchantSubscriptionChangeGatewayReq(gatewayId int64, subscriptionId string, ) *UnibeeApiMerchantSubscriptionChangeGatewayReq`

NewUnibeeApiMerchantSubscriptionChangeGatewayReq instantiates a new UnibeeApiMerchantSubscriptionChangeGatewayReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionChangeGatewayReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionChangeGatewayReqWithDefaults() *UnibeeApiMerchantSubscriptionChangeGatewayReq`

NewUnibeeApiMerchantSubscriptionChangeGatewayReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionChangeGatewayReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


