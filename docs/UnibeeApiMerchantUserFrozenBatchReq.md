# UnibeeApiMerchantUserFrozenBatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **string** | Email list of users, split by commas or semicolons, max 100. Provide either userIds or emails, not both | [optional] 
**UserIds** | Pointer to **[]int64** | UserIds, max 100. Provide either userIds or emails, not both | [optional] 

## Methods

### NewUnibeeApiMerchantUserFrozenBatchReq

`func NewUnibeeApiMerchantUserFrozenBatchReq() *UnibeeApiMerchantUserFrozenBatchReq`

NewUnibeeApiMerchantUserFrozenBatchReq instantiates a new UnibeeApiMerchantUserFrozenBatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserFrozenBatchReqWithDefaults

`func NewUnibeeApiMerchantUserFrozenBatchReqWithDefaults() *UnibeeApiMerchantUserFrozenBatchReq`

NewUnibeeApiMerchantUserFrozenBatchReqWithDefaults instantiates a new UnibeeApiMerchantUserFrozenBatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UnibeeApiMerchantUserFrozenBatchReq) GetEmails() string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UnibeeApiMerchantUserFrozenBatchReq) GetEmailsOk() (*string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UnibeeApiMerchantUserFrozenBatchReq) SetEmails(v string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UnibeeApiMerchantUserFrozenBatchReq) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetUserIds

`func (o *UnibeeApiMerchantUserFrozenBatchReq) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UnibeeApiMerchantUserFrozenBatchReq) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UnibeeApiMerchantUserFrozenBatchReq) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UnibeeApiMerchantUserFrozenBatchReq) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


