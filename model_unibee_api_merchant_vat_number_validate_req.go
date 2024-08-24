/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantVatNumberValidateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantVatNumberValidateReq{}

// UnibeeApiMerchantVatNumberValidateReq struct for UnibeeApiMerchantVatNumberValidateReq
type UnibeeApiMerchantVatNumberValidateReq struct {
	// VatNumber
	VatNumber string `json:"vatNumber"`
}

type _UnibeeApiMerchantVatNumberValidateReq UnibeeApiMerchantVatNumberValidateReq

// NewUnibeeApiMerchantVatNumberValidateReq instantiates a new UnibeeApiMerchantVatNumberValidateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantVatNumberValidateReq(vatNumber string) *UnibeeApiMerchantVatNumberValidateReq {
	this := UnibeeApiMerchantVatNumberValidateReq{}
	this.VatNumber = vatNumber
	return &this
}

// NewUnibeeApiMerchantVatNumberValidateReqWithDefaults instantiates a new UnibeeApiMerchantVatNumberValidateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantVatNumberValidateReqWithDefaults() *UnibeeApiMerchantVatNumberValidateReq {
	this := UnibeeApiMerchantVatNumberValidateReq{}
	return &this
}

// GetVatNumber returns the VatNumber field value
func (o *UnibeeApiMerchantVatNumberValidateReq) GetVatNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatNumberValidateReq) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VatNumber, true
}

// SetVatNumber sets field value
func (o *UnibeeApiMerchantVatNumberValidateReq) SetVatNumber(v string) {
	o.VatNumber = v
}

func (o UnibeeApiMerchantVatNumberValidateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantVatNumberValidateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vatNumber"] = o.VatNumber
	return toSerialize, nil
}

func (o *UnibeeApiMerchantVatNumberValidateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vatNumber",
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

	varUnibeeApiMerchantVatNumberValidateReq := _UnibeeApiMerchantVatNumberValidateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantVatNumberValidateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantVatNumberValidateReq(varUnibeeApiMerchantVatNumberValidateReq)

	return err
}

type NullableUnibeeApiMerchantVatNumberValidateReq struct {
	value *UnibeeApiMerchantVatNumberValidateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantVatNumberValidateReq) Get() *UnibeeApiMerchantVatNumberValidateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantVatNumberValidateReq) Set(val *UnibeeApiMerchantVatNumberValidateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantVatNumberValidateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantVatNumberValidateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantVatNumberValidateReq(val *UnibeeApiMerchantVatNumberValidateReq) *NullableUnibeeApiMerchantVatNumberValidateReq {
	return &NullableUnibeeApiMerchantVatNumberValidateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantVatNumberValidateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantVatNumberValidateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


