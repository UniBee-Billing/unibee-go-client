# UnibeeApiBeanPlanMultiTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalCount** | Pointer to **int64** | for stage 2+: must be &gt; 0 when intervalUnit is set; for stage 1 (index 0) used with intervalUnit to set TrialDurationTime at save | [optional] 
**IntervalUnit** | Pointer to **string** | for stage 2+: billing; for stage 1 (index 0) converted to plan TrialDurationTime at save (day/week/month≈30d/year≈365d) | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | optional metadata for this trial stage; persisted to subscription multi-trial snapshot and exposed in webhook payloads | [optional] 
**Name** | Pointer to **string** | optional display name for this trial stage; when set, used instead of plan name in multi-trial invoices | [optional] 
**TrialCompareAtPrice** | Pointer to **int64** | optional compare-at / strikethrough display price in minor units (plan currency); UI only, not used in billing | [optional] 
**TrialDuration** | Pointer to **int64** | trial duration in seconds; first stage: at least 86400 unless valid intervalUnit+intervalCount is set | [optional] 
**TrialPrice** | Pointer to **int64** | trial price; only the first trial is allowed to be 0, subsequent trials must be &gt; 0 | [optional] 

## Methods

### NewUnibeeApiBeanPlanMultiTrial

`func NewUnibeeApiBeanPlanMultiTrial() *UnibeeApiBeanPlanMultiTrial`

NewUnibeeApiBeanPlanMultiTrial instantiates a new UnibeeApiBeanPlanMultiTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanMultiTrialWithDefaults

`func NewUnibeeApiBeanPlanMultiTrialWithDefaults() *UnibeeApiBeanPlanMultiTrial`

NewUnibeeApiBeanPlanMultiTrialWithDefaults instantiates a new UnibeeApiBeanPlanMultiTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalCount

`func (o *UnibeeApiBeanPlanMultiTrial) GetIntervalCount() int64`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetIntervalCountOk() (*int64, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiBeanPlanMultiTrial) SetIntervalCount(v int64)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiBeanPlanMultiTrial) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiBeanPlanMultiTrial) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiBeanPlanMultiTrial) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiBeanPlanMultiTrial) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanPlanMultiTrial) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanPlanMultiTrial) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanPlanMultiTrial) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanPlanMultiTrial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanPlanMultiTrial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanPlanMultiTrial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrialCompareAtPrice

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialCompareAtPrice() int64`

GetTrialCompareAtPrice returns the TrialCompareAtPrice field if non-nil, zero value otherwise.

### GetTrialCompareAtPriceOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialCompareAtPriceOk() (*int64, bool)`

GetTrialCompareAtPriceOk returns a tuple with the TrialCompareAtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCompareAtPrice

`func (o *UnibeeApiBeanPlanMultiTrial) SetTrialCompareAtPrice(v int64)`

SetTrialCompareAtPrice sets TrialCompareAtPrice field to given value.

### HasTrialCompareAtPrice

`func (o *UnibeeApiBeanPlanMultiTrial) HasTrialCompareAtPrice() bool`

HasTrialCompareAtPrice returns a boolean if a field has been set.

### GetTrialDuration

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialDuration() int64`

GetTrialDuration returns the TrialDuration field if non-nil, zero value otherwise.

### GetTrialDurationOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialDurationOk() (*int64, bool)`

GetTrialDurationOk returns a tuple with the TrialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDuration

`func (o *UnibeeApiBeanPlanMultiTrial) SetTrialDuration(v int64)`

SetTrialDuration sets TrialDuration field to given value.

### HasTrialDuration

`func (o *UnibeeApiBeanPlanMultiTrial) HasTrialDuration() bool`

HasTrialDuration returns a boolean if a field has been set.

### GetTrialPrice

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialPrice() int64`

GetTrialPrice returns the TrialPrice field if non-nil, zero value otherwise.

### GetTrialPriceOk

`func (o *UnibeeApiBeanPlanMultiTrial) GetTrialPriceOk() (*int64, bool)`

GetTrialPriceOk returns a tuple with the TrialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPrice

`func (o *UnibeeApiBeanPlanMultiTrial) SetTrialPrice(v int64)`

SetTrialPrice sets TrialPrice field to given value.

### HasTrialPrice

`func (o *UnibeeApiBeanPlanMultiTrial) HasTrialPrice() bool`

HasTrialPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


