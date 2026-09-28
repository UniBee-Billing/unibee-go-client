# UnibeeApiMerchantUserChurnBulkCheckRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentlyChurnedCount** | Pointer to **int32** | Found users with no Active/Incomplete UniBee Control subscription | [optional] 
**FoundCount** | Pointer to **int32** | Emails that match a merchant user | [optional] 
**Items** | Pointer to [**[]UnibeeApiBeanUserChurnBulkCheckItem**](UnibeeApiBeanUserChurnBulkCheckItem.md) | One item per input email, same order as emails | [optional] 
**NotFoundCount** | Pointer to **int32** | Emails with no merchant user | [optional] 
**Total** | Pointer to **int32** | Number of input emails after parse/dedupe | [optional] 

## Methods

### NewUnibeeApiMerchantUserChurnBulkCheckRes

`func NewUnibeeApiMerchantUserChurnBulkCheckRes() *UnibeeApiMerchantUserChurnBulkCheckRes`

NewUnibeeApiMerchantUserChurnBulkCheckRes instantiates a new UnibeeApiMerchantUserChurnBulkCheckRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserChurnBulkCheckResWithDefaults

`func NewUnibeeApiMerchantUserChurnBulkCheckResWithDefaults() *UnibeeApiMerchantUserChurnBulkCheckRes`

NewUnibeeApiMerchantUserChurnBulkCheckResWithDefaults instantiates a new UnibeeApiMerchantUserChurnBulkCheckRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentlyChurnedCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetCurrentlyChurnedCount() int32`

GetCurrentlyChurnedCount returns the CurrentlyChurnedCount field if non-nil, zero value otherwise.

### GetCurrentlyChurnedCountOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetCurrentlyChurnedCountOk() (*int32, bool)`

GetCurrentlyChurnedCountOk returns a tuple with the CurrentlyChurnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyChurnedCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) SetCurrentlyChurnedCount(v int32)`

SetCurrentlyChurnedCount sets CurrentlyChurnedCount field to given value.

### HasCurrentlyChurnedCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) HasCurrentlyChurnedCount() bool`

HasCurrentlyChurnedCount returns a boolean if a field has been set.

### GetFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetFoundCount() int32`

GetFoundCount returns the FoundCount field if non-nil, zero value otherwise.

### GetFoundCountOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetFoundCountOk() (*int32, bool)`

GetFoundCountOk returns a tuple with the FoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) SetFoundCount(v int32)`

SetFoundCount sets FoundCount field to given value.

### HasFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) HasFoundCount() bool`

HasFoundCount returns a boolean if a field has been set.

### GetItems

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetItems() []UnibeeApiBeanUserChurnBulkCheckItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetItemsOk() (*[]UnibeeApiBeanUserChurnBulkCheckItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) SetItems(v []UnibeeApiBeanUserChurnBulkCheckItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNotFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetNotFoundCount() int32`

GetNotFoundCount returns the NotFoundCount field if non-nil, zero value otherwise.

### GetNotFoundCountOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetNotFoundCountOk() (*int32, bool)`

GetNotFoundCountOk returns a tuple with the NotFoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) SetNotFoundCount(v int32)`

SetNotFoundCount sets NotFoundCount field to given value.

### HasNotFoundCount

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) HasNotFoundCount() bool`

HasNotFoundCount returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantUserChurnBulkCheckRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


