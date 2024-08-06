/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionOnetimeAddonNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionOnetimeAddonNewRes{}

// UnibeeApiMerchantSubscriptionOnetimeAddonNewRes struct for UnibeeApiMerchantSubscriptionOnetimeAddonNewRes
type UnibeeApiMerchantSubscriptionOnetimeAddonNewRes struct {
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
	// if automatic payment is false, Gateway Link will provided that manual payment needed
	Link *string `json:"link,omitempty"`
	// true|false,automatic payment is default behavior for one-time addon purchased, payment will create attach to the purchase, when payment is success, return false, otherwise false
	Paid *bool `json:"paid,omitempty"`
	SubscriptionOnetimeAddon *UnibeeApiBeanSubscriptionOnetimeAddon `json:"subscriptionOnetimeAddon,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonNewRes instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewRes() *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonNewRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonNewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewResWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonNewRes{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscriptionOnetimeAddon returns the SubscriptionOnetimeAddon field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetSubscriptionOnetimeAddon() UnibeeApiBeanSubscriptionOnetimeAddon {
	if o == nil || IsNil(o.SubscriptionOnetimeAddon) {
		var ret UnibeeApiBeanSubscriptionOnetimeAddon
		return ret
	}
	return *o.SubscriptionOnetimeAddon
}

// GetSubscriptionOnetimeAddonOk returns a tuple with the SubscriptionOnetimeAddon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetSubscriptionOnetimeAddonOk() (*UnibeeApiBeanSubscriptionOnetimeAddon, bool) {
	if o == nil || IsNil(o.SubscriptionOnetimeAddon) {
		return nil, false
	}
	return o.SubscriptionOnetimeAddon, true
}

// HasSubscriptionOnetimeAddon returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasSubscriptionOnetimeAddon() bool {
	if o != nil && !IsNil(o.SubscriptionOnetimeAddon) {
		return true
	}

	return false
}

// SetSubscriptionOnetimeAddon gets a reference to the given UnibeeApiBeanSubscriptionOnetimeAddon and assigns it to the SubscriptionOnetimeAddon field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetSubscriptionOnetimeAddon(v UnibeeApiBeanSubscriptionOnetimeAddon) {
	o.SubscriptionOnetimeAddon = &v
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) ToMap() (map[string]interface{}, error) {
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

type NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes struct {
	value *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) Get() *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) Set(val *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes(val *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes {
	return &NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


