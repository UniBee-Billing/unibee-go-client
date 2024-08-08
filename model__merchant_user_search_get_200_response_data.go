/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantUserSearchGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUserSearchGet200ResponseData{}

// MerchantUserSearchGet200ResponseData struct for MerchantUserSearchGet200ResponseData
type MerchantUserSearchGet200ResponseData struct {
	// UserAccounts
	UserAccounts []UnibeeApiBeanDetailUserAccountDetail `json:"userAccounts,omitempty"`
}

// NewMerchantUserSearchGet200ResponseData instantiates a new MerchantUserSearchGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUserSearchGet200ResponseData() *MerchantUserSearchGet200ResponseData {
	this := MerchantUserSearchGet200ResponseData{}
	return &this
}

// NewMerchantUserSearchGet200ResponseDataWithDefaults instantiates a new MerchantUserSearchGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUserSearchGet200ResponseDataWithDefaults() *MerchantUserSearchGet200ResponseData {
	this := MerchantUserSearchGet200ResponseData{}
	return &this
}

// GetUserAccounts returns the UserAccounts field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200ResponseData) GetUserAccounts() []UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.UserAccounts) {
		var ret []UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return o.UserAccounts
}

// GetUserAccountsOk returns a tuple with the UserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200ResponseData) GetUserAccountsOk() ([]UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.UserAccounts) {
		return nil, false
	}
	return o.UserAccounts, true
}

// HasUserAccounts returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200ResponseData) HasUserAccounts() bool {
	if o != nil && !IsNil(o.UserAccounts) {
		return true
	}

	return false
}

// SetUserAccounts gets a reference to the given []UnibeeApiBeanDetailUserAccountDetail and assigns it to the UserAccounts field.
func (o *MerchantUserSearchGet200ResponseData) SetUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail) {
	o.UserAccounts = v
}

func (o MerchantUserSearchGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUserSearchGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAccounts) {
		toSerialize["userAccounts"] = o.UserAccounts
	}
	return toSerialize, nil
}

type NullableMerchantUserSearchGet200ResponseData struct {
	value *MerchantUserSearchGet200ResponseData
	isSet bool
}

func (v NullableMerchantUserSearchGet200ResponseData) Get() *MerchantUserSearchGet200ResponseData {
	return v.value
}

func (v *NullableMerchantUserSearchGet200ResponseData) Set(val *MerchantUserSearchGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUserSearchGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUserSearchGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUserSearchGet200ResponseData(val *MerchantUserSearchGet200ResponseData) *NullableMerchantUserSearchGet200ResponseData {
	return &NullableMerchantUserSearchGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantUserSearchGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUserSearchGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


