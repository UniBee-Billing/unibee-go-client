/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionResumeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionResumeReq{}

// UnibeeApiMerchantSubscriptionResumeReq struct for UnibeeApiMerchantSubscriptionResumeReq
type UnibeeApiMerchantSubscriptionResumeReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionResumeReq UnibeeApiMerchantSubscriptionResumeReq

// NewUnibeeApiMerchantSubscriptionResumeReq instantiates a new UnibeeApiMerchantSubscriptionResumeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionResumeReq(subscriptionId string) *UnibeeApiMerchantSubscriptionResumeReq {
	this := UnibeeApiMerchantSubscriptionResumeReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionResumeReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionResumeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionResumeReqWithDefaults() *UnibeeApiMerchantSubscriptionResumeReq {
	this := UnibeeApiMerchantSubscriptionResumeReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionResumeReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionResumeReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionResumeReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionResumeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionResumeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionResumeReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantSubscriptionResumeReq := _UnibeeApiMerchantSubscriptionResumeReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionResumeReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionResumeReq(varUnibeeApiMerchantSubscriptionResumeReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionResumeReq struct {
	value *UnibeeApiMerchantSubscriptionResumeReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionResumeReq) Get() *UnibeeApiMerchantSubscriptionResumeReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionResumeReq) Set(val *UnibeeApiMerchantSubscriptionResumeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionResumeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionResumeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionResumeReq(val *UnibeeApiMerchantSubscriptionResumeReq) *NullableUnibeeApiMerchantSubscriptionResumeReq {
	return &NullableUnibeeApiMerchantSubscriptionResumeReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionResumeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionResumeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


