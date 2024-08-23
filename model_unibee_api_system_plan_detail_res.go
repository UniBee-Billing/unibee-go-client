/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiSystemPlanDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemPlanDetailRes{}

// UnibeeApiSystemPlanDetailRes struct for UnibeeApiSystemPlanDetailRes
type UnibeeApiSystemPlanDetailRes struct {
	Plan *UnibeeApiBeanDetailPlanDetail `json:"plan,omitempty"`
}

// NewUnibeeApiSystemPlanDetailRes instantiates a new UnibeeApiSystemPlanDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemPlanDetailRes() *UnibeeApiSystemPlanDetailRes {
	this := UnibeeApiSystemPlanDetailRes{}
	return &this
}

// NewUnibeeApiSystemPlanDetailResWithDefaults instantiates a new UnibeeApiSystemPlanDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemPlanDetailResWithDefaults() *UnibeeApiSystemPlanDetailRes {
	this := UnibeeApiSystemPlanDetailRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiSystemPlanDetailRes) GetPlan() UnibeeApiBeanDetailPlanDetail {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanDetailPlanDetail
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemPlanDetailRes) GetPlanOk() (*UnibeeApiBeanDetailPlanDetail, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiSystemPlanDetailRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanDetailPlanDetail and assigns it to the Plan field.
func (o *UnibeeApiSystemPlanDetailRes) SetPlan(v UnibeeApiBeanDetailPlanDetail) {
	o.Plan = &v
}

func (o UnibeeApiSystemPlanDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemPlanDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiSystemPlanDetailRes struct {
	value *UnibeeApiSystemPlanDetailRes
	isSet bool
}

func (v NullableUnibeeApiSystemPlanDetailRes) Get() *UnibeeApiSystemPlanDetailRes {
	return v.value
}

func (v *NullableUnibeeApiSystemPlanDetailRes) Set(val *UnibeeApiSystemPlanDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemPlanDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemPlanDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemPlanDetailRes(val *UnibeeApiSystemPlanDetailRes) *NullableUnibeeApiSystemPlanDetailRes {
	return &NullableUnibeeApiSystemPlanDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemPlanDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemPlanDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


