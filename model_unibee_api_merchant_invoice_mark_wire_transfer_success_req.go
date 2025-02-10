/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq{}

// UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq Mark wire transfer pending invoice as success
type UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// The reason of mark action
	Reason *string `json:"reason,omitempty"`
	// The transfer number of invoice
	TransferNumber string `json:"transferNumber"`
}

type _UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq

// NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq instantiates a new UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(invoiceId string, transferNumber string) *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq {
	this := UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq{}
	this.InvoiceId = invoiceId
	this.TransferNumber = transferNumber
	return &this
}

// NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReqWithDefaults() *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq {
	this := UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetReason(v string) {
	o.Reason = &v
}

// GetTransferNumber returns the TransferNumber field value
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetTransferNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferNumber
}

// GetTransferNumberOk returns a tuple with the TransferNumber field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetTransferNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferNumber, true
}

// SetTransferNumber sets field value
func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetTransferNumber(v string) {
	o.TransferNumber = v
}

func (o UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["transferNumber"] = o.TransferNumber
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
		"transferNumber",
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

	varUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq := _UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(varUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq struct {
	value *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) Get() *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) Set(val *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(val *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) *NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq {
	return &NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


