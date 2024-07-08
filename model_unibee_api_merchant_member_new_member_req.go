/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMemberNewMemberReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberNewMemberReq{}

// UnibeeApiMerchantMemberNewMemberReq Will send email to member email provided, member can enter admin portal by email otp login
type UnibeeApiMerchantMemberNewMemberReq struct {
	// The email of member
	Email string `json:"email"`
	// The firstName of member
	FirstName *string `json:"firstName,omitempty"`
	// The lastName of member
	LastName *string `json:"lastName,omitempty"`
	// The id list of role
	RoleIds []int64 `json:"roleIds"`
}

type _UnibeeApiMerchantMemberNewMemberReq UnibeeApiMerchantMemberNewMemberReq

// NewUnibeeApiMerchantMemberNewMemberReq instantiates a new UnibeeApiMerchantMemberNewMemberReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberNewMemberReq(email string, roleIds []int64) *UnibeeApiMerchantMemberNewMemberReq {
	this := UnibeeApiMerchantMemberNewMemberReq{}
	this.Email = email
	this.RoleIds = roleIds
	return &this
}

// NewUnibeeApiMerchantMemberNewMemberReqWithDefaults instantiates a new UnibeeApiMerchantMemberNewMemberReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberNewMemberReqWithDefaults() *UnibeeApiMerchantMemberNewMemberReq {
	this := UnibeeApiMerchantMemberNewMemberReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantMemberNewMemberReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantMemberNewMemberReq) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiMerchantMemberNewMemberReq) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiMerchantMemberNewMemberReq) SetLastName(v string) {
	o.LastName = &v
}

// GetRoleIds returns the RoleIds field value
func (o *UnibeeApiMerchantMemberNewMemberReq) GetRoleIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberNewMemberReq) GetRoleIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *UnibeeApiMerchantMemberNewMemberReq) SetRoleIds(v []int64) {
	o.RoleIds = v
}

func (o UnibeeApiMerchantMemberNewMemberReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberNewMemberReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["roleIds"] = o.RoleIds
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMemberNewMemberReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"roleIds",
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

	varUnibeeApiMerchantMemberNewMemberReq := _UnibeeApiMerchantMemberNewMemberReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMemberNewMemberReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMemberNewMemberReq(varUnibeeApiMerchantMemberNewMemberReq)

	return err
}

type NullableUnibeeApiMerchantMemberNewMemberReq struct {
	value *UnibeeApiMerchantMemberNewMemberReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberNewMemberReq) Get() *UnibeeApiMerchantMemberNewMemberReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberNewMemberReq) Set(val *UnibeeApiMerchantMemberNewMemberReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberNewMemberReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberNewMemberReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberNewMemberReq(val *UnibeeApiMerchantMemberNewMemberReq) *NullableUnibeeApiMerchantMemberNewMemberReq {
	return &NullableUnibeeApiMerchantMemberNewMemberReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberNewMemberReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberNewMemberReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


