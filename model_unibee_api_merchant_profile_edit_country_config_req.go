/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantProfileEditCountryConfigReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileEditCountryConfigReq{}

// UnibeeApiMerchantProfileEditCountryConfigReq struct for UnibeeApiMerchantProfileEditCountryConfigReq
type UnibeeApiMerchantProfileEditCountryConfigReq struct {
	// CountryCode
	CountryCode string `json:"countryCode"`
	// name
	Name *string `json:"name,omitempty"`
	// VatEnable, Default true
	VatEnable *bool `json:"vatEnable,omitempty"`
}

type _UnibeeApiMerchantProfileEditCountryConfigReq UnibeeApiMerchantProfileEditCountryConfigReq

// NewUnibeeApiMerchantProfileEditCountryConfigReq instantiates a new UnibeeApiMerchantProfileEditCountryConfigReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileEditCountryConfigReq(countryCode string) *UnibeeApiMerchantProfileEditCountryConfigReq {
	this := UnibeeApiMerchantProfileEditCountryConfigReq{}
	this.CountryCode = countryCode
	return &this
}

// NewUnibeeApiMerchantProfileEditCountryConfigReqWithDefaults instantiates a new UnibeeApiMerchantProfileEditCountryConfigReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileEditCountryConfigReqWithDefaults() *UnibeeApiMerchantProfileEditCountryConfigReq {
	this := UnibeeApiMerchantProfileEditCountryConfigReq{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetName(v string) {
	o.Name = &v
}

// GetVatEnable returns the VatEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetVatEnable() bool {
	if o == nil || IsNil(o.VatEnable) {
		var ret bool
		return ret
	}
	return *o.VatEnable
}

// GetVatEnableOk returns a tuple with the VatEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetVatEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.VatEnable) {
		return nil, false
	}
	return o.VatEnable, true
}

// HasVatEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) HasVatEnable() bool {
	if o != nil && !IsNil(o.VatEnable) {
		return true
	}

	return false
}

// SetVatEnable gets a reference to the given bool and assigns it to the VatEnable field.
func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetVatEnable(v bool) {
	o.VatEnable = &v
}

func (o UnibeeApiMerchantProfileEditCountryConfigReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileEditCountryConfigReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VatEnable) {
		toSerialize["vatEnable"] = o.VatEnable
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProfileEditCountryConfigReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"countryCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantProfileEditCountryConfigReq := _UnibeeApiMerchantProfileEditCountryConfigReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProfileEditCountryConfigReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProfileEditCountryConfigReq(varUnibeeApiMerchantProfileEditCountryConfigReq)

	return err
}

type NullableUnibeeApiMerchantProfileEditCountryConfigReq struct {
	value *UnibeeApiMerchantProfileEditCountryConfigReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileEditCountryConfigReq) Get() *UnibeeApiMerchantProfileEditCountryConfigReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileEditCountryConfigReq) Set(val *UnibeeApiMerchantProfileEditCountryConfigReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileEditCountryConfigReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileEditCountryConfigReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileEditCountryConfigReq(val *UnibeeApiMerchantProfileEditCountryConfigReq) *NullableUnibeeApiMerchantProfileEditCountryConfigReq {
	return &NullableUnibeeApiMerchantProfileEditCountryConfigReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileEditCountryConfigReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileEditCountryConfigReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


