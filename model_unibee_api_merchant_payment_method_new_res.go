/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentMethodNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodNewRes{}

// UnibeeApiMerchantPaymentMethodNewRes struct for UnibeeApiMerchantPaymentMethodNewRes
type UnibeeApiMerchantPaymentMethodNewRes struct {
	Method *UnibeeApiBeanPaymentMethod `json:"method,omitempty"`
	// The gateway url to create a new method
	Url *string `json:"url,omitempty"`
}

// NewUnibeeApiMerchantPaymentMethodNewRes instantiates a new UnibeeApiMerchantPaymentMethodNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodNewRes() *UnibeeApiMerchantPaymentMethodNewRes {
	this := UnibeeApiMerchantPaymentMethodNewRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentMethodNewResWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodNewResWithDefaults() *UnibeeApiMerchantPaymentMethodNewRes {
	this := UnibeeApiMerchantPaymentMethodNewRes{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewRes) GetMethod() UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.Method) {
		var ret UnibeeApiBeanPaymentMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewRes) GetMethodOk() (*UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewRes) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given UnibeeApiBeanPaymentMethod and assigns it to the Method field.
func (o *UnibeeApiMerchantPaymentMethodNewRes) SetMethod(v UnibeeApiBeanPaymentMethod) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewRes) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewRes) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewRes) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UnibeeApiMerchantPaymentMethodNewRes) SetUrl(v string) {
	o.Url = &v
}

func (o UnibeeApiMerchantPaymentMethodNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentMethodNewRes struct {
	value *UnibeeApiMerchantPaymentMethodNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodNewRes) Get() *UnibeeApiMerchantPaymentMethodNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewRes) Set(val *UnibeeApiMerchantPaymentMethodNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodNewRes(val *UnibeeApiMerchantPaymentMethodNewRes) *NullableUnibeeApiMerchantPaymentMethodNewRes {
	return &NullableUnibeeApiMerchantPaymentMethodNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


