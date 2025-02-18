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

// checks if the MerchantCreditNewCreditRechargePost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCreditNewCreditRechargePost200ResponseData{}

// MerchantCreditNewCreditRechargePost200ResponseData struct for MerchantCreditNewCreditRechargePost200ResponseData
type MerchantCreditNewCreditRechargePost200ResponseData struct {
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

// NewMerchantCreditNewCreditRechargePost200ResponseData instantiates a new MerchantCreditNewCreditRechargePost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreditNewCreditRechargePost200ResponseData() *MerchantCreditNewCreditRechargePost200ResponseData {
	this := MerchantCreditNewCreditRechargePost200ResponseData{}
	return &this
}

// NewMerchantCreditNewCreditRechargePost200ResponseDataWithDefaults instantiates a new MerchantCreditNewCreditRechargePost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreditNewCreditRechargePost200ResponseDataWithDefaults() *MerchantCreditNewCreditRechargePost200ResponseData {
	this := MerchantCreditNewCreditRechargePost200ResponseData{}
	return &this
}

// GetCreditAccount returns the CreditAccount field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.CreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.CreditAccount
}

// GetCreditAccountOk returns a tuple with the CreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.CreditAccount) {
		return nil, false
	}
	return o.CreditAccount, true
}

// HasCreditAccount returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasCreditAccount() bool {
	if o != nil && !IsNil(o.CreditAccount) {
		return true
	}

	return false
}

// SetCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the CreditAccount field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.CreditAccount = &v
}

// GetCreditRecharge returns the CreditRecharge field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetCreditRecharge() UnibeeApiBeanCreditRecharge {
	if o == nil || IsNil(o.CreditRecharge) {
		var ret UnibeeApiBeanCreditRecharge
		return ret
	}
	return *o.CreditRecharge
}

// GetCreditRechargeOk returns a tuple with the CreditRecharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetCreditRechargeOk() (*UnibeeApiBeanCreditRecharge, bool) {
	if o == nil || IsNil(o.CreditRecharge) {
		return nil, false
	}
	return o.CreditRecharge, true
}

// HasCreditRecharge returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasCreditRecharge() bool {
	if o != nil && !IsNil(o.CreditRecharge) {
		return true
	}

	return false
}

// SetCreditRecharge gets a reference to the given UnibeeApiBeanCreditRecharge and assigns it to the CreditRecharge field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetCreditRecharge(v UnibeeApiBeanCreditRecharge) {
	o.CreditRecharge = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetLink(v string) {
	o.Link = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetPaid(v bool) {
	o.Paid = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetPayment() UnibeeApiBeanPayment {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetPaymentOk() (*UnibeeApiBeanPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPayment and assigns it to the Payment field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetPayment(v UnibeeApiBeanPayment) {
	o.Payment = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *MerchantCreditNewCreditRechargePost200ResponseData) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o MerchantCreditNewCreditRechargePost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreditNewCreditRechargePost200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantCreditNewCreditRechargePost200ResponseData struct {
	value *MerchantCreditNewCreditRechargePost200ResponseData
	isSet bool
}

func (v NullableMerchantCreditNewCreditRechargePost200ResponseData) Get() *MerchantCreditNewCreditRechargePost200ResponseData {
	return v.value
}

func (v *NullableMerchantCreditNewCreditRechargePost200ResponseData) Set(val *MerchantCreditNewCreditRechargePost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreditNewCreditRechargePost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreditNewCreditRechargePost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreditNewCreditRechargePost200ResponseData(val *MerchantCreditNewCreditRechargePost200ResponseData) *NullableMerchantCreditNewCreditRechargePost200ResponseData {
	return &NullableMerchantCreditNewCreditRechargePost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantCreditNewCreditRechargePost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreditNewCreditRechargePost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


