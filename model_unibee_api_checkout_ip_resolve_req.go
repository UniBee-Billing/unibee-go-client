/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiCheckoutIpResolveReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiCheckoutIpResolveReq{}

// UnibeeApiCheckoutIpResolveReq struct for UnibeeApiCheckoutIpResolveReq
type UnibeeApiCheckoutIpResolveReq struct {
	// ip
	Ip string `json:"ip"`
}

type _UnibeeApiCheckoutIpResolveReq UnibeeApiCheckoutIpResolveReq

// NewUnibeeApiCheckoutIpResolveReq instantiates a new UnibeeApiCheckoutIpResolveReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiCheckoutIpResolveReq(ip string) *UnibeeApiCheckoutIpResolveReq {
	this := UnibeeApiCheckoutIpResolveReq{}
	this.Ip = ip
	return &this
}

// NewUnibeeApiCheckoutIpResolveReqWithDefaults instantiates a new UnibeeApiCheckoutIpResolveReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiCheckoutIpResolveReqWithDefaults() *UnibeeApiCheckoutIpResolveReq {
	this := UnibeeApiCheckoutIpResolveReq{}
	return &this
}

// GetIp returns the Ip field value
func (o *UnibeeApiCheckoutIpResolveReq) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiCheckoutIpResolveReq) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *UnibeeApiCheckoutIpResolveReq) SetIp(v string) {
	o.Ip = v
}

func (o UnibeeApiCheckoutIpResolveReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiCheckoutIpResolveReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

func (o *UnibeeApiCheckoutIpResolveReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
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

	varUnibeeApiCheckoutIpResolveReq := _UnibeeApiCheckoutIpResolveReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiCheckoutIpResolveReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiCheckoutIpResolveReq(varUnibeeApiCheckoutIpResolveReq)

	return err
}

type NullableUnibeeApiCheckoutIpResolveReq struct {
	value *UnibeeApiCheckoutIpResolveReq
	isSet bool
}

func (v NullableUnibeeApiCheckoutIpResolveReq) Get() *UnibeeApiCheckoutIpResolveReq {
	return v.value
}

func (v *NullableUnibeeApiCheckoutIpResolveReq) Set(val *UnibeeApiCheckoutIpResolveReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiCheckoutIpResolveReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiCheckoutIpResolveReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiCheckoutIpResolveReq(val *UnibeeApiCheckoutIpResolveReq) *NullableUnibeeApiCheckoutIpResolveReq {
	return &NullableUnibeeApiCheckoutIpResolveReq{value: val, isSet: true}
}

func (v NullableUnibeeApiCheckoutIpResolveReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiCheckoutIpResolveReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


