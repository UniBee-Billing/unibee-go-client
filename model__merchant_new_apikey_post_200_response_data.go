/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantNewApikeyPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantNewApikeyPost200ResponseData{}

// MerchantNewApikeyPost200ResponseData struct for MerchantNewApikeyPost200ResponseData
type MerchantNewApikeyPost200ResponseData struct {
	// ApiKey
	ApiKey *string `json:"apiKey,omitempty"`
}

// NewMerchantNewApikeyPost200ResponseData instantiates a new MerchantNewApikeyPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantNewApikeyPost200ResponseData() *MerchantNewApikeyPost200ResponseData {
	this := MerchantNewApikeyPost200ResponseData{}
	return &this
}

// NewMerchantNewApikeyPost200ResponseDataWithDefaults instantiates a new MerchantNewApikeyPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantNewApikeyPost200ResponseDataWithDefaults() *MerchantNewApikeyPost200ResponseData {
	this := MerchantNewApikeyPost200ResponseData{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *MerchantNewApikeyPost200ResponseData) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantNewApikeyPost200ResponseData) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *MerchantNewApikeyPost200ResponseData) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *MerchantNewApikeyPost200ResponseData) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o MerchantNewApikeyPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantNewApikeyPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullableMerchantNewApikeyPost200ResponseData struct {
	value *MerchantNewApikeyPost200ResponseData
	isSet bool
}

func (v NullableMerchantNewApikeyPost200ResponseData) Get() *MerchantNewApikeyPost200ResponseData {
	return v.value
}

func (v *NullableMerchantNewApikeyPost200ResponseData) Set(val *MerchantNewApikeyPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantNewApikeyPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantNewApikeyPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantNewApikeyPost200ResponseData(val *MerchantNewApikeyPost200ResponseData) *NullableMerchantNewApikeyPost200ResponseData {
	return &NullableMerchantNewApikeyPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantNewApikeyPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantNewApikeyPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


