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

// checks if the UnibeeApiMerchantSubscriptionUpdatePreviewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdatePreviewRes{}

// UnibeeApiMerchantSubscriptionUpdatePreviewRes struct for UnibeeApiMerchantSubscriptionUpdatePreviewRes
type UnibeeApiMerchantSubscriptionUpdatePreviewRes struct {
	Currency *string `json:"currency,omitempty"`
	Invoice *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify `json:"invoice,omitempty"`
	NextPeriodInvoice *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify `json:"nextPeriodInvoice,omitempty"`
	ProrationDate *int64 `json:"prorationDate,omitempty"`
	TotalAmount *int64 `json:"totalAmount,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionUpdatePreviewRes instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUpdatePreviewRes() *UnibeeApiMerchantSubscriptionUpdatePreviewRes {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionUpdatePreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUpdatePreviewResWithDefaults() *UnibeeApiMerchantSubscriptionUpdatePreviewRes {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewRes{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetCurrency(v string) {
	o.Currency = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoice() UnibeeInternalLogicGatewayRoInvoiceDetailSimplify {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeInternalLogicGatewayRoInvoiceDetailSimplify
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoiceOk() (*UnibeeInternalLogicGatewayRoInvoiceDetailSimplify, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeInternalLogicGatewayRoInvoiceDetailSimplify and assigns it to the Invoice field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetInvoice(v UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) {
	o.Invoice = &v
}

// GetNextPeriodInvoice returns the NextPeriodInvoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoice() UnibeeInternalLogicGatewayRoInvoiceDetailSimplify {
	if o == nil || IsNil(o.NextPeriodInvoice) {
		var ret UnibeeInternalLogicGatewayRoInvoiceDetailSimplify
		return ret
	}
	return *o.NextPeriodInvoice
}

// GetNextPeriodInvoiceOk returns a tuple with the NextPeriodInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoiceOk() (*UnibeeInternalLogicGatewayRoInvoiceDetailSimplify, bool) {
	if o == nil || IsNil(o.NextPeriodInvoice) {
		return nil, false
	}
	return o.NextPeriodInvoice, true
}

// HasNextPeriodInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasNextPeriodInvoice() bool {
	if o != nil && !IsNil(o.NextPeriodInvoice) {
		return true
	}

	return false
}

// SetNextPeriodInvoice gets a reference to the given UnibeeInternalLogicGatewayRoInvoiceDetailSimplify and assigns it to the NextPeriodInvoice field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetNextPeriodInvoice(v UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) {
	o.NextPeriodInvoice = &v
}

// GetProrationDate returns the ProrationDate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetProrationDate() int64 {
	if o == nil || IsNil(o.ProrationDate) {
		var ret int64
		return ret
	}
	return *o.ProrationDate
}

// GetProrationDateOk returns a tuple with the ProrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetProrationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ProrationDate) {
		return nil, false
	}
	return o.ProrationDate, true
}

// HasProrationDate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasProrationDate() bool {
	if o != nil && !IsNil(o.ProrationDate) {
		return true
	}

	return false
}

// SetProrationDate gets a reference to the given int64 and assigns it to the ProrationDate field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetProrationDate(v int64) {
	o.ProrationDate = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.NextPeriodInvoice) {
		toSerialize["nextPeriodInvoice"] = o.NextPeriodInvoice
	}
	if !IsNil(o.ProrationDate) {
		toSerialize["prorationDate"] = o.ProrationDate
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes struct {
	value *UnibeeApiMerchantSubscriptionUpdatePreviewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) Get() *UnibeeApiMerchantSubscriptionUpdatePreviewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) Set(val *UnibeeApiMerchantSubscriptionUpdatePreviewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUpdatePreviewRes(val *UnibeeApiMerchantSubscriptionUpdatePreviewRes) *NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes {
	return &NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


