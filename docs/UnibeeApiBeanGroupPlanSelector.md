# UnibeeApiBeanGroupPlanSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **[]string** |  | [optional] 
**GroupPlanIntervalSelector** | Pointer to [**[]UnibeeApiBeanGroupPlanIntervalSelector**](UnibeeApiBeanGroupPlanIntervalSelector.md) |  | [optional] 
**Type** | Pointer to **[]int32** | 1-main plan，2-addon plan,3-onetime | [optional] 

## Methods

### NewUnibeeApiBeanGroupPlanSelector

`func NewUnibeeApiBeanGroupPlanSelector() *UnibeeApiBeanGroupPlanSelector`

NewUnibeeApiBeanGroupPlanSelector instantiates a new UnibeeApiBeanGroupPlanSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanGroupPlanSelectorWithDefaults

`func NewUnibeeApiBeanGroupPlanSelectorWithDefaults() *UnibeeApiBeanGroupPlanSelector`

NewUnibeeApiBeanGroupPlanSelectorWithDefaults instantiates a new UnibeeApiBeanGroupPlanSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiBeanGroupPlanSelector) GetCurrency() []string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanGroupPlanSelector) GetCurrencyOk() (*[]string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanGroupPlanSelector) SetCurrency(v []string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanGroupPlanSelector) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGroupPlanIntervalSelector

`func (o *UnibeeApiBeanGroupPlanSelector) GetGroupPlanIntervalSelector() []UnibeeApiBeanGroupPlanIntervalSelector`

GetGroupPlanIntervalSelector returns the GroupPlanIntervalSelector field if non-nil, zero value otherwise.

### GetGroupPlanIntervalSelectorOk

`func (o *UnibeeApiBeanGroupPlanSelector) GetGroupPlanIntervalSelectorOk() (*[]UnibeeApiBeanGroupPlanIntervalSelector, bool)`

GetGroupPlanIntervalSelectorOk returns a tuple with the GroupPlanIntervalSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPlanIntervalSelector

`func (o *UnibeeApiBeanGroupPlanSelector) SetGroupPlanIntervalSelector(v []UnibeeApiBeanGroupPlanIntervalSelector)`

SetGroupPlanIntervalSelector sets GroupPlanIntervalSelector field to given value.

### HasGroupPlanIntervalSelector

`func (o *UnibeeApiBeanGroupPlanSelector) HasGroupPlanIntervalSelector() bool`

HasGroupPlanIntervalSelector returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanGroupPlanSelector) GetType() []int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanGroupPlanSelector) GetTypeOk() (*[]int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanGroupPlanSelector) SetType(v []int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanGroupPlanSelector) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


