# UnibeeApiMerchantTaskNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportColumns** | Pointer to **[]string** | ExportColumns, the export file column list, will export all columns if not specified | [optional] 
**Format** | Pointer to **string** | The format of export file, xlsx|csv, will be xlsx if not specified | [optional] 
**Payload** | Pointer to **map[string]map[string]interface{}** | Payload | [optional] 
**Task** | **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | 

## Methods

### NewUnibeeApiMerchantTaskNewReq

`func NewUnibeeApiMerchantTaskNewReq(task string, ) *UnibeeApiMerchantTaskNewReq`

NewUnibeeApiMerchantTaskNewReq instantiates a new UnibeeApiMerchantTaskNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskNewReqWithDefaults

`func NewUnibeeApiMerchantTaskNewReqWithDefaults() *UnibeeApiMerchantTaskNewReq`

NewUnibeeApiMerchantTaskNewReqWithDefaults instantiates a new UnibeeApiMerchantTaskNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportColumns

`func (o *UnibeeApiMerchantTaskNewReq) GetExportColumns() []string`

GetExportColumns returns the ExportColumns field if non-nil, zero value otherwise.

### GetExportColumnsOk

`func (o *UnibeeApiMerchantTaskNewReq) GetExportColumnsOk() (*[]string, bool)`

GetExportColumnsOk returns a tuple with the ExportColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportColumns

`func (o *UnibeeApiMerchantTaskNewReq) SetExportColumns(v []string)`

SetExportColumns sets ExportColumns field to given value.

### HasExportColumns

`func (o *UnibeeApiMerchantTaskNewReq) HasExportColumns() bool`

HasExportColumns returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiMerchantTaskNewReq) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiMerchantTaskNewReq) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiMerchantTaskNewReq) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiMerchantTaskNewReq) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiMerchantTaskNewReq) GetPayload() map[string]map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiMerchantTaskNewReq) GetPayloadOk() (*map[string]map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiMerchantTaskNewReq) SetPayload(v map[string]map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiMerchantTaskNewReq) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


