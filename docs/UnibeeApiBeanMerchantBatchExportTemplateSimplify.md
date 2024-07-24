# UnibeeApiBeanMerchantBatchExportTemplateSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**ExportColumns** | Pointer to **[]string** | ExportColumns, the export file column list, will export all columns if not specified | [optional] 
**Format** | Pointer to **string** | The format of export file, xlsx|csv, will be xlsx if not specified | [optional] 
**MemberId** | Pointer to **int64** | member_id | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Payload** | Pointer to **map[string]string** | Payload | [optional] 
**Task** | Pointer to **string** | Task,InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | [optional] 
**TemplateId** | Pointer to **int64** | templateId | [optional] 

## Methods

### NewUnibeeApiBeanMerchantBatchExportTemplateSimplify

`func NewUnibeeApiBeanMerchantBatchExportTemplateSimplify() *UnibeeApiBeanMerchantBatchExportTemplateSimplify`

NewUnibeeApiBeanMerchantBatchExportTemplateSimplify instantiates a new UnibeeApiBeanMerchantBatchExportTemplateSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchExportTemplateSimplifyWithDefaults

`func NewUnibeeApiBeanMerchantBatchExportTemplateSimplifyWithDefaults() *UnibeeApiBeanMerchantBatchExportTemplateSimplify`

NewUnibeeApiBeanMerchantBatchExportTemplateSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantBatchExportTemplateSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetExportColumns() []string`

GetExportColumns returns the ExportColumns field if non-nil, zero value otherwise.

### GetExportColumnsOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetExportColumnsOk() (*[]string, bool)`

GetExportColumnsOk returns a tuple with the ExportColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetExportColumns(v []string)`

SetExportColumns sets ExportColumns field to given value.

### HasExportColumns

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasExportColumns() bool`

HasExportColumns returns a boolean if a field has been set.

### GetFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UnibeeApiBeanMerchantBatchExportTemplateSimplify) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


