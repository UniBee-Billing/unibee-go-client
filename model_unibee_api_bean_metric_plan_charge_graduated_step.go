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

// checks if the UnibeeApiBeanMetricPlanChargeGraduatedStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMetricPlanChargeGraduatedStep{}

// UnibeeApiBeanMetricPlanChargeGraduatedStep struct for UnibeeApiBeanMetricPlanChargeGraduatedStep
type UnibeeApiBeanMetricPlanChargeGraduatedStep struct {
	// EndValue
	EndValue *int64 `json:"endValue,omitempty"`
	// FlatAmount
	FlatAmount *float32 `json:"flatAmount,omitempty"`
	// PerAmount
	PerAmount *float32 `json:"perAmount,omitempty"`
	// StartValue
	StartValue *int64 `json:"startValue,omitempty"`
}

// NewUnibeeApiBeanMetricPlanChargeGraduatedStep instantiates a new UnibeeApiBeanMetricPlanChargeGraduatedStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMetricPlanChargeGraduatedStep() *UnibeeApiBeanMetricPlanChargeGraduatedStep {
	this := UnibeeApiBeanMetricPlanChargeGraduatedStep{}
	return &this
}

// NewUnibeeApiBeanMetricPlanChargeGraduatedStepWithDefaults instantiates a new UnibeeApiBeanMetricPlanChargeGraduatedStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMetricPlanChargeGraduatedStepWithDefaults() *UnibeeApiBeanMetricPlanChargeGraduatedStep {
	this := UnibeeApiBeanMetricPlanChargeGraduatedStep{}
	return &this
}

// GetEndValue returns the EndValue field value if set, zero value otherwise.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetEndValue() int64 {
	if o == nil || IsNil(o.EndValue) {
		var ret int64
		return ret
	}
	return *o.EndValue
}

// GetEndValueOk returns a tuple with the EndValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetEndValueOk() (*int64, bool) {
	if o == nil || IsNil(o.EndValue) {
		return nil, false
	}
	return o.EndValue, true
}

// HasEndValue returns a boolean if a field has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) HasEndValue() bool {
	if o != nil && !IsNil(o.EndValue) {
		return true
	}

	return false
}

// SetEndValue gets a reference to the given int64 and assigns it to the EndValue field.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) SetEndValue(v int64) {
	o.EndValue = &v
}

// GetFlatAmount returns the FlatAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetFlatAmount() float32 {
	if o == nil || IsNil(o.FlatAmount) {
		var ret float32
		return ret
	}
	return *o.FlatAmount
}

// GetFlatAmountOk returns a tuple with the FlatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetFlatAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.FlatAmount) {
		return nil, false
	}
	return o.FlatAmount, true
}

// HasFlatAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) HasFlatAmount() bool {
	if o != nil && !IsNil(o.FlatAmount) {
		return true
	}

	return false
}

// SetFlatAmount gets a reference to the given float32 and assigns it to the FlatAmount field.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) SetFlatAmount(v float32) {
	o.FlatAmount = &v
}

// GetPerAmount returns the PerAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetPerAmount() float32 {
	if o == nil || IsNil(o.PerAmount) {
		var ret float32
		return ret
	}
	return *o.PerAmount
}

// GetPerAmountOk returns a tuple with the PerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetPerAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.PerAmount) {
		return nil, false
	}
	return o.PerAmount, true
}

// HasPerAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) HasPerAmount() bool {
	if o != nil && !IsNil(o.PerAmount) {
		return true
	}

	return false
}

// SetPerAmount gets a reference to the given float32 and assigns it to the PerAmount field.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) SetPerAmount(v float32) {
	o.PerAmount = &v
}

// GetStartValue returns the StartValue field value if set, zero value otherwise.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetStartValue() int64 {
	if o == nil || IsNil(o.StartValue) {
		var ret int64
		return ret
	}
	return *o.StartValue
}

// GetStartValueOk returns a tuple with the StartValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) GetStartValueOk() (*int64, bool) {
	if o == nil || IsNil(o.StartValue) {
		return nil, false
	}
	return o.StartValue, true
}

// HasStartValue returns a boolean if a field has been set.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) HasStartValue() bool {
	if o != nil && !IsNil(o.StartValue) {
		return true
	}

	return false
}

// SetStartValue gets a reference to the given int64 and assigns it to the StartValue field.
func (o *UnibeeApiBeanMetricPlanChargeGraduatedStep) SetStartValue(v int64) {
	o.StartValue = &v
}

func (o UnibeeApiBeanMetricPlanChargeGraduatedStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMetricPlanChargeGraduatedStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndValue) {
		toSerialize["endValue"] = o.EndValue
	}
	if !IsNil(o.FlatAmount) {
		toSerialize["flatAmount"] = o.FlatAmount
	}
	if !IsNil(o.PerAmount) {
		toSerialize["perAmount"] = o.PerAmount
	}
	if !IsNil(o.StartValue) {
		toSerialize["startValue"] = o.StartValue
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMetricPlanChargeGraduatedStep struct {
	value *UnibeeApiBeanMetricPlanChargeGraduatedStep
	isSet bool
}

func (v NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) Get() *UnibeeApiBeanMetricPlanChargeGraduatedStep {
	return v.value
}

func (v *NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) Set(val *UnibeeApiBeanMetricPlanChargeGraduatedStep) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMetricPlanChargeGraduatedStep(val *UnibeeApiBeanMetricPlanChargeGraduatedStep) *NullableUnibeeApiBeanMetricPlanChargeGraduatedStep {
	return &NullableUnibeeApiBeanMetricPlanChargeGraduatedStep{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMetricPlanChargeGraduatedStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


