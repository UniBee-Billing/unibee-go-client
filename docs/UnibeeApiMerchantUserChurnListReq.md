# UnibeeApiMerchantUserChurnListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChurnTimeEnd** | Pointer to **int64** | ChurnTimeEnd, UTC seconds, exclusive [start,end) | [optional] 
**ChurnTimeStart** | Pointer to **int64** | ChurnTimeStart, UTC seconds, inclusive | [optional] 
**ChurnType** | Pointer to **string** | Cancelled|Expired | [optional] 
**Count** | Pointer to **int32** | Count Of Per Page | [optional] 
**CountryCode** | Pointer to **string** | Filter CountryCode | [optional] 
**Currency** | Pointer to **string** | Filter Currency | [optional] 
**Email** | Pointer to **string** | Filter Email | [optional] 
**ExternalUserId** | Pointer to **string** | Filter ExternalUserId | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**PlanIds** | Pointer to **[]int64** | Filter PlanIds | [optional] 
**ProductIds** | Pointer to **[]int64** | Filter ProductIds | [optional] 
**ReasonCategory** | Pointer to **string** | Normalized reason category | [optional] 
**SearchKey** | Pointer to **string** | Search Email|ExternalUserId|SubscriptionId|UserId | [optional] 
**SortField** | Pointer to **string** | churn_time|confirm_time|reactivated_time|id, default churn_time | [optional] 
**SortType** | Pointer to **string** | asc|desc, default desc | [optional] 
**UserId** | Pointer to **int64** | Filter UserId | [optional] 
**View** | Pointer to **string** | pending|current|reactivated|all, default all | [optional] 

## Methods

### NewUnibeeApiMerchantUserChurnListReq

`func NewUnibeeApiMerchantUserChurnListReq() *UnibeeApiMerchantUserChurnListReq`

NewUnibeeApiMerchantUserChurnListReq instantiates a new UnibeeApiMerchantUserChurnListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserChurnListReqWithDefaults

`func NewUnibeeApiMerchantUserChurnListReqWithDefaults() *UnibeeApiMerchantUserChurnListReq`

NewUnibeeApiMerchantUserChurnListReqWithDefaults instantiates a new UnibeeApiMerchantUserChurnListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChurnTimeEnd

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnTimeEnd() int64`

GetChurnTimeEnd returns the ChurnTimeEnd field if non-nil, zero value otherwise.

### GetChurnTimeEndOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnTimeEndOk() (*int64, bool)`

GetChurnTimeEndOk returns a tuple with the ChurnTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnTimeEnd

`func (o *UnibeeApiMerchantUserChurnListReq) SetChurnTimeEnd(v int64)`

SetChurnTimeEnd sets ChurnTimeEnd field to given value.

### HasChurnTimeEnd

`func (o *UnibeeApiMerchantUserChurnListReq) HasChurnTimeEnd() bool`

HasChurnTimeEnd returns a boolean if a field has been set.

### GetChurnTimeStart

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnTimeStart() int64`

GetChurnTimeStart returns the ChurnTimeStart field if non-nil, zero value otherwise.

### GetChurnTimeStartOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnTimeStartOk() (*int64, bool)`

GetChurnTimeStartOk returns a tuple with the ChurnTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnTimeStart

`func (o *UnibeeApiMerchantUserChurnListReq) SetChurnTimeStart(v int64)`

SetChurnTimeStart sets ChurnTimeStart field to given value.

### HasChurnTimeStart

`func (o *UnibeeApiMerchantUserChurnListReq) HasChurnTimeStart() bool`

HasChurnTimeStart returns a boolean if a field has been set.

### GetChurnType

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnType() string`

GetChurnType returns the ChurnType field if non-nil, zero value otherwise.

### GetChurnTypeOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetChurnTypeOk() (*string, bool)`

GetChurnTypeOk returns a tuple with the ChurnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnType

`func (o *UnibeeApiMerchantUserChurnListReq) SetChurnType(v string)`

SetChurnType sets ChurnType field to given value.

### HasChurnType

`func (o *UnibeeApiMerchantUserChurnListReq) HasChurnType() bool`

HasChurnType returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantUserChurnListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantUserChurnListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantUserChurnListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantUserChurnListReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantUserChurnListReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantUserChurnListReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantUserChurnListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantUserChurnListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantUserChurnListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantUserChurnListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserChurnListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantUserChurnListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantUserChurnListReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantUserChurnListReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantUserChurnListReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantUserChurnListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantUserChurnListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantUserChurnListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantUserChurnListReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantUserChurnListReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantUserChurnListReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetProductIds

`func (o *UnibeeApiMerchantUserChurnListReq) GetProductIds() []int64`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetProductIdsOk() (*[]int64, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *UnibeeApiMerchantUserChurnListReq) SetProductIds(v []int64)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *UnibeeApiMerchantUserChurnListReq) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetReasonCategory

`func (o *UnibeeApiMerchantUserChurnListReq) GetReasonCategory() string`

GetReasonCategory returns the ReasonCategory field if non-nil, zero value otherwise.

### GetReasonCategoryOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetReasonCategoryOk() (*string, bool)`

GetReasonCategoryOk returns a tuple with the ReasonCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCategory

`func (o *UnibeeApiMerchantUserChurnListReq) SetReasonCategory(v string)`

SetReasonCategory sets ReasonCategory field to given value.

### HasReasonCategory

`func (o *UnibeeApiMerchantUserChurnListReq) HasReasonCategory() bool`

HasReasonCategory returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantUserChurnListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantUserChurnListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantUserChurnListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantUserChurnListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantUserChurnListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantUserChurnListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantUserChurnListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantUserChurnListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantUserChurnListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserChurnListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserChurnListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserChurnListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetView

`func (o *UnibeeApiMerchantUserChurnListReq) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *UnibeeApiMerchantUserChurnListReq) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *UnibeeApiMerchantUserChurnListReq) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *UnibeeApiMerchantUserChurnListReq) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


