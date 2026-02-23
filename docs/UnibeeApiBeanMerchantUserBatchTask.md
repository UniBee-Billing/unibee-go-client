# UnibeeApiBeanMerchantUserBatchTask

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
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Payload** | Pointer to **string** | payload(json) | [optional] 
**StartTime** | Pointer to **int64** | task_start_time | [optional] 
**Status** | Pointer to **int32** | Status。0-Pending，1-Processing，2-Success，3-Failure | [optional] 
**SuccessCount** | Pointer to **int64** | success_count | [optional] 
**TaskCost** | Pointer to **int32** | task cost time(second) | [optional] 
**TaskName** | Pointer to **string** | task_name | [optional] 
**TaskType** | Pointer to **int32** | type，0-download，1-upload | [optional] 
**UploadFileUrl** | Pointer to **string** | the file url of upload type task | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanMerchantUserBatchTask

`func NewUnibeeApiBeanMerchantUserBatchTask() *UnibeeApiBeanMerchantUserBatchTask`

NewUnibeeApiBeanMerchantUserBatchTask instantiates a new UnibeeApiBeanMerchantUserBatchTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantUserBatchTaskWithDefaults

`func NewUnibeeApiBeanMerchantUserBatchTaskWithDefaults() *UnibeeApiBeanMerchantUserBatchTask`

NewUnibeeApiBeanMerchantUserBatchTaskWithDefaults instantiates a new UnibeeApiBeanMerchantUserBatchTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTaskCost

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskCost() int32`

GetTaskCost returns the TaskCost field if non-nil, zero value otherwise.

### GetTaskCostOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskCostOk() (*int32, bool)`

GetTaskCostOk returns a tuple with the TaskCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCost

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetTaskCost(v int32)`

SetTaskCost sets TaskCost field to given value.

### HasTaskCost

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasTaskCost() bool`

HasTaskCost returns a boolean if a field has been set.

### GetTaskName

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskType

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskType() int32`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetTaskTypeOk() (*int32, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetTaskType(v int32)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetUploadFileUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetUploadFileUrl() string`

GetUploadFileUrl returns the UploadFileUrl field if non-nil, zero value otherwise.

### GetUploadFileUrlOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetUploadFileUrlOk() (*string, bool)`

GetUploadFileUrlOk returns a tuple with the UploadFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFileUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetUploadFileUrl(v string)`

SetUploadFileUrl sets UploadFileUrl field to given value.

### HasUploadFileUrl

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasUploadFileUrl() bool`

HasUploadFileUrl returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanMerchantUserBatchTask) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanMerchantUserBatchTask) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanMerchantUserBatchTask) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


