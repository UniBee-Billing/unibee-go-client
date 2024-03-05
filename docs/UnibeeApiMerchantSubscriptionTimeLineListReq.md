# UnibeeApiMerchantSubscriptionTimeLineListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start WIth 0 | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**UserId** | Pointer to **int32** | Filter UserId, Default All  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionTimeLineListReq

`func NewUnibeeApiMerchantSubscriptionTimeLineListReq() *UnibeeApiMerchantSubscriptionTimeLineListReq`

NewUnibeeApiMerchantSubscriptionTimeLineListReq instantiates a new UnibeeApiMerchantSubscriptionTimeLineListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionTimeLineListReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionTimeLineListReqWithDefaults() *UnibeeApiMerchantSubscriptionTimeLineListReq`

NewUnibeeApiMerchantSubscriptionTimeLineListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionTimeLineListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionTimeLineListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


