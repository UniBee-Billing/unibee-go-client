/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantTaskEditTemplateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskEditTemplateReq{}

// UnibeeApiMerchantTaskEditTemplateReq struct for UnibeeApiMerchantTaskEditTemplateReq
type UnibeeApiMerchantTaskEditTemplateReq struct {
	// ExportColumns, the export file column list, will export all columns if not specified
	ExportColumns [][]string `json:"exportColumns,omitempty"`
	// The format of export file, xlsx|csv, will be xlsx if not specified
	Format *string `json:"format,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// Payload
	Payload *map[string]interface{} `json:"payload,omitempty"`
	// Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport
	Task *string `json:"task,omitempty"`
	// templateId
	TemplateId int64 `json:"templateId"`
}

type _UnibeeApiMerchantTaskEditTemplateReq UnibeeApiMerchantTaskEditTemplateReq

// NewUnibeeApiMerchantTaskEditTemplateReq instantiates a new UnibeeApiMerchantTaskEditTemplateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskEditTemplateReq(templateId int64) *UnibeeApiMerchantTaskEditTemplateReq {
	this := UnibeeApiMerchantTaskEditTemplateReq{}
	this.TemplateId = templateId
	return &this
}

// NewUnibeeApiMerchantTaskEditTemplateReqWithDefaults instantiates a new UnibeeApiMerchantTaskEditTemplateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskEditTemplateReqWithDefaults() *UnibeeApiMerchantTaskEditTemplateReq {
	this := UnibeeApiMerchantTaskEditTemplateReq{}
	return &this
}

// GetExportColumns returns the ExportColumns field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetExportColumns() [][]string {
	if o == nil || IsNil(o.ExportColumns) {
		var ret [][]string
		return ret
	}
	return o.ExportColumns
}

// GetExportColumnsOk returns a tuple with the ExportColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetExportColumnsOk() ([][]string, bool) {
	if o == nil || IsNil(o.ExportColumns) {
		return nil, false
	}
	return o.ExportColumns, true
}

// HasExportColumns returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) HasExportColumns() bool {
	if o != nil && !IsNil(o.ExportColumns) {
		return true
	}

	return false
}

// SetExportColumns gets a reference to the given [][]string and assigns it to the ExportColumns field.
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetExportColumns(v [][]string) {
	o.ExportColumns = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetName(v string) {
	o.Name = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetPayload() map[string]interface{} {
	if o == nil || IsNil(o.Payload) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetPayload(v map[string]interface{}) {
	o.Payload = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetTask(v string) {
	o.Task = &v
}

// GetTemplateId returns the TemplateId field value
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTemplateId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTemplateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *UnibeeApiMerchantTaskEditTemplateReq) SetTemplateId(v int64) {
	o.TemplateId = v
}

func (o UnibeeApiMerchantTaskEditTemplateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskEditTemplateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportColumns) {
		toSerialize["exportColumns"] = o.ExportColumns
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
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
	toSerialize["templateId"] = o.TemplateId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantTaskEditTemplateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantTaskEditTemplateReq := _UnibeeApiMerchantTaskEditTemplateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantTaskEditTemplateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantTaskEditTemplateReq(varUnibeeApiMerchantTaskEditTemplateReq)

	return err
}

type NullableUnibeeApiMerchantTaskEditTemplateReq struct {
	value *UnibeeApiMerchantTaskEditTemplateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskEditTemplateReq) Get() *UnibeeApiMerchantTaskEditTemplateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskEditTemplateReq) Set(val *UnibeeApiMerchantTaskEditTemplateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskEditTemplateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskEditTemplateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskEditTemplateReq(val *UnibeeApiMerchantTaskEditTemplateReq) *NullableUnibeeApiMerchantTaskEditTemplateReq {
	return &NullableUnibeeApiMerchantTaskEditTemplateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskEditTemplateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskEditTemplateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


