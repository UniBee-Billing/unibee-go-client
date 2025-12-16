# UnibeeApiMerchantDiscountBatchChildrenListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Filter by child code, fuzzy search | [optional] 
**Count** | Pointer to **int32** | Count per page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | Filter by create time end, UTC timestamp in seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | Filter by create time start, UTC timestamp in seconds | [optional] 
**Page** | Pointer to **int32** | Page number, start from 0 | [optional] 
**SortField** | Pointer to **string** | Sort field, gmt_create|gmt_modify, default gmt_create | [optional] 
**SortType** | Pointer to **string** | Sort type, asc|desc, default desc | [optional] 
**TemplateId** | **int64** | The template&#39;s Id | 

## Methods

### NewUnibeeApiMerchantDiscountBatchChildrenListReq

`func NewUnibeeApiMerchantDiscountBatchChildrenListReq(templateId int64, ) *UnibeeApiMerchantDiscountBatchChildrenListReq`

NewUnibeeApiMerchantDiscountBatchChildrenListReq instantiates a new UnibeeApiMerchantDiscountBatchChildrenListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchChildrenListReqWithDefaults

`func NewUnibeeApiMerchantDiscountBatchChildrenListReqWithDefaults() *UnibeeApiMerchantDiscountBatchChildrenListReq`

NewUnibeeApiMerchantDiscountBatchChildrenListReqWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchChildrenListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenListReq) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


