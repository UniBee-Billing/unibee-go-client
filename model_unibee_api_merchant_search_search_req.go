/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSearchSearchReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSearchSearchReq{}

// UnibeeApiMerchantSearchSearchReq struct for UnibeeApiMerchantSearchSearchReq
type UnibeeApiMerchantSearchSearchReq struct {
	// SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId
	SearchKey *string `json:"searchKey,omitempty"`
}

// NewUnibeeApiMerchantSearchSearchReq instantiates a new UnibeeApiMerchantSearchSearchReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSearchSearchReq() *UnibeeApiMerchantSearchSearchReq {
	this := UnibeeApiMerchantSearchSearchReq{}
	return &this
}

// NewUnibeeApiMerchantSearchSearchReqWithDefaults instantiates a new UnibeeApiMerchantSearchSearchReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSearchSearchReqWithDefaults() *UnibeeApiMerchantSearchSearchReq {
	this := UnibeeApiMerchantSearchSearchReq{}
	return &this
}

// GetSearchKey returns the SearchKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchSearchReq) GetSearchKey() string {
	if o == nil || IsNil(o.SearchKey) {
		var ret string
		return ret
	}
	return *o.SearchKey
}

// GetSearchKeyOk returns a tuple with the SearchKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchSearchReq) GetSearchKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SearchKey) {
		return nil, false
	}
	return o.SearchKey, true
}

// HasSearchKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchSearchReq) HasSearchKey() bool {
	if o != nil && !IsNil(o.SearchKey) {
		return true
	}

	return false
}

// SetSearchKey gets a reference to the given string and assigns it to the SearchKey field.
func (o *UnibeeApiMerchantSearchSearchReq) SetSearchKey(v string) {
	o.SearchKey = &v
}

func (o UnibeeApiMerchantSearchSearchReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSearchSearchReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchKey) {
		toSerialize["searchKey"] = o.SearchKey
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSearchSearchReq struct {
	value *UnibeeApiMerchantSearchSearchReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSearchSearchReq) Get() *UnibeeApiMerchantSearchSearchReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSearchSearchReq) Set(val *UnibeeApiMerchantSearchSearchReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSearchSearchReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSearchSearchReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSearchSearchReq(val *UnibeeApiMerchantSearchSearchReq) *NullableUnibeeApiMerchantSearchSearchReq {
	return &NullableUnibeeApiMerchantSearchSearchReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSearchSearchReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSearchSearchReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


