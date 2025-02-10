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

// checks if the UnibeeApiBeanCreditConfigStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCreditConfigStatistics{}

// UnibeeApiBeanCreditConfigStatistics struct for UnibeeApiBeanCreditConfigStatistics
type UnibeeApiBeanCreditConfigStatistics struct {
	// Id
	Id *int64 `json:"id,omitempty"`
	// the total decrement amount
	TotalDecrementAmount *int64 `json:"totalDecrementAmount,omitempty"`
	// the total increment amount
	TotalIncrementAmount *int64 `json:"totalIncrementAmount,omitempty"`
}

// NewUnibeeApiBeanCreditConfigStatistics instantiates a new UnibeeApiBeanCreditConfigStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCreditConfigStatistics() *UnibeeApiBeanCreditConfigStatistics {
	this := UnibeeApiBeanCreditConfigStatistics{}
	return &this
}

// NewUnibeeApiBeanCreditConfigStatisticsWithDefaults instantiates a new UnibeeApiBeanCreditConfigStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCreditConfigStatisticsWithDefaults() *UnibeeApiBeanCreditConfigStatistics {
	this := UnibeeApiBeanCreditConfigStatistics{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfigStatistics) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanCreditConfigStatistics) SetId(v int64) {
	o.Id = &v
}

// GetTotalDecrementAmount returns the TotalDecrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfigStatistics) GetTotalDecrementAmount() int64 {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalDecrementAmount
}

// GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) GetTotalDecrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		return nil, false
	}
	return o.TotalDecrementAmount, true
}

// HasTotalDecrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) HasTotalDecrementAmount() bool {
	if o != nil && !IsNil(o.TotalDecrementAmount) {
		return true
	}

	return false
}

// SetTotalDecrementAmount gets a reference to the given int64 and assigns it to the TotalDecrementAmount field.
func (o *UnibeeApiBeanCreditConfigStatistics) SetTotalDecrementAmount(v int64) {
	o.TotalDecrementAmount = &v
}

// GetTotalIncrementAmount returns the TotalIncrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfigStatistics) GetTotalIncrementAmount() int64 {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalIncrementAmount
}

// GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) GetTotalIncrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		return nil, false
	}
	return o.TotalIncrementAmount, true
}

// HasTotalIncrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfigStatistics) HasTotalIncrementAmount() bool {
	if o != nil && !IsNil(o.TotalIncrementAmount) {
		return true
	}

	return false
}

// SetTotalIncrementAmount gets a reference to the given int64 and assigns it to the TotalIncrementAmount field.
func (o *UnibeeApiBeanCreditConfigStatistics) SetTotalIncrementAmount(v int64) {
	o.TotalIncrementAmount = &v
}

func (o UnibeeApiBeanCreditConfigStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCreditConfigStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TotalDecrementAmount) {
		toSerialize["totalDecrementAmount"] = o.TotalDecrementAmount
	}
	if !IsNil(o.TotalIncrementAmount) {
		toSerialize["totalIncrementAmount"] = o.TotalIncrementAmount
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCreditConfigStatistics struct {
	value *UnibeeApiBeanCreditConfigStatistics
	isSet bool
}

func (v NullableUnibeeApiBeanCreditConfigStatistics) Get() *UnibeeApiBeanCreditConfigStatistics {
	return v.value
}

func (v *NullableUnibeeApiBeanCreditConfigStatistics) Set(val *UnibeeApiBeanCreditConfigStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCreditConfigStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCreditConfigStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCreditConfigStatistics(val *UnibeeApiBeanCreditConfigStatistics) *NullableUnibeeApiBeanCreditConfigStatistics {
	return &NullableUnibeeApiBeanCreditConfigStatistics{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCreditConfigStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCreditConfigStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


