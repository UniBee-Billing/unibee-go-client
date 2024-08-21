/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanCopyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanCopyReq{}

// UnibeeApiMerchantPlanCopyReq struct for UnibeeApiMerchantPlanCopyReq
type UnibeeApiMerchantPlanCopyReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanCopyReq UnibeeApiMerchantPlanCopyReq

// NewUnibeeApiMerchantPlanCopyReq instantiates a new UnibeeApiMerchantPlanCopyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanCopyReq(planId int64) *UnibeeApiMerchantPlanCopyReq {
	this := UnibeeApiMerchantPlanCopyReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanCopyReqWithDefaults instantiates a new UnibeeApiMerchantPlanCopyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanCopyReqWithDefaults() *UnibeeApiMerchantPlanCopyReq {
	this := UnibeeApiMerchantPlanCopyReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanCopyReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanCopyReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanCopyReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanCopyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanCopyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanCopyReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanCopyReq := _UnibeeApiMerchantPlanCopyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanCopyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanCopyReq(varUnibeeApiMerchantPlanCopyReq)

	return err
}

type NullableUnibeeApiMerchantPlanCopyReq struct {
	value *UnibeeApiMerchantPlanCopyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanCopyReq) Get() *UnibeeApiMerchantPlanCopyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanCopyReq) Set(val *UnibeeApiMerchantPlanCopyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanCopyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanCopyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanCopyReq(val *UnibeeApiMerchantPlanCopyReq) *NullableUnibeeApiMerchantPlanCopyReq {
	return &NullableUnibeeApiMerchantPlanCopyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanCopyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanCopyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


