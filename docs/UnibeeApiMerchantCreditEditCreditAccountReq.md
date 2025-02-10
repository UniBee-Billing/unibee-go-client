# UnibeeApiMerchantCreditEditCreditAccountReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | id of credit account | 
**PayoutEnable** | Pointer to **int32** | credit account can used or payout|apply in purchase or not, 0-no, 1-yes | [optional] 
**RechargeEnable** | Pointer to **int32** | credit account can be recharged|increment or not, 0-no, 1-yes | [optional] 

## Methods

### NewUnibeeApiMerchantCreditEditCreditAccountReq

`func NewUnibeeApiMerchantCreditEditCreditAccountReq(id int64, ) *UnibeeApiMerchantCreditEditCreditAccountReq`

NewUnibeeApiMerchantCreditEditCreditAccountReq instantiates a new UnibeeApiMerchantCreditEditCreditAccountReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditEditCreditAccountReqWithDefaults

`func NewUnibeeApiMerchantCreditEditCreditAccountReqWithDefaults() *UnibeeApiMerchantCreditEditCreditAccountReq`

NewUnibeeApiMerchantCreditEditCreditAccountReqWithDefaults instantiates a new UnibeeApiMerchantCreditEditCreditAccountReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetId(v int64)`

SetId sets Id field to given value.


### GetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiMerchantCreditEditCreditAccountReq) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


