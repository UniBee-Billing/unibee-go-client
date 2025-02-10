# UnibeeApiMerchantCreditNewCreditRechargeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** | CancelUrl | [optional] 
**CreditAccountId** | Pointer to **int64** | id of credit account, either userId&amp;currency or creditAccountId  | [optional] 
**Currency** | Pointer to **string** | currency of recharge | [optional] 
**Description** | Pointer to **string** | recharge description | [optional] 
**GatewayId** | **int64** |  | 
**Name** | Pointer to **string** | recharge name | [optional] 
**RechargeAmount** | **int64** |  | 
**ReturnUrl** | Pointer to **string** | ReturnUrl | [optional] 
**UserId** | Pointer to **int64** | id of user to recharge, either userId&amp;currency or creditAccountId  | [optional] 

## Methods

### NewUnibeeApiMerchantCreditNewCreditRechargeReq

`func NewUnibeeApiMerchantCreditNewCreditRechargeReq(gatewayId int64, rechargeAmount int64, ) *UnibeeApiMerchantCreditNewCreditRechargeReq`

NewUnibeeApiMerchantCreditNewCreditRechargeReq instantiates a new UnibeeApiMerchantCreditNewCreditRechargeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditNewCreditRechargeReqWithDefaults

`func NewUnibeeApiMerchantCreditNewCreditRechargeReqWithDefaults() *UnibeeApiMerchantCreditNewCreditRechargeReq`

NewUnibeeApiMerchantCreditNewCreditRechargeReqWithDefaults instantiates a new UnibeeApiMerchantCreditNewCreditRechargeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCreditAccountId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCreditAccountId() int64`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCreditAccountIdOk() (*int64, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCreditAccountId(v int64)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetName

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRechargeAmount

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetRechargeAmount() int64`

GetRechargeAmount returns the RechargeAmount field if non-nil, zero value otherwise.

### GetRechargeAmountOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetRechargeAmountOk() (*int64, bool)`

GetRechargeAmountOk returns a tuple with the RechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeAmount

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetRechargeAmount(v int64)`

SetRechargeAmount sets RechargeAmount field to given value.


### GetReturnUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


