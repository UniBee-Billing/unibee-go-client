/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantGatewayEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayEditReq{}

// UnibeeApiMerchantGatewayEditReq edit the exist payment gateway
type UnibeeApiMerchantGatewayEditReq struct {
	// The id of payment gateway
	GatewayId int64 `json:"gatewayId"`
	// The key of payment gateway
	GatewayKey *string `json:"gatewayKey,omitempty"`
	// The secret of payment gateway
	GatewaySecret *string `json:"gatewaySecret,omitempty"`
}

type _UnibeeApiMerchantGatewayEditReq UnibeeApiMerchantGatewayEditReq

// NewUnibeeApiMerchantGatewayEditReq instantiates a new UnibeeApiMerchantGatewayEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayEditReq(gatewayId int64) *UnibeeApiMerchantGatewayEditReq {
	this := UnibeeApiMerchantGatewayEditReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantGatewayEditReqWithDefaults instantiates a new UnibeeApiMerchantGatewayEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayEditReqWithDefaults() *UnibeeApiMerchantGatewayEditReq {
	this := UnibeeApiMerchantGatewayEditReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetGatewaySecret returns the GatewaySecret field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecret() string {
	if o == nil || IsNil(o.GatewaySecret) {
		var ret string
		return ret
	}
	return *o.GatewaySecret
}

// GetGatewaySecretOk returns a tuple with the GatewaySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecretOk() (*string, bool) {
	if o == nil || IsNil(o.GatewaySecret) {
		return nil, false
	}
	return o.GatewaySecret, true
}

// HasGatewaySecret returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewaySecret() bool {
	if o != nil && !IsNil(o.GatewaySecret) {
		return true
	}

	return false
}

// SetGatewaySecret gets a reference to the given string and assigns it to the GatewaySecret field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewaySecret(v string) {
	o.GatewaySecret = &v
}

func (o UnibeeApiMerchantGatewayEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.GatewayKey) {
		toSerialize["gatewayKey"] = o.GatewayKey
	}
	if !IsNil(o.GatewaySecret) {
		toSerialize["gatewaySecret"] = o.GatewaySecret
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewayEditReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantGatewayEditReq := _UnibeeApiMerchantGatewayEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewayEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewayEditReq(varUnibeeApiMerchantGatewayEditReq)

	return err
}

type NullableUnibeeApiMerchantGatewayEditReq struct {
	value *UnibeeApiMerchantGatewayEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayEditReq) Get() *UnibeeApiMerchantGatewayEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) Set(val *UnibeeApiMerchantGatewayEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayEditReq(val *UnibeeApiMerchantGatewayEditReq) *NullableUnibeeApiMerchantGatewayEditReq {
	return &NullableUnibeeApiMerchantGatewayEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


