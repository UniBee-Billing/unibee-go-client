/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserGetRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserGetRes{}

// UnibeeApiMerchantUserGetRes struct for UnibeeApiMerchantUserGetRes
type UnibeeApiMerchantUserGetRes struct {
	User *UnibeeApiBeanDetailUserAccountDetail `json:"user,omitempty"`
}

// NewUnibeeApiMerchantUserGetRes instantiates a new UnibeeApiMerchantUserGetRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserGetRes() *UnibeeApiMerchantUserGetRes {
	this := UnibeeApiMerchantUserGetRes{}
	return &this
}

// NewUnibeeApiMerchantUserGetResWithDefaults instantiates a new UnibeeApiMerchantUserGetRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserGetResWithDefaults() *UnibeeApiMerchantUserGetRes {
	this := UnibeeApiMerchantUserGetRes{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserGetRes) GetUser() UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserGetRes) GetUserOk() (*UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserGetRes) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanDetailUserAccountDetail and assigns it to the User field.
func (o *UnibeeApiMerchantUserGetRes) SetUser(v UnibeeApiBeanDetailUserAccountDetail) {
	o.User = &v
}

func (o UnibeeApiMerchantUserGetRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserGetRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserGetRes struct {
	value *UnibeeApiMerchantUserGetRes
	isSet bool
}

func (v NullableUnibeeApiMerchantUserGetRes) Get() *UnibeeApiMerchantUserGetRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserGetRes) Set(val *UnibeeApiMerchantUserGetRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserGetRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserGetRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserGetRes(val *UnibeeApiMerchantUserGetRes) *NullableUnibeeApiMerchantUserGetRes {
	return &NullableUnibeeApiMerchantUserGetRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserGetRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserGetRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


