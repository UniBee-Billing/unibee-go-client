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

// checks if the UnibeeApiMerchantTrackSetupSegmentReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTrackSetupSegmentReq{}

// UnibeeApiMerchantTrackSetupSegmentReq struct for UnibeeApiMerchantTrackSetupSegmentReq
type UnibeeApiMerchantTrackSetupSegmentReq struct {
	// ServerSideSecret
	ServerSideSecret string `json:"serverSideSecret"`
	// UserPortalSecret
	UserPortalSecret string `json:"userPortalSecret"`
}

type _UnibeeApiMerchantTrackSetupSegmentReq UnibeeApiMerchantTrackSetupSegmentReq

// NewUnibeeApiMerchantTrackSetupSegmentReq instantiates a new UnibeeApiMerchantTrackSetupSegmentReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTrackSetupSegmentReq(serverSideSecret string, userPortalSecret string) *UnibeeApiMerchantTrackSetupSegmentReq {
	this := UnibeeApiMerchantTrackSetupSegmentReq{}
	this.ServerSideSecret = serverSideSecret
	this.UserPortalSecret = userPortalSecret
	return &this
}

// NewUnibeeApiMerchantTrackSetupSegmentReqWithDefaults instantiates a new UnibeeApiMerchantTrackSetupSegmentReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTrackSetupSegmentReqWithDefaults() *UnibeeApiMerchantTrackSetupSegmentReq {
	this := UnibeeApiMerchantTrackSetupSegmentReq{}
	return &this
}

// GetServerSideSecret returns the ServerSideSecret field value
func (o *UnibeeApiMerchantTrackSetupSegmentReq) GetServerSideSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerSideSecret
}

// GetServerSideSecretOk returns a tuple with the ServerSideSecret field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTrackSetupSegmentReq) GetServerSideSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerSideSecret, true
}

// SetServerSideSecret sets field value
func (o *UnibeeApiMerchantTrackSetupSegmentReq) SetServerSideSecret(v string) {
	o.ServerSideSecret = v
}

// GetUserPortalSecret returns the UserPortalSecret field value
func (o *UnibeeApiMerchantTrackSetupSegmentReq) GetUserPortalSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPortalSecret
}

// GetUserPortalSecretOk returns a tuple with the UserPortalSecret field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTrackSetupSegmentReq) GetUserPortalSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPortalSecret, true
}

// SetUserPortalSecret sets field value
func (o *UnibeeApiMerchantTrackSetupSegmentReq) SetUserPortalSecret(v string) {
	o.UserPortalSecret = v
}

func (o UnibeeApiMerchantTrackSetupSegmentReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTrackSetupSegmentReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverSideSecret"] = o.ServerSideSecret
	toSerialize["userPortalSecret"] = o.UserPortalSecret
	return toSerialize, nil
}

func (o *UnibeeApiMerchantTrackSetupSegmentReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverSideSecret",
		"userPortalSecret",
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

	varUnibeeApiMerchantTrackSetupSegmentReq := _UnibeeApiMerchantTrackSetupSegmentReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantTrackSetupSegmentReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantTrackSetupSegmentReq(varUnibeeApiMerchantTrackSetupSegmentReq)

	return err
}

type NullableUnibeeApiMerchantTrackSetupSegmentReq struct {
	value *UnibeeApiMerchantTrackSetupSegmentReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTrackSetupSegmentReq) Get() *UnibeeApiMerchantTrackSetupSegmentReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTrackSetupSegmentReq) Set(val *UnibeeApiMerchantTrackSetupSegmentReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTrackSetupSegmentReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTrackSetupSegmentReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTrackSetupSegmentReq(val *UnibeeApiMerchantTrackSetupSegmentReq) *NullableUnibeeApiMerchantTrackSetupSegmentReq {
	return &NullableUnibeeApiMerchantTrackSetupSegmentReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTrackSetupSegmentReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTrackSetupSegmentReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


