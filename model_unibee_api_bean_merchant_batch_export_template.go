/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantBatchExportTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantBatchExportTemplate{}

// UnibeeApiBeanMerchantBatchExportTemplate struct for UnibeeApiBeanMerchantBatchExportTemplate
type UnibeeApiBeanMerchantBatchExportTemplate struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// ExportColumns, the export file column list, will export all columns if not specified
	ExportColumns []string `json:"exportColumns,omitempty"`
	// The format of export file, xlsx|csv, will be xlsx if not specified
	Format *string `json:"format,omitempty"`
	// member_id
	MemberId *int64 `json:"memberId,omitempty"`
	// merchant_id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// Payload
	Payload *map[string]interface{} `json:"payload,omitempty"`
	// Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport
	Task *string `json:"task,omitempty"`
	// templateId
	TemplateId *int64 `json:"templateId,omitempty"`
}

// NewUnibeeApiBeanMerchantBatchExportTemplate instantiates a new UnibeeApiBeanMerchantBatchExportTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantBatchExportTemplate() *UnibeeApiBeanMerchantBatchExportTemplate {
	this := UnibeeApiBeanMerchantBatchExportTemplate{}
	return &this
}

// NewUnibeeApiBeanMerchantBatchExportTemplateWithDefaults instantiates a new UnibeeApiBeanMerchantBatchExportTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantBatchExportTemplateWithDefaults() *UnibeeApiBeanMerchantBatchExportTemplate {
	this := UnibeeApiBeanMerchantBatchExportTemplate{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetExportColumns returns the ExportColumns field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetExportColumns() []string {
	if o == nil || IsNil(o.ExportColumns) {
		var ret []string
		return ret
	}
	return o.ExportColumns
}

// GetExportColumnsOk returns a tuple with the ExportColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetExportColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExportColumns) {
		return nil, false
	}
	return o.ExportColumns, true
}

// HasExportColumns returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasExportColumns() bool {
	if o != nil && !IsNil(o.ExportColumns) {
		return true
	}

	return false
}

// SetExportColumns gets a reference to the given []string and assigns it to the ExportColumns field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetExportColumns(v []string) {
	o.ExportColumns = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetFormat(v string) {
	o.Format = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetName(v string) {
	o.Name = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetPayload() map[string]interface{} {
	if o == nil || IsNil(o.Payload) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetPayload(v map[string]interface{}) {
	o.Payload = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetTask(v string) {
	o.Task = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTemplateId() int64 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int64
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTemplateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int64 and assigns it to the TemplateId field.
func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetTemplateId(v int64) {
	o.TemplateId = &v
}

func (o UnibeeApiBeanMerchantBatchExportTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantBatchExportTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ExportColumns) {
		toSerialize["exportColumns"] = o.ExportColumns
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantBatchExportTemplate struct {
	value *UnibeeApiBeanMerchantBatchExportTemplate
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantBatchExportTemplate) Get() *UnibeeApiBeanMerchantBatchExportTemplate {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantBatchExportTemplate) Set(val *UnibeeApiBeanMerchantBatchExportTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantBatchExportTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantBatchExportTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantBatchExportTemplate(val *UnibeeApiBeanMerchantBatchExportTemplate) *NullableUnibeeApiBeanMerchantBatchExportTemplate {
	return &NullableUnibeeApiBeanMerchantBatchExportTemplate{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantBatchExportTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantBatchExportTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


