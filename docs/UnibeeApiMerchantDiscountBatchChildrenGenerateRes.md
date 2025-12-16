# UnibeeApiMerchantDiscountBatchChildrenGenerateRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingCount** | Pointer to **int64** | Already generated count before this operation | [optional] 
**FailedCount** | Pointer to **int32** | Failed count in this operation (only for synchronous execution) | [optional] 
**GeneratedCount** | Pointer to **int32** | Newly generated count in this operation (only for synchronous execution) | [optional] 
**IsAsync** | Pointer to **bool** | Whether the generation is running asynchronously | [optional] 
**SuccessCount** | Pointer to **int32** | Successfully inserted count in this operation (only for synchronous execution) | [optional] 
**TemplateCode** | Pointer to **string** | Template code prefix | [optional] 
**TemplateId** | Pointer to **int64** | Template ID | [optional] 
**TotalTarget** | Pointer to **int64** | Target total quantity | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchChildrenGenerateRes

`func NewUnibeeApiMerchantDiscountBatchChildrenGenerateRes() *UnibeeApiMerchantDiscountBatchChildrenGenerateRes`

NewUnibeeApiMerchantDiscountBatchChildrenGenerateRes instantiates a new UnibeeApiMerchantDiscountBatchChildrenGenerateRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchChildrenGenerateResWithDefaults

`func NewUnibeeApiMerchantDiscountBatchChildrenGenerateResWithDefaults() *UnibeeApiMerchantDiscountBatchChildrenGenerateRes`

NewUnibeeApiMerchantDiscountBatchChildrenGenerateResWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchChildrenGenerateRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetExistingCount() int64`

GetExistingCount returns the ExistingCount field if non-nil, zero value otherwise.

### GetExistingCountOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetExistingCountOk() (*int64, bool)`

GetExistingCountOk returns a tuple with the ExistingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetExistingCount(v int64)`

SetExistingCount sets ExistingCount field to given value.

### HasExistingCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasExistingCount() bool`

HasExistingCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetGeneratedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetGeneratedCount() int32`

GetGeneratedCount returns the GeneratedCount field if non-nil, zero value otherwise.

### GetGeneratedCountOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetGeneratedCountOk() (*int32, bool)`

GetGeneratedCountOk returns a tuple with the GeneratedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetGeneratedCount(v int32)`

SetGeneratedCount sets GeneratedCount field to given value.

### HasGeneratedCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasGeneratedCount() bool`

HasGeneratedCount returns a boolean if a field has been set.

### GetIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetIsAsync() bool`

GetIsAsync returns the IsAsync field if non-nil, zero value otherwise.

### GetIsAsyncOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetIsAsyncOk() (*bool, bool)`

GetIsAsyncOk returns a tuple with the IsAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetIsAsync(v bool)`

SetIsAsync sets IsAsync field to given value.

### HasIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasIsAsync() bool`

HasIsAsync returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTotalTarget

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTotalTarget() int64`

GetTotalTarget returns the TotalTarget field if non-nil, zero value otherwise.

### GetTotalTargetOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) GetTotalTargetOk() (*int64, bool)`

GetTotalTargetOk returns a tuple with the TotalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTarget

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) SetTotalTarget(v int64)`

SetTotalTarget sets TotalTarget field to given value.

### HasTotalTarget

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateRes) HasTotalTarget() bool`

HasTotalTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


