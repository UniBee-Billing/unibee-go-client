# UnibeeApiMerchantDiscountBatchTemplateNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to **bool** | AdvanceConfig, enable advanced configurations | [optional] 
**BillingType** | **int32** | The billing type of the discount code, 1-one-time, 2-recurring | 
**CodePrefix** | **string** | The unique code prefix for this batch template, will be used as template identifier and child code prefix, length 1-20 | 
**Currency** | Pointer to **string** | The discount currency, available when discount_type is fixed_amount | [optional] 
**CycleLimit** | Pointer to **int32** | The count limitation of subscription cycle, 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | The discount amount, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | The discount percentage, 100&#x3D;1%, available when discount_type is percentage | [optional] 
**DiscountType** | **int32** | The discount type of the discount code, 1-percentage, 2-fixed_amount | 
**EndTime** | **int32** | The end time of discount code can effect, utc timestamp in seconds | 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata, custom key-value pairs | [optional] 
**Name** | Pointer to **string** | The batch template&#39;s display name | [optional] 
**PlanApplyGroup** | Pointer to [**UnibeeApiBeanGroupPlanSelector**](UnibeeApiBeanGroupPlanSelector.md) |  | [optional] 
**PlanApplyType** | Pointer to **int32** | Plan apply type, 0-apply for all, 1-apply for plans specified, 2-exclude for plans specified, 3-Apply to Plans by Groups, 4-Apply to Plans except by Groups | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect | [optional] 
**Quantity** | **int32** | Target number of child codes to generate, max 10000 | 
**StartTime** | **int32** | The start time of discount code can effect, utc timestamp in seconds | 
**SubscriptionLimit** | Pointer to **int32** | The limit of every subscription apply, 0-unlimited | [optional] 
**UpgradeLongPlanOnly** | Pointer to **bool** | AdvanceConfig, true-forbid for all except upgrade to longer plan | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, true-forbid for all except same interval upgrade action | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, The limit of every customer can apply, 0-unlimited | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope, 0-for all, 1-for only new user, 2-for only renewals | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchTemplateNewReq

`func NewUnibeeApiMerchantDiscountBatchTemplateNewReq(billingType int32, codePrefix string, discountType int32, endTime int32, quantity int32, startTime int32, ) *UnibeeApiMerchantDiscountBatchTemplateNewReq`

NewUnibeeApiMerchantDiscountBatchTemplateNewReq instantiates a new UnibeeApiMerchantDiscountBatchTemplateNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchTemplateNewReqWithDefaults

`func NewUnibeeApiMerchantDiscountBatchTemplateNewReqWithDefaults() *UnibeeApiMerchantDiscountBatchTemplateNewReq`

NewUnibeeApiMerchantDiscountBatchTemplateNewReqWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchTemplateNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.


### GetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCodePrefix() string`

GetCodePrefix returns the CodePrefix field if non-nil, zero value otherwise.

### GetCodePrefixOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCodePrefixOk() (*string, bool)`

GetCodePrefixOk returns a tuple with the CodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePrefix

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetCodePrefix(v string)`

SetCodePrefix sets CodePrefix field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.


### GetEndTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanApplyGroup() UnibeeApiBeanGroupPlanSelector`

GetPlanApplyGroup returns the PlanApplyGroup field if non-nil, zero value otherwise.

### GetPlanApplyGroupOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanApplyGroupOk() (*UnibeeApiBeanGroupPlanSelector, bool)`

GetPlanApplyGroupOk returns a tuple with the PlanApplyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetPlanApplyGroup(v UnibeeApiBeanGroupPlanSelector)`

SetPlanApplyGroup sets PlanApplyGroup field to given value.

### HasPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasPlanApplyGroup() bool`

HasPlanApplyGroup returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetStartTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetSubscriptionLimit() int32`

GetSubscriptionLimit returns the SubscriptionLimit field if non-nil, zero value otherwise.

### GetSubscriptionLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetSubscriptionLimitOk() (*int32, bool)`

GetSubscriptionLimitOk returns a tuple with the SubscriptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetSubscriptionLimit(v int32)`

SetSubscriptionLimit sets SubscriptionLimit field to given value.

### HasSubscriptionLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasSubscriptionLimit() bool`

HasSubscriptionLimit returns a boolean if a field has been set.

### GetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUpgradeLongPlanOnly() bool`

GetUpgradeLongPlanOnly returns the UpgradeLongPlanOnly field if non-nil, zero value otherwise.

### GetUpgradeLongPlanOnlyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUpgradeLongPlanOnlyOk() (*bool, bool)`

GetUpgradeLongPlanOnlyOk returns a tuple with the UpgradeLongPlanOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetUpgradeLongPlanOnly(v bool)`

SetUpgradeLongPlanOnly sets UpgradeLongPlanOnly field to given value.

### HasUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasUpgradeLongPlanOnly() bool`

HasUpgradeLongPlanOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiMerchantDiscountBatchTemplateNewReq) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


