# UnibeeApiMerchantPlanAddonsBindingReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **int64** | Action Type，0-override,1-add，2-delete | 
**AddonIds** | **[]int64** | Plan Ids Of Addon Type | 
**PlanId** | **int64** | PlanID | 

## Methods

### NewUnibeeApiMerchantPlanAddonsBindingReq

`func NewUnibeeApiMerchantPlanAddonsBindingReq(action int64, addonIds []int64, planId int64, ) *UnibeeApiMerchantPlanAddonsBindingReq`

NewUnibeeApiMerchantPlanAddonsBindingReq instantiates a new UnibeeApiMerchantPlanAddonsBindingReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanAddonsBindingReqWithDefaults

`func NewUnibeeApiMerchantPlanAddonsBindingReqWithDefaults() *UnibeeApiMerchantPlanAddonsBindingReq`

NewUnibeeApiMerchantPlanAddonsBindingReqWithDefaults instantiates a new UnibeeApiMerchantPlanAddonsBindingReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAction() int64`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetActionOk() (*int64, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetAction(v int64)`

SetAction sets Action field to given value.


### GetAddonIds

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.


### GetPlanId

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


