# UnibeeApiMerchantEmailHistoryListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailHistories** | Pointer to [**[]UnibeeApiBeanDetailMerchantEmailHistoryDetail**](UnibeeApiBeanDetailMerchantEmailHistoryDetail.md) | Email History Object List | [optional] 
**EmailHistoryStatistics** | Pointer to [**UnibeeApiBeanDetailMerchantEmailHistoryStatistics**](UnibeeApiBeanDetailMerchantEmailHistoryStatistics.md) |  | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantEmailHistoryListRes

`func NewUnibeeApiMerchantEmailHistoryListRes() *UnibeeApiMerchantEmailHistoryListRes`

NewUnibeeApiMerchantEmailHistoryListRes instantiates a new UnibeeApiMerchantEmailHistoryListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailHistoryListResWithDefaults

`func NewUnibeeApiMerchantEmailHistoryListResWithDefaults() *UnibeeApiMerchantEmailHistoryListRes`

NewUnibeeApiMerchantEmailHistoryListResWithDefaults instantiates a new UnibeeApiMerchantEmailHistoryListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailHistories

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetEmailHistories() []UnibeeApiBeanDetailMerchantEmailHistoryDetail`

GetEmailHistories returns the EmailHistories field if non-nil, zero value otherwise.

### GetEmailHistoriesOk

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetEmailHistoriesOk() (*[]UnibeeApiBeanDetailMerchantEmailHistoryDetail, bool)`

GetEmailHistoriesOk returns a tuple with the EmailHistories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHistories

`func (o *UnibeeApiMerchantEmailHistoryListRes) SetEmailHistories(v []UnibeeApiBeanDetailMerchantEmailHistoryDetail)`

SetEmailHistories sets EmailHistories field to given value.

### HasEmailHistories

`func (o *UnibeeApiMerchantEmailHistoryListRes) HasEmailHistories() bool`

HasEmailHistories returns a boolean if a field has been set.

### GetEmailHistoryStatistics

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetEmailHistoryStatistics() UnibeeApiBeanDetailMerchantEmailHistoryStatistics`

GetEmailHistoryStatistics returns the EmailHistoryStatistics field if non-nil, zero value otherwise.

### GetEmailHistoryStatisticsOk

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetEmailHistoryStatisticsOk() (*UnibeeApiBeanDetailMerchantEmailHistoryStatistics, bool)`

GetEmailHistoryStatisticsOk returns a tuple with the EmailHistoryStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHistoryStatistics

`func (o *UnibeeApiMerchantEmailHistoryListRes) SetEmailHistoryStatistics(v UnibeeApiBeanDetailMerchantEmailHistoryStatistics)`

SetEmailHistoryStatistics sets EmailHistoryStatistics field to given value.

### HasEmailHistoryStatistics

`func (o *UnibeeApiMerchantEmailHistoryListRes) HasEmailHistoryStatistics() bool`

HasEmailHistoryStatistics returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantEmailHistoryListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantEmailHistoryListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantEmailHistoryListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


