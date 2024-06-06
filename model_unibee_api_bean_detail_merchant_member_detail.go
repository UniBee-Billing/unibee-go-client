/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailMerchantMemberDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailMerchantMemberDetail{}

// UnibeeApiBeanDetailMerchantMemberDetail struct for UnibeeApiBeanDetailMerchantMemberDetail
type UnibeeApiBeanDetailMerchantMemberDetail struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// email
	Email *string `json:"email,omitempty"`
	// first name
	FirstName *string `json:"firstName,omitempty"`
	// userId
	Id *int64 `json:"id,omitempty"`
	// last name
	LastName *string `json:"lastName,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// mobile
	Mobile *string `json:"mobile,omitempty"`
	// permissions
	Permissions []UnibeeApiBeanMerchantRolePermission `json:"permissions,omitempty"`
	// role
	Role *string `json:"role,omitempty"`
}

// NewUnibeeApiBeanDetailMerchantMemberDetail instantiates a new UnibeeApiBeanDetailMerchantMemberDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailMerchantMemberDetail() *UnibeeApiBeanDetailMerchantMemberDetail {
	this := UnibeeApiBeanDetailMerchantMemberDetail{}
	return &this
}

// NewUnibeeApiBeanDetailMerchantMemberDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantMemberDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailMerchantMemberDetailWithDefaults() *UnibeeApiBeanDetailMerchantMemberDetail {
	this := UnibeeApiBeanDetailMerchantMemberDetail{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetId(v int64) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetLastName(v string) {
	o.LastName = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMobile(v string) {
	o.Mobile = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetPermissions() []UnibeeApiBeanMerchantRolePermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []UnibeeApiBeanMerchantRolePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetPermissionsOk() ([]UnibeeApiBeanMerchantRolePermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UnibeeApiBeanMerchantRolePermission and assigns it to the Permissions field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetPermissions(v []UnibeeApiBeanMerchantRolePermission) {
	o.Permissions = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetRole(v string) {
	o.Role = &v
}

func (o UnibeeApiBeanDetailMerchantMemberDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailMerchantMemberDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailMerchantMemberDetail struct {
	value *UnibeeApiBeanDetailMerchantMemberDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailMerchantMemberDetail) Get() *UnibeeApiBeanDetailMerchantMemberDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailMerchantMemberDetail) Set(val *UnibeeApiBeanDetailMerchantMemberDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailMerchantMemberDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailMerchantMemberDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailMerchantMemberDetail(val *UnibeeApiBeanDetailMerchantMemberDetail) *NullableUnibeeApiBeanDetailMerchantMemberDetail {
	return &NullableUnibeeApiBeanDetailMerchantMemberDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailMerchantMemberDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailMerchantMemberDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


