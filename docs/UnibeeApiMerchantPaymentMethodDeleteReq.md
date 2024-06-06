# UnibeeApiMerchantPaymentMethodDeleteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **int64** | The unique id of gateway | 
**PaymentMethodId** | **string** | The unique id of payment method | 
**UserId** | **int64** | The customer&#39;s unique id | 

## Methods

### NewUnibeeApiMerchantPaymentMethodDeleteReq

`func NewUnibeeApiMerchantPaymentMethodDeleteReq(gatewayId int64, paymentMethodId string, userId int64, ) *UnibeeApiMerchantPaymentMethodDeleteReq`

NewUnibeeApiMerchantPaymentMethodDeleteReq instantiates a new UnibeeApiMerchantPaymentMethodDeleteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentMethodDeleteReqWithDefaults

`func NewUnibeeApiMerchantPaymentMethodDeleteReqWithDefaults() *UnibeeApiMerchantPaymentMethodDeleteReq`

NewUnibeeApiMerchantPaymentMethodDeleteReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodDeleteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetPaymentMethodId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetUserId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentMethodDeleteReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


