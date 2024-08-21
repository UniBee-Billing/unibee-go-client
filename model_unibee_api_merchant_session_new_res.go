/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSessionNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSessionNewRes{}

// UnibeeApiMerchantSessionNewRes struct for UnibeeApiMerchantSessionNewRes
type UnibeeApiMerchantSessionNewRes struct {
	// ClientSession
	ClientSession *string `json:"clientSession,omitempty"`
	// ClientToken
	ClientToken *string `json:"clientToken,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// ExternalUserId
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Url
	Url *string `json:"url,omitempty"`
	// UserId
	UserId *string `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSessionNewRes instantiates a new UnibeeApiMerchantSessionNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSessionNewRes() *UnibeeApiMerchantSessionNewRes {
	this := UnibeeApiMerchantSessionNewRes{}
	return &this
}

// NewUnibeeApiMerchantSessionNewResWithDefaults instantiates a new UnibeeApiMerchantSessionNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSessionNewResWithDefaults() *UnibeeApiMerchantSessionNewRes {
	this := UnibeeApiMerchantSessionNewRes{}
	return &this
}

// GetClientSession returns the ClientSession field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetClientSession() string {
	if o == nil || IsNil(o.ClientSession) {
		var ret string
		return ret
	}
	return *o.ClientSession
}

// GetClientSessionOk returns a tuple with the ClientSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetClientSessionOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSession) {
		return nil, false
	}
	return o.ClientSession, true
}

// HasClientSession returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasClientSession() bool {
	if o != nil && !IsNil(o.ClientSession) {
		return true
	}

	return false
}

// SetClientSession gets a reference to the given string and assigns it to the ClientSession field.
func (o *UnibeeApiMerchantSessionNewRes) SetClientSession(v string) {
	o.ClientSession = &v
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetClientToken() string {
	if o == nil || IsNil(o.ClientToken) {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetClientTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientToken) {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasClientToken() bool {
	if o != nil && !IsNil(o.ClientToken) {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *UnibeeApiMerchantSessionNewRes) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantSessionNewRes) SetEmail(v string) {
	o.Email = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantSessionNewRes) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UnibeeApiMerchantSessionNewRes) SetUrl(v string) {
	o.Url = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewRes) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewRes) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewRes) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UnibeeApiMerchantSessionNewRes) SetUserId(v string) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSessionNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSessionNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientSession) {
		toSerialize["clientSession"] = o.ClientSession
	}
	if !IsNil(o.ClientToken) {
		toSerialize["clientToken"] = o.ClientToken
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSessionNewRes struct {
	value *UnibeeApiMerchantSessionNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSessionNewRes) Get() *UnibeeApiMerchantSessionNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSessionNewRes) Set(val *UnibeeApiMerchantSessionNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSessionNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSessionNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSessionNewRes(val *UnibeeApiMerchantSessionNewRes) *NullableUnibeeApiMerchantSessionNewRes {
	return &NullableUnibeeApiMerchantSessionNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSessionNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSessionNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


