# UnibeeApiBeanMerchantDiscountCodeSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingType** | Pointer to **int32** | billing_type, 1-one-time, 2-recurring | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency of discount, available when discount_type is fixed_amount | [optional] 
**CycleLimit** | Pointer to **int32** | the count limitation of subscription cycle , 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | amount of discount, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | percentage of discount, 100&#x3D;1%, available when discount_type is percentage | [optional] 
**DiscountType** | Pointer to **int32** | discount_type, 1-percentage, 2-fixed_amount | [optional] 
**EndTime** | Pointer to **int64** | end of discount available utc time, 0-invalid | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect, default effect all plans if not set | [optional] 
**StartTime** | Pointer to **int64** | start of discount available utc time | [optional] 
**Status** | Pointer to **int32** | status, 1-editable, 2-active, 3-deactive, 4-expire | [optional] 

## Methods

### NewUnibeeApiBeanMerchantDiscountCodeSimplify

`func NewUnibeeApiBeanMerchantDiscountCodeSimplify() *UnibeeApiBeanMerchantDiscountCodeSimplify`

NewUnibeeApiBeanMerchantDiscountCodeSimplify instantiates a new UnibeeApiBeanMerchantDiscountCodeSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantDiscountCodeSimplifyWithDefaults

`func NewUnibeeApiBeanMerchantDiscountCodeSimplifyWithDefaults() *UnibeeApiBeanMerchantDiscountCodeSimplify`

NewUnibeeApiBeanMerchantDiscountCodeSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantDiscountCodeSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantDiscountCodeSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


