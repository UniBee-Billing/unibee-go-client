/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantGetGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGetGet200ResponseData{}

// MerchantGetGet200ResponseData struct for MerchantGetGet200ResponseData
type MerchantGetGet200ResponseData struct {
	// Currency List
	Currency []UnibeeApiBeanCurrency `json:"Currency,omitempty"`
	// The member's role list'
	MemberRoles []UnibeeApiBeanMerchantRole `json:"MemberRoles,omitempty"`
	// TimeZone List
	TimeZone []string `json:"TimeZone,omitempty"`
	// System Env, em: daily|stage|local|prod
	Env *string `json:"env,omitempty"`
	// Gateway List
	Gateways []UnibeeApiBeanGateway `json:"gateways,omitempty"`
	// Check Member is Owner
	IsOwner *bool `json:"isOwner,omitempty"`
	// Check System Env Is Prod, true|false
	IsProd *bool `json:"isProd,omitempty"`
	Merchant *UnibeeApiBeanMerchant `json:"merchant,omitempty"`
	MerchantMember *UnibeeApiBeanDetailMerchantMemberDetail `json:"merchantMember,omitempty"`
	// OpenApiKey
	OpenApiKey *string `json:"openApiKey,omitempty"`
	// SegmentServerSideKey
	SegmentServerSideKey *string `json:"segmentServerSideKey,omitempty"`
	// SegmentUserPortalKey
	SegmentUserPortalKey *string `json:"segmentUserPortalKey,omitempty"`
	// SendGridKey
	SendGridKey *string `json:"sendGridKey,omitempty"`
	// VatSenseKey
	VatSenseKey *string `json:"vatSenseKey,omitempty"`
}

// NewMerchantGetGet200ResponseData instantiates a new MerchantGetGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGetGet200ResponseData() *MerchantGetGet200ResponseData {
	this := MerchantGetGet200ResponseData{}
	return &this
}

// NewMerchantGetGet200ResponseDataWithDefaults instantiates a new MerchantGetGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGetGet200ResponseDataWithDefaults() *MerchantGetGet200ResponseData {
	this := MerchantGetGet200ResponseData{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetCurrency() []UnibeeApiBeanCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret []UnibeeApiBeanCurrency
		return ret
	}
	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetCurrencyOk() ([]UnibeeApiBeanCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given []UnibeeApiBeanCurrency and assigns it to the Currency field.
func (o *MerchantGetGet200ResponseData) SetCurrency(v []UnibeeApiBeanCurrency) {
	o.Currency = v
}

// GetMemberRoles returns the MemberRoles field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetMemberRoles() []UnibeeApiBeanMerchantRole {
	if o == nil || IsNil(o.MemberRoles) {
		var ret []UnibeeApiBeanMerchantRole
		return ret
	}
	return o.MemberRoles
}

// GetMemberRolesOk returns a tuple with the MemberRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetMemberRolesOk() ([]UnibeeApiBeanMerchantRole, bool) {
	if o == nil || IsNil(o.MemberRoles) {
		return nil, false
	}
	return o.MemberRoles, true
}

// HasMemberRoles returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasMemberRoles() bool {
	if o != nil && !IsNil(o.MemberRoles) {
		return true
	}

	return false
}

// SetMemberRoles gets a reference to the given []UnibeeApiBeanMerchantRole and assigns it to the MemberRoles field.
func (o *MerchantGetGet200ResponseData) SetMemberRoles(v []UnibeeApiBeanMerchantRole) {
	o.MemberRoles = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetTimeZone() []string {
	if o == nil || IsNil(o.TimeZone) {
		var ret []string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetTimeZoneOk() ([]string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given []string and assigns it to the TimeZone field.
func (o *MerchantGetGet200ResponseData) SetTimeZone(v []string) {
	o.TimeZone = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetEnv() string {
	if o == nil || IsNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetEnvOk() (*string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *MerchantGetGet200ResponseData) SetEnv(v string) {
	o.Env = &v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetGateways() []UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateways) {
		var ret []UnibeeApiBeanGateway
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetGatewaysOk() ([]UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []UnibeeApiBeanGateway and assigns it to the Gateways field.
func (o *MerchantGetGet200ResponseData) SetGateways(v []UnibeeApiBeanGateway) {
	o.Gateways = v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasIsOwner() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *MerchantGetGet200ResponseData) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetIsProd returns the IsProd field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetIsProd() bool {
	if o == nil || IsNil(o.IsProd) {
		var ret bool
		return ret
	}
	return *o.IsProd
}

// GetIsProdOk returns a tuple with the IsProd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetIsProdOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProd) {
		return nil, false
	}
	return o.IsProd, true
}

// HasIsProd returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasIsProd() bool {
	if o != nil && !IsNil(o.IsProd) {
		return true
	}

	return false
}

// SetIsProd gets a reference to the given bool and assigns it to the IsProd field.
func (o *MerchantGetGet200ResponseData) SetIsProd(v bool) {
	o.IsProd = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetMerchant() UnibeeApiBeanMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchant and assigns it to the Merchant field.
func (o *MerchantGetGet200ResponseData) SetMerchant(v UnibeeApiBeanMerchant) {
	o.Merchant = &v
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the MerchantMember field.
func (o *MerchantGetGet200ResponseData) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.MerchantMember = &v
}

// GetOpenApiKey returns the OpenApiKey field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetOpenApiKey() string {
	if o == nil || IsNil(o.OpenApiKey) {
		var ret string
		return ret
	}
	return *o.OpenApiKey
}

// GetOpenApiKeyOk returns a tuple with the OpenApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetOpenApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OpenApiKey) {
		return nil, false
	}
	return o.OpenApiKey, true
}

// HasOpenApiKey returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasOpenApiKey() bool {
	if o != nil && !IsNil(o.OpenApiKey) {
		return true
	}

	return false
}

// SetOpenApiKey gets a reference to the given string and assigns it to the OpenApiKey field.
func (o *MerchantGetGet200ResponseData) SetOpenApiKey(v string) {
	o.OpenApiKey = &v
}

// GetSegmentServerSideKey returns the SegmentServerSideKey field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetSegmentServerSideKey() string {
	if o == nil || IsNil(o.SegmentServerSideKey) {
		var ret string
		return ret
	}
	return *o.SegmentServerSideKey
}

// GetSegmentServerSideKeyOk returns a tuple with the SegmentServerSideKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetSegmentServerSideKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentServerSideKey) {
		return nil, false
	}
	return o.SegmentServerSideKey, true
}

// HasSegmentServerSideKey returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasSegmentServerSideKey() bool {
	if o != nil && !IsNil(o.SegmentServerSideKey) {
		return true
	}

	return false
}

// SetSegmentServerSideKey gets a reference to the given string and assigns it to the SegmentServerSideKey field.
func (o *MerchantGetGet200ResponseData) SetSegmentServerSideKey(v string) {
	o.SegmentServerSideKey = &v
}

// GetSegmentUserPortalKey returns the SegmentUserPortalKey field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetSegmentUserPortalKey() string {
	if o == nil || IsNil(o.SegmentUserPortalKey) {
		var ret string
		return ret
	}
	return *o.SegmentUserPortalKey
}

// GetSegmentUserPortalKeyOk returns a tuple with the SegmentUserPortalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetSegmentUserPortalKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentUserPortalKey) {
		return nil, false
	}
	return o.SegmentUserPortalKey, true
}

// HasSegmentUserPortalKey returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasSegmentUserPortalKey() bool {
	if o != nil && !IsNil(o.SegmentUserPortalKey) {
		return true
	}

	return false
}

// SetSegmentUserPortalKey gets a reference to the given string and assigns it to the SegmentUserPortalKey field.
func (o *MerchantGetGet200ResponseData) SetSegmentUserPortalKey(v string) {
	o.SegmentUserPortalKey = &v
}

// GetSendGridKey returns the SendGridKey field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetSendGridKey() string {
	if o == nil || IsNil(o.SendGridKey) {
		var ret string
		return ret
	}
	return *o.SendGridKey
}

// GetSendGridKeyOk returns a tuple with the SendGridKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetSendGridKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SendGridKey) {
		return nil, false
	}
	return o.SendGridKey, true
}

// HasSendGridKey returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasSendGridKey() bool {
	if o != nil && !IsNil(o.SendGridKey) {
		return true
	}

	return false
}

// SetSendGridKey gets a reference to the given string and assigns it to the SendGridKey field.
func (o *MerchantGetGet200ResponseData) SetSendGridKey(v string) {
	o.SendGridKey = &v
}

// GetVatSenseKey returns the VatSenseKey field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetVatSenseKey() string {
	if o == nil || IsNil(o.VatSenseKey) {
		var ret string
		return ret
	}
	return *o.VatSenseKey
}

// GetVatSenseKeyOk returns a tuple with the VatSenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetVatSenseKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VatSenseKey) {
		return nil, false
	}
	return o.VatSenseKey, true
}

// HasVatSenseKey returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasVatSenseKey() bool {
	if o != nil && !IsNil(o.VatSenseKey) {
		return true
	}

	return false
}

// SetVatSenseKey gets a reference to the given string and assigns it to the VatSenseKey field.
func (o *MerchantGetGet200ResponseData) SetVatSenseKey(v string) {
	o.VatSenseKey = &v
}

func (o MerchantGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGetGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["Currency"] = o.Currency
	}
	if !IsNil(o.MemberRoles) {
		toSerialize["MemberRoles"] = o.MemberRoles
	}
	if !IsNil(o.TimeZone) {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if !IsNil(o.IsOwner) {
		toSerialize["isOwner"] = o.IsOwner
	}
	if !IsNil(o.IsProd) {
		toSerialize["isProd"] = o.IsProd
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	if !IsNil(o.OpenApiKey) {
		toSerialize["openApiKey"] = o.OpenApiKey
	}
	if !IsNil(o.SegmentServerSideKey) {
		toSerialize["segmentServerSideKey"] = o.SegmentServerSideKey
	}
	if !IsNil(o.SegmentUserPortalKey) {
		toSerialize["segmentUserPortalKey"] = o.SegmentUserPortalKey
	}
	if !IsNil(o.SendGridKey) {
		toSerialize["sendGridKey"] = o.SendGridKey
	}
	if !IsNil(o.VatSenseKey) {
		toSerialize["vatSenseKey"] = o.VatSenseKey
	}
	return toSerialize, nil
}

type NullableMerchantGetGet200ResponseData struct {
	value *MerchantGetGet200ResponseData
	isSet bool
}

func (v NullableMerchantGetGet200ResponseData) Get() *MerchantGetGet200ResponseData {
	return v.value
}

func (v *NullableMerchantGetGet200ResponseData) Set(val *MerchantGetGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGetGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGetGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGetGet200ResponseData(val *MerchantGetGet200ResponseData) *NullableMerchantGetGet200ResponseData {
	return &NullableMerchantGetGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGetGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


