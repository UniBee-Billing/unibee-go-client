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

// checks if the MerchantSubscriptionConfigGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionConfigGet200ResponseData{}

// MerchantSubscriptionConfigGet200ResponseData struct for MerchantSubscriptionConfigGet200ResponseData
type MerchantSubscriptionConfigGet200ResponseData struct {
	Config *UnibeeApiBeanSubscriptionConfig `json:"config,omitempty"`
}

// NewMerchantSubscriptionConfigGet200ResponseData instantiates a new MerchantSubscriptionConfigGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionConfigGet200ResponseData() *MerchantSubscriptionConfigGet200ResponseData {
	this := MerchantSubscriptionConfigGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionConfigGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionConfigGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionConfigGet200ResponseDataWithDefaults() *MerchantSubscriptionConfigGet200ResponseData {
	this := MerchantSubscriptionConfigGet200ResponseData{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *MerchantSubscriptionConfigGet200ResponseData) GetConfig() UnibeeApiBeanSubscriptionConfig {
	if o == nil || IsNil(o.Config) {
		var ret UnibeeApiBeanSubscriptionConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionConfigGet200ResponseData) GetConfigOk() (*UnibeeApiBeanSubscriptionConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *MerchantSubscriptionConfigGet200ResponseData) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given UnibeeApiBeanSubscriptionConfig and assigns it to the Config field.
func (o *MerchantSubscriptionConfigGet200ResponseData) SetConfig(v UnibeeApiBeanSubscriptionConfig) {
	o.Config = &v
}

func (o MerchantSubscriptionConfigGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionConfigGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionConfigGet200ResponseData struct {
	value *MerchantSubscriptionConfigGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionConfigGet200ResponseData) Get() *MerchantSubscriptionConfigGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionConfigGet200ResponseData) Set(val *MerchantSubscriptionConfigGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionConfigGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionConfigGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionConfigGet200ResponseData(val *MerchantSubscriptionConfigGet200ResponseData) *NullableMerchantSubscriptionConfigGet200ResponseData {
	return &NullableMerchantSubscriptionConfigGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionConfigGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionConfigGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


