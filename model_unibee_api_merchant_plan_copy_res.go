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

// checks if the UnibeeApiMerchantPlanCopyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanCopyRes{}

// UnibeeApiMerchantPlanCopyRes struct for UnibeeApiMerchantPlanCopyRes
type UnibeeApiMerchantPlanCopyRes struct {
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanCopyRes instantiates a new UnibeeApiMerchantPlanCopyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanCopyRes() *UnibeeApiMerchantPlanCopyRes {
	this := UnibeeApiMerchantPlanCopyRes{}
	return &this
}

// NewUnibeeApiMerchantPlanCopyResWithDefaults instantiates a new UnibeeApiMerchantPlanCopyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanCopyResWithDefaults() *UnibeeApiMerchantPlanCopyRes {
	this := UnibeeApiMerchantPlanCopyRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanCopyRes) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanCopyRes) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanCopyRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanCopyRes) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanCopyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanCopyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanCopyRes struct {
	value *UnibeeApiMerchantPlanCopyRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanCopyRes) Get() *UnibeeApiMerchantPlanCopyRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanCopyRes) Set(val *UnibeeApiMerchantPlanCopyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanCopyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanCopyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanCopyRes(val *UnibeeApiMerchantPlanCopyRes) *NullableUnibeeApiMerchantPlanCopyRes {
	return &NullableUnibeeApiMerchantPlanCopyRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanCopyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanCopyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


