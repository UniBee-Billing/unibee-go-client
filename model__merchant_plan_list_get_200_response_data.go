/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPlanListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPlanListGet200ResponseData{}

// MerchantPlanListGet200ResponseData struct for MerchantPlanListGet200ResponseData
type MerchantPlanListGet200ResponseData struct {
	// Plans
	Plans []UnibeeApiBeanDetailPlanDetail `json:"plans,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantPlanListGet200ResponseData instantiates a new MerchantPlanListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPlanListGet200ResponseData() *MerchantPlanListGet200ResponseData {
	this := MerchantPlanListGet200ResponseData{}
	return &this
}

// NewMerchantPlanListGet200ResponseDataWithDefaults instantiates a new MerchantPlanListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPlanListGet200ResponseDataWithDefaults() *MerchantPlanListGet200ResponseData {
	this := MerchantPlanListGet200ResponseData{}
	return &this
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *MerchantPlanListGet200ResponseData) GetPlans() []UnibeeApiBeanDetailPlanDetail {
	if o == nil || IsNil(o.Plans) {
		var ret []UnibeeApiBeanDetailPlanDetail
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPlanListGet200ResponseData) GetPlansOk() ([]UnibeeApiBeanDetailPlanDetail, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *MerchantPlanListGet200ResponseData) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []UnibeeApiBeanDetailPlanDetail and assigns it to the Plans field.
func (o *MerchantPlanListGet200ResponseData) SetPlans(v []UnibeeApiBeanDetailPlanDetail) {
	o.Plans = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantPlanListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPlanListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantPlanListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantPlanListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantPlanListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPlanListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantPlanListGet200ResponseData struct {
	value *MerchantPlanListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPlanListGet200ResponseData) Get() *MerchantPlanListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPlanListGet200ResponseData) Set(val *MerchantPlanListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPlanListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPlanListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPlanListGet200ResponseData(val *MerchantPlanListGet200ResponseData) *NullableMerchantPlanListGet200ResponseData {
	return &NullableMerchantPlanListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPlanListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPlanListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


