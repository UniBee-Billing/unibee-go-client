/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionOnetimeAddonNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionOnetimeAddonNewReq{}

// UnibeeApiMerchantSubscriptionOnetimeAddonNewReq Create payment for subscription onetime addon purchase
type UnibeeApiMerchantSubscriptionOnetimeAddonNewReq struct {
	// AddonId, id of one-time addon, the new payment will created base on the addon's amount'
	AddonId int64 `json:"addonId"`
	// Amount of discount
	DiscountAmount *int32 `json:"discountAmount,omitempty"`
	// DiscountCode
	DiscountCode *string `json:"discountCode,omitempty"`
	// Percentage of discount, 100=1%, ignore if discountAmount provide
	DiscountPercentage *int32 `json:"discountPercentage,omitempty"`
	// GatewayId, use user's gateway if not provide
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// Metadata，custom data
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Quantity, quantity of the new payment which one-time addon purchased
	Quantity int64 `json:"quantity"`
	// ReturnUrl, the addon's payment will redirect based on the returnUrl provided when it's back from gateway side
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// SubscriptionId, id of subscription which addon will attached, either SubscriptionId or UserId needed, The only one active subscription of userId will attach the addon
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// TaxPercentage，1000 = 10%, use subscription's taxPercentage if not provide
	TaxPercentage *int32 `json:"taxPercentage,omitempty"`
	// UserId, either SubscriptionId or UserId needed, The only one active subscription will update if userId provide instead of subscriptionId
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantSubscriptionOnetimeAddonNewReq UnibeeApiMerchantSubscriptionOnetimeAddonNewReq

// NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq(addonId int64, quantity int64) *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonNewReq{}
	this.AddonId = addonId
	this.Quantity = quantity
	return &this
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonNewReq{}
	return &this
}

// GetAddonId returns the AddonId field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetAddonId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AddonId
}

// GetAddonIdOk returns a tuple with the AddonId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetAddonIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddonId, true
}

// SetAddonId sets field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetAddonId(v int64) {
	o.AddonId = v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int32 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountAmount(v int32) {
	o.DiscountAmount = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountPercentage() int32 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int32
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int32 and assigns it to the DiscountPercentage field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountPercentage(v int32) {
	o.DiscountPercentage = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetQuantity returns the Quantity field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetQuantity(v int64) {
	o.Quantity = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetTaxPercentage() int32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetTaxPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int32 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetTaxPercentage(v int32) {
	o.TaxPercentage = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addonId"] = o.AddonId
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addonId",
		"quantity",
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

	varUnibeeApiMerchantSubscriptionOnetimeAddonNewReq := _UnibeeApiMerchantSubscriptionOnetimeAddonNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionOnetimeAddonNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionOnetimeAddonNewReq(varUnibeeApiMerchantSubscriptionOnetimeAddonNewReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq struct {
	value *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) Get() *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) Set(val *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq(val *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq {
	return &NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


