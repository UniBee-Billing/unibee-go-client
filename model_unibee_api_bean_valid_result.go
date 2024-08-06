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

// checks if the UnibeeApiBeanValidResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanValidResult{}

// UnibeeApiBeanValidResult struct for UnibeeApiBeanValidResult
type UnibeeApiBeanValidResult struct {
	CompanyAddress *string `json:"companyAddress,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	ValidateMessage *string `json:"validateMessage,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewUnibeeApiBeanValidResult instantiates a new UnibeeApiBeanValidResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanValidResult() *UnibeeApiBeanValidResult {
	this := UnibeeApiBeanValidResult{}
	return &this
}

// NewUnibeeApiBeanValidResultWithDefaults instantiates a new UnibeeApiBeanValidResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanValidResultWithDefaults() *UnibeeApiBeanValidResult {
	this := UnibeeApiBeanValidResult{}
	return &this
}

// GetCompanyAddress returns the CompanyAddress field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetCompanyAddress() string {
	if o == nil || IsNil(o.CompanyAddress) {
		var ret string
		return ret
	}
	return *o.CompanyAddress
}

// GetCompanyAddressOk returns a tuple with the CompanyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetCompanyAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyAddress) {
		return nil, false
	}
	return o.CompanyAddress, true
}

// HasCompanyAddress returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasCompanyAddress() bool {
	if o != nil && !IsNil(o.CompanyAddress) {
		return true
	}

	return false
}

// SetCompanyAddress gets a reference to the given string and assigns it to the CompanyAddress field.
func (o *UnibeeApiBeanValidResult) SetCompanyAddress(v string) {
	o.CompanyAddress = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UnibeeApiBeanValidResult) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanValidResult) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *UnibeeApiBeanValidResult) SetValid(v bool) {
	o.Valid = &v
}

// GetValidateMessage returns the ValidateMessage field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetValidateMessage() string {
	if o == nil || IsNil(o.ValidateMessage) {
		var ret string
		return ret
	}
	return *o.ValidateMessage
}

// GetValidateMessageOk returns a tuple with the ValidateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetValidateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidateMessage) {
		return nil, false
	}
	return o.ValidateMessage, true
}

// HasValidateMessage returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasValidateMessage() bool {
	if o != nil && !IsNil(o.ValidateMessage) {
		return true
	}

	return false
}

// SetValidateMessage gets a reference to the given string and assigns it to the ValidateMessage field.
func (o *UnibeeApiBeanValidResult) SetValidateMessage(v string) {
	o.ValidateMessage = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UnibeeApiBeanValidResult) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanValidResult) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UnibeeApiBeanValidResult) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UnibeeApiBeanValidResult) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o UnibeeApiBeanValidResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanValidResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyAddress) {
		toSerialize["companyAddress"] = o.CompanyAddress
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !IsNil(o.ValidateMessage) {
		toSerialize["validateMessage"] = o.ValidateMessage
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanValidResult struct {
	value *UnibeeApiBeanValidResult
	isSet bool
}

func (v NullableUnibeeApiBeanValidResult) Get() *UnibeeApiBeanValidResult {
	return v.value
}

func (v *NullableUnibeeApiBeanValidResult) Set(val *UnibeeApiBeanValidResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanValidResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanValidResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanValidResult(val *UnibeeApiBeanValidResult) *NullableUnibeeApiBeanValidResult {
	return &NullableUnibeeApiBeanValidResult{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanValidResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanValidResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


