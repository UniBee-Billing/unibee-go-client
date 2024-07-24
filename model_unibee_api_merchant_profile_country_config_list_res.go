/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProfileCountryConfigListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileCountryConfigListRes{}

// UnibeeApiMerchantProfileCountryConfigListRes struct for UnibeeApiMerchantProfileCountryConfigListRes
type UnibeeApiMerchantProfileCountryConfigListRes struct {
	// Configs
	Configs []UnibeeApiBeanMerchantCountryConfigSimplify `json:"configs,omitempty"`
}

// NewUnibeeApiMerchantProfileCountryConfigListRes instantiates a new UnibeeApiMerchantProfileCountryConfigListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileCountryConfigListRes() *UnibeeApiMerchantProfileCountryConfigListRes {
	this := UnibeeApiMerchantProfileCountryConfigListRes{}
	return &this
}

// NewUnibeeApiMerchantProfileCountryConfigListResWithDefaults instantiates a new UnibeeApiMerchantProfileCountryConfigListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileCountryConfigListResWithDefaults() *UnibeeApiMerchantProfileCountryConfigListRes {
	this := UnibeeApiMerchantProfileCountryConfigListRes{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileCountryConfigListRes) GetConfigs() []UnibeeApiBeanMerchantCountryConfigSimplify {
	if o == nil || IsNil(o.Configs) {
		var ret []UnibeeApiBeanMerchantCountryConfigSimplify
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileCountryConfigListRes) GetConfigsOk() ([]UnibeeApiBeanMerchantCountryConfigSimplify, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileCountryConfigListRes) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []UnibeeApiBeanMerchantCountryConfigSimplify and assigns it to the Configs field.
func (o *UnibeeApiMerchantProfileCountryConfigListRes) SetConfigs(v []UnibeeApiBeanMerchantCountryConfigSimplify) {
	o.Configs = v
}

func (o UnibeeApiMerchantProfileCountryConfigListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileCountryConfigListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProfileCountryConfigListRes struct {
	value *UnibeeApiMerchantProfileCountryConfigListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileCountryConfigListRes) Get() *UnibeeApiMerchantProfileCountryConfigListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileCountryConfigListRes) Set(val *UnibeeApiMerchantProfileCountryConfigListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileCountryConfigListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileCountryConfigListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileCountryConfigListRes(val *UnibeeApiMerchantProfileCountryConfigListRes) *NullableUnibeeApiMerchantProfileCountryConfigListRes {
	return &NullableUnibeeApiMerchantProfileCountryConfigListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileCountryConfigListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileCountryConfigListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


