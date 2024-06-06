/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanNewRes{}

// UnibeeApiMerchantPlanNewRes struct for UnibeeApiMerchantPlanNewRes
type UnibeeApiMerchantPlanNewRes struct {
	Plan *UnibeeApiBeanPlanSimplify `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanNewRes instantiates a new UnibeeApiMerchantPlanNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanNewRes() *UnibeeApiMerchantPlanNewRes {
	this := UnibeeApiMerchantPlanNewRes{}
	return &this
}

// NewUnibeeApiMerchantPlanNewResWithDefaults instantiates a new UnibeeApiMerchantPlanNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanNewResWithDefaults() *UnibeeApiMerchantPlanNewRes {
	this := UnibeeApiMerchantPlanNewRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewRes) GetPlan() UnibeeApiBeanPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewRes) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlanSimplify and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanNewRes) SetPlan(v UnibeeApiBeanPlanSimplify) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanNewRes struct {
	value *UnibeeApiMerchantPlanNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanNewRes) Get() *UnibeeApiMerchantPlanNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanNewRes) Set(val *UnibeeApiMerchantPlanNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanNewRes(val *UnibeeApiMerchantPlanNewRes) *NullableUnibeeApiMerchantPlanNewRes {
	return &NullableUnibeeApiMerchantPlanNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


