# UnibeeApiBeanUserAccountSimplify

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
**SubscriptionStatus** | Pointer to **int32** | sub status， 1-Pending｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | [optional] 
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

### NewUnibeeApiBeanUserAccountSimplify

`func NewUnibeeApiBeanUserAccountSimplify() *UnibeeApiBeanUserAccountSimplify`

NewUnibeeApiBeanUserAccountSimplify instantiates a new UnibeeApiBeanUserAccountSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserAccountSimplifyWithDefaults

`func NewUnibeeApiBeanUserAccountSimplifyWithDefaults() *UnibeeApiBeanUserAccountSimplify`

NewUnibeeApiBeanUserAccountSimplifyWithDefaults instantiates a new UnibeeApiBeanUserAccountSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiBeanUserAccountSimplify) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiBeanUserAccountSimplify) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiBeanUserAccountSimplify) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UnibeeApiBeanUserAccountSimplify) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UnibeeApiBeanUserAccountSimplify) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UnibeeApiBeanUserAccountSimplify) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanUserAccountSimplify) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanUserAccountSimplify) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanUserAccountSimplify) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetBirthday

`func (o *UnibeeApiBeanUserAccountSimplify) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UnibeeApiBeanUserAccountSimplify) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UnibeeApiBeanUserAccountSimplify) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiBeanUserAccountSimplify) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiBeanUserAccountSimplify) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiBeanUserAccountSimplify) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiBeanUserAccountSimplify) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiBeanUserAccountSimplify) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiBeanUserAccountSimplify) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanUserAccountSimplify) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanUserAccountSimplify) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanUserAccountSimplify) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiBeanUserAccountSimplify) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiBeanUserAccountSimplify) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiBeanUserAccountSimplify) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanUserAccountSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanUserAccountSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanUserAccountSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustom

`func (o *UnibeeApiBeanUserAccountSimplify) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UnibeeApiBeanUserAccountSimplify) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UnibeeApiBeanUserAccountSimplify) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanUserAccountSimplify) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanUserAccountSimplify) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanUserAccountSimplify) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiBeanUserAccountSimplify) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiBeanUserAccountSimplify) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiBeanUserAccountSimplify) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFacebook

`func (o *UnibeeApiBeanUserAccountSimplify) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *UnibeeApiBeanUserAccountSimplify) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *UnibeeApiBeanUserAccountSimplify) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanUserAccountSimplify) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanUserAccountSimplify) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanUserAccountSimplify) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanUserAccountSimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanUserAccountSimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanUserAccountSimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGender

`func (o *UnibeeApiBeanUserAccountSimplify) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UnibeeApiBeanUserAccountSimplify) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UnibeeApiBeanUserAccountSimplify) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanUserAccountSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanUserAccountSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanUserAccountSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRisk

`func (o *UnibeeApiBeanUserAccountSimplify) GetIsRisk() int32`

GetIsRisk returns the IsRisk field if non-nil, zero value otherwise.

### GetIsRiskOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetIsRiskOk() (*int32, bool)`

GetIsRiskOk returns a tuple with the IsRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRisk

`func (o *UnibeeApiBeanUserAccountSimplify) SetIsRisk(v int32)`

SetIsRisk sets IsRisk field to given value.

### HasIsRisk

`func (o *UnibeeApiBeanUserAccountSimplify) HasIsRisk() bool`

HasIsRisk returns a boolean if a field has been set.

### GetIsSpecial

`func (o *UnibeeApiBeanUserAccountSimplify) GetIsSpecial() int32`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetIsSpecialOk() (*int32, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *UnibeeApiBeanUserAccountSimplify) SetIsSpecial(v int32)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *UnibeeApiBeanUserAccountSimplify) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UnibeeApiBeanUserAccountSimplify) GetLastLoginAt() int64`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetLastLoginAtOk() (*int64, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UnibeeApiBeanUserAccountSimplify) SetLastLoginAt(v int64)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UnibeeApiBeanUserAccountSimplify) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanUserAccountSimplify) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanUserAccountSimplify) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanUserAccountSimplify) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedIn

`func (o *UnibeeApiBeanUserAccountSimplify) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *UnibeeApiBeanUserAccountSimplify) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *UnibeeApiBeanUserAccountSimplify) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanUserAccountSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanUserAccountSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanUserAccountSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeApiBeanUserAccountSimplify) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeApiBeanUserAccountSimplify) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeApiBeanUserAccountSimplify) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetOtherSocialInfo

`func (o *UnibeeApiBeanUserAccountSimplify) GetOtherSocialInfo() string`

GetOtherSocialInfo returns the OtherSocialInfo field if non-nil, zero value otherwise.

### GetOtherSocialInfoOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetOtherSocialInfoOk() (*string, bool)`

GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSocialInfo

`func (o *UnibeeApiBeanUserAccountSimplify) SetOtherSocialInfo(v string)`

SetOtherSocialInfo sets OtherSocialInfo field to given value.

### HasOtherSocialInfo

`func (o *UnibeeApiBeanUserAccountSimplify) HasOtherSocialInfo() bool`

HasOtherSocialInfo returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UnibeeApiBeanUserAccountSimplify) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnibeeApiBeanUserAccountSimplify) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnibeeApiBeanUserAccountSimplify) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiBeanUserAccountSimplify) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiBeanUserAccountSimplify) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiBeanUserAccountSimplify) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProfession

`func (o *UnibeeApiBeanUserAccountSimplify) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *UnibeeApiBeanUserAccountSimplify) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *UnibeeApiBeanUserAccountSimplify) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetReMark

`func (o *UnibeeApiBeanUserAccountSimplify) GetReMark() string`

GetReMark returns the ReMark field if non-nil, zero value otherwise.

### GetReMarkOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetReMarkOk() (*string, bool)`

GetReMarkOk returns a tuple with the ReMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReMark

`func (o *UnibeeApiBeanUserAccountSimplify) SetReMark(v string)`

SetReMark sets ReMark field to given value.

### HasReMark

`func (o *UnibeeApiBeanUserAccountSimplify) HasReMark() bool`

HasReMark returns a boolean if a field has been set.

### GetRecurringAmount

`func (o *UnibeeApiBeanUserAccountSimplify) GetRecurringAmount() int64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetRecurringAmountOk() (*int64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *UnibeeApiBeanUserAccountSimplify) SetRecurringAmount(v int64)`

SetRecurringAmount sets RecurringAmount field to given value.

### HasRecurringAmount

`func (o *UnibeeApiBeanUserAccountSimplify) HasRecurringAmount() bool`

HasRecurringAmount returns a boolean if a field has been set.

### GetSchool

`func (o *UnibeeApiBeanUserAccountSimplify) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *UnibeeApiBeanUserAccountSimplify) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *UnibeeApiBeanUserAccountSimplify) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanUserAccountSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanUserAccountSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanUserAccountSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanUserAccountSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanUserAccountSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *UnibeeApiBeanUserAccountSimplify) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *UnibeeApiBeanUserAccountSimplify) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionStatus() int32`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetSubscriptionStatusOk() (*int32, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *UnibeeApiBeanUserAccountSimplify) SetSubscriptionStatus(v int32)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *UnibeeApiBeanUserAccountSimplify) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanUserAccountSimplify) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanUserAccountSimplify) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanUserAccountSimplify) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTelegram

`func (o *UnibeeApiBeanUserAccountSimplify) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UnibeeApiBeanUserAccountSimplify) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UnibeeApiBeanUserAccountSimplify) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTikTok

`func (o *UnibeeApiBeanUserAccountSimplify) GetTikTok() string`

GetTikTok returns the TikTok field if non-nil, zero value otherwise.

### GetTikTokOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetTikTokOk() (*string, bool)`

GetTikTokOk returns a tuple with the TikTok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTok

`func (o *UnibeeApiBeanUserAccountSimplify) SetTikTok(v string)`

SetTikTok sets TikTok field to given value.

### HasTikTok

`func (o *UnibeeApiBeanUserAccountSimplify) HasTikTok() bool`

HasTikTok returns a boolean if a field has been set.

### GetTimeZone

`func (o *UnibeeApiBeanUserAccountSimplify) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UnibeeApiBeanUserAccountSimplify) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UnibeeApiBeanUserAccountSimplify) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanUserAccountSimplify) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanUserAccountSimplify) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanUserAccountSimplify) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiBeanUserAccountSimplify) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiBeanUserAccountSimplify) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiBeanUserAccountSimplify) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVATNumber

`func (o *UnibeeApiBeanUserAccountSimplify) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *UnibeeApiBeanUserAccountSimplify) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.

### HasVATNumber

`func (o *UnibeeApiBeanUserAccountSimplify) HasVATNumber() bool`

HasVATNumber returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeApiBeanUserAccountSimplify) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeApiBeanUserAccountSimplify) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeApiBeanUserAccountSimplify) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWeChat

`func (o *UnibeeApiBeanUserAccountSimplify) GetWeChat() string`

GetWeChat returns the WeChat field if non-nil, zero value otherwise.

### GetWeChatOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetWeChatOk() (*string, bool)`

GetWeChatOk returns a tuple with the WeChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeChat

`func (o *UnibeeApiBeanUserAccountSimplify) SetWeChat(v string)`

SetWeChat sets WeChat field to given value.

### HasWeChat

`func (o *UnibeeApiBeanUserAccountSimplify) HasWeChat() bool`

HasWeChat returns a boolean if a field has been set.

### GetWhatsAPP

`func (o *UnibeeApiBeanUserAccountSimplify) GetWhatsAPP() string`

GetWhatsAPP returns the WhatsAPP field if non-nil, zero value otherwise.

### GetWhatsAPPOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetWhatsAPPOk() (*string, bool)`

GetWhatsAPPOk returns a tuple with the WhatsAPP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAPP

`func (o *UnibeeApiBeanUserAccountSimplify) SetWhatsAPP(v string)`

SetWhatsAPP sets WhatsAPP field to given value.

### HasWhatsAPP

`func (o *UnibeeApiBeanUserAccountSimplify) HasWhatsAPP() bool`

HasWhatsAPP returns a boolean if a field has been set.

### GetZipCode

`func (o *UnibeeApiBeanUserAccountSimplify) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiBeanUserAccountSimplify) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiBeanUserAccountSimplify) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiBeanUserAccountSimplify) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


