# UnibeeApiMerchantAuthPasswordSetupOtpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**NewPassword** | **string** | The new password | 
**SetupToken** | **string** | The merchant member password setup token | 

## Methods

### NewUnibeeApiMerchantAuthPasswordSetupOtpReq

`func NewUnibeeApiMerchantAuthPasswordSetupOtpReq(email string, newPassword string, setupToken string, ) *UnibeeApiMerchantAuthPasswordSetupOtpReq`

NewUnibeeApiMerchantAuthPasswordSetupOtpReq instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthPasswordSetupOtpReqWithDefaults

`func NewUnibeeApiMerchantAuthPasswordSetupOtpReqWithDefaults() *UnibeeApiMerchantAuthPasswordSetupOtpReq`

NewUnibeeApiMerchantAuthPasswordSetupOtpReqWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewPassword

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetSetupToken

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetSetupToken() string`

GetSetupToken returns the SetupToken field if non-nil, zero value otherwise.

### GetSetupTokenOk

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetSetupTokenOk() (*string, bool)`

GetSetupTokenOk returns a tuple with the SetupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupToken

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetSetupToken(v string)`

SetSetupToken sets SetupToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


