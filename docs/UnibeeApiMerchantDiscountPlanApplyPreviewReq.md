# UnibeeApiMerchantDiscountPlanApplyPreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The discount&#39;s unique code, customize by merchant | 
**Email** | Pointer to **string** | Email | [optional] 
**ExternalPlanId** | Pointer to **string** | The externalId of plan which code to apply, either planId or externalPlanId is needed | [optional] 
**IsChangeToLongPlan** | Pointer to **bool** | IsChangeToLongPlan | [optional] 
**IsChangeToSameIntervalPlan** | Pointer to **bool** | IsChangeToSameIntervalPlan | [optional] 
**IsUpgrade** | Pointer to **bool** | IsUpgrade | [optional] 
**PlanId** | Pointer to **int64** | The id of plan which code to apply, either planId or externalPlanId is needed | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountPlanApplyPreviewReq

`func NewUnibeeApiMerchantDiscountPlanApplyPreviewReq(code string, ) *UnibeeApiMerchantDiscountPlanApplyPreviewReq`

NewUnibeeApiMerchantDiscountPlanApplyPreviewReq instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountPlanApplyPreviewReqWithDefaults

`func NewUnibeeApiMerchantDiscountPlanApplyPreviewReqWithDefaults() *UnibeeApiMerchantDiscountPlanApplyPreviewReq`

NewUnibeeApiMerchantDiscountPlanApplyPreviewReqWithDefaults instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetCode(v string)`

SetCode sets Code field to given value.


### GetEmail

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetIsChangeToLongPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToLongPlan() bool`

GetIsChangeToLongPlan returns the IsChangeToLongPlan field if non-nil, zero value otherwise.

### GetIsChangeToLongPlanOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToLongPlanOk() (*bool, bool)`

GetIsChangeToLongPlanOk returns a tuple with the IsChangeToLongPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChangeToLongPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsChangeToLongPlan(v bool)`

SetIsChangeToLongPlan sets IsChangeToLongPlan field to given value.

### HasIsChangeToLongPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsChangeToLongPlan() bool`

HasIsChangeToLongPlan returns a boolean if a field has been set.

### GetIsChangeToSameIntervalPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToSameIntervalPlan() bool`

GetIsChangeToSameIntervalPlan returns the IsChangeToSameIntervalPlan field if non-nil, zero value otherwise.

### GetIsChangeToSameIntervalPlanOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToSameIntervalPlanOk() (*bool, bool)`

GetIsChangeToSameIntervalPlanOk returns a tuple with the IsChangeToSameIntervalPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChangeToSameIntervalPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsChangeToSameIntervalPlan(v bool)`

SetIsChangeToSameIntervalPlan sets IsChangeToSameIntervalPlan field to given value.

### HasIsChangeToSameIntervalPlan

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsChangeToSameIntervalPlan() bool`

HasIsChangeToSameIntervalPlan returns a boolean if a field has been set.

### GetIsUpgrade

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsUpgrade() bool`

GetIsUpgrade returns the IsUpgrade field if non-nil, zero value otherwise.

### GetIsUpgradeOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsUpgradeOk() (*bool, bool)`

GetIsUpgradeOk returns a tuple with the IsUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgrade

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsUpgrade(v bool)`

SetIsUpgrade sets IsUpgrade field to given value.

### HasIsUpgrade

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsUpgrade() bool`

HasIsUpgrade returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


