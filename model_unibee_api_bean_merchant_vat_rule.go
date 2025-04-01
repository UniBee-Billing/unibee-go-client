/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantVatRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantVatRule{}

// UnibeeApiBeanMerchantVatRule struct for UnibeeApiBeanMerchantVatRule
type UnibeeApiBeanMerchantVatRule struct {
	GatewayNames *string `json:"gatewayNames,omitempty"`
	IgnoreVatNumber *bool `json:"ignoreVatNumber,omitempty"`
	TaxPercentage *int32 `json:"taxPercentage,omitempty"`
	ValidCountryCodes *string `json:"validCountryCodes,omitempty"`
}

// NewUnibeeApiBeanMerchantVatRule instantiates a new UnibeeApiBeanMerchantVatRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantVatRule() *UnibeeApiBeanMerchantVatRule {
	this := UnibeeApiBeanMerchantVatRule{}
	return &this
}

// NewUnibeeApiBeanMerchantVatRuleWithDefaults instantiates a new UnibeeApiBeanMerchantVatRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantVatRuleWithDefaults() *UnibeeApiBeanMerchantVatRule {
	this := UnibeeApiBeanMerchantVatRule{}
	return &this
}

// GetGatewayNames returns the GatewayNames field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantVatRule) GetGatewayNames() string {
	if o == nil || IsNil(o.GatewayNames) {
		var ret string
		return ret
	}
	return *o.GatewayNames
}

// GetGatewayNamesOk returns a tuple with the GatewayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantVatRule) GetGatewayNamesOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayNames) {
		return nil, false
	}
	return o.GatewayNames, true
}

// HasGatewayNames returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantVatRule) HasGatewayNames() bool {
	if o != nil && !IsNil(o.GatewayNames) {
		return true
	}

	return false
}

// SetGatewayNames gets a reference to the given string and assigns it to the GatewayNames field.
func (o *UnibeeApiBeanMerchantVatRule) SetGatewayNames(v string) {
	o.GatewayNames = &v
}

// GetIgnoreVatNumber returns the IgnoreVatNumber field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantVatRule) GetIgnoreVatNumber() bool {
	if o == nil || IsNil(o.IgnoreVatNumber) {
		var ret bool
		return ret
	}
	return *o.IgnoreVatNumber
}

// GetIgnoreVatNumberOk returns a tuple with the IgnoreVatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantVatRule) GetIgnoreVatNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreVatNumber) {
		return nil, false
	}
	return o.IgnoreVatNumber, true
}

// HasIgnoreVatNumber returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantVatRule) HasIgnoreVatNumber() bool {
	if o != nil && !IsNil(o.IgnoreVatNumber) {
		return true
	}

	return false
}

// SetIgnoreVatNumber gets a reference to the given bool and assigns it to the IgnoreVatNumber field.
func (o *UnibeeApiBeanMerchantVatRule) SetIgnoreVatNumber(v bool) {
	o.IgnoreVatNumber = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantVatRule) GetTaxPercentage() int32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantVatRule) GetTaxPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantVatRule) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int32 and assigns it to the TaxPercentage field.
func (o *UnibeeApiBeanMerchantVatRule) SetTaxPercentage(v int32) {
	o.TaxPercentage = &v
}

// GetValidCountryCodes returns the ValidCountryCodes field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantVatRule) GetValidCountryCodes() string {
	if o == nil || IsNil(o.ValidCountryCodes) {
		var ret string
		return ret
	}
	return *o.ValidCountryCodes
}

// GetValidCountryCodesOk returns a tuple with the ValidCountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantVatRule) GetValidCountryCodesOk() (*string, bool) {
	if o == nil || IsNil(o.ValidCountryCodes) {
		return nil, false
	}
	return o.ValidCountryCodes, true
}

// HasValidCountryCodes returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantVatRule) HasValidCountryCodes() bool {
	if o != nil && !IsNil(o.ValidCountryCodes) {
		return true
	}

	return false
}

// SetValidCountryCodes gets a reference to the given string and assigns it to the ValidCountryCodes field.
func (o *UnibeeApiBeanMerchantVatRule) SetValidCountryCodes(v string) {
	o.ValidCountryCodes = &v
}

func (o UnibeeApiBeanMerchantVatRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantVatRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayNames) {
		toSerialize["gatewayNames"] = o.GatewayNames
	}
	if !IsNil(o.IgnoreVatNumber) {
		toSerialize["ignoreVatNumber"] = o.IgnoreVatNumber
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.ValidCountryCodes) {
		toSerialize["validCountryCodes"] = o.ValidCountryCodes
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantVatRule struct {
	value *UnibeeApiBeanMerchantVatRule
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantVatRule) Get() *UnibeeApiBeanMerchantVatRule {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantVatRule) Set(val *UnibeeApiBeanMerchantVatRule) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantVatRule) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantVatRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantVatRule(val *UnibeeApiBeanMerchantVatRule) *NullableUnibeeApiBeanMerchantVatRule {
	return &NullableUnibeeApiBeanMerchantVatRule{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantVatRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantVatRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


