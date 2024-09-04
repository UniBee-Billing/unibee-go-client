/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantGatewaySetupRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewaySetupRes{}

// UnibeeApiMerchantGatewaySetupRes struct for UnibeeApiMerchantGatewaySetupRes
type UnibeeApiMerchantGatewaySetupRes struct {
	Gateway *UnibeeApiBeanGateway `json:"gateway,omitempty"`
}

// NewUnibeeApiMerchantGatewaySetupRes instantiates a new UnibeeApiMerchantGatewaySetupRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewaySetupRes() *UnibeeApiMerchantGatewaySetupRes {
	this := UnibeeApiMerchantGatewaySetupRes{}
	return &this
}

// NewUnibeeApiMerchantGatewaySetupResWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewaySetupResWithDefaults() *UnibeeApiMerchantGatewaySetupRes {
	this := UnibeeApiMerchantGatewaySetupRes{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewaySetupRes) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupRes) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewaySetupRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantGatewaySetupRes) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

func (o UnibeeApiMerchantGatewaySetupRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewaySetupRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewaySetupRes struct {
	value *UnibeeApiMerchantGatewaySetupRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewaySetupRes) Get() *UnibeeApiMerchantGatewaySetupRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewaySetupRes) Set(val *UnibeeApiMerchantGatewaySetupRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewaySetupRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewaySetupRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewaySetupRes(val *UnibeeApiMerchantGatewaySetupRes) *NullableUnibeeApiMerchantGatewaySetupRes {
	return &NullableUnibeeApiMerchantGatewaySetupRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewaySetupRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewaySetupRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


