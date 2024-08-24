/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPaymentNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentNewReq{}

// UnibeeApiMerchantPaymentNewReq struct for UnibeeApiMerchantPaymentNewReq
type UnibeeApiMerchantPaymentNewReq struct {
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
	// Whether sen invoice email to customer or not，default false
	SendInvoice *bool `json:"sendInvoice,omitempty"`
	// Total PaymentAmount, Cent, either TotalAmount&Currency or PlanId needed
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// UserId, either ExternalUserId&Email or UserId needed
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantPaymentNewReq UnibeeApiMerchantPaymentNewReq

// NewUnibeeApiMerchantPaymentNewReq instantiates a new UnibeeApiMerchantPaymentNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentNewReq(gatewayId int64) *UnibeeApiMerchantPaymentNewReq {
	this := UnibeeApiMerchantPaymentNewReq{}
	this.GatewayId = gatewayId
	var sendInvoice bool = false
	this.SendInvoice = &sendInvoice
	return &this
}

// NewUnibeeApiMerchantPaymentNewReqWithDefaults instantiates a new UnibeeApiMerchantPaymentNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentNewReqWithDefaults() *UnibeeApiMerchantPaymentNewReq {
	this := UnibeeApiMerchantPaymentNewReq{}
	var sendInvoice bool = false
	this.SendInvoice = &sendInvoice
	return &this
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *UnibeeApiMerchantPaymentNewReq) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiMerchantPaymentNewReq) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPaymentNewReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantPaymentNewReq) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantPaymentNewReq) SetEmail(v string) {
	o.Email = &v
}

// GetExternalPaymentId returns the ExternalPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentId() string {
	if o == nil || IsNil(o.ExternalPaymentId) {
		var ret string
		return ret
	}
	return *o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPaymentId) {
		return nil, false
	}
	return o.ExternalPaymentId, true
}

// HasExternalPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasExternalPaymentId() bool {
	if o != nil && !IsNil(o.ExternalPaymentId) {
		return true
	}

	return false
}

// SetExternalPaymentId gets a reference to the given string and assigns it to the ExternalPaymentId field.
func (o *UnibeeApiMerchantPaymentNewReq) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantPaymentNewReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiMerchantPaymentNewReq) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetItems() []UnibeeApiMerchantPaymentItem {
	if o == nil || IsNil(o.Items) {
		var ret []UnibeeApiMerchantPaymentItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetItemsOk() ([]UnibeeApiMerchantPaymentItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UnibeeApiMerchantPaymentItem and assigns it to the Items field.
func (o *UnibeeApiMerchantPaymentNewReq) SetItems(v []UnibeeApiMerchantPaymentItem) {
	o.Items = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPaymentNewReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantPaymentNewReq) SetName(v string) {
	o.Name = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiMerchantPaymentNewReq) SetPlanId(v int64) {
	o.PlanId = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UnibeeApiMerchantPaymentNewReq) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSendInvoice returns the SendInvoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetSendInvoice() bool {
	if o == nil || IsNil(o.SendInvoice) {
		var ret bool
		return ret
	}
	return *o.SendInvoice
}

// GetSendInvoiceOk returns a tuple with the SendInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetSendInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.SendInvoice) {
		return nil, false
	}
	return o.SendInvoice, true
}

// HasSendInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasSendInvoice() bool {
	if o != nil && !IsNil(o.SendInvoice) {
		return true
	}

	return false
}

// SetSendInvoice gets a reference to the given bool and assigns it to the SendInvoice field.
func (o *UnibeeApiMerchantPaymentNewReq) SetSendInvoice(v bool) {
	o.SendInvoice = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiMerchantPaymentNewReq) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantPaymentNewReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantPaymentNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentNewReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SendInvoice) {
		toSerialize["sendInvoice"] = o.SendInvoice
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentNewReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPaymentNewReq := _UnibeeApiMerchantPaymentNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentNewReq(varUnibeeApiMerchantPaymentNewReq)

	return err
}

type NullableUnibeeApiMerchantPaymentNewReq struct {
	value *UnibeeApiMerchantPaymentNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentNewReq) Get() *UnibeeApiMerchantPaymentNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) Set(val *UnibeeApiMerchantPaymentNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentNewReq(val *UnibeeApiMerchantPaymentNewReq) *NullableUnibeeApiMerchantPaymentNewReq {
	return &NullableUnibeeApiMerchantPaymentNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


