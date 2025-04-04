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

// checks if the UnibeeApiMerchantPaymentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentItem{}

// UnibeeApiMerchantPaymentItem struct for UnibeeApiMerchantPaymentItem
type UnibeeApiMerchantPaymentItem struct {
	// item total amount, sum(item.amount) should equal to totalAmount, cent
	Amount int64 `json:"amount"`
	// amountExcludingTax = unitAmountExcludingTax * quantity
	AmountExcludingTax *int64 `json:"amountExcludingTax,omitempty"`
	// The item currency of payment
	Currency *string `json:"currency,omitempty"`
	// The item description of payment
	Description *string `json:"description,omitempty"`
	// The item name of payment
	Name *string `json:"name,omitempty"`
	// The item quantity of payment
	Quantity *int64 `json:"quantity,omitempty"`
	// tax = amount - amountExcludingTax
	Tax *int64 `json:"tax,omitempty"`
	// The tax percentage of payment，1000 = 10%
	TaxPercentage *int64 `json:"taxPercentage,omitempty"`
	UnitAmountExcludingTax *int64 `json:"unitAmountExcludingTax,omitempty"`
}

type _UnibeeApiMerchantPaymentItem UnibeeApiMerchantPaymentItem

// NewUnibeeApiMerchantPaymentItem instantiates a new UnibeeApiMerchantPaymentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentItem(amount int64) *UnibeeApiMerchantPaymentItem {
	this := UnibeeApiMerchantPaymentItem{}
	this.Amount = amount
	return &this
}

// NewUnibeeApiMerchantPaymentItemWithDefaults instantiates a new UnibeeApiMerchantPaymentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentItemWithDefaults() *UnibeeApiMerchantPaymentItem {
	this := UnibeeApiMerchantPaymentItem{}
	return &this
}

// GetAmount returns the Amount field value
func (o *UnibeeApiMerchantPaymentItem) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UnibeeApiMerchantPaymentItem) SetAmount(v int64) {
	o.Amount = v
}

// GetAmountExcludingTax returns the AmountExcludingTax field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetAmountExcludingTax() int64 {
	if o == nil || IsNil(o.AmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.AmountExcludingTax
}

// GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.AmountExcludingTax) {
		return nil, false
	}
	return o.AmountExcludingTax, true
}

// HasAmountExcludingTax returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasAmountExcludingTax() bool {
	if o != nil && !IsNil(o.AmountExcludingTax) {
		return true
	}

	return false
}

// SetAmountExcludingTax gets a reference to the given int64 and assigns it to the AmountExcludingTax field.
func (o *UnibeeApiMerchantPaymentItem) SetAmountExcludingTax(v int64) {
	o.AmountExcludingTax = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPaymentItem) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantPaymentItem) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantPaymentItem) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantPaymentItem) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetTax() int64 {
	if o == nil || IsNil(o.Tax) {
		var ret int64
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given int64 and assigns it to the Tax field.
func (o *UnibeeApiMerchantPaymentItem) SetTax(v int64) {
	o.Tax = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetTaxPercentage() int64 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantPaymentItem) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

// GetUnitAmountExcludingTax returns the UnitAmountExcludingTax field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItem) GetUnitAmountExcludingTax() int64 {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.UnitAmountExcludingTax
}

// GetUnitAmountExcludingTaxOk returns a tuple with the UnitAmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItem) GetUnitAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		return nil, false
	}
	return o.UnitAmountExcludingTax, true
}

// HasUnitAmountExcludingTax returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItem) HasUnitAmountExcludingTax() bool {
	if o != nil && !IsNil(o.UnitAmountExcludingTax) {
		return true
	}

	return false
}

// SetUnitAmountExcludingTax gets a reference to the given int64 and assigns it to the UnitAmountExcludingTax field.
func (o *UnibeeApiMerchantPaymentItem) SetUnitAmountExcludingTax(v int64) {
	o.UnitAmountExcludingTax = &v
}

func (o UnibeeApiMerchantPaymentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AmountExcludingTax) {
		toSerialize["amountExcludingTax"] = o.AmountExcludingTax
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.UnitAmountExcludingTax) {
		toSerialize["unitAmountExcludingTax"] = o.UnitAmountExcludingTax
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varUnibeeApiMerchantPaymentItem := _UnibeeApiMerchantPaymentItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentItem)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentItem(varUnibeeApiMerchantPaymentItem)

	return err
}

type NullableUnibeeApiMerchantPaymentItem struct {
	value *UnibeeApiMerchantPaymentItem
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentItem) Get() *UnibeeApiMerchantPaymentItem {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentItem) Set(val *UnibeeApiMerchantPaymentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentItem(val *UnibeeApiMerchantPaymentItem) *NullableUnibeeApiMerchantPaymentItem {
	return &NullableUnibeeApiMerchantPaymentItem{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


