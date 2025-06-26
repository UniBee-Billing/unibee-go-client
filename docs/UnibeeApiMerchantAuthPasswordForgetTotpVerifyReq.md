# UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**NewPassword** | **string** | The new password | 
**TotpCode** | **string** | The totp code | 

## Methods

### NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReq

`func NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReq(email string, newPassword string, totpCode string, ) *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq`

NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReq instantiates a new UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReqWithDefaults

`func NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReqWithDefaults() *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq`

NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewPassword

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetTotpCode

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


