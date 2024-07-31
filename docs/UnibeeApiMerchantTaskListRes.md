# UnibeeApiMerchantTaskListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Downloads** | Pointer to [**[]UnibeeApiBeanMerchantBatchTask**](UnibeeApiBeanMerchantBatchTask.md) | Merchant Member Task List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantTaskListRes

`func NewUnibeeApiMerchantTaskListRes() *UnibeeApiMerchantTaskListRes`

NewUnibeeApiMerchantTaskListRes instantiates a new UnibeeApiMerchantTaskListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskListResWithDefaults

`func NewUnibeeApiMerchantTaskListResWithDefaults() *UnibeeApiMerchantTaskListRes`

NewUnibeeApiMerchantTaskListResWithDefaults instantiates a new UnibeeApiMerchantTaskListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloads

`func (o *UnibeeApiMerchantTaskListRes) GetDownloads() []UnibeeApiBeanMerchantBatchTask`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *UnibeeApiMerchantTaskListRes) GetDownloadsOk() (*[]UnibeeApiBeanMerchantBatchTask, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *UnibeeApiMerchantTaskListRes) SetDownloads(v []UnibeeApiBeanMerchantBatchTask)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *UnibeeApiMerchantTaskListRes) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantTaskListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantTaskListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantTaskListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantTaskListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


