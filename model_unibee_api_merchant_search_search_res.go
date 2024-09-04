/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSearchSearchRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSearchSearchRes{}

// UnibeeApiMerchantSearchSearchRes struct for UnibeeApiMerchantSearchSearchRes
type UnibeeApiMerchantSearchSearchRes struct {
	// MatchInvoice
	MatchInvoice []UnibeeApiBeanInvoice `json:"matchInvoice,omitempty"`
	// MatchUserAccounts
	MatchUserAccounts []UnibeeApiBeanDetailUserAccountDetail `json:"matchUserAccounts,omitempty"`
	PrecisionMatchObject *UnibeeApiMerchantSearchPrecisionMatchObject `json:"precisionMatchObject,omitempty"`
}

// NewUnibeeApiMerchantSearchSearchRes instantiates a new UnibeeApiMerchantSearchSearchRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSearchSearchRes() *UnibeeApiMerchantSearchSearchRes {
	this := UnibeeApiMerchantSearchSearchRes{}
	return &this
}

// NewUnibeeApiMerchantSearchSearchResWithDefaults instantiates a new UnibeeApiMerchantSearchSearchRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSearchSearchResWithDefaults() *UnibeeApiMerchantSearchSearchRes {
	this := UnibeeApiMerchantSearchSearchRes{}
	return &this
}

// GetMatchInvoice returns the MatchInvoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchSearchRes) GetMatchInvoice() []UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.MatchInvoice) {
		var ret []UnibeeApiBeanInvoice
		return ret
	}
	return o.MatchInvoice
}

// GetMatchInvoiceOk returns a tuple with the MatchInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchSearchRes) GetMatchInvoiceOk() ([]UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.MatchInvoice) {
		return nil, false
	}
	return o.MatchInvoice, true
}

// HasMatchInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchSearchRes) HasMatchInvoice() bool {
	if o != nil && !IsNil(o.MatchInvoice) {
		return true
	}

	return false
}

// SetMatchInvoice gets a reference to the given []UnibeeApiBeanInvoice and assigns it to the MatchInvoice field.
func (o *UnibeeApiMerchantSearchSearchRes) SetMatchInvoice(v []UnibeeApiBeanInvoice) {
	o.MatchInvoice = v
}

// GetMatchUserAccounts returns the MatchUserAccounts field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchSearchRes) GetMatchUserAccounts() []UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.MatchUserAccounts) {
		var ret []UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return o.MatchUserAccounts
}

// GetMatchUserAccountsOk returns a tuple with the MatchUserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchSearchRes) GetMatchUserAccountsOk() ([]UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.MatchUserAccounts) {
		return nil, false
	}
	return o.MatchUserAccounts, true
}

// HasMatchUserAccounts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchSearchRes) HasMatchUserAccounts() bool {
	if o != nil && !IsNil(o.MatchUserAccounts) {
		return true
	}

	return false
}

// SetMatchUserAccounts gets a reference to the given []UnibeeApiBeanDetailUserAccountDetail and assigns it to the MatchUserAccounts field.
func (o *UnibeeApiMerchantSearchSearchRes) SetMatchUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail) {
	o.MatchUserAccounts = v
}

// GetPrecisionMatchObject returns the PrecisionMatchObject field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchSearchRes) GetPrecisionMatchObject() UnibeeApiMerchantSearchPrecisionMatchObject {
	if o == nil || IsNil(o.PrecisionMatchObject) {
		var ret UnibeeApiMerchantSearchPrecisionMatchObject
		return ret
	}
	return *o.PrecisionMatchObject
}

// GetPrecisionMatchObjectOk returns a tuple with the PrecisionMatchObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchSearchRes) GetPrecisionMatchObjectOk() (*UnibeeApiMerchantSearchPrecisionMatchObject, bool) {
	if o == nil || IsNil(o.PrecisionMatchObject) {
		return nil, false
	}
	return o.PrecisionMatchObject, true
}

// HasPrecisionMatchObject returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchSearchRes) HasPrecisionMatchObject() bool {
	if o != nil && !IsNil(o.PrecisionMatchObject) {
		return true
	}

	return false
}

// SetPrecisionMatchObject gets a reference to the given UnibeeApiMerchantSearchPrecisionMatchObject and assigns it to the PrecisionMatchObject field.
func (o *UnibeeApiMerchantSearchSearchRes) SetPrecisionMatchObject(v UnibeeApiMerchantSearchPrecisionMatchObject) {
	o.PrecisionMatchObject = &v
}

func (o UnibeeApiMerchantSearchSearchRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSearchSearchRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchInvoice) {
		toSerialize["matchInvoice"] = o.MatchInvoice
	}
	if !IsNil(o.MatchUserAccounts) {
		toSerialize["matchUserAccounts"] = o.MatchUserAccounts
	}
	if !IsNil(o.PrecisionMatchObject) {
		toSerialize["precisionMatchObject"] = o.PrecisionMatchObject
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSearchSearchRes struct {
	value *UnibeeApiMerchantSearchSearchRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSearchSearchRes) Get() *UnibeeApiMerchantSearchSearchRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSearchSearchRes) Set(val *UnibeeApiMerchantSearchSearchRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSearchSearchRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSearchSearchRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSearchSearchRes(val *UnibeeApiMerchantSearchSearchRes) *NullableUnibeeApiMerchantSearchSearchRes {
	return &NullableUnibeeApiMerchantSearchSearchRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSearchSearchRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSearchSearchRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


