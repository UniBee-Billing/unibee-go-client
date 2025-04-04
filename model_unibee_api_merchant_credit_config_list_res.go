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

// checks if the UnibeeApiMerchantCreditConfigListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditConfigListRes{}

// UnibeeApiMerchantCreditConfigListRes struct for UnibeeApiMerchantCreditConfigListRes
type UnibeeApiMerchantCreditConfigListRes struct {
	// CreditConfig List Object
	CreditConfigs []UnibeeApiBeanCreditConfig `json:"creditConfigs,omitempty"`
}

// NewUnibeeApiMerchantCreditConfigListRes instantiates a new UnibeeApiMerchantCreditConfigListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditConfigListRes() *UnibeeApiMerchantCreditConfigListRes {
	this := UnibeeApiMerchantCreditConfigListRes{}
	return &this
}

// NewUnibeeApiMerchantCreditConfigListResWithDefaults instantiates a new UnibeeApiMerchantCreditConfigListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditConfigListResWithDefaults() *UnibeeApiMerchantCreditConfigListRes {
	this := UnibeeApiMerchantCreditConfigListRes{}
	return &this
}

// GetCreditConfigs returns the CreditConfigs field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditConfigListRes) GetCreditConfigs() []UnibeeApiBeanCreditConfig {
	if o == nil || IsNil(o.CreditConfigs) {
		var ret []UnibeeApiBeanCreditConfig
		return ret
	}
	return o.CreditConfigs
}

// GetCreditConfigsOk returns a tuple with the CreditConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditConfigListRes) GetCreditConfigsOk() ([]UnibeeApiBeanCreditConfig, bool) {
	if o == nil || IsNil(o.CreditConfigs) {
		return nil, false
	}
	return o.CreditConfigs, true
}

// HasCreditConfigs returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditConfigListRes) HasCreditConfigs() bool {
	if o != nil && !IsNil(o.CreditConfigs) {
		return true
	}

	return false
}

// SetCreditConfigs gets a reference to the given []UnibeeApiBeanCreditConfig and assigns it to the CreditConfigs field.
func (o *UnibeeApiMerchantCreditConfigListRes) SetCreditConfigs(v []UnibeeApiBeanCreditConfig) {
	o.CreditConfigs = v
}

func (o UnibeeApiMerchantCreditConfigListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditConfigListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditConfigs) {
		toSerialize["creditConfigs"] = o.CreditConfigs
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditConfigListRes struct {
	value *UnibeeApiMerchantCreditConfigListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditConfigListRes) Get() *UnibeeApiMerchantCreditConfigListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditConfigListRes) Set(val *UnibeeApiMerchantCreditConfigListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditConfigListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditConfigListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditConfigListRes(val *UnibeeApiMerchantCreditConfigListRes) *NullableUnibeeApiMerchantCreditConfigListRes {
	return &NullableUnibeeApiMerchantCreditConfigListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditConfigListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditConfigListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


