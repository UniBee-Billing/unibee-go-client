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

// checks if the UnibeeApiMerchantCreditCreditTransactionListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditCreditTransactionListRes{}

// UnibeeApiMerchantCreditCreditTransactionListRes struct for UnibeeApiMerchantCreditCreditTransactionListRes
type UnibeeApiMerchantCreditCreditTransactionListRes struct {
	// Credit Transaction List
	CreditTransactions []UnibeeApiBeanDetailCreditTransactionDetail `json:"creditTransactions,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantCreditCreditTransactionListRes instantiates a new UnibeeApiMerchantCreditCreditTransactionListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditCreditTransactionListRes() *UnibeeApiMerchantCreditCreditTransactionListRes {
	this := UnibeeApiMerchantCreditCreditTransactionListRes{}
	return &this
}

// NewUnibeeApiMerchantCreditCreditTransactionListResWithDefaults instantiates a new UnibeeApiMerchantCreditCreditTransactionListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditCreditTransactionListResWithDefaults() *UnibeeApiMerchantCreditCreditTransactionListRes {
	this := UnibeeApiMerchantCreditCreditTransactionListRes{}
	return &this
}

// GetCreditTransactions returns the CreditTransactions field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) GetCreditTransactions() []UnibeeApiBeanDetailCreditTransactionDetail {
	if o == nil || IsNil(o.CreditTransactions) {
		var ret []UnibeeApiBeanDetailCreditTransactionDetail
		return ret
	}
	return o.CreditTransactions
}

// GetCreditTransactionsOk returns a tuple with the CreditTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) GetCreditTransactionsOk() ([]UnibeeApiBeanDetailCreditTransactionDetail, bool) {
	if o == nil || IsNil(o.CreditTransactions) {
		return nil, false
	}
	return o.CreditTransactions, true
}

// HasCreditTransactions returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) HasCreditTransactions() bool {
	if o != nil && !IsNil(o.CreditTransactions) {
		return true
	}

	return false
}

// SetCreditTransactions gets a reference to the given []UnibeeApiBeanDetailCreditTransactionDetail and assigns it to the CreditTransactions field.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) SetCreditTransactions(v []UnibeeApiBeanDetailCreditTransactionDetail) {
	o.CreditTransactions = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantCreditCreditTransactionListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantCreditCreditTransactionListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditCreditTransactionListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditTransactions) {
		toSerialize["creditTransactions"] = o.CreditTransactions
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditCreditTransactionListRes struct {
	value *UnibeeApiMerchantCreditCreditTransactionListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListRes) Get() *UnibeeApiMerchantCreditCreditTransactionListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListRes) Set(val *UnibeeApiMerchantCreditCreditTransactionListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditCreditTransactionListRes(val *UnibeeApiMerchantCreditCreditTransactionListRes) *NullableUnibeeApiMerchantCreditCreditTransactionListRes {
	return &NullableUnibeeApiMerchantCreditCreditTransactionListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


