/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricEditPlanLimitRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricEditPlanLimitRes{}

// UnibeeApiMerchantMetricEditPlanLimitRes struct for UnibeeApiMerchantMetricEditPlanLimitRes
type UnibeeApiMerchantMetricEditPlanLimitRes struct {
	MerchantMetricPlanLimit *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail `json:"merchantMetricPlanLimit,omitempty"`
}

// NewUnibeeApiMerchantMetricEditPlanLimitRes instantiates a new UnibeeApiMerchantMetricEditPlanLimitRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricEditPlanLimitRes() *UnibeeApiMerchantMetricEditPlanLimitRes {
	this := UnibeeApiMerchantMetricEditPlanLimitRes{}
	return &this
}

// NewUnibeeApiMerchantMetricEditPlanLimitResWithDefaults instantiates a new UnibeeApiMerchantMetricEditPlanLimitRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricEditPlanLimitResWithDefaults() *UnibeeApiMerchantMetricEditPlanLimitRes {
	this := UnibeeApiMerchantMetricEditPlanLimitRes{}
	return &this
}

// GetMerchantMetricPlanLimit returns the MerchantMetricPlanLimit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEditPlanLimitRes) GetMerchantMetricPlanLimit() UnibeeApiBeanDetailMerchantMetricPlanLimitDetail {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		var ret UnibeeApiBeanDetailMerchantMetricPlanLimitDetail
		return ret
	}
	return *o.MerchantMetricPlanLimit
}

// GetMerchantMetricPlanLimitOk returns a tuple with the MerchantMetricPlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEditPlanLimitRes) GetMerchantMetricPlanLimitOk() (*UnibeeApiBeanDetailMerchantMetricPlanLimitDetail, bool) {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		return nil, false
	}
	return o.MerchantMetricPlanLimit, true
}

// HasMerchantMetricPlanLimit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEditPlanLimitRes) HasMerchantMetricPlanLimit() bool {
	if o != nil && !IsNil(o.MerchantMetricPlanLimit) {
		return true
	}

	return false
}

// SetMerchantMetricPlanLimit gets a reference to the given UnibeeApiBeanDetailMerchantMetricPlanLimitDetail and assigns it to the MerchantMetricPlanLimit field.
func (o *UnibeeApiMerchantMetricEditPlanLimitRes) SetMerchantMetricPlanLimit(v UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) {
	o.MerchantMetricPlanLimit = &v
}

func (o UnibeeApiMerchantMetricEditPlanLimitRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricEditPlanLimitRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricPlanLimit) {
		toSerialize["merchantMetricPlanLimit"] = o.MerchantMetricPlanLimit
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricEditPlanLimitRes struct {
	value *UnibeeApiMerchantMetricEditPlanLimitRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitRes) Get() *UnibeeApiMerchantMetricEditPlanLimitRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitRes) Set(val *UnibeeApiMerchantMetricEditPlanLimitRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricEditPlanLimitRes(val *UnibeeApiMerchantMetricEditPlanLimitRes) *NullableUnibeeApiMerchantMetricEditPlanLimitRes {
	return &NullableUnibeeApiMerchantMetricEditPlanLimitRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricEditPlanLimitRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricEditPlanLimitRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


