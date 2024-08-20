/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanPublishReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanPublishReq{}

// UnibeeApiMerchantPlanPublishReq Publish plan，a plan will display at user portal when its published
type UnibeeApiMerchantPlanPublishReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanPublishReq UnibeeApiMerchantPlanPublishReq

// NewUnibeeApiMerchantPlanPublishReq instantiates a new UnibeeApiMerchantPlanPublishReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanPublishReq(planId int64) *UnibeeApiMerchantPlanPublishReq {
	this := UnibeeApiMerchantPlanPublishReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanPublishReqWithDefaults instantiates a new UnibeeApiMerchantPlanPublishReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanPublishReqWithDefaults() *UnibeeApiMerchantPlanPublishReq {
	this := UnibeeApiMerchantPlanPublishReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanPublishReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanPublishReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanPublishReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanPublishReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanPublishReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanPublishReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanPublishReq := _UnibeeApiMerchantPlanPublishReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanPublishReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanPublishReq(varUnibeeApiMerchantPlanPublishReq)

	return err
}

type NullableUnibeeApiMerchantPlanPublishReq struct {
	value *UnibeeApiMerchantPlanPublishReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanPublishReq) Get() *UnibeeApiMerchantPlanPublishReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanPublishReq) Set(val *UnibeeApiMerchantPlanPublishReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanPublishReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanPublishReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanPublishReq(val *UnibeeApiMerchantPlanPublishReq) *NullableUnibeeApiMerchantPlanPublishReq {
	return &NullableUnibeeApiMerchantPlanPublishReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanPublishReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanPublishReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


