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

// checks if the MerchantOssFilePost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantOssFilePost200ResponseData{}

// MerchantOssFilePost200ResponseData struct for MerchantOssFilePost200ResponseData
type MerchantOssFilePost200ResponseData struct {
	// URL Of File Or Image
	Url *string `json:"url,omitempty"`
}

// NewMerchantOssFilePost200ResponseData instantiates a new MerchantOssFilePost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantOssFilePost200ResponseData() *MerchantOssFilePost200ResponseData {
	this := MerchantOssFilePost200ResponseData{}
	return &this
}

// NewMerchantOssFilePost200ResponseDataWithDefaults instantiates a new MerchantOssFilePost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantOssFilePost200ResponseDataWithDefaults() *MerchantOssFilePost200ResponseData {
	this := MerchantOssFilePost200ResponseData{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MerchantOssFilePost200ResponseData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOssFilePost200ResponseData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MerchantOssFilePost200ResponseData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MerchantOssFilePost200ResponseData) SetUrl(v string) {
	o.Url = &v
}

func (o MerchantOssFilePost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantOssFilePost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableMerchantOssFilePost200ResponseData struct {
	value *MerchantOssFilePost200ResponseData
	isSet bool
}

func (v NullableMerchantOssFilePost200ResponseData) Get() *MerchantOssFilePost200ResponseData {
	return v.value
}

func (v *NullableMerchantOssFilePost200ResponseData) Set(val *MerchantOssFilePost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantOssFilePost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantOssFilePost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantOssFilePost200ResponseData(val *MerchantOssFilePost200ResponseData) *NullableMerchantOssFilePost200ResponseData {
	return &NullableMerchantOssFilePost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantOssFilePost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantOssFilePost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


