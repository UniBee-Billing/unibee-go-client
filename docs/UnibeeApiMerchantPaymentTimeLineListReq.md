# UnibeeApiMerchantPaymentTimeLineListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountEnd** | Pointer to **int32** | The filter end amount of timeline | [optional] 
**AmountStart** | Pointer to **int32** | The filter start amount of timeline | [optional] 
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**GatewayIds** | Pointer to **[]int64** | The filter ids of gateway | [optional] 
**Page** | Pointer to **int32** | Page,Start 0 | [optional] 
**SortField** | Pointer to **string** | Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | The filter status, 0-pending, 1-success, 2-failure，3-cancel | [optional] 
**TimelineTypes** | Pointer to **[]int32** | The filter timelineType, 0-pay, 1-refund | [optional] 
**UserId** | Pointer to **int64** | Filter UserId, Default All | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentTimeLineListReq

`func NewUnibeeApiMerchantPaymentTimeLineListReq() *UnibeeApiMerchantPaymentTimeLineListReq`

NewUnibeeApiMerchantPaymentTimeLineListReq instantiates a new UnibeeApiMerchantPaymentTimeLineListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentTimeLineListReqWithDefaults

`func NewUnibeeApiMerchantPaymentTimeLineListReqWithDefaults() *UnibeeApiMerchantPaymentTimeLineListReq`

NewUnibeeApiMerchantPaymentTimeLineListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentTimeLineListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountEnd() int32`

GetAmountEnd returns the AmountEnd field if non-nil, zero value otherwise.

### GetAmountEndOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountEndOk() (*int32, bool)`

GetAmountEndOk returns a tuple with the AmountEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetAmountEnd(v int32)`

SetAmountEnd sets AmountEnd field to given value.

### HasAmountEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasAmountEnd() bool`

HasAmountEnd returns a boolean if a field has been set.

### GetAmountStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountStart() int32`

GetAmountStart returns the AmountStart field if non-nil, zero value otherwise.

### GetAmountStartOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountStartOk() (*int32, bool)`

GetAmountStartOk returns a tuple with the AmountStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetAmountStart(v int32)`

SetAmountStart sets AmountStart field to given value.

### HasAmountStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasAmountStart() bool`

HasAmountStart returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayIds

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetGatewayIds() []int64`

GetGatewayIds returns the GatewayIds field if non-nil, zero value otherwise.

### GetGatewayIdsOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetGatewayIdsOk() (*[]int64, bool)`

GetGatewayIdsOk returns a tuple with the GatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIds

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetGatewayIds(v []int64)`

SetGatewayIds sets GatewayIds field to given value.

### HasGatewayIds

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasGatewayIds() bool`

HasGatewayIds returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimelineTypes

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetTimelineTypes() []int32`

GetTimelineTypes returns the TimelineTypes field if non-nil, zero value otherwise.

### GetTimelineTypesOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetTimelineTypesOk() (*[]int32, bool)`

GetTimelineTypesOk returns a tuple with the TimelineTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTypes

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetTimelineTypes(v []int32)`

SetTimelineTypes sets TimelineTypes field to given value.

### HasTimelineTypes

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasTimelineTypes() bool`

HasTimelineTypes returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


