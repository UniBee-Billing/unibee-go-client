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

// checks if the UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq{}

// UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq Mark refund invoice success, only support Changelly and Wire Transfer
type UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// The reason of mark action
	Reason *string `json:"reason,omitempty"`
}

type _UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq

// NewUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq instantiates a new UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq(invoiceId string) *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq {
	this := UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReqWithDefaults() *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq {
	this := UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) SetReason(v string) {
	o.Reason = &v
}

func (o UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
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

	varUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq := _UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq(varUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq struct {
	value *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) Get() *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) Set(val *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq(val *UnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) *NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq {
	return &NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundInvoiceSuccessReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


