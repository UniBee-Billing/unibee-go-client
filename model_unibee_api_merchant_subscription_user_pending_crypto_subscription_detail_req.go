/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq{}

// UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq struct for UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq
type UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq struct {
	// ExternalUserId, unique, either ExternalUserId&Email or UserId needed
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// default product will use if productId not specified and subscriptionId is blank
	ProductId *int64 `json:"productId,omitempty"`
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq instantiates a new UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq{}
	return &this
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) SetProductId(v int64) {
	o.ProductId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq struct {
	value *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) Get() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) Set(val *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq(val *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq {
	return &NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


