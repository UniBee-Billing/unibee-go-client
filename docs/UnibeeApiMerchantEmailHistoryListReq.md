# UnibeeApiMerchantEmailHistoryListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Per Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart | [optional] 
**Email** | Pointer to **string** | Filter Email | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**SearchKey** | Pointer to **string** | Search Key, email or title | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 

## Methods

### NewUnibeeApiMerchantEmailHistoryListReq

`func NewUnibeeApiMerchantEmailHistoryListReq() *UnibeeApiMerchantEmailHistoryListReq`

NewUnibeeApiMerchantEmailHistoryListReq instantiates a new UnibeeApiMerchantEmailHistoryListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailHistoryListReqWithDefaults

`func NewUnibeeApiMerchantEmailHistoryListReqWithDefaults() *UnibeeApiMerchantEmailHistoryListReq`

NewUnibeeApiMerchantEmailHistoryListReqWithDefaults instantiates a new UnibeeApiMerchantEmailHistoryListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantEmailHistoryListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantEmailHistoryListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantEmailHistoryListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


