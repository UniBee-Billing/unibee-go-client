/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDeleteReq{}

// UnibeeApiMerchantMetricDeleteReq struct for UnibeeApiMerchantMetricDeleteReq
type UnibeeApiMerchantMetricDeleteReq struct {
	// MetricId
	MetricId int64 `json:"metricId"`
}

type _UnibeeApiMerchantMetricDeleteReq UnibeeApiMerchantMetricDeleteReq

// NewUnibeeApiMerchantMetricDeleteReq instantiates a new UnibeeApiMerchantMetricDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDeleteReq(metricId int64) *UnibeeApiMerchantMetricDeleteReq {
	this := UnibeeApiMerchantMetricDeleteReq{}
	this.MetricId = metricId
	return &this
}

// NewUnibeeApiMerchantMetricDeleteReqWithDefaults instantiates a new UnibeeApiMerchantMetricDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDeleteReqWithDefaults() *UnibeeApiMerchantMetricDeleteReq {
	this := UnibeeApiMerchantMetricDeleteReq{}
	return &this
}

// GetMetricId returns the MetricId field value
func (o *UnibeeApiMerchantMetricDeleteReq) GetMetricId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeleteReq) GetMetricIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *UnibeeApiMerchantMetricDeleteReq) SetMetricId(v int64) {
	o.MetricId = v
}

func (o UnibeeApiMerchantMetricDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricId"] = o.MetricId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricDeleteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricId",
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

	varUnibeeApiMerchantMetricDeleteReq := _UnibeeApiMerchantMetricDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricDeleteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricDeleteReq(varUnibeeApiMerchantMetricDeleteReq)

	return err
}

type NullableUnibeeApiMerchantMetricDeleteReq struct {
	value *UnibeeApiMerchantMetricDeleteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDeleteReq) Get() *UnibeeApiMerchantMetricDeleteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDeleteReq) Set(val *UnibeeApiMerchantMetricDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDeleteReq(val *UnibeeApiMerchantMetricDeleteReq) *NullableUnibeeApiMerchantMetricDeleteReq {
	return &NullableUnibeeApiMerchantMetricDeleteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


