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

// checks if the UnibeeApiMerchantGatewayEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayEditReq{}

// UnibeeApiMerchantGatewayEditReq Edit the exist payment gateway
type UnibeeApiMerchantGatewayEditReq struct {
	// The currency exchange for gateway payment, effect at start of payment creation when currency matched
	CurrencyExchange []UnibeeApiBeanDetailGatewayCurrencyExchange `json:"currencyExchange,omitempty"`
	// The displayName of payment gateway
	DisplayName *string `json:"displayName,omitempty"`
	// The id of payment gateway
	GatewayId int64 `json:"gatewayId"`
	// The key of payment gateway
	GatewayKey *string `json:"gatewayKey,omitempty"`
	// The logo of payment gateway
	GatewayLogo [][]string `json:"gatewayLogo,omitempty"`
	// Selected gateway payment types
	GatewayPaymentTypes []string `json:"gatewayPaymentTypes,omitempty"`
	// The secret of payment gateway
	GatewaySecret *string `json:"gatewaySecret,omitempty"`
	// The sort value of payment gateway, The bigger, the closer to the front
	Sort *int32 `json:"sort,omitempty"`
	// The sub gateway of payment gateway
	SubGateway *string `json:"subGateway,omitempty"`
}

type _UnibeeApiMerchantGatewayEditReq UnibeeApiMerchantGatewayEditReq

// NewUnibeeApiMerchantGatewayEditReq instantiates a new UnibeeApiMerchantGatewayEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayEditReq(gatewayId int64) *UnibeeApiMerchantGatewayEditReq {
	this := UnibeeApiMerchantGatewayEditReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantGatewayEditReqWithDefaults instantiates a new UnibeeApiMerchantGatewayEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayEditReqWithDefaults() *UnibeeApiMerchantGatewayEditReq {
	this := UnibeeApiMerchantGatewayEditReq{}
	return &this
}

// GetCurrencyExchange returns the CurrencyExchange field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetCurrencyExchange() []UnibeeApiBeanDetailGatewayCurrencyExchange {
	if o == nil || IsNil(o.CurrencyExchange) {
		var ret []UnibeeApiBeanDetailGatewayCurrencyExchange
		return ret
	}
	return o.CurrencyExchange
}

// GetCurrencyExchangeOk returns a tuple with the CurrencyExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetCurrencyExchangeOk() ([]UnibeeApiBeanDetailGatewayCurrencyExchange, bool) {
	if o == nil || IsNil(o.CurrencyExchange) {
		return nil, false
	}
	return o.CurrencyExchange, true
}

// HasCurrencyExchange returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasCurrencyExchange() bool {
	if o != nil && !IsNil(o.CurrencyExchange) {
		return true
	}

	return false
}

// SetCurrencyExchange gets a reference to the given []UnibeeApiBeanDetailGatewayCurrencyExchange and assigns it to the CurrencyExchange field.
func (o *UnibeeApiMerchantGatewayEditReq) SetCurrencyExchange(v []UnibeeApiBeanDetailGatewayCurrencyExchange) {
	o.CurrencyExchange = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UnibeeApiMerchantGatewayEditReq) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetGatewayLogo returns the GatewayLogo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayLogo() [][]string {
	if o == nil || IsNil(o.GatewayLogo) {
		var ret [][]string
		return ret
	}
	return o.GatewayLogo
}

// GetGatewayLogoOk returns a tuple with the GatewayLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayLogoOk() ([][]string, bool) {
	if o == nil || IsNil(o.GatewayLogo) {
		return nil, false
	}
	return o.GatewayLogo, true
}

// HasGatewayLogo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayLogo() bool {
	if o != nil && !IsNil(o.GatewayLogo) {
		return true
	}

	return false
}

// SetGatewayLogo gets a reference to the given [][]string and assigns it to the GatewayLogo field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayLogo(v [][]string) {
	o.GatewayLogo = v
}

// GetGatewayPaymentTypes returns the GatewayPaymentTypes field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayPaymentTypes() []string {
	if o == nil || IsNil(o.GatewayPaymentTypes) {
		var ret []string
		return ret
	}
	return o.GatewayPaymentTypes
}

// GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayPaymentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.GatewayPaymentTypes) {
		return nil, false
	}
	return o.GatewayPaymentTypes, true
}

// HasGatewayPaymentTypes returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayPaymentTypes() bool {
	if o != nil && !IsNil(o.GatewayPaymentTypes) {
		return true
	}

	return false
}

// SetGatewayPaymentTypes gets a reference to the given []string and assigns it to the GatewayPaymentTypes field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayPaymentTypes(v []string) {
	o.GatewayPaymentTypes = v
}

// GetGatewaySecret returns the GatewaySecret field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecret() string {
	if o == nil || IsNil(o.GatewaySecret) {
		var ret string
		return ret
	}
	return *o.GatewaySecret
}

// GetGatewaySecretOk returns a tuple with the GatewaySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecretOk() (*string, bool) {
	if o == nil || IsNil(o.GatewaySecret) {
		return nil, false
	}
	return o.GatewaySecret, true
}

// HasGatewaySecret returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasGatewaySecret() bool {
	if o != nil && !IsNil(o.GatewaySecret) {
		return true
	}

	return false
}

// SetGatewaySecret gets a reference to the given string and assigns it to the GatewaySecret field.
func (o *UnibeeApiMerchantGatewayEditReq) SetGatewaySecret(v string) {
	o.GatewaySecret = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetSort() int32 {
	if o == nil || IsNil(o.Sort) {
		var ret int32
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetSortOk() (*int32, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int32 and assigns it to the Sort field.
func (o *UnibeeApiMerchantGatewayEditReq) SetSort(v int32) {
	o.Sort = &v
}

// GetSubGateway returns the SubGateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayEditReq) GetSubGateway() string {
	if o == nil || IsNil(o.SubGateway) {
		var ret string
		return ret
	}
	return *o.SubGateway
}

// GetSubGatewayOk returns a tuple with the SubGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayEditReq) GetSubGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.SubGateway) {
		return nil, false
	}
	return o.SubGateway, true
}

// HasSubGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayEditReq) HasSubGateway() bool {
	if o != nil && !IsNil(o.SubGateway) {
		return true
	}

	return false
}

// SetSubGateway gets a reference to the given string and assigns it to the SubGateway field.
func (o *UnibeeApiMerchantGatewayEditReq) SetSubGateway(v string) {
	o.SubGateway = &v
}

func (o UnibeeApiMerchantGatewayEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyExchange) {
		toSerialize["currencyExchange"] = o.CurrencyExchange
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.GatewayKey) {
		toSerialize["gatewayKey"] = o.GatewayKey
	}
	if !IsNil(o.GatewayLogo) {
		toSerialize["gatewayLogo"] = o.GatewayLogo
	}
	if !IsNil(o.GatewayPaymentTypes) {
		toSerialize["gatewayPaymentTypes"] = o.GatewayPaymentTypes
	}
	if !IsNil(o.GatewaySecret) {
		toSerialize["gatewaySecret"] = o.GatewaySecret
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.SubGateway) {
		toSerialize["subGateway"] = o.SubGateway
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantGatewayEditReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantGatewayEditReq := _UnibeeApiMerchantGatewayEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantGatewayEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantGatewayEditReq(varUnibeeApiMerchantGatewayEditReq)

	return err
}

type NullableUnibeeApiMerchantGatewayEditReq struct {
	value *UnibeeApiMerchantGatewayEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayEditReq) Get() *UnibeeApiMerchantGatewayEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) Set(val *UnibeeApiMerchantGatewayEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayEditReq(val *UnibeeApiMerchantGatewayEditReq) *NullableUnibeeApiMerchantGatewayEditReq {
	return &NullableUnibeeApiMerchantGatewayEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


