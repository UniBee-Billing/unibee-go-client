# UnibeeApiBeanMerchantBatchTaskSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**DownloadUrl** | Pointer to **string** | download_file_url | [optional] 
**FailReason** | Pointer to **string** | reason of failure | [optional] 
**FinishTime** | Pointer to **int64** | task_finish_time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**LastUpdateTime** | Pointer to **int64** | last update utc time | [optional] 
**MemberId** | Pointer to **int64** | member_id | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Payload** | Pointer to **string** | payload(json) | [optional] 
**StartTime** | Pointer to **int64** | task_start_time | [optional] 
**Status** | Pointer to **int32** | Status。0-Pending，1-Processing，2-Success，3-Failure | [optional] 
**SuccessCount** | Pointer to **int64** | success_count | [optional] 
**TaskCost** | Pointer to **int32** | task cost time(second) | [optional] 
**TaskName** | Pointer to **string** | task_name | [optional] 
**TaskType** | Pointer to **int32** | type，0-download，1-upload | [optional] 
**UploadFileUrl** | Pointer to **string** | the file url of upload type task | [optional] 

## Methods

### NewUnibeeApiBeanMerchantBatchTaskSimplify

`func NewUnibeeApiBeanMerchantBatchTaskSimplify() *UnibeeApiBeanMerchantBatchTaskSimplify`

NewUnibeeApiBeanMerchantBatchTaskSimplify instantiates a new UnibeeApiBeanMerchantBatchTaskSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchTaskSimplifyWithDefaults

`func NewUnibeeApiBeanMerchantBatchTaskSimplifyWithDefaults() *UnibeeApiBeanMerchantBatchTaskSimplify`

NewUnibeeApiBeanMerchantBatchTaskSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantBatchTaskSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetFailReason

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetFailReason() string`

GetFailReason returns the FailReason field if non-nil, zero value otherwise.

### GetFailReasonOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetFailReasonOk() (*string, bool)`

GetFailReasonOk returns a tuple with the FailReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailReason

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetFailReason(v string)`

SetFailReason sets FailReason field to given value.

### HasFailReason

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasFailReason() bool`

HasFailReason returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTaskCost

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskCost() int32`

GetTaskCost returns the TaskCost field if non-nil, zero value otherwise.

### GetTaskCostOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskCostOk() (*int32, bool)`

GetTaskCostOk returns a tuple with the TaskCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCost

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetTaskCost(v int32)`

SetTaskCost sets TaskCost field to given value.

### HasTaskCost

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasTaskCost() bool`

HasTaskCost returns a boolean if a field has been set.

### GetTaskName

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskType

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskType() int32`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetTaskTypeOk() (*int32, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetTaskType(v int32)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetUploadFileUrl() string`

GetUploadFileUrl returns the UploadFileUrl field if non-nil, zero value otherwise.

### GetUploadFileUrlOk

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) GetUploadFileUrlOk() (*string, bool)`

GetUploadFileUrlOk returns a tuple with the UploadFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) SetUploadFileUrl(v string)`

SetUploadFileUrl sets UploadFileUrl field to given value.

### HasUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTaskSimplify) HasUploadFileUrl() bool`

HasUploadFileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


