/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricEditRes{}

// UnibeeApiMerchantMetricEditRes struct for UnibeeApiMerchantMetricEditRes
type UnibeeApiMerchantMetricEditRes struct {
	MerchantMetric *UnibeeApiBeanMerchantMetric `json:"merchantMetric,omitempty"`
}

// NewUnibeeApiMerchantMetricEditRes instantiates a new UnibeeApiMerchantMetricEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricEditRes() *UnibeeApiMerchantMetricEditRes {
	this := UnibeeApiMerchantMetricEditRes{}
	return &this
}

// NewUnibeeApiMerchantMetricEditResWithDefaults instantiates a new UnibeeApiMerchantMetricEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricEditResWithDefaults() *UnibeeApiMerchantMetricEditRes {
	this := UnibeeApiMerchantMetricEditRes{}
	return &this
}

// GetMerchantMetric returns the MerchantMetric field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEditRes) GetMerchantMetric() UnibeeApiBeanMerchantMetric {
	if o == nil || IsNil(o.MerchantMetric) {
		var ret UnibeeApiBeanMerchantMetric
		return ret
	}
	return *o.MerchantMetric
}

// GetMerchantMetricOk returns a tuple with the MerchantMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditRes) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool) {
	if o == nil || IsNil(o.MerchantMetric) {
		return nil, false
	}
	return o.MerchantMetric, true
}

// HasMerchantMetric returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEditRes) HasMerchantMetric() bool {
	if o != nil && !IsNil(o.MerchantMetric) {
		return true
	}

	return false
}

// SetMerchantMetric gets a reference to the given UnibeeApiBeanMerchantMetric and assigns it to the MerchantMetric field.
func (o *UnibeeApiMerchantMetricEditRes) SetMerchantMetric(v UnibeeApiBeanMerchantMetric) {
	o.MerchantMetric = &v
}

func (o UnibeeApiMerchantMetricEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetric) {
		toSerialize["merchantMetric"] = o.MerchantMetric
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricEditRes struct {
	value *UnibeeApiMerchantMetricEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricEditRes) Get() *UnibeeApiMerchantMetricEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricEditRes) Set(val *UnibeeApiMerchantMetricEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricEditRes(val *UnibeeApiMerchantMetricEditRes) *NullableUnibeeApiMerchantMetricEditRes {
	return &NullableUnibeeApiMerchantMetricEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


