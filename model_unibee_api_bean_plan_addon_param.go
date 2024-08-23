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

// checks if the UnibeeApiBeanPlanAddonParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPlanAddonParam{}

// UnibeeApiBeanPlanAddonParam struct for UnibeeApiBeanPlanAddonParam
type UnibeeApiBeanPlanAddonParam struct {
	// AddonPlanId
	AddonPlanId *int64 `json:"addonPlanId,omitempty"`
	// Quantity，Default 1
	Quantity *int64 `json:"quantity,omitempty"`
}

// NewUnibeeApiBeanPlanAddonParam instantiates a new UnibeeApiBeanPlanAddonParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPlanAddonParam() *UnibeeApiBeanPlanAddonParam {
	this := UnibeeApiBeanPlanAddonParam{}
	return &this
}

// NewUnibeeApiBeanPlanAddonParamWithDefaults instantiates a new UnibeeApiBeanPlanAddonParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPlanAddonParamWithDefaults() *UnibeeApiBeanPlanAddonParam {
	this := UnibeeApiBeanPlanAddonParam{}
	return &this
}

// GetAddonPlanId returns the AddonPlanId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanAddonParam) GetAddonPlanId() int64 {
	if o == nil || IsNil(o.AddonPlanId) {
		var ret int64
		return ret
	}
	return *o.AddonPlanId
}

// GetAddonPlanIdOk returns a tuple with the AddonPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanAddonParam) GetAddonPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AddonPlanId) {
		return nil, false
	}
	return o.AddonPlanId, true
}

// HasAddonPlanId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanAddonParam) HasAddonPlanId() bool {
	if o != nil && !IsNil(o.AddonPlanId) {
		return true
	}

	return false
}

// SetAddonPlanId gets a reference to the given int64 and assigns it to the AddonPlanId field.
func (o *UnibeeApiBeanPlanAddonParam) SetAddonPlanId(v int64) {
	o.AddonPlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanAddonParam) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanAddonParam) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanAddonParam) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanPlanAddonParam) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o UnibeeApiBeanPlanAddonParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPlanAddonParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonPlanId) {
		toSerialize["addonPlanId"] = o.AddonPlanId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPlanAddonParam struct {
	value *UnibeeApiBeanPlanAddonParam
	isSet bool
}

func (v NullableUnibeeApiBeanPlanAddonParam) Get() *UnibeeApiBeanPlanAddonParam {
	return v.value
}

func (v *NullableUnibeeApiBeanPlanAddonParam) Set(val *UnibeeApiBeanPlanAddonParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPlanAddonParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPlanAddonParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPlanAddonParam(val *UnibeeApiBeanPlanAddonParam) *NullableUnibeeApiBeanPlanAddonParam {
	return &NullableUnibeeApiBeanPlanAddonParam{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPlanAddonParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPlanAddonParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


