/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData{}

// MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData struct for MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData
type MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData struct {
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
	// if automatic payment is false, Gateway Link will provided that manual payment needed
	Link *string `json:"link,omitempty"`
	// true|false,automatic payment is default behavior for one-time addon purchased, payment will create attach to the purchase, when payment is success, return false, otherwise false
	Paid *bool `json:"paid,omitempty"`
	SubscriptionOnetimeAddon *UnibeeApiBeanSubscriptionOnetimeAddon `json:"subscriptionOnetimeAddon,omitempty"`
}

// NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData instantiates a new MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData() *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData {
	this := MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseDataWithDefaults() *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData {
	this := MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscriptionOnetimeAddon returns the SubscriptionOnetimeAddon field value if set, zero value otherwise.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetSubscriptionOnetimeAddon() UnibeeApiBeanSubscriptionOnetimeAddon {
	if o == nil || IsNil(o.SubscriptionOnetimeAddon) {
		var ret UnibeeApiBeanSubscriptionOnetimeAddon
		return ret
	}
	return *o.SubscriptionOnetimeAddon
}

// GetSubscriptionOnetimeAddonOk returns a tuple with the SubscriptionOnetimeAddon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetSubscriptionOnetimeAddonOk() (*UnibeeApiBeanSubscriptionOnetimeAddon, bool) {
	if o == nil || IsNil(o.SubscriptionOnetimeAddon) {
		return nil, false
	}
	return o.SubscriptionOnetimeAddon, true
}

// HasSubscriptionOnetimeAddon returns a boolean if a field has been set.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasSubscriptionOnetimeAddon() bool {
	if o != nil && !IsNil(o.SubscriptionOnetimeAddon) {
		return true
	}

	return false
}

// SetSubscriptionOnetimeAddon gets a reference to the given UnibeeApiBeanSubscriptionOnetimeAddon and assigns it to the SubscriptionOnetimeAddon field.
func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetSubscriptionOnetimeAddon(v UnibeeApiBeanSubscriptionOnetimeAddon) {
	o.SubscriptionOnetimeAddon = &v
}

func (o MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if !IsNil(o.SubscriptionOnetimeAddon) {
		toSerialize["subscriptionOnetimeAddon"] = o.SubscriptionOnetimeAddon
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData struct {
	value *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) Get() *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) Set(val *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData(val *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) *NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData {
	return &NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


