/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSessionNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSessionNewReq{}

// UnibeeApiMerchantSessionNewReq struct for UnibeeApiMerchantSessionNewReq
type UnibeeApiMerchantSessionNewReq struct {
	// Address
	Address *string `json:"address,omitempty"`
	// Email
	Email string `json:"email"`
	// ExternalUserId
	ExternalUserId string `json:"externalUserId"`
	// First Name
	FirstName *string `json:"firstName,omitempty"`
	// Last Name
	LastName *string `json:"lastName,omitempty"`
	// Password
	Password *string `json:"password,omitempty"`
	// Phone
	Phone *string `json:"phone,omitempty"`
	// ReturnUrl
	ReturnUrl *string `json:"returnUrl,omitempty"`
}

type _UnibeeApiMerchantSessionNewReq UnibeeApiMerchantSessionNewReq

// NewUnibeeApiMerchantSessionNewReq instantiates a new UnibeeApiMerchantSessionNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSessionNewReq(email string, externalUserId string) *UnibeeApiMerchantSessionNewReq {
	this := UnibeeApiMerchantSessionNewReq{}
	this.Email = email
	this.ExternalUserId = externalUserId
	return &this
}

// NewUnibeeApiMerchantSessionNewReqWithDefaults instantiates a new UnibeeApiMerchantSessionNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSessionNewReqWithDefaults() *UnibeeApiMerchantSessionNewReq {
	this := UnibeeApiMerchantSessionNewReq{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UnibeeApiMerchantSessionNewReq) SetAddress(v string) {
	o.Address = &v
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantSessionNewReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantSessionNewReq) SetEmail(v string) {
	o.Email = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *UnibeeApiMerchantSessionNewReq) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *UnibeeApiMerchantSessionNewReq) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiMerchantSessionNewReq) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiMerchantSessionNewReq) SetLastName(v string) {
	o.LastName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UnibeeApiMerchantSessionNewReq) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeApiMerchantSessionNewReq) SetPhone(v string) {
	o.Phone = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSessionNewReq) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSessionNewReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSessionNewReq) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiMerchantSessionNewReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

func (o UnibeeApiMerchantSessionNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSessionNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["email"] = o.Email
	toSerialize["externalUserId"] = o.ExternalUserId
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSessionNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"externalUserId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantSessionNewReq := _UnibeeApiMerchantSessionNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSessionNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSessionNewReq(varUnibeeApiMerchantSessionNewReq)

	return err
}

type NullableUnibeeApiMerchantSessionNewReq struct {
	value *UnibeeApiMerchantSessionNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSessionNewReq) Get() *UnibeeApiMerchantSessionNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSessionNewReq) Set(val *UnibeeApiMerchantSessionNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSessionNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSessionNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSessionNewReq(val *UnibeeApiMerchantSessionNewReq) *NullableUnibeeApiMerchantSessionNewReq {
	return &NullableUnibeeApiMerchantSessionNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSessionNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSessionNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


