/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantGatewayWireTransferEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayWireTransferEditReq{}

// UnibeeApiMerchantGatewayWireTransferEditReq Edit the wire transfer
type UnibeeApiMerchantGatewayWireTransferEditReq struct {
	Bank *UnibeeApiBeanGatewayBank `json:"bank,omitempty"`
	// The currency of wire transfer 
	Currency string `json:"currency"`
	// The id of payment gateway
	GatewayId int64 `json:"gatewayId"`
	// The minimum amount of wire transfer, cents
	MinimumAmount int64 `json:"minimumAmount"`
}

type _UnibeeApiMerchantGatewayWireTransferEditReq UnibeeApiMerchantGatewayWireTransferEditReq

// NewUnibeeApiMerchantGatewayWireTransferEditReq instantiates a new UnibeeApiMerchantGatewayWireTransferEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayWireTransferEditReq(currency string, gatewayId int64, minimumAmount int64) *UnibeeApiMerchantGatewayWireTransferEditReq {
	this := UnibeeApiMerchantGatewayWireTransferEditReq{}
	this.Currency = currency
	this.GatewayId = gatewayId
	this.MinimumAmount = minimumAmount
	return &this
}

// NewUnibeeApiMerchantGatewayWireTransferEditReqWithDefaults instantiates a new UnibeeApiMerchantGatewayWireTransferEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayWireTransferEditReqWithDefaults() *UnibeeApiMerchantGatewayWireTransferEditReq {
	this := UnibeeApiMerchantGatewayWireTransferEditReq{}
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetBank() UnibeeApiBeanGatewayBank {
	if o == nil || IsNil(o.Bank) {
		var ret UnibeeApiBeanGatewayBank
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetBankOk() (*UnibeeApiBeanGatewayBank, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given UnibeeApiBeanGatewayBank and assigns it to the Bank field.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetBank(v UnibeeApiBeanGatewayBank) {
	o.Bank = &v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetCurrency(v string) {
	o.Currency = v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetMinimumAmount returns the MinimumAmount field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetMinimumAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetMinimumAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumAmount, true
}

// SetMinimumAmount sets field value
func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetMinimumAmount(v int64) {
	o.MinimumAmount = v
}

func (o UnibeeApiMerchantGatewayWireTransferEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayWireTransferEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	toSerialize["currency"] = o.Currency
	toSerialize["gatewayId"] = o.GatewayId
	toSerialize["minimumAmount"] = o.MinimumAmount
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewayWireTransferEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"gatewayId",
		"minimumAmount",
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

	varUnibeeApiMerchantGatewayWireTransferEditReq := _UnibeeApiMerchantGatewayWireTransferEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewayWireTransferEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewayWireTransferEditReq(varUnibeeApiMerchantGatewayWireTransferEditReq)

	return err
}

type NullableUnibeeApiMerchantGatewayWireTransferEditReq struct {
	value *UnibeeApiMerchantGatewayWireTransferEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditReq) Get() *UnibeeApiMerchantGatewayWireTransferEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditReq) Set(val *UnibeeApiMerchantGatewayWireTransferEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayWireTransferEditReq(val *UnibeeApiMerchantGatewayWireTransferEditReq) *NullableUnibeeApiMerchantGatewayWireTransferEditReq {
	return &NullableUnibeeApiMerchantGatewayWireTransferEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


