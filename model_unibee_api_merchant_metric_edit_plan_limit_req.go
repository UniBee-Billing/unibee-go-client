/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricEditPlanLimitReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricEditPlanLimitReq{}

// UnibeeApiMerchantMetricEditPlanLimitReq struct for UnibeeApiMerchantMetricEditPlanLimitReq
type UnibeeApiMerchantMetricEditPlanLimitReq struct {
	// MetricLimit
	MetricLimit int64 `json:"metricLimit"`
	// MetricPlanLimitId
	MetricPlanLimitId int64 `json:"metricPlanLimitId"`
}

type _UnibeeApiMerchantMetricEditPlanLimitReq UnibeeApiMerchantMetricEditPlanLimitReq

// NewUnibeeApiMerchantMetricEditPlanLimitReq instantiates a new UnibeeApiMerchantMetricEditPlanLimitReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricEditPlanLimitReq(metricLimit int64, metricPlanLimitId int64) *UnibeeApiMerchantMetricEditPlanLimitReq {
	this := UnibeeApiMerchantMetricEditPlanLimitReq{}
	this.MetricLimit = metricLimit
	this.MetricPlanLimitId = metricPlanLimitId
	return &this
}

// NewUnibeeApiMerchantMetricEditPlanLimitReqWithDefaults instantiates a new UnibeeApiMerchantMetricEditPlanLimitReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricEditPlanLimitReqWithDefaults() *UnibeeApiMerchantMetricEditPlanLimitReq {
	this := UnibeeApiMerchantMetricEditPlanLimitReq{}
	return &this
}

// GetMetricLimit returns the MetricLimit field value
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) GetMetricLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) GetMetricLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricLimit, true
}

// SetMetricLimit sets field value
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) SetMetricLimit(v int64) {
	o.MetricLimit = v
}

// GetMetricPlanLimitId returns the MetricPlanLimitId field value
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) GetMetricPlanLimitId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricPlanLimitId
}

// GetMetricPlanLimitIdOk returns a tuple with the MetricPlanLimitId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) GetMetricPlanLimitIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricPlanLimitId, true
}

// SetMetricPlanLimitId sets field value
func (o *UnibeeApiMerchantMetricEditPlanLimitReq) SetMetricPlanLimitId(v int64) {
	o.MetricPlanLimitId = v
}

func (o UnibeeApiMerchantMetricEditPlanLimitReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricEditPlanLimitReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricLimit"] = o.MetricLimit
	toSerialize["metricPlanLimitId"] = o.MetricPlanLimitId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricEditPlanLimitReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricLimit",
		"metricPlanLimitId",
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

	varUnibeeApiMerchantMetricEditPlanLimitReq := _UnibeeApiMerchantMetricEditPlanLimitReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricEditPlanLimitReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricEditPlanLimitReq(varUnibeeApiMerchantMetricEditPlanLimitReq)

	return err
}

type NullableUnibeeApiMerchantMetricEditPlanLimitReq struct {
	value *UnibeeApiMerchantMetricEditPlanLimitReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitReq) Get() *UnibeeApiMerchantMetricEditPlanLimitReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitReq) Set(val *UnibeeApiMerchantMetricEditPlanLimitReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricEditPlanLimitReq(val *UnibeeApiMerchantMetricEditPlanLimitReq) *NullableUnibeeApiMerchantMetricEditPlanLimitReq {
	return &NullableUnibeeApiMerchantMetricEditPlanLimitReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


