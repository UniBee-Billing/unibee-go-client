/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentMethodListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentMethodListGet200ResponseData{}

// MerchantPaymentMethodListGet200ResponseData struct for MerchantPaymentMethodListGet200ResponseData
type MerchantPaymentMethodListGet200ResponseData struct {
	// Payment Method Object List
	MethodList []UnibeeApiBeanPaymentMethod `json:"methodList,omitempty"`
}

// NewMerchantPaymentMethodListGet200ResponseData instantiates a new MerchantPaymentMethodListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentMethodListGet200ResponseData() *MerchantPaymentMethodListGet200ResponseData {
	this := MerchantPaymentMethodListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentMethodListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentMethodListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentMethodListGet200ResponseDataWithDefaults() *MerchantPaymentMethodListGet200ResponseData {
	this := MerchantPaymentMethodListGet200ResponseData{}
	return &this
}

// GetMethodList returns the MethodList field value if set, zero value otherwise.
func (o *MerchantPaymentMethodListGet200ResponseData) GetMethodList() []UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.MethodList) {
		var ret []UnibeeApiBeanPaymentMethod
		return ret
	}
	return o.MethodList
}

// GetMethodListOk returns a tuple with the MethodList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentMethodListGet200ResponseData) GetMethodListOk() ([]UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.MethodList) {
		return nil, false
	}
	return o.MethodList, true
}

// HasMethodList returns a boolean if a field has been set.
func (o *MerchantPaymentMethodListGet200ResponseData) HasMethodList() bool {
	if o != nil && !IsNil(o.MethodList) {
		return true
	}

	return false
}

// SetMethodList gets a reference to the given []UnibeeApiBeanPaymentMethod and assigns it to the MethodList field.
func (o *MerchantPaymentMethodListGet200ResponseData) SetMethodList(v []UnibeeApiBeanPaymentMethod) {
	o.MethodList = v
}

func (o MerchantPaymentMethodListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentMethodListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MethodList) {
		toSerialize["methodList"] = o.MethodList
	}
	return toSerialize, nil
}

type NullableMerchantPaymentMethodListGet200ResponseData struct {
	value *MerchantPaymentMethodListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentMethodListGet200ResponseData) Get() *MerchantPaymentMethodListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentMethodListGet200ResponseData) Set(val *MerchantPaymentMethodListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentMethodListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentMethodListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentMethodListGet200ResponseData(val *MerchantPaymentMethodListGet200ResponseData) *NullableMerchantPaymentMethodListGet200ResponseData {
	return &NullableMerchantPaymentMethodListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentMethodListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentMethodListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


