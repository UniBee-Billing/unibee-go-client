# UnibeeApiMerchantMemberUpdateMemberRoleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **int64** | The unique id of member | [optional] 
**Role** | Pointer to **string** | The permission role of member | [optional] 

## Methods

### NewUnibeeApiMerchantMemberUpdateMemberRoleReq

`func NewUnibeeApiMerchantMemberUpdateMemberRoleReq() *UnibeeApiMerchantMemberUpdateMemberRoleReq`

NewUnibeeApiMerchantMemberUpdateMemberRoleReq instantiates a new UnibeeApiMerchantMemberUpdateMemberRoleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberUpdateMemberRoleReqWithDefaults

`func NewUnibeeApiMerchantMemberUpdateMemberRoleReqWithDefaults() *UnibeeApiMerchantMemberUpdateMemberRoleReq`

NewUnibeeApiMerchantMemberUpdateMemberRoleReqWithDefaults instantiates a new UnibeeApiMerchantMemberUpdateMemberRoleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetRole

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


