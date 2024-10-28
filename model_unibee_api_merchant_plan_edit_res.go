/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanEditRes{}

// UnibeeApiMerchantPlanEditRes struct for UnibeeApiMerchantPlanEditRes
type UnibeeApiMerchantPlanEditRes struct {
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanEditRes instantiates a new UnibeeApiMerchantPlanEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanEditRes() *UnibeeApiMerchantPlanEditRes {
	this := UnibeeApiMerchantPlanEditRes{}
	return &this
}

// NewUnibeeApiMerchantPlanEditResWithDefaults instantiates a new UnibeeApiMerchantPlanEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanEditResWithDefaults() *UnibeeApiMerchantPlanEditRes {
	this := UnibeeApiMerchantPlanEditRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditRes) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditRes) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanEditRes) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanEditRes struct {
	value *UnibeeApiMerchantPlanEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanEditRes) Get() *UnibeeApiMerchantPlanEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanEditRes) Set(val *UnibeeApiMerchantPlanEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanEditRes(val *UnibeeApiMerchantPlanEditRes) *NullableUnibeeApiMerchantPlanEditRes {
	return &NullableUnibeeApiMerchantPlanEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


