/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantCreditCreditTransactionListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditCreditTransactionListReq{}

// UnibeeApiMerchantCreditCreditTransactionListReq Get Credit Transaction list
type UnibeeApiMerchantCreditCreditTransactionListReq struct {
	// filter type of account, 1-main account, 2-promo credit account
	AccountType int32 `json:"accountType"`
	// Count Of Per Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// filter currency of account
	Currency *string `json:"currency,omitempty"`
	// filter email of user
	Email *string `json:"email,omitempty"`
	// Page, Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort Field，gmt_create|gmt_modify，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out
	TransactionTypes []int32 `json:"transactionTypes,omitempty"`
	// filter id of user
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantCreditCreditTransactionListReq UnibeeApiMerchantCreditCreditTransactionListReq

// NewUnibeeApiMerchantCreditCreditTransactionListReq instantiates a new UnibeeApiMerchantCreditCreditTransactionListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditCreditTransactionListReq(accountType int32) *UnibeeApiMerchantCreditCreditTransactionListReq {
	this := UnibeeApiMerchantCreditCreditTransactionListReq{}
	this.AccountType = accountType
	return &this
}

// NewUnibeeApiMerchantCreditCreditTransactionListReqWithDefaults instantiates a new UnibeeApiMerchantCreditCreditTransactionListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditCreditTransactionListReqWithDefaults() *UnibeeApiMerchantCreditCreditTransactionListReq {
	this := UnibeeApiMerchantCreditCreditTransactionListReq{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetAccountType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetAccountTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetAccountType(v int32) {
	o.AccountType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetEmail(v string) {
	o.Email = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetTransactionTypes returns the TransactionTypes field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetTransactionTypes() []int32 {
	if o == nil || IsNil(o.TransactionTypes) {
		var ret []int32
		return ret
	}
	return o.TransactionTypes
}

// GetTransactionTypesOk returns a tuple with the TransactionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetTransactionTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.TransactionTypes) {
		return nil, false
	}
	return o.TransactionTypes, true
}

// HasTransactionTypes returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasTransactionTypes() bool {
	if o != nil && !IsNil(o.TransactionTypes) {
		return true
	}

	return false
}

// SetTransactionTypes gets a reference to the given []int32 and assigns it to the TransactionTypes field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetTransactionTypes(v []int32) {
	o.TransactionTypes = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantCreditCreditTransactionListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditCreditTransactionListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.CreateTimeEnd) {
		toSerialize["createTimeEnd"] = o.CreateTimeEnd
	}
	if !IsNil(o.CreateTimeStart) {
		toSerialize["createTimeStart"] = o.CreateTimeStart
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.SortField) {
		toSerialize["sortField"] = o.SortField
	}
	if !IsNil(o.SortType) {
		toSerialize["sortType"] = o.SortType
	}
	if !IsNil(o.TransactionTypes) {
		toSerialize["transactionTypes"] = o.TransactionTypes
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantCreditCreditTransactionListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantCreditCreditTransactionListReq := _UnibeeApiMerchantCreditCreditTransactionListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantCreditCreditTransactionListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantCreditCreditTransactionListReq(varUnibeeApiMerchantCreditCreditTransactionListReq)

	return err
}

type NullableUnibeeApiMerchantCreditCreditTransactionListReq struct {
	value *UnibeeApiMerchantCreditCreditTransactionListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListReq) Get() *UnibeeApiMerchantCreditCreditTransactionListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListReq) Set(val *UnibeeApiMerchantCreditCreditTransactionListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditCreditTransactionListReq(val *UnibeeApiMerchantCreditCreditTransactionListReq) *NullableUnibeeApiMerchantCreditCreditTransactionListReq {
	return &NullableUnibeeApiMerchantCreditCreditTransactionListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditCreditTransactionListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditCreditTransactionListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


