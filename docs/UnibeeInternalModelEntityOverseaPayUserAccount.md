# UnibeeInternalModelEntityOverseaPayUserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | address | [optional] 
**AvatarUrl** | Pointer to **string** | avator url | [optional] 
**BillingType** | Pointer to **int32** | 1-recurring,2-one-time | [optional] 
**Birthday** | Pointer to **string** | brithday | [optional] 
**CompanyName** | Pointer to **string** | company name | [optional] 
**CountryCode** | Pointer to **string** | country_code | [optional] 
**CountryName** | Pointer to **string** | country_name | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Custom** | Pointer to **string** | custom | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**ExternalUserId** | Pointer to **string** | external_user_id | [optional] 
**Facebook** | Pointer to **string** | facebook | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**GatewayId** | Pointer to **string** | gateway_id | [optional] 
**Gender** | Pointer to **string** | gender | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** | userId | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**IsRisk** | Pointer to **int32** | is risk account (deperated) | [optional] 
**IsSpecial** | Pointer to **int32** | is special account（0.no，1.yes）- deperated | [optional] 
**LastLoginAt** | Pointer to **int64** | last login time, utc time | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**LinkedIn** | Pointer to **string** | linkedin | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 
**OtherSocialInfo** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | password , encrypt | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** | phone | [optional] 
**Profession** | Pointer to **string** | profession | [optional] 
**ReMark** | Pointer to **string** | note | [optional] 
**RecurringAmount** | Pointer to **int64** | total recurring amount, cent | [optional] 
**School** | Pointer to **string** | school | [optional] 
**Status** | Pointer to **int32** | 0-Active, 2-Frozen | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**SubscriptionName** | Pointer to **string** | subscription name | [optional] 
**SubscriptionStatus** | Pointer to **int32** | sub status，0-Init | 1-Create｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | [optional] 
**Telegram** | Pointer to **string** | telegram | [optional] 
**TikTok** | Pointer to **string** | tictok | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** | user name | [optional] 
**VATNumber** | Pointer to **string** | vat number | [optional] 
**Version** | Pointer to **int32** | version | [optional] 
**WeChat** | Pointer to **string** | wechat | [optional] 
**WhatsAPP** | Pointer to **string** | whats app | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayUserAccount

`func NewUnibeeInternalModelEntityOverseaPayUserAccount() *UnibeeInternalModelEntityOverseaPayUserAccount`

NewUnibeeInternalModelEntityOverseaPayUserAccount instantiates a new UnibeeInternalModelEntityOverseaPayUserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayUserAccountWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayUserAccountWithDefaults() *UnibeeInternalModelEntityOverseaPayUserAccount`

NewUnibeeInternalModelEntityOverseaPayUserAccountWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayUserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetBirthday

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustom

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFacebook

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGender

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsRisk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsRisk() int32`

GetIsRisk returns the IsRisk field if non-nil, zero value otherwise.

### GetIsRiskOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsRiskOk() (*int32, bool)`

GetIsRiskOk returns a tuple with the IsRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRisk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetIsRisk(v int32)`

SetIsRisk sets IsRisk field to given value.

### HasIsRisk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasIsRisk() bool`

HasIsRisk returns a boolean if a field has been set.

### GetIsSpecial

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsSpecial() int32`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetIsSpecialOk() (*int32, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetIsSpecial(v int32)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLastLoginAt() int64`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLastLoginAtOk() (*int64, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetLastLoginAt(v int64)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedIn

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetOtherSocialInfo

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetOtherSocialInfo() string`

GetOtherSocialInfo returns the OtherSocialInfo field if non-nil, zero value otherwise.

### GetOtherSocialInfoOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetOtherSocialInfoOk() (*string, bool)`

GetOtherSocialInfoOk returns a tuple with the OtherSocialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSocialInfo

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetOtherSocialInfo(v string)`

SetOtherSocialInfo sets OtherSocialInfo field to given value.

### HasOtherSocialInfo

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasOtherSocialInfo() bool`

HasOtherSocialInfo returns a boolean if a field has been set.

### GetPassword

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProfession

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetReMark

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetReMark() string`

GetReMark returns the ReMark field if non-nil, zero value otherwise.

### GetReMarkOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetReMarkOk() (*string, bool)`

GetReMarkOk returns a tuple with the ReMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReMark

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetReMark(v string)`

SetReMark sets ReMark field to given value.

### HasReMark

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasReMark() bool`

HasReMark returns a boolean if a field has been set.

### GetRecurringAmount

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetRecurringAmount() int64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetRecurringAmountOk() (*int64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetRecurringAmount(v int64)`

SetRecurringAmount sets RecurringAmount field to given value.

### HasRecurringAmount

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasRecurringAmount() bool`

HasRecurringAmount returns a boolean if a field has been set.

### GetSchool

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionStatus() int32`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetSubscriptionStatusOk() (*int32, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetSubscriptionStatus(v int32)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTelegram

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTikTok

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTikTok() string`

GetTikTok returns the TikTok field if non-nil, zero value otherwise.

### GetTikTokOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTikTokOk() (*string, bool)`

GetTikTokOk returns a tuple with the TikTok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTok

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetTikTok(v string)`

SetTikTok sets TikTok field to given value.

### HasTikTok

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasTikTok() bool`

HasTikTok returns a boolean if a field has been set.

### GetTimeZone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVATNumber

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.

### HasVATNumber

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasVATNumber() bool`

HasVATNumber returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWeChat

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetWeChat() string`

GetWeChat returns the WeChat field if non-nil, zero value otherwise.

### GetWeChatOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetWeChatOk() (*string, bool)`

GetWeChatOk returns a tuple with the WeChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeChat

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetWeChat(v string)`

SetWeChat sets WeChat field to given value.

### HasWeChat

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasWeChat() bool`

HasWeChat returns a boolean if a field has been set.

### GetWhatsAPP

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetWhatsAPP() string`

GetWhatsAPP returns the WhatsAPP field if non-nil, zero value otherwise.

### GetWhatsAPPOk

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) GetWhatsAPPOk() (*string, bool)`

GetWhatsAPPOk returns a tuple with the WhatsAPP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAPP

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) SetWhatsAPP(v string)`

SetWhatsAPP sets WhatsAPP field to given value.

### HasWhatsAPP

`func (o *UnibeeInternalModelEntityOverseaPayUserAccount) HasWhatsAPP() bool`

HasWhatsAPP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


