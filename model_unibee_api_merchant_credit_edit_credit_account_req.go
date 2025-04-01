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

// checks if the UnibeeApiMerchantCreditEditCreditAccountReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditEditCreditAccountReq{}

// UnibeeApiMerchantCreditEditCreditAccountReq Edit User Credit Account Config
type UnibeeApiMerchantCreditEditCreditAccountReq struct {
	// id of credit account
	Id int64 `json:"id"`
	// credit account can used or payout|apply in purchase or not, 0-no, 1-yes
	PayoutEnable *int32 `json:"payoutEnable,omitempty"`
	// credit account can be recharged|increment or not, 0-no, 1-yes
	RechargeEnable *int32 `json:"rechargeEnable,omitempty"`
}

type _UnibeeApiMerchantCreditEditCreditAccountReq UnibeeApiMerchantCreditEditCreditAccountReq

// NewUnibeeApiMerchantCreditEditCreditAccountReq instantiates a new UnibeeApiMerchantCreditEditCreditAccountReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditEditCreditAccountReq(id int64) *UnibeeApiMerchantCreditEditCreditAccountReq {
	this := UnibeeApiMerchantCreditEditCreditAccountReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantCreditEditCreditAccountReqWithDefaults instantiates a new UnibeeApiMerchantCreditEditCreditAccountReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditEditCreditAccountReqWithDefaults() *UnibeeApiMerchantCreditEditCreditAccountReq {
	this := UnibeeApiMerchantCreditEditCreditAccountReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetId(v int64) {
	o.Id = v
}

// GetPayoutEnable returns the PayoutEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetPayoutEnable() int32 {
	if o == nil || IsNil(o.PayoutEnable) {
		var ret int32
		return ret
	}
	return *o.PayoutEnable
}

// GetPayoutEnableOk returns a tuple with the PayoutEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetPayoutEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.PayoutEnable) {
		return nil, false
	}
	return o.PayoutEnable, true
}

// HasPayoutEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) HasPayoutEnable() bool {
	if o != nil && !IsNil(o.PayoutEnable) {
		return true
	}

	return false
}

// SetPayoutEnable gets a reference to the given int32 and assigns it to the PayoutEnable field.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetPayoutEnable(v int32) {
	o.PayoutEnable = &v
}

// GetRechargeEnable returns the RechargeEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetRechargeEnable() int32 {
	if o == nil || IsNil(o.RechargeEnable) {
		var ret int32
		return ret
	}
	return *o.RechargeEnable
}

// GetRechargeEnableOk returns a tuple with the RechargeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) GetRechargeEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.RechargeEnable) {
		return nil, false
	}
	return o.RechargeEnable, true
}

// HasRechargeEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) HasRechargeEnable() bool {
	if o != nil && !IsNil(o.RechargeEnable) {
		return true
	}

	return false
}

// SetRechargeEnable gets a reference to the given int32 and assigns it to the RechargeEnable field.
func (o *UnibeeApiMerchantCreditEditCreditAccountReq) SetRechargeEnable(v int32) {
	o.RechargeEnable = &v
}

func (o UnibeeApiMerchantCreditEditCreditAccountReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditEditCreditAccountReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.PayoutEnable) {
		toSerialize["payoutEnable"] = o.PayoutEnable
	}
	if !IsNil(o.RechargeEnable) {
		toSerialize["rechargeEnable"] = o.RechargeEnable
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantCreditEditCreditAccountReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantCreditEditCreditAccountReq := _UnibeeApiMerchantCreditEditCreditAccountReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantCreditEditCreditAccountReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantCreditEditCreditAccountReq(varUnibeeApiMerchantCreditEditCreditAccountReq)

	return err
}

type NullableUnibeeApiMerchantCreditEditCreditAccountReq struct {
	value *UnibeeApiMerchantCreditEditCreditAccountReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountReq) Get() *UnibeeApiMerchantCreditEditCreditAccountReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountReq) Set(val *UnibeeApiMerchantCreditEditCreditAccountReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditEditCreditAccountReq(val *UnibeeApiMerchantCreditEditCreditAccountReq) *NullableUnibeeApiMerchantCreditEditCreditAccountReq {
	return &NullableUnibeeApiMerchantCreditEditCreditAccountReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditEditCreditAccountReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditEditCreditAccountReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


