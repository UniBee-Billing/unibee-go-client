/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionAdminNoteListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionAdminNoteListGet200ResponseData{}

// MerchantSubscriptionAdminNoteListGet200ResponseData struct for MerchantSubscriptionAdminNoteListGet200ResponseData
type MerchantSubscriptionAdminNoteListGet200ResponseData struct {
	NoteLists []UnibeeApiMerchantSubscriptionAdminNoteRo `json:"noteLists,omitempty"`
}

// NewMerchantSubscriptionAdminNoteListGet200ResponseData instantiates a new MerchantSubscriptionAdminNoteListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionAdminNoteListGet200ResponseData() *MerchantSubscriptionAdminNoteListGet200ResponseData {
	this := MerchantSubscriptionAdminNoteListGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionAdminNoteListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionAdminNoteListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionAdminNoteListGet200ResponseDataWithDefaults() *MerchantSubscriptionAdminNoteListGet200ResponseData {
	this := MerchantSubscriptionAdminNoteListGet200ResponseData{}
	return &this
}

// GetNoteLists returns the NoteLists field value if set, zero value otherwise.
func (o *MerchantSubscriptionAdminNoteListGet200ResponseData) GetNoteLists() []UnibeeApiMerchantSubscriptionAdminNoteRo {
	if o == nil || IsNil(o.NoteLists) {
		var ret []UnibeeApiMerchantSubscriptionAdminNoteRo
		return ret
	}
	return o.NoteLists
}

// GetNoteListsOk returns a tuple with the NoteLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionAdminNoteListGet200ResponseData) GetNoteListsOk() ([]UnibeeApiMerchantSubscriptionAdminNoteRo, bool) {
	if o == nil || IsNil(o.NoteLists) {
		return nil, false
	}
	return o.NoteLists, true
}

// HasNoteLists returns a boolean if a field has been set.
func (o *MerchantSubscriptionAdminNoteListGet200ResponseData) HasNoteLists() bool {
	if o != nil && !IsNil(o.NoteLists) {
		return true
	}

	return false
}

// SetNoteLists gets a reference to the given []UnibeeApiMerchantSubscriptionAdminNoteRo and assigns it to the NoteLists field.
func (o *MerchantSubscriptionAdminNoteListGet200ResponseData) SetNoteLists(v []UnibeeApiMerchantSubscriptionAdminNoteRo) {
	o.NoteLists = v
}

func (o MerchantSubscriptionAdminNoteListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionAdminNoteListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoteLists) {
		toSerialize["noteLists"] = o.NoteLists
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionAdminNoteListGet200ResponseData struct {
	value *MerchantSubscriptionAdminNoteListGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionAdminNoteListGet200ResponseData) Get() *MerchantSubscriptionAdminNoteListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionAdminNoteListGet200ResponseData) Set(val *MerchantSubscriptionAdminNoteListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionAdminNoteListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionAdminNoteListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionAdminNoteListGet200ResponseData(val *MerchantSubscriptionAdminNoteListGet200ResponseData) *NullableMerchantSubscriptionAdminNoteListGet200ResponseData {
	return &NullableMerchantSubscriptionAdminNoteListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionAdminNoteListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionAdminNoteListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


