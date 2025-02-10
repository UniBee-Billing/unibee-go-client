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

// checks if the UnibeeApiSystemInvoiceBulkChannelSyncReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemInvoiceBulkChannelSyncReq{}

// UnibeeApiSystemInvoiceBulkChannelSyncReq struct for UnibeeApiSystemInvoiceBulkChannelSyncReq
type UnibeeApiSystemInvoiceBulkChannelSyncReq struct {
	// merchantId
	MerchantId string `json:"merchantId"`
}

type _UnibeeApiSystemInvoiceBulkChannelSyncReq UnibeeApiSystemInvoiceBulkChannelSyncReq

// NewUnibeeApiSystemInvoiceBulkChannelSyncReq instantiates a new UnibeeApiSystemInvoiceBulkChannelSyncReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemInvoiceBulkChannelSyncReq(merchantId string) *UnibeeApiSystemInvoiceBulkChannelSyncReq {
	this := UnibeeApiSystemInvoiceBulkChannelSyncReq{}
	this.MerchantId = merchantId
	return &this
}

// NewUnibeeApiSystemInvoiceBulkChannelSyncReqWithDefaults instantiates a new UnibeeApiSystemInvoiceBulkChannelSyncReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemInvoiceBulkChannelSyncReqWithDefaults() *UnibeeApiSystemInvoiceBulkChannelSyncReq {
	this := UnibeeApiSystemInvoiceBulkChannelSyncReq{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *UnibeeApiSystemInvoiceBulkChannelSyncReq) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceBulkChannelSyncReq) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *UnibeeApiSystemInvoiceBulkChannelSyncReq) SetMerchantId(v string) {
	o.MerchantId = v
}

func (o UnibeeApiSystemInvoiceBulkChannelSyncReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemInvoiceBulkChannelSyncReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	return toSerialize, nil
}

func (o *UnibeeApiSystemInvoiceBulkChannelSyncReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantId",
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

	varUnibeeApiSystemInvoiceBulkChannelSyncReq := _UnibeeApiSystemInvoiceBulkChannelSyncReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemInvoiceBulkChannelSyncReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemInvoiceBulkChannelSyncReq(varUnibeeApiSystemInvoiceBulkChannelSyncReq)

	return err
}

type NullableUnibeeApiSystemInvoiceBulkChannelSyncReq struct {
	value *UnibeeApiSystemInvoiceBulkChannelSyncReq
	isSet bool
}

func (v NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) Get() *UnibeeApiSystemInvoiceBulkChannelSyncReq {
	return v.value
}

func (v *NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) Set(val *UnibeeApiSystemInvoiceBulkChannelSyncReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemInvoiceBulkChannelSyncReq(val *UnibeeApiSystemInvoiceBulkChannelSyncReq) *NullableUnibeeApiSystemInvoiceBulkChannelSyncReq {
	return &NullableUnibeeApiSystemInvoiceBulkChannelSyncReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemInvoiceBulkChannelSyncReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


