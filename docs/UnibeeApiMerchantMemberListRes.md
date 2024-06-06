# UnibeeApiMerchantMemberListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMembers** | Pointer to [**[]UnibeeApiBeanMerchantMemberSimplify**](UnibeeApiBeanMerchantMemberSimplify.md) | Merchant Member Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantMemberListRes

`func NewUnibeeApiMerchantMemberListRes() *UnibeeApiMerchantMemberListRes`

NewUnibeeApiMerchantMemberListRes instantiates a new UnibeeApiMerchantMemberListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberListResWithDefaults

`func NewUnibeeApiMerchantMemberListResWithDefaults() *UnibeeApiMerchantMemberListRes`

NewUnibeeApiMerchantMemberListResWithDefaults instantiates a new UnibeeApiMerchantMemberListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMembers

`func (o *UnibeeApiMerchantMemberListRes) GetMerchantMembers() []UnibeeApiBeanMerchantMemberSimplify`

GetMerchantMembers returns the MerchantMembers field if non-nil, zero value otherwise.

### GetMerchantMembersOk

`func (o *UnibeeApiMerchantMemberListRes) GetMerchantMembersOk() (*[]UnibeeApiBeanMerchantMemberSimplify, bool)`

GetMerchantMembersOk returns a tuple with the MerchantMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMembers

`func (o *UnibeeApiMerchantMemberListRes) SetMerchantMembers(v []UnibeeApiBeanMerchantMemberSimplify)`

SetMerchantMembers sets MerchantMembers field to given value.

### HasMerchantMembers

`func (o *UnibeeApiMerchantMemberListRes) HasMerchantMembers() bool`

HasMerchantMembers returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantMemberListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantMemberListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantMemberListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantMemberListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


