# UnibeeApiBeanPublicPlanMultiTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalCount** | Pointer to **int64** | for stage 2+: must be &gt; 0 when intervalUnit is set | [optional] 
**IntervalUnit** | Pointer to **string** | for stage 2+: billing; for stage 1 converted to TrialDurationTime | [optional] 
**Name** | Pointer to **string** | optional display name for this trial stage | [optional] 
**TrialCompareAtPrice** | Pointer to **int64** | optional compare-at / strikethrough display price in minor units | [optional] 
**TrialDuration** | Pointer to **int64** | trial duration in seconds | [optional] 
**TrialPrice** | Pointer to **int64** | trial price; only the first trial is allowed to be 0 | [optional] 

## Methods

### NewUnibeeApiBeanPublicPlanMultiTrial

`func NewUnibeeApiBeanPublicPlanMultiTrial() *UnibeeApiBeanPublicPlanMultiTrial`

NewUnibeeApiBeanPublicPlanMultiTrial instantiates a new UnibeeApiBeanPublicPlanMultiTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPublicPlanMultiTrialWithDefaults

`func NewUnibeeApiBeanPublicPlanMultiTrialWithDefaults() *UnibeeApiBeanPublicPlanMultiTrial`

NewUnibeeApiBeanPublicPlanMultiTrialWithDefaults instantiates a new UnibeeApiBeanPublicPlanMultiTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalCount

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetIntervalCount() int64`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetIntervalCountOk() (*int64, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetIntervalCount(v int64)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrialCompareAtPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialCompareAtPrice() int64`

GetTrialCompareAtPrice returns the TrialCompareAtPrice field if non-nil, zero value otherwise.

### GetTrialCompareAtPriceOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialCompareAtPriceOk() (*int64, bool)`

GetTrialCompareAtPriceOk returns a tuple with the TrialCompareAtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCompareAtPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetTrialCompareAtPrice(v int64)`

SetTrialCompareAtPrice sets TrialCompareAtPrice field to given value.

### HasTrialCompareAtPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasTrialCompareAtPrice() bool`

HasTrialCompareAtPrice returns a boolean if a field has been set.

### GetTrialDuration

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialDuration() int64`

GetTrialDuration returns the TrialDuration field if non-nil, zero value otherwise.

### GetTrialDurationOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialDurationOk() (*int64, bool)`

GetTrialDurationOk returns a tuple with the TrialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDuration

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetTrialDuration(v int64)`

SetTrialDuration sets TrialDuration field to given value.

### HasTrialDuration

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasTrialDuration() bool`

HasTrialDuration returns a boolean if a field has been set.

### GetTrialPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialPrice() int64`

GetTrialPrice returns the TrialPrice field if non-nil, zero value otherwise.

### GetTrialPriceOk

`func (o *UnibeeApiBeanPublicPlanMultiTrial) GetTrialPriceOk() (*int64, bool)`

GetTrialPriceOk returns a tuple with the TrialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) SetTrialPrice(v int64)`

SetTrialPrice sets TrialPrice field to given value.

### HasTrialPrice

`func (o *UnibeeApiBeanPublicPlanMultiTrial) HasTrialPrice() bool`

HasTrialPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


