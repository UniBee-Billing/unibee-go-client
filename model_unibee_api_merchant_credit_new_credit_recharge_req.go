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

// checks if the UnibeeApiMerchantCreditNewCreditRechargeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditNewCreditRechargeReq{}

// UnibeeApiMerchantCreditNewCreditRechargeReq New Credit Recharge
type UnibeeApiMerchantCreditNewCreditRechargeReq struct {
	// CancelUrl
	CancelUrl *string `json:"cancelUrl,omitempty"`
	// id of credit account, either userId&currency or creditAccountId 
	CreditAccountId *int64 `json:"creditAccountId,omitempty"`
	// currency of recharge
	Currency *string `json:"currency,omitempty"`
	// recharge description
	Description *string `json:"description,omitempty"`
	GatewayId int64 `json:"gatewayId"`
	// recharge name
	Name *string `json:"name,omitempty"`
	RechargeAmount int64 `json:"rechargeAmount"`
	// ReturnUrl
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// id of user to recharge, either userId&currency or creditAccountId 
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantCreditNewCreditRechargeReq UnibeeApiMerchantCreditNewCreditRechargeReq

// NewUnibeeApiMerchantCreditNewCreditRechargeReq instantiates a new UnibeeApiMerchantCreditNewCreditRechargeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditNewCreditRechargeReq(gatewayId int64, rechargeAmount int64) *UnibeeApiMerchantCreditNewCreditRechargeReq {
	this := UnibeeApiMerchantCreditNewCreditRechargeReq{}
	this.GatewayId = gatewayId
	this.RechargeAmount = rechargeAmount
	return &this
}

// NewUnibeeApiMerchantCreditNewCreditRechargeReqWithDefaults instantiates a new UnibeeApiMerchantCreditNewCreditRechargeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditNewCreditRechargeReqWithDefaults() *UnibeeApiMerchantCreditNewCreditRechargeReq {
	this := UnibeeApiMerchantCreditNewCreditRechargeReq{}
	return &this
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetCreditAccountId returns the CreditAccountId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCreditAccountId() int64 {
	if o == nil || IsNil(o.CreditAccountId) {
		var ret int64
		return ret
	}
	return *o.CreditAccountId
}

// GetCreditAccountIdOk returns a tuple with the CreditAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCreditAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CreditAccountId) {
		return nil, false
	}
	return o.CreditAccountId, true
}

// HasCreditAccountId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCreditAccountId() bool {
	if o != nil && !IsNil(o.CreditAccountId) {
		return true
	}

	return false
}

// SetCreditAccountId gets a reference to the given int64 and assigns it to the CreditAccountId field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCreditAccountId(v int64) {
	o.CreditAccountId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetDescription(v string) {
	o.Description = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetName(v string) {
	o.Name = &v
}

// GetRechargeAmount returns the RechargeAmount field value
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetRechargeAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RechargeAmount
}

// GetRechargeAmountOk returns a tuple with the RechargeAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetRechargeAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RechargeAmount, true
}

// SetRechargeAmount sets field value
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetRechargeAmount(v int64) {
	o.RechargeAmount = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantCreditNewCreditRechargeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditNewCreditRechargeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if !IsNil(o.CreditAccountId) {
		toSerialize["creditAccountId"] = o.CreditAccountId
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["rechargeAmount"] = o.RechargeAmount
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantCreditNewCreditRechargeReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
		"rechargeAmount",
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

	varUnibeeApiMerchantCreditNewCreditRechargeReq := _UnibeeApiMerchantCreditNewCreditRechargeReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantCreditNewCreditRechargeReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantCreditNewCreditRechargeReq(varUnibeeApiMerchantCreditNewCreditRechargeReq)

	return err
}

type NullableUnibeeApiMerchantCreditNewCreditRechargeReq struct {
	value *UnibeeApiMerchantCreditNewCreditRechargeReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeReq) Get() *UnibeeApiMerchantCreditNewCreditRechargeReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeReq) Set(val *UnibeeApiMerchantCreditNewCreditRechargeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditNewCreditRechargeReq(val *UnibeeApiMerchantCreditNewCreditRechargeReq) *NullableUnibeeApiMerchantCreditNewCreditRechargeReq {
	return &NullableUnibeeApiMerchantCreditNewCreditRechargeReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditNewCreditRechargeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditNewCreditRechargeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


