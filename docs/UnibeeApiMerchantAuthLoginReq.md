# UnibeeApiMerchantAuthLoginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**Password** | **string** | The merchant member password | 
**Provider** | Pointer to **string** | Connect to OAuth provider | [optional] 
**ProviderId** | Pointer to **string** | Connect to OAuth ProviderId | [optional] 
**TotpCode** | Pointer to **string** | The totp code | [optional] 

## Methods

### NewUnibeeApiMerchantAuthLoginReq

`func NewUnibeeApiMerchantAuthLoginReq(email string, password string, ) *UnibeeApiMerchantAuthLoginReq`

NewUnibeeApiMerchantAuthLoginReq instantiates a new UnibeeApiMerchantAuthLoginReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthLoginReqWithDefaults

`func NewUnibeeApiMerchantAuthLoginReqWithDefaults() *UnibeeApiMerchantAuthLoginReq`

NewUnibeeApiMerchantAuthLoginReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthLoginReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthLoginReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthLoginReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UnibeeApiMerchantAuthLoginReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeApiMerchantAuthLoginReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeApiMerchantAuthLoginReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProvider

`func (o *UnibeeApiMerchantAuthLoginReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UnibeeApiMerchantAuthLoginReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UnibeeApiMerchantAuthLoginReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UnibeeApiMerchantAuthLoginReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderId

`func (o *UnibeeApiMerchantAuthLoginReq) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *UnibeeApiMerchantAuthLoginReq) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *UnibeeApiMerchantAuthLoginReq) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *UnibeeApiMerchantAuthLoginReq) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetTotpCode

`func (o *UnibeeApiMerchantAuthLoginReq) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *UnibeeApiMerchantAuthLoginReq) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *UnibeeApiMerchantAuthLoginReq) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *UnibeeApiMerchantAuthLoginReq) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


