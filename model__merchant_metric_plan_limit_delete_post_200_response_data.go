/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMetricPlanLimitDeletePost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricPlanLimitDeletePost200ResponseData{}

// MerchantMetricPlanLimitDeletePost200ResponseData struct for MerchantMetricPlanLimitDeletePost200ResponseData
type MerchantMetricPlanLimitDeletePost200ResponseData struct {
	MerchantMetricPlanLimit *UnibeeApiBeanMerchantMetricPlanLimit `json:"merchantMetricPlanLimit,omitempty"`
}

// NewMerchantMetricPlanLimitDeletePost200ResponseData instantiates a new MerchantMetricPlanLimitDeletePost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricPlanLimitDeletePost200ResponseData() *MerchantMetricPlanLimitDeletePost200ResponseData {
	this := MerchantMetricPlanLimitDeletePost200ResponseData{}
	return &this
}

// NewMerchantMetricPlanLimitDeletePost200ResponseDataWithDefaults instantiates a new MerchantMetricPlanLimitDeletePost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricPlanLimitDeletePost200ResponseDataWithDefaults() *MerchantMetricPlanLimitDeletePost200ResponseData {
	this := MerchantMetricPlanLimitDeletePost200ResponseData{}
	return &this
}

// GetMerchantMetricPlanLimit returns the MerchantMetricPlanLimit field value if set, zero value otherwise.
func (o *MerchantMetricPlanLimitDeletePost200ResponseData) GetMerchantMetricPlanLimit() UnibeeApiBeanMerchantMetricPlanLimit {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		var ret UnibeeApiBeanMerchantMetricPlanLimit
		return ret
	}
	return *o.MerchantMetricPlanLimit
}

// GetMerchantMetricPlanLimitOk returns a tuple with the MerchantMetricPlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricPlanLimitDeletePost200ResponseData) GetMerchantMetricPlanLimitOk() (*UnibeeApiBeanMerchantMetricPlanLimit, bool) {
	if o == nil || IsNil(o.MerchantMetricPlanLimit) {
		return nil, false
	}
	return o.MerchantMetricPlanLimit, true
}

// HasMerchantMetricPlanLimit returns a boolean if a field has been set.
func (o *MerchantMetricPlanLimitDeletePost200ResponseData) HasMerchantMetricPlanLimit() bool {
	if o != nil && !IsNil(o.MerchantMetricPlanLimit) {
		return true
	}

	return false
}

// SetMerchantMetricPlanLimit gets a reference to the given UnibeeApiBeanMerchantMetricPlanLimit and assigns it to the MerchantMetricPlanLimit field.
func (o *MerchantMetricPlanLimitDeletePost200ResponseData) SetMerchantMetricPlanLimit(v UnibeeApiBeanMerchantMetricPlanLimit) {
	o.MerchantMetricPlanLimit = &v
}

func (o MerchantMetricPlanLimitDeletePost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricPlanLimitDeletePost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricPlanLimit) {
		toSerialize["merchantMetricPlanLimit"] = o.MerchantMetricPlanLimit
	}
	return toSerialize, nil
}

type NullableMerchantMetricPlanLimitDeletePost200ResponseData struct {
	value *MerchantMetricPlanLimitDeletePost200ResponseData
	isSet bool
}

func (v NullableMerchantMetricPlanLimitDeletePost200ResponseData) Get() *MerchantMetricPlanLimitDeletePost200ResponseData {
	return v.value
}

func (v *NullableMerchantMetricPlanLimitDeletePost200ResponseData) Set(val *MerchantMetricPlanLimitDeletePost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricPlanLimitDeletePost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricPlanLimitDeletePost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricPlanLimitDeletePost200ResponseData(val *MerchantMetricPlanLimitDeletePost200ResponseData) *NullableMerchantMetricPlanLimitDeletePost200ResponseData {
	return &NullableMerchantMetricPlanLimitDeletePost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMetricPlanLimitDeletePost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricPlanLimitDeletePost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


