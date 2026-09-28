# UnibeeApiBeanUserChurnBulkCheckSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelReason** | Pointer to **string** | subscription.cancel_reason | [optional] 
**CurrentPeriodEnd** | Pointer to **int64** | Current period end, UTC seconds | [optional] 
**GmtModify** | Pointer to **int64** | Subscription gmt_modify, UTC seconds | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**Status** | Pointer to **string** | Subscription status name | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 

## Methods

### NewUnibeeApiBeanUserChurnBulkCheckSub

`func NewUnibeeApiBeanUserChurnBulkCheckSub() *UnibeeApiBeanUserChurnBulkCheckSub`

NewUnibeeApiBeanUserChurnBulkCheckSub instantiates a new UnibeeApiBeanUserChurnBulkCheckSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserChurnBulkCheckSubWithDefaults

`func NewUnibeeApiBeanUserChurnBulkCheckSubWithDefaults() *UnibeeApiBeanUserChurnBulkCheckSub`

NewUnibeeApiBeanUserChurnBulkCheckSubWithDefaults instantiates a new UnibeeApiBeanUserChurnBulkCheckSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelReason

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckSub) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


