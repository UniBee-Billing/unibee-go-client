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

// checks if the MerchantPaymentMethodNewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentMethodNewPost200ResponseData{}

// MerchantPaymentMethodNewPost200ResponseData struct for MerchantPaymentMethodNewPost200ResponseData
type MerchantPaymentMethodNewPost200ResponseData struct {
	Method *UnibeeApiBeanPaymentMethod `json:"method,omitempty"`
	// The gateway url to create a new method
	Url *string `json:"url,omitempty"`
}

// NewMerchantPaymentMethodNewPost200ResponseData instantiates a new MerchantPaymentMethodNewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentMethodNewPost200ResponseData() *MerchantPaymentMethodNewPost200ResponseData {
	this := MerchantPaymentMethodNewPost200ResponseData{}
	return &this
}

// NewMerchantPaymentMethodNewPost200ResponseDataWithDefaults instantiates a new MerchantPaymentMethodNewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentMethodNewPost200ResponseDataWithDefaults() *MerchantPaymentMethodNewPost200ResponseData {
	this := MerchantPaymentMethodNewPost200ResponseData{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MerchantPaymentMethodNewPost200ResponseData) GetMethod() UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.Method) {
		var ret UnibeeApiBeanPaymentMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentMethodNewPost200ResponseData) GetMethodOk() (*UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MerchantPaymentMethodNewPost200ResponseData) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given UnibeeApiBeanPaymentMethod and assigns it to the Method field.
func (o *MerchantPaymentMethodNewPost200ResponseData) SetMethod(v UnibeeApiBeanPaymentMethod) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MerchantPaymentMethodNewPost200ResponseData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentMethodNewPost200ResponseData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MerchantPaymentMethodNewPost200ResponseData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MerchantPaymentMethodNewPost200ResponseData) SetUrl(v string) {
	o.Url = &v
}

func (o MerchantPaymentMethodNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentMethodNewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableMerchantPaymentMethodNewPost200ResponseData struct {
	value *MerchantPaymentMethodNewPost200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentMethodNewPost200ResponseData) Get() *MerchantPaymentMethodNewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentMethodNewPost200ResponseData) Set(val *MerchantPaymentMethodNewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentMethodNewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentMethodNewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentMethodNewPost200ResponseData(val *MerchantPaymentMethodNewPost200ResponseData) *NullableMerchantPaymentMethodNewPost200ResponseData {
	return &NullableMerchantPaymentMethodNewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentMethodNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentMethodNewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


