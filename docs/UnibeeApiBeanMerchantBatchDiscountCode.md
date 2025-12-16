# UnibeeApiBeanMerchantBatchDiscountCode

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
**InvoiceId** | Pointer to **string** | Invoice ID where the code was used | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，&gt; 0, Deleted, the deleted utc time | [optional] 
**IsRedeemed** | Pointer to **bool** | Whether the code has been redeemed | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**ParentTemplateCode** | Pointer to **string** | parentTemplateCode | [optional] 
**PaymentId** | Pointer to **string** | Payment ID where the code was used | [optional] 
**PlanApplyGroup** | Pointer to [**UnibeeApiBeanGroupPlanSelector**](UnibeeApiBeanGroupPlanSelector.md) |  | [optional] 
**PlanApplyType** | Pointer to **int32** | plan apply type, 0-apply for all, 1-apply for plans specified, 2-exclude for plans specified, 3-Apply to Plans by Groups(Billing Period Included), 4-Apply to Plans except by Groups(Billing Period Included) | [optional] 
**PlanIds** | Pointer to **[]int64** | Ids of plan which discount code can effect, default effect all plans if not set | [optional] 
**Quantity** | Pointer to **int64** | quantity of code, 0-no limit | [optional] 
**RedeemedAt** | Pointer to **int64** | UTC timestamp when code was redeemed, 0 if not redeemed | [optional] 
**RedeemedByEmail** | Pointer to **string** | Email of user who redeemed the code | [optional] 
**StartTime** | Pointer to **int64** | start of discount available utc time | [optional] 
**Status** | Pointer to **int32** | status, 1-editable, 2-active, 3-deactivate, 4-expire, 10-archive | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription ID where the code was used | [optional] 
**Type** | Pointer to **int32** | type, 1-external discount code, 2-batch code template, 3-batch code | [optional] 
**UpgradeLongerOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except upgrade to longer plan if set true | [optional] 
**UpgradeOnly** | Pointer to **bool** | AdvanceConfig, 0-false,1-true, will forbid for all except same interval upgrade action if set true | [optional] 
**UserLimit** | Pointer to **int32** | AdvanceConfig, The limit of every customer can apply, the recurring apply not involved, 0-unlimited | [optional] 
**UserScope** | Pointer to **int32** | AdvanceConfig, Apply user scope,0-for all, 1-for only new user, 2-for only renewals, renewals is upgrade&amp;downgrade&amp;renew | [optional] 

## Methods

### NewUnibeeApiBeanMerchantBatchDiscountCode

`func NewUnibeeApiBeanMerchantBatchDiscountCode() *UnibeeApiBeanMerchantBatchDiscountCode`

NewUnibeeApiBeanMerchantBatchDiscountCode instantiates a new UnibeeApiBeanMerchantBatchDiscountCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantBatchDiscountCodeWithDefaults

`func NewUnibeeApiBeanMerchantBatchDiscountCodeWithDefaults() *UnibeeApiBeanMerchantBatchDiscountCode`

NewUnibeeApiBeanMerchantBatchDiscountCodeWithDefaults instantiates a new UnibeeApiBeanMerchantBatchDiscountCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetAdvance() bool`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetAdvanceOk() (*bool, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetAdvance(v bool)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCodePrefix() string`

GetCodePrefix returns the CodePrefix field if non-nil, zero value otherwise.

### GetCodePrefixOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCodePrefixOk() (*string, bool)`

GetCodePrefixOk returns a tuple with the CodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetCodePrefix(v string)`

SetCodePrefix sets CodePrefix field to given value.

### HasCodePrefix

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasCodePrefix() bool`

HasCodePrefix returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsRedeemed

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetIsRedeemed() bool`

GetIsRedeemed returns the IsRedeemed field if non-nil, zero value otherwise.

### GetIsRedeemedOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetIsRedeemedOk() (*bool, bool)`

GetIsRedeemedOk returns a tuple with the IsRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedeemed

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetIsRedeemed(v bool)`

SetIsRedeemed sets IsRedeemed field to given value.

### HasIsRedeemed

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasIsRedeemed() bool`

HasIsRedeemed returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTemplateCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetParentTemplateCode() string`

GetParentTemplateCode returns the ParentTemplateCode field if non-nil, zero value otherwise.

### GetParentTemplateCodeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetParentTemplateCodeOk() (*string, bool)`

GetParentTemplateCodeOk returns a tuple with the ParentTemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTemplateCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetParentTemplateCode(v string)`

SetParentTemplateCode sets ParentTemplateCode field to given value.

### HasParentTemplateCode

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasParentTemplateCode() bool`

HasParentTemplateCode returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanApplyGroup() UnibeeApiBeanGroupPlanSelector`

GetPlanApplyGroup returns the PlanApplyGroup field if non-nil, zero value otherwise.

### GetPlanApplyGroupOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanApplyGroupOk() (*UnibeeApiBeanGroupPlanSelector, bool)`

GetPlanApplyGroupOk returns a tuple with the PlanApplyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetPlanApplyGroup(v UnibeeApiBeanGroupPlanSelector)`

SetPlanApplyGroup sets PlanApplyGroup field to given value.

### HasPlanApplyGroup

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasPlanApplyGroup() bool`

HasPlanApplyGroup returns a boolean if a field has been set.

### GetPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanApplyType() int32`

GetPlanApplyType returns the PlanApplyType field if non-nil, zero value otherwise.

### GetPlanApplyTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanApplyTypeOk() (*int32, bool)`

GetPlanApplyTypeOk returns a tuple with the PlanApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetPlanApplyType(v int32)`

SetPlanApplyType sets PlanApplyType field to given value.

### HasPlanApplyType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasPlanApplyType() bool`

HasPlanApplyType returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRedeemedAt

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetRedeemedAt() int64`

GetRedeemedAt returns the RedeemedAt field if non-nil, zero value otherwise.

### GetRedeemedAtOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetRedeemedAtOk() (*int64, bool)`

GetRedeemedAtOk returns a tuple with the RedeemedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedAt

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetRedeemedAt(v int64)`

SetRedeemedAt sets RedeemedAt field to given value.

### HasRedeemedAt

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasRedeemedAt() bool`

HasRedeemedAt returns a boolean if a field has been set.

### GetRedeemedByEmail

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetRedeemedByEmail() string`

GetRedeemedByEmail returns the RedeemedByEmail field if non-nil, zero value otherwise.

### GetRedeemedByEmailOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetRedeemedByEmailOk() (*string, bool)`

GetRedeemedByEmailOk returns a tuple with the RedeemedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedByEmail

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetRedeemedByEmail(v string)`

SetRedeemedByEmail sets RedeemedByEmail field to given value.

### HasRedeemedByEmail

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasRedeemedByEmail() bool`

HasRedeemedByEmail returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUpgradeLongerOnly() bool`

GetUpgradeLongerOnly returns the UpgradeLongerOnly field if non-nil, zero value otherwise.

### GetUpgradeLongerOnlyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUpgradeLongerOnlyOk() (*bool, bool)`

GetUpgradeLongerOnlyOk returns a tuple with the UpgradeLongerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetUpgradeLongerOnly(v bool)`

SetUpgradeLongerOnly sets UpgradeLongerOnly field to given value.

### HasUpgradeLongerOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasUpgradeLongerOnly() bool`

HasUpgradeLongerOnly returns a boolean if a field has been set.

### GetUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUpgradeOnly() bool`

GetUpgradeOnly returns the UpgradeOnly field if non-nil, zero value otherwise.

### GetUpgradeOnlyOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUpgradeOnlyOk() (*bool, bool)`

GetUpgradeOnlyOk returns a tuple with the UpgradeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetUpgradeOnly(v bool)`

SetUpgradeOnly sets UpgradeOnly field to given value.

### HasUpgradeOnly

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasUpgradeOnly() bool`

HasUpgradeOnly returns a boolean if a field has been set.

### GetUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUserScope() int32`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) GetUserScopeOk() (*int32, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) SetUserScope(v int32)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *UnibeeApiBeanMerchantBatchDiscountCode) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


