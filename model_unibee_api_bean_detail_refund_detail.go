/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailRefundDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailRefundDetail{}

// UnibeeApiBeanDetailRefundDetail struct for UnibeeApiBeanDetailRefundDetail
type UnibeeApiBeanDetailRefundDetail struct {
	Payment *UnibeeApiBeanPaymentSimplify `json:"payment,omitempty"`
	Refund *UnibeeApiBeanRefundSimplify `json:"refund,omitempty"`
	User *UnibeeApiBeanUserAccountSimplify `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailRefundDetail instantiates a new UnibeeApiBeanDetailRefundDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailRefundDetail() *UnibeeApiBeanDetailRefundDetail {
	this := UnibeeApiBeanDetailRefundDetail{}
	return &this
}

// NewUnibeeApiBeanDetailRefundDetailWithDefaults instantiates a new UnibeeApiBeanDetailRefundDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailRefundDetailWithDefaults() *UnibeeApiBeanDetailRefundDetail {
	this := UnibeeApiBeanDetailRefundDetail{}
	return &this
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailRefundDetail) GetPayment() UnibeeApiBeanPaymentSimplify {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPaymentSimplify
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailRefundDetail) GetPaymentOk() (*UnibeeApiBeanPaymentSimplify, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailRefundDetail) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPaymentSimplify and assigns it to the Payment field.
func (o *UnibeeApiBeanDetailRefundDetail) SetPayment(v UnibeeApiBeanPaymentSimplify) {
	o.Payment = &v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailRefundDetail) GetRefund() UnibeeApiBeanRefundSimplify {
	if o == nil || IsNil(o.Refund) {
		var ret UnibeeApiBeanRefundSimplify
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailRefundDetail) GetRefundOk() (*UnibeeApiBeanRefundSimplify, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailRefundDetail) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given UnibeeApiBeanRefundSimplify and assigns it to the Refund field.
func (o *UnibeeApiBeanDetailRefundDetail) SetRefund(v UnibeeApiBeanRefundSimplify) {
	o.Refund = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailRefundDetail) GetUser() UnibeeApiBeanUserAccountSimplify {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccountSimplify
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailRefundDetail) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailRefundDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccountSimplify and assigns it to the User field.
func (o *UnibeeApiBeanDetailRefundDetail) SetUser(v UnibeeApiBeanUserAccountSimplify) {
	o.User = &v
}

func (o UnibeeApiBeanDetailRefundDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailRefundDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailRefundDetail struct {
	value *UnibeeApiBeanDetailRefundDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailRefundDetail) Get() *UnibeeApiBeanDetailRefundDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailRefundDetail) Set(val *UnibeeApiBeanDetailRefundDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailRefundDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailRefundDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailRefundDetail(val *UnibeeApiBeanDetailRefundDetail) *NullableUnibeeApiBeanDetailRefundDetail {
	return &NullableUnibeeApiBeanDetailRefundDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailRefundDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailRefundDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


