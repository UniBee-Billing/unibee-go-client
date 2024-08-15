/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantVatCountryListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantVatCountryListRes{}

// UnibeeApiMerchantVatCountryListRes struct for UnibeeApiMerchantVatCountryListRes
type UnibeeApiMerchantVatCountryListRes struct {
	// VatCountryList
	VatCountryList []UnibeeApiBeanVatCountryRate `json:"vatCountryList,omitempty"`
}

// NewUnibeeApiMerchantVatCountryListRes instantiates a new UnibeeApiMerchantVatCountryListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantVatCountryListRes() *UnibeeApiMerchantVatCountryListRes {
	this := UnibeeApiMerchantVatCountryListRes{}
	return &this
}

// NewUnibeeApiMerchantVatCountryListResWithDefaults instantiates a new UnibeeApiMerchantVatCountryListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantVatCountryListResWithDefaults() *UnibeeApiMerchantVatCountryListRes {
	this := UnibeeApiMerchantVatCountryListRes{}
	return &this
}

// GetVatCountryList returns the VatCountryList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantVatCountryListRes) GetVatCountryList() []UnibeeApiBeanVatCountryRate {
	if o == nil || IsNil(o.VatCountryList) {
		var ret []UnibeeApiBeanVatCountryRate
		return ret
	}
	return o.VatCountryList
}

// GetVatCountryListOk returns a tuple with the VatCountryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatCountryListRes) GetVatCountryListOk() ([]UnibeeApiBeanVatCountryRate, bool) {
	if o == nil || IsNil(o.VatCountryList) {
		return nil, false
	}
	return o.VatCountryList, true
}

// HasVatCountryList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantVatCountryListRes) HasVatCountryList() bool {
	if o != nil && !IsNil(o.VatCountryList) {
		return true
	}

	return false
}

// SetVatCountryList gets a reference to the given []UnibeeApiBeanVatCountryRate and assigns it to the VatCountryList field.
func (o *UnibeeApiMerchantVatCountryListRes) SetVatCountryList(v []UnibeeApiBeanVatCountryRate) {
	o.VatCountryList = v
}

func (o UnibeeApiMerchantVatCountryListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantVatCountryListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VatCountryList) {
		toSerialize["vatCountryList"] = o.VatCountryList
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantVatCountryListRes struct {
	value *UnibeeApiMerchantVatCountryListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantVatCountryListRes) Get() *UnibeeApiMerchantVatCountryListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantVatCountryListRes) Set(val *UnibeeApiMerchantVatCountryListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantVatCountryListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantVatCountryListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantVatCountryListRes(val *UnibeeApiMerchantVatCountryListRes) *NullableUnibeeApiMerchantVatCountryListRes {
	return &NullableUnibeeApiMerchantVatCountryListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantVatCountryListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantVatCountryListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


