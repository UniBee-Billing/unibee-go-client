# UnibeeApiBeanDetailUserAccountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | address | [optional] 
**AvatarUrl** | Pointer to **string** | avator url | [optional] 
**BillingType** | Pointer to **int32** | 1-recurring,2-one-time | [optional] 
**Birthday** | Pointer to **string** | brithday | [optional] 
**City** | Pointer to **string** | city | [optional] 
**CompanyName** | Pointer to **string** | company name | [optional] 
**CountryCode** | Pointer to **string** | country_code | [optional] 
**CountryName** | Pointer to **string** | country_name | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**CreditAccounts** | Pointer to [**[]UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) | creditAccounts | [optional] 
**Custom** | Pointer to **string** | custom | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**ExternalUserId** | Pointer to **string** | external_user_id | [optional] 
**Facebook** | Pointer to **string** | facebook | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayPaymentType** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** | gender | [optional] 
**Id** | Pointer to **int64** | userId | [optional] 
**IsRisk** | Pointer to **int32** | is risk account (deperated) | [optional] 
**IsSpecial** | Pointer to **int32** | is special account（0.no，1.yes）- deperated | [optional] 
**Language** | Pointer to **string** | User Language, en|ru|cn|vi|bp | [optional] 
**LastLoginAt** | Pointer to **int64** | last login time, utc time | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**LinkedIn** | Pointer to **string** | linkedin | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 
**OtherSocialInfo** | Pointer to **string** |  | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** | phone | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**Profession** | Pointer to **string** | profession | [optional] 
**PromoCreditAccounts** | Pointer to [**[]UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) | promoCreditAccounts | [optional] 
**RecurringAmount** | Pointer to **int64** | total recurring amount, cent | [optional] 
**RegistrationNumber** | Pointer to **string** | RegistrationNumber | [optional] 
**School** | Pointer to **string** | school | [optional] 
**Status** | Pointer to **int32** | 0-Active, 2-Suspend | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**SubscriptionName** | Pointer to **string** | subscription name | [optional] 
**SubscriptionStatus** | Pointer to **int32** | sub status， 1-Pending｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | [optional] 
**TaxPercentage** | Pointer to **int64** | taxPercentage，1000 &#x3D; 10% | [optional] 
**Telegram** | Pointer to **string** | telegram | [optional] 
**TikTok** | Pointer to **string** | tictok | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int64** | User type, 1-Individual|2-organization | [optional] 
**UserName** | Pointer to **string** | user name | [optional] 
**VATNumber** | Pointer to **string** | vat number | [optional] 
**Version** | Pointer to **int32** | version | [optional] 
**WeChat** | Pointer to **string** | wechat | [optional] 
**WhatsAPP** | Pointer to **string** | whats app | [optional] 
**ZipCode** | Pointer to **string** | zip_code | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserAccountDetail

`func NewUnibeeApiBeanDetailUserAccountDetail() *UnibeeApiBeanDetailUserAccountDetail`

NewUnibeeApiBeanDetailUserAccountDetail instantiates a new UnibeeApiBeanDetailUserAccountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserAccountDetailWithDefaults

`func NewUnibeeApiBeanDetailUserAccountDetailWithDefaults() *UnibeeApiBeanDetailUserAccountDetail`

NewUnibeeApiBeanDetailUserAccountDetailWithDefaults instantiates a new UnibeeApiBeanDetailUserAccountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetBirthday

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCreditAccounts() []UnibeeApiBeanCreditAccount`

GetCreditAccounts returns the CreditAccounts field if non-nil, zero value otherwise.

### GetCreditAccountsOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCreditAccountsOk() (*[]UnibeeApiBeanCreditAccount, bool)`

GetCreditAccountsOk returns a tuple with the CreditAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCreditAccounts(v []UnibeeApiBeanCreditAccount)`

SetCreditAccounts sets CreditAccounts field to given value.

### HasCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCreditAccounts() bool`

HasCreditAccounts returns a boolean if a field has been set.

### GetCustom

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFacebook

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetGender

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRisk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetIsRisk() int32`

GetIsRisk returns the IsRisk field if non-nil, zero value otherwise.

### GetIsRiskOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetIsRiskOk() (*int32, bool)`

GetIsRiskOk returns a tuple with the IsRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRisk

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetIsRisk(v int32)`

SetIsRisk sets IsRisk field to given value.

### HasIsRisk

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasIsRisk() bool`

HasIsRisk returns a boolean if a field has been set.

### GetIsSpecial

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetIsSpecial() int32`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetIsSpecialOk() (*int32, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetIsSpecial(v int32)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### GetLanguage

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLastLoginAt() int64`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLastLoginAtOk() (*int64, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetLastLoginAt(v int64)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedIn

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetOtherSocialInfo

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetOtherSocialInfo() string`

GetOtherSocialInfo returns the OtherSocialInfo field if non-nil, zero value otherwise.

### GetOtherSocialInfoOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetOtherSocialInfoOk() (*string, bool)`

GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSocialInfo

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetOtherSocialInfo(v string)`

SetOtherSocialInfo sets OtherSocialInfo field to given value.

### HasOtherSocialInfo

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasOtherSocialInfo() bool`

HasOtherSocialInfo returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProfession

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetPromoCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPromoCreditAccounts() []UnibeeApiBeanCreditAccount`

GetPromoCreditAccounts returns the PromoCreditAccounts field if non-nil, zero value otherwise.

### GetPromoCreditAccountsOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetPromoCreditAccountsOk() (*[]UnibeeApiBeanCreditAccount, bool)`

GetPromoCreditAccountsOk returns a tuple with the PromoCreditAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetPromoCreditAccounts(v []UnibeeApiBeanCreditAccount)`

SetPromoCreditAccounts sets PromoCreditAccounts field to given value.

### HasPromoCreditAccounts

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasPromoCreditAccounts() bool`

HasPromoCreditAccounts returns a boolean if a field has been set.

### GetRecurringAmount

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetRecurringAmount() int64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetRecurringAmountOk() (*int64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetRecurringAmount(v int64)`

SetRecurringAmount sets RecurringAmount field to given value.

### HasRecurringAmount

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasRecurringAmount() bool`

HasRecurringAmount returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetSchool

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionStatus() int32`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetSubscriptionStatusOk() (*int32, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetSubscriptionStatus(v int32)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTelegram

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTikTok

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTikTok() string`

GetTikTok returns the TikTok field if non-nil, zero value otherwise.

### GetTikTokOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTikTokOk() (*string, bool)`

GetTikTokOk returns a tuple with the TikTok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTok

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetTikTok(v string)`

SetTikTok sets TikTok field to given value.

### HasTikTok

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasTikTok() bool`

HasTikTok returns a boolean if a field has been set.

### GetTimeZone

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVATNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.

### HasVATNumber

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasVATNumber() bool`

HasVATNumber returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWeChat

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetWeChat() string`

GetWeChat returns the WeChat field if non-nil, zero value otherwise.

### GetWeChatOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetWeChatOk() (*string, bool)`

GetWeChatOk returns a tuple with the WeChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeChat

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetWeChat(v string)`

SetWeChat sets WeChat field to given value.

### HasWeChat

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasWeChat() bool`

HasWeChat returns a boolean if a field has been set.

### GetWhatsAPP

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetWhatsAPP() string`

GetWhatsAPP returns the WhatsAPP field if non-nil, zero value otherwise.

### GetWhatsAPPOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetWhatsAPPOk() (*string, bool)`

GetWhatsAPPOk returns a tuple with the WhatsAPP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAPP

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetWhatsAPP(v string)`

SetWhatsAPP sets WhatsAPP field to given value.

### HasWhatsAPP

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasWhatsAPP() bool`

HasWhatsAPP returns a boolean if a field has been set.

### GetZipCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiBeanDetailUserAccountDetail) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiBeanDetailUserAccountDetail) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


