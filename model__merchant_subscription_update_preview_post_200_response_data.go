/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionUpdatePreviewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionUpdatePreviewPost200ResponseData{}

// MerchantSubscriptionUpdatePreviewPost200ResponseData struct for MerchantSubscriptionUpdatePreviewPost200ResponseData
type MerchantSubscriptionUpdatePreviewPost200ResponseData struct {
	Currency *string `json:"currency,omitempty"`
	Discount *UnibeeApiBeanMerchantDiscountCodeSimplify `json:"discount,omitempty"`
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	Invoice *UnibeeApiBeanInvoiceSimplify `json:"invoice,omitempty"`
	NextPeriodInvoice *UnibeeApiBeanInvoiceSimplify `json:"nextPeriodInvoice,omitempty"`
	OriginAmount *int64 `json:"originAmount,omitempty"`
	ProrationDate *int64 `json:"prorationDate,omitempty"`
	TotalAmount *int64 `json:"totalAmount,omitempty"`
}

// NewMerchantSubscriptionUpdatePreviewPost200ResponseData instantiates a new MerchantSubscriptionUpdatePreviewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionUpdatePreviewPost200ResponseData() *MerchantSubscriptionUpdatePreviewPost200ResponseData {
	this := MerchantSubscriptionUpdatePreviewPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionUpdatePreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUpdatePreviewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionUpdatePreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionUpdatePreviewPost200ResponseData {
	this := MerchantSubscriptionUpdatePreviewPost200ResponseData{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetCurrency(v string) {
	o.Currency = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscount() UnibeeApiBeanMerchantDiscountCodeSimplify {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCodeSimplify
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCodeSimplify, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCodeSimplify and assigns it to the Discount field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetDiscount(v UnibeeApiBeanMerchantDiscountCodeSimplify) {
	o.Discount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoiceSimplify {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoiceSimplify
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoiceSimplify and assigns it to the Invoice field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoiceSimplify) {
	o.Invoice = &v
}

// GetNextPeriodInvoice returns the NextPeriodInvoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetNextPeriodInvoice() UnibeeApiBeanInvoiceSimplify {
	if o == nil || IsNil(o.NextPeriodInvoice) {
		var ret UnibeeApiBeanInvoiceSimplify
		return ret
	}
	return *o.NextPeriodInvoice
}

// GetNextPeriodInvoiceOk returns a tuple with the NextPeriodInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetNextPeriodInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool) {
	if o == nil || IsNil(o.NextPeriodInvoice) {
		return nil, false
	}
	return o.NextPeriodInvoice, true
}

// HasNextPeriodInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasNextPeriodInvoice() bool {
	if o != nil && !IsNil(o.NextPeriodInvoice) {
		return true
	}

	return false
}

// SetNextPeriodInvoice gets a reference to the given UnibeeApiBeanInvoiceSimplify and assigns it to the NextPeriodInvoice field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetNextPeriodInvoice(v UnibeeApiBeanInvoiceSimplify) {
	o.NextPeriodInvoice = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetOriginAmount() int64 {
	if o == nil || IsNil(o.OriginAmount) {
		var ret int64
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetOriginAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given int64 and assigns it to the OriginAmount field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetOriginAmount(v int64) {
	o.OriginAmount = &v
}

// GetProrationDate returns the ProrationDate field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetProrationDate() int64 {
	if o == nil || IsNil(o.ProrationDate) {
		var ret int64
		return ret
	}
	return *o.ProrationDate
}

// GetProrationDateOk returns a tuple with the ProrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetProrationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ProrationDate) {
		return nil, false
	}
	return o.ProrationDate, true
}

// HasProrationDate returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasProrationDate() bool {
	if o != nil && !IsNil(o.ProrationDate) {
		return true
	}

	return false
}

// SetProrationDate gets a reference to the given int64 and assigns it to the ProrationDate field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetProrationDate(v int64) {
	o.ProrationDate = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

func (o MerchantSubscriptionUpdatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionUpdatePreviewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.NextPeriodInvoice) {
		toSerialize["nextPeriodInvoice"] = o.NextPeriodInvoice
	}
	if !IsNil(o.OriginAmount) {
		toSerialize["originAmount"] = o.OriginAmount
	}
	if !IsNil(o.ProrationDate) {
		toSerialize["prorationDate"] = o.ProrationDate
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionUpdatePreviewPost200ResponseData struct {
	value *MerchantSubscriptionUpdatePreviewPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) Get() *MerchantSubscriptionUpdatePreviewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) Set(val *MerchantSubscriptionUpdatePreviewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionUpdatePreviewPost200ResponseData(val *MerchantSubscriptionUpdatePreviewPost200ResponseData) *NullableMerchantSubscriptionUpdatePreviewPost200ResponseData {
	return &NullableMerchantSubscriptionUpdatePreviewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionUpdatePreviewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


