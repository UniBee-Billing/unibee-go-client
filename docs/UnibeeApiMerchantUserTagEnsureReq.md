# UnibeeApiMerchantUserTagEnsureReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutMs** | Pointer to **int32** | Wait timeout milliseconds, default 3000 | [optional] 
**UserId** | **int64** |  | 
**WaitForReady** | Pointer to **bool** | Wait for calculation to finish | [optional] 

## Methods

### NewUnibeeApiMerchantUserTagEnsureReq

`func NewUnibeeApiMerchantUserTagEnsureReq(userId int64, ) *UnibeeApiMerchantUserTagEnsureReq`

NewUnibeeApiMerchantUserTagEnsureReq instantiates a new UnibeeApiMerchantUserTagEnsureReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserTagEnsureReqWithDefaults

`func NewUnibeeApiMerchantUserTagEnsureReqWithDefaults() *UnibeeApiMerchantUserTagEnsureReq`

NewUnibeeApiMerchantUserTagEnsureReqWithDefaults instantiates a new UnibeeApiMerchantUserTagEnsureReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutMs

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *UnibeeApiMerchantUserTagEnsureReq) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *UnibeeApiMerchantUserTagEnsureReq) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserTagEnsureReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetWaitForReady

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetWaitForReady() bool`

GetWaitForReady returns the WaitForReady field if non-nil, zero value otherwise.

### GetWaitForReadyOk

`func (o *UnibeeApiMerchantUserTagEnsureReq) GetWaitForReadyOk() (*bool, bool)`

GetWaitForReadyOk returns a tuple with the WaitForReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForReady

`func (o *UnibeeApiMerchantUserTagEnsureReq) SetWaitForReady(v bool)`

SetWaitForReady sets WaitForReady field to given value.

### HasWaitForReady

`func (o *UnibeeApiMerchantUserTagEnsureReq) HasWaitForReady() bool`

HasWaitForReady returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


