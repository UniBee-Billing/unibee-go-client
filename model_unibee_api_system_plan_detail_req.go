/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemPlanDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemPlanDetailReq{}

// UnibeeApiSystemPlanDetailReq struct for UnibeeApiSystemPlanDetailReq
type UnibeeApiSystemPlanDetailReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiSystemPlanDetailReq UnibeeApiSystemPlanDetailReq

// NewUnibeeApiSystemPlanDetailReq instantiates a new UnibeeApiSystemPlanDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemPlanDetailReq(planId int64) *UnibeeApiSystemPlanDetailReq {
	this := UnibeeApiSystemPlanDetailReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiSystemPlanDetailReqWithDefaults instantiates a new UnibeeApiSystemPlanDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemPlanDetailReqWithDefaults() *UnibeeApiSystemPlanDetailReq {
	this := UnibeeApiSystemPlanDetailReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiSystemPlanDetailReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemPlanDetailReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiSystemPlanDetailReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiSystemPlanDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemPlanDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiSystemPlanDetailReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiSystemPlanDetailReq := _UnibeeApiSystemPlanDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemPlanDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemPlanDetailReq(varUnibeeApiSystemPlanDetailReq)

	return err
}

type NullableUnibeeApiSystemPlanDetailReq struct {
	value *UnibeeApiSystemPlanDetailReq
	isSet bool
}

func (v NullableUnibeeApiSystemPlanDetailReq) Get() *UnibeeApiSystemPlanDetailReq {
	return v.value
}

func (v *NullableUnibeeApiSystemPlanDetailReq) Set(val *UnibeeApiSystemPlanDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemPlanDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemPlanDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemPlanDetailReq(val *UnibeeApiSystemPlanDetailReq) *NullableUnibeeApiSystemPlanDetailReq {
	return &NullableUnibeeApiSystemPlanDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemPlanDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemPlanDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


