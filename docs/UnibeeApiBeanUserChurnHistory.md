# UnibeeApiBeanUserChurnHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChurnTime** | Pointer to **int64** | Original churn time, UTC seconds | [optional] 
**ChurnType** | Pointer to **string** | Cancelled|Expired | [optional] 
**ConfirmTime** | Pointer to **int64** | Confirm time, UTC seconds | [optional] 
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**DaysToReactivate** | Pointer to **int32** | Days from churn to reactivate | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**ProductId** | Pointer to **int64** | ProductId | [optional] 
**ProductName** | Pointer to **string** | ProductName | [optional] 
**Reactivated** | Pointer to **bool** | Whether reactivated | [optional] 
**ReactivatedTime** | Pointer to **int64** | Reactivated time, UTC seconds | [optional] 
**ReactivatingPlanId** | Pointer to **int64** | Reactivating plan id | [optional] 
**ReactivatingSubscriptionId** | Pointer to **string** | Reactivating subscription id | [optional] 
**ReasonCategory** | Pointer to **string** | Normalized reason | [optional] 
**ReasonRaw** | Pointer to **string** | Raw cancel reason | [optional] 
**Status** | Pointer to **string** | pending|confirmed|recovered | [optional] 
**TriggerSubscriptionId** | Pointer to **string** | Trigger subscription id | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiBeanUserChurnHistory

`func NewUnibeeApiBeanUserChurnHistory() *UnibeeApiBeanUserChurnHistory`

NewUnibeeApiBeanUserChurnHistory instantiates a new UnibeeApiBeanUserChurnHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserChurnHistoryWithDefaults

`func NewUnibeeApiBeanUserChurnHistoryWithDefaults() *UnibeeApiBeanUserChurnHistory`

NewUnibeeApiBeanUserChurnHistoryWithDefaults instantiates a new UnibeeApiBeanUserChurnHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChurnTime

`func (o *UnibeeApiBeanUserChurnHistory) GetChurnTime() int64`

GetChurnTime returns the ChurnTime field if non-nil, zero value otherwise.

### GetChurnTimeOk

`func (o *UnibeeApiBeanUserChurnHistory) GetChurnTimeOk() (*int64, bool)`

GetChurnTimeOk returns a tuple with the ChurnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnTime

`func (o *UnibeeApiBeanUserChurnHistory) SetChurnTime(v int64)`

SetChurnTime sets ChurnTime field to given value.

### HasChurnTime

`func (o *UnibeeApiBeanUserChurnHistory) HasChurnTime() bool`

HasChurnTime returns a boolean if a field has been set.

### GetChurnType

`func (o *UnibeeApiBeanUserChurnHistory) GetChurnType() string`

GetChurnType returns the ChurnType field if non-nil, zero value otherwise.

### GetChurnTypeOk

`func (o *UnibeeApiBeanUserChurnHistory) GetChurnTypeOk() (*string, bool)`

GetChurnTypeOk returns a tuple with the ChurnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnType

`func (o *UnibeeApiBeanUserChurnHistory) SetChurnType(v string)`

SetChurnType sets ChurnType field to given value.

### HasChurnType

`func (o *UnibeeApiBeanUserChurnHistory) HasChurnType() bool`

HasChurnType returns a boolean if a field has been set.

### GetConfirmTime

`func (o *UnibeeApiBeanUserChurnHistory) GetConfirmTime() int64`

GetConfirmTime returns the ConfirmTime field if non-nil, zero value otherwise.

### GetConfirmTimeOk

`func (o *UnibeeApiBeanUserChurnHistory) GetConfirmTimeOk() (*int64, bool)`

GetConfirmTimeOk returns a tuple with the ConfirmTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTime

`func (o *UnibeeApiBeanUserChurnHistory) SetConfirmTime(v int64)`

SetConfirmTime sets ConfirmTime field to given value.

### HasConfirmTime

`func (o *UnibeeApiBeanUserChurnHistory) HasConfirmTime() bool`

HasConfirmTime returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanUserChurnHistory) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanUserChurnHistory) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanUserChurnHistory) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanUserChurnHistory) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanUserChurnHistory) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanUserChurnHistory) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanUserChurnHistory) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanUserChurnHistory) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDaysToReactivate

`func (o *UnibeeApiBeanUserChurnHistory) GetDaysToReactivate() int32`

GetDaysToReactivate returns the DaysToReactivate field if non-nil, zero value otherwise.

### GetDaysToReactivateOk

`func (o *UnibeeApiBeanUserChurnHistory) GetDaysToReactivateOk() (*int32, bool)`

GetDaysToReactivateOk returns a tuple with the DaysToReactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToReactivate

`func (o *UnibeeApiBeanUserChurnHistory) SetDaysToReactivate(v int32)`

SetDaysToReactivate sets DaysToReactivate field to given value.

### HasDaysToReactivate

`func (o *UnibeeApiBeanUserChurnHistory) HasDaysToReactivate() bool`

HasDaysToReactivate returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanUserChurnHistory) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanUserChurnHistory) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanUserChurnHistory) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanUserChurnHistory) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiBeanUserChurnHistory) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiBeanUserChurnHistory) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiBeanUserChurnHistory) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanUserChurnHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanUserChurnHistory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanUserChurnHistory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanUserChurnHistory) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanUserChurnHistory) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanUserChurnHistory) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiBeanUserChurnHistory) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiBeanUserChurnHistory) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiBeanUserChurnHistory) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiBeanUserChurnHistory) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiBeanUserChurnHistory) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiBeanUserChurnHistory) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiBeanUserChurnHistory) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanUserChurnHistory) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanUserChurnHistory) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanUserChurnHistory) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanUserChurnHistory) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetReactivated

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivated() bool`

GetReactivated returns the Reactivated field if non-nil, zero value otherwise.

### GetReactivatedOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatedOk() (*bool, bool)`

GetReactivatedOk returns a tuple with the Reactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivated

`func (o *UnibeeApiBeanUserChurnHistory) SetReactivated(v bool)`

SetReactivated sets Reactivated field to given value.

### HasReactivated

`func (o *UnibeeApiBeanUserChurnHistory) HasReactivated() bool`

HasReactivated returns a boolean if a field has been set.

### GetReactivatedTime

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatedTime() int64`

GetReactivatedTime returns the ReactivatedTime field if non-nil, zero value otherwise.

### GetReactivatedTimeOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatedTimeOk() (*int64, bool)`

GetReactivatedTimeOk returns a tuple with the ReactivatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatedTime

`func (o *UnibeeApiBeanUserChurnHistory) SetReactivatedTime(v int64)`

SetReactivatedTime sets ReactivatedTime field to given value.

### HasReactivatedTime

`func (o *UnibeeApiBeanUserChurnHistory) HasReactivatedTime() bool`

HasReactivatedTime returns a boolean if a field has been set.

### GetReactivatingPlanId

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatingPlanId() int64`

GetReactivatingPlanId returns the ReactivatingPlanId field if non-nil, zero value otherwise.

### GetReactivatingPlanIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatingPlanIdOk() (*int64, bool)`

GetReactivatingPlanIdOk returns a tuple with the ReactivatingPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatingPlanId

`func (o *UnibeeApiBeanUserChurnHistory) SetReactivatingPlanId(v int64)`

SetReactivatingPlanId sets ReactivatingPlanId field to given value.

### HasReactivatingPlanId

`func (o *UnibeeApiBeanUserChurnHistory) HasReactivatingPlanId() bool`

HasReactivatingPlanId returns a boolean if a field has been set.

### GetReactivatingSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatingSubscriptionId() string`

GetReactivatingSubscriptionId returns the ReactivatingSubscriptionId field if non-nil, zero value otherwise.

### GetReactivatingSubscriptionIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReactivatingSubscriptionIdOk() (*string, bool)`

GetReactivatingSubscriptionIdOk returns a tuple with the ReactivatingSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatingSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) SetReactivatingSubscriptionId(v string)`

SetReactivatingSubscriptionId sets ReactivatingSubscriptionId field to given value.

### HasReactivatingSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) HasReactivatingSubscriptionId() bool`

HasReactivatingSubscriptionId returns a boolean if a field has been set.

### GetReasonCategory

`func (o *UnibeeApiBeanUserChurnHistory) GetReasonCategory() string`

GetReasonCategory returns the ReasonCategory field if non-nil, zero value otherwise.

### GetReasonCategoryOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReasonCategoryOk() (*string, bool)`

GetReasonCategoryOk returns a tuple with the ReasonCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCategory

`func (o *UnibeeApiBeanUserChurnHistory) SetReasonCategory(v string)`

SetReasonCategory sets ReasonCategory field to given value.

### HasReasonCategory

`func (o *UnibeeApiBeanUserChurnHistory) HasReasonCategory() bool`

HasReasonCategory returns a boolean if a field has been set.

### GetReasonRaw

`func (o *UnibeeApiBeanUserChurnHistory) GetReasonRaw() string`

GetReasonRaw returns the ReasonRaw field if non-nil, zero value otherwise.

### GetReasonRawOk

`func (o *UnibeeApiBeanUserChurnHistory) GetReasonRawOk() (*string, bool)`

GetReasonRawOk returns a tuple with the ReasonRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRaw

`func (o *UnibeeApiBeanUserChurnHistory) SetReasonRaw(v string)`

SetReasonRaw sets ReasonRaw field to given value.

### HasReasonRaw

`func (o *UnibeeApiBeanUserChurnHistory) HasReasonRaw() bool`

HasReasonRaw returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanUserChurnHistory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanUserChurnHistory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanUserChurnHistory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanUserChurnHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) GetTriggerSubscriptionId() string`

GetTriggerSubscriptionId returns the TriggerSubscriptionId field if non-nil, zero value otherwise.

### GetTriggerSubscriptionIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetTriggerSubscriptionIdOk() (*string, bool)`

GetTriggerSubscriptionIdOk returns a tuple with the TriggerSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) SetTriggerSubscriptionId(v string)`

SetTriggerSubscriptionId sets TriggerSubscriptionId field to given value.

### HasTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnHistory) HasTriggerSubscriptionId() bool`

HasTriggerSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanUserChurnHistory) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanUserChurnHistory) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanUserChurnHistory) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanUserChurnHistory) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


