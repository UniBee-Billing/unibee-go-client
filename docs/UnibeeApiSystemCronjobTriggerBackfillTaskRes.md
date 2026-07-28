# UnibeeApiSystemCronjobTriggerBackfillTaskRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTasks** | Pointer to **[]string** | All task names accepted by this endpoint | [optional] 
**BatchSizeHint** | Pointer to **string** | Approximate rows processed per invocation | [optional] 
**LockTaskName** | Pointer to **string** | Redis lock scope shared with @hourly cron | [optional] 
**Message** | Pointer to **string** | Human-readable result | [optional] 
**TaskName** | Pointer to **string** | Task that was executed | [optional] 

## Methods

### NewUnibeeApiSystemCronjobTriggerBackfillTaskRes

`func NewUnibeeApiSystemCronjobTriggerBackfillTaskRes() *UnibeeApiSystemCronjobTriggerBackfillTaskRes`

NewUnibeeApiSystemCronjobTriggerBackfillTaskRes instantiates a new UnibeeApiSystemCronjobTriggerBackfillTaskRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemCronjobTriggerBackfillTaskResWithDefaults

`func NewUnibeeApiSystemCronjobTriggerBackfillTaskResWithDefaults() *UnibeeApiSystemCronjobTriggerBackfillTaskRes`

NewUnibeeApiSystemCronjobTriggerBackfillTaskResWithDefaults instantiates a new UnibeeApiSystemCronjobTriggerBackfillTaskRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTasks

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetAllowedTasks() []string`

GetAllowedTasks returns the AllowedTasks field if non-nil, zero value otherwise.

### GetAllowedTasksOk

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetAllowedTasksOk() (*[]string, bool)`

GetAllowedTasksOk returns a tuple with the AllowedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTasks

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) SetAllowedTasks(v []string)`

SetAllowedTasks sets AllowedTasks field to given value.

### HasAllowedTasks

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) HasAllowedTasks() bool`

HasAllowedTasks returns a boolean if a field has been set.

### GetBatchSizeHint

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetBatchSizeHint() string`

GetBatchSizeHint returns the BatchSizeHint field if non-nil, zero value otherwise.

### GetBatchSizeHintOk

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetBatchSizeHintOk() (*string, bool)`

GetBatchSizeHintOk returns a tuple with the BatchSizeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSizeHint

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) SetBatchSizeHint(v string)`

SetBatchSizeHint sets BatchSizeHint field to given value.

### HasBatchSizeHint

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) HasBatchSizeHint() bool`

HasBatchSizeHint returns a boolean if a field has been set.

### GetLockTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetLockTaskName() string`

GetLockTaskName returns the LockTaskName field if non-nil, zero value otherwise.

### GetLockTaskNameOk

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetLockTaskNameOk() (*string, bool)`

GetLockTaskNameOk returns a tuple with the LockTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) SetLockTaskName(v string)`

SetLockTaskName sets LockTaskName field to given value.

### HasLockTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) HasLockTaskName() bool`

HasLockTaskName returns a boolean if a field has been set.

### GetMessage

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *UnibeeApiSystemCronjobTriggerBackfillTaskRes) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


