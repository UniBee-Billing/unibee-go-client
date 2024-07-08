/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentMethodNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodNewReq{}

// UnibeeApiMerchantPaymentMethodNewReq struct for UnibeeApiMerchantPaymentMethodNewReq
type UnibeeApiMerchantPaymentMethodNewReq struct {
	// The currency of payment method
	Currency *string `json:"currency,omitempty"`
	// The unique id of gateway
	GatewayId int64 `json:"gatewayId"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The redirect url when method created return back
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The id of subscription that want to attach
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	Type *string `json:"type,omitempty"`
	// The customer's unique id
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantPaymentMethodNewReq UnibeeApiMerchantPaymentMethodNewReq

// NewUnibeeApiMerchantPaymentMethodNewReq instantiates a new UnibeeApiMerchantPaymentMethodNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodNewReq(gatewayId int64, userId int64) *UnibeeApiMerchantPaymentMethodNewReq {
	this := UnibeeApiMerchantPaymentMethodNewReq{}
	this.GatewayId = gatewayId
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantPaymentMethodNewReqWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodNewReqWithDefaults() *UnibeeApiMerchantPaymentMethodNewReq {
	this := UnibeeApiMerchantPaymentMethodNewReq{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetType(v string) {
	o.Type = &v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodNewReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantPaymentMethodNewReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantPaymentMethodNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentMethodNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
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

	varUnibeeApiMerchantPaymentMethodNewReq := _UnibeeApiMerchantPaymentMethodNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentMethodNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentMethodNewReq(varUnibeeApiMerchantPaymentMethodNewReq)

	return err
}

type NullableUnibeeApiMerchantPaymentMethodNewReq struct {
	value *UnibeeApiMerchantPaymentMethodNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodNewReq) Get() *UnibeeApiMerchantPaymentMethodNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewReq) Set(val *UnibeeApiMerchantPaymentMethodNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodNewReq(val *UnibeeApiMerchantPaymentMethodNewReq) *NullableUnibeeApiMerchantPaymentMethodNewReq {
	return &NullableUnibeeApiMerchantPaymentMethodNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


