/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionConfigUpdateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionConfigUpdateRes{}

// UnibeeApiMerchantSubscriptionConfigUpdateRes struct for UnibeeApiMerchantSubscriptionConfigUpdateRes
type UnibeeApiMerchantSubscriptionConfigUpdateRes struct {
	Config *UnibeeApiBeanSubscriptionConfig `json:"config,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionConfigUpdateRes instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionConfigUpdateRes() *UnibeeApiMerchantSubscriptionConfigUpdateRes {
	this := UnibeeApiMerchantSubscriptionConfigUpdateRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionConfigUpdateResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionConfigUpdateResWithDefaults() *UnibeeApiMerchantSubscriptionConfigUpdateRes {
	this := UnibeeApiMerchantSubscriptionConfigUpdateRes{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateRes) GetConfig() UnibeeApiBeanSubscriptionConfig {
	if o == nil || IsNil(o.Config) {
		var ret UnibeeApiBeanSubscriptionConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateRes) GetConfigOk() (*UnibeeApiBeanSubscriptionConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateRes) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given UnibeeApiBeanSubscriptionConfig and assigns it to the Config field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateRes) SetConfig(v UnibeeApiBeanSubscriptionConfig) {
	o.Config = &v
}

func (o UnibeeApiMerchantSubscriptionConfigUpdateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionConfigUpdateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionConfigUpdateRes struct {
	value *UnibeeApiMerchantSubscriptionConfigUpdateRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) Get() *UnibeeApiMerchantSubscriptionConfigUpdateRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) Set(val *UnibeeApiMerchantSubscriptionConfigUpdateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionConfigUpdateRes(val *UnibeeApiMerchantSubscriptionConfigUpdateRes) *NullableUnibeeApiMerchantSubscriptionConfigUpdateRes {
	return &NullableUnibeeApiMerchantSubscriptionConfigUpdateRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


