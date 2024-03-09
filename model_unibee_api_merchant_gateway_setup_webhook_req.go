/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantGatewaySetupWebhookReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewaySetupWebhookReq{}

// UnibeeApiMerchantGatewaySetupWebhookReq struct for UnibeeApiMerchantGatewaySetupWebhookReq
type UnibeeApiMerchantGatewaySetupWebhookReq struct {
	// GatewayId
	GatewayId int64 `json:"gatewayId"`
}

type _UnibeeApiMerchantGatewaySetupWebhookReq UnibeeApiMerchantGatewaySetupWebhookReq

// NewUnibeeApiMerchantGatewaySetupWebhookReq instantiates a new UnibeeApiMerchantGatewaySetupWebhookReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewaySetupWebhookReq(gatewayId int64) *UnibeeApiMerchantGatewaySetupWebhookReq {
	this := UnibeeApiMerchantGatewaySetupWebhookReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantGatewaySetupWebhookReqWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupWebhookReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewaySetupWebhookReqWithDefaults() *UnibeeApiMerchantGatewaySetupWebhookReq {
	this := UnibeeApiMerchantGatewaySetupWebhookReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantGatewaySetupWebhookReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

func (o UnibeeApiMerchantGatewaySetupWebhookReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewaySetupWebhookReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewaySetupWebhookReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantGatewaySetupWebhookReq := _UnibeeApiMerchantGatewaySetupWebhookReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewaySetupWebhookReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewaySetupWebhookReq(varUnibeeApiMerchantGatewaySetupWebhookReq)

	return err
}

type NullableUnibeeApiMerchantGatewaySetupWebhookReq struct {
	value *UnibeeApiMerchantGatewaySetupWebhookReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookReq) Get() *UnibeeApiMerchantGatewaySetupWebhookReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookReq) Set(val *UnibeeApiMerchantGatewaySetupWebhookReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewaySetupWebhookReq(val *UnibeeApiMerchantGatewaySetupWebhookReq) *NullableUnibeeApiMerchantGatewaySetupWebhookReq {
	return &NullableUnibeeApiMerchantGatewaySetupWebhookReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewaySetupWebhookReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewaySetupWebhookReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


