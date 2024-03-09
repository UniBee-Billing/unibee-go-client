/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoMerchantMemberSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoMerchantMemberSimplify{}

// UnibeeInternalLogicGatewayRoMerchantMemberSimplify struct for UnibeeInternalLogicGatewayRoMerchantMemberSimplify
type UnibeeInternalLogicGatewayRoMerchantMemberSimplify struct {
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
}

// NewUnibeeInternalLogicGatewayRoMerchantMemberSimplify instantiates a new UnibeeInternalLogicGatewayRoMerchantMemberSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoMerchantMemberSimplify() *UnibeeInternalLogicGatewayRoMerchantMemberSimplify {
	this := UnibeeInternalLogicGatewayRoMerchantMemberSimplify{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoMerchantMemberSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoMerchantMemberSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoMerchantMemberSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoMerchantMemberSimplify {
	this := UnibeeInternalLogicGatewayRoMerchantMemberSimplify{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetId(v int64) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetLastName(v string) {
	o.LastName = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

func (o UnibeeInternalLogicGatewayRoMerchantMemberSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoMerchantMemberSimplify) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify struct {
	value *UnibeeInternalLogicGatewayRoMerchantMemberSimplify
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) Get() *UnibeeInternalLogicGatewayRoMerchantMemberSimplify {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) Set(val *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify(val *UnibeeInternalLogicGatewayRoMerchantMemberSimplify) *NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify {
	return &NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoMerchantMemberSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


