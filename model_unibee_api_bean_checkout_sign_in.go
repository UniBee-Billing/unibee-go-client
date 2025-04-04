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

// checks if the UnibeeApiBeanCheckoutSignIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCheckoutSignIn{}

// UnibeeApiBeanCheckoutSignIn struct for UnibeeApiBeanCheckoutSignIn
type UnibeeApiBeanCheckoutSignIn struct {
	// whether contain active or incomplete subscription or not
	DuplicateSubscription *bool `json:"DuplicateSubscription,omitempty"`
	// should redirect to sign in page
	Redirect *bool `json:"redirect,omitempty"`
	// redirect url
	Url *string `json:"url,omitempty"`
}

// NewUnibeeApiBeanCheckoutSignIn instantiates a new UnibeeApiBeanCheckoutSignIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCheckoutSignIn() *UnibeeApiBeanCheckoutSignIn {
	this := UnibeeApiBeanCheckoutSignIn{}
	return &this
}

// NewUnibeeApiBeanCheckoutSignInWithDefaults instantiates a new UnibeeApiBeanCheckoutSignIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCheckoutSignInWithDefaults() *UnibeeApiBeanCheckoutSignIn {
	this := UnibeeApiBeanCheckoutSignIn{}
	return &this
}

// GetDuplicateSubscription returns the DuplicateSubscription field value if set, zero value otherwise.
func (o *UnibeeApiBeanCheckoutSignIn) GetDuplicateSubscription() bool {
	if o == nil || IsNil(o.DuplicateSubscription) {
		var ret bool
		return ret
	}
	return *o.DuplicateSubscription
}

// GetDuplicateSubscriptionOk returns a tuple with the DuplicateSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCheckoutSignIn) GetDuplicateSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.DuplicateSubscription) {
		return nil, false
	}
	return o.DuplicateSubscription, true
}

// HasDuplicateSubscription returns a boolean if a field has been set.
func (o *UnibeeApiBeanCheckoutSignIn) HasDuplicateSubscription() bool {
	if o != nil && !IsNil(o.DuplicateSubscription) {
		return true
	}

	return false
}

// SetDuplicateSubscription gets a reference to the given bool and assigns it to the DuplicateSubscription field.
func (o *UnibeeApiBeanCheckoutSignIn) SetDuplicateSubscription(v bool) {
	o.DuplicateSubscription = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *UnibeeApiBeanCheckoutSignIn) GetRedirect() bool {
	if o == nil || IsNil(o.Redirect) {
		var ret bool
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCheckoutSignIn) GetRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *UnibeeApiBeanCheckoutSignIn) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given bool and assigns it to the Redirect field.
func (o *UnibeeApiBeanCheckoutSignIn) SetRedirect(v bool) {
	o.Redirect = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UnibeeApiBeanCheckoutSignIn) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCheckoutSignIn) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanCheckoutSignIn) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UnibeeApiBeanCheckoutSignIn) SetUrl(v string) {
	o.Url = &v
}

func (o UnibeeApiBeanCheckoutSignIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCheckoutSignIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DuplicateSubscription) {
		toSerialize["DuplicateSubscription"] = o.DuplicateSubscription
	}
	if !IsNil(o.Redirect) {
		toSerialize["redirect"] = o.Redirect
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCheckoutSignIn struct {
	value *UnibeeApiBeanCheckoutSignIn
	isSet bool
}

func (v NullableUnibeeApiBeanCheckoutSignIn) Get() *UnibeeApiBeanCheckoutSignIn {
	return v.value
}

func (v *NullableUnibeeApiBeanCheckoutSignIn) Set(val *UnibeeApiBeanCheckoutSignIn) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCheckoutSignIn) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCheckoutSignIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCheckoutSignIn(val *UnibeeApiBeanCheckoutSignIn) *NullableUnibeeApiBeanCheckoutSignIn {
	return &NullableUnibeeApiBeanCheckoutSignIn{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCheckoutSignIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCheckoutSignIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


