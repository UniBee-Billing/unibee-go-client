/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantCountryConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantCountryConfig{}

// UnibeeApiBeanMerchantCountryConfig struct for UnibeeApiBeanMerchantCountryConfig
type UnibeeApiBeanMerchantCountryConfig struct {
	CountryCode *string `json:"countryCode,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUnibeeApiBeanMerchantCountryConfig instantiates a new UnibeeApiBeanMerchantCountryConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantCountryConfig() *UnibeeApiBeanMerchantCountryConfig {
	this := UnibeeApiBeanMerchantCountryConfig{}
	return &this
}

// NewUnibeeApiBeanMerchantCountryConfigWithDefaults instantiates a new UnibeeApiBeanMerchantCountryConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantCountryConfigWithDefaults() *UnibeeApiBeanMerchantCountryConfig {
	this := UnibeeApiBeanMerchantCountryConfig{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfig) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanMerchantCountryConfig) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfig) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantCountryConfig) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanMerchantCountryConfig) SetName(v string) {
	o.Name = &v
}

func (o UnibeeApiBeanMerchantCountryConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantCountryConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantCountryConfig struct {
	value *UnibeeApiBeanMerchantCountryConfig
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantCountryConfig) Get() *UnibeeApiBeanMerchantCountryConfig {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantCountryConfig) Set(val *UnibeeApiBeanMerchantCountryConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantCountryConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantCountryConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantCountryConfig(val *UnibeeApiBeanMerchantCountryConfig) *NullableUnibeeApiBeanMerchantCountryConfig {
	return &NullableUnibeeApiBeanMerchantCountryConfig{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantCountryConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantCountryConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


