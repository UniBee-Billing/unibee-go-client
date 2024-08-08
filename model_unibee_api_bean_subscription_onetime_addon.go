/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanSubscriptionOnetimeAddon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanSubscriptionOnetimeAddon{}

// UnibeeApiBeanSubscriptionOnetimeAddon struct for UnibeeApiBeanSubscriptionOnetimeAddon
type UnibeeApiBeanSubscriptionOnetimeAddon struct {
	// onetime addonId
	AddonId *int64 `json:"addonId,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// 0-UnDeleted，1-Deleted
	IsDeleted *int32 `json:"isDeleted,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// PaymentId
	PaymentId *string `json:"paymentId,omitempty"`
	// PaymentLink
	PaymentLink *string `json:"paymentLink,omitempty"`
	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// status, 1-create, 2-paid, 3-cancel, 4-expired
	Status *int32 `json:"status,omitempty"`
	// subscription_id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewUnibeeApiBeanSubscriptionOnetimeAddon instantiates a new UnibeeApiBeanSubscriptionOnetimeAddon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanSubscriptionOnetimeAddon() *UnibeeApiBeanSubscriptionOnetimeAddon {
	this := UnibeeApiBeanSubscriptionOnetimeAddon{}
	return &this
}

// NewUnibeeApiBeanSubscriptionOnetimeAddonWithDefaults instantiates a new UnibeeApiBeanSubscriptionOnetimeAddon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanSubscriptionOnetimeAddonWithDefaults() *UnibeeApiBeanSubscriptionOnetimeAddon {
	this := UnibeeApiBeanSubscriptionOnetimeAddon{}
	return &this
}

// GetAddonId returns the AddonId field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetAddonId() int64 {
	if o == nil || IsNil(o.AddonId) {
		var ret int64
		return ret
	}
	return *o.AddonId
}

// GetAddonIdOk returns a tuple with the AddonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetAddonIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AddonId) {
		return nil, false
	}
	return o.AddonId, true
}

// HasAddonId returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasAddonId() bool {
	if o != nil && !IsNil(o.AddonId) {
		return true
	}

	return false
}

// SetAddonId gets a reference to the given int64 and assigns it to the AddonId field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetAddonId(v int64) {
	o.AddonId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetId(v int64) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetIsDeleted() int32 {
	if o == nil || IsNil(o.IsDeleted) {
		var ret int32
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetIsDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given int32 and assigns it to the IsDeleted field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetIsDeleted(v int32) {
	o.IsDeleted = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetPaymentLink returns the PaymentLink field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetPaymentLink() string {
	if o == nil || IsNil(o.PaymentLink) {
		var ret string
		return ret
	}
	return *o.PaymentLink
}

// GetPaymentLinkOk returns a tuple with the PaymentLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetPaymentLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentLink) {
		return nil, false
	}
	return o.PaymentLink, true
}

// HasPaymentLink returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasPaymentLink() bool {
	if o != nil && !IsNil(o.PaymentLink) {
		return true
	}

	return false
}

// SetPaymentLink gets a reference to the given string and assigns it to the PaymentLink field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetPaymentLink(v string) {
	o.PaymentLink = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanSubscriptionOnetimeAddon) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o UnibeeApiBeanSubscriptionOnetimeAddon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanSubscriptionOnetimeAddon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonId) {
		toSerialize["addonId"] = o.AddonId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.PaymentLink) {
		toSerialize["paymentLink"] = o.PaymentLink
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanSubscriptionOnetimeAddon struct {
	value *UnibeeApiBeanSubscriptionOnetimeAddon
	isSet bool
}

func (v NullableUnibeeApiBeanSubscriptionOnetimeAddon) Get() *UnibeeApiBeanSubscriptionOnetimeAddon {
	return v.value
}

func (v *NullableUnibeeApiBeanSubscriptionOnetimeAddon) Set(val *UnibeeApiBeanSubscriptionOnetimeAddon) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanSubscriptionOnetimeAddon) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanSubscriptionOnetimeAddon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanSubscriptionOnetimeAddon(val *UnibeeApiBeanSubscriptionOnetimeAddon) *NullableUnibeeApiBeanSubscriptionOnetimeAddon {
	return &NullableUnibeeApiBeanSubscriptionOnetimeAddon{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanSubscriptionOnetimeAddon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanSubscriptionOnetimeAddon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


