/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentMethodGetGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentMethodGetGet200ResponseData{}

// MerchantPaymentMethodGetGet200ResponseData struct for MerchantPaymentMethodGetGet200ResponseData
type MerchantPaymentMethodGetGet200ResponseData struct {
	Method *UnibeeApiBeanPaymentMethod `json:"method,omitempty"`
}

// NewMerchantPaymentMethodGetGet200ResponseData instantiates a new MerchantPaymentMethodGetGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentMethodGetGet200ResponseData() *MerchantPaymentMethodGetGet200ResponseData {
	this := MerchantPaymentMethodGetGet200ResponseData{}
	return &this
}

// NewMerchantPaymentMethodGetGet200ResponseDataWithDefaults instantiates a new MerchantPaymentMethodGetGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentMethodGetGet200ResponseDataWithDefaults() *MerchantPaymentMethodGetGet200ResponseData {
	this := MerchantPaymentMethodGetGet200ResponseData{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MerchantPaymentMethodGetGet200ResponseData) GetMethod() UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.Method) {
		var ret UnibeeApiBeanPaymentMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentMethodGetGet200ResponseData) GetMethodOk() (*UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MerchantPaymentMethodGetGet200ResponseData) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given UnibeeApiBeanPaymentMethod and assigns it to the Method field.
func (o *MerchantPaymentMethodGetGet200ResponseData) SetMethod(v UnibeeApiBeanPaymentMethod) {
	o.Method = &v
}

func (o MerchantPaymentMethodGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentMethodGetGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullableMerchantPaymentMethodGetGet200ResponseData struct {
	value *MerchantPaymentMethodGetGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentMethodGetGet200ResponseData) Get() *MerchantPaymentMethodGetGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentMethodGetGet200ResponseData) Set(val *MerchantPaymentMethodGetGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentMethodGetGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentMethodGetGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentMethodGetGet200ResponseData(val *MerchantPaymentMethodGetGet200ResponseData) *NullableMerchantPaymentMethodGetGet200ResponseData {
	return &NullableMerchantPaymentMethodGetGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentMethodGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentMethodGetGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


