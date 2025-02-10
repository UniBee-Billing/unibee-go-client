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

// checks if the UnibeeApiMerchantSubscriptionAdminNoteListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionAdminNoteListReq{}

// UnibeeApiMerchantSubscriptionAdminNoteListReq struct for UnibeeApiMerchantSubscriptionAdminNoteListReq
type UnibeeApiMerchantSubscriptionAdminNoteListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionAdminNoteListReq UnibeeApiMerchantSubscriptionAdminNoteListReq

// NewUnibeeApiMerchantSubscriptionAdminNoteListReq instantiates a new UnibeeApiMerchantSubscriptionAdminNoteListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionAdminNoteListReq(subscriptionId string) *UnibeeApiMerchantSubscriptionAdminNoteListReq {
	this := UnibeeApiMerchantSubscriptionAdminNoteListReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionAdminNoteListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionAdminNoteListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionAdminNoteListReqWithDefaults() *UnibeeApiMerchantSubscriptionAdminNoteListReq {
	this := UnibeeApiMerchantSubscriptionAdminNoteListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionAdminNoteListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionAdminNoteListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionAdminNoteListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionAdminNoteListReq := _UnibeeApiMerchantSubscriptionAdminNoteListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionAdminNoteListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionAdminNoteListReq(varUnibeeApiMerchantSubscriptionAdminNoteListReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionAdminNoteListReq struct {
	value *UnibeeApiMerchantSubscriptionAdminNoteListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) Get() *UnibeeApiMerchantSubscriptionAdminNoteListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) Set(val *UnibeeApiMerchantSubscriptionAdminNoteListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionAdminNoteListReq(val *UnibeeApiMerchantSubscriptionAdminNoteListReq) *NullableUnibeeApiMerchantSubscriptionAdminNoteListReq {
	return &NullableUnibeeApiMerchantSubscriptionAdminNoteListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionAdminNoteListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


