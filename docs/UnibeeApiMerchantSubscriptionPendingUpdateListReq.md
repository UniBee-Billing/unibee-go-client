# UnibeeApiMerchantSubscriptionPendingUpdateListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start WIth 0 | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 

## Methods

### NewUnibeeApiMerchantSubscriptionPendingUpdateListReq

`func NewUnibeeApiMerchantSubscriptionPendingUpdateListReq(subscriptionId string, ) *UnibeeApiMerchantSubscriptionPendingUpdateListReq`

NewUnibeeApiMerchantSubscriptionPendingUpdateListReq instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionPendingUpdateListReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionPendingUpdateListReqWithDefaults() *UnibeeApiMerchantSubscriptionPendingUpdateListReq`

NewUnibeeApiMerchantSubscriptionPendingUpdateListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionPendingUpdateListReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


