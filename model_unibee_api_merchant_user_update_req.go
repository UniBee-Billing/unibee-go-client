/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantUserUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserUpdateReq{}

// UnibeeApiMerchantUserUpdateReq struct for UnibeeApiMerchantUserUpdateReq
type UnibeeApiMerchantUserUpdateReq struct {
	// LinkedIn
	LinkedIn *string `json:"LinkedIn,omitempty"`
	// Billing Address
	Address string `json:"address"`
	// city
	City *string `json:"city,omitempty"`
	// Company Name
	CompanyName *string `json:"companyName,omitempty"`
	// Country Code
	CountryCode *string `json:"countryCode,omitempty"`
	// Country Name
	CountryName *string `json:"countryName,omitempty"`
	// Email
	Email string `json:"email"`
	// Facebook
	Facebook *string `json:"facebook,omitempty"`
	// First name
	FirstName string `json:"firstName"`
	// GatewayId
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// User Language, en|ru|cn|vi|bp
	Language *string `json:"language,omitempty"`
	// Last Name
	LastName string `json:"lastName"`
	// Other Social Info
	OtherSocialInfo *string `json:"otherSocialInfo,omitempty"`
	// PaymentMethodId of gateway, available for card type gateway, payment automatic will enable if set
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	// Phone
	Phone *string `json:"phone,omitempty"`
	// Telegram
	Telegram *string `json:"telegram,omitempty"`
	// Tiktok
	Tiktok *string `json:"tiktok,omitempty"`
	// User type, 1-Individual|2-organization
	Type *int32 `json:"type,omitempty"`
	// User Id
	UserId int64 `json:"userId"`
	// VAT Number
	VATNumber *string `json:"vATNumber,omitempty"`
	// WeChat
	WeChat *string `json:"weChat,omitempty"`
	// WhatsApp
	WhatsApp *string `json:"whatsApp,omitempty"`
	// zip_code
	ZipCode *string `json:"zipCode,omitempty"`
}

type _UnibeeApiMerchantUserUpdateReq UnibeeApiMerchantUserUpdateReq

// NewUnibeeApiMerchantUserUpdateReq instantiates a new UnibeeApiMerchantUserUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserUpdateReq(address string, email string, firstName string, lastName string, userId int64) *UnibeeApiMerchantUserUpdateReq {
	this := UnibeeApiMerchantUserUpdateReq{}
	this.Address = address
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantUserUpdateReqWithDefaults instantiates a new UnibeeApiMerchantUserUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserUpdateReqWithDefaults() *UnibeeApiMerchantUserUpdateReq {
	this := UnibeeApiMerchantUserUpdateReq{}
	return &this
}

// GetLinkedIn returns the LinkedIn field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetLinkedIn() string {
	if o == nil || IsNil(o.LinkedIn) {
		var ret string
		return ret
	}
	return *o.LinkedIn
}

// GetLinkedInOk returns a tuple with the LinkedIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetLinkedInOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedIn) {
		return nil, false
	}
	return o.LinkedIn, true
}

// HasLinkedIn returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasLinkedIn() bool {
	if o != nil && !IsNil(o.LinkedIn) {
		return true
	}

	return false
}

// SetLinkedIn gets a reference to the given string and assigns it to the LinkedIn field.
func (o *UnibeeApiMerchantUserUpdateReq) SetLinkedIn(v string) {
	o.LinkedIn = &v
}

// GetAddress returns the Address field value
func (o *UnibeeApiMerchantUserUpdateReq) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UnibeeApiMerchantUserUpdateReq) SetAddress(v string) {
	o.Address = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UnibeeApiMerchantUserUpdateReq) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UnibeeApiMerchantUserUpdateReq) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiMerchantUserUpdateReq) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetCountryName() string {
	if o == nil || IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasCountryName() bool {
	if o != nil && !IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *UnibeeApiMerchantUserUpdateReq) SetCountryName(v string) {
	o.CountryName = &v
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantUserUpdateReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantUserUpdateReq) SetEmail(v string) {
	o.Email = v
}

// GetFacebook returns the Facebook field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetFacebook() string {
	if o == nil || IsNil(o.Facebook) {
		var ret string
		return ret
	}
	return *o.Facebook
}

// GetFacebookOk returns a tuple with the Facebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetFacebookOk() (*string, bool) {
	if o == nil || IsNil(o.Facebook) {
		return nil, false
	}
	return o.Facebook, true
}

// HasFacebook returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasFacebook() bool {
	if o != nil && !IsNil(o.Facebook) {
		return true
	}

	return false
}

// SetFacebook gets a reference to the given string and assigns it to the Facebook field.
func (o *UnibeeApiMerchantUserUpdateReq) SetFacebook(v string) {
	o.Facebook = &v
}

// GetFirstName returns the FirstName field value
func (o *UnibeeApiMerchantUserUpdateReq) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UnibeeApiMerchantUserUpdateReq) SetFirstName(v string) {
	o.FirstName = v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantUserUpdateReq) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UnibeeApiMerchantUserUpdateReq) SetLanguage(v string) {
	o.Language = &v
}

// GetLastName returns the LastName field value
func (o *UnibeeApiMerchantUserUpdateReq) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UnibeeApiMerchantUserUpdateReq) SetLastName(v string) {
	o.LastName = v
}

// GetOtherSocialInfo returns the OtherSocialInfo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetOtherSocialInfo() string {
	if o == nil || IsNil(o.OtherSocialInfo) {
		var ret string
		return ret
	}
	return *o.OtherSocialInfo
}

// GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetOtherSocialInfoOk() (*string, bool) {
	if o == nil || IsNil(o.OtherSocialInfo) {
		return nil, false
	}
	return o.OtherSocialInfo, true
}

// HasOtherSocialInfo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasOtherSocialInfo() bool {
	if o != nil && !IsNil(o.OtherSocialInfo) {
		return true
	}

	return false
}

// SetOtherSocialInfo gets a reference to the given string and assigns it to the OtherSocialInfo field.
func (o *UnibeeApiMerchantUserUpdateReq) SetOtherSocialInfo(v string) {
	o.OtherSocialInfo = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *UnibeeApiMerchantUserUpdateReq) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeApiMerchantUserUpdateReq) SetPhone(v string) {
	o.Phone = &v
}

// GetTelegram returns the Telegram field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetTelegram() string {
	if o == nil || IsNil(o.Telegram) {
		var ret string
		return ret
	}
	return *o.Telegram
}

// GetTelegramOk returns a tuple with the Telegram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetTelegramOk() (*string, bool) {
	if o == nil || IsNil(o.Telegram) {
		return nil, false
	}
	return o.Telegram, true
}

// HasTelegram returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasTelegram() bool {
	if o != nil && !IsNil(o.Telegram) {
		return true
	}

	return false
}

// SetTelegram gets a reference to the given string and assigns it to the Telegram field.
func (o *UnibeeApiMerchantUserUpdateReq) SetTelegram(v string) {
	o.Telegram = &v
}

// GetTiktok returns the Tiktok field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetTiktok() string {
	if o == nil || IsNil(o.Tiktok) {
		var ret string
		return ret
	}
	return *o.Tiktok
}

// GetTiktokOk returns a tuple with the Tiktok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetTiktokOk() (*string, bool) {
	if o == nil || IsNil(o.Tiktok) {
		return nil, false
	}
	return o.Tiktok, true
}

// HasTiktok returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasTiktok() bool {
	if o != nil && !IsNil(o.Tiktok) {
		return true
	}

	return false
}

// SetTiktok gets a reference to the given string and assigns it to the Tiktok field.
func (o *UnibeeApiMerchantUserUpdateReq) SetTiktok(v string) {
	o.Tiktok = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiMerchantUserUpdateReq) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantUserUpdateReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantUserUpdateReq) SetUserId(v int64) {
	o.UserId = v
}

// GetVATNumber returns the VATNumber field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetVATNumber() string {
	if o == nil || IsNil(o.VATNumber) {
		var ret string
		return ret
	}
	return *o.VATNumber
}

// GetVATNumberOk returns a tuple with the VATNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetVATNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VATNumber) {
		return nil, false
	}
	return o.VATNumber, true
}

// HasVATNumber returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasVATNumber() bool {
	if o != nil && !IsNil(o.VATNumber) {
		return true
	}

	return false
}

// SetVATNumber gets a reference to the given string and assigns it to the VATNumber field.
func (o *UnibeeApiMerchantUserUpdateReq) SetVATNumber(v string) {
	o.VATNumber = &v
}

// GetWeChat returns the WeChat field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetWeChat() string {
	if o == nil || IsNil(o.WeChat) {
		var ret string
		return ret
	}
	return *o.WeChat
}

// GetWeChatOk returns a tuple with the WeChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetWeChatOk() (*string, bool) {
	if o == nil || IsNil(o.WeChat) {
		return nil, false
	}
	return o.WeChat, true
}

// HasWeChat returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasWeChat() bool {
	if o != nil && !IsNil(o.WeChat) {
		return true
	}

	return false
}

// SetWeChat gets a reference to the given string and assigns it to the WeChat field.
func (o *UnibeeApiMerchantUserUpdateReq) SetWeChat(v string) {
	o.WeChat = &v
}

// GetWhatsApp returns the WhatsApp field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetWhatsApp() string {
	if o == nil || IsNil(o.WhatsApp) {
		var ret string
		return ret
	}
	return *o.WhatsApp
}

// GetWhatsAppOk returns a tuple with the WhatsApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetWhatsAppOk() (*string, bool) {
	if o == nil || IsNil(o.WhatsApp) {
		return nil, false
	}
	return o.WhatsApp, true
}

// HasWhatsApp returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasWhatsApp() bool {
	if o != nil && !IsNil(o.WhatsApp) {
		return true
	}

	return false
}

// SetWhatsApp gets a reference to the given string and assigns it to the WhatsApp field.
func (o *UnibeeApiMerchantUserUpdateReq) SetWhatsApp(v string) {
	o.WhatsApp = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserUpdateReq) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserUpdateReq) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserUpdateReq) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UnibeeApiMerchantUserUpdateReq) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o UnibeeApiMerchantUserUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkedIn) {
		toSerialize["LinkedIn"] = o.LinkedIn
	}
	toSerialize["address"] = o.Address
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CountryName) {
		toSerialize["countryName"] = o.CountryName
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.Facebook) {
		toSerialize["facebook"] = o.Facebook
	}
	toSerialize["firstName"] = o.FirstName
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	toSerialize["lastName"] = o.LastName
	if !IsNil(o.OtherSocialInfo) {
		toSerialize["otherSocialInfo"] = o.OtherSocialInfo
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Telegram) {
		toSerialize["telegram"] = o.Telegram
	}
	if !IsNil(o.Tiktok) {
		toSerialize["tiktok"] = o.Tiktok
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.VATNumber) {
		toSerialize["vATNumber"] = o.VATNumber
	}
	if !IsNil(o.WeChat) {
		toSerialize["weChat"] = o.WeChat
	}
	if !IsNil(o.WhatsApp) {
		toSerialize["whatsApp"] = o.WhatsApp
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zipCode"] = o.ZipCode
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantUserUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"email",
		"firstName",
		"lastName",
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

	varUnibeeApiMerchantUserUpdateReq := _UnibeeApiMerchantUserUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantUserUpdateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantUserUpdateReq(varUnibeeApiMerchantUserUpdateReq)

	return err
}

type NullableUnibeeApiMerchantUserUpdateReq struct {
	value *UnibeeApiMerchantUserUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserUpdateReq) Get() *UnibeeApiMerchantUserUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserUpdateReq) Set(val *UnibeeApiMerchantUserUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserUpdateReq(val *UnibeeApiMerchantUserUpdateReq) *NullableUnibeeApiMerchantUserUpdateReq {
	return &NullableUnibeeApiMerchantUserUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


