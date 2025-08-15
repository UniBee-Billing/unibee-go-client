# UnibeeApiMerchantAuthSetupOAuthReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**NewPassword** | Pointer to **string** | The new password | [optional] 
**SetupToken** | **string** | The merchant member password setup token | 

## Methods

### NewUnibeeApiMerchantAuthSetupOAuthReq

`func NewUnibeeApiMerchantAuthSetupOAuthReq(email string, setupToken string, ) *UnibeeApiMerchantAuthSetupOAuthReq`

NewUnibeeApiMerchantAuthSetupOAuthReq instantiates a new UnibeeApiMerchantAuthSetupOAuthReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthSetupOAuthReqWithDefaults

`func NewUnibeeApiMerchantAuthSetupOAuthReqWithDefaults() *UnibeeApiMerchantAuthSetupOAuthReq`

NewUnibeeApiMerchantAuthSetupOAuthReqWithDefaults instantiates a new UnibeeApiMerchantAuthSetupOAuthReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewPassword

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetSetupToken

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetSetupToken() string`

GetSetupToken returns the SetupToken field if non-nil, zero value otherwise.

### GetSetupTokenOk

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) GetSetupTokenOk() (*string, bool)`

GetSetupTokenOk returns a tuple with the SetupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupToken

`func (o *UnibeeApiMerchantAuthSetupOAuthReq) SetSetupToken(v string)`

SetSetupToken sets SetupToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


