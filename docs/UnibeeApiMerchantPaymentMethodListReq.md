# UnibeeApiMerchantPaymentMethodListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **int64** | GatewayId | 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentMethodListReq

`func NewUnibeeApiMerchantPaymentMethodListReq(gatewayId int64, ) *UnibeeApiMerchantPaymentMethodListReq`

NewUnibeeApiMerchantPaymentMethodListReq instantiates a new UnibeeApiMerchantPaymentMethodListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentMethodListReqWithDefaults

`func NewUnibeeApiMerchantPaymentMethodListReqWithDefaults() *UnibeeApiMerchantPaymentMethodListReq`

NewUnibeeApiMerchantPaymentMethodListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodListReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetPaymentId

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentMethodListReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiMerchantPaymentMethodListReq) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentMethodListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentMethodListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantPaymentMethodListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


