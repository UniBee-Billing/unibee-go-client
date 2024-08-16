/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentRefundCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentRefundCancelReq{}

// UnibeeApiMerchantPaymentRefundCancelReq struct for UnibeeApiMerchantPaymentRefundCancelReq
type UnibeeApiMerchantPaymentRefundCancelReq struct {
	// The unique id of refund
	RefundId string `json:"refundId"`
}

type _UnibeeApiMerchantPaymentRefundCancelReq UnibeeApiMerchantPaymentRefundCancelReq

// NewUnibeeApiMerchantPaymentRefundCancelReq instantiates a new UnibeeApiMerchantPaymentRefundCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentRefundCancelReq(refundId string) *UnibeeApiMerchantPaymentRefundCancelReq {
	this := UnibeeApiMerchantPaymentRefundCancelReq{}
	this.RefundId = refundId
	return &this
}

// NewUnibeeApiMerchantPaymentRefundCancelReqWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentRefundCancelReqWithDefaults() *UnibeeApiMerchantPaymentRefundCancelReq {
	this := UnibeeApiMerchantPaymentRefundCancelReq{}
	return &this
}

// GetRefundId returns the RefundId field value
func (o *UnibeeApiMerchantPaymentRefundCancelReq) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundCancelReq) GetRefundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *UnibeeApiMerchantPaymentRefundCancelReq) SetRefundId(v string) {
	o.RefundId = v
}

func (o UnibeeApiMerchantPaymentRefundCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentRefundCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refundId"] = o.RefundId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentRefundCancelReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"refundId",
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

	varUnibeeApiMerchantPaymentRefundCancelReq := _UnibeeApiMerchantPaymentRefundCancelReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentRefundCancelReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentRefundCancelReq(varUnibeeApiMerchantPaymentRefundCancelReq)

	return err
}

type NullableUnibeeApiMerchantPaymentRefundCancelReq struct {
	value *UnibeeApiMerchantPaymentRefundCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentRefundCancelReq) Get() *UnibeeApiMerchantPaymentRefundCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentRefundCancelReq) Set(val *UnibeeApiMerchantPaymentRefundCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentRefundCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentRefundCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentRefundCancelReq(val *UnibeeApiMerchantPaymentRefundCancelReq) *NullableUnibeeApiMerchantPaymentRefundCancelReq {
	return &NullableUnibeeApiMerchantPaymentRefundCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentRefundCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentRefundCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


