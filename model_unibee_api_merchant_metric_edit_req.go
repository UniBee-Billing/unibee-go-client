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

// checks if the UnibeeApiMerchantMetricEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricEditReq{}

// UnibeeApiMerchantMetricEditReq struct for UnibeeApiMerchantMetricEditReq
type UnibeeApiMerchantMetricEditReq struct {
	// MetricDescription
	MetricDescription *string `json:"metricDescription,omitempty"`
	// MetricId
	MetricId int64 `json:"metricId"`
	// MetricName
	MetricName string `json:"metricName"`
}

type _UnibeeApiMerchantMetricEditReq UnibeeApiMerchantMetricEditReq

// NewUnibeeApiMerchantMetricEditReq instantiates a new UnibeeApiMerchantMetricEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricEditReq(metricId int64, metricName string) *UnibeeApiMerchantMetricEditReq {
	this := UnibeeApiMerchantMetricEditReq{}
	this.MetricId = metricId
	this.MetricName = metricName
	return &this
}

// NewUnibeeApiMerchantMetricEditReqWithDefaults instantiates a new UnibeeApiMerchantMetricEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricEditReqWithDefaults() *UnibeeApiMerchantMetricEditReq {
	this := UnibeeApiMerchantMetricEditReq{}
	return &this
}

// GetMetricDescription returns the MetricDescription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEditReq) GetMetricDescription() string {
	if o == nil || IsNil(o.MetricDescription) {
		var ret string
		return ret
	}
	return *o.MetricDescription
}

// GetMetricDescriptionOk returns a tuple with the MetricDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditReq) GetMetricDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDescription) {
		return nil, false
	}
	return o.MetricDescription, true
}

// HasMetricDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEditReq) HasMetricDescription() bool {
	if o != nil && !IsNil(o.MetricDescription) {
		return true
	}

	return false
}

// SetMetricDescription gets a reference to the given string and assigns it to the MetricDescription field.
func (o *UnibeeApiMerchantMetricEditReq) SetMetricDescription(v string) {
	o.MetricDescription = &v
}

// GetMetricId returns the MetricId field value
func (o *UnibeeApiMerchantMetricEditReq) GetMetricId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditReq) GetMetricIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *UnibeeApiMerchantMetricEditReq) SetMetricId(v int64) {
	o.MetricId = v
}

// GetMetricName returns the MetricName field value
func (o *UnibeeApiMerchantMetricEditReq) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditReq) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *UnibeeApiMerchantMetricEditReq) SetMetricName(v string) {
	o.MetricName = v
}

func (o UnibeeApiMerchantMetricEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricDescription) {
		toSerialize["metricDescription"] = o.MetricDescription
	}
	toSerialize["metricId"] = o.MetricId
	toSerialize["metricName"] = o.MetricName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricId",
		"metricName",
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

	varUnibeeApiMerchantMetricEditReq := _UnibeeApiMerchantMetricEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricEditReq(varUnibeeApiMerchantMetricEditReq)

	return err
}

type NullableUnibeeApiMerchantMetricEditReq struct {
	value *UnibeeApiMerchantMetricEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricEditReq) Get() *UnibeeApiMerchantMetricEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricEditReq) Set(val *UnibeeApiMerchantMetricEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricEditReq(val *UnibeeApiMerchantMetricEditReq) *NullableUnibeeApiMerchantMetricEditReq {
	return &NullableUnibeeApiMerchantMetricEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


