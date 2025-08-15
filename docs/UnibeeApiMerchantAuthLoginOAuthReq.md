# UnibeeApiMerchantAuthLoginOAuthReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant member email address | 
**TotpCode** | Pointer to **string** | The totp code | [optional] 

## Methods

### NewUnibeeApiMerchantAuthLoginOAuthReq

`func NewUnibeeApiMerchantAuthLoginOAuthReq(email string, ) *UnibeeApiMerchantAuthLoginOAuthReq`

NewUnibeeApiMerchantAuthLoginOAuthReq instantiates a new UnibeeApiMerchantAuthLoginOAuthReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthLoginOAuthReqWithDefaults

`func NewUnibeeApiMerchantAuthLoginOAuthReqWithDefaults() *UnibeeApiMerchantAuthLoginOAuthReq`

NewUnibeeApiMerchantAuthLoginOAuthReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOAuthReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTotpCode

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *UnibeeApiMerchantAuthLoginOAuthReq) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


