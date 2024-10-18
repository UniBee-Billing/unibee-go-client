/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantVatNumberValidateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantVatNumberValidateRes{}

// UnibeeApiMerchantVatNumberValidateRes struct for UnibeeApiMerchantVatNumberValidateRes
type UnibeeApiMerchantVatNumberValidateRes struct {
	VatNumberValidate *UnibeeApiBeanValidResult `json:"vatNumberValidate,omitempty"`
}

// NewUnibeeApiMerchantVatNumberValidateRes instantiates a new UnibeeApiMerchantVatNumberValidateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantVatNumberValidateRes() *UnibeeApiMerchantVatNumberValidateRes {
	this := UnibeeApiMerchantVatNumberValidateRes{}
	return &this
}

// NewUnibeeApiMerchantVatNumberValidateResWithDefaults instantiates a new UnibeeApiMerchantVatNumberValidateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantVatNumberValidateResWithDefaults() *UnibeeApiMerchantVatNumberValidateRes {
	this := UnibeeApiMerchantVatNumberValidateRes{}
	return &this
}

// GetVatNumberValidate returns the VatNumberValidate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantVatNumberValidateRes) GetVatNumberValidate() UnibeeApiBeanValidResult {
	if o == nil || IsNil(o.VatNumberValidate) {
		var ret UnibeeApiBeanValidResult
		return ret
	}
	return *o.VatNumberValidate
}

// GetVatNumberValidateOk returns a tuple with the VatNumberValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatNumberValidateRes) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool) {
	if o == nil || IsNil(o.VatNumberValidate) {
		return nil, false
	}
	return o.VatNumberValidate, true
}

// HasVatNumberValidate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantVatNumberValidateRes) HasVatNumberValidate() bool {
	if o != nil && !IsNil(o.VatNumberValidate) {
		return true
	}

	return false
}

// SetVatNumberValidate gets a reference to the given UnibeeApiBeanValidResult and assigns it to the VatNumberValidate field.
func (o *UnibeeApiMerchantVatNumberValidateRes) SetVatNumberValidate(v UnibeeApiBeanValidResult) {
	o.VatNumberValidate = &v
}

func (o UnibeeApiMerchantVatNumberValidateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantVatNumberValidateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VatNumberValidate) {
		toSerialize["vatNumberValidate"] = o.VatNumberValidate
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantVatNumberValidateRes struct {
	value *UnibeeApiMerchantVatNumberValidateRes
	isSet bool
}

func (v NullableUnibeeApiMerchantVatNumberValidateRes) Get() *UnibeeApiMerchantVatNumberValidateRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantVatNumberValidateRes) Set(val *UnibeeApiMerchantVatNumberValidateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantVatNumberValidateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantVatNumberValidateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantVatNumberValidateRes(val *UnibeeApiMerchantVatNumberValidateRes) *NullableUnibeeApiMerchantVatNumberValidateRes {
	return &NullableUnibeeApiMerchantVatNumberValidateRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantVatNumberValidateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantVatNumberValidateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


