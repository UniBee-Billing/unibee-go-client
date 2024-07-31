/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoicePdfGenerateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoicePdfGenerateReq{}

// UnibeeApiMerchantInvoicePdfGenerateReq struct for UnibeeApiMerchantInvoicePdfGenerateReq
type UnibeeApiMerchantInvoicePdfGenerateReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// Whether sen invoice email to user or not，default false
	SendUserEmail *bool `json:"sendUserEmail,omitempty"`
}

type _UnibeeApiMerchantInvoicePdfGenerateReq UnibeeApiMerchantInvoicePdfGenerateReq

// NewUnibeeApiMerchantInvoicePdfGenerateReq instantiates a new UnibeeApiMerchantInvoicePdfGenerateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoicePdfGenerateReq(invoiceId string) *UnibeeApiMerchantInvoicePdfGenerateReq {
	this := UnibeeApiMerchantInvoicePdfGenerateReq{}
	this.InvoiceId = invoiceId
	var sendUserEmail bool = false
	this.SendUserEmail = &sendUserEmail
	return &this
}

// NewUnibeeApiMerchantInvoicePdfGenerateReqWithDefaults instantiates a new UnibeeApiMerchantInvoicePdfGenerateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoicePdfGenerateReqWithDefaults() *UnibeeApiMerchantInvoicePdfGenerateReq {
	this := UnibeeApiMerchantInvoicePdfGenerateReq{}
	var sendUserEmail bool = false
	this.SendUserEmail = &sendUserEmail
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetSendUserEmail returns the SendUserEmail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetSendUserEmail() bool {
	if o == nil || IsNil(o.SendUserEmail) {
		var ret bool
		return ret
	}
	return *o.SendUserEmail
}

// GetSendUserEmailOk returns a tuple with the SendUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetSendUserEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendUserEmail) {
		return nil, false
	}
	return o.SendUserEmail, true
}

// HasSendUserEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) HasSendUserEmail() bool {
	if o != nil && !IsNil(o.SendUserEmail) {
		return true
	}

	return false
}

// SetSendUserEmail gets a reference to the given bool and assigns it to the SendUserEmail field.
func (o *UnibeeApiMerchantInvoicePdfGenerateReq) SetSendUserEmail(v bool) {
	o.SendUserEmail = &v
}

func (o UnibeeApiMerchantInvoicePdfGenerateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoicePdfGenerateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	if !IsNil(o.SendUserEmail) {
		toSerialize["sendUserEmail"] = o.SendUserEmail
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoicePdfGenerateReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantInvoicePdfGenerateReq := _UnibeeApiMerchantInvoicePdfGenerateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoicePdfGenerateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoicePdfGenerateReq(varUnibeeApiMerchantInvoicePdfGenerateReq)

	return err
}

type NullableUnibeeApiMerchantInvoicePdfGenerateReq struct {
	value *UnibeeApiMerchantInvoicePdfGenerateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoicePdfGenerateReq) Get() *UnibeeApiMerchantInvoicePdfGenerateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoicePdfGenerateReq) Set(val *UnibeeApiMerchantInvoicePdfGenerateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoicePdfGenerateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoicePdfGenerateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoicePdfGenerateReq(val *UnibeeApiMerchantInvoicePdfGenerateReq) *NullableUnibeeApiMerchantInvoicePdfGenerateReq {
	return &NullableUnibeeApiMerchantInvoicePdfGenerateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoicePdfGenerateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoicePdfGenerateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


