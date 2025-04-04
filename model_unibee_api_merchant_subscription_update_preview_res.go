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

// checks if the UnibeeApiMerchantSubscriptionUpdatePreviewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdatePreviewRes{}

// UnibeeApiMerchantSubscriptionUpdatePreviewRes struct for UnibeeApiMerchantSubscriptionUpdatePreviewRes
type UnibeeApiMerchantSubscriptionUpdatePreviewRes struct {
	// apply promo credit or not
	ApplyPromoCredit *bool `json:"applyPromoCredit,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Discount *UnibeeApiBeanMerchantDiscountCode `json:"discount,omitempty"`
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	DiscountMessage *string `json:"discountMessage,omitempty"`
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
	NextPeriodInvoice *UnibeeApiBeanInvoice `json:"nextPeriodInvoice,omitempty"`
	OriginAmount *int64 `json:"originAmount,omitempty"`
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

// GetApplyPromoCredit returns the ApplyPromoCredit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetApplyPromoCredit() bool {
	if o == nil || IsNil(o.ApplyPromoCredit) {
		var ret bool
		return ret
	}
	return *o.ApplyPromoCredit
}

// GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetApplyPromoCreditOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyPromoCredit) {
		return nil, false
	}
	return o.ApplyPromoCredit, true
}

// HasApplyPromoCredit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasApplyPromoCredit() bool {
	if o != nil && !IsNil(o.ApplyPromoCredit) {
		return true
	}

	return false
}

// SetApplyPromoCredit gets a reference to the given bool and assigns it to the ApplyPromoCredit field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetApplyPromoCredit(v bool) {
	o.ApplyPromoCredit = &v
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

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the Discount field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode) {
	o.Discount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountMessage returns the DiscountMessage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountMessage() string {
	if o == nil || IsNil(o.DiscountMessage) {
		var ret string
		return ret
	}
	return *o.DiscountMessage
}

// GetDiscountMessageOk returns a tuple with the DiscountMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountMessage) {
		return nil, false
	}
	return o.DiscountMessage, true
}

// HasDiscountMessage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscountMessage() bool {
	if o != nil && !IsNil(o.DiscountMessage) {
		return true
	}

	return false
}

// SetDiscountMessage gets a reference to the given string and assigns it to the DiscountMessage field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscountMessage(v string) {
	o.DiscountMessage = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
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

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetNextPeriodInvoice returns the NextPeriodInvoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.NextPeriodInvoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.NextPeriodInvoice
}

// GetNextPeriodInvoiceOk returns a tuple with the NextPeriodInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
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

// SetNextPeriodInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the NextPeriodInvoice field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetNextPeriodInvoice(v UnibeeApiBeanInvoice) {
	o.NextPeriodInvoice = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetOriginAmount() int64 {
	if o == nil || IsNil(o.OriginAmount) {
		var ret int64
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetOriginAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given int64 and assigns it to the OriginAmount field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetOriginAmount(v int64) {
	o.OriginAmount = &v
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
	if !IsNil(o.ApplyPromoCredit) {
		toSerialize["applyPromoCredit"] = o.ApplyPromoCredit
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountMessage) {
		toSerialize["discountMessage"] = o.DiscountMessage
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


