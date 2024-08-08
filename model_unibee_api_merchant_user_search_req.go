/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserSearchReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserSearchReq{}

// UnibeeApiMerchantUserSearchReq struct for UnibeeApiMerchantUserSearchReq
type UnibeeApiMerchantUserSearchReq struct {
	// SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId
	SearchKey *string `json:"searchKey,omitempty"`
}

// NewUnibeeApiMerchantUserSearchReq instantiates a new UnibeeApiMerchantUserSearchReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserSearchReq() *UnibeeApiMerchantUserSearchReq {
	this := UnibeeApiMerchantUserSearchReq{}
	return &this
}

// NewUnibeeApiMerchantUserSearchReqWithDefaults instantiates a new UnibeeApiMerchantUserSearchReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserSearchReqWithDefaults() *UnibeeApiMerchantUserSearchReq {
	this := UnibeeApiMerchantUserSearchReq{}
	return &this
}

// GetSearchKey returns the SearchKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserSearchReq) GetSearchKey() string {
	if o == nil || IsNil(o.SearchKey) {
		var ret string
		return ret
	}
	return *o.SearchKey
}

// GetSearchKeyOk returns a tuple with the SearchKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserSearchReq) GetSearchKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SearchKey) {
		return nil, false
	}
	return o.SearchKey, true
}

// HasSearchKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserSearchReq) HasSearchKey() bool {
	if o != nil && !IsNil(o.SearchKey) {
		return true
	}

	return false
}

// SetSearchKey gets a reference to the given string and assigns it to the SearchKey field.
func (o *UnibeeApiMerchantUserSearchReq) SetSearchKey(v string) {
	o.SearchKey = &v
}

func (o UnibeeApiMerchantUserSearchReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserSearchReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchKey) {
		toSerialize["searchKey"] = o.SearchKey
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserSearchReq struct {
	value *UnibeeApiMerchantUserSearchReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserSearchReq) Get() *UnibeeApiMerchantUserSearchReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserSearchReq) Set(val *UnibeeApiMerchantUserSearchReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserSearchReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserSearchReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserSearchReq(val *UnibeeApiMerchantUserSearchReq) *NullableUnibeeApiMerchantUserSearchReq {
	return &NullableUnibeeApiMerchantUserSearchReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserSearchReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserSearchReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


