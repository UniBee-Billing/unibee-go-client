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

// checks if the UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq{}

// UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq struct for UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq
type UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq

// NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq instantiates a new UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq(subscriptionId string) *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq {
	this := UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReqWithDefaults() *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq {
	this := UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq := _UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq(varUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq struct {
	value *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) Get() *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) Set(val *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq(val *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) *NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq {
	return &NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


