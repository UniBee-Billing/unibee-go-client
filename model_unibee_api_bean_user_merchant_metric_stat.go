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

// checks if the UnibeeApiBeanUserMerchantMetricStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanUserMerchantMetricStat{}

// UnibeeApiBeanUserMerchantMetricStat struct for UnibeeApiBeanUserMerchantMetricStat
type UnibeeApiBeanUserMerchantMetricStat struct {
	CurrentUseValue *int64 `json:"CurrentUseValue,omitempty"`
	MetricLimit *UnibeeApiBeanPlanMetricLimitDetail `json:"MetricLimit,omitempty"`
}

// NewUnibeeApiBeanUserMerchantMetricStat instantiates a new UnibeeApiBeanUserMerchantMetricStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanUserMerchantMetricStat() *UnibeeApiBeanUserMerchantMetricStat {
	this := UnibeeApiBeanUserMerchantMetricStat{}
	return &this
}

// NewUnibeeApiBeanUserMerchantMetricStatWithDefaults instantiates a new UnibeeApiBeanUserMerchantMetricStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanUserMerchantMetricStatWithDefaults() *UnibeeApiBeanUserMerchantMetricStat {
	this := UnibeeApiBeanUserMerchantMetricStat{}
	return &this
}

// GetCurrentUseValue returns the CurrentUseValue field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMerchantMetricStat) GetCurrentUseValue() int64 {
	if o == nil || IsNil(o.CurrentUseValue) {
		var ret int64
		return ret
	}
	return *o.CurrentUseValue
}

// GetCurrentUseValueOk returns a tuple with the CurrentUseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMerchantMetricStat) GetCurrentUseValueOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUseValue) {
		return nil, false
	}
	return o.CurrentUseValue, true
}

// HasCurrentUseValue returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMerchantMetricStat) HasCurrentUseValue() bool {
	if o != nil && !IsNil(o.CurrentUseValue) {
		return true
	}

	return false
}

// SetCurrentUseValue gets a reference to the given int64 and assigns it to the CurrentUseValue field.
func (o *UnibeeApiBeanUserMerchantMetricStat) SetCurrentUseValue(v int64) {
	o.CurrentUseValue = &v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMerchantMetricStat) GetMetricLimit() UnibeeApiBeanPlanMetricLimitDetail {
	if o == nil || IsNil(o.MetricLimit) {
		var ret UnibeeApiBeanPlanMetricLimitDetail
		return ret
	}
	return *o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMerchantMetricStat) GetMetricLimitOk() (*UnibeeApiBeanPlanMetricLimitDetail, bool) {
	if o == nil || IsNil(o.MetricLimit) {
		return nil, false
	}
	return o.MetricLimit, true
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMerchantMetricStat) HasMetricLimit() bool {
	if o != nil && !IsNil(o.MetricLimit) {
		return true
	}

	return false
}

// SetMetricLimit gets a reference to the given UnibeeApiBeanPlanMetricLimitDetail and assigns it to the MetricLimit field.
func (o *UnibeeApiBeanUserMerchantMetricStat) SetMetricLimit(v UnibeeApiBeanPlanMetricLimitDetail) {
	o.MetricLimit = &v
}

func (o UnibeeApiBeanUserMerchantMetricStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanUserMerchantMetricStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentUseValue) {
		toSerialize["CurrentUseValue"] = o.CurrentUseValue
	}
	if !IsNil(o.MetricLimit) {
		toSerialize["MetricLimit"] = o.MetricLimit
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanUserMerchantMetricStat struct {
	value *UnibeeApiBeanUserMerchantMetricStat
	isSet bool
}

func (v NullableUnibeeApiBeanUserMerchantMetricStat) Get() *UnibeeApiBeanUserMerchantMetricStat {
	return v.value
}

func (v *NullableUnibeeApiBeanUserMerchantMetricStat) Set(val *UnibeeApiBeanUserMerchantMetricStat) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanUserMerchantMetricStat) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanUserMerchantMetricStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanUserMerchantMetricStat(val *UnibeeApiBeanUserMerchantMetricStat) *NullableUnibeeApiBeanUserMerchantMetricStat {
	return &NullableUnibeeApiBeanUserMerchantMetricStat{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanUserMerchantMetricStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanUserMerchantMetricStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


