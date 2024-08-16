/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionAdminNoteListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionAdminNoteListRes{}

// UnibeeApiMerchantSubscriptionAdminNoteListRes struct for UnibeeApiMerchantSubscriptionAdminNoteListRes
type UnibeeApiMerchantSubscriptionAdminNoteListRes struct {
	NoteLists []UnibeeApiMerchantSubscriptionAdminNoteRo `json:"noteLists,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionAdminNoteListRes instantiates a new UnibeeApiMerchantSubscriptionAdminNoteListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionAdminNoteListRes() *UnibeeApiMerchantSubscriptionAdminNoteListRes {
	this := UnibeeApiMerchantSubscriptionAdminNoteListRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionAdminNoteListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionAdminNoteListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionAdminNoteListResWithDefaults() *UnibeeApiMerchantSubscriptionAdminNoteListRes {
	this := UnibeeApiMerchantSubscriptionAdminNoteListRes{}
	return &this
}

// GetNoteLists returns the NoteLists field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListRes) GetNoteLists() []UnibeeApiMerchantSubscriptionAdminNoteRo {
	if o == nil || IsNil(o.NoteLists) {
		var ret []UnibeeApiMerchantSubscriptionAdminNoteRo
		return ret
	}
	return o.NoteLists
}

// GetNoteListsOk returns a tuple with the NoteLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListRes) GetNoteListsOk() ([]UnibeeApiMerchantSubscriptionAdminNoteRo, bool) {
	if o == nil || IsNil(o.NoteLists) {
		return nil, false
	}
	return o.NoteLists, true
}

// HasNoteLists returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListRes) HasNoteLists() bool {
	if o != nil && !IsNil(o.NoteLists) {
		return true
	}

	return false
}

// SetNoteLists gets a reference to the given []UnibeeApiMerchantSubscriptionAdminNoteRo and assigns it to the NoteLists field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListRes) SetNoteLists(v []UnibeeApiMerchantSubscriptionAdminNoteRo) {
	o.NoteLists = v
}

func (o UnibeeApiMerchantSubscriptionAdminNoteListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionAdminNoteListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoteLists) {
		toSerialize["noteLists"] = o.NoteLists
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionAdminNoteListRes struct {
	value *UnibeeApiMerchantSubscriptionAdminNoteListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) Get() *UnibeeApiMerchantSubscriptionAdminNoteListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) Set(val *UnibeeApiMerchantSubscriptionAdminNoteListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionAdminNoteListRes(val *UnibeeApiMerchantSubscriptionAdminNoteListRes) *NullableUnibeeApiMerchantSubscriptionAdminNoteListRes {
	return &NullableUnibeeApiMerchantSubscriptionAdminNoteListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


