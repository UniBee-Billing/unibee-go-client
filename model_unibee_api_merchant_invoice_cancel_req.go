/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceCancelReq{}

// UnibeeApiMerchantInvoiceCancelReq struct for UnibeeApiMerchantInvoiceCancelReq
type UnibeeApiMerchantInvoiceCancelReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
}

type _UnibeeApiMerchantInvoiceCancelReq UnibeeApiMerchantInvoiceCancelReq

// NewUnibeeApiMerchantInvoiceCancelReq instantiates a new UnibeeApiMerchantInvoiceCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceCancelReq(invoiceId string) *UnibeeApiMerchantInvoiceCancelReq {
	this := UnibeeApiMerchantInvoiceCancelReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceCancelReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceCancelReqWithDefaults() *UnibeeApiMerchantInvoiceCancelReq {
	this := UnibeeApiMerchantInvoiceCancelReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceCancelReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceCancelReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceCancelReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

func (o UnibeeApiMerchantInvoiceCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceCancelReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantInvoiceCancelReq := _UnibeeApiMerchantInvoiceCancelReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceCancelReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceCancelReq(varUnibeeApiMerchantInvoiceCancelReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceCancelReq struct {
	value *UnibeeApiMerchantInvoiceCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) Get() *UnibeeApiMerchantInvoiceCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) Set(val *UnibeeApiMerchantInvoiceCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceCancelReq(val *UnibeeApiMerchantInvoiceCancelReq) *NullableUnibeeApiMerchantInvoiceCancelReq {
	return &NullableUnibeeApiMerchantInvoiceCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


