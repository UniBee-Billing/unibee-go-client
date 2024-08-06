/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPlanAddonDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPlanAddonDetail{}

// UnibeeApiBeanPlanAddonDetail struct for UnibeeApiBeanPlanAddonDetail
type UnibeeApiBeanPlanAddonDetail struct {
	AddonPlan *UnibeeApiBeanPlan `json:"addonPlan,omitempty"`
	// Quantity
	Quantity *int64 `json:"quantity,omitempty"`
}

// NewUnibeeApiBeanPlanAddonDetail instantiates a new UnibeeApiBeanPlanAddonDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPlanAddonDetail() *UnibeeApiBeanPlanAddonDetail {
	this := UnibeeApiBeanPlanAddonDetail{}
	return &this
}

// NewUnibeeApiBeanPlanAddonDetailWithDefaults instantiates a new UnibeeApiBeanPlanAddonDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPlanAddonDetailWithDefaults() *UnibeeApiBeanPlanAddonDetail {
	this := UnibeeApiBeanPlanAddonDetail{}
	return &this
}

// GetAddonPlan returns the AddonPlan field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanAddonDetail) GetAddonPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.AddonPlan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.AddonPlan
}

// GetAddonPlanOk returns a tuple with the AddonPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanAddonDetail) GetAddonPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.AddonPlan) {
		return nil, false
	}
	return o.AddonPlan, true
}

// HasAddonPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanAddonDetail) HasAddonPlan() bool {
	if o != nil && !IsNil(o.AddonPlan) {
		return true
	}

	return false
}

// SetAddonPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the AddonPlan field.
func (o *UnibeeApiBeanPlanAddonDetail) SetAddonPlan(v UnibeeApiBeanPlan) {
	o.AddonPlan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanAddonDetail) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanAddonDetail) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanAddonDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanPlanAddonDetail) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o UnibeeApiBeanPlanAddonDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPlanAddonDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonPlan) {
		toSerialize["addonPlan"] = o.AddonPlan
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPlanAddonDetail struct {
	value *UnibeeApiBeanPlanAddonDetail
	isSet bool
}

func (v NullableUnibeeApiBeanPlanAddonDetail) Get() *UnibeeApiBeanPlanAddonDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanPlanAddonDetail) Set(val *UnibeeApiBeanPlanAddonDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPlanAddonDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPlanAddonDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPlanAddonDetail(val *UnibeeApiBeanPlanAddonDetail) *NullableUnibeeApiBeanPlanAddonDetail {
	return &NullableUnibeeApiBeanPlanAddonDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPlanAddonDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPlanAddonDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


