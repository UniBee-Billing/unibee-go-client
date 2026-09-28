# UnibeeApiSystemUserTagRebuildReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **int64** |  | 
**UserId** | Pointer to **int64** | Single user id shortcut | [optional] 
**UserIds** | Pointer to **[]int64** | User ids to rebuild; required unless omitted for empty | [optional] 

## Methods

### NewUnibeeApiSystemUserTagRebuildReq

`func NewUnibeeApiSystemUserTagRebuildReq(merchantId int64, ) *UnibeeApiSystemUserTagRebuildReq`

NewUnibeeApiSystemUserTagRebuildReq instantiates a new UnibeeApiSystemUserTagRebuildReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemUserTagRebuildReqWithDefaults

`func NewUnibeeApiSystemUserTagRebuildReqWithDefaults() *UnibeeApiSystemUserTagRebuildReq`

NewUnibeeApiSystemUserTagRebuildReqWithDefaults instantiates a new UnibeeApiSystemUserTagRebuildReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *UnibeeApiSystemUserTagRebuildReq) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiSystemUserTagRebuildReq) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiSystemUserTagRebuildReq) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.


### GetUserId

`func (o *UnibeeApiSystemUserTagRebuildReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiSystemUserTagRebuildReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiSystemUserTagRebuildReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiSystemUserTagRebuildReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIds

`func (o *UnibeeApiSystemUserTagRebuildReq) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UnibeeApiSystemUserTagRebuildReq) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UnibeeApiSystemUserTagRebuildReq) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UnibeeApiSystemUserTagRebuildReq) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


