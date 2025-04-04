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

// checks if the MerchantCreditDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditDetailGet200ResponseData{}

// MerchantCreditDetailGet200ResponseData struct for MerchantCreditDetailGet200ResponseData
type MerchantCreditDetailGet200ResponseData struct {
	CreditAccount *UnibeeApiBeanDetailCreditAccountDetail `json:"creditAccount,omitempty"`
	// Credit Transaction List
	CreditTransactions []UnibeeApiBeanCreditTransaction `json:"creditTransactions,omitempty"`
}

// NewMerchantCreditDetailGet200ResponseData instantiates a new MerchantCreditDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditDetailGet200ResponseData() *MerchantCreditDetailGet200ResponseData {
	this := MerchantCreditDetailGet200ResponseData{}
	return &this
}

// NewMerchantCreditDetailGet200ResponseDataWithDefaults instantiates a new MerchantCreditDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditDetailGet200ResponseDataWithDefaults() *MerchantCreditDetailGet200ResponseData {
	this := MerchantCreditDetailGet200ResponseData{}
	return &this
}

// GetCreditAccount returns the CreditAccount field value if set, zero value otherwise.
func (o *MerchantCreditDetailGet200ResponseData) GetCreditAccount() UnibeeApiBeanDetailCreditAccountDetail {
	if o == nil || IsNil(o.CreditAccount) {
		var ret UnibeeApiBeanDetailCreditAccountDetail
		return ret
	}
	return *o.CreditAccount
}

// GetCreditAccountOk returns a tuple with the CreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditDetailGet200ResponseData) GetCreditAccountOk() (*UnibeeApiBeanDetailCreditAccountDetail, bool) {
	if o == nil || IsNil(o.CreditAccount) {
		return nil, false
	}
	return o.CreditAccount, true
}

// HasCreditAccount returns a boolean if a field has been set.
func (o *MerchantCreditDetailGet200ResponseData) HasCreditAccount() bool {
	if o != nil && !IsNil(o.CreditAccount) {
		return true
	}

	return false
}

// SetCreditAccount gets a reference to the given UnibeeApiBeanDetailCreditAccountDetail and assigns it to the CreditAccount field.
func (o *MerchantCreditDetailGet200ResponseData) SetCreditAccount(v UnibeeApiBeanDetailCreditAccountDetail) {
	o.CreditAccount = &v
}

// GetCreditTransactions returns the CreditTransactions field value if set, zero value otherwise.
func (o *MerchantCreditDetailGet200ResponseData) GetCreditTransactions() []UnibeeApiBeanCreditTransaction {
	if o == nil || IsNil(o.CreditTransactions) {
		var ret []UnibeeApiBeanCreditTransaction
		return ret
	}
	return o.CreditTransactions
}

// GetCreditTransactionsOk returns a tuple with the CreditTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditDetailGet200ResponseData) GetCreditTransactionsOk() ([]UnibeeApiBeanCreditTransaction, bool) {
	if o == nil || IsNil(o.CreditTransactions) {
		return nil, false
	}
	return o.CreditTransactions, true
}

// HasCreditTransactions returns a boolean if a field has been set.
func (o *MerchantCreditDetailGet200ResponseData) HasCreditTransactions() bool {
	if o != nil && !IsNil(o.CreditTransactions) {
		return true
	}

	return false
}

// SetCreditTransactions gets a reference to the given []UnibeeApiBeanCreditTransaction and assigns it to the CreditTransactions field.
func (o *MerchantCreditDetailGet200ResponseData) SetCreditTransactions(v []UnibeeApiBeanCreditTransaction) {
	o.CreditTransactions = v
}

func (o MerchantCreditDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditAccount) {
		toSerialize["creditAccount"] = o.CreditAccount
	}
	if !IsNil(o.CreditTransactions) {
		toSerialize["creditTransactions"] = o.CreditTransactions
	}
	return toSerialize, nil
}

type NullableMerchantCreditDetailGet200ResponseData struct {
	value *MerchantCreditDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantCreditDetailGet200ResponseData) Get() *MerchantCreditDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditDetailGet200ResponseData) Set(val *MerchantCreditDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditDetailGet200ResponseData(val *MerchantCreditDetailGet200ResponseData) *NullableMerchantCreditDetailGet200ResponseData {
	return &NullableMerchantCreditDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


