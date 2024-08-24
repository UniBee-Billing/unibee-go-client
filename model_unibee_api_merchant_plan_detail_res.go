/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanDetailRes{}

// UnibeeApiMerchantPlanDetailRes struct for UnibeeApiMerchantPlanDetailRes
type UnibeeApiMerchantPlanDetailRes struct {
	Plan *UnibeeApiBeanDetailPlanDetail `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanDetailRes instantiates a new UnibeeApiMerchantPlanDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanDetailRes() *UnibeeApiMerchantPlanDetailRes {
	this := UnibeeApiMerchantPlanDetailRes{}
	return &this
}

// NewUnibeeApiMerchantPlanDetailResWithDefaults instantiates a new UnibeeApiMerchantPlanDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanDetailResWithDefaults() *UnibeeApiMerchantPlanDetailRes {
	this := UnibeeApiMerchantPlanDetailRes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanDetailRes) GetPlan() UnibeeApiBeanDetailPlanDetail {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanDetailPlanDetail
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanDetailRes) GetPlanOk() (*UnibeeApiBeanDetailPlanDetail, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanDetailRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanDetailPlanDetail and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanDetailRes) SetPlan(v UnibeeApiBeanDetailPlanDetail) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanDetailRes struct {
	value *UnibeeApiMerchantPlanDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanDetailRes) Get() *UnibeeApiMerchantPlanDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanDetailRes) Set(val *UnibeeApiMerchantPlanDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanDetailRes(val *UnibeeApiMerchantPlanDetailRes) *NullableUnibeeApiMerchantPlanDetailRes {
	return &NullableUnibeeApiMerchantPlanDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


