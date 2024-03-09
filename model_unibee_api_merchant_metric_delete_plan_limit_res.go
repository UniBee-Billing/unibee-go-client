/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricDeletePlanLimitRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDeletePlanLimitRes{}

// UnibeeApiMerchantMetricDeletePlanLimitRes struct for UnibeeApiMerchantMetricDeletePlanLimitRes
type UnibeeApiMerchantMetricDeletePlanLimitRes struct {
	MerchantMetricPlanLimit *UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo `json:"merchantMetricPlanLimit,omitempty"`
}

// NewUnibeeApiMerchantMetricDeletePlanLimitRes instantiates a new UnibeeApiMerchantMetricDeletePlanLimitRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDeletePlanLimitRes() *UnibeeApiMerchantMetricDeletePlanLimitRes {
	this := UnibeeApiMerchantMetricDeletePlanLimitRes{}
	return &this
}

// NewUnibeeApiMerchantMetricDeletePlanLimitResWithDefaults instantiates a new UnibeeApiMerchantMetricDeletePlanLimitRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDeletePlanLimitResWithDefaults() *UnibeeApiMerchantMetricDeletePlanLimitRes {
	this := UnibeeApiMerchantMetricDeletePlanLimitRes{}
	return &this
}

// GetMerchantMetricPlanLimit returns the MerchantMetricPlanLimit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricDeletePlanLimitRes) GetMerchantMetricPlanLimit() UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		var ret UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo
		return ret
	}
	return *o.MerchantMetricPlanLimit
}

// GetMerchantMetricPlanLimitOk returns a tuple with the MerchantMetricPlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeletePlanLimitRes) GetMerchantMetricPlanLimitOk() (*UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo, bool) {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		return nil, false
	}
	return o.MerchantMetricPlanLimit, true
}

// HasMerchantMetricPlanLimit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricDeletePlanLimitRes) HasMerchantMetricPlanLimit() bool {
	if o != nil && !IsNil(o.MerchantMetricPlanLimit) {
		return true
	}

	return false
}

// SetMerchantMetricPlanLimit gets a reference to the given UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo and assigns it to the MerchantMetricPlanLimit field.
func (o *UnibeeApiMerchantMetricDeletePlanLimitRes) SetMerchantMetricPlanLimit(v UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo) {
	o.MerchantMetricPlanLimit = &v
}

func (o UnibeeApiMerchantMetricDeletePlanLimitRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDeletePlanLimitRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricPlanLimit) {
		toSerialize["merchantMetricPlanLimit"] = o.MerchantMetricPlanLimit
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricDeletePlanLimitRes struct {
	value *UnibeeApiMerchantMetricDeletePlanLimitRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitRes) Get() *UnibeeApiMerchantMetricDeletePlanLimitRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitRes) Set(val *UnibeeApiMerchantMetricDeletePlanLimitRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDeletePlanLimitRes(val *UnibeeApiMerchantMetricDeletePlanLimitRes) *NullableUnibeeApiMerchantMetricDeletePlanLimitRes {
	return &NullableUnibeeApiMerchantMetricDeletePlanLimitRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDeletePlanLimitRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDeletePlanLimitRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


