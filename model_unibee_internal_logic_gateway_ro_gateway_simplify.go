/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoGatewaySimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoGatewaySimplify{}

// UnibeeInternalLogicGatewayRoGatewaySimplify struct for UnibeeInternalLogicGatewayRoGatewaySimplify
type UnibeeInternalLogicGatewayRoGatewaySimplify struct {
	GatewayId *int64 `json:"gatewayId,omitempty"`
	GatewayLogo *string `json:"gatewayLogo,omitempty"`
	GatewayName *string `json:"gatewayName,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoGatewaySimplify instantiates a new UnibeeInternalLogicGatewayRoGatewaySimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoGatewaySimplify() *UnibeeInternalLogicGatewayRoGatewaySimplify {
	this := UnibeeInternalLogicGatewayRoGatewaySimplify{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoGatewaySimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoGatewaySimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoGatewaySimplifyWithDefaults() *UnibeeInternalLogicGatewayRoGatewaySimplify {
	this := UnibeeInternalLogicGatewayRoGatewaySimplify{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayLogo returns the GatewayLogo field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayLogo() string {
	if o == nil || IsNil(o.GatewayLogo) {
		var ret string
		return ret
	}
	return *o.GatewayLogo
}

// GetGatewayLogoOk returns a tuple with the GatewayLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayLogoOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayLogo) {
		return nil, false
	}
	return o.GatewayLogo, true
}

// HasGatewayLogo returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) HasGatewayLogo() bool {
	if o != nil && !IsNil(o.GatewayLogo) {
		return true
	}

	return false
}

// SetGatewayLogo gets a reference to the given string and assigns it to the GatewayLogo field.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) SetGatewayLogo(v string) {
	o.GatewayLogo = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *UnibeeInternalLogicGatewayRoGatewaySimplify) SetGatewayName(v string) {
	o.GatewayName = &v
}

func (o UnibeeInternalLogicGatewayRoGatewaySimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoGatewaySimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayLogo) {
		toSerialize["gatewayLogo"] = o.GatewayLogo
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gatewayName"] = o.GatewayName
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoGatewaySimplify struct {
	value *UnibeeInternalLogicGatewayRoGatewaySimplify
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoGatewaySimplify) Get() *UnibeeInternalLogicGatewayRoGatewaySimplify {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewaySimplify) Set(val *UnibeeInternalLogicGatewayRoGatewaySimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoGatewaySimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewaySimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoGatewaySimplify(val *UnibeeInternalLogicGatewayRoGatewaySimplify) *NullableUnibeeInternalLogicGatewayRoGatewaySimplify {
	return &NullableUnibeeInternalLogicGatewayRoGatewaySimplify{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoGatewaySimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewaySimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


