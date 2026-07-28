# UnibeeApiBeanPlanTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Tag match mode: any|all. Default any when tagIds is non-empty. | [optional] 
**TagIds** | Pointer to **[]int64** | Plan tag ids. Omit planTagFilter or pass empty tagIds to skip tag filtering. | [optional] 

## Methods

### NewUnibeeApiBeanPlanTagFilter

`func NewUnibeeApiBeanPlanTagFilter() *UnibeeApiBeanPlanTagFilter`

NewUnibeeApiBeanPlanTagFilter instantiates a new UnibeeApiBeanPlanTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanTagFilterWithDefaults

`func NewUnibeeApiBeanPlanTagFilterWithDefaults() *UnibeeApiBeanPlanTagFilter`

NewUnibeeApiBeanPlanTagFilterWithDefaults instantiates a new UnibeeApiBeanPlanTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UnibeeApiBeanPlanTagFilter) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UnibeeApiBeanPlanTagFilter) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UnibeeApiBeanPlanTagFilter) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UnibeeApiBeanPlanTagFilter) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTagIds

`func (o *UnibeeApiBeanPlanTagFilter) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *UnibeeApiBeanPlanTagFilter) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *UnibeeApiBeanPlanTagFilter) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *UnibeeApiBeanPlanTagFilter) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


