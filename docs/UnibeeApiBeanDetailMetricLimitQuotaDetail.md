# UnibeeApiBeanDetailMetricLimitQuotaDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustmentTime** | Pointer to **int64** | adjustment time (for manual) | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MerchantMemberEmail** | Pointer to **string** | operator email (for manual) | [optional] 
**MerchantMemberId** | Pointer to **int64** | operator id (for manual) | [optional] 
**PreviousPeriodLimit** | Pointer to **int64** | previous period limit (for carryover) | [optional] 
**PreviousPeriodUsed** | Pointer to **int64** | previous period used (for carryover) | [optional] 
**QuotaAmount** | Pointer to **int64** | quota amount | [optional] 
**QuotaType** | Pointer to **string** | carryover or manual | [optional] 
**Reason** | Pointer to **string** | reason/note | [optional] 

## Methods

### NewUnibeeApiBeanDetailMetricLimitQuotaDetail

`func NewUnibeeApiBeanDetailMetricLimitQuotaDetail() *UnibeeApiBeanDetailMetricLimitQuotaDetail`

NewUnibeeApiBeanDetailMetricLimitQuotaDetail instantiates a new UnibeeApiBeanDetailMetricLimitQuotaDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMetricLimitQuotaDetailWithDefaults

`func NewUnibeeApiBeanDetailMetricLimitQuotaDetailWithDefaults() *UnibeeApiBeanDetailMetricLimitQuotaDetail`

NewUnibeeApiBeanDetailMetricLimitQuotaDetailWithDefaults instantiates a new UnibeeApiBeanDetailMetricLimitQuotaDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustmentTime

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetAdjustmentTime() int64`

GetAdjustmentTime returns the AdjustmentTime field if non-nil, zero value otherwise.

### GetAdjustmentTimeOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetAdjustmentTimeOk() (*int64, bool)`

GetAdjustmentTimeOk returns a tuple with the AdjustmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTime

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetAdjustmentTime(v int64)`

SetAdjustmentTime sets AdjustmentTime field to given value.

### HasAdjustmentTime

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasAdjustmentTime() bool`

HasAdjustmentTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantMemberEmail

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetMerchantMemberEmail() string`

GetMerchantMemberEmail returns the MerchantMemberEmail field if non-nil, zero value otherwise.

### GetMerchantMemberEmailOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetMerchantMemberEmailOk() (*string, bool)`

GetMerchantMemberEmailOk returns a tuple with the MerchantMemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberEmail

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetMerchantMemberEmail(v string)`

SetMerchantMemberEmail sets MerchantMemberEmail field to given value.

### HasMerchantMemberEmail

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasMerchantMemberEmail() bool`

HasMerchantMemberEmail returns a boolean if a field has been set.

### GetMerchantMemberId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetMerchantMemberId() int64`

GetMerchantMemberId returns the MerchantMemberId field if non-nil, zero value otherwise.

### GetMerchantMemberIdOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetMerchantMemberIdOk() (*int64, bool)`

GetMerchantMemberIdOk returns a tuple with the MerchantMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetMerchantMemberId(v int64)`

SetMerchantMemberId sets MerchantMemberId field to given value.

### HasMerchantMemberId

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasMerchantMemberId() bool`

HasMerchantMemberId returns a boolean if a field has been set.

### GetPreviousPeriodLimit

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetPreviousPeriodLimit() int64`

GetPreviousPeriodLimit returns the PreviousPeriodLimit field if non-nil, zero value otherwise.

### GetPreviousPeriodLimitOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetPreviousPeriodLimitOk() (*int64, bool)`

GetPreviousPeriodLimitOk returns a tuple with the PreviousPeriodLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPeriodLimit

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetPreviousPeriodLimit(v int64)`

SetPreviousPeriodLimit sets PreviousPeriodLimit field to given value.

### HasPreviousPeriodLimit

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasPreviousPeriodLimit() bool`

HasPreviousPeriodLimit returns a boolean if a field has been set.

### GetPreviousPeriodUsed

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetPreviousPeriodUsed() int64`

GetPreviousPeriodUsed returns the PreviousPeriodUsed field if non-nil, zero value otherwise.

### GetPreviousPeriodUsedOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetPreviousPeriodUsedOk() (*int64, bool)`

GetPreviousPeriodUsedOk returns a tuple with the PreviousPeriodUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPeriodUsed

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetPreviousPeriodUsed(v int64)`

SetPreviousPeriodUsed sets PreviousPeriodUsed field to given value.

### HasPreviousPeriodUsed

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasPreviousPeriodUsed() bool`

HasPreviousPeriodUsed returns a boolean if a field has been set.

### GetQuotaAmount

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetQuotaAmount() int64`

GetQuotaAmount returns the QuotaAmount field if non-nil, zero value otherwise.

### GetQuotaAmountOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetQuotaAmountOk() (*int64, bool)`

GetQuotaAmountOk returns a tuple with the QuotaAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaAmount

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetQuotaAmount(v int64)`

SetQuotaAmount sets QuotaAmount field to given value.

### HasQuotaAmount

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasQuotaAmount() bool`

HasQuotaAmount returns a boolean if a field has been set.

### GetQuotaType

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.

### GetReason

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiBeanDetailMetricLimitQuotaDetail) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


