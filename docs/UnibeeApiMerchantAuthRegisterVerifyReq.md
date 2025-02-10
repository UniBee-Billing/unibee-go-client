# UnibeeApiMerchantAuthRegisterVerifyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**VerificationCode** | **string** | OTP Code, received from email | 

## Methods

### NewUnibeeApiMerchantAuthRegisterVerifyReq

`func NewUnibeeApiMerchantAuthRegisterVerifyReq(email string, verificationCode string, ) *UnibeeApiMerchantAuthRegisterVerifyReq`

NewUnibeeApiMerchantAuthRegisterVerifyReq instantiates a new UnibeeApiMerchantAuthRegisterVerifyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthRegisterVerifyReqWithDefaults

`func NewUnibeeApiMerchantAuthRegisterVerifyReqWithDefaults() *UnibeeApiMerchantAuthRegisterVerifyReq`

NewUnibeeApiMerchantAuthRegisterVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterVerifyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerificationCode

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *UnibeeApiMerchantAuthRegisterVerifyReq) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


