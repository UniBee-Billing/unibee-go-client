# UnibeeApiMerchantUserFrozenBatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email from the request when suspending by emails, or the resolved account email when suspending by userIds | [optional] 
**Message** | Pointer to **string** | Empty on newly suspended; already suspend on idempotent success; failure reason otherwise | [optional] 
**Success** | Pointer to **bool** | Whether this user was suspended or already suspended | [optional] 
**UserId** | Pointer to **int64** | UserId, 0 when the email was not found | [optional] 

## Methods

### NewUnibeeApiMerchantUserFrozenBatchItem

`func NewUnibeeApiMerchantUserFrozenBatchItem() *UnibeeApiMerchantUserFrozenBatchItem`

NewUnibeeApiMerchantUserFrozenBatchItem instantiates a new UnibeeApiMerchantUserFrozenBatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserFrozenBatchItemWithDefaults

`func NewUnibeeApiMerchantUserFrozenBatchItemWithDefaults() *UnibeeApiMerchantUserFrozenBatchItem`

NewUnibeeApiMerchantUserFrozenBatchItemWithDefaults instantiates a new UnibeeApiMerchantUserFrozenBatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserFrozenBatchItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantUserFrozenBatchItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMessage

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnibeeApiMerchantUserFrozenBatchItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnibeeApiMerchantUserFrozenBatchItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UnibeeApiMerchantUserFrozenBatchItem) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UnibeeApiMerchantUserFrozenBatchItem) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserFrozenBatchItem) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserFrozenBatchItem) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserFrozenBatchItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


