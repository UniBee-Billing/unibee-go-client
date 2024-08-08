# UnibeeApiMerchantTaskEditTemplateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportColumns** | Pointer to **[][]string** | ExportColumns, the export file column list, will export all columns if not specified | [optional] 
**Format** | Pointer to **string** | The format of export file, xlsx|csv, will be xlsx if not specified | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Payload** | Pointer to **map[string]map[string]interface{}** | Payload | [optional] 
**Task** | Pointer to **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | [optional] 
**TemplateId** | **int64** | templateId | 

## Methods

### NewUnibeeApiMerchantTaskEditTemplateReq

`func NewUnibeeApiMerchantTaskEditTemplateReq(templateId int64, ) *UnibeeApiMerchantTaskEditTemplateReq`

NewUnibeeApiMerchantTaskEditTemplateReq instantiates a new UnibeeApiMerchantTaskEditTemplateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskEditTemplateReqWithDefaults

`func NewUnibeeApiMerchantTaskEditTemplateReqWithDefaults() *UnibeeApiMerchantTaskEditTemplateReq`

NewUnibeeApiMerchantTaskEditTemplateReqWithDefaults instantiates a new UnibeeApiMerchantTaskEditTemplateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportColumns

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetExportColumns() [][]string`

GetExportColumns returns the ExportColumns field if non-nil, zero value otherwise.

### GetExportColumnsOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetExportColumnsOk() (*[][]string, bool)`

GetExportColumnsOk returns a tuple with the ExportColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportColumns

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetExportColumns(v [][]string)`

SetExportColumns sets ExportColumns field to given value.

### HasExportColumns

`func (o *UnibeeApiMerchantTaskEditTemplateReq) HasExportColumns() bool`

HasExportColumns returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiMerchantTaskEditTemplateReq) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantTaskEditTemplateReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetPayload() map[string]map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetPayloadOk() (*map[string]map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetPayload(v map[string]map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiMerchantTaskEditTemplateReq) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *UnibeeApiMerchantTaskEditTemplateReq) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiMerchantTaskEditTemplateReq) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiMerchantTaskEditTemplateReq) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


