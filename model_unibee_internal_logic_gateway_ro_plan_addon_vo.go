/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoPlanAddonVo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoPlanAddonVo{}

// UnibeeInternalLogicGatewayRoPlanAddonVo struct for UnibeeInternalLogicGatewayRoPlanAddonVo
type UnibeeInternalLogicGatewayRoPlanAddonVo struct {
	AddonPlan *UnibeeInternalLogicGatewayRoPlanSimplify `json:"addonPlan,omitempty"`
	// Quantity
	Quantity *int64 `json:"quantity,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoPlanAddonVo instantiates a new UnibeeInternalLogicGatewayRoPlanAddonVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoPlanAddonVo() *UnibeeInternalLogicGatewayRoPlanAddonVo {
	this := UnibeeInternalLogicGatewayRoPlanAddonVo{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoPlanAddonVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoPlanAddonVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoPlanAddonVoWithDefaults() *UnibeeInternalLogicGatewayRoPlanAddonVo {
	this := UnibeeInternalLogicGatewayRoPlanAddonVo{}
	return &this
}

// GetAddonPlan returns the AddonPlan field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) GetAddonPlan() UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.AddonPlan) {
		var ret UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return *o.AddonPlan
}

// GetAddonPlanOk returns a tuple with the AddonPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) GetAddonPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.AddonPlan) {
		return nil, false
	}
	return o.AddonPlan, true
}

// HasAddonPlan returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) HasAddonPlan() bool {
	if o != nil && !IsNil(o.AddonPlan) {
		return true
	}

	return false
}

// SetAddonPlan gets a reference to the given UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the AddonPlan field.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) SetAddonPlan(v UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.AddonPlan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeInternalLogicGatewayRoPlanAddonVo) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o UnibeeInternalLogicGatewayRoPlanAddonVo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoPlanAddonVo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonPlan) {
		toSerialize["addonPlan"] = o.AddonPlan
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoPlanAddonVo struct {
	value *UnibeeInternalLogicGatewayRoPlanAddonVo
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoPlanAddonVo) Get() *UnibeeInternalLogicGatewayRoPlanAddonVo {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanAddonVo) Set(val *UnibeeInternalLogicGatewayRoPlanAddonVo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoPlanAddonVo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanAddonVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoPlanAddonVo(val *UnibeeInternalLogicGatewayRoPlanAddonVo) *NullableUnibeeInternalLogicGatewayRoPlanAddonVo {
	return &NullableUnibeeInternalLogicGatewayRoPlanAddonVo{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoPlanAddonVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanAddonVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


