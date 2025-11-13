# UnibeeApiMerchantSubscriptionListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The filter email of subscription user | [optional] 
**AmountEnd** | Pointer to **int32** | The filter end amount of subscription | [optional] 
**AmountStart** | Pointer to **int32** | The filter start amount of subscription | [optional] 
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Currency** | Pointer to **string** | The currency of subscription | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**PlanIds** | Pointer to **[]int64** | The filter ids of plan | [optional] 
**ProductIds** | Pointer to **[]int64** | The filter ids of product, invalid if planIds is provided | [optional] 
**SearchKey** | Pointer to **string** | Search SubscriptionId|Email | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionListReq

`func NewUnibeeApiMerchantSubscriptionListReq() *UnibeeApiMerchantSubscriptionListReq`

NewUnibeeApiMerchantSubscriptionListReq instantiates a new UnibeeApiMerchantSubscriptionListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionListReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionListReqWithDefaults() *UnibeeApiMerchantSubscriptionListReq`

NewUnibeeApiMerchantSubscriptionListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAmountEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountEnd() int32`

GetAmountEnd returns the AmountEnd field if non-nil, zero value otherwise.

### GetAmountEndOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountEndOk() (*int32, bool)`

GetAmountEndOk returns a tuple with the AmountEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) SetAmountEnd(v int32)`

SetAmountEnd sets AmountEnd field to given value.

### HasAmountEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) HasAmountEnd() bool`

HasAmountEnd returns a boolean if a field has been set.

### GetAmountStart

`func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountStart() int32`

GetAmountStart returns the AmountStart field if non-nil, zero value otherwise.

### GetAmountStartOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountStartOk() (*int32, bool)`

GetAmountStartOk returns a tuple with the AmountStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountStart

`func (o *UnibeeApiMerchantSubscriptionListReq) SetAmountStart(v int32)`

SetAmountStart sets AmountStart field to given value.

### HasAmountStart

`func (o *UnibeeApiMerchantSubscriptionListReq) HasAmountStart() bool`

HasAmountStart returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantSubscriptionListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantSubscriptionListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantSubscriptionListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantSubscriptionListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantSubscriptionListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantSubscriptionListReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantSubscriptionListReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantSubscriptionListReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetProductIds

`func (o *UnibeeApiMerchantSubscriptionListReq) GetProductIds() []int64`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetProductIdsOk() (*[]int64, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *UnibeeApiMerchantSubscriptionListReq) SetProductIds(v []int64)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *UnibeeApiMerchantSubscriptionListReq) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantSubscriptionListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantSubscriptionListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantSubscriptionListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantSubscriptionListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantSubscriptionListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantSubscriptionListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantSubscriptionListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantSubscriptionListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantSubscriptionListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


