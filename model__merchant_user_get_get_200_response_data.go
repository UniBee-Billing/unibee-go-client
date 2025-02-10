/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantUserGetGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUserGetGet200ResponseData{}

// MerchantUserGetGet200ResponseData struct for MerchantUserGetGet200ResponseData
type MerchantUserGetGet200ResponseData struct {
	User *UnibeeApiBeanDetailUserAccountDetail `json:"user,omitempty"`
}

// NewMerchantUserGetGet200ResponseData instantiates a new MerchantUserGetGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUserGetGet200ResponseData() *MerchantUserGetGet200ResponseData {
	this := MerchantUserGetGet200ResponseData{}
	return &this
}

// NewMerchantUserGetGet200ResponseDataWithDefaults instantiates a new MerchantUserGetGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUserGetGet200ResponseDataWithDefaults() *MerchantUserGetGet200ResponseData {
	this := MerchantUserGetGet200ResponseData{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantUserGetGet200ResponseData) GetUser() UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserGetGet200ResponseData) GetUserOk() (*UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantUserGetGet200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanDetailUserAccountDetail and assigns it to the User field.
func (o *MerchantUserGetGet200ResponseData) SetUser(v UnibeeApiBeanDetailUserAccountDetail) {
	o.User = &v
}

func (o MerchantUserGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUserGetGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableMerchantUserGetGet200ResponseData struct {
	value *MerchantUserGetGet200ResponseData
	isSet bool
}

func (v NullableMerchantUserGetGet200ResponseData) Get() *MerchantUserGetGet200ResponseData {
	return v.value
}

func (v *NullableMerchantUserGetGet200ResponseData) Set(val *MerchantUserGetGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUserGetGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUserGetGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUserGetGet200ResponseData(val *MerchantUserGetGet200ResponseData) *NullableMerchantUserGetGet200ResponseData {
	return &NullableMerchantUserGetGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantUserGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUserGetGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


