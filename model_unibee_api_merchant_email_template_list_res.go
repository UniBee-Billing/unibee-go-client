/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantEmailTemplateListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailTemplateListRes{}

// UnibeeApiMerchantEmailTemplateListRes struct for UnibeeApiMerchantEmailTemplateListRes
type UnibeeApiMerchantEmailTemplateListRes struct {
	// Email Template Object List
	EmailTemplateList []UnibeeApiBeanMerchantEmailTemplate `json:"emailTemplateList,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantEmailTemplateListRes instantiates a new UnibeeApiMerchantEmailTemplateListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailTemplateListRes() *UnibeeApiMerchantEmailTemplateListRes {
	this := UnibeeApiMerchantEmailTemplateListRes{}
	return &this
}

// NewUnibeeApiMerchantEmailTemplateListResWithDefaults instantiates a new UnibeeApiMerchantEmailTemplateListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailTemplateListResWithDefaults() *UnibeeApiMerchantEmailTemplateListRes {
	this := UnibeeApiMerchantEmailTemplateListRes{}
	return &this
}

// GetEmailTemplateList returns the EmailTemplateList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailTemplateListRes) GetEmailTemplateList() []UnibeeApiBeanMerchantEmailTemplate {
	if o == nil || IsNil(o.EmailTemplateList) {
		var ret []UnibeeApiBeanMerchantEmailTemplate
		return ret
	}
	return o.EmailTemplateList
}

// GetEmailTemplateListOk returns a tuple with the EmailTemplateList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateListRes) GetEmailTemplateListOk() ([]UnibeeApiBeanMerchantEmailTemplate, bool) {
	if o == nil || IsNil(o.EmailTemplateList) {
		return nil, false
	}
	return o.EmailTemplateList, true
}

// HasEmailTemplateList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailTemplateListRes) HasEmailTemplateList() bool {
	if o != nil && !IsNil(o.EmailTemplateList) {
		return true
	}

	return false
}

// SetEmailTemplateList gets a reference to the given []UnibeeApiBeanMerchantEmailTemplate and assigns it to the EmailTemplateList field.
func (o *UnibeeApiMerchantEmailTemplateListRes) SetEmailTemplateList(v []UnibeeApiBeanMerchantEmailTemplate) {
	o.EmailTemplateList = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailTemplateListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailTemplateListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantEmailTemplateListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantEmailTemplateListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailTemplateListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailTemplateList) {
		toSerialize["emailTemplateList"] = o.EmailTemplateList
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantEmailTemplateListRes struct {
	value *UnibeeApiMerchantEmailTemplateListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailTemplateListRes) Get() *UnibeeApiMerchantEmailTemplateListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailTemplateListRes) Set(val *UnibeeApiMerchantEmailTemplateListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailTemplateListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailTemplateListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailTemplateListRes(val *UnibeeApiMerchantEmailTemplateListRes) *NullableUnibeeApiMerchantEmailTemplateListRes {
	return &NullableUnibeeApiMerchantEmailTemplateListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailTemplateListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailTemplateListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


