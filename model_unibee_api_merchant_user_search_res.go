/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserSearchRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserSearchRes{}

// UnibeeApiMerchantUserSearchRes struct for UnibeeApiMerchantUserSearchRes
type UnibeeApiMerchantUserSearchRes struct {
	// UserAccounts
	UserAccounts []UnibeeApiBeanDetailUserAccountDetail `json:"userAccounts,omitempty"`
}

// NewUnibeeApiMerchantUserSearchRes instantiates a new UnibeeApiMerchantUserSearchRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserSearchRes() *UnibeeApiMerchantUserSearchRes {
	this := UnibeeApiMerchantUserSearchRes{}
	return &this
}

// NewUnibeeApiMerchantUserSearchResWithDefaults instantiates a new UnibeeApiMerchantUserSearchRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserSearchResWithDefaults() *UnibeeApiMerchantUserSearchRes {
	this := UnibeeApiMerchantUserSearchRes{}
	return &this
}

// GetUserAccounts returns the UserAccounts field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserSearchRes) GetUserAccounts() []UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.UserAccounts) {
		var ret []UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return o.UserAccounts
}

// GetUserAccountsOk returns a tuple with the UserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserSearchRes) GetUserAccountsOk() ([]UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.UserAccounts) {
		return nil, false
	}
	return o.UserAccounts, true
}

// HasUserAccounts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserSearchRes) HasUserAccounts() bool {
	if o != nil && !IsNil(o.UserAccounts) {
		return true
	}

	return false
}

// SetUserAccounts gets a reference to the given []UnibeeApiBeanDetailUserAccountDetail and assigns it to the UserAccounts field.
func (o *UnibeeApiMerchantUserSearchRes) SetUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail) {
	o.UserAccounts = v
}

func (o UnibeeApiMerchantUserSearchRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserSearchRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAccounts) {
		toSerialize["userAccounts"] = o.UserAccounts
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserSearchRes struct {
	value *UnibeeApiMerchantUserSearchRes
	isSet bool
}

func (v NullableUnibeeApiMerchantUserSearchRes) Get() *UnibeeApiMerchantUserSearchRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserSearchRes) Set(val *UnibeeApiMerchantUserSearchRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserSearchRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserSearchRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserSearchRes(val *UnibeeApiMerchantUserSearchRes) *NullableUnibeeApiMerchantUserSearchRes {
	return &NullableUnibeeApiMerchantUserSearchRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserSearchRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserSearchRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


