# UnibeeApiMerchantDiscountBatchTemplateListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTemplateCount** | Pointer to **int64** | Total count of active template codes | [optional] 
**Templates** | Pointer to [**[]UnibeeApiBeanMerchantBatchDiscountCodeTemplate**](UnibeeApiBeanMerchantBatchDiscountCodeTemplate.md) | Batch template list | [optional] 
**Total** | Pointer to **int32** | Total count | [optional] 
**TotalChildCodeCount** | Pointer to **int64** | Total count of all child codes | [optional] 
**UsageRate** | Pointer to **float32** | Usage rate (usedChildCodeCount / totalChildCodeCount), 0-1 | [optional] 
**UsedChildCodeCount** | Pointer to **int64** | Total count of used child codes | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchTemplateListRes

`func NewUnibeeApiMerchantDiscountBatchTemplateListRes() *UnibeeApiMerchantDiscountBatchTemplateListRes`

NewUnibeeApiMerchantDiscountBatchTemplateListRes instantiates a new UnibeeApiMerchantDiscountBatchTemplateListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchTemplateListResWithDefaults

`func NewUnibeeApiMerchantDiscountBatchTemplateListResWithDefaults() *UnibeeApiMerchantDiscountBatchTemplateListRes`

NewUnibeeApiMerchantDiscountBatchTemplateListResWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchTemplateListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTemplateCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetActiveTemplateCount() int64`

GetActiveTemplateCount returns the ActiveTemplateCount field if non-nil, zero value otherwise.

### GetActiveTemplateCountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetActiveTemplateCountOk() (*int64, bool)`

GetActiveTemplateCountOk returns a tuple with the ActiveTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTemplateCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetActiveTemplateCount(v int64)`

SetActiveTemplateCount sets ActiveTemplateCount field to given value.

### HasActiveTemplateCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasActiveTemplateCount() bool`

HasActiveTemplateCount returns a boolean if a field has been set.

### GetTemplates

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTemplates() []UnibeeApiBeanMerchantBatchDiscountCodeTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTemplatesOk() (*[]UnibeeApiBeanMerchantBatchDiscountCodeTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetTemplates(v []UnibeeApiBeanMerchantBatchDiscountCodeTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTotalChildCodeCount() int64`

GetTotalChildCodeCount returns the TotalChildCodeCount field if non-nil, zero value otherwise.

### GetTotalChildCodeCountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetTotalChildCodeCountOk() (*int64, bool)`

GetTotalChildCodeCountOk returns a tuple with the TotalChildCodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetTotalChildCodeCount(v int64)`

SetTotalChildCodeCount sets TotalChildCodeCount field to given value.

### HasTotalChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasTotalChildCodeCount() bool`

HasTotalChildCodeCount returns a boolean if a field has been set.

### GetUsageRate

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetUsageRate() float32`

GetUsageRate returns the UsageRate field if non-nil, zero value otherwise.

### GetUsageRateOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetUsageRateOk() (*float32, bool)`

GetUsageRateOk returns a tuple with the UsageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRate

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetUsageRate(v float32)`

SetUsageRate sets UsageRate field to given value.

### HasUsageRate

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasUsageRate() bool`

HasUsageRate returns a boolean if a field has been set.

### GetUsedChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetUsedChildCodeCount() int64`

GetUsedChildCodeCount returns the UsedChildCodeCount field if non-nil, zero value otherwise.

### GetUsedChildCodeCountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) GetUsedChildCodeCountOk() (*int64, bool)`

GetUsedChildCodeCountOk returns a tuple with the UsedChildCodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) SetUsedChildCodeCount(v int64)`

SetUsedChildCodeCount sets UsedChildCodeCount field to given value.

### HasUsedChildCodeCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListRes) HasUsedChildCodeCount() bool`

HasUsedChildCodeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


