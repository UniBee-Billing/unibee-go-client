/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantCreditPromoCreditDecrementReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditPromoCreditDecrementReq{}

// UnibeeApiMerchantCreditPromoCreditDecrementReq Decrease user promo credit amount, the amount after decreased should greater than 0
type UnibeeApiMerchantCreditPromoCreditDecrementReq struct {
	// The Amount to decrease, should greater than 0
	Amount int64 `json:"amount"`
	// currency of recharge
	Currency string `json:"currency"`
	// description of increase action
	Description *string `json:"description,omitempty"`
	// name of increase action
	Name *string `json:"name,omitempty"`
	// filter id of user
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantCreditPromoCreditDecrementReq UnibeeApiMerchantCreditPromoCreditDecrementReq

// NewUnibeeApiMerchantCreditPromoCreditDecrementReq instantiates a new UnibeeApiMerchantCreditPromoCreditDecrementReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditPromoCreditDecrementReq(amount int64, currency string, userId int64) *UnibeeApiMerchantCreditPromoCreditDecrementReq {
	this := UnibeeApiMerchantCreditPromoCreditDecrementReq{}
	this.Amount = amount
	this.Currency = currency
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantCreditPromoCreditDecrementReqWithDefaults instantiates a new UnibeeApiMerchantCreditPromoCreditDecrementReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditPromoCreditDecrementReqWithDefaults() *UnibeeApiMerchantCreditPromoCreditDecrementReq {
	this := UnibeeApiMerchantCreditPromoCreditDecrementReq{}
	return &this
}

// GetAmount returns the Amount field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetAmount(v int64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantCreditPromoCreditDecrementReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditPromoCreditDecrementReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"userId",
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

	varUnibeeApiMerchantCreditPromoCreditDecrementReq := _UnibeeApiMerchantCreditPromoCreditDecrementReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantCreditPromoCreditDecrementReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantCreditPromoCreditDecrementReq(varUnibeeApiMerchantCreditPromoCreditDecrementReq)

	return err
}

type NullableUnibeeApiMerchantCreditPromoCreditDecrementReq struct {
	value *UnibeeApiMerchantCreditPromoCreditDecrementReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) Get() *UnibeeApiMerchantCreditPromoCreditDecrementReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) Set(val *UnibeeApiMerchantCreditPromoCreditDecrementReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditPromoCreditDecrementReq(val *UnibeeApiMerchantCreditPromoCreditDecrementReq) *NullableUnibeeApiMerchantCreditPromoCreditDecrementReq {
	return &NullableUnibeeApiMerchantCreditPromoCreditDecrementReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditPromoCreditDecrementReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


