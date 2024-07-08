# UnibeeApiMerchantTaskNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **map[string]string** | Payload | [optional] 
**SkipColumnIndexes** | Pointer to **[]int32** | SkipColumnIndexes, the column will be skipped in the export file if its index specified, start from 0 | [optional] 
**Task** | Pointer to **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | [optional] 

## Methods

### NewUnibeeApiMerchantTaskNewReq

`func NewUnibeeApiMerchantTaskNewReq() *UnibeeApiMerchantTaskNewReq`

NewUnibeeApiMerchantTaskNewReq instantiates a new UnibeeApiMerchantTaskNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskNewReqWithDefaults

`func NewUnibeeApiMerchantTaskNewReqWithDefaults() *UnibeeApiMerchantTaskNewReq`

NewUnibeeApiMerchantTaskNewReqWithDefaults instantiates a new UnibeeApiMerchantTaskNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *UnibeeApiMerchantTaskNewReq) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiMerchantTaskNewReq) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiMerchantTaskNewReq) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiMerchantTaskNewReq) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSkipColumnIndexes

`func (o *UnibeeApiMerchantTaskNewReq) GetSkipColumnIndexes() []int32`

GetSkipColumnIndexes returns the SkipColumnIndexes field if non-nil, zero value otherwise.

### GetSkipColumnIndexesOk

`func (o *UnibeeApiMerchantTaskNewReq) GetSkipColumnIndexesOk() (*[]int32, bool)`

GetSkipColumnIndexesOk returns a tuple with the SkipColumnIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipColumnIndexes

`func (o *UnibeeApiMerchantTaskNewReq) SetSkipColumnIndexes(v []int32)`

SetSkipColumnIndexes sets SkipColumnIndexes field to given value.

### HasSkipColumnIndexes

`func (o *UnibeeApiMerchantTaskNewReq) HasSkipColumnIndexes() bool`

HasSkipColumnIndexes returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiMerchantTaskNewReq) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiMerchantTaskNewReq) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiMerchantTaskNewReq) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *UnibeeApiMerchantTaskNewReq) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


