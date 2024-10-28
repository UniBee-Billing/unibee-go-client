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

// checks if the MerchantUserListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUserListGet200ResponseData{}

// MerchantUserListGet200ResponseData struct for MerchantUserListGet200ResponseData
type MerchantUserListGet200ResponseData struct {
	// Total
	Total *int32 `json:"total,omitempty"`
	// User Account Object List
	UserAccounts []UnibeeApiBeanDetailUserAccountDetail `json:"userAccounts,omitempty"`
}

// NewMerchantUserListGet200ResponseData instantiates a new MerchantUserListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUserListGet200ResponseData() *MerchantUserListGet200ResponseData {
	this := MerchantUserListGet200ResponseData{}
	return &this
}

// NewMerchantUserListGet200ResponseDataWithDefaults instantiates a new MerchantUserListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUserListGet200ResponseDataWithDefaults() *MerchantUserListGet200ResponseData {
	this := MerchantUserListGet200ResponseData{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantUserListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantUserListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantUserListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

// GetUserAccounts returns the UserAccounts field value if set, zero value otherwise.
func (o *MerchantUserListGet200ResponseData) GetUserAccounts() []UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.UserAccounts) {
		var ret []UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return o.UserAccounts
}

// GetUserAccountsOk returns a tuple with the UserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserListGet200ResponseData) GetUserAccountsOk() ([]UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.UserAccounts) {
		return nil, false
	}
	return o.UserAccounts, true
}

// HasUserAccounts returns a boolean if a field has been set.
func (o *MerchantUserListGet200ResponseData) HasUserAccounts() bool {
	if o != nil && !IsNil(o.UserAccounts) {
		return true
	}

	return false
}

// SetUserAccounts gets a reference to the given []UnibeeApiBeanDetailUserAccountDetail and assigns it to the UserAccounts field.
func (o *MerchantUserListGet200ResponseData) SetUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail) {
	o.UserAccounts = v
}

func (o MerchantUserListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUserListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.UserAccounts) {
		toSerialize["userAccounts"] = o.UserAccounts
	}
	return toSerialize, nil
}

type NullableMerchantUserListGet200ResponseData struct {
	value *MerchantUserListGet200ResponseData
	isSet bool
}

func (v NullableMerchantUserListGet200ResponseData) Get() *MerchantUserListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantUserListGet200ResponseData) Set(val *MerchantUserListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUserListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUserListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUserListGet200ResponseData(val *MerchantUserListGet200ResponseData) *NullableMerchantUserListGet200ResponseData {
	return &NullableMerchantUserListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantUserListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUserListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


