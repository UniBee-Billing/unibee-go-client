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

// checks if the UnibeeApiMerchantCreditNewCreditRechargeRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditNewCreditRechargeRes{}

// UnibeeApiMerchantCreditNewCreditRechargeRes struct for UnibeeApiMerchantCreditNewCreditRechargeRes
type UnibeeApiMerchantCreditNewCreditRechargeRes struct {
	CreditAccount *UnibeeApiBeanCreditAccount `json:"creditAccount,omitempty"`
	CreditRecharge *UnibeeApiBeanCreditRecharge `json:"creditRecharge,omitempty"`
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
	Link *string `json:"link,omitempty"`
	Merchant *UnibeeApiBeanMerchant `json:"merchant,omitempty"`
	// Paid，true|false
	Paid *bool `json:"paid,omitempty"`
	Payment *UnibeeApiBeanPayment `json:"payment,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiMerchantCreditNewCreditRechargeRes instantiates a new UnibeeApiMerchantCreditNewCreditRechargeRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditNewCreditRechargeRes() *UnibeeApiMerchantCreditNewCreditRechargeRes {
	this := UnibeeApiMerchantCreditNewCreditRechargeRes{}
	return &this
}

// NewUnibeeApiMerchantCreditNewCreditRechargeResWithDefaults instantiates a new UnibeeApiMerchantCreditNewCreditRechargeRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditNewCreditRechargeResWithDefaults() *UnibeeApiMerchantCreditNewCreditRechargeRes {
	this := UnibeeApiMerchantCreditNewCreditRechargeRes{}
	return &this
}

// GetCreditAccount returns the CreditAccount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.CreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.CreditAccount
}

// GetCreditAccountOk returns a tuple with the CreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.CreditAccount) {
		return nil, false
	}
	return o.CreditAccount, true
}

// HasCreditAccount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasCreditAccount() bool {
	if o != nil && !IsNil(o.CreditAccount) {
		return true
	}

	return false
}

// SetCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the CreditAccount field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.CreditAccount = &v
}

// GetCreditRecharge returns the CreditRecharge field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditRecharge() UnibeeApiBeanCreditRecharge {
	if o == nil || IsNil(o.CreditRecharge) {
		var ret UnibeeApiBeanCreditRecharge
		return ret
	}
	return *o.CreditRecharge
}

// GetCreditRechargeOk returns a tuple with the CreditRecharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditRechargeOk() (*UnibeeApiBeanCreditRecharge, bool) {
	if o == nil || IsNil(o.CreditRecharge) {
		return nil, false
	}
	return o.CreditRecharge, true
}

// HasCreditRecharge returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasCreditRecharge() bool {
	if o != nil && !IsNil(o.CreditRecharge) {
		return true
	}

	return false
}

// SetCreditRecharge gets a reference to the given UnibeeApiBeanCreditRecharge and assigns it to the CreditRecharge field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetCreditRecharge(v UnibeeApiBeanCreditRecharge) {
	o.CreditRecharge = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetLink(v string) {
	o.Link = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetPaid(v bool) {
	o.Paid = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPayment() UnibeeApiBeanPayment {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaymentOk() (*UnibeeApiBeanPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPayment and assigns it to the Payment field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetPayment(v UnibeeApiBeanPayment) {
	o.Payment = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiMerchantCreditNewCreditRechargeRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditNewCreditRechargeRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditAccount) {
		toSerialize["creditAccount"] = o.CreditAccount
	}
	if !IsNil(o.CreditRecharge) {
		toSerialize["creditRecharge"] = o.CreditRecharge
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditNewCreditRechargeRes struct {
	value *UnibeeApiMerchantCreditNewCreditRechargeRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeRes) Get() *UnibeeApiMerchantCreditNewCreditRechargeRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeRes) Set(val *UnibeeApiMerchantCreditNewCreditRechargeRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditNewCreditRechargeRes(val *UnibeeApiMerchantCreditNewCreditRechargeRes) *NullableUnibeeApiMerchantCreditNewCreditRechargeRes {
	return &NullableUnibeeApiMerchantCreditNewCreditRechargeRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


