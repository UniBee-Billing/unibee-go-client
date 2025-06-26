# UnibeeApiMerchantAuthClearTotpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**TotpResumeCode** | **string** | TOTP Resume Code | 

## Methods

### NewUnibeeApiMerchantAuthClearTotpReq

`func NewUnibeeApiMerchantAuthClearTotpReq(email string, totpResumeCode string, ) *UnibeeApiMerchantAuthClearTotpReq`

NewUnibeeApiMerchantAuthClearTotpReq instantiates a new UnibeeApiMerchantAuthClearTotpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthClearTotpReqWithDefaults

`func NewUnibeeApiMerchantAuthClearTotpReqWithDefaults() *UnibeeApiMerchantAuthClearTotpReq`

NewUnibeeApiMerchantAuthClearTotpReqWithDefaults instantiates a new UnibeeApiMerchantAuthClearTotpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthClearTotpReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthClearTotpReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthClearTotpReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTotpResumeCode

`func (o *UnibeeApiMerchantAuthClearTotpReq) GetTotpResumeCode() string`

GetTotpResumeCode returns the TotpResumeCode field if non-nil, zero value otherwise.

### GetTotpResumeCodeOk

`func (o *UnibeeApiMerchantAuthClearTotpReq) GetTotpResumeCodeOk() (*string, bool)`

GetTotpResumeCodeOk returns a tuple with the TotpResumeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpResumeCode

`func (o *UnibeeApiMerchantAuthClearTotpReq) SetTotpResumeCode(v string)`

SetTotpResumeCode sets TotpResumeCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


