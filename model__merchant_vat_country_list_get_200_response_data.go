/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantVatCountryListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantVatCountryListGet200ResponseData{}

// MerchantVatCountryListGet200ResponseData struct for MerchantVatCountryListGet200ResponseData
type MerchantVatCountryListGet200ResponseData struct {
	// VatCountryList
	VatCountryList []UnibeeApiBeanVatCountryRate `json:"vatCountryList,omitempty"`
}

// NewMerchantVatCountryListGet200ResponseData instantiates a new MerchantVatCountryListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantVatCountryListGet200ResponseData() *MerchantVatCountryListGet200ResponseData {
	this := MerchantVatCountryListGet200ResponseData{}
	return &this
}

// NewMerchantVatCountryListGet200ResponseDataWithDefaults instantiates a new MerchantVatCountryListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantVatCountryListGet200ResponseDataWithDefaults() *MerchantVatCountryListGet200ResponseData {
	this := MerchantVatCountryListGet200ResponseData{}
	return &this
}

// GetVatCountryList returns the VatCountryList field value if set, zero value otherwise.
func (o *MerchantVatCountryListGet200ResponseData) GetVatCountryList() []UnibeeApiBeanVatCountryRate {
	if o == nil || IsNil(o.VatCountryList) {
		var ret []UnibeeApiBeanVatCountryRate
		return ret
	}
	return o.VatCountryList
}

// GetVatCountryListOk returns a tuple with the VatCountryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantVatCountryListGet200ResponseData) GetVatCountryListOk() ([]UnibeeApiBeanVatCountryRate, bool) {
	if o == nil || IsNil(o.VatCountryList) {
		return nil, false
	}
	return o.VatCountryList, true
}

// HasVatCountryList returns a boolean if a field has been set.
func (o *MerchantVatCountryListGet200ResponseData) HasVatCountryList() bool {
	if o != nil && !IsNil(o.VatCountryList) {
		return true
	}

	return false
}

// SetVatCountryList gets a reference to the given []UnibeeApiBeanVatCountryRate and assigns it to the VatCountryList field.
func (o *MerchantVatCountryListGet200ResponseData) SetVatCountryList(v []UnibeeApiBeanVatCountryRate) {
	o.VatCountryList = v
}

func (o MerchantVatCountryListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantVatCountryListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VatCountryList) {
		toSerialize["vatCountryList"] = o.VatCountryList
	}
	return toSerialize, nil
}

type NullableMerchantVatCountryListGet200ResponseData struct {
	value *MerchantVatCountryListGet200ResponseData
	isSet bool
}

func (v NullableMerchantVatCountryListGet200ResponseData) Get() *MerchantVatCountryListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantVatCountryListGet200ResponseData) Set(val *MerchantVatCountryListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantVatCountryListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantVatCountryListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantVatCountryListGet200ResponseData(val *MerchantVatCountryListGet200ResponseData) *NullableMerchantVatCountryListGet200ResponseData {
	return &NullableMerchantVatCountryListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantVatCountryListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantVatCountryListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


