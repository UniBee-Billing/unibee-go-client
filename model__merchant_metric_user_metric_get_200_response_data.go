/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMetricUserMetricGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricUserMetricGet200ResponseData{}

// MerchantMetricUserMetricGet200ResponseData struct for MerchantMetricUserMetricGet200ResponseData
type MerchantMetricUserMetricGet200ResponseData struct {
	UserMetric *UnibeeApiBeanDetailUserMetric `json:"userMetric,omitempty"`
}

// NewMerchantMetricUserMetricGet200ResponseData instantiates a new MerchantMetricUserMetricGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricUserMetricGet200ResponseData() *MerchantMetricUserMetricGet200ResponseData {
	this := MerchantMetricUserMetricGet200ResponseData{}
	return &this
}

// NewMerchantMetricUserMetricGet200ResponseDataWithDefaults instantiates a new MerchantMetricUserMetricGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricUserMetricGet200ResponseDataWithDefaults() *MerchantMetricUserMetricGet200ResponseData {
	this := MerchantMetricUserMetricGet200ResponseData{}
	return &this
}

// GetUserMetric returns the UserMetric field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200ResponseData) GetUserMetric() UnibeeApiBeanDetailUserMetric {
	if o == nil || IsNil(o.UserMetric) {
		var ret UnibeeApiBeanDetailUserMetric
		return ret
	}
	return *o.UserMetric
}

// GetUserMetricOk returns a tuple with the UserMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200ResponseData) GetUserMetricOk() (*UnibeeApiBeanDetailUserMetric, bool) {
	if o == nil || IsNil(o.UserMetric) {
		return nil, false
	}
	return o.UserMetric, true
}

// HasUserMetric returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200ResponseData) HasUserMetric() bool {
	if o != nil && !IsNil(o.UserMetric) {
		return true
	}

	return false
}

// SetUserMetric gets a reference to the given UnibeeApiBeanDetailUserMetric and assigns it to the UserMetric field.
func (o *MerchantMetricUserMetricGet200ResponseData) SetUserMetric(v UnibeeApiBeanDetailUserMetric) {
	o.UserMetric = &v
}

func (o MerchantMetricUserMetricGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricUserMetricGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserMetric) {
		toSerialize["userMetric"] = o.UserMetric
	}
	return toSerialize, nil
}

type NullableMerchantMetricUserMetricGet200ResponseData struct {
	value *MerchantMetricUserMetricGet200ResponseData
	isSet bool
}

func (v NullableMerchantMetricUserMetricGet200ResponseData) Get() *MerchantMetricUserMetricGet200ResponseData {
	return v.value
}

func (v *NullableMerchantMetricUserMetricGet200ResponseData) Set(val *MerchantMetricUserMetricGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricUserMetricGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricUserMetricGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricUserMetricGet200ResponseData(val *MerchantMetricUserMetricGet200ResponseData) *NullableMerchantMetricUserMetricGet200ResponseData {
	return &NullableMerchantMetricUserMetricGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMetricUserMetricGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricUserMetricGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


