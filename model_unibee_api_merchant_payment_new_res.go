/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentNewRes{}

// UnibeeApiMerchantPaymentNewRes struct for UnibeeApiMerchantPaymentNewRes
type UnibeeApiMerchantPaymentNewRes struct {
	Action map[string]interface{} `json:"action,omitempty"`
	// The external unique id of payment
	ExternalPaymentId *string `json:"externalPaymentId,omitempty"`
	Link *string `json:"link,omitempty"`
	// The unique id of payment
	PaymentId *string `json:"paymentId,omitempty"`
	// Status, 10-Created|20-Success|30-Failed|40-Cancelled
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiMerchantPaymentNewRes instantiates a new UnibeeApiMerchantPaymentNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentNewRes() *UnibeeApiMerchantPaymentNewRes {
	this := UnibeeApiMerchantPaymentNewRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentNewResWithDefaults instantiates a new UnibeeApiMerchantPaymentNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentNewResWithDefaults() *UnibeeApiMerchantPaymentNewRes {
	this := UnibeeApiMerchantPaymentNewRes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewRes) GetAction() map[string]interface{} {
	if o == nil || IsNil(o.Action) {
		var ret map[string]interface{}
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewRes) GetActionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Action) {
		return map[string]interface{}{}, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewRes) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given map[string]interface{} and assigns it to the Action field.
func (o *UnibeeApiMerchantPaymentNewRes) SetAction(v map[string]interface{}) {
	o.Action = v
}

// GetExternalPaymentId returns the ExternalPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewRes) GetExternalPaymentId() string {
	if o == nil || IsNil(o.ExternalPaymentId) {
		var ret string
		return ret
	}
	return *o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewRes) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPaymentId) {
		return nil, false
	}
	return o.ExternalPaymentId, true
}

// HasExternalPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewRes) HasExternalPaymentId() bool {
	if o != nil && !IsNil(o.ExternalPaymentId) {
		return true
	}

	return false
}

// SetExternalPaymentId gets a reference to the given string and assigns it to the ExternalPaymentId field.
func (o *UnibeeApiMerchantPaymentNewRes) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantPaymentNewRes) SetLink(v string) {
	o.Link = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewRes) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewRes) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewRes) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiMerchantPaymentNewRes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewRes) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewRes) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewRes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantPaymentNewRes) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiMerchantPaymentNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ExternalPaymentId) {
		toSerialize["externalPaymentId"] = o.ExternalPaymentId
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentNewRes struct {
	value *UnibeeApiMerchantPaymentNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentNewRes) Get() *UnibeeApiMerchantPaymentNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentNewRes) Set(val *UnibeeApiMerchantPaymentNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentNewRes(val *UnibeeApiMerchantPaymentNewRes) *NullableUnibeeApiMerchantPaymentNewRes {
	return &NullableUnibeeApiMerchantPaymentNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


