/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantTaskExportColumnListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskExportColumnListReq{}

// UnibeeApiMerchantTaskExportColumnListReq struct for UnibeeApiMerchantTaskExportColumnListReq
type UnibeeApiMerchantTaskExportColumnListReq struct {
	// Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport
	Task *string `json:"task,omitempty"`
}

// NewUnibeeApiMerchantTaskExportColumnListReq instantiates a new UnibeeApiMerchantTaskExportColumnListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskExportColumnListReq() *UnibeeApiMerchantTaskExportColumnListReq {
	this := UnibeeApiMerchantTaskExportColumnListReq{}
	return &this
}

// NewUnibeeApiMerchantTaskExportColumnListReqWithDefaults instantiates a new UnibeeApiMerchantTaskExportColumnListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskExportColumnListReqWithDefaults() *UnibeeApiMerchantTaskExportColumnListReq {
	this := UnibeeApiMerchantTaskExportColumnListReq{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportColumnListReq) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportColumnListReq) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportColumnListReq) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *UnibeeApiMerchantTaskExportColumnListReq) SetTask(v string) {
	o.Task = &v
}

func (o UnibeeApiMerchantTaskExportColumnListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskExportColumnListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskExportColumnListReq struct {
	value *UnibeeApiMerchantTaskExportColumnListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskExportColumnListReq) Get() *UnibeeApiMerchantTaskExportColumnListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListReq) Set(val *UnibeeApiMerchantTaskExportColumnListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskExportColumnListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskExportColumnListReq(val *UnibeeApiMerchantTaskExportColumnListReq) *NullableUnibeeApiMerchantTaskExportColumnListReq {
	return &NullableUnibeeApiMerchantTaskExportColumnListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskExportColumnListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


