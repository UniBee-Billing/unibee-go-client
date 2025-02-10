# UnibeeApiBeanDetailMerchantDiscountCodeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will enable all advance config if set 1 | [optional] 
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
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，&gt; 0, Deleted, the deleted utc time | [optional] 
**LiveQuantity** | Pointer to **int64** | the live quantity of code | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PlanApplyType** | Pointer to **int32** | plan apply type, 0-apply for all, 1-apply for plans specified, 2-exclude for plans specified | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect, default effect all plans if not set | [optional] 
**Plans** | Pointer to [**[]UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) | plans which discount code can effect, default effect all plans if not set | [optional] 
**Quantity** | Pointer to **int64** | quantity of code, 0-no limit, will not change | [optional] 
**QuantityUsed** | Pointer to **int64** | quantity used count of code | [optional] 
**StartTime** | Pointer to **int64** | start of discount available utc time | [optional] 
**Status** | Pointer to **int32** | status, 1-editable, 2-active, 3-deactive, 4-expire, 10-archive | [optional] 
**UpgradeLongerOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except upgrade to longer plan if set 1 | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except upgrade action if set 1 | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, The limit of every customer can apply, the recurring apply not involved, 0-unlimited\&quot; | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope,0-for all, 1-for only new user, 2-for only renewals, renewals is upgrade&amp;downgrade&amp;renew | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantDiscountCodeDetail

`func NewUnibeeApiBeanDetailMerchantDiscountCodeDetail() *UnibeeApiBeanDetailMerchantDiscountCodeDetail`

NewUnibeeApiBeanDetailMerchantDiscountCodeDetail instantiates a new UnibeeApiBeanDetailMerchantDiscountCodeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantDiscountCodeDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantDiscountCodeDetailWithDefaults() *UnibeeApiBeanDetailMerchantDiscountCodeDetail`

NewUnibeeApiBeanDetailMerchantDiscountCodeDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantDiscountCodeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLiveQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetLiveQuantity() int64`

GetLiveQuantity returns the LiveQuantity field if non-nil, zero value otherwise.

### GetLiveQuantityOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetLiveQuantityOk() (*int64, bool)`

GetLiveQuantityOk returns a tuple with the LiveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetLiveQuantity(v int64)`

SetLiveQuantity sets LiveQuantity field to given value.

### HasLiveQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasLiveQuantity() bool`

HasLiveQuantity returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetPlans

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlans() []UnibeeApiBeanPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetPlansOk() (*[]UnibeeApiBeanPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetPlans(v []UnibeeApiBeanPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityUsed

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetQuantityUsed() int64`

GetQuantityUsed returns the QuantityUsed field if non-nil, zero value otherwise.

### GetQuantityUsedOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetQuantityUsedOk() (*int64, bool)`

GetQuantityUsedOk returns a tuple with the QuantityUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUsed

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetQuantityUsed(v int64)`

SetQuantityUsed sets QuantityUsed field to given value.

### HasQuantityUsed

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasQuantityUsed() bool`

HasQuantityUsed returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeLongerOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUpgradeLongerOnly() bool`

GetUpgradeLongerOnly returns the UpgradeLongerOnly field if non-nil, zero value otherwise.

### GetUpgradeLongerOnlyOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUpgradeLongerOnlyOk() (*bool, bool)`

GetUpgradeLongerOnlyOk returns a tuple with the UpgradeLongerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongerOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetUpgradeLongerOnly(v bool)`

SetUpgradeLongerOnly sets UpgradeLongerOnly field to given value.

### HasUpgradeLongerOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasUpgradeLongerOnly() bool`

HasUpgradeLongerOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiBeanDetailMerchantDiscountCodeDetail) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


