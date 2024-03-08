# UnibeeApiMerchantPaymentRefundListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**GatewayId** | Pointer to **int64** | GatewayId | [optional] 
**PaymentId** | **string** | PaymentId | 
**Status** | Pointer to **int32** | Status,10-create|20-success|30-Failed|40-Reverse | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentRefundListReq

`func NewUnibeeApiMerchantPaymentRefundListReq(paymentId string, ) *UnibeeApiMerchantPaymentRefundListReq`

NewUnibeeApiMerchantPaymentRefundListReq instantiates a new UnibeeApiMerchantPaymentRefundListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentRefundListReqWithDefaults

`func NewUnibeeApiMerchantPaymentRefundListReqWithDefaults() *UnibeeApiMerchantPaymentRefundListReq`

NewUnibeeApiMerchantPaymentRefundListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPaymentRefundListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantPaymentRefundListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentRefundListReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetStatus

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantPaymentRefundListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentRefundListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentRefundListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantPaymentRefundListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


