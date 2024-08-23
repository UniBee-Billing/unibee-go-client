/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantAuthRegisterVerifyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthRegisterVerifyRes{}

// UnibeeApiMerchantAuthRegisterVerifyRes struct for UnibeeApiMerchantAuthRegisterVerifyRes
type UnibeeApiMerchantAuthRegisterVerifyRes struct {
	MerchantMember *UnibeeApiBeanDetailMerchantMemberDetail `json:"merchantMember,omitempty"`
}

// NewUnibeeApiMerchantAuthRegisterVerifyRes instantiates a new UnibeeApiMerchantAuthRegisterVerifyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthRegisterVerifyRes() *UnibeeApiMerchantAuthRegisterVerifyRes {
	this := UnibeeApiMerchantAuthRegisterVerifyRes{}
	return &this
}

// NewUnibeeApiMerchantAuthRegisterVerifyResWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterVerifyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthRegisterVerifyResWithDefaults() *UnibeeApiMerchantAuthRegisterVerifyRes {
	this := UnibeeApiMerchantAuthRegisterVerifyRes{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthRegisterVerifyRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterVerifyRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthRegisterVerifyRes) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the MerchantMember field.
func (o *UnibeeApiMerchantAuthRegisterVerifyRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.MerchantMember = &v
}

func (o UnibeeApiMerchantAuthRegisterVerifyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthRegisterVerifyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantAuthRegisterVerifyRes struct {
	value *UnibeeApiMerchantAuthRegisterVerifyRes
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyRes) Get() *UnibeeApiMerchantAuthRegisterVerifyRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyRes) Set(val *UnibeeApiMerchantAuthRegisterVerifyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthRegisterVerifyRes(val *UnibeeApiMerchantAuthRegisterVerifyRes) *NullableUnibeeApiMerchantAuthRegisterVerifyRes {
	return &NullableUnibeeApiMerchantAuthRegisterVerifyRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


