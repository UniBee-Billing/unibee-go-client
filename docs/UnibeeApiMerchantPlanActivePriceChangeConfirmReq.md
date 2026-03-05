# UnibeeApiMerchantPlanActivePriceChangeConfirmReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmOldAmount** | **int64** | Old amount expected by caller, used for optimistic locking | 
**NewAmount** | **int64** | New amount (in smallest currency unit, e.g. cents) | 
**PlanId** | **int64** | Target planId, must be active | 
**Reason** | Pointer to **string** | Reason or note for this price change | [optional] 

## Methods

### NewUnibeeApiMerchantPlanActivePriceChangeConfirmReq

`func NewUnibeeApiMerchantPlanActivePriceChangeConfirmReq(confirmOldAmount int64, newAmount int64, planId int64, ) *UnibeeApiMerchantPlanActivePriceChangeConfirmReq`

NewUnibeeApiMerchantPlanActivePriceChangeConfirmReq instantiates a new UnibeeApiMerchantPlanActivePriceChangeConfirmReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanActivePriceChangeConfirmReqWithDefaults

`func NewUnibeeApiMerchantPlanActivePriceChangeConfirmReqWithDefaults() *UnibeeApiMerchantPlanActivePriceChangeConfirmReq`

NewUnibeeApiMerchantPlanActivePriceChangeConfirmReqWithDefaults instantiates a new UnibeeApiMerchantPlanActivePriceChangeConfirmReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetConfirmOldAmount() int64`

GetConfirmOldAmount returns the ConfirmOldAmount field if non-nil, zero value otherwise.

### GetConfirmOldAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetConfirmOldAmountOk() (*int64, bool)`

GetConfirmOldAmountOk returns a tuple with the ConfirmOldAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) SetConfirmOldAmount(v int64)`

SetConfirmOldAmount sets ConfirmOldAmount field to given value.


### GetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetNewAmount() int64`

GetNewAmount returns the NewAmount field if non-nil, zero value otherwise.

### GetNewAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetNewAmountOk() (*int64, bool)`

GetNewAmountOk returns a tuple with the NewAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) SetNewAmount(v int64)`

SetNewAmount sets NewAmount field to given value.


### GetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetReason

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmReq) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


