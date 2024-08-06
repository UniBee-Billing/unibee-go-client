/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantGatewayEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayEditRes{}

// UnibeeApiMerchantGatewayEditRes struct for UnibeeApiMerchantGatewayEditRes
type UnibeeApiMerchantGatewayEditRes struct {
	Gateway *UnibeeApiBeanGateway `json:"gateway,omitempty"`
}

// NewUnibeeApiMerchantGatewayEditRes instantiates a new UnibeeApiMerchantGatewayEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayEditRes() *UnibeeApiMerchantGatewayEditRes {
	this := UnibeeApiMerchantGatewayEditRes{}
	return &this
}

// NewUnibeeApiMerchantGatewayEditResWithDefaults instantiates a new UnibeeApiMerchantGatewayEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayEditResWithDefaults() *UnibeeApiMerchantGatewayEditRes {
	this := UnibeeApiMerchantGatewayEditRes{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditRes) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditRes) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantGatewayEditRes) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

func (o UnibeeApiMerchantGatewayEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayEditRes struct {
	value *UnibeeApiMerchantGatewayEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayEditRes) Get() *UnibeeApiMerchantGatewayEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayEditRes) Set(val *UnibeeApiMerchantGatewayEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayEditRes(val *UnibeeApiMerchantGatewayEditRes) *NullableUnibeeApiMerchantGatewayEditRes {
	return &NullableUnibeeApiMerchantGatewayEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


