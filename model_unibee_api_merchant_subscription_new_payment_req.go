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

// checks if the UnibeeApiMerchantSubscriptionNewPaymentReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionNewPaymentReq{}

// UnibeeApiMerchantSubscriptionNewPaymentReq struct for UnibeeApiMerchantSubscriptionNewPaymentReq
type UnibeeApiMerchantSubscriptionNewPaymentReq struct {
	// CancelUrl
	CancelUrl *string `json:"cancelUrl,omitempty"`
	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`
	// Currency, either Currency&TotalAmount or PlanId needed
	Currency *string `json:"currency,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Email, either ExternalUserId&Email or UserId needed
	Email *string `json:"email,omitempty"`
	// ExternalPaymentId should unique for payment
	ExternalPaymentId *string `json:"externalPaymentId,omitempty"`
	// ExternalUserId, unique, either ExternalUserId&Email or UserId needed
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// who pay the gas, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	// GatewayId
	GatewayId int64 `json:"gatewayId"`
	// Items
	Items []UnibeeApiMerchantPaymentItem `json:"items,omitempty"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// PlanId, either TotalAmount&Currency or PlanId needed
	PlanId *int64 `json:"planId,omitempty"`
	// Redirect Url
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// Total PaymentAmount, Cent, either TotalAmount&Currency or PlanId needed
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// UserId, either ExternalUserId&Email or UserId needed
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantSubscriptionNewPaymentReq UnibeeApiMerchantSubscriptionNewPaymentReq

// NewUnibeeApiMerchantSubscriptionNewPaymentReq instantiates a new UnibeeApiMerchantSubscriptionNewPaymentReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionNewPaymentReq(gatewayId int64) *UnibeeApiMerchantSubscriptionNewPaymentReq {
	this := UnibeeApiMerchantSubscriptionNewPaymentReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantSubscriptionNewPaymentReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionNewPaymentReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionNewPaymentReqWithDefaults() *UnibeeApiMerchantSubscriptionNewPaymentReq {
	this := UnibeeApiMerchantSubscriptionNewPaymentReq{}
	return &this
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetEmail(v string) {
	o.Email = &v
}

// GetExternalPaymentId returns the ExternalPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalPaymentId() string {
	if o == nil || IsNil(o.ExternalPaymentId) {
		var ret string
		return ret
	}
	return *o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPaymentId) {
		return nil, false
	}
	return o.ExternalPaymentId, true
}

// HasExternalPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasExternalPaymentId() bool {
	if o != nil && !IsNil(o.ExternalPaymentId) {
		return true
	}

	return false
}

// SetExternalPaymentId gets a reference to the given string and assigns it to the ExternalPaymentId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetItems() []UnibeeApiMerchantPaymentItem {
	if o == nil || IsNil(o.Items) {
		var ret []UnibeeApiMerchantPaymentItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetItemsOk() ([]UnibeeApiMerchantPaymentItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UnibeeApiMerchantPaymentItem and assigns it to the Items field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetItems(v []UnibeeApiMerchantPaymentItem) {
	o.Items = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetName(v string) {
	o.Name = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetPlanId(v int64) {
	o.PlanId = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionNewPaymentReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionNewPaymentReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalPaymentId) {
		toSerialize["externalPaymentId"] = o.ExternalPaymentId
	}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.GasPayer) {
		toSerialize["gasPayer"] = o.GasPayer
	}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
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

	varUnibeeApiMerchantSubscriptionNewPaymentReq := _UnibeeApiMerchantSubscriptionNewPaymentReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionNewPaymentReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionNewPaymentReq(varUnibeeApiMerchantSubscriptionNewPaymentReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionNewPaymentReq struct {
	value *UnibeeApiMerchantSubscriptionNewPaymentReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentReq) Get() *UnibeeApiMerchantSubscriptionNewPaymentReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentReq) Set(val *UnibeeApiMerchantSubscriptionNewPaymentReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionNewPaymentReq(val *UnibeeApiMerchantSubscriptionNewPaymentReq) *NullableUnibeeApiMerchantSubscriptionNewPaymentReq {
	return &NullableUnibeeApiMerchantSubscriptionNewPaymentReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


