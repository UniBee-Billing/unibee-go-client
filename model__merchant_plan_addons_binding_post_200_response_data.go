/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPlanAddonsBindingPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPlanAddonsBindingPost200ResponseData{}

// MerchantPlanAddonsBindingPost200ResponseData struct for MerchantPlanAddonsBindingPost200ResponseData
type MerchantPlanAddonsBindingPost200ResponseData struct {
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
}

// NewMerchantPlanAddonsBindingPost200ResponseData instantiates a new MerchantPlanAddonsBindingPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPlanAddonsBindingPost200ResponseData() *MerchantPlanAddonsBindingPost200ResponseData {
	this := MerchantPlanAddonsBindingPost200ResponseData{}
	return &this
}

// NewMerchantPlanAddonsBindingPost200ResponseDataWithDefaults instantiates a new MerchantPlanAddonsBindingPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPlanAddonsBindingPost200ResponseDataWithDefaults() *MerchantPlanAddonsBindingPost200ResponseData {
	this := MerchantPlanAddonsBindingPost200ResponseData{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantPlanAddonsBindingPost200ResponseData) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPlanAddonsBindingPost200ResponseData) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantPlanAddonsBindingPost200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *MerchantPlanAddonsBindingPost200ResponseData) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

func (o MerchantPlanAddonsBindingPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPlanAddonsBindingPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableMerchantPlanAddonsBindingPost200ResponseData struct {
	value *MerchantPlanAddonsBindingPost200ResponseData
	isSet bool
}

func (v NullableMerchantPlanAddonsBindingPost200ResponseData) Get() *MerchantPlanAddonsBindingPost200ResponseData {
	return v.value
}

func (v *NullableMerchantPlanAddonsBindingPost200ResponseData) Set(val *MerchantPlanAddonsBindingPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPlanAddonsBindingPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPlanAddonsBindingPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPlanAddonsBindingPost200ResponseData(val *MerchantPlanAddonsBindingPost200ResponseData) *NullableMerchantPlanAddonsBindingPost200ResponseData {
	return &NullableMerchantPlanAddonsBindingPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPlanAddonsBindingPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPlanAddonsBindingPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


