# MerchantUserChurnListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChurnRecords** | Pointer to [**[]UnibeeApiBeanUserChurnHistory**](UnibeeApiBeanUserChurnHistory.md) | Churn episode list | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantUserChurnListGet200ResponseData

`func NewMerchantUserChurnListGet200ResponseData() *MerchantUserChurnListGet200ResponseData`

NewMerchantUserChurnListGet200ResponseData instantiates a new MerchantUserChurnListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantUserChurnListGet200ResponseDataWithDefaults

`func NewMerchantUserChurnListGet200ResponseDataWithDefaults() *MerchantUserChurnListGet200ResponseData`

NewMerchantUserChurnListGet200ResponseDataWithDefaults instantiates a new MerchantUserChurnListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChurnRecords

`func (o *MerchantUserChurnListGet200ResponseData) GetChurnRecords() []UnibeeApiBeanUserChurnHistory`

GetChurnRecords returns the ChurnRecords field if non-nil, zero value otherwise.

### GetChurnRecordsOk

`func (o *MerchantUserChurnListGet200ResponseData) GetChurnRecordsOk() (*[]UnibeeApiBeanUserChurnHistory, bool)`

GetChurnRecordsOk returns a tuple with the ChurnRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnRecords

`func (o *MerchantUserChurnListGet200ResponseData) SetChurnRecords(v []UnibeeApiBeanUserChurnHistory)`

SetChurnRecords sets ChurnRecords field to given value.

### HasChurnRecords

`func (o *MerchantUserChurnListGet200ResponseData) HasChurnRecords() bool`

HasChurnRecords returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantUserChurnListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantUserChurnListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantUserChurnListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantUserChurnListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


