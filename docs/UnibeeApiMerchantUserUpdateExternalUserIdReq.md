# UnibeeApiMerchantUserUpdateExternalUserIdReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of user, either Email or UserId needed | [optional] 
**ExternalUserId** | **string** | Target ExternalUserId | 
**UserId** | Pointer to **int32** | The id of user, either Email or UserId needed | [optional] 

## Methods

### NewUnibeeApiMerchantUserUpdateExternalUserIdReq

`func NewUnibeeApiMerchantUserUpdateExternalUserIdReq(externalUserId string, ) *UnibeeApiMerchantUserUpdateExternalUserIdReq`

NewUnibeeApiMerchantUserUpdateExternalUserIdReq instantiates a new UnibeeApiMerchantUserUpdateExternalUserIdReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserUpdateExternalUserIdReqWithDefaults

`func NewUnibeeApiMerchantUserUpdateExternalUserIdReqWithDefaults() *UnibeeApiMerchantUserUpdateExternalUserIdReq`

NewUnibeeApiMerchantUserUpdateExternalUserIdReqWithDefaults instantiates a new UnibeeApiMerchantUserUpdateExternalUserIdReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.


### GetUserId

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserUpdateExternalUserIdReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


