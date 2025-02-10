/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentNewPaymentRefundRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentNewPaymentRefundRes{}

// UnibeeApiMerchantPaymentNewPaymentRefundRes struct for UnibeeApiMerchantPaymentNewPaymentRefundRes
type UnibeeApiMerchantPaymentNewPaymentRefundRes struct {
	// ExternalRefundId
	ExternalRefundId *string `json:"externalRefundId,omitempty"`
	// PaymentId
	PaymentId *string `json:"paymentId,omitempty"`
	// RefundId
	RefundId *string `json:"refundId,omitempty"`
	// Status,10-create|20-success|30-Failed|40-Reverse
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiMerchantPaymentNewPaymentRefundRes instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentNewPaymentRefundRes() *UnibeeApiMerchantPaymentNewPaymentRefundRes {
	this := UnibeeApiMerchantPaymentNewPaymentRefundRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentNewPaymentRefundResWithDefaults instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentNewPaymentRefundResWithDefaults() *UnibeeApiMerchantPaymentNewPaymentRefundRes {
	this := UnibeeApiMerchantPaymentNewPaymentRefundRes{}
	return &this
}

// GetExternalRefundId returns the ExternalRefundId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetExternalRefundId() string {
	if o == nil || IsNil(o.ExternalRefundId) {
		var ret string
		return ret
	}
	return *o.ExternalRefundId
}

// GetExternalRefundIdOk returns a tuple with the ExternalRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetExternalRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalRefundId) {
		return nil, false
	}
	return o.ExternalRefundId, true
}

// HasExternalRefundId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) HasExternalRefundId() bool {
	if o != nil && !IsNil(o.ExternalRefundId) {
		return true
	}

	return false
}

// SetExternalRefundId gets a reference to the given string and assigns it to the ExternalRefundId field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) SetExternalRefundId(v string) {
	o.ExternalRefundId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) SetRefundId(v string) {
	o.RefundId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundRes) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiMerchantPaymentNewPaymentRefundRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentNewPaymentRefundRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalRefundId) {
		toSerialize["externalRefundId"] = o.ExternalRefundId
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentNewPaymentRefundRes struct {
	value *UnibeeApiMerchantPaymentNewPaymentRefundRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) Get() *UnibeeApiMerchantPaymentNewPaymentRefundRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) Set(val *UnibeeApiMerchantPaymentNewPaymentRefundRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentNewPaymentRefundRes(val *UnibeeApiMerchantPaymentNewPaymentRefundRes) *NullableUnibeeApiMerchantPaymentNewPaymentRefundRes {
	return &NullableUnibeeApiMerchantPaymentNewPaymentRefundRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


