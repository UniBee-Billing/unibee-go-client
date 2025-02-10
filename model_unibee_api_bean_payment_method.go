/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPaymentMethod{}

// UnibeeApiBeanPaymentMethod struct for UnibeeApiBeanPaymentMethod
type UnibeeApiBeanPaymentMethod struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Id *string `json:"id,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewUnibeeApiBeanPaymentMethod instantiates a new UnibeeApiBeanPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPaymentMethod() *UnibeeApiBeanPaymentMethod {
	this := UnibeeApiBeanPaymentMethod{}
	return &this
}

// NewUnibeeApiBeanPaymentMethodWithDefaults instantiates a new UnibeeApiBeanPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPaymentMethodWithDefaults() *UnibeeApiBeanPaymentMethod {
	this := UnibeeApiBeanPaymentMethod{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentMethod) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentMethod) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentMethod) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *UnibeeApiBeanPaymentMethod) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentMethod) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UnibeeApiBeanPaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentMethod) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentMethod) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentMethod) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *UnibeeApiBeanPaymentMethod) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentMethod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentMethod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnibeeApiBeanPaymentMethod) SetType(v string) {
	o.Type = &v
}

func (o UnibeeApiBeanPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPaymentMethod struct {
	value *UnibeeApiBeanPaymentMethod
	isSet bool
}

func (v NullableUnibeeApiBeanPaymentMethod) Get() *UnibeeApiBeanPaymentMethod {
	return v.value
}

func (v *NullableUnibeeApiBeanPaymentMethod) Set(val *UnibeeApiBeanPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPaymentMethod(val *UnibeeApiBeanPaymentMethod) *NullableUnibeeApiBeanPaymentMethod {
	return &NullableUnibeeApiBeanPaymentMethod{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


