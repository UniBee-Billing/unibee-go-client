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

// checks if the UnibeeApiBeanInvoiceItemSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanInvoiceItemSimplify{}

// UnibeeApiBeanInvoiceItemSimplify struct for UnibeeApiBeanInvoiceItemSimplify
type UnibeeApiBeanInvoiceItemSimplify struct {
	Amount *int64 `json:"amount,omitempty"`
	AmountExcludingTax *int64 `json:"amountExcludingTax,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	Name *string `json:"name,omitempty"`
	OriginAmount *int64 `json:"originAmount,omitempty"`
	OriginUnitAmountExcludeTax *int64 `json:"originUnitAmountExcludeTax,omitempty"`
	PdfDescription *string `json:"pdfDescription,omitempty"`
	PeriodEnd *int64 `json:"periodEnd,omitempty"`
	PeriodStart *int64 `json:"periodStart,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Proration *bool `json:"proration,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	Tax *int64 `json:"tax,omitempty"`
	// Tax Percentage，1000 = 10%
	TaxPercentage *int64 `json:"taxPercentage,omitempty"`
	UnitAmountExcludingTax *int64 `json:"unitAmountExcludingTax,omitempty"`
}

// NewUnibeeApiBeanInvoiceItemSimplify instantiates a new UnibeeApiBeanInvoiceItemSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanInvoiceItemSimplify() *UnibeeApiBeanInvoiceItemSimplify {
	this := UnibeeApiBeanInvoiceItemSimplify{}
	return &this
}

// NewUnibeeApiBeanInvoiceItemSimplifyWithDefaults instantiates a new UnibeeApiBeanInvoiceItemSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanInvoiceItemSimplifyWithDefaults() *UnibeeApiBeanInvoiceItemSimplify {
	this := UnibeeApiBeanInvoiceItemSimplify{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetAmount(v int64) {
	o.Amount = &v
}

// GetAmountExcludingTax returns the AmountExcludingTax field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountExcludingTax() int64 {
	if o == nil || IsNil(o.AmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.AmountExcludingTax
}

// GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.AmountExcludingTax) {
		return nil, false
	}
	return o.AmountExcludingTax, true
}

// HasAmountExcludingTax returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasAmountExcludingTax() bool {
	if o != nil && !IsNil(o.AmountExcludingTax) {
		return true
	}

	return false
}

// SetAmountExcludingTax gets a reference to the given int64 and assigns it to the AmountExcludingTax field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetAmountExcludingTax(v int64) {
	o.AmountExcludingTax = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetName(v string) {
	o.Name = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginAmount() int64 {
	if o == nil || IsNil(o.OriginAmount) {
		var ret int64
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given int64 and assigns it to the OriginAmount field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetOriginAmount(v int64) {
	o.OriginAmount = &v
}

// GetOriginUnitAmountExcludeTax returns the OriginUnitAmountExcludeTax field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginUnitAmountExcludeTax() int64 {
	if o == nil || IsNil(o.OriginUnitAmountExcludeTax) {
		var ret int64
		return ret
	}
	return *o.OriginUnitAmountExcludeTax
}

// GetOriginUnitAmountExcludeTaxOk returns a tuple with the OriginUnitAmountExcludeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginUnitAmountExcludeTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginUnitAmountExcludeTax) {
		return nil, false
	}
	return o.OriginUnitAmountExcludeTax, true
}

// HasOriginUnitAmountExcludeTax returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasOriginUnitAmountExcludeTax() bool {
	if o != nil && !IsNil(o.OriginUnitAmountExcludeTax) {
		return true
	}

	return false
}

// SetOriginUnitAmountExcludeTax gets a reference to the given int64 and assigns it to the OriginUnitAmountExcludeTax field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetOriginUnitAmountExcludeTax(v int64) {
	o.OriginUnitAmountExcludeTax = &v
}

// GetPdfDescription returns the PdfDescription field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPdfDescription() string {
	if o == nil || IsNil(o.PdfDescription) {
		var ret string
		return ret
	}
	return *o.PdfDescription
}

// GetPdfDescriptionOk returns a tuple with the PdfDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPdfDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PdfDescription) {
		return nil, false
	}
	return o.PdfDescription, true
}

// HasPdfDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasPdfDescription() bool {
	if o != nil && !IsNil(o.PdfDescription) {
		return true
	}

	return false
}

// SetPdfDescription gets a reference to the given string and assigns it to the PdfDescription field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetPdfDescription(v string) {
	o.PdfDescription = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodEnd() int64 {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret int64
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodEndOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given int64 and assigns it to the PeriodEnd field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetPeriodEnd(v int64) {
	o.PeriodEnd = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodStart() int64 {
	if o == nil || IsNil(o.PeriodStart) {
		var ret int64
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodStartOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given int64 and assigns it to the PeriodStart field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetPeriodStart(v int64) {
	o.PeriodStart = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetProration returns the Proration field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetProration() bool {
	if o == nil || IsNil(o.Proration) {
		var ret bool
		return ret
	}
	return *o.Proration
}

// GetProrationOk returns a tuple with the Proration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetProrationOk() (*bool, bool) {
	if o == nil || IsNil(o.Proration) {
		return nil, false
	}
	return o.Proration, true
}

// HasProration returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasProration() bool {
	if o != nil && !IsNil(o.Proration) {
		return true
	}

	return false
}

// SetProration gets a reference to the given bool and assigns it to the Proration field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetProration(v bool) {
	o.Proration = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetTax() int64 {
	if o == nil || IsNil(o.Tax) {
		var ret int64
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given int64 and assigns it to the Tax field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetTax(v int64) {
	o.Tax = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxPercentage() int64 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

// GetUnitAmountExcludingTax returns the UnitAmountExcludingTax field value if set, zero value otherwise.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetUnitAmountExcludingTax() int64 {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.UnitAmountExcludingTax
}

// GetUnitAmountExcludingTaxOk returns a tuple with the UnitAmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) GetUnitAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		return nil, false
	}
	return o.UnitAmountExcludingTax, true
}

// HasUnitAmountExcludingTax returns a boolean if a field has been set.
func (o *UnibeeApiBeanInvoiceItemSimplify) HasUnitAmountExcludingTax() bool {
	if o != nil && !IsNil(o.UnitAmountExcludingTax) {
		return true
	}

	return false
}

// SetUnitAmountExcludingTax gets a reference to the given int64 and assigns it to the UnitAmountExcludingTax field.
func (o *UnibeeApiBeanInvoiceItemSimplify) SetUnitAmountExcludingTax(v int64) {
	o.UnitAmountExcludingTax = &v
}

func (o UnibeeApiBeanInvoiceItemSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanInvoiceItemSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountExcludingTax) {
		toSerialize["amountExcludingTax"] = o.AmountExcludingTax
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OriginAmount) {
		toSerialize["originAmount"] = o.OriginAmount
	}
	if !IsNil(o.OriginUnitAmountExcludeTax) {
		toSerialize["originUnitAmountExcludeTax"] = o.OriginUnitAmountExcludeTax
	}
	if !IsNil(o.PdfDescription) {
		toSerialize["pdfDescription"] = o.PdfDescription
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["periodEnd"] = o.PeriodEnd
	}
	if !IsNil(o.PeriodStart) {
		toSerialize["periodStart"] = o.PeriodStart
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Proration) {
		toSerialize["proration"] = o.Proration
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

type NullableUnibeeApiBeanInvoiceItemSimplify struct {
	value *UnibeeApiBeanInvoiceItemSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanInvoiceItemSimplify) Get() *UnibeeApiBeanInvoiceItemSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanInvoiceItemSimplify) Set(val *UnibeeApiBeanInvoiceItemSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanInvoiceItemSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanInvoiceItemSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanInvoiceItemSimplify(val *UnibeeApiBeanInvoiceItemSimplify) *NullableUnibeeApiBeanInvoiceItemSimplify {
	return &NullableUnibeeApiBeanInvoiceItemSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanInvoiceItemSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanInvoiceItemSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


