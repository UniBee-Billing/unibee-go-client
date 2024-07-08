/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantBalanceUserDetailQueryRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantBalanceUserDetailQueryRes{}

// UnibeeApiMerchantBalanceUserDetailQueryRes struct for UnibeeApiMerchantBalanceUserDetailQueryRes
type UnibeeApiMerchantBalanceUserDetailQueryRes struct {
	Balance *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"balance,omitempty"`
	CashBalance []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"cashBalance,omitempty"`
	InvoiceCreditBalance []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"invoiceCreditBalance,omitempty"`
}

// NewUnibeeApiMerchantBalanceUserDetailQueryRes instantiates a new UnibeeApiMerchantBalanceUserDetailQueryRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantBalanceUserDetailQueryRes() *UnibeeApiMerchantBalanceUserDetailQueryRes {
	this := UnibeeApiMerchantBalanceUserDetailQueryRes{}
	return &this
}

// NewUnibeeApiMerchantBalanceUserDetailQueryResWithDefaults instantiates a new UnibeeApiMerchantBalanceUserDetailQueryRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantBalanceUserDetailQueryResWithDefaults() *UnibeeApiMerchantBalanceUserDetailQueryRes {
	this := UnibeeApiMerchantBalanceUserDetailQueryRes{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetBalance() UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Balance) {
		var ret UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetBalanceOk() (*UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Balance field.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) SetBalance(v UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Balance = &v
}

// GetCashBalance returns the CashBalance field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetCashBalance() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.CashBalance) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.CashBalance
}

// GetCashBalanceOk returns a tuple with the CashBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetCashBalanceOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.CashBalance) {
		return nil, false
	}
	return o.CashBalance, true
}

// HasCashBalance returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) HasCashBalance() bool {
	if o != nil && !IsNil(o.CashBalance) {
		return true
	}

	return false
}

// SetCashBalance gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the CashBalance field.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) SetCashBalance(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.CashBalance = v
}

// GetInvoiceCreditBalance returns the InvoiceCreditBalance field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetInvoiceCreditBalance() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.InvoiceCreditBalance) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.InvoiceCreditBalance
}

// GetInvoiceCreditBalanceOk returns a tuple with the InvoiceCreditBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) GetInvoiceCreditBalanceOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.InvoiceCreditBalance) {
		return nil, false
	}
	return o.InvoiceCreditBalance, true
}

// HasInvoiceCreditBalance returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) HasInvoiceCreditBalance() bool {
	if o != nil && !IsNil(o.InvoiceCreditBalance) {
		return true
	}

	return false
}

// SetInvoiceCreditBalance gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the InvoiceCreditBalance field.
func (o *UnibeeApiMerchantBalanceUserDetailQueryRes) SetInvoiceCreditBalance(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.InvoiceCreditBalance = v
}

func (o UnibeeApiMerchantBalanceUserDetailQueryRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantBalanceUserDetailQueryRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.CashBalance) {
		toSerialize["cashBalance"] = o.CashBalance
	}
	if !IsNil(o.InvoiceCreditBalance) {
		toSerialize["invoiceCreditBalance"] = o.InvoiceCreditBalance
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantBalanceUserDetailQueryRes struct {
	value *UnibeeApiMerchantBalanceUserDetailQueryRes
	isSet bool
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryRes) Get() *UnibeeApiMerchantBalanceUserDetailQueryRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryRes) Set(val *UnibeeApiMerchantBalanceUserDetailQueryRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantBalanceUserDetailQueryRes(val *UnibeeApiMerchantBalanceUserDetailQueryRes) *NullableUnibeeApiMerchantBalanceUserDetailQueryRes {
	return &NullableUnibeeApiMerchantBalanceUserDetailQueryRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


