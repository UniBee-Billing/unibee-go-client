/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceNewInvoiceItemParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceNewInvoiceItemParam{}

// UnibeeApiMerchantInvoiceNewInvoiceItemParam struct for UnibeeApiMerchantInvoiceNewInvoiceItemParam
type UnibeeApiMerchantInvoiceNewInvoiceItemParam struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	UnitAmountExcludingTax *int64 `json:"unitAmountExcludingTax,omitempty"`
}

// NewUnibeeApiMerchantInvoiceNewInvoiceItemParam instantiates a new UnibeeApiMerchantInvoiceNewInvoiceItemParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceNewInvoiceItemParam() *UnibeeApiMerchantInvoiceNewInvoiceItemParam {
	this := UnibeeApiMerchantInvoiceNewInvoiceItemParam{}
	return &this
}

// NewUnibeeApiMerchantInvoiceNewInvoiceItemParamWithDefaults instantiates a new UnibeeApiMerchantInvoiceNewInvoiceItemParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceNewInvoiceItemParamWithDefaults() *UnibeeApiMerchantInvoiceNewInvoiceItemParam {
	this := UnibeeApiMerchantInvoiceNewInvoiceItemParam{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetUnitAmountExcludingTax returns the UnitAmountExcludingTax field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetUnitAmountExcludingTax() int64 {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.UnitAmountExcludingTax
}

// GetUnitAmountExcludingTaxOk returns a tuple with the UnitAmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) GetUnitAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitAmountExcludingTax) {
		return nil, false
	}
	return o.UnitAmountExcludingTax, true
}

// HasUnitAmountExcludingTax returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) HasUnitAmountExcludingTax() bool {
	if o != nil && !IsNil(o.UnitAmountExcludingTax) {
		return true
	}

	return false
}

// SetUnitAmountExcludingTax gets a reference to the given int64 and assigns it to the UnitAmountExcludingTax field.
func (o *UnibeeApiMerchantInvoiceNewInvoiceItemParam) SetUnitAmountExcludingTax(v int64) {
	o.UnitAmountExcludingTax = &v
}

func (o UnibeeApiMerchantInvoiceNewInvoiceItemParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceNewInvoiceItemParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UnitAmountExcludingTax) {
		toSerialize["unitAmountExcludingTax"] = o.UnitAmountExcludingTax
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam struct {
	value *UnibeeApiMerchantInvoiceNewInvoiceItemParam
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) Get() *UnibeeApiMerchantInvoiceNewInvoiceItemParam {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) Set(val *UnibeeApiMerchantInvoiceNewInvoiceItemParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceNewInvoiceItemParam(val *UnibeeApiMerchantInvoiceNewInvoiceItemParam) *NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam {
	return &NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceNewInvoiceItemParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


