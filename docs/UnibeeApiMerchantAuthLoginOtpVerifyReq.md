# UnibeeApiMerchantAuthLoginOtpVerifyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**VerificationCode** | **string** | OTP Code, received from email | 

## Methods

### NewUnibeeApiMerchantAuthLoginOtpVerifyReq

`func NewUnibeeApiMerchantAuthLoginOtpVerifyReq(email string, verificationCode string, ) *UnibeeApiMerchantAuthLoginOtpVerifyReq`

NewUnibeeApiMerchantAuthLoginOtpVerifyReq instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthLoginOtpVerifyReqWithDefaults

`func NewUnibeeApiMerchantAuthLoginOtpVerifyReqWithDefaults() *UnibeeApiMerchantAuthLoginOtpVerifyReq`

NewUnibeeApiMerchantAuthLoginOtpVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerificationCode

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


