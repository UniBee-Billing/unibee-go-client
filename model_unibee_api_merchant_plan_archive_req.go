/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanArchiveReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanArchiveReq{}

// UnibeeApiMerchantPlanArchiveReq struct for UnibeeApiMerchantPlanArchiveReq
type UnibeeApiMerchantPlanArchiveReq struct {
	// Hard Archive
	HardArchive *bool `json:"hardArchive,omitempty"`
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanArchiveReq UnibeeApiMerchantPlanArchiveReq

// NewUnibeeApiMerchantPlanArchiveReq instantiates a new UnibeeApiMerchantPlanArchiveReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanArchiveReq(planId int64) *UnibeeApiMerchantPlanArchiveReq {
	this := UnibeeApiMerchantPlanArchiveReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanArchiveReqWithDefaults instantiates a new UnibeeApiMerchantPlanArchiveReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanArchiveReqWithDefaults() *UnibeeApiMerchantPlanArchiveReq {
	this := UnibeeApiMerchantPlanArchiveReq{}
	return &this
}

// GetHardArchive returns the HardArchive field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanArchiveReq) GetHardArchive() bool {
	if o == nil || IsNil(o.HardArchive) {
		var ret bool
		return ret
	}
	return *o.HardArchive
}

// GetHardArchiveOk returns a tuple with the HardArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanArchiveReq) GetHardArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.HardArchive) {
		return nil, false
	}
	return o.HardArchive, true
}

// HasHardArchive returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanArchiveReq) HasHardArchive() bool {
	if o != nil && !IsNil(o.HardArchive) {
		return true
	}

	return false
}

// SetHardArchive gets a reference to the given bool and assigns it to the HardArchive field.
func (o *UnibeeApiMerchantPlanArchiveReq) SetHardArchive(v bool) {
	o.HardArchive = &v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanArchiveReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanArchiveReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanArchiveReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanArchiveReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanArchiveReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardArchive) {
		toSerialize["hardArchive"] = o.HardArchive
	}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanArchiveReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanArchiveReq := _UnibeeApiMerchantPlanArchiveReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanArchiveReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanArchiveReq(varUnibeeApiMerchantPlanArchiveReq)

	return err
}

type NullableUnibeeApiMerchantPlanArchiveReq struct {
	value *UnibeeApiMerchantPlanArchiveReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanArchiveReq) Get() *UnibeeApiMerchantPlanArchiveReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanArchiveReq) Set(val *UnibeeApiMerchantPlanArchiveReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanArchiveReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanArchiveReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanArchiveReq(val *UnibeeApiMerchantPlanArchiveReq) *NullableUnibeeApiMerchantPlanArchiveReq {
	return &NullableUnibeeApiMerchantPlanArchiveReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanArchiveReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanArchiveReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


