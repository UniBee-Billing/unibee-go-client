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

// checks if the MerchantCreditCreditAccountListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditCreditAccountListGet200ResponseData{}

// MerchantCreditCreditAccountListGet200ResponseData struct for MerchantCreditCreditAccountListGet200ResponseData
type MerchantCreditCreditAccountListGet200ResponseData struct {
	// Credit Account List
	CreditAccounts []UnibeeApiBeanDetailCreditAccountDetail `json:"creditAccounts,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantCreditCreditAccountListGet200ResponseData instantiates a new MerchantCreditCreditAccountListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditCreditAccountListGet200ResponseData() *MerchantCreditCreditAccountListGet200ResponseData {
	this := MerchantCreditCreditAccountListGet200ResponseData{}
	return &this
}

// NewMerchantCreditCreditAccountListGet200ResponseDataWithDefaults instantiates a new MerchantCreditCreditAccountListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditCreditAccountListGet200ResponseDataWithDefaults() *MerchantCreditCreditAccountListGet200ResponseData {
	this := MerchantCreditCreditAccountListGet200ResponseData{}
	return &this
}

// GetCreditAccounts returns the CreditAccounts field value if set, zero value otherwise.
func (o *MerchantCreditCreditAccountListGet200ResponseData) GetCreditAccounts() []UnibeeApiBeanDetailCreditAccountDetail {
	if o == nil || IsNil(o.CreditAccounts) {
		var ret []UnibeeApiBeanDetailCreditAccountDetail
		return ret
	}
	return o.CreditAccounts
}

// GetCreditAccountsOk returns a tuple with the CreditAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditCreditAccountListGet200ResponseData) GetCreditAccountsOk() ([]UnibeeApiBeanDetailCreditAccountDetail, bool) {
	if o == nil || IsNil(o.CreditAccounts) {
		return nil, false
	}
	return o.CreditAccounts, true
}

// HasCreditAccounts returns a boolean if a field has been set.
func (o *MerchantCreditCreditAccountListGet200ResponseData) HasCreditAccounts() bool {
	if o != nil && !IsNil(o.CreditAccounts) {
		return true
	}

	return false
}

// SetCreditAccounts gets a reference to the given []UnibeeApiBeanDetailCreditAccountDetail and assigns it to the CreditAccounts field.
func (o *MerchantCreditCreditAccountListGet200ResponseData) SetCreditAccounts(v []UnibeeApiBeanDetailCreditAccountDetail) {
	o.CreditAccounts = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantCreditCreditAccountListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditCreditAccountListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantCreditCreditAccountListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantCreditCreditAccountListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantCreditCreditAccountListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditCreditAccountListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditAccounts) {
		toSerialize["creditAccounts"] = o.CreditAccounts
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantCreditCreditAccountListGet200ResponseData struct {
	value *MerchantCreditCreditAccountListGet200ResponseData
	isSet bool
}

func (v NullableMerchantCreditCreditAccountListGet200ResponseData) Get() *MerchantCreditCreditAccountListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditCreditAccountListGet200ResponseData) Set(val *MerchantCreditCreditAccountListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditCreditAccountListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditCreditAccountListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditCreditAccountListGet200ResponseData(val *MerchantCreditCreditAccountListGet200ResponseData) *NullableMerchantCreditCreditAccountListGet200ResponseData {
	return &NullableMerchantCreditCreditAccountListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditCreditAccountListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditCreditAccountListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


