# UnibeeApiBeanDetailUserAdminNoteDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | CreateTime, UTC Time | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**MerchantMemberId** | Pointer to **int64** | merchant_user_id | [optional] 
**Note** | Pointer to **string** | note | [optional] 
**UserAccount** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserAdminNoteDetail

`func NewUnibeeApiBeanDetailUserAdminNoteDetail() *UnibeeApiBeanDetailUserAdminNoteDetail`

NewUnibeeApiBeanDetailUserAdminNoteDetail instantiates a new UnibeeApiBeanDetailUserAdminNoteDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserAdminNoteDetailWithDefaults

`func NewUnibeeApiBeanDetailUserAdminNoteDetailWithDefaults() *UnibeeApiBeanDetailUserAdminNoteDetail`

NewUnibeeApiBeanDetailUserAdminNoteDetailWithDefaults instantiates a new UnibeeApiBeanDetailUserAdminNoteDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantMember

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetMerchantMemberId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberId() int64`

GetMerchantMemberId returns the MerchantMemberId field if non-nil, zero value otherwise.

### GetMerchantMemberIdOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetMerchantMemberIdOk() (*int64, bool)`

GetMerchantMemberIdOk returns a tuple with the MerchantMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetMerchantMemberId(v int64)`

SetMerchantMemberId sets MerchantMemberId field to given value.

### HasMerchantMemberId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasMerchantMemberId() bool`

HasMerchantMemberId returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUserAccount

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserAccount() UnibeeApiBeanUserAccount`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserAccountOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetUserAccount(v UnibeeApiBeanUserAccount)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailUserAdminNoteDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


