/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricDeletePlanLimitReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDeletePlanLimitReq{}

// UnibeeApiMerchantMetricDeletePlanLimitReq struct for UnibeeApiMerchantMetricDeletePlanLimitReq
type UnibeeApiMerchantMetricDeletePlanLimitReq struct {
	// MetricPlanLimitId
	MetricPlanLimitId int64 `json:"metricPlanLimitId"`
}

type _UnibeeApiMerchantMetricDeletePlanLimitReq UnibeeApiMerchantMetricDeletePlanLimitReq

// NewUnibeeApiMerchantMetricDeletePlanLimitReq instantiates a new UnibeeApiMerchantMetricDeletePlanLimitReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDeletePlanLimitReq(metricPlanLimitId int64) *UnibeeApiMerchantMetricDeletePlanLimitReq {
	this := UnibeeApiMerchantMetricDeletePlanLimitReq{}
	this.MetricPlanLimitId = metricPlanLimitId
	return &this
}

// NewUnibeeApiMerchantMetricDeletePlanLimitReqWithDefaults instantiates a new UnibeeApiMerchantMetricDeletePlanLimitReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDeletePlanLimitReqWithDefaults() *UnibeeApiMerchantMetricDeletePlanLimitReq {
	this := UnibeeApiMerchantMetricDeletePlanLimitReq{}
	return &this
}

// GetMetricPlanLimitId returns the MetricPlanLimitId field value
func (o *UnibeeApiMerchantMetricDeletePlanLimitReq) GetMetricPlanLimitId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricPlanLimitId
}

// GetMetricPlanLimitIdOk returns a tuple with the MetricPlanLimitId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeletePlanLimitReq) GetMetricPlanLimitIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricPlanLimitId, true
}

// SetMetricPlanLimitId sets field value
func (o *UnibeeApiMerchantMetricDeletePlanLimitReq) SetMetricPlanLimitId(v int64) {
	o.MetricPlanLimitId = v
}

func (o UnibeeApiMerchantMetricDeletePlanLimitReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDeletePlanLimitReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricPlanLimitId"] = o.MetricPlanLimitId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricDeletePlanLimitReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUnibeeApiMerchantMetricDeletePlanLimitReq := _UnibeeApiMerchantMetricDeletePlanLimitReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricDeletePlanLimitReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricDeletePlanLimitReq(varUnibeeApiMerchantMetricDeletePlanLimitReq)

	return err
}

type NullableUnibeeApiMerchantMetricDeletePlanLimitReq struct {
	value *UnibeeApiMerchantMetricDeletePlanLimitReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitReq) Get() *UnibeeApiMerchantMetricDeletePlanLimitReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitReq) Set(val *UnibeeApiMerchantMetricDeletePlanLimitReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDeletePlanLimitReq(val *UnibeeApiMerchantMetricDeletePlanLimitReq) *NullableUnibeeApiMerchantMetricDeletePlanLimitReq {
	return &NullableUnibeeApiMerchantMetricDeletePlanLimitReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


