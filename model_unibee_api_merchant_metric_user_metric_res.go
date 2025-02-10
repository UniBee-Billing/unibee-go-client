/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricUserMetricRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricUserMetricRes{}

// UnibeeApiMerchantMetricUserMetricRes struct for UnibeeApiMerchantMetricUserMetricRes
type UnibeeApiMerchantMetricUserMetricRes struct {
	UserMetric *UnibeeApiMerchantMetricUserMetric `json:"userMetric,omitempty"`
}

// NewUnibeeApiMerchantMetricUserMetricRes instantiates a new UnibeeApiMerchantMetricUserMetricRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricUserMetricRes() *UnibeeApiMerchantMetricUserMetricRes {
	this := UnibeeApiMerchantMetricUserMetricRes{}
	return &this
}

// NewUnibeeApiMerchantMetricUserMetricResWithDefaults instantiates a new UnibeeApiMerchantMetricUserMetricRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricUserMetricResWithDefaults() *UnibeeApiMerchantMetricUserMetricRes {
	this := UnibeeApiMerchantMetricUserMetricRes{}
	return &this
}

// GetUserMetric returns the UserMetric field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetricRes) GetUserMetric() UnibeeApiMerchantMetricUserMetric {
	if o == nil || IsNil(o.UserMetric) {
		var ret UnibeeApiMerchantMetricUserMetric
		return ret
	}
	return *o.UserMetric
}

// GetUserMetricOk returns a tuple with the UserMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetricRes) GetUserMetricOk() (*UnibeeApiMerchantMetricUserMetric, bool) {
	if o == nil || IsNil(o.UserMetric) {
		return nil, false
	}
	return o.UserMetric, true
}

// HasUserMetric returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetricRes) HasUserMetric() bool {
	if o != nil && !IsNil(o.UserMetric) {
		return true
	}

	return false
}

// SetUserMetric gets a reference to the given UnibeeApiMerchantMetricUserMetric and assigns it to the UserMetric field.
func (o *UnibeeApiMerchantMetricUserMetricRes) SetUserMetric(v UnibeeApiMerchantMetricUserMetric) {
	o.UserMetric = &v
}

func (o UnibeeApiMerchantMetricUserMetricRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricUserMetricRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserMetric) {
		toSerialize["userMetric"] = o.UserMetric
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricUserMetricRes struct {
	value *UnibeeApiMerchantMetricUserMetricRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricUserMetricRes) Get() *UnibeeApiMerchantMetricUserMetricRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricUserMetricRes) Set(val *UnibeeApiMerchantMetricUserMetricRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricUserMetricRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricUserMetricRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricUserMetricRes(val *UnibeeApiMerchantMetricUserMetricRes) *NullableUnibeeApiMerchantMetricUserMetricRes {
	return &NullableUnibeeApiMerchantMetricUserMetricRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricUserMetricRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricUserMetricRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


