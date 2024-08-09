/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiBeanBulkMetricLimitPlanBindingParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanBulkMetricLimitPlanBindingParam{}

// UnibeeApiBeanBulkMetricLimitPlanBindingParam struct for UnibeeApiBeanBulkMetricLimitPlanBindingParam
type UnibeeApiBeanBulkMetricLimitPlanBindingParam struct {
	// MetricId
	MetricId int64 `json:"metricId"`
	// MetricLimit
	MetricLimit int64 `json:"metricLimit"`
}

type _UnibeeApiBeanBulkMetricLimitPlanBindingParam UnibeeApiBeanBulkMetricLimitPlanBindingParam

// NewUnibeeApiBeanBulkMetricLimitPlanBindingParam instantiates a new UnibeeApiBeanBulkMetricLimitPlanBindingParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanBulkMetricLimitPlanBindingParam(metricId int64, metricLimit int64) *UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	this := UnibeeApiBeanBulkMetricLimitPlanBindingParam{}
	this.MetricId = metricId
	this.MetricLimit = metricLimit
	return &this
}

// NewUnibeeApiBeanBulkMetricLimitPlanBindingParamWithDefaults instantiates a new UnibeeApiBeanBulkMetricLimitPlanBindingParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanBulkMetricLimitPlanBindingParamWithDefaults() *UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	this := UnibeeApiBeanBulkMetricLimitPlanBindingParam{}
	return &this
}

// GetMetricId returns the MetricId field value
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) GetMetricId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) GetMetricIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) SetMetricId(v int64) {
	o.MetricId = v
}

// GetMetricLimit returns the MetricLimit field value
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) GetMetricLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) GetMetricLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricLimit, true
}

// SetMetricLimit sets field value
func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) SetMetricLimit(v int64) {
	o.MetricLimit = v
}

func (o UnibeeApiBeanBulkMetricLimitPlanBindingParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanBulkMetricLimitPlanBindingParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricId"] = o.MetricId
	toSerialize["metricLimit"] = o.MetricLimit
	return toSerialize, nil
}

func (o *UnibeeApiBeanBulkMetricLimitPlanBindingParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricId",
		"metricLimit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiBeanBulkMetricLimitPlanBindingParam := _UnibeeApiBeanBulkMetricLimitPlanBindingParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiBeanBulkMetricLimitPlanBindingParam)

	if err != nil {
		return err
	}

	*o = UnibeeApiBeanBulkMetricLimitPlanBindingParam(varUnibeeApiBeanBulkMetricLimitPlanBindingParam)

	return err
}

type NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam struct {
	value *UnibeeApiBeanBulkMetricLimitPlanBindingParam
	isSet bool
}

func (v NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) Get() *UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	return v.value
}

func (v *NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) Set(val *UnibeeApiBeanBulkMetricLimitPlanBindingParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanBulkMetricLimitPlanBindingParam(val *UnibeeApiBeanBulkMetricLimitPlanBindingParam) *NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam {
	return &NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanBulkMetricLimitPlanBindingParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


