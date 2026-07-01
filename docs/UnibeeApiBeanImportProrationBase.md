# UnibeeApiBeanImportProrationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPaidAmount** | Pointer to **int64** | Historical paid amount in minor units (cent), used directly as the proration credit base | [optional] 
**StartTime** | Pointer to **int64** | Optional. Historical period start time (utc seconds). When &gt; 0, overrides the proration window PeriodStart | [optional] 

## Methods

### NewUnibeeApiBeanImportProrationBase

`func NewUnibeeApiBeanImportProrationBase() *UnibeeApiBeanImportProrationBase`

NewUnibeeApiBeanImportProrationBase instantiates a new UnibeeApiBeanImportProrationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanImportProrationBaseWithDefaults

`func NewUnibeeApiBeanImportProrationBaseWithDefaults() *UnibeeApiBeanImportProrationBase`

NewUnibeeApiBeanImportProrationBaseWithDefaults instantiates a new UnibeeApiBeanImportProrationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPaidAmount

`func (o *UnibeeApiBeanImportProrationBase) GetLastPaidAmount() int64`

GetLastPaidAmount returns the LastPaidAmount field if non-nil, zero value otherwise.

### GetLastPaidAmountOk

`func (o *UnibeeApiBeanImportProrationBase) GetLastPaidAmountOk() (*int64, bool)`

GetLastPaidAmountOk returns a tuple with the LastPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaidAmount

`func (o *UnibeeApiBeanImportProrationBase) SetLastPaidAmount(v int64)`

SetLastPaidAmount sets LastPaidAmount field to given value.

### HasLastPaidAmount

`func (o *UnibeeApiBeanImportProrationBase) HasLastPaidAmount() bool`

HasLastPaidAmount returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiBeanImportProrationBase) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiBeanImportProrationBase) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiBeanImportProrationBase) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiBeanImportProrationBase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


