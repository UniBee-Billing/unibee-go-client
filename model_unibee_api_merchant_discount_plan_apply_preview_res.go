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

// checks if the UnibeeApiMerchantDiscountPlanApplyPreviewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountPlanApplyPreviewRes{}

// UnibeeApiMerchantDiscountPlanApplyPreviewRes struct for UnibeeApiMerchantDiscountPlanApplyPreviewRes
type UnibeeApiMerchantDiscountPlanApplyPreviewRes struct {
	// The discount amount can apply to plan
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	DiscountCode *UnibeeApiBeanMerchantDiscountCode `json:"discountCode,omitempty"`
	// The apply preview failure reason
	FailureReason *string `json:"failureReason,omitempty"`
	// The apply preview result, true or false
	Valid *bool `json:"valid,omitempty"`
}

// NewUnibeeApiMerchantDiscountPlanApplyPreviewRes instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountPlanApplyPreviewRes() *UnibeeApiMerchantDiscountPlanApplyPreviewRes {
	this := UnibeeApiMerchantDiscountPlanApplyPreviewRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountPlanApplyPreviewResWithDefaults instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountPlanApplyPreviewResWithDefaults() *UnibeeApiMerchantDiscountPlanApplyPreviewRes {
	this := UnibeeApiMerchantDiscountPlanApplyPreviewRes{}
	return &this
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.DiscountCode) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode) {
	o.DiscountCode = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetValid(v bool) {
	o.Valid = &v
}

func (o UnibeeApiMerchantDiscountPlanApplyPreviewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountPlanApplyPreviewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes struct {
	value *UnibeeApiMerchantDiscountPlanApplyPreviewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) Get() *UnibeeApiMerchantDiscountPlanApplyPreviewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) Set(val *UnibeeApiMerchantDiscountPlanApplyPreviewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountPlanApplyPreviewRes(val *UnibeeApiMerchantDiscountPlanApplyPreviewRes) *NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes {
	return &NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


