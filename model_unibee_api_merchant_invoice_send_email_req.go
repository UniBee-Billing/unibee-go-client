/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceSendEmailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceSendEmailReq{}

// UnibeeApiMerchantInvoiceSendEmailReq struct for UnibeeApiMerchantInvoiceSendEmailReq
type UnibeeApiMerchantInvoiceSendEmailReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
}

type _UnibeeApiMerchantInvoiceSendEmailReq UnibeeApiMerchantInvoiceSendEmailReq

// NewUnibeeApiMerchantInvoiceSendEmailReq instantiates a new UnibeeApiMerchantInvoiceSendEmailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceSendEmailReq(invoiceId string) *UnibeeApiMerchantInvoiceSendEmailReq {
	this := UnibeeApiMerchantInvoiceSendEmailReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceSendEmailReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceSendEmailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceSendEmailReqWithDefaults() *UnibeeApiMerchantInvoiceSendEmailReq {
	this := UnibeeApiMerchantInvoiceSendEmailReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceSendEmailReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceSendEmailReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceSendEmailReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

func (o UnibeeApiMerchantInvoiceSendEmailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceSendEmailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceSendEmailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantInvoiceSendEmailReq := _UnibeeApiMerchantInvoiceSendEmailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceSendEmailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceSendEmailReq(varUnibeeApiMerchantInvoiceSendEmailReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceSendEmailReq struct {
	value *UnibeeApiMerchantInvoiceSendEmailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceSendEmailReq) Get() *UnibeeApiMerchantInvoiceSendEmailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceSendEmailReq) Set(val *UnibeeApiMerchantInvoiceSendEmailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceSendEmailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceSendEmailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceSendEmailReq(val *UnibeeApiMerchantInvoiceSendEmailReq) *NullableUnibeeApiMerchantInvoiceSendEmailReq {
	return &NullableUnibeeApiMerchantInvoiceSendEmailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceSendEmailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceSendEmailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


