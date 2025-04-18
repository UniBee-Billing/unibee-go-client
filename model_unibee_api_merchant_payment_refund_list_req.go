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

// checks if the UnibeeApiMerchantPaymentRefundListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentRefundListReq{}

// UnibeeApiMerchantPaymentRefundListReq struct for UnibeeApiMerchantPaymentRefundListReq
type UnibeeApiMerchantPaymentRefundListReq struct {
	// Currency
	Currency *string `json:"currency,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// GatewayId
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// PaymentId
	PaymentId string `json:"paymentId"`
	// Status,10-create|20-success|30-Failed|40-Reverse
	Status *int32 `json:"status,omitempty"`
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantPaymentRefundListReq UnibeeApiMerchantPaymentRefundListReq

// NewUnibeeApiMerchantPaymentRefundListReq instantiates a new UnibeeApiMerchantPaymentRefundListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentRefundListReq(paymentId string) *UnibeeApiMerchantPaymentRefundListReq {
	this := UnibeeApiMerchantPaymentRefundListReq{}
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiMerchantPaymentRefundListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentRefundListReqWithDefaults() *UnibeeApiMerchantPaymentRefundListReq {
	this := UnibeeApiMerchantPaymentRefundListReq{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPaymentRefundListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantPaymentRefundListReq) SetEmail(v string) {
	o.Email = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantPaymentRefundListReq) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiMerchantPaymentRefundListReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiMerchantPaymentRefundListReq) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantPaymentRefundListReq) SetStatus(v int32) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantPaymentRefundListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantPaymentRefundListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentRefundListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	toSerialize["paymentId"] = o.PaymentId
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentRefundListReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPaymentRefundListReq := _UnibeeApiMerchantPaymentRefundListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentRefundListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentRefundListReq(varUnibeeApiMerchantPaymentRefundListReq)

	return err
}

type NullableUnibeeApiMerchantPaymentRefundListReq struct {
	value *UnibeeApiMerchantPaymentRefundListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentRefundListReq) Get() *UnibeeApiMerchantPaymentRefundListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentRefundListReq) Set(val *UnibeeApiMerchantPaymentRefundListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentRefundListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentRefundListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentRefundListReq(val *UnibeeApiMerchantPaymentRefundListReq) *NullableUnibeeApiMerchantPaymentRefundListReq {
	return &NullableUnibeeApiMerchantPaymentRefundListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentRefundListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentRefundListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


