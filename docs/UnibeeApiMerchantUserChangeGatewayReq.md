# UnibeeApiMerchantUserChangeGatewayReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **int64** | GatewayId | 
**GatewayPaymentType** | Pointer to **string** | GatewayPaymentType | [optional] 
**GatewayUserId** | Pointer to **string** | GatewayUserId, verify and save GatewayUserId via gateway | [optional] 
**PaymentMethodId** | Pointer to **string** | PaymentMethodId of gateway, available for card type gateway, payment automatic will enable if set | [optional] 
**UserId** | **int64** | User Id | 

## Methods

### NewUnibeeApiMerchantUserChangeGatewayReq

`func NewUnibeeApiMerchantUserChangeGatewayReq(gatewayId int64, userId int64, ) *UnibeeApiMerchantUserChangeGatewayReq`

NewUnibeeApiMerchantUserChangeGatewayReq instantiates a new UnibeeApiMerchantUserChangeGatewayReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserChangeGatewayReqWithDefaults

`func NewUnibeeApiMerchantUserChangeGatewayReqWithDefaults() *UnibeeApiMerchantUserChangeGatewayReq`

NewUnibeeApiMerchantUserChangeGatewayReqWithDefaults instantiates a new UnibeeApiMerchantUserChangeGatewayReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetGatewayPaymentType

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantUserChangeGatewayReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetGatewayUserId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayUserId() string`

GetGatewayUserId returns the GatewayUserId field if non-nil, zero value otherwise.

### GetGatewayUserIdOk

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayUserIdOk() (*string, bool)`

GetGatewayUserIdOk returns a tuple with the GatewayUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUserId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayUserId(v string)`

SetGatewayUserId sets GatewayUserId field to given value.

### HasGatewayUserId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) HasGatewayUserId() bool`

HasGatewayUserId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserChangeGatewayReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserChangeGatewayReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


