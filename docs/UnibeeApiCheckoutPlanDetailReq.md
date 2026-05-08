# UnibeeApiCheckoutPlanDetailReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Optional. When set and supported by merchant multi-currency config, plan amount, currency, and multi-trial prices are returned in this currency; otherwise default behavior | [optional] 
**PlanId** | **int64** | PlanId | 

## Methods

### NewUnibeeApiCheckoutPlanDetailReq

`func NewUnibeeApiCheckoutPlanDetailReq(planId int64, ) *UnibeeApiCheckoutPlanDetailReq`

NewUnibeeApiCheckoutPlanDetailReq instantiates a new UnibeeApiCheckoutPlanDetailReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutPlanDetailReqWithDefaults

`func NewUnibeeApiCheckoutPlanDetailReqWithDefaults() *UnibeeApiCheckoutPlanDetailReq`

NewUnibeeApiCheckoutPlanDetailReqWithDefaults instantiates a new UnibeeApiCheckoutPlanDetailReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiCheckoutPlanDetailReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiCheckoutPlanDetailReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiCheckoutPlanDetailReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiCheckoutPlanDetailReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiCheckoutPlanDetailReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiCheckoutPlanDetailReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiCheckoutPlanDetailReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


