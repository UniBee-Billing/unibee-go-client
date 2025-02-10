/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantGatewayEditCountryConfigReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayEditCountryConfigReq{}

// UnibeeApiMerchantGatewayEditCountryConfigReq Edit country config for payment gateway, to enable or disable the payment for countryCode, default is enable
type UnibeeApiMerchantGatewayEditCountryConfigReq struct {
	// The country config of payment gateway, a map with countryCode as key, and value for enable or disable
	CountryConfig *map[string]bool `json:"countryConfig,omitempty"`
	// The id of payment gateway
	GatewayId int64 `json:"gatewayId"`
}

type _UnibeeApiMerchantGatewayEditCountryConfigReq UnibeeApiMerchantGatewayEditCountryConfigReq

// NewUnibeeApiMerchantGatewayEditCountryConfigReq instantiates a new UnibeeApiMerchantGatewayEditCountryConfigReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayEditCountryConfigReq(gatewayId int64) *UnibeeApiMerchantGatewayEditCountryConfigReq {
	this := UnibeeApiMerchantGatewayEditCountryConfigReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantGatewayEditCountryConfigReqWithDefaults instantiates a new UnibeeApiMerchantGatewayEditCountryConfigReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayEditCountryConfigReqWithDefaults() *UnibeeApiMerchantGatewayEditCountryConfigReq {
	this := UnibeeApiMerchantGatewayEditCountryConfigReq{}
	return &this
}

// GetCountryConfig returns the CountryConfig field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetCountryConfig() map[string]bool {
	if o == nil || IsNil(o.CountryConfig) {
		var ret map[string]bool
		return ret
	}
	return *o.CountryConfig
}

// GetCountryConfigOk returns a tuple with the CountryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetCountryConfigOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.CountryConfig) {
		return nil, false
	}
	return o.CountryConfig, true
}

// HasCountryConfig returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) HasCountryConfig() bool {
	if o != nil && !IsNil(o.CountryConfig) {
		return true
	}

	return false
}

// SetCountryConfig gets a reference to the given map[string]bool and assigns it to the CountryConfig field.
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) SetCountryConfig(v map[string]bool) {
	o.CountryConfig = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

func (o UnibeeApiMerchantGatewayEditCountryConfigReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayEditCountryConfigReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryConfig) {
		toSerialize["countryConfig"] = o.CountryConfig
	}
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantGatewayEditCountryConfigReq := _UnibeeApiMerchantGatewayEditCountryConfigReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewayEditCountryConfigReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewayEditCountryConfigReq(varUnibeeApiMerchantGatewayEditCountryConfigReq)

	return err
}

type NullableUnibeeApiMerchantGatewayEditCountryConfigReq struct {
	value *UnibeeApiMerchantGatewayEditCountryConfigReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayEditCountryConfigReq) Get() *UnibeeApiMerchantGatewayEditCountryConfigReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayEditCountryConfigReq) Set(val *UnibeeApiMerchantGatewayEditCountryConfigReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayEditCountryConfigReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayEditCountryConfigReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayEditCountryConfigReq(val *UnibeeApiMerchantGatewayEditCountryConfigReq) *NullableUnibeeApiMerchantGatewayEditCountryConfigReq {
	return &NullableUnibeeApiMerchantGatewayEditCountryConfigReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayEditCountryConfigReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayEditCountryConfigReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


