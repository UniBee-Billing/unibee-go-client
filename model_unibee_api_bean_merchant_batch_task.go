/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantBatchTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantBatchTask{}

// UnibeeApiBeanMerchantBatchTask struct for UnibeeApiBeanMerchantBatchTask
type UnibeeApiBeanMerchantBatchTask struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// download_file_url
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// reason of failure
	FailureReason *string `json:"failureReason,omitempty"`
	// task_finish_time
	FinishTime *int64 `json:"finishTime,omitempty"`
	// format of file
	Format *string `json:"format,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// last update utc time
	LastUpdateTime *int64 `json:"lastUpdateTime,omitempty"`
	// member_id
	MemberId *int64 `json:"memberId,omitempty"`
	// merchant_id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// payload(json)
	Payload *string `json:"payload,omitempty"`
	// task_start_time
	StartTime *int64 `json:"startTime,omitempty"`
	// Status。0-Pending，1-Processing，2-Success，3-Failure
	Status *int32 `json:"status,omitempty"`
	// success_count
	SuccessCount *int64 `json:"successCount,omitempty"`
	// task cost time(second)
	TaskCost *int32 `json:"taskCost,omitempty"`
	// task_name
	TaskName *string `json:"taskName,omitempty"`
	// type，0-download，1-upload
	TaskType *int32 `json:"taskType,omitempty"`
	// the file url of upload type task
	UploadFileUrl *string `json:"uploadFileUrl,omitempty"`
}

// NewUnibeeApiBeanMerchantBatchTask instantiates a new UnibeeApiBeanMerchantBatchTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantBatchTask() *UnibeeApiBeanMerchantBatchTask {
	this := UnibeeApiBeanMerchantBatchTask{}
	return &this
}

// NewUnibeeApiBeanMerchantBatchTaskWithDefaults instantiates a new UnibeeApiBeanMerchantBatchTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantBatchTaskWithDefaults() *UnibeeApiBeanMerchantBatchTask {
	this := UnibeeApiBeanMerchantBatchTask{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantBatchTask) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *UnibeeApiBeanMerchantBatchTask) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *UnibeeApiBeanMerchantBatchTask) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetFinishTime() int64 {
	if o == nil || IsNil(o.FinishTime) {
		var ret int64
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetFinishTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.FinishTime) {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasFinishTime() bool {
	if o != nil && !IsNil(o.FinishTime) {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given int64 and assigns it to the FinishTime field.
func (o *UnibeeApiBeanMerchantBatchTask) SetFinishTime(v int64) {
	o.FinishTime = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UnibeeApiBeanMerchantBatchTask) SetFormat(v string) {
	o.Format = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantBatchTask) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetLastUpdateTime() int64 {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret int64
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetLastUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given int64 and assigns it to the LastUpdateTime field.
func (o *UnibeeApiBeanMerchantBatchTask) SetLastUpdateTime(v int64) {
	o.LastUpdateTime = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiBeanMerchantBatchTask) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantBatchTask) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *UnibeeApiBeanMerchantBatchTask) SetPayload(v string) {
	o.Payload = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *UnibeeApiBeanMerchantBatchTask) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanMerchantBatchTask) SetStatus(v int32) {
	o.Status = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetSuccessCount() int64 {
	if o == nil || IsNil(o.SuccessCount) {
		var ret int64
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetSuccessCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SuccessCount) {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasSuccessCount() bool {
	if o != nil && !IsNil(o.SuccessCount) {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int64 and assigns it to the SuccessCount field.
func (o *UnibeeApiBeanMerchantBatchTask) SetSuccessCount(v int64) {
	o.SuccessCount = &v
}

// GetTaskCost returns the TaskCost field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskCost() int32 {
	if o == nil || IsNil(o.TaskCost) {
		var ret int32
		return ret
	}
	return *o.TaskCost
}

// GetTaskCostOk returns a tuple with the TaskCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskCostOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskCost) {
		return nil, false
	}
	return o.TaskCost, true
}

// HasTaskCost returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasTaskCost() bool {
	if o != nil && !IsNil(o.TaskCost) {
		return true
	}

	return false
}

// SetTaskCost gets a reference to the given int32 and assigns it to the TaskCost field.
func (o *UnibeeApiBeanMerchantBatchTask) SetTaskCost(v int32) {
	o.TaskCost = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskName() string {
	if o == nil || IsNil(o.TaskName) {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaskName) {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasTaskName() bool {
	if o != nil && !IsNil(o.TaskName) {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *UnibeeApiBeanMerchantBatchTask) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskType() int32 {
	if o == nil || IsNil(o.TaskType) {
		var ret int32
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetTaskTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given int32 and assigns it to the TaskType field.
func (o *UnibeeApiBeanMerchantBatchTask) SetTaskType(v int32) {
	o.TaskType = &v
}

// GetUploadFileUrl returns the UploadFileUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchTask) GetUploadFileUrl() string {
	if o == nil || IsNil(o.UploadFileUrl) {
		var ret string
		return ret
	}
	return *o.UploadFileUrl
}

// GetUploadFileUrlOk returns a tuple with the UploadFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchTask) GetUploadFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UploadFileUrl) {
		return nil, false
	}
	return o.UploadFileUrl, true
}

// HasUploadFileUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchTask) HasUploadFileUrl() bool {
	if o != nil && !IsNil(o.UploadFileUrl) {
		return true
	}

	return false
}

// SetUploadFileUrl gets a reference to the given string and assigns it to the UploadFileUrl field.
func (o *UnibeeApiBeanMerchantBatchTask) SetUploadFileUrl(v string) {
	o.UploadFileUrl = &v
}

func (o UnibeeApiBeanMerchantBatchTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantBatchTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.FinishTime) {
		toSerialize["finishTime"] = o.FinishTime
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SuccessCount) {
		toSerialize["successCount"] = o.SuccessCount
	}
	if !IsNil(o.TaskCost) {
		toSerialize["taskCost"] = o.TaskCost
	}
	if !IsNil(o.TaskName) {
		toSerialize["taskName"] = o.TaskName
	}
	if !IsNil(o.TaskType) {
		toSerialize["taskType"] = o.TaskType
	}
	if !IsNil(o.UploadFileUrl) {
		toSerialize["uploadFileUrl"] = o.UploadFileUrl
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantBatchTask struct {
	value *UnibeeApiBeanMerchantBatchTask
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantBatchTask) Get() *UnibeeApiBeanMerchantBatchTask {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantBatchTask) Set(val *UnibeeApiBeanMerchantBatchTask) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantBatchTask) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantBatchTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantBatchTask(val *UnibeeApiBeanMerchantBatchTask) *NullableUnibeeApiBeanMerchantBatchTask {
	return &NullableUnibeeApiBeanMerchantBatchTask{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantBatchTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantBatchTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


