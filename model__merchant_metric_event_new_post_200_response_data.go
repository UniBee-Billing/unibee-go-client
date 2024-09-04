/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMetricEventNewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricEventNewPost200ResponseData{}

// MerchantMetricEventNewPost200ResponseData struct for MerchantMetricEventNewPost200ResponseData
type MerchantMetricEventNewPost200ResponseData struct {
	MerchantMetricEvent *UnibeeApiBeanMerchantMetricEvent `json:"merchantMetricEvent,omitempty"`
}

// NewMerchantMetricEventNewPost200ResponseData instantiates a new MerchantMetricEventNewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricEventNewPost200ResponseData() *MerchantMetricEventNewPost200ResponseData {
	this := MerchantMetricEventNewPost200ResponseData{}
	return &this
}

// NewMerchantMetricEventNewPost200ResponseDataWithDefaults instantiates a new MerchantMetricEventNewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricEventNewPost200ResponseDataWithDefaults() *MerchantMetricEventNewPost200ResponseData {
	this := MerchantMetricEventNewPost200ResponseData{}
	return &this
}

// GetMerchantMetricEvent returns the MerchantMetricEvent field value if set, zero value otherwise.
func (o *MerchantMetricEventNewPost200ResponseData) GetMerchantMetricEvent() UnibeeApiBeanMerchantMetricEvent {
	if o == nil || IsNil(o.MerchantMetricEvent) {
		var ret UnibeeApiBeanMerchantMetricEvent
		return ret
	}
	return *o.MerchantMetricEvent
}

// GetMerchantMetricEventOk returns a tuple with the MerchantMetricEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricEventNewPost200ResponseData) GetMerchantMetricEventOk() (*UnibeeApiBeanMerchantMetricEvent, bool) {
	if o == nil || IsNil(o.MerchantMetricEvent) {
		return nil, false
	}
	return o.MerchantMetricEvent, true
}

// HasMerchantMetricEvent returns a boolean if a field has been set.
func (o *MerchantMetricEventNewPost200ResponseData) HasMerchantMetricEvent() bool {
	if o != nil && !IsNil(o.MerchantMetricEvent) {
		return true
	}

	return false
}

// SetMerchantMetricEvent gets a reference to the given UnibeeApiBeanMerchantMetricEvent and assigns it to the MerchantMetricEvent field.
func (o *MerchantMetricEventNewPost200ResponseData) SetMerchantMetricEvent(v UnibeeApiBeanMerchantMetricEvent) {
	o.MerchantMetricEvent = &v
}

func (o MerchantMetricEventNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricEventNewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricEvent) {
		toSerialize["merchantMetricEvent"] = o.MerchantMetricEvent
	}
	return toSerialize, nil
}

type NullableMerchantMetricEventNewPost200ResponseData struct {
	value *MerchantMetricEventNewPost200ResponseData
	isSet bool
}

func (v NullableMerchantMetricEventNewPost200ResponseData) Get() *MerchantMetricEventNewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantMetricEventNewPost200ResponseData) Set(val *MerchantMetricEventNewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricEventNewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricEventNewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricEventNewPost200ResponseData(val *MerchantMetricEventNewPost200ResponseData) *NullableMerchantMetricEventNewPost200ResponseData {
	return &NullableMerchantMetricEventNewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMetricEventNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricEventNewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


