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

// checks if the UnibeeApiMerchantGatewayListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayListReq{}

// UnibeeApiMerchantGatewayListReq struct for UnibeeApiMerchantGatewayListReq
type UnibeeApiMerchantGatewayListReq struct {
	// Filter archive gateway or not, default all
	Archive *bool `json:"archive,omitempty"`
}

// NewUnibeeApiMerchantGatewayListReq instantiates a new UnibeeApiMerchantGatewayListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayListReq() *UnibeeApiMerchantGatewayListReq {
	this := UnibeeApiMerchantGatewayListReq{}
	return &this
}

// NewUnibeeApiMerchantGatewayListReqWithDefaults instantiates a new UnibeeApiMerchantGatewayListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayListReqWithDefaults() *UnibeeApiMerchantGatewayListReq {
	this := UnibeeApiMerchantGatewayListReq{}
	return &this
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayListReq) GetArchive() bool {
	if o == nil || IsNil(o.Archive) {
		var ret bool
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayListReq) GetArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayListReq) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given bool and assigns it to the Archive field.
func (o *UnibeeApiMerchantGatewayListReq) SetArchive(v bool) {
	o.Archive = &v
}

func (o UnibeeApiMerchantGatewayListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayListReq struct {
	value *UnibeeApiMerchantGatewayListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayListReq) Get() *UnibeeApiMerchantGatewayListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayListReq) Set(val *UnibeeApiMerchantGatewayListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayListReq(val *UnibeeApiMerchantGatewayListReq) *NullableUnibeeApiMerchantGatewayListReq {
	return &NullableUnibeeApiMerchantGatewayListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


