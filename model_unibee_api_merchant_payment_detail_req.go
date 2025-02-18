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

// checks if the UnibeeApiMerchantPaymentDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentDetailReq{}

// UnibeeApiMerchantPaymentDetailReq struct for UnibeeApiMerchantPaymentDetailReq
type UnibeeApiMerchantPaymentDetailReq struct {
	// The unique id of payment
	PaymentId string `json:"paymentId"`
}

type _UnibeeApiMerchantPaymentDetailReq UnibeeApiMerchantPaymentDetailReq

// NewUnibeeApiMerchantPaymentDetailReq instantiates a new UnibeeApiMerchantPaymentDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentDetailReq(paymentId string) *UnibeeApiMerchantPaymentDetailReq {
	this := UnibeeApiMerchantPaymentDetailReq{}
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiMerchantPaymentDetailReqWithDefaults instantiates a new UnibeeApiMerchantPaymentDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentDetailReqWithDefaults() *UnibeeApiMerchantPaymentDetailReq {
	this := UnibeeApiMerchantPaymentDetailReq{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiMerchantPaymentDetailReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentDetailReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiMerchantPaymentDetailReq) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o UnibeeApiMerchantPaymentDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPaymentDetailReq := _UnibeeApiMerchantPaymentDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentDetailReq(varUnibeeApiMerchantPaymentDetailReq)

	return err
}

type NullableUnibeeApiMerchantPaymentDetailReq struct {
	value *UnibeeApiMerchantPaymentDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentDetailReq) Get() *UnibeeApiMerchantPaymentDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentDetailReq) Set(val *UnibeeApiMerchantPaymentDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentDetailReq(val *UnibeeApiMerchantPaymentDetailReq) *NullableUnibeeApiMerchantPaymentDetailReq {
	return &NullableUnibeeApiMerchantPaymentDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


