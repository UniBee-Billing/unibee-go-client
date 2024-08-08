/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantOssFileUploadRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantOssFileUploadRes{}

// UnibeeApiMerchantOssFileUploadRes struct for UnibeeApiMerchantOssFileUploadRes
type UnibeeApiMerchantOssFileUploadRes struct {
	// URL Of File Or Image
	Url *string `json:"url,omitempty"`
}

// NewUnibeeApiMerchantOssFileUploadRes instantiates a new UnibeeApiMerchantOssFileUploadRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantOssFileUploadRes() *UnibeeApiMerchantOssFileUploadRes {
	this := UnibeeApiMerchantOssFileUploadRes{}
	return &this
}

// NewUnibeeApiMerchantOssFileUploadResWithDefaults instantiates a new UnibeeApiMerchantOssFileUploadRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantOssFileUploadResWithDefaults() *UnibeeApiMerchantOssFileUploadRes {
	this := UnibeeApiMerchantOssFileUploadRes{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UnibeeApiMerchantOssFileUploadRes) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantOssFileUploadRes) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantOssFileUploadRes) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UnibeeApiMerchantOssFileUploadRes) SetUrl(v string) {
	o.Url = &v
}

func (o UnibeeApiMerchantOssFileUploadRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantOssFileUploadRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantOssFileUploadRes struct {
	value *UnibeeApiMerchantOssFileUploadRes
	isSet bool
}

func (v NullableUnibeeApiMerchantOssFileUploadRes) Get() *UnibeeApiMerchantOssFileUploadRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantOssFileUploadRes) Set(val *UnibeeApiMerchantOssFileUploadRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantOssFileUploadRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantOssFileUploadRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantOssFileUploadRes(val *UnibeeApiMerchantOssFileUploadRes) *NullableUnibeeApiMerchantOssFileUploadRes {
	return &NullableUnibeeApiMerchantOssFileUploadRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantOssFileUploadRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantOssFileUploadRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


