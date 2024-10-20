/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanExpireReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanExpireReq{}

// UnibeeApiMerchantPlanExpireReq struct for UnibeeApiMerchantPlanExpireReq
type UnibeeApiMerchantPlanExpireReq struct {
	// Code From Email
	EmailCode int64 `json:"emailCode"`
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanExpireReq UnibeeApiMerchantPlanExpireReq

// NewUnibeeApiMerchantPlanExpireReq instantiates a new UnibeeApiMerchantPlanExpireReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanExpireReq(emailCode int64, planId int64) *UnibeeApiMerchantPlanExpireReq {
	this := UnibeeApiMerchantPlanExpireReq{}
	this.EmailCode = emailCode
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanExpireReqWithDefaults instantiates a new UnibeeApiMerchantPlanExpireReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanExpireReqWithDefaults() *UnibeeApiMerchantPlanExpireReq {
	this := UnibeeApiMerchantPlanExpireReq{}
	return &this
}

// GetEmailCode returns the EmailCode field value
func (o *UnibeeApiMerchantPlanExpireReq) GetEmailCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EmailCode
}

// GetEmailCodeOk returns a tuple with the EmailCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanExpireReq) GetEmailCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailCode, true
}

// SetEmailCode sets field value
func (o *UnibeeApiMerchantPlanExpireReq) SetEmailCode(v int64) {
	o.EmailCode = v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanExpireReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanExpireReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanExpireReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanExpireReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanExpireReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailCode"] = o.EmailCode
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanExpireReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailCode",
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

	varUnibeeApiMerchantPlanExpireReq := _UnibeeApiMerchantPlanExpireReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanExpireReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanExpireReq(varUnibeeApiMerchantPlanExpireReq)

	return err
}

type NullableUnibeeApiMerchantPlanExpireReq struct {
	value *UnibeeApiMerchantPlanExpireReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanExpireReq) Get() *UnibeeApiMerchantPlanExpireReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanExpireReq) Set(val *UnibeeApiMerchantPlanExpireReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanExpireReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanExpireReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanExpireReq(val *UnibeeApiMerchantPlanExpireReq) *NullableUnibeeApiMerchantPlanExpireReq {
	return &NullableUnibeeApiMerchantPlanExpireReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanExpireReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanExpireReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


