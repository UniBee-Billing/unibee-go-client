# UnibeeApiMerchantDiscountBatchTemplateEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to **bool** | AdvanceConfig, enable advanced configurations | [optional] 
**BillingType** | Pointer to **int32** | The billing type, 1-one-time, 2-recurring | [optional] 
**CodePrefix** | **string** | The code prefix, must match existing value, cannot be modified | 
**Currency** | Pointer to **string** | The discount currency | [optional] 
**CycleLimit** | Pointer to **int32** | The count limitation of subscription cycle, 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | The discount amount, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | The discount percentage, 100&#x3D;1% | [optional] 
**DiscountType** | Pointer to **int32** | The discount type, 1-percentage, 2-fixed_amount | [optional] 
**EndTime** | Pointer to **int32** | The end time, editable after activate, utc timestamp in seconds | [optional] 
**Id** | **int64** | The template&#39;s Id | 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata, custom key-value pairs | [optional] 
**Name** | Pointer to **string** | The batch template&#39;s display name | [optional] 
**PlanApplyGroup** | Pointer to [**UnibeeApiBeanGroupPlanSelector**](UnibeeApiBeanGroupPlanSelector.md) |  | [optional] 
**PlanApplyType** | Pointer to **int32** | Plan apply type, 0-apply for all, 1-specified, 2-exclude, 3-by groups, 4-except groups | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect | [optional] 
**Quantity** | Pointer to **int32** | Target number of child codes, can only increase after activation, max 10000 | [optional] 
**StartTime** | Pointer to **int32** | The start time, editable after activate, utc timestamp in seconds | [optional] 
**SubscriptionLimit** | Pointer to **int32** | The limit of every subscription apply, 0-unlimited | [optional] 
**UpgradeLongPlanOnly** | Pointer to **bool** | AdvanceConfig, true-only for upgrade to longer plan | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, true-only for same interval upgrade | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, Per customer limit, 0-unlimited | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope, 0-all, 1-new user, 2-renewals | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchTemplateEditReq

`func NewUnibeeApiMerchantDiscountBatchTemplateEditReq(codePrefix string, id int64, ) *UnibeeApiMerchantDiscountBatchTemplateEditReq`

NewUnibeeApiMerchantDiscountBatchTemplateEditReq instantiates a new UnibeeApiMerchantDiscountBatchTemplateEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchTemplateEditReqWithDefaults

`func NewUnibeeApiMerchantDiscountBatchTemplateEditReqWithDefaults() *UnibeeApiMerchantDiscountBatchTemplateEditReq`

NewUnibeeApiMerchantDiscountBatchTemplateEditReqWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchTemplateEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCodePrefix() string`

GetCodePrefix returns the CodePrefix field if non-nil, zero value otherwise.

### GetCodePrefixOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCodePrefixOk() (*string, bool)`

GetCodePrefixOk returns a tuple with the CodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetCodePrefix(v string)`

SetCodePrefix sets CodePrefix field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetId(v int64)`

SetId sets Id field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanApplyGroup() UnibeeApiBeanGroupPlanSelector`

GetPlanApplyGroup returns the PlanApplyGroup field if non-nil, zero value otherwise.

### GetPlanApplyGroupOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanApplyGroupOk() (*UnibeeApiBeanGroupPlanSelector, bool)`

GetPlanApplyGroupOk returns a tuple with the PlanApplyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetPlanApplyGroup(v UnibeeApiBeanGroupPlanSelector)`

SetPlanApplyGroup sets PlanApplyGroup field to given value.

### HasPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasPlanApplyGroup() bool`

HasPlanApplyGroup returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetSubscriptionLimit() int32`

GetSubscriptionLimit returns the SubscriptionLimit field if non-nil, zero value otherwise.

### GetSubscriptionLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetSubscriptionLimitOk() (*int32, bool)`

GetSubscriptionLimitOk returns a tuple with the SubscriptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetSubscriptionLimit(v int32)`

SetSubscriptionLimit sets SubscriptionLimit field to given value.

### HasSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasSubscriptionLimit() bool`

HasSubscriptionLimit returns a boolean if a field has been set.

### GetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUpgradeLongPlanOnly() bool`

GetUpgradeLongPlanOnly returns the UpgradeLongPlanOnly field if non-nil, zero value otherwise.

### GetUpgradeLongPlanOnlyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUpgradeLongPlanOnlyOk() (*bool, bool)`

GetUpgradeLongPlanOnlyOk returns a tuple with the UpgradeLongPlanOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetUpgradeLongPlanOnly(v bool)`

SetUpgradeLongPlanOnly sets UpgradeLongPlanOnly field to given value.

### HasUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasUpgradeLongPlanOnly() bool`

HasUpgradeLongPlanOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateEditReq) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


