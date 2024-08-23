/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}

// UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq Cancel subscription at period end, the subscription will not turn to 'cancelled' at once but will cancelled at period end time, no invoice will generate, the flag 'cancelAtPeriodEnd' of subscription will be enabled
type UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq struct {
	// default product will use if productId not specified and subscriptionId is blank
	ProductId *int64 `json:"productId,omitempty"`
	// SubscriptionId, id of subscription, either SubscriptionId or UserId needed, The only one active subscription of userId will effect
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// UserId, either SubscriptionId or UserId needed, The only one active subscription will effect if userId provide instead of subscriptionId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetProductId(v int64) {
	o.ProductId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq struct {
	value *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Get() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Set(val *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(val *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	return &NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


