# UnibeeApiMerchantPlanActivePriceChangeConfirmRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAffectedSubscriptions** | Pointer to **int64** | Active/processing subscriptions whose future billing will be affected | [optional] 
**NewAmount** | Pointer to **int64** | New amount after change | [optional] 
**OldAmount** | Pointer to **int64** | Old amount before change | [optional] 
**PlanId** | Pointer to **int64** | Target planId | [optional] 
**TotalAffectedSubscriptions** | Pointer to **int64** | Total subscriptions that have used or are using this plan (by status set) | [optional] 

## Methods

### NewUnibeeApiMerchantPlanActivePriceChangeConfirmRes

`func NewUnibeeApiMerchantPlanActivePriceChangeConfirmRes() *UnibeeApiMerchantPlanActivePriceChangeConfirmRes`

NewUnibeeApiMerchantPlanActivePriceChangeConfirmRes instantiates a new UnibeeApiMerchantPlanActivePriceChangeConfirmRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanActivePriceChangeConfirmResWithDefaults

`func NewUnibeeApiMerchantPlanActivePriceChangeConfirmResWithDefaults() *UnibeeApiMerchantPlanActivePriceChangeConfirmRes`

NewUnibeeApiMerchantPlanActivePriceChangeConfirmResWithDefaults instantiates a new UnibeeApiMerchantPlanActivePriceChangeConfirmRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetActiveAffectedSubscriptions() int64`

GetActiveAffectedSubscriptions returns the ActiveAffectedSubscriptions field if non-nil, zero value otherwise.

### GetActiveAffectedSubscriptionsOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetActiveAffectedSubscriptionsOk() (*int64, bool)`

GetActiveAffectedSubscriptionsOk returns a tuple with the ActiveAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) SetActiveAffectedSubscriptions(v int64)`

SetActiveAffectedSubscriptions sets ActiveAffectedSubscriptions field to given value.

### HasActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) HasActiveAffectedSubscriptions() bool`

HasActiveAffectedSubscriptions returns a boolean if a field has been set.

### GetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetNewAmount() int64`

GetNewAmount returns the NewAmount field if non-nil, zero value otherwise.

### GetNewAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetNewAmountOk() (*int64, bool)`

GetNewAmountOk returns a tuple with the NewAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) SetNewAmount(v int64)`

SetNewAmount sets NewAmount field to given value.

### HasNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) HasNewAmount() bool`

HasNewAmount returns a boolean if a field has been set.

### GetOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetOldAmount() int64`

GetOldAmount returns the OldAmount field if non-nil, zero value otherwise.

### GetOldAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetOldAmountOk() (*int64, bool)`

GetOldAmountOk returns a tuple with the OldAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) SetOldAmount(v int64)`

SetOldAmount sets OldAmount field to given value.

### HasOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) HasOldAmount() bool`

HasOldAmount returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetTotalAffectedSubscriptions() int64`

GetTotalAffectedSubscriptions returns the TotalAffectedSubscriptions field if non-nil, zero value otherwise.

### GetTotalAffectedSubscriptionsOk

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) GetTotalAffectedSubscriptionsOk() (*int64, bool)`

GetTotalAffectedSubscriptionsOk returns a tuple with the TotalAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) SetTotalAffectedSubscriptions(v int64)`

SetTotalAffectedSubscriptions sets TotalAffectedSubscriptions field to given value.

### HasTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangeConfirmRes) HasTotalAffectedSubscriptions() bool`

HasTotalAffectedSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


