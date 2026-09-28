# MerchantUserChurnBulkCheckPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentlyChurnedCount** | Pointer to **int32** | Found users with no Active/Incomplete UniBee Control subscription | [optional] 
**FoundCount** | Pointer to **int32** | Emails that match a merchant user | [optional] 
**Items** | Pointer to [**[]UnibeeApiBeanUserChurnBulkCheckItem**](UnibeeApiBeanUserChurnBulkCheckItem.md) | One item per input email, same order as emails | [optional] 
**NotFoundCount** | Pointer to **int32** | Emails with no merchant user | [optional] 
**Total** | Pointer to **int32** | Number of input emails after parse/dedupe | [optional] 

## Methods

### NewMerchantUserChurnBulkCheckPost200ResponseData

`func NewMerchantUserChurnBulkCheckPost200ResponseData() *MerchantUserChurnBulkCheckPost200ResponseData`

NewMerchantUserChurnBulkCheckPost200ResponseData instantiates a new MerchantUserChurnBulkCheckPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantUserChurnBulkCheckPost200ResponseDataWithDefaults

`func NewMerchantUserChurnBulkCheckPost200ResponseDataWithDefaults() *MerchantUserChurnBulkCheckPost200ResponseData`

NewMerchantUserChurnBulkCheckPost200ResponseDataWithDefaults instantiates a new MerchantUserChurnBulkCheckPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentlyChurnedCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetCurrentlyChurnedCount() int32`

GetCurrentlyChurnedCount returns the CurrentlyChurnedCount field if non-nil, zero value otherwise.

### GetCurrentlyChurnedCountOk

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetCurrentlyChurnedCountOk() (*int32, bool)`

GetCurrentlyChurnedCountOk returns a tuple with the CurrentlyChurnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyChurnedCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) SetCurrentlyChurnedCount(v int32)`

SetCurrentlyChurnedCount sets CurrentlyChurnedCount field to given value.

### HasCurrentlyChurnedCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) HasCurrentlyChurnedCount() bool`

HasCurrentlyChurnedCount returns a boolean if a field has been set.

### GetFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetFoundCount() int32`

GetFoundCount returns the FoundCount field if non-nil, zero value otherwise.

### GetFoundCountOk

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetFoundCountOk() (*int32, bool)`

GetFoundCountOk returns a tuple with the FoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) SetFoundCount(v int32)`

SetFoundCount sets FoundCount field to given value.

### HasFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) HasFoundCount() bool`

HasFoundCount returns a boolean if a field has been set.

### GetItems

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetItems() []UnibeeApiBeanUserChurnBulkCheckItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetItemsOk() (*[]UnibeeApiBeanUserChurnBulkCheckItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) SetItems(v []UnibeeApiBeanUserChurnBulkCheckItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNotFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetNotFoundCount() int32`

GetNotFoundCount returns the NotFoundCount field if non-nil, zero value otherwise.

### GetNotFoundCountOk

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetNotFoundCountOk() (*int32, bool)`

GetNotFoundCountOk returns a tuple with the NotFoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) SetNotFoundCount(v int32)`

SetNotFoundCount sets NotFoundCount field to given value.

### HasNotFoundCount

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) HasNotFoundCount() bool`

HasNotFoundCount returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantUserChurnBulkCheckPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


