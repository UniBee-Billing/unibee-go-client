/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceNewReq{}

// UnibeeApiMerchantInvoiceNewReq struct for UnibeeApiMerchantInvoiceNewReq
type UnibeeApiMerchantInvoiceNewReq struct {
	// Currency
	Currency string `json:"currency"`
	Finish *bool `json:"finish,omitempty"`
	// Gateway Id
	GatewayId int64 `json:"gatewayId"`
	Lines []UnibeeApiMerchantInvoiceNewInvoiceItemParam `json:"lines,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// TaxScale，1000 represent 10%
	TaxScale int64 `json:"taxScale"`
	// UserId
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantInvoiceNewReq UnibeeApiMerchantInvoiceNewReq

// NewUnibeeApiMerchantInvoiceNewReq instantiates a new UnibeeApiMerchantInvoiceNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceNewReq(currency string, gatewayId int64, taxScale int64, userId int64) *UnibeeApiMerchantInvoiceNewReq {
	this := UnibeeApiMerchantInvoiceNewReq{}
	this.Currency = currency
	this.GatewayId = gatewayId
	this.TaxScale = taxScale
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantInvoiceNewReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceNewReqWithDefaults() *UnibeeApiMerchantInvoiceNewReq {
	this := UnibeeApiMerchantInvoiceNewReq{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantInvoiceNewReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantInvoiceNewReq) SetCurrency(v string) {
	o.Currency = v
}

// GetFinish returns the Finish field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewReq) GetFinish() bool {
	if o == nil || IsNil(o.Finish) {
		var ret bool
		return ret
	}
	return *o.Finish
}

// GetFinishOk returns a tuple with the Finish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetFinishOk() (*bool, bool) {
	if o == nil || IsNil(o.Finish) {
		return nil, false
	}
	return o.Finish, true
}

// HasFinish returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) HasFinish() bool {
	if o != nil && !IsNil(o.Finish) {
		return true
	}

	return false
}

// SetFinish gets a reference to the given bool and assigns it to the Finish field.
func (o *UnibeeApiMerchantInvoiceNewReq) SetFinish(v bool) {
	o.Finish = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantInvoiceNewReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantInvoiceNewReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewReq) GetLines() []UnibeeApiMerchantInvoiceNewInvoiceItemParam {
	if o == nil || IsNil(o.Lines) {
		var ret []UnibeeApiMerchantInvoiceNewInvoiceItemParam
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetLinesOk() ([]UnibeeApiMerchantInvoiceNewInvoiceItemParam, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []UnibeeApiMerchantInvoiceNewInvoiceItemParam and assigns it to the Lines field.
func (o *UnibeeApiMerchantInvoiceNewReq) SetLines(v []UnibeeApiMerchantInvoiceNewInvoiceItemParam) {
	o.Lines = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantInvoiceNewReq) SetName(v string) {
	o.Name = &v
}

// GetTaxScale returns the TaxScale field value
func (o *UnibeeApiMerchantInvoiceNewReq) GetTaxScale() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TaxScale
}

// GetTaxScaleOk returns a tuple with the TaxScale field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetTaxScaleOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxScale, true
}

// SetTaxScale sets field value
func (o *UnibeeApiMerchantInvoiceNewReq) SetTaxScale(v int64) {
	o.TaxScale = v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantInvoiceNewReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantInvoiceNewReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantInvoiceNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Finish) {
		toSerialize["finish"] = o.Finish
	}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["taxScale"] = o.TaxScale
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"gatewayId",
		"taxScale",
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

	varUnibeeApiMerchantInvoiceNewReq := _UnibeeApiMerchantInvoiceNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceNewReq(varUnibeeApiMerchantInvoiceNewReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceNewReq struct {
	value *UnibeeApiMerchantInvoiceNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceNewReq) Get() *UnibeeApiMerchantInvoiceNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceNewReq) Set(val *UnibeeApiMerchantInvoiceNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceNewReq(val *UnibeeApiMerchantInvoiceNewReq) *NullableUnibeeApiMerchantInvoiceNewReq {
	return &NullableUnibeeApiMerchantInvoiceNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


