# UnibeeApiMerchantMetricEventListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count OF Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**MetricIds** | Pointer to **[]int64** | Filter MetricIds, Default All | [optional] 
**Page** | Pointer to **int32** | Page,Start 0 | [optional] 
**SortField** | Pointer to **string** | Sort，user_id|gmt_create，Default gmt_create | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**UserIds** | Pointer to **[]int64** | Filter UserIds, Default All | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventListReq

`func NewUnibeeApiMerchantMetricEventListReq() *UnibeeApiMerchantMetricEventListReq`

NewUnibeeApiMerchantMetricEventListReq instantiates a new UnibeeApiMerchantMetricEventListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventListReqWithDefaults

`func NewUnibeeApiMerchantMetricEventListReqWithDefaults() *UnibeeApiMerchantMetricEventListReq`

NewUnibeeApiMerchantMetricEventListReqWithDefaults instantiates a new UnibeeApiMerchantMetricEventListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantMetricEventListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantMetricEventListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantMetricEventListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantMetricEventListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantMetricEventListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantMetricEventListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantMetricEventListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetMetricIds

`func (o *UnibeeApiMerchantMetricEventListReq) GetMetricIds() []int64`

GetMetricIds returns the MetricIds field if non-nil, zero value otherwise.

### GetMetricIdsOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetMetricIdsOk() (*[]int64, bool)`

GetMetricIdsOk returns a tuple with the MetricIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIds

`func (o *UnibeeApiMerchantMetricEventListReq) SetMetricIds(v []int64)`

SetMetricIds sets MetricIds field to given value.

### HasMetricIds

`func (o *UnibeeApiMerchantMetricEventListReq) HasMetricIds() bool`

HasMetricIds returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantMetricEventListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantMetricEventListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantMetricEventListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantMetricEventListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantMetricEventListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantMetricEventListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantMetricEventListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantMetricEventListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantMetricEventListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetUserIds

`func (o *UnibeeApiMerchantMetricEventListReq) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UnibeeApiMerchantMetricEventListReq) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UnibeeApiMerchantMetricEventListReq) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UnibeeApiMerchantMetricEventListReq) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


