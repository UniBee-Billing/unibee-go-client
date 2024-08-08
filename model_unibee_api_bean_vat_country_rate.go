/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanVatCountryRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanVatCountryRate{}

// UnibeeApiBeanVatCountryRate struct for UnibeeApiBeanVatCountryRate
type UnibeeApiBeanVatCountryRate struct {
	CountryCode *string `json:"countryCode,omitempty"`
	CountryName *string `json:"countryName,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	// TaxId
	Id *int64 `json:"id,omitempty"`
	IsEU *bool `json:"isEU,omitempty"`
	// Tax税率，万分位，1000 表示 10%
	StandardTaxPercentage *int64 `json:"standardTaxPercentage,omitempty"`
	// vat support,true or false
	VatSupport *bool `json:"vatSupport,omitempty"`
}

// NewUnibeeApiBeanVatCountryRate instantiates a new UnibeeApiBeanVatCountryRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanVatCountryRate() *UnibeeApiBeanVatCountryRate {
	this := UnibeeApiBeanVatCountryRate{}
	return &this
}

// NewUnibeeApiBeanVatCountryRateWithDefaults instantiates a new UnibeeApiBeanVatCountryRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanVatCountryRateWithDefaults() *UnibeeApiBeanVatCountryRate {
	this := UnibeeApiBeanVatCountryRate{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanVatCountryRate) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetCountryName() string {
	if o == nil || IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasCountryName() bool {
	if o != nil && !IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *UnibeeApiBeanVatCountryRate) SetCountryName(v string) {
	o.CountryName = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UnibeeApiBeanVatCountryRate) SetGateway(v string) {
	o.Gateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanVatCountryRate) SetId(v int64) {
	o.Id = &v
}

// GetIsEU returns the IsEU field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetIsEU() bool {
	if o == nil || IsNil(o.IsEU) {
		var ret bool
		return ret
	}
	return *o.IsEU
}

// GetIsEUOk returns a tuple with the IsEU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetIsEUOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEU) {
		return nil, false
	}
	return o.IsEU, true
}

// HasIsEU returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasIsEU() bool {
	if o != nil && !IsNil(o.IsEU) {
		return true
	}

	return false
}

// SetIsEU gets a reference to the given bool and assigns it to the IsEU field.
func (o *UnibeeApiBeanVatCountryRate) SetIsEU(v bool) {
	o.IsEU = &v
}

// GetStandardTaxPercentage returns the StandardTaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetStandardTaxPercentage() int64 {
	if o == nil || IsNil(o.StandardTaxPercentage) {
		var ret int64
		return ret
	}
	return *o.StandardTaxPercentage
}

// GetStandardTaxPercentageOk returns a tuple with the StandardTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetStandardTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.StandardTaxPercentage) {
		return nil, false
	}
	return o.StandardTaxPercentage, true
}

// HasStandardTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasStandardTaxPercentage() bool {
	if o != nil && !IsNil(o.StandardTaxPercentage) {
		return true
	}

	return false
}

// SetStandardTaxPercentage gets a reference to the given int64 and assigns it to the StandardTaxPercentage field.
func (o *UnibeeApiBeanVatCountryRate) SetStandardTaxPercentage(v int64) {
	o.StandardTaxPercentage = &v
}

// GetVatSupport returns the VatSupport field value if set, zero value otherwise.
func (o *UnibeeApiBeanVatCountryRate) GetVatSupport() bool {
	if o == nil || IsNil(o.VatSupport) {
		var ret bool
		return ret
	}
	return *o.VatSupport
}

// GetVatSupportOk returns a tuple with the VatSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanVatCountryRate) GetVatSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.VatSupport) {
		return nil, false
	}
	return o.VatSupport, true
}

// HasVatSupport returns a boolean if a field has been set.
func (o *UnibeeApiBeanVatCountryRate) HasVatSupport() bool {
	if o != nil && !IsNil(o.VatSupport) {
		return true
	}

	return false
}

// SetVatSupport gets a reference to the given bool and assigns it to the VatSupport field.
func (o *UnibeeApiBeanVatCountryRate) SetVatSupport(v bool) {
	o.VatSupport = &v
}

func (o UnibeeApiBeanVatCountryRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanVatCountryRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CountryName) {
		toSerialize["countryName"] = o.CountryName
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsEU) {
		toSerialize["isEU"] = o.IsEU
	}
	if !IsNil(o.StandardTaxPercentage) {
		toSerialize["standardTaxPercentage"] = o.StandardTaxPercentage
	}
	if !IsNil(o.VatSupport) {
		toSerialize["vatSupport"] = o.VatSupport
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanVatCountryRate struct {
	value *UnibeeApiBeanVatCountryRate
	isSet bool
}

func (v NullableUnibeeApiBeanVatCountryRate) Get() *UnibeeApiBeanVatCountryRate {
	return v.value
}

func (v *NullableUnibeeApiBeanVatCountryRate) Set(val *UnibeeApiBeanVatCountryRate) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanVatCountryRate) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanVatCountryRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanVatCountryRate(val *UnibeeApiBeanVatCountryRate) *NullableUnibeeApiBeanVatCountryRate {
	return &NullableUnibeeApiBeanVatCountryRate{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanVatCountryRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanVatCountryRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


