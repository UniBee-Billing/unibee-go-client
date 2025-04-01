/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthSessionReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthSessionReq{}

// UnibeeApiMerchantAuthSessionReq Session login
type UnibeeApiMerchantAuthSessionReq struct {
	// The session
	Session string `json:"session"`
}

type _UnibeeApiMerchantAuthSessionReq UnibeeApiMerchantAuthSessionReq

// NewUnibeeApiMerchantAuthSessionReq instantiates a new UnibeeApiMerchantAuthSessionReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthSessionReq(session string) *UnibeeApiMerchantAuthSessionReq {
	this := UnibeeApiMerchantAuthSessionReq{}
	this.Session = session
	return &this
}

// NewUnibeeApiMerchantAuthSessionReqWithDefaults instantiates a new UnibeeApiMerchantAuthSessionReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthSessionReqWithDefaults() *UnibeeApiMerchantAuthSessionReq {
	this := UnibeeApiMerchantAuthSessionReq{}
	return &this
}

// GetSession returns the Session field value
func (o *UnibeeApiMerchantAuthSessionReq) GetSession() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthSessionReq) GetSessionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *UnibeeApiMerchantAuthSessionReq) SetSession(v string) {
	o.Session = v
}

func (o UnibeeApiMerchantAuthSessionReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthSessionReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthSessionReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session",
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

	varUnibeeApiMerchantAuthSessionReq := _UnibeeApiMerchantAuthSessionReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthSessionReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthSessionReq(varUnibeeApiMerchantAuthSessionReq)

	return err
}

type NullableUnibeeApiMerchantAuthSessionReq struct {
	value *UnibeeApiMerchantAuthSessionReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthSessionReq) Get() *UnibeeApiMerchantAuthSessionReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthSessionReq) Set(val *UnibeeApiMerchantAuthSessionReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthSessionReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthSessionReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthSessionReq(val *UnibeeApiMerchantAuthSessionReq) *NullableUnibeeApiMerchantAuthSessionReq {
	return &NullableUnibeeApiMerchantAuthSessionReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthSessionReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthSessionReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


