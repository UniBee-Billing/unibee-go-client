# UnibeeApiMerchantMemberListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Email** | Pointer to **string** | Search Filter Email | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**RoleIds** | Pointer to **[]int64** | The member roleId if specified&#39; | [optional] 
**SearchKey** | Pointer to **string** | Search Key, FirstName,LastName or Email | [optional] 

## Methods

### NewUnibeeApiMerchantMemberListReq

`func NewUnibeeApiMerchantMemberListReq() *UnibeeApiMerchantMemberListReq`

NewUnibeeApiMerchantMemberListReq instantiates a new UnibeeApiMerchantMemberListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberListReqWithDefaults

`func NewUnibeeApiMerchantMemberListReqWithDefaults() *UnibeeApiMerchantMemberListReq`

NewUnibeeApiMerchantMemberListReqWithDefaults instantiates a new UnibeeApiMerchantMemberListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantMemberListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantMemberListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantMemberListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantMemberListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantMemberListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantMemberListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantMemberListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantMemberListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantMemberListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantMemberListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantMemberListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantMemberListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantMemberListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMemberListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMemberListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMemberListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantMemberListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantMemberListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantMemberListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantMemberListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRoleIds

`func (o *UnibeeApiMerchantMemberListReq) GetRoleIds() []int64`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UnibeeApiMerchantMemberListReq) GetRoleIdsOk() (*[]int64, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UnibeeApiMerchantMemberListReq) SetRoleIds(v []int64)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UnibeeApiMerchantMemberListReq) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantMemberListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantMemberListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantMemberListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantMemberListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


