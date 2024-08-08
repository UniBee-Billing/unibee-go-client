/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemInvoiceChannelSyncReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemInvoiceChannelSyncReq{}

// UnibeeApiSystemInvoiceChannelSyncReq struct for UnibeeApiSystemInvoiceChannelSyncReq
type UnibeeApiSystemInvoiceChannelSyncReq struct {
	// invoiceId
	InvoiceId string `json:"invoiceId"`
	// merchantId
	MerchantId string `json:"merchantId"`
}

type _UnibeeApiSystemInvoiceChannelSyncReq UnibeeApiSystemInvoiceChannelSyncReq

// NewUnibeeApiSystemInvoiceChannelSyncReq instantiates a new UnibeeApiSystemInvoiceChannelSyncReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemInvoiceChannelSyncReq(invoiceId string, merchantId string) *UnibeeApiSystemInvoiceChannelSyncReq {
	this := UnibeeApiSystemInvoiceChannelSyncReq{}
	this.InvoiceId = invoiceId
	this.MerchantId = merchantId
	return &this
}

// NewUnibeeApiSystemInvoiceChannelSyncReqWithDefaults instantiates a new UnibeeApiSystemInvoiceChannelSyncReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemInvoiceChannelSyncReqWithDefaults() *UnibeeApiSystemInvoiceChannelSyncReq {
	this := UnibeeApiSystemInvoiceChannelSyncReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiSystemInvoiceChannelSyncReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceChannelSyncReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiSystemInvoiceChannelSyncReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetMerchantId returns the MerchantId field value
func (o *UnibeeApiSystemInvoiceChannelSyncReq) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceChannelSyncReq) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *UnibeeApiSystemInvoiceChannelSyncReq) SetMerchantId(v string) {
	o.MerchantId = v
}

func (o UnibeeApiSystemInvoiceChannelSyncReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemInvoiceChannelSyncReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	toSerialize["merchantId"] = o.MerchantId
	return toSerialize, nil
}

func (o *UnibeeApiSystemInvoiceChannelSyncReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
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

	varUnibeeApiSystemInvoiceChannelSyncReq := _UnibeeApiSystemInvoiceChannelSyncReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemInvoiceChannelSyncReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemInvoiceChannelSyncReq(varUnibeeApiSystemInvoiceChannelSyncReq)

	return err
}

type NullableUnibeeApiSystemInvoiceChannelSyncReq struct {
	value *UnibeeApiSystemInvoiceChannelSyncReq
	isSet bool
}

func (v NullableUnibeeApiSystemInvoiceChannelSyncReq) Get() *UnibeeApiSystemInvoiceChannelSyncReq {
	return v.value
}

func (v *NullableUnibeeApiSystemInvoiceChannelSyncReq) Set(val *UnibeeApiSystemInvoiceChannelSyncReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemInvoiceChannelSyncReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemInvoiceChannelSyncReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemInvoiceChannelSyncReq(val *UnibeeApiSystemInvoiceChannelSyncReq) *NullableUnibeeApiSystemInvoiceChannelSyncReq {
	return &NullableUnibeeApiSystemInvoiceChannelSyncReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemInvoiceChannelSyncReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemInvoiceChannelSyncReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


