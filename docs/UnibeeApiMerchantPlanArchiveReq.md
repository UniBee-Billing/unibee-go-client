# UnibeeApiMerchantPlanArchiveReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPlanId** | Pointer to **string** | The ExternalId of plan, either externalPlanId or planId should be set | [optional] 
**HardArchive** | Pointer to **bool** | Hard Archive | [optional] 
**MigrationPlanId** | Pointer to **int64** | Optional. Hard archive only: persist a target main plan id for later subscription migration. Stored in plan additional settings JSON (same place as autoUpgradePlanId, separate key). 0 or omit &#x3D; no migration target. Does not migrate subscriptions in this request. | [optional] 
**PlanId** | Pointer to **int64** | The Id of plan, either planId or externalPlanId should be set | [optional] 

## Methods

### NewUnibeeApiMerchantPlanArchiveReq

`func NewUnibeeApiMerchantPlanArchiveReq() *UnibeeApiMerchantPlanArchiveReq`

NewUnibeeApiMerchantPlanArchiveReq instantiates a new UnibeeApiMerchantPlanArchiveReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanArchiveReqWithDefaults

`func NewUnibeeApiMerchantPlanArchiveReqWithDefaults() *UnibeeApiMerchantPlanArchiveReq`

NewUnibeeApiMerchantPlanArchiveReqWithDefaults instantiates a new UnibeeApiMerchantPlanArchiveReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantPlanArchiveReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetHardArchive

`func (o *UnibeeApiMerchantPlanArchiveReq) GetHardArchive() bool`

GetHardArchive returns the HardArchive field if non-nil, zero value otherwise.

### GetHardArchiveOk

`func (o *UnibeeApiMerchantPlanArchiveReq) GetHardArchiveOk() (*bool, bool)`

GetHardArchiveOk returns a tuple with the HardArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardArchive

`func (o *UnibeeApiMerchantPlanArchiveReq) SetHardArchive(v bool)`

SetHardArchive sets HardArchive field to given value.

### HasHardArchive

`func (o *UnibeeApiMerchantPlanArchiveReq) HasHardArchive() bool`

HasHardArchive returns a boolean if a field has been set.

### GetMigrationPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) GetMigrationPlanId() int64`

GetMigrationPlanId returns the MigrationPlanId field if non-nil, zero value otherwise.

### GetMigrationPlanIdOk

`func (o *UnibeeApiMerchantPlanArchiveReq) GetMigrationPlanIdOk() (*int64, bool)`

GetMigrationPlanIdOk returns a tuple with the MigrationPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) SetMigrationPlanId(v int64)`

SetMigrationPlanId sets MigrationPlanId field to given value.

### HasMigrationPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) HasMigrationPlanId() bool`

HasMigrationPlanId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanArchiveReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantPlanArchiveReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


