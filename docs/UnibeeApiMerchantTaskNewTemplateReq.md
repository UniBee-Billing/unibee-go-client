# UnibeeApiMerchantTaskNewTemplateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportColumns** | Pointer to **[]string** | ExportColumns, the export file column list, will export all columns if not specified, first char should lower case | [optional] 
**Format** | Pointer to **string** | The format of export file, xlsx|csv, will be xlsx if not specified | [optional] 
**Name** | **string** | name | 
**Payload** | Pointer to **map[string]string** | Payload | [optional] 
**Task** | **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | 

## Methods

### NewUnibeeApiMerchantTaskNewTemplateReq

`func NewUnibeeApiMerchantTaskNewTemplateReq(name string, task string, ) *UnibeeApiMerchantTaskNewTemplateReq`

NewUnibeeApiMerchantTaskNewTemplateReq instantiates a new UnibeeApiMerchantTaskNewTemplateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskNewTemplateReqWithDefaults

`func NewUnibeeApiMerchantTaskNewTemplateReqWithDefaults() *UnibeeApiMerchantTaskNewTemplateReq`

NewUnibeeApiMerchantTaskNewTemplateReqWithDefaults instantiates a new UnibeeApiMerchantTaskNewTemplateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportColumns

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetExportColumns() []string`

GetExportColumns returns the ExportColumns field if non-nil, zero value otherwise.

### GetExportColumnsOk

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetExportColumnsOk() (*[]string, bool)`

GetExportColumnsOk returns a tuple with the ExportColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportColumns

`func (o *UnibeeApiMerchantTaskNewTemplateReq) SetExportColumns(v []string)`

SetExportColumns sets ExportColumns field to given value.

### HasExportColumns

`func (o *UnibeeApiMerchantTaskNewTemplateReq) HasExportColumns() bool`

HasExportColumns returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiMerchantTaskNewTemplateReq) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiMerchantTaskNewTemplateReq) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantTaskNewTemplateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiMerchantTaskNewTemplateReq) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiMerchantTaskNewTemplateReq) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiMerchantTaskNewTemplateReq) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiMerchantTaskNewTemplateReq) SetTask(v string)`

SetTask sets Task field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


