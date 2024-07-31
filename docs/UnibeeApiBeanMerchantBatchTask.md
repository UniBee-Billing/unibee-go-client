# UnibeeApiBeanMerchantBatchTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**DownloadUrl** | Pointer to **string** | download_file_url | [optional] 
**FailureReason** | Pointer to **string** | reason of failure | [optional] 
**FinishTime** | Pointer to **int64** | task_finish_time | [optional] 
**Format** | Pointer to **string** | format of file | [optional] 
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

### NewUnibeeApiBeanMerchantBatchTask

`func NewUnibeeApiBeanMerchantBatchTask() *UnibeeApiBeanMerchantBatchTask`

NewUnibeeApiBeanMerchantBatchTask instantiates a new UnibeeApiBeanMerchantBatchTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchTaskWithDefaults

`func NewUnibeeApiBeanMerchantBatchTaskWithDefaults() *UnibeeApiBeanMerchantBatchTask`

NewUnibeeApiBeanMerchantBatchTaskWithDefaults instantiates a new UnibeeApiBeanMerchantBatchTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchTask) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchTask) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchTask) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTask) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTask) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *UnibeeApiBeanMerchantBatchTask) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeApiBeanMerchantBatchTask) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiBeanMerchantBatchTask) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeApiBeanMerchantBatchTask) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanMerchantBatchTask) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanMerchantBatchTask) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanMerchantBatchTask) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiBeanMerchantBatchTask) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiBeanMerchantBatchTask) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiBeanMerchantBatchTask) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantBatchTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantBatchTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantBatchTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTask) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTask) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeApiBeanMerchantBatchTask) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiBeanMerchantBatchTask) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiBeanMerchantBatchTask) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiBeanMerchantBatchTask) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchTask) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchTask) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchTask) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiBeanMerchantBatchTask) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiBeanMerchantBatchTask) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiBeanMerchantBatchTask) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantBatchTask) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantBatchTask) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantBatchTask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantBatchTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantBatchTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantBatchTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTask) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTask) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiBeanMerchantBatchTask) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTaskCost

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskCost() int32`

GetTaskCost returns the TaskCost field if non-nil, zero value otherwise.

### GetTaskCostOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskCostOk() (*int32, bool)`

GetTaskCostOk returns a tuple with the TaskCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCost

`func (o *UnibeeApiBeanMerchantBatchTask) SetTaskCost(v int32)`

SetTaskCost sets TaskCost field to given value.

### HasTaskCost

`func (o *UnibeeApiBeanMerchantBatchTask) HasTaskCost() bool`

HasTaskCost returns a boolean if a field has been set.

### GetTaskName

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *UnibeeApiBeanMerchantBatchTask) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *UnibeeApiBeanMerchantBatchTask) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskType

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskType() int32`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetTaskTypeOk() (*int32, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UnibeeApiBeanMerchantBatchTask) SetTaskType(v int32)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UnibeeApiBeanMerchantBatchTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTask) GetUploadFileUrl() string`

GetUploadFileUrl returns the UploadFileUrl field if non-nil, zero value otherwise.

### GetUploadFileUrlOk

`func (o *UnibeeApiBeanMerchantBatchTask) GetUploadFileUrlOk() (*string, bool)`

GetUploadFileUrlOk returns a tuple with the UploadFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTask) SetUploadFileUrl(v string)`

SetUploadFileUrl sets UploadFileUrl field to given value.

### HasUploadFileUrl

`func (o *UnibeeApiBeanMerchantBatchTask) HasUploadFileUrl() bool`

HasUploadFileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


