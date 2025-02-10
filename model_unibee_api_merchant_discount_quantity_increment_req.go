/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountQuantityIncrementReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountQuantityIncrementReq{}

// UnibeeApiMerchantDiscountQuantityIncrementReq Increase discount code quantity, if original quantity is 0, increase greater than 0 will enable quantity control
type UnibeeApiMerchantDiscountQuantityIncrementReq struct {
	// The discount quantity amount to increase, should greater than 0
	Amount *int64 `json:"amount,omitempty"`
	// The discount's Id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantDiscountQuantityIncrementReq UnibeeApiMerchantDiscountQuantityIncrementReq

// NewUnibeeApiMerchantDiscountQuantityIncrementReq instantiates a new UnibeeApiMerchantDiscountQuantityIncrementReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountQuantityIncrementReq(id int64) *UnibeeApiMerchantDiscountQuantityIncrementReq {
	this := UnibeeApiMerchantDiscountQuantityIncrementReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountQuantityIncrementReqWithDefaults instantiates a new UnibeeApiMerchantDiscountQuantityIncrementReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountQuantityIncrementReqWithDefaults() *UnibeeApiMerchantDiscountQuantityIncrementReq {
	this := UnibeeApiMerchantDiscountQuantityIncrementReq{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) SetAmount(v int64) {
	o.Amount = &v
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantDiscountQuantityIncrementReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountQuantityIncrementReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUnibeeApiMerchantDiscountQuantityIncrementReq := _UnibeeApiMerchantDiscountQuantityIncrementReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountQuantityIncrementReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountQuantityIncrementReq(varUnibeeApiMerchantDiscountQuantityIncrementReq)

	return err
}

type NullableUnibeeApiMerchantDiscountQuantityIncrementReq struct {
	value *UnibeeApiMerchantDiscountQuantityIncrementReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountQuantityIncrementReq) Get() *UnibeeApiMerchantDiscountQuantityIncrementReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountQuantityIncrementReq) Set(val *UnibeeApiMerchantDiscountQuantityIncrementReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountQuantityIncrementReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountQuantityIncrementReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountQuantityIncrementReq(val *UnibeeApiMerchantDiscountQuantityIncrementReq) *NullableUnibeeApiMerchantDiscountQuantityIncrementReq {
	return &NullableUnibeeApiMerchantDiscountQuantityIncrementReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountQuantityIncrementReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountQuantityIncrementReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


