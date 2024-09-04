# UnibeeApiMerchantUserChangeEmailReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUserId** | Pointer to **string** | The externalUserId of user, either ExternalUserId or UserId needed | [optional] 
**NewEmail** | **string** | Target Email want to change | 
**UserId** | Pointer to **int64** | The id of user, either ExternalUserId or UserId needed | [optional] 

## Methods

### NewUnibeeApiMerchantUserChangeEmailReq

`func NewUnibeeApiMerchantUserChangeEmailReq(newEmail string, ) *UnibeeApiMerchantUserChangeEmailReq`

NewUnibeeApiMerchantUserChangeEmailReq instantiates a new UnibeeApiMerchantUserChangeEmailReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserChangeEmailReqWithDefaults

`func NewUnibeeApiMerchantUserChangeEmailReqWithDefaults() *UnibeeApiMerchantUserChangeEmailReq`

NewUnibeeApiMerchantUserChangeEmailReqWithDefaults instantiates a new UnibeeApiMerchantUserChangeEmailReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetNewEmail

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *UnibeeApiMerchantUserChangeEmailReq) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.


### GetUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserChangeEmailReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserChangeEmailReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


