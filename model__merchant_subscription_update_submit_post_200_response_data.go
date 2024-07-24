/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionUpdateSubmitPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionUpdateSubmitPost200ResponseData{}

// MerchantSubscriptionUpdateSubmitPost200ResponseData struct for MerchantSubscriptionUpdateSubmitPost200ResponseData
type MerchantSubscriptionUpdateSubmitPost200ResponseData struct {
	Link *string `json:"link,omitempty"`
	Paid *bool `json:"paid,omitempty"`
	SubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"subscriptionPendingUpdate,omitempty"`
}

// NewMerchantSubscriptionUpdateSubmitPost200ResponseData instantiates a new MerchantSubscriptionUpdateSubmitPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionUpdateSubmitPost200ResponseData() *MerchantSubscriptionUpdateSubmitPost200ResponseData {
	this := MerchantSubscriptionUpdateSubmitPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionUpdateSubmitPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUpdateSubmitPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionUpdateSubmitPost200ResponseDataWithDefaults() *MerchantSubscriptionUpdateSubmitPost200ResponseData {
	this := MerchantSubscriptionUpdateSubmitPost200ResponseData{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscriptionPendingUpdate returns the SubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.SubscriptionPendingUpdate
}

// GetSubscriptionPendingUpdateOk returns a tuple with the SubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) GetSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		return nil, false
	}
	return o.SubscriptionPendingUpdate, true
}

// HasSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) HasSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.SubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the SubscriptionPendingUpdate field.
func (o *MerchantSubscriptionUpdateSubmitPost200ResponseData) SetSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.SubscriptionPendingUpdate = &v
}

func (o MerchantSubscriptionUpdateSubmitPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionUpdateSubmitPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if !IsNil(o.SubscriptionPendingUpdate) {
		toSerialize["subscriptionPendingUpdate"] = o.SubscriptionPendingUpdate
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionUpdateSubmitPost200ResponseData struct {
	value *MerchantSubscriptionUpdateSubmitPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) Get() *MerchantSubscriptionUpdateSubmitPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) Set(val *MerchantSubscriptionUpdateSubmitPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionUpdateSubmitPost200ResponseData(val *MerchantSubscriptionUpdateSubmitPost200ResponseData) *NullableMerchantSubscriptionUpdateSubmitPost200ResponseData {
	return &NullableMerchantSubscriptionUpdateSubmitPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionUpdateSubmitPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


