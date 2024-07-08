# MerchantMemberListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMembers** | Pointer to [**[]UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) | Merchant Member Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantMemberListGet200ResponseData

`func NewMerchantMemberListGet200ResponseData() *MerchantMemberListGet200ResponseData`

NewMerchantMemberListGet200ResponseData instantiates a new MerchantMemberListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantMemberListGet200ResponseDataWithDefaults

`func NewMerchantMemberListGet200ResponseDataWithDefaults() *MerchantMemberListGet200ResponseData`

NewMerchantMemberListGet200ResponseDataWithDefaults instantiates a new MerchantMemberListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMembers

`func (o *MerchantMemberListGet200ResponseData) GetMerchantMembers() []UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMembers returns the MerchantMembers field if non-nil, zero value otherwise.

### GetMerchantMembersOk

`func (o *MerchantMemberListGet200ResponseData) GetMerchantMembersOk() (*[]UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMembersOk returns a tuple with the MerchantMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMembers

`func (o *MerchantMemberListGet200ResponseData) SetMerchantMembers(v []UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMembers sets MerchantMembers field to given value.

### HasMerchantMembers

`func (o *MerchantMemberListGet200ResponseData) HasMerchantMembers() bool`

HasMerchantMembers returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantMemberListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantMemberListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantMemberListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantMemberListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


