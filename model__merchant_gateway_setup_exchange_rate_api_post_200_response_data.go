/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantGatewaySetupExchangeRateApiPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGatewaySetupExchangeRateApiPost200ResponseData{}

// MerchantGatewaySetupExchangeRateApiPost200ResponseData struct for MerchantGatewaySetupExchangeRateApiPost200ResponseData
type MerchantGatewaySetupExchangeRateApiPost200ResponseData struct {
	// The hide star key of exchange rate api
	Data *string `json:"data,omitempty"`
}

// NewMerchantGatewaySetupExchangeRateApiPost200ResponseData instantiates a new MerchantGatewaySetupExchangeRateApiPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGatewaySetupExchangeRateApiPost200ResponseData() *MerchantGatewaySetupExchangeRateApiPost200ResponseData {
	this := MerchantGatewaySetupExchangeRateApiPost200ResponseData{}
	return &this
}

// NewMerchantGatewaySetupExchangeRateApiPost200ResponseDataWithDefaults instantiates a new MerchantGatewaySetupExchangeRateApiPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGatewaySetupExchangeRateApiPost200ResponseDataWithDefaults() *MerchantGatewaySetupExchangeRateApiPost200ResponseData {
	this := MerchantGatewaySetupExchangeRateApiPost200ResponseData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MerchantGatewaySetupExchangeRateApiPost200ResponseData) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGatewaySetupExchangeRateApiPost200ResponseData) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MerchantGatewaySetupExchangeRateApiPost200ResponseData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *MerchantGatewaySetupExchangeRateApiPost200ResponseData) SetData(v string) {
	o.Data = &v
}

func (o MerchantGatewaySetupExchangeRateApiPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGatewaySetupExchangeRateApiPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData struct {
	value *MerchantGatewaySetupExchangeRateApiPost200ResponseData
	isSet bool
}

func (v NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) Get() *MerchantGatewaySetupExchangeRateApiPost200ResponseData {
	return v.value
}

func (v *NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) Set(val *MerchantGatewaySetupExchangeRateApiPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGatewaySetupExchangeRateApiPost200ResponseData(val *MerchantGatewaySetupExchangeRateApiPost200ResponseData) *NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData {
	return &NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGatewaySetupExchangeRateApiPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


