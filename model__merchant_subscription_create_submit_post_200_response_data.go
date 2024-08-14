/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionCreateSubmitPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionCreateSubmitPost200ResponseData{}

// MerchantSubscriptionCreateSubmitPost200ResponseData struct for MerchantSubscriptionCreateSubmitPost200ResponseData
type MerchantSubscriptionCreateSubmitPost200ResponseData struct {
	Link *string `json:"link,omitempty"`
	Paid *bool `json:"paid,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
	// token
	Token *string `json:"token,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewMerchantSubscriptionCreateSubmitPost200ResponseData instantiates a new MerchantSubscriptionCreateSubmitPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionCreateSubmitPost200ResponseData() *MerchantSubscriptionCreateSubmitPost200ResponseData {
	this := MerchantSubscriptionCreateSubmitPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionCreateSubmitPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionCreateSubmitPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionCreateSubmitPost200ResponseDataWithDefaults() *MerchantSubscriptionCreateSubmitPost200ResponseData {
	this := MerchantSubscriptionCreateSubmitPost200ResponseData{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) SetLink(v string) {
	o.Link = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) SetToken(v string) {
	o.Token = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *MerchantSubscriptionCreateSubmitPost200ResponseData) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o MerchantSubscriptionCreateSubmitPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionCreateSubmitPost200ResponseData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionCreateSubmitPost200ResponseData struct {
	value *MerchantSubscriptionCreateSubmitPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionCreateSubmitPost200ResponseData) Get() *MerchantSubscriptionCreateSubmitPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionCreateSubmitPost200ResponseData) Set(val *MerchantSubscriptionCreateSubmitPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionCreateSubmitPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionCreateSubmitPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionCreateSubmitPost200ResponseData(val *MerchantSubscriptionCreateSubmitPost200ResponseData) *NullableMerchantSubscriptionCreateSubmitPost200ResponseData {
	return &NullableMerchantSubscriptionCreateSubmitPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionCreateSubmitPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionCreateSubmitPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


