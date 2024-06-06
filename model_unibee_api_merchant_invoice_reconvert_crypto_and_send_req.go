/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq{}

// UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq struct for UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq
type UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
}

type _UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq

// NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq instantiates a new UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(invoiceId string) *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq {
	this := UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceReconvertCryptoAndSendReqWithDefaults() *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq {
	this := UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

func (o UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq := _UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(varUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq struct {
	value *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) Get() *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) Set(val *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq(val *UnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) *NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq {
	return &NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceReconvertCryptoAndSendReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


