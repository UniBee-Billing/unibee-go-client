/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantBalanceUserBalanceQueryGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantBalanceUserBalanceQueryGet200ResponseData{}

// MerchantBalanceUserBalanceQueryGet200ResponseData struct for MerchantBalanceUserBalanceQueryGet200ResponseData
type MerchantBalanceUserBalanceQueryGet200ResponseData struct {
	Balance *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"balance,omitempty"`
	CashBalance []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"cashBalance,omitempty"`
	InvoiceCreditBalance []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"invoiceCreditBalance,omitempty"`
}

// NewMerchantBalanceUserBalanceQueryGet200ResponseData instantiates a new MerchantBalanceUserBalanceQueryGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantBalanceUserBalanceQueryGet200ResponseData() *MerchantBalanceUserBalanceQueryGet200ResponseData {
	this := MerchantBalanceUserBalanceQueryGet200ResponseData{}
	return &this
}

// NewMerchantBalanceUserBalanceQueryGet200ResponseDataWithDefaults instantiates a new MerchantBalanceUserBalanceQueryGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantBalanceUserBalanceQueryGet200ResponseDataWithDefaults() *MerchantBalanceUserBalanceQueryGet200ResponseData {
	this := MerchantBalanceUserBalanceQueryGet200ResponseData{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetBalance() UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Balance) {
		var ret UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetBalanceOk() (*UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Balance field.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) SetBalance(v UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Balance = &v
}

// GetCashBalance returns the CashBalance field value if set, zero value otherwise.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetCashBalance() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.CashBalance) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.CashBalance
}

// GetCashBalanceOk returns a tuple with the CashBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetCashBalanceOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.CashBalance) {
		return nil, false
	}
	return o.CashBalance, true
}

// HasCashBalance returns a boolean if a field has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) HasCashBalance() bool {
	if o != nil && !IsNil(o.CashBalance) {
		return true
	}

	return false
}

// SetCashBalance gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the CashBalance field.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) SetCashBalance(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.CashBalance = v
}

// GetInvoiceCreditBalance returns the InvoiceCreditBalance field value if set, zero value otherwise.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetInvoiceCreditBalance() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.InvoiceCreditBalance) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.InvoiceCreditBalance
}

// GetInvoiceCreditBalanceOk returns a tuple with the InvoiceCreditBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) GetInvoiceCreditBalanceOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.InvoiceCreditBalance) {
		return nil, false
	}
	return o.InvoiceCreditBalance, true
}

// HasInvoiceCreditBalance returns a boolean if a field has been set.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) HasInvoiceCreditBalance() bool {
	if o != nil && !IsNil(o.InvoiceCreditBalance) {
		return true
	}

	return false
}

// SetInvoiceCreditBalance gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the InvoiceCreditBalance field.
func (o *MerchantBalanceUserBalanceQueryGet200ResponseData) SetInvoiceCreditBalance(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.InvoiceCreditBalance = v
}

func (o MerchantBalanceUserBalanceQueryGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantBalanceUserBalanceQueryGet200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantBalanceUserBalanceQueryGet200ResponseData struct {
	value *MerchantBalanceUserBalanceQueryGet200ResponseData
	isSet bool
}

func (v NullableMerchantBalanceUserBalanceQueryGet200ResponseData) Get() *MerchantBalanceUserBalanceQueryGet200ResponseData {
	return v.value
}

func (v *NullableMerchantBalanceUserBalanceQueryGet200ResponseData) Set(val *MerchantBalanceUserBalanceQueryGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantBalanceUserBalanceQueryGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantBalanceUserBalanceQueryGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantBalanceUserBalanceQueryGet200ResponseData(val *MerchantBalanceUserBalanceQueryGet200ResponseData) *NullableMerchantBalanceUserBalanceQueryGet200ResponseData {
	return &NullableMerchantBalanceUserBalanceQueryGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantBalanceUserBalanceQueryGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantBalanceUserBalanceQueryGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


