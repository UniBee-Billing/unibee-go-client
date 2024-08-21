/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionConfigRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionConfigRes{}

// UnibeeApiMerchantSubscriptionConfigRes struct for UnibeeApiMerchantSubscriptionConfigRes
type UnibeeApiMerchantSubscriptionConfigRes struct {
	Config *UnibeeApiBeanSubscriptionConfig `json:"config,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionConfigRes instantiates a new UnibeeApiMerchantSubscriptionConfigRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionConfigRes() *UnibeeApiMerchantSubscriptionConfigRes {
	this := UnibeeApiMerchantSubscriptionConfigRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionConfigResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionConfigRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionConfigResWithDefaults() *UnibeeApiMerchantSubscriptionConfigRes {
	this := UnibeeApiMerchantSubscriptionConfigRes{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigRes) GetConfig() UnibeeApiBeanSubscriptionConfig {
	if o == nil || IsNil(o.Config) {
		var ret UnibeeApiBeanSubscriptionConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigRes) GetConfigOk() (*UnibeeApiBeanSubscriptionConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigRes) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given UnibeeApiBeanSubscriptionConfig and assigns it to the Config field.
func (o *UnibeeApiMerchantSubscriptionConfigRes) SetConfig(v UnibeeApiBeanSubscriptionConfig) {
	o.Config = &v
}

func (o UnibeeApiMerchantSubscriptionConfigRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionConfigRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionConfigRes struct {
	value *UnibeeApiMerchantSubscriptionConfigRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionConfigRes) Get() *UnibeeApiMerchantSubscriptionConfigRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigRes) Set(val *UnibeeApiMerchantSubscriptionConfigRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionConfigRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionConfigRes(val *UnibeeApiMerchantSubscriptionConfigRes) *NullableUnibeeApiMerchantSubscriptionConfigRes {
	return &NullableUnibeeApiMerchantSubscriptionConfigRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionConfigRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


