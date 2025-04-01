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

// checks if the UnibeeApiBeanDetailCreditTransactionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailCreditTransactionDetail{}

// UnibeeApiBeanDetailCreditTransactionDetail struct for UnibeeApiBeanDetailCreditTransactionDetail
type UnibeeApiBeanDetailCreditTransactionDetail struct {
	// type of credit account, 1-main recharge account, 2-promo credit account
	AccountType *int32 `json:"accountType,omitempty"`
	AdminMember *UnibeeApiBeanMerchantMember `json:"adminMember,omitempty"`
	// business id
	BizId *string `json:"bizId,omitempty"`
	By *string `json:"by,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	CreditAccount *UnibeeApiBeanCreditAccount `json:"creditAccount,omitempty"`
	// the credit amount after transaction,cent
	CreditAmountAfter *int64 `json:"creditAmountAfter,omitempty"`
	// the credit amount before transaction,cent
	CreditAmountBefore *int64 `json:"creditAmountBefore,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// delta amount,cent
	DeltaAmount *int64 `json:"deltaAmount,omitempty"`
	// delta currency amount, in cent
	DeltaCurrencyAmount *int64 `json:"deltaCurrencyAmount,omitempty"`
	// recharge transaction description
	Description *string `json:"description,omitempty"`
	// ExchangeRate for transaction, keep two decimal places，multiply by 100 saved, 1 currency = 1 credit * (exchange_rate/100), main account fixed rate to 100
	ExchangeRate *int64 `json:"exchangeRate,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// invoice_id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// recharge transaction title
	Name *string `json:"name,omitempty"`
	// unique id for timeline
	TransactionId *string `json:"transactionId,omitempty"`
	// transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out
	TransactionType *int32 `json:"transactionType,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailCreditTransactionDetail instantiates a new UnibeeApiBeanDetailCreditTransactionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailCreditTransactionDetail() *UnibeeApiBeanDetailCreditTransactionDetail {
	this := UnibeeApiBeanDetailCreditTransactionDetail{}
	return &this
}

// NewUnibeeApiBeanDetailCreditTransactionDetailWithDefaults instantiates a new UnibeeApiBeanDetailCreditTransactionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailCreditTransactionDetailWithDefaults() *UnibeeApiBeanDetailCreditTransactionDetail {
	this := UnibeeApiBeanDetailCreditTransactionDetail{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAccountType() int32 {
	if o == nil || IsNil(o.AccountType) {
		var ret int32
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAccountTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given int32 and assigns it to the AccountType field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetAccountType(v int32) {
	o.AccountType = &v
}

// GetAdminMember returns the AdminMember field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAdminMember() UnibeeApiBeanMerchantMember {
	if o == nil || IsNil(o.AdminMember) {
		var ret UnibeeApiBeanMerchantMember
		return ret
	}
	return *o.AdminMember
}

// GetAdminMemberOk returns a tuple with the AdminMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAdminMemberOk() (*UnibeeApiBeanMerchantMember, bool) {
	if o == nil || IsNil(o.AdminMember) {
		return nil, false
	}
	return o.AdminMember, true
}

// HasAdminMember returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasAdminMember() bool {
	if o != nil && !IsNil(o.AdminMember) {
		return true
	}

	return false
}

// SetAdminMember gets a reference to the given UnibeeApiBeanMerchantMember and assigns it to the AdminMember field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetAdminMember(v UnibeeApiBeanMerchantMember) {
	o.AdminMember = &v
}

// GetBizId returns the BizId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBizId() string {
	if o == nil || IsNil(o.BizId) {
		var ret string
		return ret
	}
	return *o.BizId
}

// GetBizIdOk returns a tuple with the BizId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBizIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizId) {
		return nil, false
	}
	return o.BizId, true
}

// HasBizId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasBizId() bool {
	if o != nil && !IsNil(o.BizId) {
		return true
	}

	return false
}

// SetBizId gets a reference to the given string and assigns it to the BizId field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetBizId(v string) {
	o.BizId = &v
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBy() string {
	if o == nil || IsNil(o.By) {
		var ret string
		return ret
	}
	return *o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetByOk() (*string, bool) {
	if o == nil || IsNil(o.By) {
		return nil, false
	}
	return o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasBy() bool {
	if o != nil && !IsNil(o.By) {
		return true
	}

	return false
}

// SetBy gets a reference to the given string and assigns it to the By field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetBy(v string) {
	o.By = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCreditAccount returns the CreditAccount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAccount() UnibeeApiBeanCreditAccount {
	if o == nil || IsNil(o.CreditAccount) {
		var ret UnibeeApiBeanCreditAccount
		return ret
	}
	return *o.CreditAccount
}

// GetCreditAccountOk returns a tuple with the CreditAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool) {
	if o == nil || IsNil(o.CreditAccount) {
		return nil, false
	}
	return o.CreditAccount, true
}

// HasCreditAccount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAccount() bool {
	if o != nil && !IsNil(o.CreditAccount) {
		return true
	}

	return false
}

// SetCreditAccount gets a reference to the given UnibeeApiBeanCreditAccount and assigns it to the CreditAccount field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAccount(v UnibeeApiBeanCreditAccount) {
	o.CreditAccount = &v
}

// GetCreditAmountAfter returns the CreditAmountAfter field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountAfter() int64 {
	if o == nil || IsNil(o.CreditAmountAfter) {
		var ret int64
		return ret
	}
	return *o.CreditAmountAfter
}

// GetCreditAmountAfterOk returns a tuple with the CreditAmountAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountAfterOk() (*int64, bool) {
	if o == nil || IsNil(o.CreditAmountAfter) {
		return nil, false
	}
	return o.CreditAmountAfter, true
}

// HasCreditAmountAfter returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAmountAfter() bool {
	if o != nil && !IsNil(o.CreditAmountAfter) {
		return true
	}

	return false
}

// SetCreditAmountAfter gets a reference to the given int64 and assigns it to the CreditAmountAfter field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAmountAfter(v int64) {
	o.CreditAmountAfter = &v
}

// GetCreditAmountBefore returns the CreditAmountBefore field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountBefore() int64 {
	if o == nil || IsNil(o.CreditAmountBefore) {
		var ret int64
		return ret
	}
	return *o.CreditAmountBefore
}

// GetCreditAmountBeforeOk returns a tuple with the CreditAmountBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountBeforeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreditAmountBefore) {
		return nil, false
	}
	return o.CreditAmountBefore, true
}

// HasCreditAmountBefore returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAmountBefore() bool {
	if o != nil && !IsNil(o.CreditAmountBefore) {
		return true
	}

	return false
}

// SetCreditAmountBefore gets a reference to the given int64 and assigns it to the CreditAmountBefore field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAmountBefore(v int64) {
	o.CreditAmountBefore = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetDeltaAmount returns the DeltaAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaAmount() int64 {
	if o == nil || IsNil(o.DeltaAmount) {
		var ret int64
		return ret
	}
	return *o.DeltaAmount
}

// GetDeltaAmountOk returns a tuple with the DeltaAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DeltaAmount) {
		return nil, false
	}
	return o.DeltaAmount, true
}

// HasDeltaAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDeltaAmount() bool {
	if o != nil && !IsNil(o.DeltaAmount) {
		return true
	}

	return false
}

// SetDeltaAmount gets a reference to the given int64 and assigns it to the DeltaAmount field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDeltaAmount(v int64) {
	o.DeltaAmount = &v
}

// GetDeltaCurrencyAmount returns the DeltaCurrencyAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaCurrencyAmount() int64 {
	if o == nil || IsNil(o.DeltaCurrencyAmount) {
		var ret int64
		return ret
	}
	return *o.DeltaCurrencyAmount
}

// GetDeltaCurrencyAmountOk returns a tuple with the DeltaCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaCurrencyAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DeltaCurrencyAmount) {
		return nil, false
	}
	return o.DeltaCurrencyAmount, true
}

// HasDeltaCurrencyAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDeltaCurrencyAmount() bool {
	if o != nil && !IsNil(o.DeltaCurrencyAmount) {
		return true
	}

	return false
}

// SetDeltaCurrencyAmount gets a reference to the given int64 and assigns it to the DeltaCurrencyAmount field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDeltaCurrencyAmount(v int64) {
	o.DeltaCurrencyAmount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDescription(v string) {
	o.Description = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetExchangeRate() int64 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret int64
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetExchangeRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given int64 and assigns it to the ExchangeRate field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetExchangeRate(v int64) {
	o.ExchangeRate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetId(v int64) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetName(v string) {
	o.Name = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionType() int32 {
	if o == nil || IsNil(o.TransactionType) {
		var ret int32
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given int32 and assigns it to the TransactionType field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetTransactionType(v int32) {
	o.TransactionType = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiBeanDetailCreditTransactionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailCreditTransactionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.AdminMember) {
		toSerialize["adminMember"] = o.AdminMember
	}
	if !IsNil(o.BizId) {
		toSerialize["bizId"] = o.BizId
	}
	if !IsNil(o.By) {
		toSerialize["by"] = o.By
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.CreditAccount) {
		toSerialize["creditAccount"] = o.CreditAccount
	}
	if !IsNil(o.CreditAmountAfter) {
		toSerialize["creditAmountAfter"] = o.CreditAmountAfter
	}
	if !IsNil(o.CreditAmountBefore) {
		toSerialize["creditAmountBefore"] = o.CreditAmountBefore
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DeltaAmount) {
		toSerialize["deltaAmount"] = o.DeltaAmount
	}
	if !IsNil(o.DeltaCurrencyAmount) {
		toSerialize["deltaCurrencyAmount"] = o.DeltaCurrencyAmount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailCreditTransactionDetail struct {
	value *UnibeeApiBeanDetailCreditTransactionDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailCreditTransactionDetail) Get() *UnibeeApiBeanDetailCreditTransactionDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailCreditTransactionDetail) Set(val *UnibeeApiBeanDetailCreditTransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailCreditTransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailCreditTransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailCreditTransactionDetail(val *UnibeeApiBeanDetailCreditTransactionDetail) *NullableUnibeeApiBeanDetailCreditTransactionDetail {
	return &NullableUnibeeApiBeanDetailCreditTransactionDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailCreditTransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailCreditTransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


