# UnibeeApiMerchantPlanDeleteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPlanId** | Pointer to **string** | The ExternalId of plan, either externalPlanId or planId should be set | [optional] 
**PlanId** | Pointer to **int64** | The Id of plan, either planId or externalPlanId should be set | [optional] 

## Methods

### NewUnibeeApiMerchantPlanDeleteReq

`func NewUnibeeApiMerchantPlanDeleteReq() *UnibeeApiMerchantPlanDeleteReq`

NewUnibeeApiMerchantPlanDeleteReq instantiates a new UnibeeApiMerchantPlanDeleteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanDeleteReqWithDefaults

`func NewUnibeeApiMerchantPlanDeleteReqWithDefaults() *UnibeeApiMerchantPlanDeleteReq`

NewUnibeeApiMerchantPlanDeleteReqWithDefaults instantiates a new UnibeeApiMerchantPlanDeleteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantPlanDeleteReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanDeleteReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantPlanDeleteReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


