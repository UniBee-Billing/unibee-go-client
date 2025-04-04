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

// checks if the UnibeeApiMerchantMetricDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDetailRes{}

// UnibeeApiMerchantMetricDetailRes struct for UnibeeApiMerchantMetricDetailRes
type UnibeeApiMerchantMetricDetailRes struct {
	MerchantMetric *UnibeeApiBeanMerchantMetric `json:"merchantMetric,omitempty"`
}

// NewUnibeeApiMerchantMetricDetailRes instantiates a new UnibeeApiMerchantMetricDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDetailRes() *UnibeeApiMerchantMetricDetailRes {
	this := UnibeeApiMerchantMetricDetailRes{}
	return &this
}

// NewUnibeeApiMerchantMetricDetailResWithDefaults instantiates a new UnibeeApiMerchantMetricDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDetailResWithDefaults() *UnibeeApiMerchantMetricDetailRes {
	this := UnibeeApiMerchantMetricDetailRes{}
	return &this
}

// GetMerchantMetric returns the MerchantMetric field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricDetailRes) GetMerchantMetric() UnibeeApiBeanMerchantMetric {
	if o == nil || IsNil(o.MerchantMetric) {
		var ret UnibeeApiBeanMerchantMetric
		return ret
	}
	return *o.MerchantMetric
}

// GetMerchantMetricOk returns a tuple with the MerchantMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDetailRes) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool) {
	if o == nil || IsNil(o.MerchantMetric) {
		return nil, false
	}
	return o.MerchantMetric, true
}

// HasMerchantMetric returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricDetailRes) HasMerchantMetric() bool {
	if o != nil && !IsNil(o.MerchantMetric) {
		return true
	}

	return false
}

// SetMerchantMetric gets a reference to the given UnibeeApiBeanMerchantMetric and assigns it to the MerchantMetric field.
func (o *UnibeeApiMerchantMetricDetailRes) SetMerchantMetric(v UnibeeApiBeanMerchantMetric) {
	o.MerchantMetric = &v
}

func (o UnibeeApiMerchantMetricDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetric) {
		toSerialize["merchantMetric"] = o.MerchantMetric
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricDetailRes struct {
	value *UnibeeApiMerchantMetricDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDetailRes) Get() *UnibeeApiMerchantMetricDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDetailRes) Set(val *UnibeeApiMerchantMetricDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDetailRes(val *UnibeeApiMerchantMetricDetailRes) *NullableUnibeeApiMerchantMetricDetailRes {
	return &NullableUnibeeApiMerchantMetricDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


