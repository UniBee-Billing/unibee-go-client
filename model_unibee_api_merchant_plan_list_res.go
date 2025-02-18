/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanListRes{}

// UnibeeApiMerchantPlanListRes struct for UnibeeApiMerchantPlanListRes
type UnibeeApiMerchantPlanListRes struct {
	// Plans
	Plans []UnibeeApiBeanDetailPlanDetail `json:"plans,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantPlanListRes instantiates a new UnibeeApiMerchantPlanListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanListRes() *UnibeeApiMerchantPlanListRes {
	this := UnibeeApiMerchantPlanListRes{}
	return &this
}

// NewUnibeeApiMerchantPlanListResWithDefaults instantiates a new UnibeeApiMerchantPlanListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanListResWithDefaults() *UnibeeApiMerchantPlanListRes {
	this := UnibeeApiMerchantPlanListRes{}
	return &this
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListRes) GetPlans() []UnibeeApiBeanDetailPlanDetail {
	if o == nil || IsNil(o.Plans) {
		var ret []UnibeeApiBeanDetailPlanDetail
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListRes) GetPlansOk() ([]UnibeeApiBeanDetailPlanDetail, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListRes) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []UnibeeApiBeanDetailPlanDetail and assigns it to the Plans field.
func (o *UnibeeApiMerchantPlanListRes) SetPlans(v []UnibeeApiBeanDetailPlanDetail) {
	o.Plans = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantPlanListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantPlanListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanListRes struct {
	value *UnibeeApiMerchantPlanListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanListRes) Get() *UnibeeApiMerchantPlanListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanListRes) Set(val *UnibeeApiMerchantPlanListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanListRes(val *UnibeeApiMerchantPlanListRes) *NullableUnibeeApiMerchantPlanListRes {
	return &NullableUnibeeApiMerchantPlanListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


