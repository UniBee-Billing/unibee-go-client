# UnibeeApiMerchantPaymentPaymentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**UnibeeInternalLogicGatewayRoPaymentSimplify**](UnibeeInternalLogicGatewayRoPaymentSimplify.md) |  | [optional] 
**User** | Pointer to [**UnibeeInternalLogicGatewayRoUserAccountSimplify**](UnibeeInternalLogicGatewayRoUserAccountSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentPaymentDetail

`func NewUnibeeApiMerchantPaymentPaymentDetail() *UnibeeApiMerchantPaymentPaymentDetail`

NewUnibeeApiMerchantPaymentPaymentDetail instantiates a new UnibeeApiMerchantPaymentPaymentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentPaymentDetailWithDefaults

`func NewUnibeeApiMerchantPaymentPaymentDetailWithDefaults() *UnibeeApiMerchantPaymentPaymentDetail`

NewUnibeeApiMerchantPaymentPaymentDetailWithDefaults instantiates a new UnibeeApiMerchantPaymentPaymentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *UnibeeApiMerchantPaymentPaymentDetail) GetPayment() UnibeeInternalLogicGatewayRoPaymentSimplify`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiMerchantPaymentPaymentDetail) GetPaymentOk() (*UnibeeInternalLogicGatewayRoPaymentSimplify, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiMerchantPaymentPaymentDetail) SetPayment(v UnibeeInternalLogicGatewayRoPaymentSimplify)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiMerchantPaymentPaymentDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantPaymentPaymentDetail) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantPaymentPaymentDetail) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantPaymentPaymentDetail) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantPaymentPaymentDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


