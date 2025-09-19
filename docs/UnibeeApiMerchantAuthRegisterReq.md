# UnibeeApiMerchantAuthRegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Company Name | [optional] 
**CountryCode** | Pointer to **string** | Country Code | [optional] 
**CountryName** | Pointer to **string** | Country Name | [optional] 
**Email** | **string** | The merchant owner&#39;s email address | 
**FirstName** | **string** | The merchant owner&#39;s first name | 
**LastName** | **string** | The merchant owner&#39;s last name | 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**Password** | **string** | The owner&#39;s password | 
**Phone** | Pointer to **string** | The owner&#39;s Phone | [optional] 
**UserName** | Pointer to **string** | The owner&#39;s UserName | [optional] 

## Methods

### NewUnibeeApiMerchantAuthRegisterReq

`func NewUnibeeApiMerchantAuthRegisterReq(email string, firstName string, lastName string, password string, ) *UnibeeApiMerchantAuthRegisterReq`

NewUnibeeApiMerchantAuthRegisterReq instantiates a new UnibeeApiMerchantAuthRegisterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthRegisterReqWithDefaults

`func NewUnibeeApiMerchantAuthRegisterReqWithDefaults() *UnibeeApiMerchantAuthRegisterReq`

NewUnibeeApiMerchantAuthRegisterReqWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiMerchantAuthRegisterReq) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiMerchantAuthRegisterReq) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantAuthRegisterReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantAuthRegisterReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiMerchantAuthRegisterReq) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiMerchantAuthRegisterReq) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantAuthRegisterReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthRegisterReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UnibeeApiMerchantAuthRegisterReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantAuthRegisterReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UnibeeApiMerchantAuthRegisterReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantAuthRegisterReq) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantAuthRegisterReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantAuthRegisterReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantAuthRegisterReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *UnibeeApiMerchantAuthRegisterReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeApiMerchantAuthRegisterReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPhone

`func (o *UnibeeApiMerchantAuthRegisterReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiMerchantAuthRegisterReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiMerchantAuthRegisterReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiMerchantAuthRegisterReq) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiMerchantAuthRegisterReq) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiMerchantAuthRegisterReq) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiMerchantAuthRegisterReq) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


