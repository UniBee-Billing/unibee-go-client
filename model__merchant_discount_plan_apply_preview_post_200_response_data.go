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

// checks if the MerchantDiscountPlanApplyPreviewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantDiscountPlanApplyPreviewPost200ResponseData{}

// MerchantDiscountPlanApplyPreviewPost200ResponseData struct for MerchantDiscountPlanApplyPreviewPost200ResponseData
type MerchantDiscountPlanApplyPreviewPost200ResponseData struct {
	// The discount amount can apply to plan
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	DiscountCode *UnibeeApiBeanMerchantDiscountCode `json:"discountCode,omitempty"`
	// The apply preview failure reason
	FailureReason *string `json:"failureReason,omitempty"`
	// The apply preview result, true or false
	Valid *bool `json:"valid,omitempty"`
}

// NewMerchantDiscountPlanApplyPreviewPost200ResponseData instantiates a new MerchantDiscountPlanApplyPreviewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDiscountPlanApplyPreviewPost200ResponseData() *MerchantDiscountPlanApplyPreviewPost200ResponseData {
	this := MerchantDiscountPlanApplyPreviewPost200ResponseData{}
	return &this
}

// NewMerchantDiscountPlanApplyPreviewPost200ResponseDataWithDefaults instantiates a new MerchantDiscountPlanApplyPreviewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDiscountPlanApplyPreviewPost200ResponseDataWithDefaults() *MerchantDiscountPlanApplyPreviewPost200ResponseData {
	this := MerchantDiscountPlanApplyPreviewPost200ResponseData{}
	return &this
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.DiscountCode) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the DiscountCode field.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode) {
	o.DiscountCode = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetValid(v bool) {
	o.Valid = &v
}

func (o MerchantDiscountPlanApplyPreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantDiscountPlanApplyPreviewPost200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantDiscountPlanApplyPreviewPost200ResponseData struct {
	value *MerchantDiscountPlanApplyPreviewPost200ResponseData
	isSet bool
}

func (v NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) Get() *MerchantDiscountPlanApplyPreviewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) Set(val *MerchantDiscountPlanApplyPreviewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDiscountPlanApplyPreviewPost200ResponseData(val *MerchantDiscountPlanApplyPreviewPost200ResponseData) *NullableMerchantDiscountPlanApplyPreviewPost200ResponseData {
	return &NullableMerchantDiscountPlanApplyPreviewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDiscountPlanApplyPreviewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


