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

// checks if the UnibeeApiMerchantTaskExportTemplateListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskExportTemplateListRes{}

// UnibeeApiMerchantTaskExportTemplateListRes struct for UnibeeApiMerchantTaskExportTemplateListRes
type UnibeeApiMerchantTaskExportTemplateListRes struct {
	// Merchant Member Export Template List
	Templates []UnibeeApiBeanMerchantBatchExportTemplateSimplify `json:"templates,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantTaskExportTemplateListRes instantiates a new UnibeeApiMerchantTaskExportTemplateListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskExportTemplateListRes() *UnibeeApiMerchantTaskExportTemplateListRes {
	this := UnibeeApiMerchantTaskExportTemplateListRes{}
	return &this
}

// NewUnibeeApiMerchantTaskExportTemplateListResWithDefaults instantiates a new UnibeeApiMerchantTaskExportTemplateListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskExportTemplateListResWithDefaults() *UnibeeApiMerchantTaskExportTemplateListRes {
	this := UnibeeApiMerchantTaskExportTemplateListRes{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) GetTemplates() []UnibeeApiBeanMerchantBatchExportTemplateSimplify {
	if o == nil || IsNil(o.Templates) {
		var ret []UnibeeApiBeanMerchantBatchExportTemplateSimplify
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) GetTemplatesOk() ([]UnibeeApiBeanMerchantBatchExportTemplateSimplify, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []UnibeeApiBeanMerchantBatchExportTemplateSimplify and assigns it to the Templates field.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) SetTemplates(v []UnibeeApiBeanMerchantBatchExportTemplateSimplify) {
	o.Templates = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantTaskExportTemplateListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantTaskExportTemplateListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskExportTemplateListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskExportTemplateListRes struct {
	value *UnibeeApiMerchantTaskExportTemplateListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListRes) Get() *UnibeeApiMerchantTaskExportTemplateListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListRes) Set(val *UnibeeApiMerchantTaskExportTemplateListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskExportTemplateListRes(val *UnibeeApiMerchantTaskExportTemplateListRes) *NullableUnibeeApiMerchantTaskExportTemplateListRes {
	return &NullableUnibeeApiMerchantTaskExportTemplateListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskExportTemplateListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskExportTemplateListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


