/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentCaptureReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentCaptureReq{}

// UnibeeApiMerchantPaymentCaptureReq struct for UnibeeApiMerchantPaymentCaptureReq
type UnibeeApiMerchantPaymentCaptureReq struct {
	// CaptureAmount, Cent
	CaptureAmount int64 `json:"captureAmount"`
	// Currency
	Currency string `json:"currency"`
	// ExternalCaptureId
	ExternalCaptureId string `json:"externalCaptureId"`
	// PaymentId
	PaymentId string `json:"paymentId"`
}

type _UnibeeApiMerchantPaymentCaptureReq UnibeeApiMerchantPaymentCaptureReq

// NewUnibeeApiMerchantPaymentCaptureReq instantiates a new UnibeeApiMerchantPaymentCaptureReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentCaptureReq(captureAmount int64, currency string, externalCaptureId string, paymentId string) *UnibeeApiMerchantPaymentCaptureReq {
	this := UnibeeApiMerchantPaymentCaptureReq{}
	this.CaptureAmount = captureAmount
	this.Currency = currency
	this.ExternalCaptureId = externalCaptureId
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiMerchantPaymentCaptureReqWithDefaults instantiates a new UnibeeApiMerchantPaymentCaptureReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentCaptureReqWithDefaults() *UnibeeApiMerchantPaymentCaptureReq {
	this := UnibeeApiMerchantPaymentCaptureReq{}
	return &this
}

// GetCaptureAmount returns the CaptureAmount field value
func (o *UnibeeApiMerchantPaymentCaptureReq) GetCaptureAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CaptureAmount
}

// GetCaptureAmountOk returns a tuple with the CaptureAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentCaptureReq) GetCaptureAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptureAmount, true
}

// SetCaptureAmount sets field value
func (o *UnibeeApiMerchantPaymentCaptureReq) SetCaptureAmount(v int64) {
	o.CaptureAmount = v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantPaymentCaptureReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentCaptureReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantPaymentCaptureReq) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalCaptureId returns the ExternalCaptureId field value
func (o *UnibeeApiMerchantPaymentCaptureReq) GetExternalCaptureId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCaptureId
}

// GetExternalCaptureIdOk returns a tuple with the ExternalCaptureId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentCaptureReq) GetExternalCaptureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCaptureId, true
}

// SetExternalCaptureId sets field value
func (o *UnibeeApiMerchantPaymentCaptureReq) SetExternalCaptureId(v string) {
	o.ExternalCaptureId = v
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiMerchantPaymentCaptureReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentCaptureReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiMerchantPaymentCaptureReq) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o UnibeeApiMerchantPaymentCaptureReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentCaptureReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["captureAmount"] = o.CaptureAmount
	toSerialize["currency"] = o.Currency
	toSerialize["externalCaptureId"] = o.ExternalCaptureId
	toSerialize["paymentId"] = o.PaymentId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentCaptureReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"captureAmount",
		"currency",
		"externalCaptureId",
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

	varUnibeeApiMerchantPaymentCaptureReq := _UnibeeApiMerchantPaymentCaptureReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentCaptureReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentCaptureReq(varUnibeeApiMerchantPaymentCaptureReq)

	return err
}

type NullableUnibeeApiMerchantPaymentCaptureReq struct {
	value *UnibeeApiMerchantPaymentCaptureReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentCaptureReq) Get() *UnibeeApiMerchantPaymentCaptureReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentCaptureReq) Set(val *UnibeeApiMerchantPaymentCaptureReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentCaptureReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentCaptureReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentCaptureReq(val *UnibeeApiMerchantPaymentCaptureReq) *NullableUnibeeApiMerchantPaymentCaptureReq {
	return &NullableUnibeeApiMerchantPaymentCaptureReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentCaptureReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentCaptureReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


