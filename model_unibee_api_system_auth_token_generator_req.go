/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemAuthTokenGeneratorReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemAuthTokenGeneratorReq{}

// UnibeeApiSystemAuthTokenGeneratorReq struct for UnibeeApiSystemAuthTokenGeneratorReq
type UnibeeApiSystemAuthTokenGeneratorReq struct {
	Email string `json:"email"`
	// default daily
	Env *string `json:"env,omitempty"`
	MerchantId *int64 `json:"merchantId,omitempty"`
	// 0-Admin Portal, 1-User Portal, Default 0
	PortalType *int64 `json:"portalType,omitempty"`
	// default 1
	RedisDatabase *int32 `json:"redisDatabase,omitempty"`
}

type _UnibeeApiSystemAuthTokenGeneratorReq UnibeeApiSystemAuthTokenGeneratorReq

// NewUnibeeApiSystemAuthTokenGeneratorReq instantiates a new UnibeeApiSystemAuthTokenGeneratorReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemAuthTokenGeneratorReq(email string) *UnibeeApiSystemAuthTokenGeneratorReq {
	this := UnibeeApiSystemAuthTokenGeneratorReq{}
	this.Email = email
	var env string = "daily"
	this.Env = &env
	var merchantId int64 = 15621
	this.MerchantId = &merchantId
	var portalType int64 = 0
	this.PortalType = &portalType
	var redisDatabase int32 = 1
	this.RedisDatabase = &redisDatabase
	return &this
}

// NewUnibeeApiSystemAuthTokenGeneratorReqWithDefaults instantiates a new UnibeeApiSystemAuthTokenGeneratorReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemAuthTokenGeneratorReqWithDefaults() *UnibeeApiSystemAuthTokenGeneratorReq {
	this := UnibeeApiSystemAuthTokenGeneratorReq{}
	var email string = ""
	this.Email = email
	var env string = "daily"
	this.Env = &env
	var merchantId int64 = 15621
	this.MerchantId = &merchantId
	var portalType int64 = 0
	this.PortalType = &portalType
	var redisDatabase int32 = 1
	this.RedisDatabase = &redisDatabase
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetEmail(v string) {
	o.Email = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEnv() string {
	if o == nil || IsNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEnvOk() (*string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetEnv(v string) {
	o.Env = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPortalType returns the PortalType field value if set, zero value otherwise.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetPortalType() int64 {
	if o == nil || IsNil(o.PortalType) {
		var ret int64
		return ret
	}
	return *o.PortalType
}

// GetPortalTypeOk returns a tuple with the PortalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetPortalTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.PortalType) {
		return nil, false
	}
	return o.PortalType, true
}

// HasPortalType returns a boolean if a field has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasPortalType() bool {
	if o != nil && !IsNil(o.PortalType) {
		return true
	}

	return false
}

// SetPortalType gets a reference to the given int64 and assigns it to the PortalType field.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetPortalType(v int64) {
	o.PortalType = &v
}

// GetRedisDatabase returns the RedisDatabase field value if set, zero value otherwise.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetRedisDatabase() int32 {
	if o == nil || IsNil(o.RedisDatabase) {
		var ret int32
		return ret
	}
	return *o.RedisDatabase
}

// GetRedisDatabaseOk returns a tuple with the RedisDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetRedisDatabaseOk() (*int32, bool) {
	if o == nil || IsNil(o.RedisDatabase) {
		return nil, false
	}
	return o.RedisDatabase, true
}

// HasRedisDatabase returns a boolean if a field has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasRedisDatabase() bool {
	if o != nil && !IsNil(o.RedisDatabase) {
		return true
	}

	return false
}

// SetRedisDatabase gets a reference to the given int32 and assigns it to the RedisDatabase field.
func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetRedisDatabase(v int32) {
	o.RedisDatabase = &v
}

func (o UnibeeApiSystemAuthTokenGeneratorReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemAuthTokenGeneratorReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.PortalType) {
		toSerialize["portalType"] = o.PortalType
	}
	if !IsNil(o.RedisDatabase) {
		toSerialize["redisDatabase"] = o.RedisDatabase
	}
	return toSerialize, nil
}

func (o *UnibeeApiSystemAuthTokenGeneratorReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUnibeeApiSystemAuthTokenGeneratorReq := _UnibeeApiSystemAuthTokenGeneratorReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemAuthTokenGeneratorReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemAuthTokenGeneratorReq(varUnibeeApiSystemAuthTokenGeneratorReq)

	return err
}

type NullableUnibeeApiSystemAuthTokenGeneratorReq struct {
	value *UnibeeApiSystemAuthTokenGeneratorReq
	isSet bool
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorReq) Get() *UnibeeApiSystemAuthTokenGeneratorReq {
	return v.value
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorReq) Set(val *UnibeeApiSystemAuthTokenGeneratorReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemAuthTokenGeneratorReq(val *UnibeeApiSystemAuthTokenGeneratorReq) *NullableUnibeeApiSystemAuthTokenGeneratorReq {
	return &NullableUnibeeApiSystemAuthTokenGeneratorReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


