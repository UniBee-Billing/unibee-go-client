# UnibeeApiBeanNewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**City** | Pointer to **string** | city | [optional] 
**CompanyName** | Pointer to **string** | company name | [optional] 
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**Email** | **string** | Email | 
**ExternalUserId** | Pointer to **string** | ExternalUserId | [optional] 
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Phone** | Pointer to **string** | Phone | [optional] 
**Type** | Pointer to **int64** | User type, 1-Individual|2-organization | [optional] 
**UserName** | Pointer to **string** | UserName | [optional] 
**VatNumber** | Pointer to **string** | vat number | [optional] 
**ZipCode** | Pointer to **string** | zip_code | [optional] 

## Methods

### NewUnibeeApiBeanNewUser

`func NewUnibeeApiBeanNewUser(email string, ) *UnibeeApiBeanNewUser`

NewUnibeeApiBeanNewUser instantiates a new UnibeeApiBeanNewUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanNewUserWithDefaults

`func NewUnibeeApiBeanNewUserWithDefaults() *UnibeeApiBeanNewUser`

NewUnibeeApiBeanNewUserWithDefaults instantiates a new UnibeeApiBeanNewUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiBeanNewUser) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiBeanNewUser) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiBeanNewUser) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiBeanNewUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiBeanNewUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiBeanNewUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiBeanNewUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiBeanNewUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiBeanNewUser) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiBeanNewUser) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiBeanNewUser) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiBeanNewUser) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanNewUser) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanNewUser) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanNewUser) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanNewUser) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanNewUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanNewUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanNewUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalUserId

`func (o *UnibeeApiBeanNewUser) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiBeanNewUser) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiBeanNewUser) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiBeanNewUser) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanNewUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanNewUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanNewUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanNewUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanNewUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanNewUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanNewUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanNewUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiBeanNewUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiBeanNewUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiBeanNewUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiBeanNewUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanNewUser) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanNewUser) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanNewUser) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanNewUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiBeanNewUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiBeanNewUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiBeanNewUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiBeanNewUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanNewUser) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanNewUser) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanNewUser) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanNewUser) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *UnibeeApiBeanNewUser) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiBeanNewUser) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiBeanNewUser) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiBeanNewUser) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


