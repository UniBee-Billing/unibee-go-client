/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantCreditPromoCreditDecrementPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditPromoCreditDecrementPost200ResponseData{}

// MerchantCreditPromoCreditDecrementPost200ResponseData struct for MerchantCreditPromoCreditDecrementPost200ResponseData
type MerchantCreditPromoCreditDecrementPost200ResponseData struct {
	UserPromoCreditAccount *UnibeeApiBeanCreditAccount `json:"UserPromoCreditAccount,omitempty"`
}

// NewMerchantCreditPromoCreditDecrementPost200ResponseData instantiates a new MerchantCreditPromoCreditDecrementPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditPromoCreditDecrementPost200ResponseData() *MerchantCreditPromoCreditDecrementPost200ResponseData {
	this := MerchantCreditPromoCreditDecrementPost200ResponseData{}
	return &this
}

// NewMerchantCreditPromoCreditDecrementPost200ResponseDataWithDefaults instantiates a new MerchantCreditPromoCreditDecrementPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditPromoCreditDecrementPost200ResponseDataWithDefaults() *MerchantCreditPromoCreditDecrementPost200ResponseData {
	this := MerchantCreditPromoCreditDecrementPost200ResponseData{}
	return &this
}

// GetUserPromoCreditAccount returns the UserPromoCreditAccount field value if set, zero value otherwise.
func (o *MerchantCreditPromoCreditDecrementPost200ResponseData) GetUserPromoCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.UserPromoCreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.UserPromoCreditAccount
}

// GetUserPromoCreditAccountOk returns a tuple with the UserPromoCreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditPromoCreditDecrementPost200ResponseData) GetUserPromoCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.UserPromoCreditAccount) {
		return nil, false
	}
	return o.UserPromoCreditAccount, true
}

// HasUserPromoCreditAccount returns a boolean if a field has been set.
func (o *MerchantCreditPromoCreditDecrementPost200ResponseData) HasUserPromoCreditAccount() bool {
	if o != nil && !IsNil(o.UserPromoCreditAccount) {
		return true
	}

	return false
}

// SetUserPromoCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the UserPromoCreditAccount field.
func (o *MerchantCreditPromoCreditDecrementPost200ResponseData) SetUserPromoCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.UserPromoCreditAccount = &v
}

func (o MerchantCreditPromoCreditDecrementPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditPromoCreditDecrementPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserPromoCreditAccount) {
		toSerialize["UserPromoCreditAccount"] = o.UserPromoCreditAccount
	}
	return toSerialize, nil
}

type NullableMerchantCreditPromoCreditDecrementPost200ResponseData struct {
	value *MerchantCreditPromoCreditDecrementPost200ResponseData
	isSet bool
}

func (v NullableMerchantCreditPromoCreditDecrementPost200ResponseData) Get() *MerchantCreditPromoCreditDecrementPost200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditPromoCreditDecrementPost200ResponseData) Set(val *MerchantCreditPromoCreditDecrementPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditPromoCreditDecrementPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditPromoCreditDecrementPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditPromoCreditDecrementPost200ResponseData(val *MerchantCreditPromoCreditDecrementPost200ResponseData) *NullableMerchantCreditPromoCreditDecrementPost200ResponseData {
	return &NullableMerchantCreditPromoCreditDecrementPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditPromoCreditDecrementPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditPromoCreditDecrementPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


