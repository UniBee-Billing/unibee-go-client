/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentNewPaymentRefundReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentNewPaymentRefundReq{}

// UnibeeApiMerchantPaymentNewPaymentRefundReq struct for UnibeeApiMerchantPaymentNewPaymentRefundReq
type UnibeeApiMerchantPaymentNewPaymentRefundReq struct {
	// Currency
	Currency string `json:"currency"`
	// ExternalRefundId
	ExternalRefundId string `json:"externalRefundId"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// PaymentId
	PaymentId string `json:"paymentId"`
	// Reason
	Reason *string `json:"reason,omitempty"`
	// RefundAmount, Cent
	RefundAmount int64 `json:"refundAmount"`
}

type _UnibeeApiMerchantPaymentNewPaymentRefundReq UnibeeApiMerchantPaymentNewPaymentRefundReq

// NewUnibeeApiMerchantPaymentNewPaymentRefundReq instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentNewPaymentRefundReq(currency string, externalRefundId string, paymentId string, refundAmount int64) *UnibeeApiMerchantPaymentNewPaymentRefundReq {
	this := UnibeeApiMerchantPaymentNewPaymentRefundReq{}
	this.Currency = currency
	this.ExternalRefundId = externalRefundId
	this.PaymentId = paymentId
	this.RefundAmount = refundAmount
	return &this
}

// NewUnibeeApiMerchantPaymentNewPaymentRefundReqWithDefaults instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentNewPaymentRefundReqWithDefaults() *UnibeeApiMerchantPaymentNewPaymentRefundReq {
	this := UnibeeApiMerchantPaymentNewPaymentRefundReq{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalRefundId returns the ExternalRefundId field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetExternalRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalRefundId
}

// GetExternalRefundIdOk returns a tuple with the ExternalRefundId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetExternalRefundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalRefundId, true
}

// SetExternalRefundId sets field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetExternalRefundId(v string) {
	o.ExternalRefundId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetReason(v string) {
	o.Reason = &v
}

// GetRefundAmount returns the RefundAmount field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetRefundAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetRefundAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetRefundAmount(v int64) {
	o.RefundAmount = v
}

func (o UnibeeApiMerchantPaymentNewPaymentRefundReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentNewPaymentRefundReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["externalRefundId"] = o.ExternalRefundId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["paymentId"] = o.PaymentId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["refundAmount"] = o.RefundAmount
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"externalRefundId",
		"paymentId",
		"refundAmount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantPaymentNewPaymentRefundReq := _UnibeeApiMerchantPaymentNewPaymentRefundReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentNewPaymentRefundReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentNewPaymentRefundReq(varUnibeeApiMerchantPaymentNewPaymentRefundReq)

	return err
}

type NullableUnibeeApiMerchantPaymentNewPaymentRefundReq struct {
	value *UnibeeApiMerchantPaymentNewPaymentRefundReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) Get() *UnibeeApiMerchantPaymentNewPaymentRefundReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) Set(val *UnibeeApiMerchantPaymentNewPaymentRefundReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentNewPaymentRefundReq(val *UnibeeApiMerchantPaymentNewPaymentRefundReq) *NullableUnibeeApiMerchantPaymentNewPaymentRefundReq {
	return &NullableUnibeeApiMerchantPaymentNewPaymentRefundReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentNewPaymentRefundReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


