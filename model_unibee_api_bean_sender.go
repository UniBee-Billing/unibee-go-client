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

// checks if the UnibeeApiBeanSender type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanSender{}

// UnibeeApiBeanSender struct for UnibeeApiBeanSender
type UnibeeApiBeanSender struct {
	// address
	Address *string `json:"address,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
}

// NewUnibeeApiBeanSender instantiates a new UnibeeApiBeanSender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanSender() *UnibeeApiBeanSender {
	this := UnibeeApiBeanSender{}
	return &this
}

// NewUnibeeApiBeanSenderWithDefaults instantiates a new UnibeeApiBeanSender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanSenderWithDefaults() *UnibeeApiBeanSender {
	this := UnibeeApiBeanSender{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnibeeApiBeanSender) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSender) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnibeeApiBeanSender) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UnibeeApiBeanSender) SetAddress(v string) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanSender) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSender) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanSender) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanSender) SetName(v string) {
	o.Name = &v
}

func (o UnibeeApiBeanSender) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanSender) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanSender struct {
	value *UnibeeApiBeanSender
	isSet bool
}

func (v NullableUnibeeApiBeanSender) Get() *UnibeeApiBeanSender {
	return v.value
}

func (v *NullableUnibeeApiBeanSender) Set(val *UnibeeApiBeanSender) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanSender) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanSender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanSender(val *UnibeeApiBeanSender) *NullableUnibeeApiBeanSender {
	return &NullableUnibeeApiBeanSender{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanSender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanSender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


