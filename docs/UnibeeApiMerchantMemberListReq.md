# UnibeeApiMerchantMemberListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**RoleIds** | Pointer to **[]int64** | The member&#39;s roleId if specified&#39; | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


