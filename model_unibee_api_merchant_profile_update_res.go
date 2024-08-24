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

// checks if the UnibeeApiMerchantProfileUpdateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileUpdateRes{}

// UnibeeApiMerchantProfileUpdateRes struct for UnibeeApiMerchantProfileUpdateRes
type UnibeeApiMerchantProfileUpdateRes struct {
	Merchant *UnibeeApiBeanMerchant `json:"merchant,omitempty"`
}

// NewUnibeeApiMerchantProfileUpdateRes instantiates a new UnibeeApiMerchantProfileUpdateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileUpdateRes() *UnibeeApiMerchantProfileUpdateRes {
	this := UnibeeApiMerchantProfileUpdateRes{}
	return &this
}

// NewUnibeeApiMerchantProfileUpdateResWithDefaults instantiates a new UnibeeApiMerchantProfileUpdateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileUpdateResWithDefaults() *UnibeeApiMerchantProfileUpdateRes {
	this := UnibeeApiMerchantProfileUpdateRes{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateRes) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateRes) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *UnibeeApiMerchantProfileUpdateRes) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

func (o UnibeeApiMerchantProfileUpdateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileUpdateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProfileUpdateRes struct {
	value *UnibeeApiMerchantProfileUpdateRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileUpdateRes) Get() *UnibeeApiMerchantProfileUpdateRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileUpdateRes) Set(val *UnibeeApiMerchantProfileUpdateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileUpdateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileUpdateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileUpdateRes(val *UnibeeApiMerchantProfileUpdateRes) *NullableUnibeeApiMerchantProfileUpdateRes {
	return &NullableUnibeeApiMerchantProfileUpdateRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileUpdateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileUpdateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


