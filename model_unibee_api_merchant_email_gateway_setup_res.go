/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantEmailGatewaySetupRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailGatewaySetupRes{}

// UnibeeApiMerchantEmailGatewaySetupRes struct for UnibeeApiMerchantEmailGatewaySetupRes
type UnibeeApiMerchantEmailGatewaySetupRes struct {
	// The hide star data
	Data *string `json:"data,omitempty"`
}

// NewUnibeeApiMerchantEmailGatewaySetupRes instantiates a new UnibeeApiMerchantEmailGatewaySetupRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailGatewaySetupRes() *UnibeeApiMerchantEmailGatewaySetupRes {
	this := UnibeeApiMerchantEmailGatewaySetupRes{}
	return &this
}

// NewUnibeeApiMerchantEmailGatewaySetupResWithDefaults instantiates a new UnibeeApiMerchantEmailGatewaySetupRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailGatewaySetupResWithDefaults() *UnibeeApiMerchantEmailGatewaySetupRes {
	this := UnibeeApiMerchantEmailGatewaySetupRes{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailGatewaySetupRes) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupRes) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupRes) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *UnibeeApiMerchantEmailGatewaySetupRes) SetData(v string) {
	o.Data = &v
}

func (o UnibeeApiMerchantEmailGatewaySetupRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailGatewaySetupRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantEmailGatewaySetupRes struct {
	value *UnibeeApiMerchantEmailGatewaySetupRes
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupRes) Get() *UnibeeApiMerchantEmailGatewaySetupRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupRes) Set(val *UnibeeApiMerchantEmailGatewaySetupRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailGatewaySetupRes(val *UnibeeApiMerchantEmailGatewaySetupRes) *NullableUnibeeApiMerchantEmailGatewaySetupRes {
	return &NullableUnibeeApiMerchantEmailGatewaySetupRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


