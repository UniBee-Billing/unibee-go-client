# UnibeeApiMerchantDiscountBatchTemplateListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingType** | Pointer to **[]int32** | Filter by billing_type, 1-one-time, 2-recurring | [optional] 
**CodePrefix** | Pointer to **string** | Filter by code prefix | [optional] 
**Count** | Pointer to **int32** | Count per page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | Filter by create time end, UTC timestamp in seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | Filter by create time start, UTC timestamp in seconds | [optional] 
**Currency** | Pointer to **string** | Filter by currency | [optional] 
**DiscountType** | Pointer to **[]int32** | Filter by discount_type, 1-percentage, 2-fixed_amount | [optional] 
**Page** | Pointer to **int32** | Page number, start from 0 | [optional] 
**SearchKey** | Pointer to **string** | Search by code prefix or name | [optional] 
**SortField** | Pointer to **string** | Sort field, gmt_create|gmt_modify, default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort type, asc|desc, default desc | [optional] 
**Status** | Pointer to **[]int32** | Filter by status, 1-editable, 2-active, 3-deactive, 4-expire, 10-archive | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchTemplateListReq

`func NewUnibeeApiMerchantDiscountBatchTemplateListReq() *UnibeeApiMerchantDiscountBatchTemplateListReq`

NewUnibeeApiMerchantDiscountBatchTemplateListReq instantiates a new UnibeeApiMerchantDiscountBatchTemplateListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchTemplateListReqWithDefaults

`func NewUnibeeApiMerchantDiscountBatchTemplateListReqWithDefaults() *UnibeeApiMerchantDiscountBatchTemplateListReq`

NewUnibeeApiMerchantDiscountBatchTemplateListReqWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchTemplateListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetBillingType() []int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetBillingTypeOk() (*[]int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetBillingType(v []int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCodePrefix() string`

GetCodePrefix returns the CodePrefix field if non-nil, zero value otherwise.

### GetCodePrefixOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCodePrefixOk() (*string, bool)`

GetCodePrefixOk returns a tuple with the CodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetCodePrefix(v string)`

SetCodePrefix sets CodePrefix field to given value.

### HasCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasCodePrefix() bool`

HasCodePrefix returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetDiscountType() []int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetDiscountTypeOk() (*[]int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetDiscountType(v []int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantDiscountBatchTemplateListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


