# UnibeeApiMerchantDiscountNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will enable all advance config if set true | [optional] 
**BillingType** | **int32** | The billing type of the discount code, 1-one-time, 2-recurring, define the situation the code can be used, the code of one-time billing_type can used for all situation that effect only once, the code of recurring billing_tye can only used for subscription purchase | 
**Code** | **string** | The discount&#39;s unique code, customize by merchant | 
**Currency** | Pointer to **string** | The discount currency of discount code, available when discount_type is fixed_amount | [optional] 
**CycleLimit** | Pointer to **int32** | The count limitation of subscription cycle, each subscription is valid separately , 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | The discount amount of the discount code, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | The discount percentage of discount code, 100&#x3D;1%, available when discount_type is percentage | [optional] 
**DiscountType** | **int32** | The discount type of the discount code, 1-percentage, 2-fixed_amount, the discountType of code, the discountPercentage will be effect when discountType is percentage, the discountAmount and currency will be effect when discountTYpe is fixed_amount | 
**EndTime** | **int32** | The end time of discount code can effect, utc time | 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**Name** | Pointer to **string** | The discount&#39;s name | [optional] 
**PlanApplyGroup** | Pointer to [**UnibeeApiBeanGroupPlanSelector**](UnibeeApiBeanGroupPlanSelector.md) |  | [optional] 
**PlanApplyType** | Pointer to **int32** | plan apply type, 0-apply for all, 1-apply for plans specified, 2-exclude for plans specified, 3-Apply to Plans by Groups(Billing Period Included), 4-Apply to Plans except by Groups(Billing Period Included) | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect, default effect all plans if not set | [optional] 
**Quantity** | Pointer to **int32** | Quantity of code, default 0, set 0 to disable quantity management | [optional] 
**StartTime** | **int32** | The start time of discount code can effect, utc time | 
**UpgradeLongPlanOnly** | Pointer to **bool** | AdvanceConfig, true or false, will forbid for all except upgrade to longer plan if set true | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, true or false, will forbid for all except same interval upgrade action if set true | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, The limit of every customer can apply, the recurring apply not involved, 0-unlimited | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope,0-for all, 1-for only new user, 2-for only renewals, renewals is upgrade&amp;downgrade&amp;renew | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountNewReq

`func NewUnibeeApiMerchantDiscountNewReq(billingType int32, code string, discountType int32, endTime int32, startTime int32, ) *UnibeeApiMerchantDiscountNewReq`

NewUnibeeApiMerchantDiscountNewReq instantiates a new UnibeeApiMerchantDiscountNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountNewReqWithDefaults

`func NewUnibeeApiMerchantDiscountNewReqWithDefaults() *UnibeeApiMerchantDiscountNewReq`

NewUnibeeApiMerchantDiscountNewReqWithDefaults instantiates a new UnibeeApiMerchantDiscountNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiMerchantDiscountNewReq) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiMerchantDiscountNewReq) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiMerchantDiscountNewReq) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiMerchantDiscountNewReq) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiMerchantDiscountNewReq) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.


### GetCode

`func (o *UnibeeApiMerchantDiscountNewReq) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiMerchantDiscountNewReq) SetCode(v string)`

SetCode sets Code field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantDiscountNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantDiscountNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantDiscountNewReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiMerchantDiscountNewReq) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiMerchantDiscountNewReq) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiMerchantDiscountNewReq) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantDiscountNewReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantDiscountNewReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.


### GetEndTime

`func (o *UnibeeApiMerchantDiscountNewReq) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiMerchantDiscountNewReq) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantDiscountNewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantDiscountNewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantDiscountNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantDiscountNewReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantDiscountNewReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantDiscountNewReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanApplyGroup() UnibeeApiBeanGroupPlanSelector`

GetPlanApplyGroup returns the PlanApplyGroup field if non-nil, zero value otherwise.

### GetPlanApplyGroupOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanApplyGroupOk() (*UnibeeApiBeanGroupPlanSelector, bool)`

GetPlanApplyGroupOk returns a tuple with the PlanApplyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountNewReq) SetPlanApplyGroup(v UnibeeApiBeanGroupPlanSelector)`

SetPlanApplyGroup sets PlanApplyGroup field to given value.

### HasPlanApplyGroup

`func (o *UnibeeApiMerchantDiscountNewReq) HasPlanApplyGroup() bool`

HasPlanApplyGroup returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiMerchantDiscountNewReq) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiMerchantDiscountNewReq) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantDiscountNewReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantDiscountNewReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantDiscountNewReq) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantDiscountNewReq) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantDiscountNewReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiMerchantDiscountNewReq) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiMerchantDiscountNewReq) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountNewReq) GetUpgradeLongPlanOnly() bool`

GetUpgradeLongPlanOnly returns the UpgradeLongPlanOnly field if non-nil, zero value otherwise.

### GetUpgradeLongPlanOnlyOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetUpgradeLongPlanOnlyOk() (*bool, bool)`

GetUpgradeLongPlanOnlyOk returns a tuple with the UpgradeLongPlanOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountNewReq) SetUpgradeLongPlanOnly(v bool)`

SetUpgradeLongPlanOnly sets UpgradeLongPlanOnly field to given value.

### HasUpgradeLongPlanOnly

`func (o *UnibeeApiMerchantDiscountNewReq) HasUpgradeLongPlanOnly() bool`

HasUpgradeLongPlanOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountNewReq) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiMerchantDiscountNewReq) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiMerchantDiscountNewReq) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiMerchantDiscountNewReq) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiMerchantDiscountNewReq) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiMerchantDiscountNewReq) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiMerchantDiscountNewReq) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiMerchantDiscountNewReq) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiMerchantDiscountNewReq) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiMerchantDiscountNewReq) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


