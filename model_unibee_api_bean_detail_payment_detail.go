/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailPaymentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailPaymentDetail{}

// UnibeeApiBeanDetailPaymentDetail struct for UnibeeApiBeanDetailPaymentDetail
type UnibeeApiBeanDetailPaymentDetail struct {
	Gateway *UnibeeApiBeanGateway `json:"gateway,omitempty"`
	Invoice *UnibeeApiBeanDetailInvoiceDetail `json:"invoice,omitempty"`
	Payment *UnibeeApiBeanPayment `json:"payment,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailPaymentDetail instantiates a new UnibeeApiBeanDetailPaymentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailPaymentDetail() *UnibeeApiBeanDetailPaymentDetail {
	this := UnibeeApiBeanDetailPaymentDetail{}
	return &this
}

// NewUnibeeApiBeanDetailPaymentDetailWithDefaults instantiates a new UnibeeApiBeanDetailPaymentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailPaymentDetailWithDefaults() *UnibeeApiBeanDetailPaymentDetail {
	this := UnibeeApiBeanDetailPaymentDetail{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentDetail) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *UnibeeApiBeanDetailPaymentDetail) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentDetail) GetInvoice() UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanDetailInvoiceDetail and assigns it to the Invoice field.
func (o *UnibeeApiBeanDetailPaymentDetail) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail) {
	o.Invoice = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentDetail) GetPayment() UnibeeApiBeanPayment {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPayment and assigns it to the Payment field.
func (o *UnibeeApiBeanDetailPaymentDetail) SetPayment(v UnibeeApiBeanPayment) {
	o.Payment = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentDetail) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiBeanDetailPaymentDetail) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiBeanDetailPaymentDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailPaymentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailPaymentDetail struct {
	value *UnibeeApiBeanDetailPaymentDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailPaymentDetail) Get() *UnibeeApiBeanDetailPaymentDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailPaymentDetail) Set(val *UnibeeApiBeanDetailPaymentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailPaymentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailPaymentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailPaymentDetail(val *UnibeeApiBeanDetailPaymentDetail) *NullableUnibeeApiBeanDetailPaymentDetail {
	return &NullableUnibeeApiBeanDetailPaymentDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailPaymentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailPaymentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


