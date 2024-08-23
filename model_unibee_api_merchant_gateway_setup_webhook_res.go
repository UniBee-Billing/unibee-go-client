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

// checks if the UnibeeApiMerchantGatewaySetupWebhookRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewaySetupWebhookRes{}

// UnibeeApiMerchantGatewaySetupWebhookRes struct for UnibeeApiMerchantGatewaySetupWebhookRes
type UnibeeApiMerchantGatewaySetupWebhookRes struct {
	// The webhook endpoint url of payment gateway, if gateway is stripe, the url will setting to stripe by api automatic
	GatewayWebhookUrl *string `json:"gatewayWebhookUrl,omitempty"`
}

// NewUnibeeApiMerchantGatewaySetupWebhookRes instantiates a new UnibeeApiMerchantGatewaySetupWebhookRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewaySetupWebhookRes() *UnibeeApiMerchantGatewaySetupWebhookRes {
	this := UnibeeApiMerchantGatewaySetupWebhookRes{}
	return &this
}

// NewUnibeeApiMerchantGatewaySetupWebhookResWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupWebhookRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewaySetupWebhookResWithDefaults() *UnibeeApiMerchantGatewaySetupWebhookRes {
	this := UnibeeApiMerchantGatewaySetupWebhookRes{}
	return &this
}

// GetGatewayWebhookUrl returns the GatewayWebhookUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewaySetupWebhookRes) GetGatewayWebhookUrl() string {
	if o == nil || IsNil(o.GatewayWebhookUrl) {
		var ret string
		return ret
	}
	return *o.GatewayWebhookUrl
}

// GetGatewayWebhookUrlOk returns a tuple with the GatewayWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupWebhookRes) GetGatewayWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayWebhookUrl) {
		return nil, false
	}
	return o.GatewayWebhookUrl, true
}

// HasGatewayWebhookUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewaySetupWebhookRes) HasGatewayWebhookUrl() bool {
	if o != nil && !IsNil(o.GatewayWebhookUrl) {
		return true
	}

	return false
}

// SetGatewayWebhookUrl gets a reference to the given string and assigns it to the GatewayWebhookUrl field.
func (o *UnibeeApiMerchantGatewaySetupWebhookRes) SetGatewayWebhookUrl(v string) {
	o.GatewayWebhookUrl = &v
}

func (o UnibeeApiMerchantGatewaySetupWebhookRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewaySetupWebhookRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayWebhookUrl) {
		toSerialize["gatewayWebhookUrl"] = o.GatewayWebhookUrl
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewaySetupWebhookRes struct {
	value *UnibeeApiMerchantGatewaySetupWebhookRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookRes) Get() *UnibeeApiMerchantGatewaySetupWebhookRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookRes) Set(val *UnibeeApiMerchantGatewaySetupWebhookRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewaySetupWebhookRes(val *UnibeeApiMerchantGatewaySetupWebhookRes) *NullableUnibeeApiMerchantGatewaySetupWebhookRes {
	return &NullableUnibeeApiMerchantGatewaySetupWebhookRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


