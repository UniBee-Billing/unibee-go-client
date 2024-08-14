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

// checks if the MerchantGatewaySetupWebhookPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGatewaySetupWebhookPost200ResponseData{}

// MerchantGatewaySetupWebhookPost200ResponseData struct for MerchantGatewaySetupWebhookPost200ResponseData
type MerchantGatewaySetupWebhookPost200ResponseData struct {
	// The webhook endpoint url of payment gateway, if gateway is stripe, the url will setting to stripe by api automatic
	GatewayWebhookUrl *string `json:"gatewayWebhookUrl,omitempty"`
}

// NewMerchantGatewaySetupWebhookPost200ResponseData instantiates a new MerchantGatewaySetupWebhookPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGatewaySetupWebhookPost200ResponseData() *MerchantGatewaySetupWebhookPost200ResponseData {
	this := MerchantGatewaySetupWebhookPost200ResponseData{}
	return &this
}

// NewMerchantGatewaySetupWebhookPost200ResponseDataWithDefaults instantiates a new MerchantGatewaySetupWebhookPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGatewaySetupWebhookPost200ResponseDataWithDefaults() *MerchantGatewaySetupWebhookPost200ResponseData {
	this := MerchantGatewaySetupWebhookPost200ResponseData{}
	return &this
}

// GetGatewayWebhookUrl returns the GatewayWebhookUrl field value if set, zero value otherwise.
func (o *MerchantGatewaySetupWebhookPost200ResponseData) GetGatewayWebhookUrl() string {
	if o == nil || IsNil(o.GatewayWebhookUrl) {
		var ret string
		return ret
	}
	return *o.GatewayWebhookUrl
}

// GetGatewayWebhookUrlOk returns a tuple with the GatewayWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGatewaySetupWebhookPost200ResponseData) GetGatewayWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayWebhookUrl) {
		return nil, false
	}
	return o.GatewayWebhookUrl, true
}

// HasGatewayWebhookUrl returns a boolean if a field has been set.
func (o *MerchantGatewaySetupWebhookPost200ResponseData) HasGatewayWebhookUrl() bool {
	if o != nil && !IsNil(o.GatewayWebhookUrl) {
		return true
	}

	return false
}

// SetGatewayWebhookUrl gets a reference to the given string and assigns it to the GatewayWebhookUrl field.
func (o *MerchantGatewaySetupWebhookPost200ResponseData) SetGatewayWebhookUrl(v string) {
	o.GatewayWebhookUrl = &v
}

func (o MerchantGatewaySetupWebhookPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGatewaySetupWebhookPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayWebhookUrl) {
		toSerialize["gatewayWebhookUrl"] = o.GatewayWebhookUrl
	}
	return toSerialize, nil
}

type NullableMerchantGatewaySetupWebhookPost200ResponseData struct {
	value *MerchantGatewaySetupWebhookPost200ResponseData
	isSet bool
}

func (v NullableMerchantGatewaySetupWebhookPost200ResponseData) Get() *MerchantGatewaySetupWebhookPost200ResponseData {
	return v.value
}

func (v *NullableMerchantGatewaySetupWebhookPost200ResponseData) Set(val *MerchantGatewaySetupWebhookPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGatewaySetupWebhookPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGatewaySetupWebhookPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGatewaySetupWebhookPost200ResponseData(val *MerchantGatewaySetupWebhookPost200ResponseData) *NullableMerchantGatewaySetupWebhookPost200ResponseData {
	return &NullableMerchantGatewaySetupWebhookPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGatewaySetupWebhookPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGatewaySetupWebhookPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


