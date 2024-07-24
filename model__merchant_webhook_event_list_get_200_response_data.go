/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantWebhookEventListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantWebhookEventListGet200ResponseData{}

// MerchantWebhookEventListGet200ResponseData struct for MerchantWebhookEventListGet200ResponseData
type MerchantWebhookEventListGet200ResponseData struct {
	// EventList
	EventList []string `json:"eventList,omitempty"`
}

// NewMerchantWebhookEventListGet200ResponseData instantiates a new MerchantWebhookEventListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantWebhookEventListGet200ResponseData() *MerchantWebhookEventListGet200ResponseData {
	this := MerchantWebhookEventListGet200ResponseData{}
	return &this
}

// NewMerchantWebhookEventListGet200ResponseDataWithDefaults instantiates a new MerchantWebhookEventListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWebhookEventListGet200ResponseDataWithDefaults() *MerchantWebhookEventListGet200ResponseData {
	this := MerchantWebhookEventListGet200ResponseData{}
	return &this
}

// GetEventList returns the EventList field value if set, zero value otherwise.
func (o *MerchantWebhookEventListGet200ResponseData) GetEventList() []string {
	if o == nil || IsNil(o.EventList) {
		var ret []string
		return ret
	}
	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookEventListGet200ResponseData) GetEventListOk() ([]string, bool) {
	if o == nil || IsNil(o.EventList) {
		return nil, false
	}
	return o.EventList, true
}

// HasEventList returns a boolean if a field has been set.
func (o *MerchantWebhookEventListGet200ResponseData) HasEventList() bool {
	if o != nil && !IsNil(o.EventList) {
		return true
	}

	return false
}

// SetEventList gets a reference to the given []string and assigns it to the EventList field.
func (o *MerchantWebhookEventListGet200ResponseData) SetEventList(v []string) {
	o.EventList = v
}

func (o MerchantWebhookEventListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantWebhookEventListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventList) {
		toSerialize["eventList"] = o.EventList
	}
	return toSerialize, nil
}

type NullableMerchantWebhookEventListGet200ResponseData struct {
	value *MerchantWebhookEventListGet200ResponseData
	isSet bool
}

func (v NullableMerchantWebhookEventListGet200ResponseData) Get() *MerchantWebhookEventListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantWebhookEventListGet200ResponseData) Set(val *MerchantWebhookEventListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantWebhookEventListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantWebhookEventListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantWebhookEventListGet200ResponseData(val *MerchantWebhookEventListGet200ResponseData) *NullableMerchantWebhookEventListGet200ResponseData {
	return &NullableMerchantWebhookEventListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantWebhookEventListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantWebhookEventListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


