/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionRenewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionRenewPost200ResponseData{}

// MerchantSubscriptionRenewPost200ResponseData struct for MerchantSubscriptionRenewPost200ResponseData
type MerchantSubscriptionRenewPost200ResponseData struct {
	Link *string `json:"link,omitempty"`
	Paid *bool `json:"paid,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
}

// NewMerchantSubscriptionRenewPost200ResponseData instantiates a new MerchantSubscriptionRenewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionRenewPost200ResponseData() *MerchantSubscriptionRenewPost200ResponseData {
	this := MerchantSubscriptionRenewPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionRenewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionRenewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionRenewPost200ResponseDataWithDefaults() *MerchantSubscriptionRenewPost200ResponseData {
	this := MerchantSubscriptionRenewPost200ResponseData{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MerchantSubscriptionRenewPost200ResponseData) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *MerchantSubscriptionRenewPost200ResponseData) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionRenewPost200ResponseData) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *MerchantSubscriptionRenewPost200ResponseData) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

func (o MerchantSubscriptionRenewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionRenewPost200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantSubscriptionRenewPost200ResponseData struct {
	value *MerchantSubscriptionRenewPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionRenewPost200ResponseData) Get() *MerchantSubscriptionRenewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionRenewPost200ResponseData) Set(val *MerchantSubscriptionRenewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionRenewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionRenewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionRenewPost200ResponseData(val *MerchantSubscriptionRenewPost200ResponseData) *NullableMerchantSubscriptionRenewPost200ResponseData {
	return &NullableMerchantSubscriptionRenewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionRenewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionRenewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


