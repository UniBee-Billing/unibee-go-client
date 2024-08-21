/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricNewPlanLimitReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewPlanLimitReq{}

// UnibeeApiMerchantMetricNewPlanLimitReq struct for UnibeeApiMerchantMetricNewPlanLimitReq
type UnibeeApiMerchantMetricNewPlanLimitReq struct {
	// MetricId
	MetricId int64 `json:"metricId"`
	// MetricLimit
	MetricLimit int64 `json:"metricLimit"`
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantMetricNewPlanLimitReq UnibeeApiMerchantMetricNewPlanLimitReq

// NewUnibeeApiMerchantMetricNewPlanLimitReq instantiates a new UnibeeApiMerchantMetricNewPlanLimitReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewPlanLimitReq(metricId int64, metricLimit int64, planId int64) *UnibeeApiMerchantMetricNewPlanLimitReq {
	this := UnibeeApiMerchantMetricNewPlanLimitReq{}
	this.MetricId = metricId
	this.MetricLimit = metricLimit
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantMetricNewPlanLimitReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewPlanLimitReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewPlanLimitReqWithDefaults() *UnibeeApiMerchantMetricNewPlanLimitReq {
	this := UnibeeApiMerchantMetricNewPlanLimitReq{}
	return &this
}

// GetMetricId returns the MetricId field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetMetricId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetMetricIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) SetMetricId(v int64) {
	o.MetricId = v
}

// GetMetricLimit returns the MetricLimit field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetMetricLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetMetricLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricLimit, true
}

// SetMetricLimit sets field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) SetMetricLimit(v int64) {
	o.MetricLimit = v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantMetricNewPlanLimitReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantMetricNewPlanLimitReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewPlanLimitReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricId"] = o.MetricId
	toSerialize["metricLimit"] = o.MetricLimit
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricNewPlanLimitReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricId",
		"metricLimit",
		"planId",
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

	varUnibeeApiMerchantMetricNewPlanLimitReq := _UnibeeApiMerchantMetricNewPlanLimitReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricNewPlanLimitReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricNewPlanLimitReq(varUnibeeApiMerchantMetricNewPlanLimitReq)

	return err
}

type NullableUnibeeApiMerchantMetricNewPlanLimitReq struct {
	value *UnibeeApiMerchantMetricNewPlanLimitReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitReq) Get() *UnibeeApiMerchantMetricNewPlanLimitReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitReq) Set(val *UnibeeApiMerchantMetricNewPlanLimitReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewPlanLimitReq(val *UnibeeApiMerchantMetricNewPlanLimitReq) *NullableUnibeeApiMerchantMetricNewPlanLimitReq {
	return &NullableUnibeeApiMerchantMetricNewPlanLimitReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


