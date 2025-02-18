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

// checks if the MerchantCreditGetPromoConfigStatisticsGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditGetPromoConfigStatisticsGet200ResponseData{}

// MerchantCreditGetPromoConfigStatisticsGet200ResponseData struct for MerchantCreditGetPromoConfigStatisticsGet200ResponseData
type MerchantCreditGetPromoConfigStatisticsGet200ResponseData struct {
	CreditConfigStatistics *UnibeeApiBeanCreditConfigStatistics `json:"creditConfigStatistics,omitempty"`
}

// NewMerchantCreditGetPromoConfigStatisticsGet200ResponseData instantiates a new MerchantCreditGetPromoConfigStatisticsGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditGetPromoConfigStatisticsGet200ResponseData() *MerchantCreditGetPromoConfigStatisticsGet200ResponseData {
	this := MerchantCreditGetPromoConfigStatisticsGet200ResponseData{}
	return &this
}

// NewMerchantCreditGetPromoConfigStatisticsGet200ResponseDataWithDefaults instantiates a new MerchantCreditGetPromoConfigStatisticsGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditGetPromoConfigStatisticsGet200ResponseDataWithDefaults() *MerchantCreditGetPromoConfigStatisticsGet200ResponseData {
	this := MerchantCreditGetPromoConfigStatisticsGet200ResponseData{}
	return &this
}

// GetCreditConfigStatistics returns the CreditConfigStatistics field value if set, zero value otherwise.
func (o *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) GetCreditConfigStatistics() UnibeeApiBeanCreditConfigStatistics {
	if o == nil || IsNil(o.CreditConfigStatistics) {
		var ret UnibeeApiBeanCreditConfigStatistics
		return ret
	}
	return *o.CreditConfigStatistics
}

// GetCreditConfigStatisticsOk returns a tuple with the CreditConfigStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) GetCreditConfigStatisticsOk() (*UnibeeApiBeanCreditConfigStatistics, bool) {
	if o == nil || IsNil(o.CreditConfigStatistics) {
		return nil, false
	}
	return o.CreditConfigStatistics, true
}

// HasCreditConfigStatistics returns a boolean if a field has been set.
func (o *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) HasCreditConfigStatistics() bool {
	if o != nil && !IsNil(o.CreditConfigStatistics) {
		return true
	}

	return false
}

// SetCreditConfigStatistics gets a reference to the given UnibeeApiBeanCreditConfigStatistics and assigns it to the CreditConfigStatistics field.
func (o *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) SetCreditConfigStatistics(v UnibeeApiBeanCreditConfigStatistics) {
	o.CreditConfigStatistics = &v
}

func (o MerchantCreditGetPromoConfigStatisticsGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditGetPromoConfigStatisticsGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditConfigStatistics) {
		toSerialize["creditConfigStatistics"] = o.CreditConfigStatistics
	}
	return toSerialize, nil
}

type NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData struct {
	value *MerchantCreditGetPromoConfigStatisticsGet200ResponseData
	isSet bool
}

func (v NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) Get() *MerchantCreditGetPromoConfigStatisticsGet200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) Set(val *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData(val *MerchantCreditGetPromoConfigStatisticsGet200ResponseData) *NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData {
	return &NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditGetPromoConfigStatisticsGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


