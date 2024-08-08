/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantGatewaySetupReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewaySetupReq{}

// UnibeeApiMerchantGatewaySetupReq Setup the payment gateway
type UnibeeApiMerchantGatewaySetupReq struct {
	// The key of payment gateway
	GatewayKey *string `json:"gatewayKey,omitempty"`
	// The name of payment gateway, stripe|paypal|changelly
	GatewayName string `json:"gatewayName"`
	// The secret of payment gateway
	GatewaySecret *string `json:"gatewaySecret,omitempty"`
}

type _UnibeeApiMerchantGatewaySetupReq UnibeeApiMerchantGatewaySetupReq

// NewUnibeeApiMerchantGatewaySetupReq instantiates a new UnibeeApiMerchantGatewaySetupReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewaySetupReq(gatewayName string) *UnibeeApiMerchantGatewaySetupReq {
	this := UnibeeApiMerchantGatewaySetupReq{}
	this.GatewayName = gatewayName
	return &this
}

// NewUnibeeApiMerchantGatewaySetupReqWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewaySetupReqWithDefaults() *UnibeeApiMerchantGatewaySetupReq {
	this := UnibeeApiMerchantGatewaySetupReq{}
	return &this
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetGatewayName returns the GatewayName field value
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayName, true
}

// SetGatewayName sets field value
func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayName(v string) {
	o.GatewayName = v
}

// GetGatewaySecret returns the GatewaySecret field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecret() string {
	if o == nil || IsNil(o.GatewaySecret) {
		var ret string
		return ret
	}
	return *o.GatewaySecret
}

// GetGatewaySecretOk returns a tuple with the GatewaySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecretOk() (*string, bool) {
	if o == nil || IsNil(o.GatewaySecret) {
		return nil, false
	}
	return o.GatewaySecret, true
}

// HasGatewaySecret returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewaySecret() bool {
	if o != nil && !IsNil(o.GatewaySecret) {
		return true
	}

	return false
}

// SetGatewaySecret gets a reference to the given string and assigns it to the GatewaySecret field.
func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewaySecret(v string) {
	o.GatewaySecret = &v
}

func (o UnibeeApiMerchantGatewaySetupReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewaySetupReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayKey) {
		toSerialize["gatewayKey"] = o.GatewayKey
	}
	toSerialize["gatewayName"] = o.GatewayName
	if !IsNil(o.GatewaySecret) {
		toSerialize["gatewaySecret"] = o.GatewaySecret
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewaySetupReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayName",
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

	varUnibeeApiMerchantGatewaySetupReq := _UnibeeApiMerchantGatewaySetupReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewaySetupReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewaySetupReq(varUnibeeApiMerchantGatewaySetupReq)

	return err
}

type NullableUnibeeApiMerchantGatewaySetupReq struct {
	value *UnibeeApiMerchantGatewaySetupReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewaySetupReq) Get() *UnibeeApiMerchantGatewaySetupReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewaySetupReq) Set(val *UnibeeApiMerchantGatewaySetupReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewaySetupReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewaySetupReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewaySetupReq(val *UnibeeApiMerchantGatewaySetupReq) *NullableUnibeeApiMerchantGatewaySetupReq {
	return &NullableUnibeeApiMerchantGatewaySetupReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewaySetupReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewaySetupReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


