# UnibeeApiMerchantUserUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedIn** | Pointer to **string** | LinkedIn | [optional] 
**Address** | Pointer to **string** | Billing Address | [optional] 
**City** | Pointer to **string** | city | [optional] 
**CompanyName** | Pointer to **string** | Company Name | [optional] 
**CountryCode** | Pointer to **string** | Country Code | [optional] 
**CountryName** | Pointer to **string** | Country Name | [optional] 
**Email** | Pointer to **string** | The email of user, either Email or UserId needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId | [optional] 
**Facebook** | Pointer to **string** | Facebook | [optional] 
**FirstName** | Pointer to **string** | First name | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId | [optional] 
**Language** | Pointer to **string** | User Language, en|ru|cn|vi|bp | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**OtherSocialInfo** | Pointer to **string** | Other Social Info | [optional] 
**PaymentMethodId** | Pointer to **string** | PaymentMethodId of gateway, available for card type gateway, payment automatic will enable if set | [optional] 
**Phone** | Pointer to **string** | Phone | [optional] 
**RegistrationNumber** | Pointer to **string** | RegistrationNumber | [optional] 
**Telegram** | Pointer to **string** | Telegram | [optional] 
**Tiktok** | Pointer to **string** | Tiktok | [optional] 
**Type** | Pointer to **int32** | User type, 1-Individual|2-organization | [optional] 
**UserId** | Pointer to **int32** | The id of user, either Email or UserId needed | [optional] 
**VATNumber** | Pointer to **string** | VAT Number | [optional] 
**WeChat** | Pointer to **string** | WeChat | [optional] 
**WhatsApp** | Pointer to **string** | WhatsApp | [optional] 
**ZipCode** | Pointer to **string** | zip_code | [optional] 

## Methods

### NewUnibeeApiMerchantUserUpdateReq

`func NewUnibeeApiMerchantUserUpdateReq() *UnibeeApiMerchantUserUpdateReq`

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

### HasAddress

`func (o *UnibeeApiMerchantUserUpdateReq) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiMerchantUserUpdateReq) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiMerchantUserUpdateReq) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiMerchantUserUpdateReq) HasCity() bool`

HasCity returns a boolean if a field has been set.

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

### HasCountryCode

`func (o *UnibeeApiMerchantUserUpdateReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

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

### HasCountryName

`func (o *UnibeeApiMerchantUserUpdateReq) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

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

### HasEmail

`func (o *UnibeeApiMerchantUserUpdateReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantUserUpdateReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantUserUpdateReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantUserUpdateReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

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

### HasFirstName

`func (o *UnibeeApiMerchantUserUpdateReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantUserUpdateReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantUserUpdateReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantUserUpdateReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetLanguage

`func (o *UnibeeApiMerchantUserUpdateReq) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UnibeeApiMerchantUserUpdateReq) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UnibeeApiMerchantUserUpdateReq) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

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

### HasLastName

`func (o *UnibeeApiMerchantUserUpdateReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantUserUpdateReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantUserUpdateReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantUserUpdateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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

### GetPaymentMethodId

`func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiMerchantUserUpdateReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiMerchantUserUpdateReq) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

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

### GetRegistrationNumber

`func (o *UnibeeApiMerchantUserUpdateReq) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *UnibeeApiMerchantUserUpdateReq) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *UnibeeApiMerchantUserUpdateReq) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

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

### GetType

`func (o *UnibeeApiMerchantUserUpdateReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantUserUpdateReq) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantUserUpdateReq) HasType() bool`

HasType returns a boolean if a field has been set.

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

### HasUserId

`func (o *UnibeeApiMerchantUserUpdateReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

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

### GetZipCode

`func (o *UnibeeApiMerchantUserUpdateReq) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiMerchantUserUpdateReq) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiMerchantUserUpdateReq) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiMerchantUserUpdateReq) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


