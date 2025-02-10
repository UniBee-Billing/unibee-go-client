/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentItemListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentItemListGet200ResponseData{}

// MerchantPaymentItemListGet200ResponseData struct for MerchantPaymentItemListGet200ResponseData
type MerchantPaymentItemListGet200ResponseData struct {
	// Payment Item Object List
	PaymentItems []UnibeeApiBeanPaymentItem `json:"paymentItems,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantPaymentItemListGet200ResponseData instantiates a new MerchantPaymentItemListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentItemListGet200ResponseData() *MerchantPaymentItemListGet200ResponseData {
	this := MerchantPaymentItemListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentItemListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentItemListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentItemListGet200ResponseDataWithDefaults() *MerchantPaymentItemListGet200ResponseData {
	this := MerchantPaymentItemListGet200ResponseData{}
	return &this
}

// GetPaymentItems returns the PaymentItems field value if set, zero value otherwise.
func (o *MerchantPaymentItemListGet200ResponseData) GetPaymentItems() []UnibeeApiBeanPaymentItem {
	if o == nil || IsNil(o.PaymentItems) {
		var ret []UnibeeApiBeanPaymentItem
		return ret
	}
	return o.PaymentItems
}

// GetPaymentItemsOk returns a tuple with the PaymentItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentItemListGet200ResponseData) GetPaymentItemsOk() ([]UnibeeApiBeanPaymentItem, bool) {
	if o == nil || IsNil(o.PaymentItems) {
		return nil, false
	}
	return o.PaymentItems, true
}

// HasPaymentItems returns a boolean if a field has been set.
func (o *MerchantPaymentItemListGet200ResponseData) HasPaymentItems() bool {
	if o != nil && !IsNil(o.PaymentItems) {
		return true
	}

	return false
}

// SetPaymentItems gets a reference to the given []UnibeeApiBeanPaymentItem and assigns it to the PaymentItems field.
func (o *MerchantPaymentItemListGet200ResponseData) SetPaymentItems(v []UnibeeApiBeanPaymentItem) {
	o.PaymentItems = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantPaymentItemListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentItemListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantPaymentItemListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantPaymentItemListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantPaymentItemListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentItemListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentItems) {
		toSerialize["paymentItems"] = o.PaymentItems
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantPaymentItemListGet200ResponseData struct {
	value *MerchantPaymentItemListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentItemListGet200ResponseData) Get() *MerchantPaymentItemListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentItemListGet200ResponseData) Set(val *MerchantPaymentItemListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentItemListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentItemListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentItemListGet200ResponseData(val *MerchantPaymentItemListGet200ResponseData) *NullableMerchantPaymentItemListGet200ResponseData {
	return &NullableMerchantPaymentItemListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentItemListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentItemListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


