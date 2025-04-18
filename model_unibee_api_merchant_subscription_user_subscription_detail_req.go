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

// checks if the UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}

// UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct for UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq
type UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct {
	// ExternalUserId, unique, either ExternalUserId&Email or UserId needed
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// default product will use if productId not specified and subscriptionId is blank
	ProductId *int64 `json:"productId,omitempty"`
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}
	return &this
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetProductId(v int64) {
	o.ProductId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) ToMap() (map[string]interface{}, error) {
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

type NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct {
	value *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Get() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Set(val *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(val *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	return &NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


