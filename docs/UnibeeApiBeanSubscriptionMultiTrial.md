# UnibeeApiBeanSubscriptionMultiTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**CurrentIndex** | Pointer to **int32** |  | [optional] 
**Finished** | Pointer to **bool** |  | [optional] 
**PlanId** | Pointer to **int64** |  | [optional] 
**RegularPrice** | Pointer to **int64** | main plan recurring total ex tax in subscription currency; API only, not in link | [optional] 
**StartedAt** | Pointer to **int64** |  | [optional] 
**Trials** | Pointer to [**[]UnibeeApiBeanSubscriptionMultiTrialItem**](UnibeeApiBeanSubscriptionMultiTrialItem.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewUnibeeApiBeanSubscriptionMultiTrial

`func NewUnibeeApiBeanSubscriptionMultiTrial() *UnibeeApiBeanSubscriptionMultiTrial`

NewUnibeeApiBeanSubscriptionMultiTrial instantiates a new UnibeeApiBeanSubscriptionMultiTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionMultiTrialWithDefaults

`func NewUnibeeApiBeanSubscriptionMultiTrialWithDefaults() *UnibeeApiBeanSubscriptionMultiTrial`

NewUnibeeApiBeanSubscriptionMultiTrialWithDefaults instantiates a new UnibeeApiBeanSubscriptionMultiTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentIndex

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetCurrentIndex() int32`

GetCurrentIndex returns the CurrentIndex field if non-nil, zero value otherwise.

### GetCurrentIndexOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetCurrentIndexOk() (*int32, bool)`

GetCurrentIndexOk returns a tuple with the CurrentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIndex

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetCurrentIndex(v int32)`

SetCurrentIndex sets CurrentIndex field to given value.

### HasCurrentIndex

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasCurrentIndex() bool`

HasCurrentIndex returns a boolean if a field has been set.

### GetFinished

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetRegularPrice

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetRegularPrice() int64`

GetRegularPrice returns the RegularPrice field if non-nil, zero value otherwise.

### GetRegularPriceOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetRegularPriceOk() (*int64, bool)`

GetRegularPriceOk returns a tuple with the RegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPrice

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetRegularPrice(v int64)`

SetRegularPrice sets RegularPrice field to given value.

### HasRegularPrice

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasRegularPrice() bool`

HasRegularPrice returns a boolean if a field has been set.

### GetStartedAt

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTrials

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetTrials() []UnibeeApiBeanSubscriptionMultiTrialItem`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetTrialsOk() (*[]UnibeeApiBeanSubscriptionMultiTrialItem, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetTrials(v []UnibeeApiBeanSubscriptionMultiTrialItem)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeApiBeanSubscriptionMultiTrial) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeApiBeanSubscriptionMultiTrial) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeApiBeanSubscriptionMultiTrial) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


