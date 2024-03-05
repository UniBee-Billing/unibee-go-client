/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoVatCountryRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoVatCountryRate{}

// UnibeeInternalLogicGatewayRoVatCountryRate struct for UnibeeInternalLogicGatewayRoVatCountryRate
type UnibeeInternalLogicGatewayRoVatCountryRate struct {
	CountryCode *string `json:"countryCode,omitempty"`
	CountryName *string `json:"countryName,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	// TaxId
	Id *int64 `json:"id,omitempty"`
	// Tax税率，万分位，1000 表示 10%
	StandardTaxPercentage *int64 `json:"standardTaxPercentage,omitempty"`
	// vat support,true or false
	VatSupport *bool `json:"vatSupport,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoVatCountryRate instantiates a new UnibeeInternalLogicGatewayRoVatCountryRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoVatCountryRate() *UnibeeInternalLogicGatewayRoVatCountryRate {
	this := UnibeeInternalLogicGatewayRoVatCountryRate{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoVatCountryRateWithDefaults instantiates a new UnibeeInternalLogicGatewayRoVatCountryRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoVatCountryRateWithDefaults() *UnibeeInternalLogicGatewayRoVatCountryRate {
	this := UnibeeInternalLogicGatewayRoVatCountryRate{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetCountryName() string {
	if o == nil || IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasCountryName() bool {
	if o != nil && !IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetCountryName(v string) {
	o.CountryName = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetGateway(v string) {
	o.Gateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetId(v int64) {
	o.Id = &v
}

// GetStandardTaxPercentage returns the StandardTaxPercentage field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetStandardTaxPercentage() int64 {
	if o == nil || IsNil(o.StandardTaxPercentage) {
		var ret int64
		return ret
	}
	return *o.StandardTaxPercentage
}

// GetStandardTaxPercentageOk returns a tuple with the StandardTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetStandardTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.StandardTaxPercentage) {
		return nil, false
	}
	return o.StandardTaxPercentage, true
}

// HasStandardTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasStandardTaxPercentage() bool {
	if o != nil && !IsNil(o.StandardTaxPercentage) {
		return true
	}

	return false
}

// SetStandardTaxPercentage gets a reference to the given int64 and assigns it to the StandardTaxPercentage field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetStandardTaxPercentage(v int64) {
	o.StandardTaxPercentage = &v
}

// GetVatSupport returns the VatSupport field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetVatSupport() bool {
	if o == nil || IsNil(o.VatSupport) {
		var ret bool
		return ret
	}
	return *o.VatSupport
}

// GetVatSupportOk returns a tuple with the VatSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) GetVatSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.VatSupport) {
		return nil, false
	}
	return o.VatSupport, true
}

// HasVatSupport returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) HasVatSupport() bool {
	if o != nil && !IsNil(o.VatSupport) {
		return true
	}

	return false
}

// SetVatSupport gets a reference to the given bool and assigns it to the VatSupport field.
func (o *UnibeeInternalLogicGatewayRoVatCountryRate) SetVatSupport(v bool) {
	o.VatSupport = &v
}

func (o UnibeeInternalLogicGatewayRoVatCountryRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoVatCountryRate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StandardTaxPercentage) {
		toSerialize["standardTaxPercentage"] = o.StandardTaxPercentage
	}
	if !IsNil(o.VatSupport) {
		toSerialize["vatSupport"] = o.VatSupport
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoVatCountryRate struct {
	value *UnibeeInternalLogicGatewayRoVatCountryRate
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoVatCountryRate) Get() *UnibeeInternalLogicGatewayRoVatCountryRate {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoVatCountryRate) Set(val *UnibeeInternalLogicGatewayRoVatCountryRate) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoVatCountryRate) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoVatCountryRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoVatCountryRate(val *UnibeeInternalLogicGatewayRoVatCountryRate) *NullableUnibeeInternalLogicGatewayRoVatCountryRate {
	return &NullableUnibeeInternalLogicGatewayRoVatCountryRate{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoVatCountryRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoVatCountryRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


