/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantVatVatNumberValidatePost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantVatVatNumberValidatePost200ResponseData{}

// MerchantVatVatNumberValidatePost200ResponseData struct for MerchantVatVatNumberValidatePost200ResponseData
type MerchantVatVatNumberValidatePost200ResponseData struct {
	VatNumberValidate *UnibeeApiBeanValidResult `json:"vatNumberValidate,omitempty"`
}

// NewMerchantVatVatNumberValidatePost200ResponseData instantiates a new MerchantVatVatNumberValidatePost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantVatVatNumberValidatePost200ResponseData() *MerchantVatVatNumberValidatePost200ResponseData {
	this := MerchantVatVatNumberValidatePost200ResponseData{}
	return &this
}

// NewMerchantVatVatNumberValidatePost200ResponseDataWithDefaults instantiates a new MerchantVatVatNumberValidatePost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantVatVatNumberValidatePost200ResponseDataWithDefaults() *MerchantVatVatNumberValidatePost200ResponseData {
	this := MerchantVatVatNumberValidatePost200ResponseData{}
	return &this
}

// GetVatNumberValidate returns the VatNumberValidate field value if set, zero value otherwise.
func (o *MerchantVatVatNumberValidatePost200ResponseData) GetVatNumberValidate() UnibeeApiBeanValidResult {
	if o == nil || IsNil(o.VatNumberValidate) {
		var ret UnibeeApiBeanValidResult
		return ret
	}
	return *o.VatNumberValidate
}

// GetVatNumberValidateOk returns a tuple with the VatNumberValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantVatVatNumberValidatePost200ResponseData) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool) {
	if o == nil || IsNil(o.VatNumberValidate) {
		return nil, false
	}
	return o.VatNumberValidate, true
}

// HasVatNumberValidate returns a boolean if a field has been set.
func (o *MerchantVatVatNumberValidatePost200ResponseData) HasVatNumberValidate() bool {
	if o != nil && !IsNil(o.VatNumberValidate) {
		return true
	}

	return false
}

// SetVatNumberValidate gets a reference to the given UnibeeApiBeanValidResult and assigns it to the VatNumberValidate field.
func (o *MerchantVatVatNumberValidatePost200ResponseData) SetVatNumberValidate(v UnibeeApiBeanValidResult) {
	o.VatNumberValidate = &v
}

func (o MerchantVatVatNumberValidatePost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantVatVatNumberValidatePost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VatNumberValidate) {
		toSerialize["vatNumberValidate"] = o.VatNumberValidate
	}
	return toSerialize, nil
}

type NullableMerchantVatVatNumberValidatePost200ResponseData struct {
	value *MerchantVatVatNumberValidatePost200ResponseData
	isSet bool
}

func (v NullableMerchantVatVatNumberValidatePost200ResponseData) Get() *MerchantVatVatNumberValidatePost200ResponseData {
	return v.value
}

func (v *NullableMerchantVatVatNumberValidatePost200ResponseData) Set(val *MerchantVatVatNumberValidatePost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantVatVatNumberValidatePost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantVatVatNumberValidatePost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantVatVatNumberValidatePost200ResponseData(val *MerchantVatVatNumberValidatePost200ResponseData) *NullableMerchantVatVatNumberValidatePost200ResponseData {
	return &NullableMerchantVatVatNumberValidatePost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantVatVatNumberValidatePost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantVatVatNumberValidatePost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


