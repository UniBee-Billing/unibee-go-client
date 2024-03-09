/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanUnPublishReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanUnPublishReq{}

// UnibeeApiMerchantPlanUnPublishReq struct for UnibeeApiMerchantPlanUnPublishReq
type UnibeeApiMerchantPlanUnPublishReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanUnPublishReq UnibeeApiMerchantPlanUnPublishReq

// NewUnibeeApiMerchantPlanUnPublishReq instantiates a new UnibeeApiMerchantPlanUnPublishReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanUnPublishReq(planId int64) *UnibeeApiMerchantPlanUnPublishReq {
	this := UnibeeApiMerchantPlanUnPublishReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanUnPublishReqWithDefaults instantiates a new UnibeeApiMerchantPlanUnPublishReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanUnPublishReqWithDefaults() *UnibeeApiMerchantPlanUnPublishReq {
	this := UnibeeApiMerchantPlanUnPublishReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanUnPublishReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanUnPublishReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanUnPublishReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanUnPublishReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanUnPublishReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanUnPublishReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanUnPublishReq := _UnibeeApiMerchantPlanUnPublishReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanUnPublishReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanUnPublishReq(varUnibeeApiMerchantPlanUnPublishReq)

	return err
}

type NullableUnibeeApiMerchantPlanUnPublishReq struct {
	value *UnibeeApiMerchantPlanUnPublishReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanUnPublishReq) Get() *UnibeeApiMerchantPlanUnPublishReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanUnPublishReq) Set(val *UnibeeApiMerchantPlanUnPublishReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanUnPublishReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanUnPublishReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanUnPublishReq(val *UnibeeApiMerchantPlanUnPublishReq) *NullableUnibeeApiMerchantPlanUnPublishReq {
	return &NullableUnibeeApiMerchantPlanUnPublishReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanUnPublishReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanUnPublishReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


