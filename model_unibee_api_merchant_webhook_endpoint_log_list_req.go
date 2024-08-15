/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantWebhookEndpointLogListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookEndpointLogListReq{}

// UnibeeApiMerchantWebhookEndpointLogListReq struct for UnibeeApiMerchantWebhookEndpointLogListReq
type UnibeeApiMerchantWebhookEndpointLogListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// EndpointId
	EndpointId int64 `json:"endpointId"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
}

type _UnibeeApiMerchantWebhookEndpointLogListReq UnibeeApiMerchantWebhookEndpointLogListReq

// NewUnibeeApiMerchantWebhookEndpointLogListReq instantiates a new UnibeeApiMerchantWebhookEndpointLogListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookEndpointLogListReq(endpointId int64) *UnibeeApiMerchantWebhookEndpointLogListReq {
	this := UnibeeApiMerchantWebhookEndpointLogListReq{}
	this.EndpointId = endpointId
	return &this
}

// NewUnibeeApiMerchantWebhookEndpointLogListReqWithDefaults instantiates a new UnibeeApiMerchantWebhookEndpointLogListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookEndpointLogListReqWithDefaults() *UnibeeApiMerchantWebhookEndpointLogListReq {
	this := UnibeeApiMerchantWebhookEndpointLogListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) SetCount(v int32) {
	o.Count = &v
}

// GetEndpointId returns the EndpointId field value
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetEndpointId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetEndpointIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) SetEndpointId(v int64) {
	o.EndpointId = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantWebhookEndpointLogListReq) SetPage(v int32) {
	o.Page = &v
}

func (o UnibeeApiMerchantWebhookEndpointLogListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookEndpointLogListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["endpointId"] = o.EndpointId
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantWebhookEndpointLogListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpointId",
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

	varUnibeeApiMerchantWebhookEndpointLogListReq := _UnibeeApiMerchantWebhookEndpointLogListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantWebhookEndpointLogListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantWebhookEndpointLogListReq(varUnibeeApiMerchantWebhookEndpointLogListReq)

	return err
}

type NullableUnibeeApiMerchantWebhookEndpointLogListReq struct {
	value *UnibeeApiMerchantWebhookEndpointLogListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListReq) Get() *UnibeeApiMerchantWebhookEndpointLogListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListReq) Set(val *UnibeeApiMerchantWebhookEndpointLogListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookEndpointLogListReq(val *UnibeeApiMerchantWebhookEndpointLogListReq) *NullableUnibeeApiMerchantWebhookEndpointLogListReq {
	return &NullableUnibeeApiMerchantWebhookEndpointLogListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


