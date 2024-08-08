# UnibeeApiBeanExternalDiscountParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CycleLimit** | Pointer to **int32** | the count limitation of subscription recurring cycle, recurring need enable if cycleLimit set | [optional] 
**DiscountAmount** | Pointer to **int32** | Amount of discount | [optional] 
**DiscountPercentage** | Pointer to **int32** | Percentage of discount, 100&#x3D;1%, ignore if discountAmount set | [optional] 
**EndTime** | Pointer to **int32** | end of discount available utc time | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**Recurring** | Pointer to **bool** | Discount recurring enable, default false | [optional] 

## Methods

### NewUnibeeApiBeanExternalDiscountParam

`func NewUnibeeApiBeanExternalDiscountParam() *UnibeeApiBeanExternalDiscountParam`

NewUnibeeApiBeanExternalDiscountParam instantiates a new UnibeeApiBeanExternalDiscountParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanExternalDiscountParamWithDefaults

`func NewUnibeeApiBeanExternalDiscountParamWithDefaults() *UnibeeApiBeanExternalDiscountParam`

NewUnibeeApiBeanExternalDiscountParamWithDefaults instantiates a new UnibeeApiBeanExternalDiscountParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycleLimit

`func (o *UnibeeApiBeanExternalDiscountParam) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanExternalDiscountParam) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanExternalDiscountParam) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanExternalDiscountParam) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanExternalDiscountParam) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountPercentage() int32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountPercentageOk() (*int32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanExternalDiscountParam) SetDiscountPercentage(v int32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanExternalDiscountParam) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiBeanExternalDiscountParam) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiBeanExternalDiscountParam) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiBeanExternalDiscountParam) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanExternalDiscountParam) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanExternalDiscountParam) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanExternalDiscountParam) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecurring

`func (o *UnibeeApiBeanExternalDiscountParam) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *UnibeeApiBeanExternalDiscountParam) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *UnibeeApiBeanExternalDiscountParam) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *UnibeeApiBeanExternalDiscountParam) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


