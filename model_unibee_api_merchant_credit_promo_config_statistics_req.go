/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantCreditPromoConfigStatisticsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditPromoConfigStatisticsReq{}

// UnibeeApiMerchantCreditPromoConfigStatisticsReq struct for UnibeeApiMerchantCreditPromoConfigStatisticsReq
type UnibeeApiMerchantCreditPromoConfigStatisticsReq struct {
	// currency
	Currency *string `json:"currency,omitempty"`
}

// NewUnibeeApiMerchantCreditPromoConfigStatisticsReq instantiates a new UnibeeApiMerchantCreditPromoConfigStatisticsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditPromoConfigStatisticsReq() *UnibeeApiMerchantCreditPromoConfigStatisticsReq {
	this := UnibeeApiMerchantCreditPromoConfigStatisticsReq{}
	return &this
}

// NewUnibeeApiMerchantCreditPromoConfigStatisticsReqWithDefaults instantiates a new UnibeeApiMerchantCreditPromoConfigStatisticsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditPromoConfigStatisticsReqWithDefaults() *UnibeeApiMerchantCreditPromoConfigStatisticsReq {
	this := UnibeeApiMerchantCreditPromoConfigStatisticsReq{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditPromoConfigStatisticsReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoConfigStatisticsReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditPromoConfigStatisticsReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantCreditPromoConfigStatisticsReq) SetCurrency(v string) {
	o.Currency = &v
}

func (o UnibeeApiMerchantCreditPromoConfigStatisticsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditPromoConfigStatisticsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq struct {
	value *UnibeeApiMerchantCreditPromoConfigStatisticsReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) Get() *UnibeeApiMerchantCreditPromoConfigStatisticsReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) Set(val *UnibeeApiMerchantCreditPromoConfigStatisticsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditPromoConfigStatisticsReq(val *UnibeeApiMerchantCreditPromoConfigStatisticsReq) *NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq {
	return &NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigStatisticsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


