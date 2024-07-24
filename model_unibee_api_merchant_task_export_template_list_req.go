/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantTaskExportTemplateListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskExportTemplateListReq{}

// UnibeeApiMerchantTaskExportTemplateListReq struct for UnibeeApiMerchantTaskExportTemplateListReq
type UnibeeApiMerchantTaskExportTemplateListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
	// Filter Task, Optional, InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport
	Task *string `json:"task,omitempty"`
}

// NewUnibeeApiMerchantTaskExportTemplateListReq instantiates a new UnibeeApiMerchantTaskExportTemplateListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskExportTemplateListReq() *UnibeeApiMerchantTaskExportTemplateListReq {
	this := UnibeeApiMerchantTaskExportTemplateListReq{}
	return &this
}

// NewUnibeeApiMerchantTaskExportTemplateListReqWithDefaults instantiates a new UnibeeApiMerchantTaskExportTemplateListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskExportTemplateListReqWithDefaults() *UnibeeApiMerchantTaskExportTemplateListReq {
	this := UnibeeApiMerchantTaskExportTemplateListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetPage(v int32) {
	o.Page = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetTask(v string) {
	o.Task = &v
}

func (o UnibeeApiMerchantTaskExportTemplateListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskExportTemplateListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskExportTemplateListReq struct {
	value *UnibeeApiMerchantTaskExportTemplateListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListReq) Get() *UnibeeApiMerchantTaskExportTemplateListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListReq) Set(val *UnibeeApiMerchantTaskExportTemplateListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskExportTemplateListReq(val *UnibeeApiMerchantTaskExportTemplateListReq) *NullableUnibeeApiMerchantTaskExportTemplateListReq {
	return &NullableUnibeeApiMerchantTaskExportTemplateListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


