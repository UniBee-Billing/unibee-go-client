/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserUpdateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserUpdateRes{}

// UnibeeApiMerchantUserUpdateRes struct for UnibeeApiMerchantUserUpdateRes
type UnibeeApiMerchantUserUpdateRes struct {
	User *UnibeeApiBeanDetailUserAccountDetail `json:"user,omitempty"`
}

// NewUnibeeApiMerchantUserUpdateRes instantiates a new UnibeeApiMerchantUserUpdateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserUpdateRes() *UnibeeApiMerchantUserUpdateRes {
	this := UnibeeApiMerchantUserUpdateRes{}
	return &this
}

// NewUnibeeApiMerchantUserUpdateResWithDefaults instantiates a new UnibeeApiMerchantUserUpdateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserUpdateResWithDefaults() *UnibeeApiMerchantUserUpdateRes {
	this := UnibeeApiMerchantUserUpdateRes{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateRes) GetUser() UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateRes) GetUserOk() (*UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateRes) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanDetailUserAccountDetail and assigns it to the User field.
func (o *UnibeeApiMerchantUserUpdateRes) SetUser(v UnibeeApiBeanDetailUserAccountDetail) {
	o.User = &v
}

func (o UnibeeApiMerchantUserUpdateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserUpdateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserUpdateRes struct {
	value *UnibeeApiMerchantUserUpdateRes
	isSet bool
}

func (v NullableUnibeeApiMerchantUserUpdateRes) Get() *UnibeeApiMerchantUserUpdateRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserUpdateRes) Set(val *UnibeeApiMerchantUserUpdateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserUpdateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserUpdateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserUpdateRes(val *UnibeeApiMerchantUserUpdateRes) *NullableUnibeeApiMerchantUserUpdateRes {
	return &NullableUnibeeApiMerchantUserUpdateRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserUpdateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserUpdateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


