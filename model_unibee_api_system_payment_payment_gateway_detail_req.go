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

// checks if the UnibeeApiSystemPaymentPaymentGatewayDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemPaymentPaymentGatewayDetailReq{}

// UnibeeApiSystemPaymentPaymentGatewayDetailReq struct for UnibeeApiSystemPaymentPaymentGatewayDetailReq
type UnibeeApiSystemPaymentPaymentGatewayDetailReq struct {
	// PaymentId
	PaymentId string `json:"paymentId"`
}

type _UnibeeApiSystemPaymentPaymentGatewayDetailReq UnibeeApiSystemPaymentPaymentGatewayDetailReq

// NewUnibeeApiSystemPaymentPaymentGatewayDetailReq instantiates a new UnibeeApiSystemPaymentPaymentGatewayDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemPaymentPaymentGatewayDetailReq(paymentId string) *UnibeeApiSystemPaymentPaymentGatewayDetailReq {
	this := UnibeeApiSystemPaymentPaymentGatewayDetailReq{}
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiSystemPaymentPaymentGatewayDetailReqWithDefaults instantiates a new UnibeeApiSystemPaymentPaymentGatewayDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemPaymentPaymentGatewayDetailReqWithDefaults() *UnibeeApiSystemPaymentPaymentGatewayDetailReq {
	this := UnibeeApiSystemPaymentPaymentGatewayDetailReq{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiSystemPaymentPaymentGatewayDetailReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemPaymentPaymentGatewayDetailReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiSystemPaymentPaymentGatewayDetailReq) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o UnibeeApiSystemPaymentPaymentGatewayDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemPaymentPaymentGatewayDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	return toSerialize, nil
}

func (o *UnibeeApiSystemPaymentPaymentGatewayDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiSystemPaymentPaymentGatewayDetailReq := _UnibeeApiSystemPaymentPaymentGatewayDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemPaymentPaymentGatewayDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemPaymentPaymentGatewayDetailReq(varUnibeeApiSystemPaymentPaymentGatewayDetailReq)

	return err
}

type NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq struct {
	value *UnibeeApiSystemPaymentPaymentGatewayDetailReq
	isSet bool
}

func (v NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) Get() *UnibeeApiSystemPaymentPaymentGatewayDetailReq {
	return v.value
}

func (v *NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) Set(val *UnibeeApiSystemPaymentPaymentGatewayDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemPaymentPaymentGatewayDetailReq(val *UnibeeApiSystemPaymentPaymentGatewayDetailReq) *NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq {
	return &NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemPaymentPaymentGatewayDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


