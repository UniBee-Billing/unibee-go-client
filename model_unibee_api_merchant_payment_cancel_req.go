/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentCancelReq{}

// UnibeeApiMerchantPaymentCancelReq struct for UnibeeApiMerchantPaymentCancelReq
type UnibeeApiMerchantPaymentCancelReq struct {
	// The unique id of payment
	PaymentId string `json:"paymentId"`
}

type _UnibeeApiMerchantPaymentCancelReq UnibeeApiMerchantPaymentCancelReq

// NewUnibeeApiMerchantPaymentCancelReq instantiates a new UnibeeApiMerchantPaymentCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentCancelReq(paymentId string) *UnibeeApiMerchantPaymentCancelReq {
	this := UnibeeApiMerchantPaymentCancelReq{}
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiMerchantPaymentCancelReqWithDefaults instantiates a new UnibeeApiMerchantPaymentCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentCancelReqWithDefaults() *UnibeeApiMerchantPaymentCancelReq {
	this := UnibeeApiMerchantPaymentCancelReq{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiMerchantPaymentCancelReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentCancelReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiMerchantPaymentCancelReq) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o UnibeeApiMerchantPaymentCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentCancelReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"paymentId",
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

	varUnibeeApiMerchantPaymentCancelReq := _UnibeeApiMerchantPaymentCancelReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentCancelReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentCancelReq(varUnibeeApiMerchantPaymentCancelReq)

	return err
}

type NullableUnibeeApiMerchantPaymentCancelReq struct {
	value *UnibeeApiMerchantPaymentCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentCancelReq) Get() *UnibeeApiMerchantPaymentCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentCancelReq) Set(val *UnibeeApiMerchantPaymentCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentCancelReq(val *UnibeeApiMerchantPaymentCancelReq) *NullableUnibeeApiMerchantPaymentCancelReq {
	return &NullableUnibeeApiMerchantPaymentCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


