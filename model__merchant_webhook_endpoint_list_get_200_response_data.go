/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantWebhookEndpointListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantWebhookEndpointListGet200ResponseData{}

// MerchantWebhookEndpointListGet200ResponseData struct for MerchantWebhookEndpointListGet200ResponseData
type MerchantWebhookEndpointListGet200ResponseData struct {
	// EndpointList
	EndpointList []UnibeeApiBeanMerchantWebhookEndpoint `json:"endpointList,omitempty"`
}

// NewMerchantWebhookEndpointListGet200ResponseData instantiates a new MerchantWebhookEndpointListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantWebhookEndpointListGet200ResponseData() *MerchantWebhookEndpointListGet200ResponseData {
	this := MerchantWebhookEndpointListGet200ResponseData{}
	return &this
}

// NewMerchantWebhookEndpointListGet200ResponseDataWithDefaults instantiates a new MerchantWebhookEndpointListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWebhookEndpointListGet200ResponseDataWithDefaults() *MerchantWebhookEndpointListGet200ResponseData {
	this := MerchantWebhookEndpointListGet200ResponseData{}
	return &this
}

// GetEndpointList returns the EndpointList field value if set, zero value otherwise.
func (o *MerchantWebhookEndpointListGet200ResponseData) GetEndpointList() []UnibeeApiBeanMerchantWebhookEndpoint {
	if o == nil || IsNil(o.EndpointList) {
		var ret []UnibeeApiBeanMerchantWebhookEndpoint
		return ret
	}
	return o.EndpointList
}

// GetEndpointListOk returns a tuple with the EndpointList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookEndpointListGet200ResponseData) GetEndpointListOk() ([]UnibeeApiBeanMerchantWebhookEndpoint, bool) {
	if o == nil || IsNil(o.EndpointList) {
		return nil, false
	}
	return o.EndpointList, true
}

// HasEndpointList returns a boolean if a field has been set.
func (o *MerchantWebhookEndpointListGet200ResponseData) HasEndpointList() bool {
	if o != nil && !IsNil(o.EndpointList) {
		return true
	}

	return false
}

// SetEndpointList gets a reference to the given []UnibeeApiBeanMerchantWebhookEndpoint and assigns it to the EndpointList field.
func (o *MerchantWebhookEndpointListGet200ResponseData) SetEndpointList(v []UnibeeApiBeanMerchantWebhookEndpoint) {
	o.EndpointList = v
}

func (o MerchantWebhookEndpointListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantWebhookEndpointListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointList) {
		toSerialize["endpointList"] = o.EndpointList
	}
	return toSerialize, nil
}

type NullableMerchantWebhookEndpointListGet200ResponseData struct {
	value *MerchantWebhookEndpointListGet200ResponseData
	isSet bool
}

func (v NullableMerchantWebhookEndpointListGet200ResponseData) Get() *MerchantWebhookEndpointListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantWebhookEndpointListGet200ResponseData) Set(val *MerchantWebhookEndpointListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantWebhookEndpointListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantWebhookEndpointListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantWebhookEndpointListGet200ResponseData(val *MerchantWebhookEndpointListGet200ResponseData) *NullableMerchantWebhookEndpointListGet200ResponseData {
	return &NullableMerchantWebhookEndpointListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantWebhookEndpointListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantWebhookEndpointListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


