/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantEmailTemplateListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantEmailTemplateListGet200ResponseData{}

// MerchantEmailTemplateListGet200ResponseData struct for MerchantEmailTemplateListGet200ResponseData
type MerchantEmailTemplateListGet200ResponseData struct {
	// Email Template Object List
	EmailTemplateList []UnibeeApiBeanMerchantEmailTemplate `json:"emailTemplateList,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantEmailTemplateListGet200ResponseData instantiates a new MerchantEmailTemplateListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantEmailTemplateListGet200ResponseData() *MerchantEmailTemplateListGet200ResponseData {
	this := MerchantEmailTemplateListGet200ResponseData{}
	return &this
}

// NewMerchantEmailTemplateListGet200ResponseDataWithDefaults instantiates a new MerchantEmailTemplateListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantEmailTemplateListGet200ResponseDataWithDefaults() *MerchantEmailTemplateListGet200ResponseData {
	this := MerchantEmailTemplateListGet200ResponseData{}
	return &this
}

// GetEmailTemplateList returns the EmailTemplateList field value if set, zero value otherwise.
func (o *MerchantEmailTemplateListGet200ResponseData) GetEmailTemplateList() []UnibeeApiBeanMerchantEmailTemplate {
	if o == nil || IsNil(o.EmailTemplateList) {
		var ret []UnibeeApiBeanMerchantEmailTemplate
		return ret
	}
	return o.EmailTemplateList
}

// GetEmailTemplateListOk returns a tuple with the EmailTemplateList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantEmailTemplateListGet200ResponseData) GetEmailTemplateListOk() ([]UnibeeApiBeanMerchantEmailTemplate, bool) {
	if o == nil || IsNil(o.EmailTemplateList) {
		return nil, false
	}
	return o.EmailTemplateList, true
}

// HasEmailTemplateList returns a boolean if a field has been set.
func (o *MerchantEmailTemplateListGet200ResponseData) HasEmailTemplateList() bool {
	if o != nil && !IsNil(o.EmailTemplateList) {
		return true
	}

	return false
}

// SetEmailTemplateList gets a reference to the given []UnibeeApiBeanMerchantEmailTemplate and assigns it to the EmailTemplateList field.
func (o *MerchantEmailTemplateListGet200ResponseData) SetEmailTemplateList(v []UnibeeApiBeanMerchantEmailTemplate) {
	o.EmailTemplateList = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantEmailTemplateListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantEmailTemplateListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantEmailTemplateListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantEmailTemplateListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantEmailTemplateListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantEmailTemplateListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailTemplateList) {
		toSerialize["emailTemplateList"] = o.EmailTemplateList
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantEmailTemplateListGet200ResponseData struct {
	value *MerchantEmailTemplateListGet200ResponseData
	isSet bool
}

func (v NullableMerchantEmailTemplateListGet200ResponseData) Get() *MerchantEmailTemplateListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantEmailTemplateListGet200ResponseData) Set(val *MerchantEmailTemplateListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantEmailTemplateListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantEmailTemplateListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantEmailTemplateListGet200ResponseData(val *MerchantEmailTemplateListGet200ResponseData) *NullableMerchantEmailTemplateListGet200ResponseData {
	return &NullableMerchantEmailTemplateListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantEmailTemplateListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantEmailTemplateListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


