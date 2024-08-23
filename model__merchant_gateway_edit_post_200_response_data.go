/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantGatewayEditPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGatewayEditPost200ResponseData{}

// MerchantGatewayEditPost200ResponseData struct for MerchantGatewayEditPost200ResponseData
type MerchantGatewayEditPost200ResponseData struct {
	Gateway *UnibeeApiBeanGateway `json:"gateway,omitempty"`
}

// NewMerchantGatewayEditPost200ResponseData instantiates a new MerchantGatewayEditPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGatewayEditPost200ResponseData() *MerchantGatewayEditPost200ResponseData {
	this := MerchantGatewayEditPost200ResponseData{}
	return &this
}

// NewMerchantGatewayEditPost200ResponseDataWithDefaults instantiates a new MerchantGatewayEditPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGatewayEditPost200ResponseDataWithDefaults() *MerchantGatewayEditPost200ResponseData {
	this := MerchantGatewayEditPost200ResponseData{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantGatewayEditPost200ResponseData) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGatewayEditPost200ResponseData) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantGatewayEditPost200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *MerchantGatewayEditPost200ResponseData) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

func (o MerchantGatewayEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGatewayEditPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableMerchantGatewayEditPost200ResponseData struct {
	value *MerchantGatewayEditPost200ResponseData
	isSet bool
}

func (v NullableMerchantGatewayEditPost200ResponseData) Get() *MerchantGatewayEditPost200ResponseData {
	return v.value
}

func (v *NullableMerchantGatewayEditPost200ResponseData) Set(val *MerchantGatewayEditPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGatewayEditPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGatewayEditPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGatewayEditPost200ResponseData(val *MerchantGatewayEditPost200ResponseData) *NullableMerchantGatewayEditPost200ResponseData {
	return &NullableMerchantGatewayEditPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGatewayEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGatewayEditPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


