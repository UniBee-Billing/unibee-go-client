/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantWebhookEndpointListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookEndpointListRes{}

// UnibeeApiMerchantWebhookEndpointListRes struct for UnibeeApiMerchantWebhookEndpointListRes
type UnibeeApiMerchantWebhookEndpointListRes struct {
	// EndpointList
	EndpointList []UnibeeApiBeanMerchantWebhookEndpoint `json:"endpointList,omitempty"`
}

// NewUnibeeApiMerchantWebhookEndpointListRes instantiates a new UnibeeApiMerchantWebhookEndpointListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookEndpointListRes() *UnibeeApiMerchantWebhookEndpointListRes {
	this := UnibeeApiMerchantWebhookEndpointListRes{}
	return &this
}

// NewUnibeeApiMerchantWebhookEndpointListResWithDefaults instantiates a new UnibeeApiMerchantWebhookEndpointListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookEndpointListResWithDefaults() *UnibeeApiMerchantWebhookEndpointListRes {
	this := UnibeeApiMerchantWebhookEndpointListRes{}
	return &this
}

// GetEndpointList returns the EndpointList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookEndpointListRes) GetEndpointList() []UnibeeApiBeanMerchantWebhookEndpoint {
	if o == nil || IsNil(o.EndpointList) {
		var ret []UnibeeApiBeanMerchantWebhookEndpoint
		return ret
	}
	return o.EndpointList
}

// GetEndpointListOk returns a tuple with the EndpointList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEndpointListRes) GetEndpointListOk() ([]UnibeeApiBeanMerchantWebhookEndpoint, bool) {
	if o == nil || IsNil(o.EndpointList) {
		return nil, false
	}
	return o.EndpointList, true
}

// HasEndpointList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookEndpointListRes) HasEndpointList() bool {
	if o != nil && !IsNil(o.EndpointList) {
		return true
	}

	return false
}

// SetEndpointList gets a reference to the given []UnibeeApiBeanMerchantWebhookEndpoint and assigns it to the EndpointList field.
func (o *UnibeeApiMerchantWebhookEndpointListRes) SetEndpointList(v []UnibeeApiBeanMerchantWebhookEndpoint) {
	o.EndpointList = v
}

func (o UnibeeApiMerchantWebhookEndpointListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookEndpointListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointList) {
		toSerialize["endpointList"] = o.EndpointList
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantWebhookEndpointListRes struct {
	value *UnibeeApiMerchantWebhookEndpointListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookEndpointListRes) Get() *UnibeeApiMerchantWebhookEndpointListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookEndpointListRes) Set(val *UnibeeApiMerchantWebhookEndpointListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookEndpointListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookEndpointListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookEndpointListRes(val *UnibeeApiMerchantWebhookEndpointListRes) *NullableUnibeeApiMerchantWebhookEndpointListRes {
	return &NullableUnibeeApiMerchantWebhookEndpointListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookEndpointListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookEndpointListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


