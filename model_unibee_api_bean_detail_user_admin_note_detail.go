/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailUserAdminNoteDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailUserAdminNoteDetail{}

// UnibeeApiBeanDetailUserAdminNoteDetail struct for UnibeeApiBeanDetailUserAdminNoteDetail
type UnibeeApiBeanDetailUserAdminNoteDetail struct {
	// CreateTime, UTC Time
	CreateTime *int64 `json:"createTime,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	MerchantMember *UnibeeApiBeanDetailMerchantMemberDetail `json:"merchantMember,omitempty"`
	// merchant_user_id
	MerchantMemberId *int64 `json:"merchantMemberId,omitempty"`
	// note
	Note *string `json:"note,omitempty"`
	UserAccount *UnibeeApiBeanUserAccount `json:"userAccount,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanDetailUserAdminNoteDetail instantiates a new UnibeeApiBeanDetailUserAdminNoteDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailUserAdminNoteDetail() *UnibeeApiBeanDetailUserAdminNoteDetail {
	this := UnibeeApiBeanDetailUserAdminNoteDetail{}
	return &this
}

// NewUnibeeApiBeanDetailUserAdminNoteDetailWithDefaults instantiates a new UnibeeApiBeanDetailUserAdminNoteDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailUserAdminNoteDetailWithDefaults() *UnibeeApiBeanDetailUserAdminNoteDetail {
	this := UnibeeApiBeanDetailUserAdminNoteDetail{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetId(v int64) {
	o.Id = &v
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the MerchantMember field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.MerchantMember = &v
}

// GetMerchantMemberId returns the MerchantMemberId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberId() int64 {
	if o == nil || IsNil(o.MerchantMemberId) {
		var ret int64
		return ret
	}
	return *o.MerchantMemberId
}

// GetMerchantMemberIdOk returns a tuple with the MerchantMemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantMemberId) {
		return nil, false
	}
	return o.MerchantMemberId, true
}

// HasMerchantMemberId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasMerchantMemberId() bool {
	if o != nil && !IsNil(o.MerchantMemberId) {
		return true
	}

	return false
}

// SetMerchantMemberId gets a reference to the given int64 and assigns it to the MerchantMemberId field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetMerchantMemberId(v int64) {
	o.MerchantMemberId = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetNote(v string) {
	o.Note = &v
}

// GetUserAccount returns the UserAccount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserAccount() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.UserAccount) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.UserAccount
}

// GetUserAccountOk returns a tuple with the UserAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserAccountOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.UserAccount) {
		return nil, false
	}
	return o.UserAccount, true
}

// HasUserAccount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasUserAccount() bool {
	if o != nil && !IsNil(o.UserAccount) {
		return true
	}

	return false
}

// SetUserAccount gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the UserAccount field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetUserAccount(v UnibeeApiBeanUserAccount) {
	o.UserAccount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanDetailUserAdminNoteDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailUserAdminNoteDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	if !IsNil(o.MerchantMemberId) {
		toSerialize["merchantMemberId"] = o.MerchantMemberId
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.UserAccount) {
		toSerialize["userAccount"] = o.UserAccount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailUserAdminNoteDetail struct {
	value *UnibeeApiBeanDetailUserAdminNoteDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailUserAdminNoteDetail) Get() *UnibeeApiBeanDetailUserAdminNoteDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailUserAdminNoteDetail) Set(val *UnibeeApiBeanDetailUserAdminNoteDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailUserAdminNoteDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailUserAdminNoteDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailUserAdminNoteDetail(val *UnibeeApiBeanDetailUserAdminNoteDetail) *NullableUnibeeApiBeanDetailUserAdminNoteDetail {
	return &NullableUnibeeApiBeanDetailUserAdminNoteDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailUserAdminNoteDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailUserAdminNoteDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


