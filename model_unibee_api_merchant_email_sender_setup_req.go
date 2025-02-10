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

// checks if the UnibeeApiMerchantEmailSenderSetupReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailSenderSetupReq{}

// UnibeeApiMerchantEmailSenderSetupReq struct for UnibeeApiMerchantEmailSenderSetupReq
type UnibeeApiMerchantEmailSenderSetupReq struct {
	// The address of email sender, like 'no-reply@unibee.dev'
	Address string `json:"address"`
	// The name of email sender, like 'no-reply'
	Name string `json:"name"`
}

type _UnibeeApiMerchantEmailSenderSetupReq UnibeeApiMerchantEmailSenderSetupReq

// NewUnibeeApiMerchantEmailSenderSetupReq instantiates a new UnibeeApiMerchantEmailSenderSetupReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailSenderSetupReq(address string, name string) *UnibeeApiMerchantEmailSenderSetupReq {
	this := UnibeeApiMerchantEmailSenderSetupReq{}
	this.Address = address
	this.Name = name
	return &this
}

// NewUnibeeApiMerchantEmailSenderSetupReqWithDefaults instantiates a new UnibeeApiMerchantEmailSenderSetupReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailSenderSetupReqWithDefaults() *UnibeeApiMerchantEmailSenderSetupReq {
	this := UnibeeApiMerchantEmailSenderSetupReq{}
	return &this
}

// GetAddress returns the Address field value
func (o *UnibeeApiMerchantEmailSenderSetupReq) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailSenderSetupReq) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UnibeeApiMerchantEmailSenderSetupReq) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value
func (o *UnibeeApiMerchantEmailSenderSetupReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailSenderSetupReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UnibeeApiMerchantEmailSenderSetupReq) SetName(v string) {
	o.Name = v
}

func (o UnibeeApiMerchantEmailSenderSetupReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailSenderSetupReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailSenderSetupReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"name",
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

	varUnibeeApiMerchantEmailSenderSetupReq := _UnibeeApiMerchantEmailSenderSetupReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailSenderSetupReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailSenderSetupReq(varUnibeeApiMerchantEmailSenderSetupReq)

	return err
}

type NullableUnibeeApiMerchantEmailSenderSetupReq struct {
	value *UnibeeApiMerchantEmailSenderSetupReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailSenderSetupReq) Get() *UnibeeApiMerchantEmailSenderSetupReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailSenderSetupReq) Set(val *UnibeeApiMerchantEmailSenderSetupReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailSenderSetupReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailSenderSetupReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailSenderSetupReq(val *UnibeeApiMerchantEmailSenderSetupReq) *NullableUnibeeApiMerchantEmailSenderSetupReq {
	return &NullableUnibeeApiMerchantEmailSenderSetupReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailSenderSetupReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailSenderSetupReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


