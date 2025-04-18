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

// checks if the UnibeeApiBeanDetailGatewaySort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailGatewaySort{}

// UnibeeApiBeanDetailGatewaySort struct for UnibeeApiBeanDetailGatewaySort
type UnibeeApiBeanDetailGatewaySort struct {
	// The gateway id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// Required, The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium
	GatewayName *string `json:"gatewayName,omitempty"`
	// Required, The sort value of payment gateway, should greater than 0, The bigger, the closer to the front
	Sort *int64 `json:"sort,omitempty"`
}

// NewUnibeeApiBeanDetailGatewaySort instantiates a new UnibeeApiBeanDetailGatewaySort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailGatewaySort() *UnibeeApiBeanDetailGatewaySort {
	this := UnibeeApiBeanDetailGatewaySort{}
	return &this
}

// NewUnibeeApiBeanDetailGatewaySortWithDefaults instantiates a new UnibeeApiBeanDetailGatewaySort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailGatewaySortWithDefaults() *UnibeeApiBeanDetailGatewaySort {
	this := UnibeeApiBeanDetailGatewaySort{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGatewaySort) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanDetailGatewaySort) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGatewaySort) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *UnibeeApiBeanDetailGatewaySort) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGatewaySort) GetSort() int64 {
	if o == nil || IsNil(o.Sort) {
		var ret int64
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGatewaySort) GetSortOk() (*int64, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGatewaySort) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int64 and assigns it to the Sort field.
func (o *UnibeeApiBeanDetailGatewaySort) SetSort(v int64) {
	o.Sort = &v
}

func (o UnibeeApiBeanDetailGatewaySort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailGatewaySort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gatewayName"] = o.GatewayName
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailGatewaySort struct {
	value *UnibeeApiBeanDetailGatewaySort
	isSet bool
}

func (v NullableUnibeeApiBeanDetailGatewaySort) Get() *UnibeeApiBeanDetailGatewaySort {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailGatewaySort) Set(val *UnibeeApiBeanDetailGatewaySort) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailGatewaySort) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailGatewaySort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailGatewaySort(val *UnibeeApiBeanDetailGatewaySort) *NullableUnibeeApiBeanDetailGatewaySort {
	return &NullableUnibeeApiBeanDetailGatewaySort{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailGatewaySort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailGatewaySort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


