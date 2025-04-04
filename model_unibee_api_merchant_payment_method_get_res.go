/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentMethodGetRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodGetRes{}

// UnibeeApiMerchantPaymentMethodGetRes struct for UnibeeApiMerchantPaymentMethodGetRes
type UnibeeApiMerchantPaymentMethodGetRes struct {
	Method *UnibeeApiBeanPaymentMethod `json:"method,omitempty"`
}

// NewUnibeeApiMerchantPaymentMethodGetRes instantiates a new UnibeeApiMerchantPaymentMethodGetRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodGetRes() *UnibeeApiMerchantPaymentMethodGetRes {
	this := UnibeeApiMerchantPaymentMethodGetRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentMethodGetResWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodGetRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodGetResWithDefaults() *UnibeeApiMerchantPaymentMethodGetRes {
	this := UnibeeApiMerchantPaymentMethodGetRes{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodGetRes) GetMethod() UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.Method) {
		var ret UnibeeApiBeanPaymentMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodGetRes) GetMethodOk() (*UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodGetRes) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given UnibeeApiBeanPaymentMethod and assigns it to the Method field.
func (o *UnibeeApiMerchantPaymentMethodGetRes) SetMethod(v UnibeeApiBeanPaymentMethod) {
	o.Method = &v
}

func (o UnibeeApiMerchantPaymentMethodGetRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodGetRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentMethodGetRes struct {
	value *UnibeeApiMerchantPaymentMethodGetRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodGetRes) Get() *UnibeeApiMerchantPaymentMethodGetRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetRes) Set(val *UnibeeApiMerchantPaymentMethodGetRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodGetRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodGetRes(val *UnibeeApiMerchantPaymentMethodGetRes) *NullableUnibeeApiMerchantPaymentMethodGetRes {
	return &NullableUnibeeApiMerchantPaymentMethodGetRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodGetRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodGetRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


