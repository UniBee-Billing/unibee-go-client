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

// checks if the UnibeeApiMerchantProfileGetLicenseRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileGetLicenseRes{}

// UnibeeApiMerchantProfileGetLicenseRes struct for UnibeeApiMerchantProfileGetLicenseRes
type UnibeeApiMerchantProfileGetLicenseRes struct {
	License *UnibeeInternalLogicMiddlewareLicense `json:"license,omitempty"`
	Merchant *UnibeeApiBeanMerchant `json:"merchant,omitempty"`
}

// NewUnibeeApiMerchantProfileGetLicenseRes instantiates a new UnibeeApiMerchantProfileGetLicenseRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileGetLicenseRes() *UnibeeApiMerchantProfileGetLicenseRes {
	this := UnibeeApiMerchantProfileGetLicenseRes{}
	return &this
}

// NewUnibeeApiMerchantProfileGetLicenseResWithDefaults instantiates a new UnibeeApiMerchantProfileGetLicenseRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileGetLicenseResWithDefaults() *UnibeeApiMerchantProfileGetLicenseRes {
	this := UnibeeApiMerchantProfileGetLicenseRes{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileGetLicenseRes) GetLicense() UnibeeInternalLogicMiddlewareLicense {
	if o == nil || IsNil(o.License) {
		var ret UnibeeInternalLogicMiddlewareLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileGetLicenseRes) GetLicenseOk() (*UnibeeInternalLogicMiddlewareLicense, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileGetLicenseRes) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given UnibeeInternalLogicMiddlewareLicense and assigns it to the License field.
func (o *UnibeeApiMerchantProfileGetLicenseRes) SetLicense(v UnibeeInternalLogicMiddlewareLicense) {
	o.License = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileGetLicenseRes) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *UnibeeApiMerchantProfileGetLicenseRes) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

func (o UnibeeApiMerchantProfileGetLicenseRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileGetLicenseRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProfileGetLicenseRes struct {
	value *UnibeeApiMerchantProfileGetLicenseRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileGetLicenseRes) Get() *UnibeeApiMerchantProfileGetLicenseRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileGetLicenseRes) Set(val *UnibeeApiMerchantProfileGetLicenseRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileGetLicenseRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileGetLicenseRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileGetLicenseRes(val *UnibeeApiMerchantProfileGetLicenseRes) *NullableUnibeeApiMerchantProfileGetLicenseRes {
	return &NullableUnibeeApiMerchantProfileGetLicenseRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileGetLicenseRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileGetLicenseRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


