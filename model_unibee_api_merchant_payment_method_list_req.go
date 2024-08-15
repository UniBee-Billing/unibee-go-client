/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentMethodListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodListReq{}

// UnibeeApiMerchantPaymentMethodListReq The method list of payment gateway
type UnibeeApiMerchantPaymentMethodListReq struct {
	// The unique id of gateway
	GatewayId int64 `json:"gatewayId"`
	// The unique id of payment
	PaymentId *string `json:"paymentId,omitempty"`
	// The id of user
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantPaymentMethodListReq UnibeeApiMerchantPaymentMethodListReq

// NewUnibeeApiMerchantPaymentMethodListReq instantiates a new UnibeeApiMerchantPaymentMethodListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodListReq(gatewayId int64) *UnibeeApiMerchantPaymentMethodListReq {
	this := UnibeeApiMerchantPaymentMethodListReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantPaymentMethodListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodListReqWithDefaults() *UnibeeApiMerchantPaymentMethodListReq {
	this := UnibeeApiMerchantPaymentMethodListReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantPaymentMethodListReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodListReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantPaymentMethodListReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodListReq) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodListReq) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodListReq) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiMerchantPaymentMethodListReq) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantPaymentMethodListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantPaymentMethodListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentMethodListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
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

	varUnibeeApiMerchantPaymentMethodListReq := _UnibeeApiMerchantPaymentMethodListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentMethodListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentMethodListReq(varUnibeeApiMerchantPaymentMethodListReq)

	return err
}

type NullableUnibeeApiMerchantPaymentMethodListReq struct {
	value *UnibeeApiMerchantPaymentMethodListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodListReq) Get() *UnibeeApiMerchantPaymentMethodListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodListReq) Set(val *UnibeeApiMerchantPaymentMethodListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodListReq(val *UnibeeApiMerchantPaymentMethodListReq) *NullableUnibeeApiMerchantPaymentMethodListReq {
	return &NullableUnibeeApiMerchantPaymentMethodListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


