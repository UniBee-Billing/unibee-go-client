/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPlanProductParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPlanProductParam{}

// UnibeeApiBeanPlanProductParam struct for UnibeeApiBeanPlanProductParam
type UnibeeApiBeanPlanProductParam struct {
	// Description
	Description *string `json:"description,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
}

// NewUnibeeApiBeanPlanProductParam instantiates a new UnibeeApiBeanPlanProductParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPlanProductParam() *UnibeeApiBeanPlanProductParam {
	this := UnibeeApiBeanPlanProductParam{}
	return &this
}

// NewUnibeeApiBeanPlanProductParamWithDefaults instantiates a new UnibeeApiBeanPlanProductParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPlanProductParamWithDefaults() *UnibeeApiBeanPlanProductParam {
	this := UnibeeApiBeanPlanProductParam{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanProductParam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanProductParam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanProductParam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanPlanProductParam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlanProductParam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlanProductParam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlanProductParam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanPlanProductParam) SetName(v string) {
	o.Name = &v
}

func (o UnibeeApiBeanPlanProductParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPlanProductParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPlanProductParam struct {
	value *UnibeeApiBeanPlanProductParam
	isSet bool
}

func (v NullableUnibeeApiBeanPlanProductParam) Get() *UnibeeApiBeanPlanProductParam {
	return v.value
}

func (v *NullableUnibeeApiBeanPlanProductParam) Set(val *UnibeeApiBeanPlanProductParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPlanProductParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPlanProductParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPlanProductParam(val *UnibeeApiBeanPlanProductParam) *NullableUnibeeApiBeanPlanProductParam {
	return &NullableUnibeeApiBeanPlanProductParam{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPlanProductParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPlanProductParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


