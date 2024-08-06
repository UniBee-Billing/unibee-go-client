/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailGatewaySetupReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailGatewaySetupReq{}

// UnibeeApiMerchantEmailGatewaySetupReq struct for UnibeeApiMerchantEmailGatewaySetupReq
type UnibeeApiMerchantEmailGatewaySetupReq struct {
	// Whether setup the gateway as default or not, default is true
	IsDefault *bool `json:"IsDefault,omitempty"`
	// The setup data of email gateway
	Data string `json:"data"`
	// The name of email gateway, 'sendgrid' or other for future updates
	GatewayName string `json:"gatewayName"`
}

type _UnibeeApiMerchantEmailGatewaySetupReq UnibeeApiMerchantEmailGatewaySetupReq

// NewUnibeeApiMerchantEmailGatewaySetupReq instantiates a new UnibeeApiMerchantEmailGatewaySetupReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailGatewaySetupReq(data string, gatewayName string) *UnibeeApiMerchantEmailGatewaySetupReq {
	this := UnibeeApiMerchantEmailGatewaySetupReq{}
	var isDefault bool = true
	this.IsDefault = &isDefault
	this.Data = data
	this.GatewayName = gatewayName
	return &this
}

// NewUnibeeApiMerchantEmailGatewaySetupReqWithDefaults instantiates a new UnibeeApiMerchantEmailGatewaySetupReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailGatewaySetupReqWithDefaults() *UnibeeApiMerchantEmailGatewaySetupReq {
	this := UnibeeApiMerchantEmailGatewaySetupReq{}
	var isDefault bool = true
	this.IsDefault = &isDefault
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetData returns the Data field value
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetData(v string) {
	o.Data = v
}

// GetGatewayName returns the GatewayName field value
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetGatewayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetGatewayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayName, true
}

// SetGatewayName sets field value
func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetGatewayName(v string) {
	o.GatewayName = v
}

func (o UnibeeApiMerchantEmailGatewaySetupReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailGatewaySetupReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	toSerialize["data"] = o.Data
	toSerialize["gatewayName"] = o.GatewayName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailGatewaySetupReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"gatewayName",
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

	varUnibeeApiMerchantEmailGatewaySetupReq := _UnibeeApiMerchantEmailGatewaySetupReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailGatewaySetupReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailGatewaySetupReq(varUnibeeApiMerchantEmailGatewaySetupReq)

	return err
}

type NullableUnibeeApiMerchantEmailGatewaySetupReq struct {
	value *UnibeeApiMerchantEmailGatewaySetupReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupReq) Get() *UnibeeApiMerchantEmailGatewaySetupReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupReq) Set(val *UnibeeApiMerchantEmailGatewaySetupReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailGatewaySetupReq(val *UnibeeApiMerchantEmailGatewaySetupReq) *NullableUnibeeApiMerchantEmailGatewaySetupReq {
	return &NullableUnibeeApiMerchantEmailGatewaySetupReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailGatewaySetupReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailGatewaySetupReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


