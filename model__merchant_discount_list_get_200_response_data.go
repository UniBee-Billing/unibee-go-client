/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantDiscountListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantDiscountListGet200ResponseData{}

// MerchantDiscountListGet200ResponseData struct for MerchantDiscountListGet200ResponseData
type MerchantDiscountListGet200ResponseData struct {
	// Discount Object List
	Discounts []UnibeeApiBeanMerchantDiscountCode `json:"discounts,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantDiscountListGet200ResponseData instantiates a new MerchantDiscountListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDiscountListGet200ResponseData() *MerchantDiscountListGet200ResponseData {
	this := MerchantDiscountListGet200ResponseData{}
	return &this
}

// NewMerchantDiscountListGet200ResponseDataWithDefaults instantiates a new MerchantDiscountListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDiscountListGet200ResponseDataWithDefaults() *MerchantDiscountListGet200ResponseData {
	this := MerchantDiscountListGet200ResponseData{}
	return &this
}

// GetDiscounts returns the Discounts field value if set, zero value otherwise.
func (o *MerchantDiscountListGet200ResponseData) GetDiscounts() []UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discounts) {
		var ret []UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountListGet200ResponseData) GetDiscountsOk() ([]UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discounts) {
		return nil, false
	}
	return o.Discounts, true
}

// HasDiscounts returns a boolean if a field has been set.
func (o *MerchantDiscountListGet200ResponseData) HasDiscounts() bool {
	if o != nil && !IsNil(o.Discounts) {
		return true
	}

	return false
}

// SetDiscounts gets a reference to the given []UnibeeApiBeanMerchantDiscountCode and assigns it to the Discounts field.
func (o *MerchantDiscountListGet200ResponseData) SetDiscounts(v []UnibeeApiBeanMerchantDiscountCode) {
	o.Discounts = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantDiscountListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantDiscountListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantDiscountListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantDiscountListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantDiscountListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discounts) {
		toSerialize["discounts"] = o.Discounts
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantDiscountListGet200ResponseData struct {
	value *MerchantDiscountListGet200ResponseData
	isSet bool
}

func (v NullableMerchantDiscountListGet200ResponseData) Get() *MerchantDiscountListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantDiscountListGet200ResponseData) Set(val *MerchantDiscountListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDiscountListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDiscountListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDiscountListGet200ResponseData(val *MerchantDiscountListGet200ResponseData) *NullableMerchantDiscountListGet200ResponseData {
	return &NullableMerchantDiscountListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantDiscountListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDiscountListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


