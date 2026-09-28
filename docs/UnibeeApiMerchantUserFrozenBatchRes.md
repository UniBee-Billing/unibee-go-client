# UnibeeApiMerchantUserFrozenBatchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedCount** | Pointer to **int32** | Failed | [optional] 
**Items** | Pointer to [**[]UnibeeApiMerchantUserFrozenBatchItem**](UnibeeApiMerchantUserFrozenBatchItem.md) | Per-user result, same order as unique userIds or emails | [optional] 
**SuccessCount** | Pointer to **int32** | Succeeded, including already suspended | [optional] 
**Total** | Pointer to **int32** | Number of unique userIds or emails processed | [optional] 

## Methods

### NewUnibeeApiMerchantUserFrozenBatchRes

`func NewUnibeeApiMerchantUserFrozenBatchRes() *UnibeeApiMerchantUserFrozenBatchRes`

NewUnibeeApiMerchantUserFrozenBatchRes instantiates a new UnibeeApiMerchantUserFrozenBatchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserFrozenBatchResWithDefaults

`func NewUnibeeApiMerchantUserFrozenBatchResWithDefaults() *UnibeeApiMerchantUserFrozenBatchRes`

NewUnibeeApiMerchantUserFrozenBatchResWithDefaults instantiates a new UnibeeApiMerchantUserFrozenBatchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetItems

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetItems() []UnibeeApiMerchantUserFrozenBatchItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetItemsOk() (*[]UnibeeApiMerchantUserFrozenBatchItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiMerchantUserFrozenBatchRes) SetItems(v []UnibeeApiMerchantUserFrozenBatchItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UnibeeApiMerchantUserFrozenBatchRes) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiMerchantUserFrozenBatchRes) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantUserFrozenBatchRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantUserFrozenBatchRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantUserFrozenBatchRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


