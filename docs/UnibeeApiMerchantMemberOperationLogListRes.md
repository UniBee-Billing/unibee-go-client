# UnibeeApiMerchantMemberOperationLogListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOperationLogs** | Pointer to [**[]UnibeeApiBeanDetailMerchantOperationLogDetail**](UnibeeApiBeanDetailMerchantOperationLogDetail.md) | Merchant Member Operation Log List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantMemberOperationLogListRes

`func NewUnibeeApiMerchantMemberOperationLogListRes() *UnibeeApiMerchantMemberOperationLogListRes`

NewUnibeeApiMerchantMemberOperationLogListRes instantiates a new UnibeeApiMerchantMemberOperationLogListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberOperationLogListResWithDefaults

`func NewUnibeeApiMerchantMemberOperationLogListResWithDefaults() *UnibeeApiMerchantMemberOperationLogListRes`

NewUnibeeApiMerchantMemberOperationLogListResWithDefaults instantiates a new UnibeeApiMerchantMemberOperationLogListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOperationLogs

`func (o *UnibeeApiMerchantMemberOperationLogListRes) GetMerchantOperationLogs() []UnibeeApiBeanDetailMerchantOperationLogDetail`

GetMerchantOperationLogs returns the MerchantOperationLogs field if non-nil, zero value otherwise.

### GetMerchantOperationLogsOk

`func (o *UnibeeApiMerchantMemberOperationLogListRes) GetMerchantOperationLogsOk() (*[]UnibeeApiBeanDetailMerchantOperationLogDetail, bool)`

GetMerchantOperationLogsOk returns a tuple with the MerchantOperationLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOperationLogs

`func (o *UnibeeApiMerchantMemberOperationLogListRes) SetMerchantOperationLogs(v []UnibeeApiBeanDetailMerchantOperationLogDetail)`

SetMerchantOperationLogs sets MerchantOperationLogs field to given value.

### HasMerchantOperationLogs

`func (o *UnibeeApiMerchantMemberOperationLogListRes) HasMerchantOperationLogs() bool`

HasMerchantOperationLogs returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantMemberOperationLogListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantMemberOperationLogListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantMemberOperationLogListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantMemberOperationLogListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


