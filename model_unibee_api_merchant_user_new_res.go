/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserNewRes{}

// UnibeeApiMerchantUserNewRes struct for UnibeeApiMerchantUserNewRes
type UnibeeApiMerchantUserNewRes struct {
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiMerchantUserNewRes instantiates a new UnibeeApiMerchantUserNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserNewRes() *UnibeeApiMerchantUserNewRes {
	this := UnibeeApiMerchantUserNewRes{}
	return &this
}

// NewUnibeeApiMerchantUserNewResWithDefaults instantiates a new UnibeeApiMerchantUserNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserNewResWithDefaults() *UnibeeApiMerchantUserNewRes {
	this := UnibeeApiMerchantUserNewRes{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserNewRes) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserNewRes) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserNewRes) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiMerchantUserNewRes) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiMerchantUserNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserNewRes struct {
	value *UnibeeApiMerchantUserNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantUserNewRes) Get() *UnibeeApiMerchantUserNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserNewRes) Set(val *UnibeeApiMerchantUserNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserNewRes(val *UnibeeApiMerchantUserNewRes) *NullableUnibeeApiMerchantUserNewRes {
	return &NullableUnibeeApiMerchantUserNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


