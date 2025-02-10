# UnibeeApiMerchantUserNewAdminNoteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | **string** | Note | 
**UserId** | **int64** | The id of user, either ExternalUserId or UserId needed | 

## Methods

### NewUnibeeApiMerchantUserNewAdminNoteReq

`func NewUnibeeApiMerchantUserNewAdminNoteReq(note string, userId int64, ) *UnibeeApiMerchantUserNewAdminNoteReq`

NewUnibeeApiMerchantUserNewAdminNoteReq instantiates a new UnibeeApiMerchantUserNewAdminNoteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserNewAdminNoteReqWithDefaults

`func NewUnibeeApiMerchantUserNewAdminNoteReqWithDefaults() *UnibeeApiMerchantUserNewAdminNoteReq`

NewUnibeeApiMerchantUserNewAdminNoteReqWithDefaults instantiates a new UnibeeApiMerchantUserNewAdminNoteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) SetNote(v string)`

SetNote sets Note field to given value.


### GetUserId

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserNewAdminNoteReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


