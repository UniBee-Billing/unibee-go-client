/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCancelReq{}

// UnibeeApiMerchantSubscriptionCancelReq Cancel subscription immediately, no proration invoice will generate
type UnibeeApiMerchantSubscriptionCancelReq struct {
	// Default false
	// Deprecated
	InvoiceNow *bool `json:"invoiceNow,omitempty"`
	// default product will use if productId not specified and subscriptionId is blank
	ProductId *int64 `json:"productId,omitempty"`
	// Prorate Generate Invoice，Default false
	// Deprecated
	Prorate *bool `json:"prorate,omitempty"`
	// SubscriptionId, id of subscription, either SubscriptionId or UserId needed, The only one active subscription of userId will effect
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// UserId, either SubscriptionId or UserId needed, The only one active subscription will effect if userId provide instead of subscriptionId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionCancelReq instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCancelReq() *UnibeeApiMerchantSubscriptionCancelReq {
	this := UnibeeApiMerchantSubscriptionCancelReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelReq {
	this := UnibeeApiMerchantSubscriptionCancelReq{}
	return &this
}

// GetInvoiceNow returns the InvoiceNow field value if set, zero value otherwise.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNow() bool {
	if o == nil || IsNil(o.InvoiceNow) {
		var ret bool
		return ret
	}
	return *o.InvoiceNow
}

// GetInvoiceNowOk returns a tuple with the InvoiceNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNowOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceNow) {
		return nil, false
	}
	return o.InvoiceNow, true
}

// HasInvoiceNow returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasInvoiceNow() bool {
	if o != nil && !IsNil(o.InvoiceNow) {
		return true
	}

	return false
}

// SetInvoiceNow gets a reference to the given bool and assigns it to the InvoiceNow field.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetInvoiceNow(v bool) {
	o.InvoiceNow = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetProductId(v int64) {
	o.ProductId = &v
}

// GetProrate returns the Prorate field value if set, zero value otherwise.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrate() bool {
	if o == nil || IsNil(o.Prorate) {
		var ret bool
		return ret
	}
	return *o.Prorate
}

// GetProrateOk returns a tuple with the Prorate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrateOk() (*bool, bool) {
	if o == nil || IsNil(o.Prorate) {
		return nil, false
	}
	return o.Prorate, true
}

// HasProrate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasProrate() bool {
	if o != nil && !IsNil(o.Prorate) {
		return true
	}

	return false
}

// SetProrate gets a reference to the given bool and assigns it to the Prorate field.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetProrate(v bool) {
	o.Prorate = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceNow) {
		toSerialize["invoiceNow"] = o.InvoiceNow
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.Prorate) {
		toSerialize["prorate"] = o.Prorate
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionCancelReq struct {
	value *UnibeeApiMerchantSubscriptionCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) Get() *UnibeeApiMerchantSubscriptionCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) Set(val *UnibeeApiMerchantSubscriptionCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCancelReq(val *UnibeeApiMerchantSubscriptionCancelReq) *NullableUnibeeApiMerchantSubscriptionCancelReq {
	return &NullableUnibeeApiMerchantSubscriptionCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


