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

// checks if the UnibeeApiMerchantSubscriptionChangeGatewayReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionChangeGatewayReq{}

// UnibeeApiMerchantSubscriptionChangeGatewayReq struct for UnibeeApiMerchantSubscriptionChangeGatewayReq
type UnibeeApiMerchantSubscriptionChangeGatewayReq struct {
	// GatewayId
	GatewayId int64 `json:"gatewayId"`
	// PaymentMethodId
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionChangeGatewayReq UnibeeApiMerchantSubscriptionChangeGatewayReq

// NewUnibeeApiMerchantSubscriptionChangeGatewayReq instantiates a new UnibeeApiMerchantSubscriptionChangeGatewayReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionChangeGatewayReq(gatewayId int64, subscriptionId string) *UnibeeApiMerchantSubscriptionChangeGatewayReq {
	this := UnibeeApiMerchantSubscriptionChangeGatewayReq{}
	this.GatewayId = gatewayId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionChangeGatewayReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionChangeGatewayReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionChangeGatewayReqWithDefaults() *UnibeeApiMerchantSubscriptionChangeGatewayReq {
	this := UnibeeApiMerchantSubscriptionChangeGatewayReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionChangeGatewayReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionChangeGatewayReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionChangeGatewayReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionChangeGatewayReq := _UnibeeApiMerchantSubscriptionChangeGatewayReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionChangeGatewayReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionChangeGatewayReq(varUnibeeApiMerchantSubscriptionChangeGatewayReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionChangeGatewayReq struct {
	value *UnibeeApiMerchantSubscriptionChangeGatewayReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) Get() *UnibeeApiMerchantSubscriptionChangeGatewayReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) Set(val *UnibeeApiMerchantSubscriptionChangeGatewayReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionChangeGatewayReq(val *UnibeeApiMerchantSubscriptionChangeGatewayReq) *NullableUnibeeApiMerchantSubscriptionChangeGatewayReq {
	return &NullableUnibeeApiMerchantSubscriptionChangeGatewayReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionChangeGatewayReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


