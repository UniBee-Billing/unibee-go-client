/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantCountryConfigListPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCountryConfigListPost200ResponseData{}

// MerchantCountryConfigListPost200ResponseData struct for MerchantCountryConfigListPost200ResponseData
type MerchantCountryConfigListPost200ResponseData struct {
	// Configs
	Configs []UnibeeApiBeanMerchantCountryConfig `json:"configs,omitempty"`
}

// NewMerchantCountryConfigListPost200ResponseData instantiates a new MerchantCountryConfigListPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCountryConfigListPost200ResponseData() *MerchantCountryConfigListPost200ResponseData {
	this := MerchantCountryConfigListPost200ResponseData{}
	return &this
}

// NewMerchantCountryConfigListPost200ResponseDataWithDefaults instantiates a new MerchantCountryConfigListPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCountryConfigListPost200ResponseDataWithDefaults() *MerchantCountryConfigListPost200ResponseData {
	this := MerchantCountryConfigListPost200ResponseData{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *MerchantCountryConfigListPost200ResponseData) GetConfigs() []UnibeeApiBeanMerchantCountryConfig {
	if o == nil || IsNil(o.Configs) {
		var ret []UnibeeApiBeanMerchantCountryConfig
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCountryConfigListPost200ResponseData) GetConfigsOk() ([]UnibeeApiBeanMerchantCountryConfig, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *MerchantCountryConfigListPost200ResponseData) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []UnibeeApiBeanMerchantCountryConfig and assigns it to the Configs field.
func (o *MerchantCountryConfigListPost200ResponseData) SetConfigs(v []UnibeeApiBeanMerchantCountryConfig) {
	o.Configs = v
}

func (o MerchantCountryConfigListPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCountryConfigListPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	return toSerialize, nil
}

type NullableMerchantCountryConfigListPost200ResponseData struct {
	value *MerchantCountryConfigListPost200ResponseData
	isSet bool
}

func (v NullableMerchantCountryConfigListPost200ResponseData) Get() *MerchantCountryConfigListPost200ResponseData {
	return v.value
}

func (v *NullableMerchantCountryConfigListPost200ResponseData) Set(val *MerchantCountryConfigListPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCountryConfigListPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCountryConfigListPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCountryConfigListPost200ResponseData(val *MerchantCountryConfigListPost200ResponseData) *NullableMerchantCountryConfigListPost200ResponseData {
	return &NullableMerchantCountryConfigListPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCountryConfigListPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCountryConfigListPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


