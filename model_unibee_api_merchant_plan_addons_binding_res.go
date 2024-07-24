/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanAddonsBindingRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanAddonsBindingRes{}

// UnibeeApiMerchantPlanAddonsBindingRes struct for UnibeeApiMerchantPlanAddonsBindingRes
type UnibeeApiMerchantPlanAddonsBindingRes struct {
	Plan *UnibeeApiBeanPlanSimplify `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanAddonsBindingRes instantiates a new UnibeeApiMerchantPlanAddonsBindingRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanAddonsBindingRes() *UnibeeApiMerchantPlanAddonsBindingRes {
	this := UnibeeApiMerchantPlanAddonsBindingRes{}
	return &this
}

// NewUnibeeApiMerchantPlanAddonsBindingResWithDefaults instantiates a new UnibeeApiMerchantPlanAddonsBindingRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanAddonsBindingResWithDefaults() *UnibeeApiMerchantPlanAddonsBindingRes {
	this := UnibeeApiMerchantPlanAddonsBindingRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanAddonsBindingRes) GetPlan() UnibeeApiBeanPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingRes) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlanSimplify and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanAddonsBindingRes) SetPlan(v UnibeeApiBeanPlanSimplify) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanAddonsBindingRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanAddonsBindingRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanAddonsBindingRes struct {
	value *UnibeeApiMerchantPlanAddonsBindingRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingRes) Get() *UnibeeApiMerchantPlanAddonsBindingRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingRes) Set(val *UnibeeApiMerchantPlanAddonsBindingRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanAddonsBindingRes(val *UnibeeApiMerchantPlanAddonsBindingRes) *NullableUnibeeApiMerchantPlanAddonsBindingRes {
	return &NullableUnibeeApiMerchantPlanAddonsBindingRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


