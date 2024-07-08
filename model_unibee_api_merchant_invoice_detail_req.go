/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceDetailReq{}

// UnibeeApiMerchantInvoiceDetailReq Get detail of invoice
type UnibeeApiMerchantInvoiceDetailReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
}

type _UnibeeApiMerchantInvoiceDetailReq UnibeeApiMerchantInvoiceDetailReq

// NewUnibeeApiMerchantInvoiceDetailReq instantiates a new UnibeeApiMerchantInvoiceDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceDetailReq(invoiceId string) *UnibeeApiMerchantInvoiceDetailReq {
	this := UnibeeApiMerchantInvoiceDetailReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceDetailReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceDetailReqWithDefaults() *UnibeeApiMerchantInvoiceDetailReq {
	this := UnibeeApiMerchantInvoiceDetailReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceDetailReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceDetailReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceDetailReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

func (o UnibeeApiMerchantInvoiceDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantInvoiceDetailReq := _UnibeeApiMerchantInvoiceDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceDetailReq(varUnibeeApiMerchantInvoiceDetailReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceDetailReq struct {
	value *UnibeeApiMerchantInvoiceDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceDetailReq) Get() *UnibeeApiMerchantInvoiceDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceDetailReq) Set(val *UnibeeApiMerchantInvoiceDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceDetailReq(val *UnibeeApiMerchantInvoiceDetailReq) *NullableUnibeeApiMerchantInvoiceDetailReq {
	return &NullableUnibeeApiMerchantInvoiceDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


