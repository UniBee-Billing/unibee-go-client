# UnibeeApiMerchantUserAdminNoteListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**UserId** | **int64** | The id of user, either ExternalUserId or UserId needed | 

## Methods

### NewUnibeeApiMerchantUserAdminNoteListReq

`func NewUnibeeApiMerchantUserAdminNoteListReq(userId int64, ) *UnibeeApiMerchantUserAdminNoteListReq`

NewUnibeeApiMerchantUserAdminNoteListReq instantiates a new UnibeeApiMerchantUserAdminNoteListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserAdminNoteListReqWithDefaults

`func NewUnibeeApiMerchantUserAdminNoteListReqWithDefaults() *UnibeeApiMerchantUserAdminNoteListReq`

NewUnibeeApiMerchantUserAdminNoteListReqWithDefaults instantiates a new UnibeeApiMerchantUserAdminNoteListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantUserAdminNoteListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantUserAdminNoteListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantUserAdminNoteListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantUserAdminNoteListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserAdminNoteListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserAdminNoteListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


