# UnibeeApiBeanMerchantBatchDiscountCodeTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will enable all advance config if set true | [optional] 
**BillingType** | Pointer to **int32** | billing_type, 1-one-time, 2-recurring | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**CodePrefix** | Pointer to **string** | codePrefix | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency of discount, available when discount_type is fixed_amount | [optional] 
**CycleLimit** | Pointer to **int32** | the count limitation of subscription cycle , 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | amount of discount, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | percentage of discount, 100&#x3D;1%, available when discount_type is percentage | [optional] 
**DiscountType** | Pointer to **int32** | discount_type, 1-percentage, 2-fixed_amount | [optional] 
**EndTime** | Pointer to **int64** | end of discount available utc time, 0-invalid | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，&gt; 0, Deleted, the deleted utc time | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PlanApplyGroup** | Pointer to [**UnibeeApiBeanGroupPlanSelector**](UnibeeApiBeanGroupPlanSelector.md) |  | [optional] 
**PlanApplyType** | Pointer to **int32** | plan apply type, 0-apply for all, 1-apply for plans specified, 2-exclude for plans specified, 3-Apply to Plans by Groups(Billing Period Included), 4-Apply to Plans except by Groups(Billing Period Included) | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect, default effect all plans if not set | [optional] 
**Quantity** | Pointer to **int64** | quantity of code, 0-no limit | [optional] 
**StartTime** | Pointer to **int64** | start of discount available utc time | [optional] 
**Status** | Pointer to **int32** | status, 1-editable, 2-active, 3-deactivate, 4-expire, 10-archive | [optional] 
**Type** | Pointer to **int32** | type, 1-external discount code, 2-batch code template, 3-batch code | [optional] 
**UpgradeLongerOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except upgrade to longer plan if set true | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except same interval upgrade action if set true | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, The limit of every customer can apply, the recurring apply not involved, 0-unlimited | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope,0-for all, 1-for only new user, 2-for only renewals, renewals is upgrade&amp;downgrade&amp;renew | [optional] 

## Methods

### NewUnibeeApiBeanMerchantBatchDiscountCodeTemplate

`func NewUnibeeApiBeanMerchantBatchDiscountCodeTemplate() *UnibeeApiBeanMerchantBatchDiscountCodeTemplate`

NewUnibeeApiBeanMerchantBatchDiscountCodeTemplate instantiates a new UnibeeApiBeanMerchantBatchDiscountCodeTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchDiscountCodeTemplateWithDefaults

`func NewUnibeeApiBeanMerchantBatchDiscountCodeTemplateWithDefaults() *UnibeeApiBeanMerchantBatchDiscountCodeTemplate`

NewUnibeeApiBeanMerchantBatchDiscountCodeTemplateWithDefaults instantiates a new UnibeeApiBeanMerchantBatchDiscountCodeTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCodePrefix() string`

GetCodePrefix returns the CodePrefix field if non-nil, zero value otherwise.

### GetCodePrefixOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCodePrefixOk() (*string, bool)`

GetCodePrefixOk returns a tuple with the CodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetCodePrefix(v string)`

SetCodePrefix sets CodePrefix field to given value.

### HasCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasCodePrefix() bool`

HasCodePrefix returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanApplyGroup() UnibeeApiBeanGroupPlanSelector`

GetPlanApplyGroup returns the PlanApplyGroup field if non-nil, zero value otherwise.

### GetPlanApplyGroupOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanApplyGroupOk() (*UnibeeApiBeanGroupPlanSelector, bool)`

GetPlanApplyGroupOk returns a tuple with the PlanApplyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetPlanApplyGroup(v UnibeeApiBeanGroupPlanSelector)`

SetPlanApplyGroup sets PlanApplyGroup field to given value.

### HasPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasPlanApplyGroup() bool`

HasPlanApplyGroup returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUpgradeLongerOnly() bool`

GetUpgradeLongerOnly returns the UpgradeLongerOnly field if non-nil, zero value otherwise.

### GetUpgradeLongerOnlyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUpgradeLongerOnlyOk() (*bool, bool)`

GetUpgradeLongerOnlyOk returns a tuple with the UpgradeLongerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetUpgradeLongerOnly(v bool)`

SetUpgradeLongerOnly sets UpgradeLongerOnly field to given value.

### HasUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasUpgradeLongerOnly() bool`

HasUpgradeLongerOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCodeTemplate) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


