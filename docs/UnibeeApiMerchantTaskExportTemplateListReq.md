# UnibeeApiMerchantTaskExportTemplateListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**Task** | Pointer to **string** | Filter Task, Optional, InvoiceExport|UserExport|SubscriptionExport|TransactionExport|DiscountExport|UserDiscountExport | [optional] 

## Methods

### NewUnibeeApiMerchantTaskExportTemplateListReq

`func NewUnibeeApiMerchantTaskExportTemplateListReq() *UnibeeApiMerchantTaskExportTemplateListReq`

NewUnibeeApiMerchantTaskExportTemplateListReq instantiates a new UnibeeApiMerchantTaskExportTemplateListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTaskExportTemplateListReqWithDefaults

`func NewUnibeeApiMerchantTaskExportTemplateListReqWithDefaults() *UnibeeApiMerchantTaskExportTemplateListReq`

NewUnibeeApiMerchantTaskExportTemplateListReqWithDefaults instantiates a new UnibeeApiMerchantTaskExportTemplateListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTask

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *UnibeeApiMerchantTaskExportTemplateListReq) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


