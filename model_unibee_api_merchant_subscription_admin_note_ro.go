/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionAdminNoteRo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionAdminNoteRo{}

// UnibeeApiMerchantSubscriptionAdminNoteRo struct for UnibeeApiMerchantSubscriptionAdminNoteRo
type UnibeeApiMerchantSubscriptionAdminNoteRo struct {
	// CreateTime, UTC Time
	CreateTime *int64 `json:"createTime,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// FirstName
	FirstName *string `json:"firstName,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// LastName
	LastName *string `json:"lastName,omitempty"`
	// Mobile
	Mobile *string `json:"mobile,omitempty"`
	// Note
	Note *string `json:"note,omitempty"`
	// SubscriptionId
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// UserName
	UserName *string `json:"userName,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionAdminNoteRo instantiates a new UnibeeApiMerchantSubscriptionAdminNoteRo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionAdminNoteRo() *UnibeeApiMerchantSubscriptionAdminNoteRo {
	this := UnibeeApiMerchantSubscriptionAdminNoteRo{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionAdminNoteRoWithDefaults instantiates a new UnibeeApiMerchantSubscriptionAdminNoteRo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionAdminNoteRoWithDefaults() *UnibeeApiMerchantSubscriptionAdminNoteRo {
	this := UnibeeApiMerchantSubscriptionAdminNoteRo{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetId(v int64) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetLastName(v string) {
	o.LastName = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetMobile(v string) {
	o.Mobile = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetNote(v string) {
	o.Note = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteRo) SetUserName(v string) {
	o.UserName = &v
}

func (o UnibeeApiMerchantSubscriptionAdminNoteRo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionAdminNoteRo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionAdminNoteRo struct {
	value *UnibeeApiMerchantSubscriptionAdminNoteRo
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteRo) Get() *UnibeeApiMerchantSubscriptionAdminNoteRo {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteRo) Set(val *UnibeeApiMerchantSubscriptionAdminNoteRo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteRo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteRo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionAdminNoteRo(val *UnibeeApiMerchantSubscriptionAdminNoteRo) *NullableUnibeeApiMerchantSubscriptionAdminNoteRo {
	return &NullableUnibeeApiMerchantSubscriptionAdminNoteRo{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteRo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteRo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


