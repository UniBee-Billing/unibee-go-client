# UnibeeApiBeanUserChurnBulkCheckEpisode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChurnTime** | Pointer to **int64** | Churn time, UTC seconds | [optional] 
**ChurnType** | Pointer to **string** | Cancelled|Expired | [optional] 
**Id** | Pointer to **int64** | Churn episode id | [optional] 
**ReactivatedTime** | Pointer to **int64** | Reactivated time, UTC seconds | [optional] 
**ReasonCategory** | Pointer to **string** | Normalized reason | [optional] 
**ReasonRaw** | Pointer to **string** | Raw cancel reason | [optional] 
**Status** | Pointer to **string** | pending|confirmed|recovered | [optional] 
**TriggerSubscriptionId** | Pointer to **string** | Trigger subscription id | [optional] 

## Methods

### NewUnibeeApiBeanUserChurnBulkCheckEpisode

`func NewUnibeeApiBeanUserChurnBulkCheckEpisode() *UnibeeApiBeanUserChurnBulkCheckEpisode`

NewUnibeeApiBeanUserChurnBulkCheckEpisode instantiates a new UnibeeApiBeanUserChurnBulkCheckEpisode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserChurnBulkCheckEpisodeWithDefaults

`func NewUnibeeApiBeanUserChurnBulkCheckEpisodeWithDefaults() *UnibeeApiBeanUserChurnBulkCheckEpisode`

NewUnibeeApiBeanUserChurnBulkCheckEpisodeWithDefaults instantiates a new UnibeeApiBeanUserChurnBulkCheckEpisode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChurnTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetChurnTime() int64`

GetChurnTime returns the ChurnTime field if non-nil, zero value otherwise.

### GetChurnTimeOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetChurnTimeOk() (*int64, bool)`

GetChurnTimeOk returns a tuple with the ChurnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetChurnTime(v int64)`

SetChurnTime sets ChurnTime field to given value.

### HasChurnTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasChurnTime() bool`

HasChurnTime returns a boolean if a field has been set.

### GetChurnType

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetChurnType() string`

GetChurnType returns the ChurnType field if non-nil, zero value otherwise.

### GetChurnTypeOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetChurnTypeOk() (*string, bool)`

GetChurnTypeOk returns a tuple with the ChurnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnType

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetChurnType(v string)`

SetChurnType sets ChurnType field to given value.

### HasChurnType

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasChurnType() bool`

HasChurnType returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReactivatedTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReactivatedTime() int64`

GetReactivatedTime returns the ReactivatedTime field if non-nil, zero value otherwise.

### GetReactivatedTimeOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReactivatedTimeOk() (*int64, bool)`

GetReactivatedTimeOk returns a tuple with the ReactivatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatedTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetReactivatedTime(v int64)`

SetReactivatedTime sets ReactivatedTime field to given value.

### HasReactivatedTime

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasReactivatedTime() bool`

HasReactivatedTime returns a boolean if a field has been set.

### GetReasonCategory

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReasonCategory() string`

GetReasonCategory returns the ReasonCategory field if non-nil, zero value otherwise.

### GetReasonCategoryOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReasonCategoryOk() (*string, bool)`

GetReasonCategoryOk returns a tuple with the ReasonCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCategory

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetReasonCategory(v string)`

SetReasonCategory sets ReasonCategory field to given value.

### HasReasonCategory

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasReasonCategory() bool`

HasReasonCategory returns a boolean if a field has been set.

### GetReasonRaw

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReasonRaw() string`

GetReasonRaw returns the ReasonRaw field if non-nil, zero value otherwise.

### GetReasonRawOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetReasonRawOk() (*string, bool)`

GetReasonRawOk returns a tuple with the ReasonRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRaw

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetReasonRaw(v string)`

SetReasonRaw sets ReasonRaw field to given value.

### HasReasonRaw

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasReasonRaw() bool`

HasReasonRaw returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetTriggerSubscriptionId() string`

GetTriggerSubscriptionId returns the TriggerSubscriptionId field if non-nil, zero value otherwise.

### GetTriggerSubscriptionIdOk

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) GetTriggerSubscriptionIdOk() (*string, bool)`

GetTriggerSubscriptionIdOk returns a tuple with the TriggerSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) SetTriggerSubscriptionId(v string)`

SetTriggerSubscriptionId sets TriggerSubscriptionId field to given value.

### HasTriggerSubscriptionId

`func (o *UnibeeApiBeanUserChurnBulkCheckEpisode) HasTriggerSubscriptionId() bool`

HasTriggerSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


