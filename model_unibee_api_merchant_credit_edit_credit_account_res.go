/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantCreditEditCreditAccountRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditEditCreditAccountRes{}

// UnibeeApiMerchantCreditEditCreditAccountRes struct for UnibeeApiMerchantCreditEditCreditAccountRes
type UnibeeApiMerchantCreditEditCreditAccountRes struct {
	UserCreditAccount *UnibeeApiBeanCreditAccount `json:"UserCreditAccount,omitempty"`
}

// NewUnibeeApiMerchantCreditEditCreditAccountRes instantiates a new UnibeeApiMerchantCreditEditCreditAccountRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditEditCreditAccountRes() *UnibeeApiMerchantCreditEditCreditAccountRes {
	this := UnibeeApiMerchantCreditEditCreditAccountRes{}
	return &this
}

// NewUnibeeApiMerchantCreditEditCreditAccountResWithDefaults instantiates a new UnibeeApiMerchantCreditEditCreditAccountRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditEditCreditAccountResWithDefaults() *UnibeeApiMerchantCreditEditCreditAccountRes {
	this := UnibeeApiMerchantCreditEditCreditAccountRes{}
	return &this
}

// GetUserCreditAccount returns the UserCreditAccount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditCreditAccountRes) GetUserCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.UserCreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.UserCreditAccount
}

// GetUserCreditAccountOk returns a tuple with the UserCreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountRes) GetUserCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.UserCreditAccount) {
		return nil, false
	}
	return o.UserCreditAccount, true
}

// HasUserCreditAccount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountRes) HasUserCreditAccount() bool {
	if o != nil && !IsNil(o.UserCreditAccount) {
		return true
	}

	return false
}

// SetUserCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the UserCreditAccount field.
func (o *UnibeeApiMerchantCreditEditCreditAccountRes) SetUserCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.UserCreditAccount = &v
}

func (o UnibeeApiMerchantCreditEditCreditAccountRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditEditCreditAccountRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserCreditAccount) {
		toSerialize["UserCreditAccount"] = o.UserCreditAccount
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditEditCreditAccountRes struct {
	value *UnibeeApiMerchantCreditEditCreditAccountRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountRes) Get() *UnibeeApiMerchantCreditEditCreditAccountRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountRes) Set(val *UnibeeApiMerchantCreditEditCreditAccountRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditEditCreditAccountRes(val *UnibeeApiMerchantCreditEditCreditAccountRes) *NullableUnibeeApiMerchantCreditEditCreditAccountRes {
	return &NullableUnibeeApiMerchantCreditEditCreditAccountRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


