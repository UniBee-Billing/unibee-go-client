/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceRefundReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceRefundReq{}

// UnibeeApiMerchantInvoiceRefundReq Create payment refund for paid invoice
type UnibeeApiMerchantInvoiceRefundReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// The reason of refund
	Reason string `json:"reason"`
	// The amount of refund
	RefundAmount int64 `json:"refundAmount"`
	// The out refund number
	RefundNo *string `json:"refundNo,omitempty"`
}

type _UnibeeApiMerchantInvoiceRefundReq UnibeeApiMerchantInvoiceRefundReq

// NewUnibeeApiMerchantInvoiceRefundReq instantiates a new UnibeeApiMerchantInvoiceRefundReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceRefundReq(invoiceId string, reason string, refundAmount int64) *UnibeeApiMerchantInvoiceRefundReq {
	this := UnibeeApiMerchantInvoiceRefundReq{}
	this.InvoiceId = invoiceId
	this.Reason = reason
	this.RefundAmount = refundAmount
	return &this
}

// NewUnibeeApiMerchantInvoiceRefundReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceRefundReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceRefundReqWithDefaults() *UnibeeApiMerchantInvoiceRefundReq {
	this := UnibeeApiMerchantInvoiceRefundReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceRefundReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceRefundReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceRefundReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReason returns the Reason field value
func (o *UnibeeApiMerchantInvoiceRefundReq) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceRefundReq) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *UnibeeApiMerchantInvoiceRefundReq) SetReason(v string) {
	o.Reason = v
}

// GetRefundAmount returns the RefundAmount field value
func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *UnibeeApiMerchantInvoiceRefundReq) SetRefundAmount(v int64) {
	o.RefundAmount = v
}

// GetRefundNo returns the RefundNo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundNo() string {
	if o == nil || IsNil(o.RefundNo) {
		var ret string
		return ret
	}
	return *o.RefundNo
}

// GetRefundNoOk returns a tuple with the RefundNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundNoOk() (*string, bool) {
	if o == nil || IsNil(o.RefundNo) {
		return nil, false
	}
	return o.RefundNo, true
}

// HasRefundNo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceRefundReq) HasRefundNo() bool {
	if o != nil && !IsNil(o.RefundNo) {
		return true
	}

	return false
}

// SetRefundNo gets a reference to the given string and assigns it to the RefundNo field.
func (o *UnibeeApiMerchantInvoiceRefundReq) SetRefundNo(v string) {
	o.RefundNo = &v
}

func (o UnibeeApiMerchantInvoiceRefundReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceRefundReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	toSerialize["reason"] = o.Reason
	toSerialize["refundAmount"] = o.RefundAmount
	if !IsNil(o.RefundNo) {
		toSerialize["refundNo"] = o.RefundNo
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceRefundReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
		"reason",
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

	varUnibeeApiMerchantInvoiceRefundReq := _UnibeeApiMerchantInvoiceRefundReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceRefundReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceRefundReq(varUnibeeApiMerchantInvoiceRefundReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceRefundReq struct {
	value *UnibeeApiMerchantInvoiceRefundReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceRefundReq) Get() *UnibeeApiMerchantInvoiceRefundReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceRefundReq) Set(val *UnibeeApiMerchantInvoiceRefundReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceRefundReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceRefundReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceRefundReq(val *UnibeeApiMerchantInvoiceRefundReq) *NullableUnibeeApiMerchantInvoiceRefundReq {
	return &NullableUnibeeApiMerchantInvoiceRefundReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceRefundReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceRefundReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


