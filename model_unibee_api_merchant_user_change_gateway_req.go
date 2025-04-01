/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantUserChangeGatewayReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserChangeGatewayReq{}

// UnibeeApiMerchantUserChangeGatewayReq struct for UnibeeApiMerchantUserChangeGatewayReq
type UnibeeApiMerchantUserChangeGatewayReq struct {
	// GatewayId
	GatewayId int64 `json:"gatewayId"`
	// GatewayPaymentType
	GatewayPaymentType *string `json:"gatewayPaymentType,omitempty"`
	// GatewayUserId, verify and save GatewayUserId via gateway
	GatewayUserId *string `json:"gatewayUserId,omitempty"`
	// PaymentMethodId of gateway, available for card type gateway, payment automatic will enable if set
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	// User Id
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantUserChangeGatewayReq UnibeeApiMerchantUserChangeGatewayReq

// NewUnibeeApiMerchantUserChangeGatewayReq instantiates a new UnibeeApiMerchantUserChangeGatewayReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserChangeGatewayReq(gatewayId int64, userId int64) *UnibeeApiMerchantUserChangeGatewayReq {
	this := UnibeeApiMerchantUserChangeGatewayReq{}
	this.GatewayId = gatewayId
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantUserChangeGatewayReqWithDefaults instantiates a new UnibeeApiMerchantUserChangeGatewayReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserChangeGatewayReqWithDefaults() *UnibeeApiMerchantUserChangeGatewayReq {
	this := UnibeeApiMerchantUserChangeGatewayReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetGatewayPaymentType returns the GatewayPaymentType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayPaymentType() string {
	if o == nil || IsNil(o.GatewayPaymentType) {
		var ret string
		return ret
	}
	return *o.GatewayPaymentType
}

// GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayPaymentType) {
		return nil, false
	}
	return o.GatewayPaymentType, true
}

// HasGatewayPaymentType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) HasGatewayPaymentType() bool {
	if o != nil && !IsNil(o.GatewayPaymentType) {
		return true
	}

	return false
}

// SetGatewayPaymentType gets a reference to the given string and assigns it to the GatewayPaymentType field.
func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayPaymentType(v string) {
	o.GatewayPaymentType = &v
}

// GetGatewayUserId returns the GatewayUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayUserId() string {
	if o == nil || IsNil(o.GatewayUserId) {
		var ret string
		return ret
	}
	return *o.GatewayUserId
}

// GetGatewayUserIdOk returns a tuple with the GatewayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetGatewayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayUserId) {
		return nil, false
	}
	return o.GatewayUserId, true
}

// HasGatewayUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) HasGatewayUserId() bool {
	if o != nil && !IsNil(o.GatewayUserId) {
		return true
	}

	return false
}

// SetGatewayUserId gets a reference to the given string and assigns it to the GatewayUserId field.
func (o *UnibeeApiMerchantUserChangeGatewayReq) SetGatewayUserId(v string) {
	o.GatewayUserId = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *UnibeeApiMerchantUserChangeGatewayReq) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeGatewayReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantUserChangeGatewayReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantUserChangeGatewayReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserChangeGatewayReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.GatewayPaymentType) {
		toSerialize["gatewayPaymentType"] = o.GatewayPaymentType
	}
	if !IsNil(o.GatewayUserId) {
		toSerialize["gatewayUserId"] = o.GatewayUserId
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantUserChangeGatewayReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
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

	varUnibeeApiMerchantUserChangeGatewayReq := _UnibeeApiMerchantUserChangeGatewayReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantUserChangeGatewayReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantUserChangeGatewayReq(varUnibeeApiMerchantUserChangeGatewayReq)

	return err
}

type NullableUnibeeApiMerchantUserChangeGatewayReq struct {
	value *UnibeeApiMerchantUserChangeGatewayReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserChangeGatewayReq) Get() *UnibeeApiMerchantUserChangeGatewayReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserChangeGatewayReq) Set(val *UnibeeApiMerchantUserChangeGatewayReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserChangeGatewayReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserChangeGatewayReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserChangeGatewayReq(val *UnibeeApiMerchantUserChangeGatewayReq) *NullableUnibeeApiMerchantUserChangeGatewayReq {
	return &NullableUnibeeApiMerchantUserChangeGatewayReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserChangeGatewayReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserChangeGatewayReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


