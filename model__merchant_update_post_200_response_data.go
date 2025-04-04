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

// checks if the MerchantUpdatePost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUpdatePost200ResponseData{}

// MerchantUpdatePost200ResponseData struct for MerchantUpdatePost200ResponseData
type MerchantUpdatePost200ResponseData struct {
	Merchant *UnibeeApiBeanMerchant `json:"merchant,omitempty"`
}

// NewMerchantUpdatePost200ResponseData instantiates a new MerchantUpdatePost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUpdatePost200ResponseData() *MerchantUpdatePost200ResponseData {
	this := MerchantUpdatePost200ResponseData{}
	return &this
}

// NewMerchantUpdatePost200ResponseDataWithDefaults instantiates a new MerchantUpdatePost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUpdatePost200ResponseDataWithDefaults() *MerchantUpdatePost200ResponseData {
	this := MerchantUpdatePost200ResponseData{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *MerchantUpdatePost200ResponseData) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUpdatePost200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *MerchantUpdatePost200ResponseData) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *MerchantUpdatePost200ResponseData) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

func (o MerchantUpdatePost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUpdatePost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	return toSerialize, nil
}

type NullableMerchantUpdatePost200ResponseData struct {
	value *MerchantUpdatePost200ResponseData
	isSet bool
}

func (v NullableMerchantUpdatePost200ResponseData) Get() *MerchantUpdatePost200ResponseData {
	return v.value
}

func (v *NullableMerchantUpdatePost200ResponseData) Set(val *MerchantUpdatePost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUpdatePost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUpdatePost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUpdatePost200ResponseData(val *MerchantUpdatePost200ResponseData) *NullableMerchantUpdatePost200ResponseData {
	return &NullableMerchantUpdatePost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantUpdatePost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUpdatePost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


