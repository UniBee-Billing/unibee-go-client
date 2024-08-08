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

// checks if the UnibeeApiMerchantGatewayWireTransferSetupReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayWireTransferSetupReq{}

// UnibeeApiMerchantGatewayWireTransferSetupReq Setup the wire transfer
type UnibeeApiMerchantGatewayWireTransferSetupReq struct {
	Bank *UnibeeApiBeanGatewayBank `json:"bank,omitempty"`
	// The currency of wire transfer 
	Currency string `json:"currency"`
	// The minimum amount of wire transfer, cents
	MinimumAmount int64 `json:"minimumAmount"`
}

type _UnibeeApiMerchantGatewayWireTransferSetupReq UnibeeApiMerchantGatewayWireTransferSetupReq

// NewUnibeeApiMerchantGatewayWireTransferSetupReq instantiates a new UnibeeApiMerchantGatewayWireTransferSetupReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayWireTransferSetupReq(currency string, minimumAmount int64) *UnibeeApiMerchantGatewayWireTransferSetupReq {
	this := UnibeeApiMerchantGatewayWireTransferSetupReq{}
	this.Currency = currency
	this.MinimumAmount = minimumAmount
	return &this
}

// NewUnibeeApiMerchantGatewayWireTransferSetupReqWithDefaults instantiates a new UnibeeApiMerchantGatewayWireTransferSetupReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayWireTransferSetupReqWithDefaults() *UnibeeApiMerchantGatewayWireTransferSetupReq {
	this := UnibeeApiMerchantGatewayWireTransferSetupReq{}
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBank() UnibeeApiBeanGatewayBank {
	if o == nil || IsNil(o.Bank) {
		var ret UnibeeApiBeanGatewayBank
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBankOk() (*UnibeeApiBeanGatewayBank, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given UnibeeApiBeanGatewayBank and assigns it to the Bank field.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetBank(v UnibeeApiBeanGatewayBank) {
	o.Bank = &v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetCurrency(v string) {
	o.Currency = v
}

// GetMinimumAmount returns the MinimumAmount field value
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetMinimumAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetMinimumAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumAmount, true
}

// SetMinimumAmount sets field value
func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetMinimumAmount(v int64) {
	o.MinimumAmount = v
}

func (o UnibeeApiMerchantGatewayWireTransferSetupReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayWireTransferSetupReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	toSerialize["currency"] = o.Currency
	toSerialize["minimumAmount"] = o.MinimumAmount
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
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

	varUnibeeApiMerchantGatewayWireTransferSetupReq := _UnibeeApiMerchantGatewayWireTransferSetupReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewayWireTransferSetupReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewayWireTransferSetupReq(varUnibeeApiMerchantGatewayWireTransferSetupReq)

	return err
}

type NullableUnibeeApiMerchantGatewayWireTransferSetupReq struct {
	value *UnibeeApiMerchantGatewayWireTransferSetupReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayWireTransferSetupReq) Get() *UnibeeApiMerchantGatewayWireTransferSetupReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferSetupReq) Set(val *UnibeeApiMerchantGatewayWireTransferSetupReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayWireTransferSetupReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferSetupReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayWireTransferSetupReq(val *UnibeeApiMerchantGatewayWireTransferSetupReq) *NullableUnibeeApiMerchantGatewayWireTransferSetupReq {
	return &NullableUnibeeApiMerchantGatewayWireTransferSetupReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayWireTransferSetupReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferSetupReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


