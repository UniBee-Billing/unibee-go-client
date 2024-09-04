/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthRegisterReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthRegisterReq{}

// UnibeeApiMerchantAuthRegisterReq Register merchant with owner, send email to owner's email address with OTP code, only open for cloud version; the owner account will create automatic for standalone version
type UnibeeApiMerchantAuthRegisterReq struct {
	// The merchant owner's email address
	Email string `json:"email"`
	// The merchant owner's first name
	FirstName string `json:"firstName"`
	// The merchant owner's last name
	LastName string `json:"lastName"`
	// The owner's password
	Password string `json:"password"`
	// The owner's Phone
	Phone *string `json:"phone,omitempty"`
	// The owner's UserName
	UserName *string `json:"userName,omitempty"`
}

type _UnibeeApiMerchantAuthRegisterReq UnibeeApiMerchantAuthRegisterReq

// NewUnibeeApiMerchantAuthRegisterReq instantiates a new UnibeeApiMerchantAuthRegisterReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthRegisterReq(email string, firstName string, lastName string, password string) *UnibeeApiMerchantAuthRegisterReq {
	this := UnibeeApiMerchantAuthRegisterReq{}
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.Password = password
	return &this
}

// NewUnibeeApiMerchantAuthRegisterReqWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthRegisterReqWithDefaults() *UnibeeApiMerchantAuthRegisterReq {
	this := UnibeeApiMerchantAuthRegisterReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthRegisterReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthRegisterReq) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *UnibeeApiMerchantAuthRegisterReq) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UnibeeApiMerchantAuthRegisterReq) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UnibeeApiMerchantAuthRegisterReq) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UnibeeApiMerchantAuthRegisterReq) SetLastName(v string) {
	o.LastName = v
}

// GetPassword returns the Password field value
func (o *UnibeeApiMerchantAuthRegisterReq) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UnibeeApiMerchantAuthRegisterReq) SetPassword(v string) {
	o.Password = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthRegisterReq) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeApiMerchantAuthRegisterReq) SetPhone(v string) {
	o.Phone = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthRegisterReq) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthRegisterReq) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UnibeeApiMerchantAuthRegisterReq) SetUserName(v string) {
	o.UserName = &v
}

func (o UnibeeApiMerchantAuthRegisterReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthRegisterReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["password"] = o.Password
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthRegisterReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"firstName",
		"lastName",
		"password",
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

	varUnibeeApiMerchantAuthRegisterReq := _UnibeeApiMerchantAuthRegisterReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthRegisterReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthRegisterReq(varUnibeeApiMerchantAuthRegisterReq)

	return err
}

type NullableUnibeeApiMerchantAuthRegisterReq struct {
	value *UnibeeApiMerchantAuthRegisterReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthRegisterReq) Get() *UnibeeApiMerchantAuthRegisterReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthRegisterReq) Set(val *UnibeeApiMerchantAuthRegisterReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthRegisterReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthRegisterReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthRegisterReq(val *UnibeeApiMerchantAuthRegisterReq) *NullableUnibeeApiMerchantAuthRegisterReq {
	return &NullableUnibeeApiMerchantAuthRegisterReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthRegisterReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthRegisterReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


