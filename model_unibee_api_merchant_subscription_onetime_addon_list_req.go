/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionOnetimeAddonListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionOnetimeAddonListReq{}

// UnibeeApiMerchantSubscriptionOnetimeAddonListReq struct for UnibeeApiMerchantSubscriptionOnetimeAddonListReq
type UnibeeApiMerchantSubscriptionOnetimeAddonListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
	// UserId
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantSubscriptionOnetimeAddonListReq UnibeeApiMerchantSubscriptionOnetimeAddonListReq

// NewUnibeeApiMerchantSubscriptionOnetimeAddonListReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionOnetimeAddonListReq(userId int64) *UnibeeApiMerchantSubscriptionOnetimeAddonListReq {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonListReq{}
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionOnetimeAddonListReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonListReq {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetPage(v int32) {
	o.Page = &v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varUnibeeApiMerchantSubscriptionOnetimeAddonListReq := _UnibeeApiMerchantSubscriptionOnetimeAddonListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionOnetimeAddonListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionOnetimeAddonListReq(varUnibeeApiMerchantSubscriptionOnetimeAddonListReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq struct {
	value *UnibeeApiMerchantSubscriptionOnetimeAddonListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) Get() *UnibeeApiMerchantSubscriptionOnetimeAddonListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) Set(val *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq(val *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq {
	return &NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


