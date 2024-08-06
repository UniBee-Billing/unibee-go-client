/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricNewPlanLimitRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewPlanLimitRes{}

// UnibeeApiMerchantMetricNewPlanLimitRes struct for UnibeeApiMerchantMetricNewPlanLimitRes
type UnibeeApiMerchantMetricNewPlanLimitRes struct {
	MerchantMetricPlanLimit *UnibeeApiBeanMerchantMetricPlanLimit `json:"merchantMetricPlanLimit,omitempty"`
}

// NewUnibeeApiMerchantMetricNewPlanLimitRes instantiates a new UnibeeApiMerchantMetricNewPlanLimitRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewPlanLimitRes() *UnibeeApiMerchantMetricNewPlanLimitRes {
	this := UnibeeApiMerchantMetricNewPlanLimitRes{}
	return &this
}

// NewUnibeeApiMerchantMetricNewPlanLimitResWithDefaults instantiates a new UnibeeApiMerchantMetricNewPlanLimitRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewPlanLimitResWithDefaults() *UnibeeApiMerchantMetricNewPlanLimitRes {
	this := UnibeeApiMerchantMetricNewPlanLimitRes{}
	return &this
}

// GetMerchantMetricPlanLimit returns the MerchantMetricPlanLimit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewPlanLimitRes) GetMerchantMetricPlanLimit() UnibeeApiBeanMerchantMetricPlanLimit {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		var ret UnibeeApiBeanMerchantMetricPlanLimit
		return ret
	}
	return *o.MerchantMetricPlanLimit
}

// GetMerchantMetricPlanLimitOk returns a tuple with the MerchantMetricPlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewPlanLimitRes) GetMerchantMetricPlanLimitOk() (*UnibeeApiBeanMerchantMetricPlanLimit, bool) {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		return nil, false
	}
	return o.MerchantMetricPlanLimit, true
}

// HasMerchantMetricPlanLimit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewPlanLimitRes) HasMerchantMetricPlanLimit() bool {
	if o != nil && !IsNil(o.MerchantMetricPlanLimit) {
		return true
	}

	return false
}

// SetMerchantMetricPlanLimit gets a reference to the given UnibeeApiBeanMerchantMetricPlanLimit and assigns it to the MerchantMetricPlanLimit field.
func (o *UnibeeApiMerchantMetricNewPlanLimitRes) SetMerchantMetricPlanLimit(v UnibeeApiBeanMerchantMetricPlanLimit) {
	o.MerchantMetricPlanLimit = &v
}

func (o UnibeeApiMerchantMetricNewPlanLimitRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewPlanLimitRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricPlanLimit) {
		toSerialize["merchantMetricPlanLimit"] = o.MerchantMetricPlanLimit
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricNewPlanLimitRes struct {
	value *UnibeeApiMerchantMetricNewPlanLimitRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitRes) Get() *UnibeeApiMerchantMetricNewPlanLimitRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitRes) Set(val *UnibeeApiMerchantMetricNewPlanLimitRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewPlanLimitRes(val *UnibeeApiMerchantMetricNewPlanLimitRes) *NullableUnibeeApiMerchantMetricNewPlanLimitRes {
	return &NullableUnibeeApiMerchantMetricNewPlanLimitRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewPlanLimitRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewPlanLimitRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


