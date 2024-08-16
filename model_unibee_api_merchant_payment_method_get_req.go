/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentMethodGetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodGetReq{}

// UnibeeApiMerchantPaymentMethodGetReq The method of payment gateway
type UnibeeApiMerchantPaymentMethodGetReq struct {
	// The unique id of gateway
	GatewayId int64 `json:"gatewayId"`
	// The unique id of payment method
	PaymentMethodId string `json:"paymentMethodId"`
	// The customer's unique id
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantPaymentMethodGetReq UnibeeApiMerchantPaymentMethodGetReq

// NewUnibeeApiMerchantPaymentMethodGetReq instantiates a new UnibeeApiMerchantPaymentMethodGetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodGetReq(gatewayId int64, paymentMethodId string, userId int64) *UnibeeApiMerchantPaymentMethodGetReq {
	this := UnibeeApiMerchantPaymentMethodGetReq{}
	this.GatewayId = gatewayId
	this.PaymentMethodId = paymentMethodId
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantPaymentMethodGetReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodGetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodGetReqWithDefaults() *UnibeeApiMerchantPaymentMethodGetReq {
	this := UnibeeApiMerchantPaymentMethodGetReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetPaymentMethodId returns the PaymentMethodId field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetPaymentMethodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodId, true
}

// SetPaymentMethodId sets field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) SetPaymentMethodId(v string) {
	o.PaymentMethodId = v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodGetReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantPaymentMethodGetReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantPaymentMethodGetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodGetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	toSerialize["paymentMethodId"] = o.PaymentMethodId
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentMethodGetReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
		"paymentMethodId",
		"userId",
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

	varUnibeeApiMerchantPaymentMethodGetReq := _UnibeeApiMerchantPaymentMethodGetReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentMethodGetReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentMethodGetReq(varUnibeeApiMerchantPaymentMethodGetReq)

	return err
}

type NullableUnibeeApiMerchantPaymentMethodGetReq struct {
	value *UnibeeApiMerchantPaymentMethodGetReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodGetReq) Get() *UnibeeApiMerchantPaymentMethodGetReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetReq) Set(val *UnibeeApiMerchantPaymentMethodGetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodGetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodGetReq(val *UnibeeApiMerchantPaymentMethodGetReq) *NullableUnibeeApiMerchantPaymentMethodGetReq {
	return &NullableUnibeeApiMerchantPaymentMethodGetReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodGetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


