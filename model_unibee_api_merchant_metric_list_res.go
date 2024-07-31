/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricListRes{}

// UnibeeApiMerchantMetricListRes struct for UnibeeApiMerchantMetricListRes
type UnibeeApiMerchantMetricListRes struct {
	// MerchantMetrics
	MerchantMetrics []UnibeeApiBeanMerchantMetric `json:"merchantMetrics,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantMetricListRes instantiates a new UnibeeApiMerchantMetricListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricListRes() *UnibeeApiMerchantMetricListRes {
	this := UnibeeApiMerchantMetricListRes{}
	return &this
}

// NewUnibeeApiMerchantMetricListResWithDefaults instantiates a new UnibeeApiMerchantMetricListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricListResWithDefaults() *UnibeeApiMerchantMetricListRes {
	this := UnibeeApiMerchantMetricListRes{}
	return &this
}

// GetMerchantMetrics returns the MerchantMetrics field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricListRes) GetMerchantMetrics() []UnibeeApiBeanMerchantMetric {
	if o == nil || IsNil(o.MerchantMetrics) {
		var ret []UnibeeApiBeanMerchantMetric
		return ret
	}
	return o.MerchantMetrics
}

// GetMerchantMetricsOk returns a tuple with the MerchantMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricListRes) GetMerchantMetricsOk() ([]UnibeeApiBeanMerchantMetric, bool) {
	if o == nil || IsNil(o.MerchantMetrics) {
		return nil, false
	}
	return o.MerchantMetrics, true
}

// HasMerchantMetrics returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricListRes) HasMerchantMetrics() bool {
	if o != nil && !IsNil(o.MerchantMetrics) {
		return true
	}

	return false
}

// SetMerchantMetrics gets a reference to the given []UnibeeApiBeanMerchantMetric and assigns it to the MerchantMetrics field.
func (o *UnibeeApiMerchantMetricListRes) SetMerchantMetrics(v []UnibeeApiBeanMerchantMetric) {
	o.MerchantMetrics = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantMetricListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantMetricListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetrics) {
		toSerialize["merchantMetrics"] = o.MerchantMetrics
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricListRes struct {
	value *UnibeeApiMerchantMetricListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricListRes) Get() *UnibeeApiMerchantMetricListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricListRes) Set(val *UnibeeApiMerchantMetricListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricListRes(val *UnibeeApiMerchantMetricListRes) *NullableUnibeeApiMerchantMetricListRes {
	return &NullableUnibeeApiMerchantMetricListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


