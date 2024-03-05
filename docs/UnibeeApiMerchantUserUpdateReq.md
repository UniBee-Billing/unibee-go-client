# UnibeeApiMerchantUserUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedIn** | Pointer to **string** | LinkedIn | [optional] 
**Address** | **string** | Billing Address | 
**CompanyName** | Pointer to **string** | Company Name | [optional] 
**CountryCode** | **string** | Country Code | 
**CountryName** | **string** | Country Name | 
**Email** | **string** | Email | 
**Facebook** | Pointer to **string** | Facebook | [optional] 
**FirstName** | **string** | First name | 
**LastName** | **string** | Last Name | 
**OtherSocialInfo** | Pointer to **string** | Other Social Info | [optional] 
**PaymentMethod** | Pointer to **string** | Payment Method | [optional] 
**Phone** | Pointer to **string** | Phone | [optional] 
**Telegram** | Pointer to **string** | Telegram | [optional] 
**Tiktok** | Pointer to **string** | Tiktok | [optional] 
**UserId** | **int32** | User Id | 
**VATNumber** | Pointer to **string** | VAT Number | [optional] 
**WeChat** | Pointer to **string** | WeChat | [optional] 
**WhatsApp** | Pointer to **string** | WhatsApp | [optional] 

## Methods

### NewUnibeeApiMerchantUserUpdateReq

`func NewUnibeeApiMerchantUserUpdateReq(address string, countryCode string, countryName string, email string, firstName string, lastName string, userId int32, ) *UnibeeApiMerchantUserUpdateReq`

NewUnibeeApiMerchantUserUpdateReq instantiates a new UnibeeApiMerchantUserUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserUpdateReqWithDefaults

`func NewUnibeeApiMerchantUserUpdateReqWithDefaults() *UnibeeApiMerchantUserUpdateReq`

NewUnibeeApiMerchantUserUpdateReqWithDefaults instantiates a new UnibeeApiMerchantUserUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedIn

`func (o *UnibeeApiMerchantUserUpdateReq) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *UnibeeApiMerchantUserUpdateReq) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *UnibeeApiMerchantUserUpdateReq) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetAddress

`func (o *UnibeeApiMerchantUserUpdateReq) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiMerchantUserUpdateReq) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCompanyName

`func (o *UnibeeApiMerchantUserUpdateReq) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiMerchantUserUpdateReq) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiMerchantUserUpdateReq) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantUserUpdateReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantUserUpdateReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCountryName

`func (o *UnibeeApiMerchantUserUpdateReq) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiMerchantUserUpdateReq) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetEmail

`func (o *UnibeeApiMerchantUserUpdateReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserUpdateReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFacebook

`func (o *UnibeeApiMerchantUserUpdateReq) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *UnibeeApiMerchantUserUpdateReq) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *UnibeeApiMerchantUserUpdateReq) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiMerchantUserUpdateReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantUserUpdateReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UnibeeApiMerchantUserUpdateReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantUserUpdateReq) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOtherSocialInfo

`func (o *UnibeeApiMerchantUserUpdateReq) GetOtherSocialInfo() string`

GetOtherSocialInfo returns the OtherSocialInfo field if non-nil, zero value otherwise.

### GetOtherSocialInfoOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetOtherSocialInfoOk() (*string, bool)`

GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSocialInfo

`func (o *UnibeeApiMerchantUserUpdateReq) SetOtherSocialInfo(v string)`

SetOtherSocialInfo sets OtherSocialInfo field to given value.

### HasOtherSocialInfo

`func (o *UnibeeApiMerchantUserUpdateReq) HasOtherSocialInfo() bool`

HasOtherSocialInfo returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnibeeApiMerchantUserUpdateReq) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnibeeApiMerchantUserUpdateReq) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiMerchantUserUpdateReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiMerchantUserUpdateReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiMerchantUserUpdateReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTelegram

`func (o *UnibeeApiMerchantUserUpdateReq) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UnibeeApiMerchantUserUpdateReq) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UnibeeApiMerchantUserUpdateReq) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTiktok

`func (o *UnibeeApiMerchantUserUpdateReq) GetTiktok() string`

GetTiktok returns the Tiktok field if non-nil, zero value otherwise.

### GetTiktokOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetTiktokOk() (*string, bool)`

GetTiktokOk returns a tuple with the Tiktok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktok

`func (o *UnibeeApiMerchantUserUpdateReq) SetTiktok(v string)`

SetTiktok sets Tiktok field to given value.

### HasTiktok

`func (o *UnibeeApiMerchantUserUpdateReq) HasTiktok() bool`

HasTiktok returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserUpdateReq) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserUpdateReq) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetVATNumber

`func (o *UnibeeApiMerchantUserUpdateReq) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *UnibeeApiMerchantUserUpdateReq) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.

### HasVATNumber

`func (o *UnibeeApiMerchantUserUpdateReq) HasVATNumber() bool`

HasVATNumber returns a boolean if a field has been set.

### GetWeChat

`func (o *UnibeeApiMerchantUserUpdateReq) GetWeChat() string`

GetWeChat returns the WeChat field if non-nil, zero value otherwise.

### GetWeChatOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetWeChatOk() (*string, bool)`

GetWeChatOk returns a tuple with the WeChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeChat

`func (o *UnibeeApiMerchantUserUpdateReq) SetWeChat(v string)`

SetWeChat sets WeChat field to given value.

### HasWeChat

`func (o *UnibeeApiMerchantUserUpdateReq) HasWeChat() bool`

HasWeChat returns a boolean if a field has been set.

### GetWhatsApp

`func (o *UnibeeApiMerchantUserUpdateReq) GetWhatsApp() string`

GetWhatsApp returns the WhatsApp field if non-nil, zero value otherwise.

### GetWhatsAppOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetWhatsAppOk() (*string, bool)`

GetWhatsAppOk returns a tuple with the WhatsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsApp

`func (o *UnibeeApiMerchantUserUpdateReq) SetWhatsApp(v string)`

SetWhatsApp sets WhatsApp field to given value.

### HasWhatsApp

`func (o *UnibeeApiMerchantUserUpdateReq) HasWhatsApp() bool`

HasWhatsApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


