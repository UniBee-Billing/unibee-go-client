# UnibeeApiMerchantPaymentRefundDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**UnibeeInternalLogicGatewayRoPaymentSimplify**](UnibeeInternalLogicGatewayRoPaymentSimplify.md) |  | [optional] 
**Refund** | Pointer to [**UnibeeInternalLogicGatewayRoRefundSimplify**](UnibeeInternalLogicGatewayRoRefundSimplify.md) |  | [optional] 
**User** | Pointer to [**UnibeeInternalLogicGatewayRoUserAccountSimplify**](UnibeeInternalLogicGatewayRoUserAccountSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentRefundDetail

`func NewUnibeeApiMerchantPaymentRefundDetail() *UnibeeApiMerchantPaymentRefundDetail`

NewUnibeeApiMerchantPaymentRefundDetail instantiates a new UnibeeApiMerchantPaymentRefundDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentRefundDetailWithDefaults

`func NewUnibeeApiMerchantPaymentRefundDetailWithDefaults() *UnibeeApiMerchantPaymentRefundDetail`

NewUnibeeApiMerchantPaymentRefundDetailWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetPayment() UnibeeInternalLogicGatewayRoPaymentSimplify`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetPaymentOk() (*UnibeeInternalLogicGatewayRoPaymentSimplify, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiMerchantPaymentRefundDetail) SetPayment(v UnibeeInternalLogicGatewayRoPaymentSimplify)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiMerchantPaymentRefundDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetRefund() UnibeeInternalLogicGatewayRoRefundSimplify`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetRefundOk() (*UnibeeInternalLogicGatewayRoRefundSimplify, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeApiMerchantPaymentRefundDetail) SetRefund(v UnibeeInternalLogicGatewayRoRefundSimplify)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeApiMerchantPaymentRefundDetail) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantPaymentRefundDetail) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantPaymentRefundDetail) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantPaymentRefundDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


