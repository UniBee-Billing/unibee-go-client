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

// checks if the MerchantGatewayEditSortPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGatewayEditSortPost200ResponseData{}

// MerchantGatewayEditSortPost200ResponseData struct for MerchantGatewayEditSortPost200ResponseData
type MerchantGatewayEditSortPost200ResponseData struct {
	// Payment Gateway Setup Object List
	Gateways []UnibeeApiBeanDetailGateway `json:"gateways,omitempty"`
}

// NewMerchantGatewayEditSortPost200ResponseData instantiates a new MerchantGatewayEditSortPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGatewayEditSortPost200ResponseData() *MerchantGatewayEditSortPost200ResponseData {
	this := MerchantGatewayEditSortPost200ResponseData{}
	return &this
}

// NewMerchantGatewayEditSortPost200ResponseDataWithDefaults instantiates a new MerchantGatewayEditSortPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGatewayEditSortPost200ResponseDataWithDefaults() *MerchantGatewayEditSortPost200ResponseData {
	this := MerchantGatewayEditSortPost200ResponseData{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *MerchantGatewayEditSortPost200ResponseData) GetGateways() []UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateways) {
		var ret []UnibeeApiBeanDetailGateway
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGatewayEditSortPost200ResponseData) GetGatewaysOk() ([]UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *MerchantGatewayEditSortPost200ResponseData) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []UnibeeApiBeanDetailGateway and assigns it to the Gateways field.
func (o *MerchantGatewayEditSortPost200ResponseData) SetGateways(v []UnibeeApiBeanDetailGateway) {
	o.Gateways = v
}

func (o MerchantGatewayEditSortPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGatewayEditSortPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableMerchantGatewayEditSortPost200ResponseData struct {
	value *MerchantGatewayEditSortPost200ResponseData
	isSet bool
}

func (v NullableMerchantGatewayEditSortPost200ResponseData) Get() *MerchantGatewayEditSortPost200ResponseData {
	return v.value
}

func (v *NullableMerchantGatewayEditSortPost200ResponseData) Set(val *MerchantGatewayEditSortPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGatewayEditSortPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGatewayEditSortPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGatewayEditSortPost200ResponseData(val *MerchantGatewayEditSortPost200ResponseData) *NullableMerchantGatewayEditSortPost200ResponseData {
	return &NullableMerchantGatewayEditSortPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGatewayEditSortPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGatewayEditSortPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


