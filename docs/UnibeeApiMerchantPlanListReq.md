# UnibeeApiMerchantPlanListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Per Page | [optional] 
**Currency** | Pointer to **string** | Filter Currency | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**ProductIds** | Pointer to **[]int64** | filter id list of product, default all | [optional] 
**PublishStatus** | Pointer to **int32** | Filter, Default All，PublishStatus，1-UnPublished，2-Published | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-Expired | [optional] 
**Type** | Pointer to **[]int32** | 1-main plan，2-addon plan | [optional] 

## Methods

### NewUnibeeApiMerchantPlanListReq

`func NewUnibeeApiMerchantPlanListReq() *UnibeeApiMerchantPlanListReq`

NewUnibeeApiMerchantPlanListReq instantiates a new UnibeeApiMerchantPlanListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanListReqWithDefaults

`func NewUnibeeApiMerchantPlanListReqWithDefaults() *UnibeeApiMerchantPlanListReq`

NewUnibeeApiMerchantPlanListReqWithDefaults instantiates a new UnibeeApiMerchantPlanListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantPlanListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantPlanListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantPlanListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantPlanListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPlanListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPlanListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPlanListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPlanListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantPlanListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantPlanListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantPlanListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantPlanListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetProductIds

`func (o *UnibeeApiMerchantPlanListReq) GetProductIds() []int64`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *UnibeeApiMerchantPlanListReq) GetProductIdsOk() (*[]int64, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *UnibeeApiMerchantPlanListReq) SetProductIds(v []int64)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *UnibeeApiMerchantPlanListReq) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetPublishStatus

`func (o *UnibeeApiMerchantPlanListReq) GetPublishStatus() int32`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *UnibeeApiMerchantPlanListReq) GetPublishStatusOk() (*int32, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *UnibeeApiMerchantPlanListReq) SetPublishStatus(v int32)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *UnibeeApiMerchantPlanListReq) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantPlanListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantPlanListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantPlanListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantPlanListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantPlanListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantPlanListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantPlanListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantPlanListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantPlanListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantPlanListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantPlanListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantPlanListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantPlanListReq) GetType() []int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantPlanListReq) GetTypeOk() (*[]int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantPlanListReq) SetType(v []int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantPlanListReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


