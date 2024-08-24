/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantUserNewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUserNewPost200ResponseData{}

// MerchantUserNewPost200ResponseData struct for MerchantUserNewPost200ResponseData
type MerchantUserNewPost200ResponseData struct {
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewMerchantUserNewPost200ResponseData instantiates a new MerchantUserNewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUserNewPost200ResponseData() *MerchantUserNewPost200ResponseData {
	this := MerchantUserNewPost200ResponseData{}
	return &this
}

// NewMerchantUserNewPost200ResponseDataWithDefaults instantiates a new MerchantUserNewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUserNewPost200ResponseDataWithDefaults() *MerchantUserNewPost200ResponseData {
	this := MerchantUserNewPost200ResponseData{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantUserNewPost200ResponseData) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserNewPost200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantUserNewPost200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *MerchantUserNewPost200ResponseData) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o MerchantUserNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUserNewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableMerchantUserNewPost200ResponseData struct {
	value *MerchantUserNewPost200ResponseData
	isSet bool
}

func (v NullableMerchantUserNewPost200ResponseData) Get() *MerchantUserNewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantUserNewPost200ResponseData) Set(val *MerchantUserNewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUserNewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUserNewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUserNewPost200ResponseData(val *MerchantUserNewPost200ResponseData) *NullableMerchantUserNewPost200ResponseData {
	return &NullableMerchantUserNewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantUserNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUserNewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


