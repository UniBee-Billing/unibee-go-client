/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceEditReq{}

// UnibeeApiMerchantInvoiceEditReq Edit invoice of pending status
type UnibeeApiMerchantInvoiceEditReq struct {
	// The currency of invoice
	Currency *string `json:"currency,omitempty"`
	Finish *bool `json:"finish,omitempty"`
	// The gateway id of invoice
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	Lines []UnibeeApiMerchantInvoiceNewInvoiceItemParam `json:"lines,omitempty"`
	// The name of invoice
	Name *string `json:"name,omitempty"`
	// The tax percentage of invoice，1000=10%
	TaxPercentage *int64 `json:"taxPercentage,omitempty"`
}

type _UnibeeApiMerchantInvoiceEditReq UnibeeApiMerchantInvoiceEditReq

// NewUnibeeApiMerchantInvoiceEditReq instantiates a new UnibeeApiMerchantInvoiceEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceEditReq(invoiceId string) *UnibeeApiMerchantInvoiceEditReq {
	this := UnibeeApiMerchantInvoiceEditReq{}
	this.InvoiceId = invoiceId
	return &this
}

// NewUnibeeApiMerchantInvoiceEditReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceEditReqWithDefaults() *UnibeeApiMerchantInvoiceEditReq {
	this := UnibeeApiMerchantInvoiceEditReq{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetFinish returns the Finish field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetFinish() bool {
	if o == nil || IsNil(o.Finish) {
		var ret bool
		return ret
	}
	return *o.Finish
}

// GetFinishOk returns a tuple with the Finish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetFinishOk() (*bool, bool) {
	if o == nil || IsNil(o.Finish) {
		return nil, false
	}
	return o.Finish, true
}

// HasFinish returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasFinish() bool {
	if o != nil && !IsNil(o.Finish) {
		return true
	}

	return false
}

// SetFinish gets a reference to the given bool and assigns it to the Finish field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetFinish(v bool) {
	o.Finish = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceEditReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceEditReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetLines() []UnibeeApiMerchantInvoiceNewInvoiceItemParam {
	if o == nil || IsNil(o.Lines) {
		var ret []UnibeeApiMerchantInvoiceNewInvoiceItemParam
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetLinesOk() ([]UnibeeApiMerchantInvoiceNewInvoiceItemParam, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []UnibeeApiMerchantInvoiceNewInvoiceItemParam and assigns it to the Lines field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetLines(v []UnibeeApiMerchantInvoiceNewInvoiceItemParam) {
	o.Lines = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetName(v string) {
	o.Name = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditReq) GetTaxPercentage() int64 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditReq) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantInvoiceEditReq) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

func (o UnibeeApiMerchantInvoiceEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Finish) {
		toSerialize["finish"] = o.Finish
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	toSerialize["invoiceId"] = o.InvoiceId
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
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

	varUnibeeApiMerchantInvoiceEditReq := _UnibeeApiMerchantInvoiceEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceEditReq(varUnibeeApiMerchantInvoiceEditReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceEditReq struct {
	value *UnibeeApiMerchantInvoiceEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceEditReq) Get() *UnibeeApiMerchantInvoiceEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceEditReq) Set(val *UnibeeApiMerchantInvoiceEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceEditReq(val *UnibeeApiMerchantInvoiceEditReq) *NullableUnibeeApiMerchantInvoiceEditReq {
	return &NullableUnibeeApiMerchantInvoiceEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


