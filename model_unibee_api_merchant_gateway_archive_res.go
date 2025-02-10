/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantGatewayArchiveRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayArchiveRes{}

// UnibeeApiMerchantGatewayArchiveRes struct for UnibeeApiMerchantGatewayArchiveRes
type UnibeeApiMerchantGatewayArchiveRes struct {
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
}

// NewUnibeeApiMerchantGatewayArchiveRes instantiates a new UnibeeApiMerchantGatewayArchiveRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayArchiveRes() *UnibeeApiMerchantGatewayArchiveRes {
	this := UnibeeApiMerchantGatewayArchiveRes{}
	return &this
}

// NewUnibeeApiMerchantGatewayArchiveResWithDefaults instantiates a new UnibeeApiMerchantGatewayArchiveRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayArchiveResWithDefaults() *UnibeeApiMerchantGatewayArchiveRes {
	this := UnibeeApiMerchantGatewayArchiveRes{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayArchiveRes) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayArchiveRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayArchiveRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantGatewayArchiveRes) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

func (o UnibeeApiMerchantGatewayArchiveRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayArchiveRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayArchiveRes struct {
	value *UnibeeApiMerchantGatewayArchiveRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayArchiveRes) Get() *UnibeeApiMerchantGatewayArchiveRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayArchiveRes) Set(val *UnibeeApiMerchantGatewayArchiveRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayArchiveRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayArchiveRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayArchiveRes(val *UnibeeApiMerchantGatewayArchiveRes) *NullableUnibeeApiMerchantGatewayArchiveRes {
	return &NullableUnibeeApiMerchantGatewayArchiveRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayArchiveRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayArchiveRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


