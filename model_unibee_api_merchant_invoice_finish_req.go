/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceFinishReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceFinishReq{}

// UnibeeApiMerchantInvoiceFinishReq Finish invoice, generate invoice link
type UnibeeApiMerchantInvoiceFinishReq struct {
	// Due Day Of Invoice Payment
	DaysUtilDue int32 `json:"daysUtilDue"`
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
}

type _UnibeeApiMerchantInvoiceFinishReq UnibeeApiMerchantInvoiceFinishReq

// NewUnibeeApiMerchantInvoiceFinishReq instantiates a new UnibeeApiMerchantInvoiceFinishReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceFinishReq(daysUtilDue int32, invoiceId string) *UnibeeApiMerchantInvoiceFinishReq {
	this := UnibeeApiMerchantInvoiceFinishReq{}
	this.DaysUtilDue = daysUtilDue
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceFinishReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceFinishReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceFinishReqWithDefaults() *UnibeeApiMerchantInvoiceFinishReq {
	this := UnibeeApiMerchantInvoiceFinishReq{}
	return &this
}

// GetDaysUtilDue returns the DaysUtilDue field value
func (o *UnibeeApiMerchantInvoiceFinishReq) GetDaysUtilDue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysUtilDue
}

// GetDaysUtilDueOk returns a tuple with the DaysUtilDue field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceFinishReq) GetDaysUtilDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysUtilDue, true
}

// SetDaysUtilDue sets field value
func (o *UnibeeApiMerchantInvoiceFinishReq) SetDaysUtilDue(v int32) {
	o.DaysUtilDue = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceFinishReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceFinishReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceFinishReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

func (o UnibeeApiMerchantInvoiceFinishReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceFinishReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["daysUtilDue"] = o.DaysUtilDue
	toSerialize["invoiceId"] = o.InvoiceId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceFinishReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"daysUtilDue",
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

	varUnibeeApiMerchantInvoiceFinishReq := _UnibeeApiMerchantInvoiceFinishReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceFinishReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceFinishReq(varUnibeeApiMerchantInvoiceFinishReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceFinishReq struct {
	value *UnibeeApiMerchantInvoiceFinishReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceFinishReq) Get() *UnibeeApiMerchantInvoiceFinishReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceFinishReq) Set(val *UnibeeApiMerchantInvoiceFinishReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceFinishReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceFinishReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceFinishReq(val *UnibeeApiMerchantInvoiceFinishReq) *NullableUnibeeApiMerchantInvoiceFinishReq {
	return &NullableUnibeeApiMerchantInvoiceFinishReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceFinishReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceFinishReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


