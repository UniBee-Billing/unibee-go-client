# UnibeeApiBeanCreateSubmitDenyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | Email globs, match any one. Local/domain: exact, *, prefix*, *suffix, *contains*, pre*suf (max two * each). Domain also supports [hex:N]suffix, e.g. *@[hex:32].com. [random] matches a part with no aligned 4+ letter English word and high randomness, e.g. [random]@[random].com | [optional] 
**Enabled** | Pointer to **bool** | Optional. Default true. False skips the rule | [optional] 
**Id** | Pointer to **string** | Optional. Deny log only | [optional] 
**PlanIds** | Pointer to **[]int64** | Optional. Empty &#x3D; all plans | [optional] 

## Methods

### NewUnibeeApiBeanCreateSubmitDenyRule

`func NewUnibeeApiBeanCreateSubmitDenyRule() *UnibeeApiBeanCreateSubmitDenyRule`

NewUnibeeApiBeanCreateSubmitDenyRule instantiates a new UnibeeApiBeanCreateSubmitDenyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreateSubmitDenyRuleWithDefaults

`func NewUnibeeApiBeanCreateSubmitDenyRuleWithDefaults() *UnibeeApiBeanCreateSubmitDenyRule`

NewUnibeeApiBeanCreateSubmitDenyRuleWithDefaults instantiates a new UnibeeApiBeanCreateSubmitDenyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UnibeeApiBeanCreateSubmitDenyRule) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UnibeeApiBeanCreateSubmitDenyRule) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetEnabled

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnibeeApiBeanCreateSubmitDenyRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UnibeeApiBeanCreateSubmitDenyRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanCreateSubmitDenyRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanCreateSubmitDenyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiBeanCreateSubmitDenyRule) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiBeanCreateSubmitDenyRule) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiBeanCreateSubmitDenyRule) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


