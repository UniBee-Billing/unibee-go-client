/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionDetailReq{}

// UnibeeApiMerchantSubscriptionDetailReq struct for UnibeeApiMerchantSubscriptionDetailReq
type UnibeeApiMerchantSubscriptionDetailReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionDetailReq UnibeeApiMerchantSubscriptionDetailReq

// NewUnibeeApiMerchantSubscriptionDetailReq instantiates a new UnibeeApiMerchantSubscriptionDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionDetailReq(subscriptionId string) *UnibeeApiMerchantSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionDetailReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionDetailReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionDetailReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionDetailReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionDetailReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionDetailReq := _UnibeeApiMerchantSubscriptionDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionDetailReq(varUnibeeApiMerchantSubscriptionDetailReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionDetailReq struct {
	value *UnibeeApiMerchantSubscriptionDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionDetailReq) Get() *UnibeeApiMerchantSubscriptionDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailReq) Set(val *UnibeeApiMerchantSubscriptionDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionDetailReq(val *UnibeeApiMerchantSubscriptionDetailReq) *NullableUnibeeApiMerchantSubscriptionDetailReq {
	return &NullableUnibeeApiMerchantSubscriptionDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


