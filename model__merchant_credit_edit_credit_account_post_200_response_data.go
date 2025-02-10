/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantCreditEditCreditAccountPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditEditCreditAccountPost200ResponseData{}

// MerchantCreditEditCreditAccountPost200ResponseData struct for MerchantCreditEditCreditAccountPost200ResponseData
type MerchantCreditEditCreditAccountPost200ResponseData struct {
	UserCreditAccount *UnibeeApiBeanCreditAccount `json:"UserCreditAccount,omitempty"`
}

// NewMerchantCreditEditCreditAccountPost200ResponseData instantiates a new MerchantCreditEditCreditAccountPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditEditCreditAccountPost200ResponseData() *MerchantCreditEditCreditAccountPost200ResponseData {
	this := MerchantCreditEditCreditAccountPost200ResponseData{}
	return &this
}

// NewMerchantCreditEditCreditAccountPost200ResponseDataWithDefaults instantiates a new MerchantCreditEditCreditAccountPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditEditCreditAccountPost200ResponseDataWithDefaults() *MerchantCreditEditCreditAccountPost200ResponseData {
	this := MerchantCreditEditCreditAccountPost200ResponseData{}
	return &this
}

// GetUserCreditAccount returns the UserCreditAccount field value if set, zero value otherwise.
func (o *MerchantCreditEditCreditAccountPost200ResponseData) GetUserCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.UserCreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.UserCreditAccount
}

// GetUserCreditAccountOk returns a tuple with the UserCreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditEditCreditAccountPost200ResponseData) GetUserCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.UserCreditAccount) {
		return nil, false
	}
	return o.UserCreditAccount, true
}

// HasUserCreditAccount returns a boolean if a field has been set.
func (o *MerchantCreditEditCreditAccountPost200ResponseData) HasUserCreditAccount() bool {
	if o != nil && !IsNil(o.UserCreditAccount) {
		return true
	}

	return false
}

// SetUserCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the UserCreditAccount field.
func (o *MerchantCreditEditCreditAccountPost200ResponseData) SetUserCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.UserCreditAccount = &v
}

func (o MerchantCreditEditCreditAccountPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditEditCreditAccountPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserCreditAccount) {
		toSerialize["UserCreditAccount"] = o.UserCreditAccount
	}
	return toSerialize, nil
}

type NullableMerchantCreditEditCreditAccountPost200ResponseData struct {
	value *MerchantCreditEditCreditAccountPost200ResponseData
	isSet bool
}

func (v NullableMerchantCreditEditCreditAccountPost200ResponseData) Get() *MerchantCreditEditCreditAccountPost200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditEditCreditAccountPost200ResponseData) Set(val *MerchantCreditEditCreditAccountPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditEditCreditAccountPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditEditCreditAccountPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditEditCreditAccountPost200ResponseData(val *MerchantCreditEditCreditAccountPost200ResponseData) *NullableMerchantCreditEditCreditAccountPost200ResponseData {
	return &NullableMerchantCreditEditCreditAccountPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditEditCreditAccountPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditEditCreditAccountPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


