# UnibeeApiBeanUserAccount

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
**Custom** | Pointer to **string** | custom | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**ExternalUserId** | Pointer to **string** | external_user_id | [optional] 
**Facebook** | Pointer to **string** | facebook | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**Gender** | Pointer to **string** | gender | [optional] 
**Id** | Pointer to **int64** | userId | [optional] 
**IsRisk** | Pointer to **int32** | is risk account (deperated) | [optional] 
**IsSpecial** | Pointer to **int32** | is special account（0.no，1.yes）- deperated | [optional] 
**Language** | Pointer to **string** | User Language, en|ru|cn|vi|bp | [optional] 
**LastLoginAt** | Pointer to **int64** | last login time, utc time | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**LinkedIn** | Pointer to **string** | linkedin | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 
**OtherSocialInfo** | Pointer to **string** |  | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** | phone | [optional] 
**Profession** | Pointer to **string** | profession | [optional] 
**ReMark** | Pointer to **string** | note | [optional] 
**RecurringAmount** | Pointer to **int64** | total recurring amount, cent | [optional] 
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

### NewUnibeeApiBeanUserAccount

`func NewUnibeeApiBeanUserAccount() *UnibeeApiBeanUserAccount`

NewUnibeeApiBeanUserAccount instantiates a new UnibeeApiBeanUserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserAccountWithDefaults

`func NewUnibeeApiBeanUserAccountWithDefaults() *UnibeeApiBeanUserAccount`

NewUnibeeApiBeanUserAccountWithDefaults instantiates a new UnibeeApiBeanUserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiBeanUserAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiBeanUserAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiBeanUserAccount) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiBeanUserAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UnibeeApiBeanUserAccount) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UnibeeApiBeanUserAccount) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UnibeeApiBeanUserAccount) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UnibeeApiBeanUserAccount) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanUserAccount) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanUserAccount) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanUserAccount) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanUserAccount) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetBirthday

`func (o *UnibeeApiBeanUserAccount) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UnibeeApiBeanUserAccount) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UnibeeApiBeanUserAccount) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UnibeeApiBeanUserAccount) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiBeanUserAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiBeanUserAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiBeanUserAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiBeanUserAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiBeanUserAccount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiBeanUserAccount) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiBeanUserAccount) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiBeanUserAccount) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanUserAccount) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanUserAccount) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanUserAccount) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanUserAccount) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiBeanUserAccount) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiBeanUserAccount) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiBeanUserAccount) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiBeanUserAccount) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanUserAccount) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanUserAccount) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanUserAccount) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanUserAccount) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustom

`func (o *UnibeeApiBeanUserAccount) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UnibeeApiBeanUserAccount) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UnibeeApiBeanUserAccount) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UnibeeApiBeanUserAccount) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanUserAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanUserAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanUserAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanUserAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiBeanUserAccount) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiBeanUserAccount) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiBeanUserAccount) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiBeanUserAccount) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFacebook

`func (o *UnibeeApiBeanUserAccount) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *UnibeeApiBeanUserAccount) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *UnibeeApiBeanUserAccount) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *UnibeeApiBeanUserAccount) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanUserAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanUserAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanUserAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanUserAccount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanUserAccount) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanUserAccount) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanUserAccount) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanUserAccount) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGender

`func (o *UnibeeApiBeanUserAccount) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UnibeeApiBeanUserAccount) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UnibeeApiBeanUserAccount) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UnibeeApiBeanUserAccount) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanUserAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanUserAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanUserAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanUserAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRisk

`func (o *UnibeeApiBeanUserAccount) GetIsRisk() int32`

GetIsRisk returns the IsRisk field if non-nil, zero value otherwise.

### GetIsRiskOk

`func (o *UnibeeApiBeanUserAccount) GetIsRiskOk() (*int32, bool)`

GetIsRiskOk returns a tuple with the IsRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRisk

`func (o *UnibeeApiBeanUserAccount) SetIsRisk(v int32)`

SetIsRisk sets IsRisk field to given value.

### HasIsRisk

`func (o *UnibeeApiBeanUserAccount) HasIsRisk() bool`

HasIsRisk returns a boolean if a field has been set.

### GetIsSpecial

`func (o *UnibeeApiBeanUserAccount) GetIsSpecial() int32`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *UnibeeApiBeanUserAccount) GetIsSpecialOk() (*int32, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *UnibeeApiBeanUserAccount) SetIsSpecial(v int32)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *UnibeeApiBeanUserAccount) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### GetLanguage

`func (o *UnibeeApiBeanUserAccount) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UnibeeApiBeanUserAccount) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UnibeeApiBeanUserAccount) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UnibeeApiBeanUserAccount) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UnibeeApiBeanUserAccount) GetLastLoginAt() int64`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UnibeeApiBeanUserAccount) GetLastLoginAtOk() (*int64, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UnibeeApiBeanUserAccount) SetLastLoginAt(v int64)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UnibeeApiBeanUserAccount) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanUserAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanUserAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanUserAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanUserAccount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedIn

`func (o *UnibeeApiBeanUserAccount) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *UnibeeApiBeanUserAccount) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *UnibeeApiBeanUserAccount) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *UnibeeApiBeanUserAccount) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanUserAccount) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanUserAccount) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanUserAccount) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanUserAccount) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeApiBeanUserAccount) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeApiBeanUserAccount) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeApiBeanUserAccount) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeApiBeanUserAccount) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetOtherSocialInfo

`func (o *UnibeeApiBeanUserAccount) GetOtherSocialInfo() string`

GetOtherSocialInfo returns the OtherSocialInfo field if non-nil, zero value otherwise.

### GetOtherSocialInfoOk

`func (o *UnibeeApiBeanUserAccount) GetOtherSocialInfoOk() (*string, bool)`

GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSocialInfo

`func (o *UnibeeApiBeanUserAccount) SetOtherSocialInfo(v string)`

SetOtherSocialInfo sets OtherSocialInfo field to given value.

### HasOtherSocialInfo

`func (o *UnibeeApiBeanUserAccount) HasOtherSocialInfo() bool`

HasOtherSocialInfo returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UnibeeApiBeanUserAccount) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnibeeApiBeanUserAccount) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnibeeApiBeanUserAccount) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnibeeApiBeanUserAccount) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiBeanUserAccount) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiBeanUserAccount) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiBeanUserAccount) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiBeanUserAccount) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProfession

`func (o *UnibeeApiBeanUserAccount) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *UnibeeApiBeanUserAccount) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *UnibeeApiBeanUserAccount) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *UnibeeApiBeanUserAccount) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetReMark

`func (o *UnibeeApiBeanUserAccount) GetReMark() string`

GetReMark returns the ReMark field if non-nil, zero value otherwise.

### GetReMarkOk

`func (o *UnibeeApiBeanUserAccount) GetReMarkOk() (*string, bool)`

GetReMarkOk returns a tuple with the ReMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReMark

`func (o *UnibeeApiBeanUserAccount) SetReMark(v string)`

SetReMark sets ReMark field to given value.

### HasReMark

`func (o *UnibeeApiBeanUserAccount) HasReMark() bool`

HasReMark returns a boolean if a field has been set.

### GetRecurringAmount

`func (o *UnibeeApiBeanUserAccount) GetRecurringAmount() int64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *UnibeeApiBeanUserAccount) GetRecurringAmountOk() (*int64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *UnibeeApiBeanUserAccount) SetRecurringAmount(v int64)`

SetRecurringAmount sets RecurringAmount field to given value.

### HasRecurringAmount

`func (o *UnibeeApiBeanUserAccount) HasRecurringAmount() bool`

HasRecurringAmount returns a boolean if a field has been set.

### GetSchool

`func (o *UnibeeApiBeanUserAccount) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *UnibeeApiBeanUserAccount) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *UnibeeApiBeanUserAccount) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *UnibeeApiBeanUserAccount) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanUserAccount) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanUserAccount) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanUserAccount) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanUserAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanUserAccount) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanUserAccount) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *UnibeeApiBeanUserAccount) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *UnibeeApiBeanUserAccount) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionStatus() int32`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *UnibeeApiBeanUserAccount) GetSubscriptionStatusOk() (*int32, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *UnibeeApiBeanUserAccount) SetSubscriptionStatus(v int32)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *UnibeeApiBeanUserAccount) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanUserAccount) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanUserAccount) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanUserAccount) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanUserAccount) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTelegram

`func (o *UnibeeApiBeanUserAccount) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UnibeeApiBeanUserAccount) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UnibeeApiBeanUserAccount) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UnibeeApiBeanUserAccount) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTikTok

`func (o *UnibeeApiBeanUserAccount) GetTikTok() string`

GetTikTok returns the TikTok field if non-nil, zero value otherwise.

### GetTikTokOk

`func (o *UnibeeApiBeanUserAccount) GetTikTokOk() (*string, bool)`

GetTikTokOk returns a tuple with the TikTok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTok

`func (o *UnibeeApiBeanUserAccount) SetTikTok(v string)`

SetTikTok sets TikTok field to given value.

### HasTikTok

`func (o *UnibeeApiBeanUserAccount) HasTikTok() bool`

HasTikTok returns a boolean if a field has been set.

### GetTimeZone

`func (o *UnibeeApiBeanUserAccount) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UnibeeApiBeanUserAccount) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UnibeeApiBeanUserAccount) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UnibeeApiBeanUserAccount) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanUserAccount) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanUserAccount) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanUserAccount) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanUserAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiBeanUserAccount) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiBeanUserAccount) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiBeanUserAccount) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiBeanUserAccount) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVATNumber

`func (o *UnibeeApiBeanUserAccount) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *UnibeeApiBeanUserAccount) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *UnibeeApiBeanUserAccount) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.

### HasVATNumber

`func (o *UnibeeApiBeanUserAccount) HasVATNumber() bool`

HasVATNumber returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeApiBeanUserAccount) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeApiBeanUserAccount) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeApiBeanUserAccount) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeApiBeanUserAccount) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWeChat

`func (o *UnibeeApiBeanUserAccount) GetWeChat() string`

GetWeChat returns the WeChat field if non-nil, zero value otherwise.

### GetWeChatOk

`func (o *UnibeeApiBeanUserAccount) GetWeChatOk() (*string, bool)`

GetWeChatOk returns a tuple with the WeChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeChat

`func (o *UnibeeApiBeanUserAccount) SetWeChat(v string)`

SetWeChat sets WeChat field to given value.

### HasWeChat

`func (o *UnibeeApiBeanUserAccount) HasWeChat() bool`

HasWeChat returns a boolean if a field has been set.

### GetWhatsAPP

`func (o *UnibeeApiBeanUserAccount) GetWhatsAPP() string`

GetWhatsAPP returns the WhatsAPP field if non-nil, zero value otherwise.

### GetWhatsAPPOk

`func (o *UnibeeApiBeanUserAccount) GetWhatsAPPOk() (*string, bool)`

GetWhatsAPPOk returns a tuple with the WhatsAPP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAPP

`func (o *UnibeeApiBeanUserAccount) SetWhatsAPP(v string)`

SetWhatsAPP sets WhatsAPP field to given value.

### HasWhatsAPP

`func (o *UnibeeApiBeanUserAccount) HasWhatsAPP() bool`

HasWhatsAPP returns a boolean if a field has been set.

### GetZipCode

`func (o *UnibeeApiBeanUserAccount) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiBeanUserAccount) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiBeanUserAccount) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiBeanUserAccount) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


