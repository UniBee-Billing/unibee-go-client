# UnibeeApiBeanMerchantBatchExportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**ExportColumns** | Pointer to **[]string** | ExportColumns, the export file column list, will export all columns if not specified | [optional] 
**Format** | Pointer to **string** | The format of export file, xlsx|csv, will be xlsx if not specified | [optional] 
**MemberId** | Pointer to **int64** | member_id | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Payload** | Pointer to **map[string]map[string]interface{}** | Payload | [optional] 
**Task** | Pointer to **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | [optional] 
**TemplateId** | Pointer to **int64** | templateId | [optional] 

## Methods

### NewUnibeeApiBeanMerchantBatchExportTemplate

`func NewUnibeeApiBeanMerchantBatchExportTemplate() *UnibeeApiBeanMerchantBatchExportTemplate`

NewUnibeeApiBeanMerchantBatchExportTemplate instantiates a new UnibeeApiBeanMerchantBatchExportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchExportTemplateWithDefaults

`func NewUnibeeApiBeanMerchantBatchExportTemplateWithDefaults() *UnibeeApiBeanMerchantBatchExportTemplate`

NewUnibeeApiBeanMerchantBatchExportTemplateWithDefaults instantiates a new UnibeeApiBeanMerchantBatchExportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetExportColumns() []string`

GetExportColumns returns the ExportColumns field if non-nil, zero value otherwise.

### GetExportColumnsOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetExportColumnsOk() (*[]string, bool)`

GetExportColumnsOk returns a tuple with the ExportColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetExportColumns(v []string)`

SetExportColumns sets ExportColumns field to given value.

### HasExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasExportColumns() bool`

HasExportColumns returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetPayload() map[string]map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetPayloadOk() (*map[string]map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetPayload(v map[string]map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


