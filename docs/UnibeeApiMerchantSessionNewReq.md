# UnibeeApiMerchantSessionNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**Email** | **string** | Email | 
**ExternalUserId** | **string** | ExternalUserId | 
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Phone** | Pointer to **string** | Phone | [optional] 
**ReturnUrl** | Pointer to **string** | ReturnUrl | [optional] 

## Methods

### NewUnibeeApiMerchantSessionNewReq

`func NewUnibeeApiMerchantSessionNewReq(email string, externalUserId string, ) *UnibeeApiMerchantSessionNewReq`

NewUnibeeApiMerchantSessionNewReq instantiates a new UnibeeApiMerchantSessionNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSessionNewReqWithDefaults

`func NewUnibeeApiMerchantSessionNewReqWithDefaults() *UnibeeApiMerchantSessionNewReq`

NewUnibeeApiMerchantSessionNewReqWithDefaults instantiates a new UnibeeApiMerchantSessionNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiMerchantSessionNewReq) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiMerchantSessionNewReq) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiMerchantSessionNewReq) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiMerchantSessionNewReq) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSessionNewReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSessionNewReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSessionNewReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalUserId

`func (o *UnibeeApiMerchantSessionNewReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSessionNewReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSessionNewReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.


### GetFirstName

`func (o *UnibeeApiMerchantSessionNewReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantSessionNewReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantSessionNewReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantSessionNewReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantSessionNewReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantSessionNewReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantSessionNewReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantSessionNewReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *UnibeeApiMerchantSessionNewReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeApiMerchantSessionNewReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeApiMerchantSessionNewReq) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UnibeeApiMerchantSessionNewReq) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UnibeeApiMerchantSessionNewReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnibeeApiMerchantSessionNewReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnibeeApiMerchantSessionNewReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnibeeApiMerchantSessionNewReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiMerchantSessionNewReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantSessionNewReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantSessionNewReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantSessionNewReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


