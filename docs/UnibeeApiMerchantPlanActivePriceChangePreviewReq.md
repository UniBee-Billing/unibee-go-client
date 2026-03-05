# UnibeeApiMerchantPlanActivePriceChangePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAmount** | **int64** | New amount (in smallest currency unit, e.g. cents) | 
**PlanId** | **int64** | Target planId, must be active | 

## Methods

### NewUnibeeApiMerchantPlanActivePriceChangePreviewReq

`func NewUnibeeApiMerchantPlanActivePriceChangePreviewReq(newAmount int64, planId int64, ) *UnibeeApiMerchantPlanActivePriceChangePreviewReq`

NewUnibeeApiMerchantPlanActivePriceChangePreviewReq instantiates a new UnibeeApiMerchantPlanActivePriceChangePreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanActivePriceChangePreviewReqWithDefaults

`func NewUnibeeApiMerchantPlanActivePriceChangePreviewReqWithDefaults() *UnibeeApiMerchantPlanActivePriceChangePreviewReq`

NewUnibeeApiMerchantPlanActivePriceChangePreviewReqWithDefaults instantiates a new UnibeeApiMerchantPlanActivePriceChangePreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) GetNewAmount() int64`

GetNewAmount returns the NewAmount field if non-nil, zero value otherwise.

### GetNewAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) GetNewAmountOk() (*int64, bool)`

GetNewAmountOk returns a tuple with the NewAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) SetNewAmount(v int64)`

SetNewAmount sets NewAmount field to given value.


### GetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


