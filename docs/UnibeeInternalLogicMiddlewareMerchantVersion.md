# UnibeeInternalLogicMiddlewareMerchantVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int64** | EndTime,UTC, The End Time Of Plan,0 for free | [optional] 
**Expired** | Pointer to **bool** | Expired | [optional] 
**IsPaid** | Pointer to **bool** | IsPaid | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicMiddlewarePlan**](UnibeeInternalLogicMiddlewarePlan.md) |  | [optional] 
**StartTime** | Pointer to **int64** | StartTime,UTC, The Start Time Of Plan,0 for free | [optional] 

## Methods

### NewUnibeeInternalLogicMiddlewareMerchantVersion

`func NewUnibeeInternalLogicMiddlewareMerchantVersion() *UnibeeInternalLogicMiddlewareMerchantVersion`

NewUnibeeInternalLogicMiddlewareMerchantVersion instantiates a new UnibeeInternalLogicMiddlewareMerchantVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicMiddlewareMerchantVersionWithDefaults

`func NewUnibeeInternalLogicMiddlewareMerchantVersionWithDefaults() *UnibeeInternalLogicMiddlewareMerchantVersion`

NewUnibeeInternalLogicMiddlewareMerchantVersionWithDefaults instantiates a new UnibeeInternalLogicMiddlewareMerchantVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpired

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetIsPaid

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetName

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetPlan() UnibeeInternalLogicMiddlewarePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetPlanOk() (*UnibeeInternalLogicMiddlewarePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetPlan(v UnibeeInternalLogicMiddlewarePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeInternalLogicMiddlewareMerchantVersion) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


