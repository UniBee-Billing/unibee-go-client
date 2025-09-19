# UnibeeApiMerchantAuthRegisterOAuthReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Company Name | [optional] 
**CountryCode** | Pointer to **string** | Country Code | [optional] 
**CountryName** | Pointer to **string** | Country Name | [optional] 
**Email** | **string** | The merchant member email address | 
**FirstName** | Pointer to **string** | The merchant owner&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The merchant owner&#39;s last name | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**Password** | Pointer to **string** | The owner&#39;s password | [optional] 
**Phone** | Pointer to **string** | The owner&#39;s Phone | [optional] 
**UserName** | Pointer to **string** | The owner&#39;s UserName | [optional] 

## Methods

### NewUnibeeApiMerchantAuthRegisterOAuthReq

`func NewUnibeeApiMerchantAuthRegisterOAuthReq(email string, ) *UnibeeApiMerchantAuthRegisterOAuthReq`

NewUnibeeApiMerchantAuthRegisterOAuthReq instantiates a new UnibeeApiMerchantAuthRegisterOAuthReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthRegisterOAuthReqWithDefaults

`func NewUnibeeApiMerchantAuthRegisterOAuthReqWithDefaults() *UnibeeApiMerchantAuthRegisterOAuthReq`

NewUnibeeApiMerchantAuthRegisterOAuthReqWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterOAuthReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeApiMerchantAuthRegisterOAuthReq) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


