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

// checks if the UnibeeApiMerchantCreditDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditDetailRes{}

// UnibeeApiMerchantCreditDetailRes struct for UnibeeApiMerchantCreditDetailRes
type UnibeeApiMerchantCreditDetailRes struct {
	CreditAccount *UnibeeApiBeanDetailCreditAccountDetail `json:"creditAccount,omitempty"`
	// Credit Transaction List
	CreditTransactions []UnibeeApiBeanCreditTransaction `json:"creditTransactions,omitempty"`
}

// NewUnibeeApiMerchantCreditDetailRes instantiates a new UnibeeApiMerchantCreditDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditDetailRes() *UnibeeApiMerchantCreditDetailRes {
	this := UnibeeApiMerchantCreditDetailRes{}
	return &this
}

// NewUnibeeApiMerchantCreditDetailResWithDefaults instantiates a new UnibeeApiMerchantCreditDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditDetailResWithDefaults() *UnibeeApiMerchantCreditDetailRes {
	this := UnibeeApiMerchantCreditDetailRes{}
	return &this
}

// GetCreditAccount returns the CreditAccount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditDetailRes) GetCreditAccount() UnibeeApiBeanDetailCreditAccountDetail {
	if o == nil || IsNil(o.CreditAccount) {
		var ret UnibeeApiBeanDetailCreditAccountDetail
		return ret
	}
	return *o.CreditAccount
}

// GetCreditAccountOk returns a tuple with the CreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditDetailRes) GetCreditAccountOk() (*UnibeeApiBeanDetailCreditAccountDetail, bool) {
	if o == nil || IsNil(o.CreditAccount) {
		return nil, false
	}
	return o.CreditAccount, true
}

// HasCreditAccount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditDetailRes) HasCreditAccount() bool {
	if o != nil && !IsNil(o.CreditAccount) {
		return true
	}

	return false
}

// SetCreditAccount gets a reference to the given UnibeeApiBeanDetailCreditAccountDetail and assigns it to the CreditAccount field.
func (o *UnibeeApiMerchantCreditDetailRes) SetCreditAccount(v UnibeeApiBeanDetailCreditAccountDetail) {
	o.CreditAccount = &v
}

// GetCreditTransactions returns the CreditTransactions field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditDetailRes) GetCreditTransactions() []UnibeeApiBeanCreditTransaction {
	if o == nil || IsNil(o.CreditTransactions) {
		var ret []UnibeeApiBeanCreditTransaction
		return ret
	}
	return o.CreditTransactions
}

// GetCreditTransactionsOk returns a tuple with the CreditTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditDetailRes) GetCreditTransactionsOk() ([]UnibeeApiBeanCreditTransaction, bool) {
	if o == nil || IsNil(o.CreditTransactions) {
		return nil, false
	}
	return o.CreditTransactions, true
}

// HasCreditTransactions returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditDetailRes) HasCreditTransactions() bool {
	if o != nil && !IsNil(o.CreditTransactions) {
		return true
	}

	return false
}

// SetCreditTransactions gets a reference to the given []UnibeeApiBeanCreditTransaction and assigns it to the CreditTransactions field.
func (o *UnibeeApiMerchantCreditDetailRes) SetCreditTransactions(v []UnibeeApiBeanCreditTransaction) {
	o.CreditTransactions = v
}

func (o UnibeeApiMerchantCreditDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditAccount) {
		toSerialize["creditAccount"] = o.CreditAccount
	}
	if !IsNil(o.CreditTransactions) {
		toSerialize["creditTransactions"] = o.CreditTransactions
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditDetailRes struct {
	value *UnibeeApiMerchantCreditDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditDetailRes) Get() *UnibeeApiMerchantCreditDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditDetailRes) Set(val *UnibeeApiMerchantCreditDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditDetailRes(val *UnibeeApiMerchantCreditDetailRes) *NullableUnibeeApiMerchantCreditDetailRes {
	return &NullableUnibeeApiMerchantCreditDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


