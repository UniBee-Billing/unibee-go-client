/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionRenewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionRenewRes{}

// UnibeeApiMerchantSubscriptionRenewRes struct for UnibeeApiMerchantSubscriptionRenewRes
type UnibeeApiMerchantSubscriptionRenewRes struct {
	Link *string `json:"link,omitempty"`
	Paid *bool `json:"paid,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionRenewRes instantiates a new UnibeeApiMerchantSubscriptionRenewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionRenewRes() *UnibeeApiMerchantSubscriptionRenewRes {
	this := UnibeeApiMerchantSubscriptionRenewRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionRenewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionRenewResWithDefaults() *UnibeeApiMerchantSubscriptionRenewRes {
	this := UnibeeApiMerchantSubscriptionRenewRes{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantSubscriptionRenewRes) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *UnibeeApiMerchantSubscriptionRenewRes) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewRes) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *UnibeeApiMerchantSubscriptionRenewRes) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

func (o UnibeeApiMerchantSubscriptionRenewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionRenewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionRenewRes struct {
	value *UnibeeApiMerchantSubscriptionRenewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionRenewRes) Get() *UnibeeApiMerchantSubscriptionRenewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewRes) Set(val *UnibeeApiMerchantSubscriptionRenewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionRenewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionRenewRes(val *UnibeeApiMerchantSubscriptionRenewRes) *NullableUnibeeApiMerchantSubscriptionRenewRes {
	return &NullableUnibeeApiMerchantSubscriptionRenewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionRenewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


