# UnibeeApiBeanUserChurnBulkCheckItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChurnHistoryEpisodeCount** | Pointer to **int32** | user_churn_history episode count | [optional] 
**CurrentlyChurned** | Pointer to **bool** | True when found and there is no Active(2)/Incomplete(7) UniBee Control subscription | [optional] 
**EffectiveSubscriptions** | Pointer to [**[]UnibeeApiBeanUserChurnBulkCheckSub**](UnibeeApiBeanUserChurnBulkCheckSub.md) | Active/Incomplete UniBee Control subscriptions | [optional] 
**ExternalUserIds** | Pointer to **[]string** | Matched external user ids | [optional] 
**Found** | Pointer to **bool** | Whether a merchant user was found | [optional] 
**InputEmail** | Pointer to **string** | Original input email | [optional] 
**LatestChurn** | Pointer to [**UnibeeApiBeanUserChurnBulkCheckEpisode**](UnibeeApiBeanUserChurnBulkCheckEpisode.md) |  | [optional] 
**LatestSubscription** | Pointer to [**UnibeeApiBeanUserChurnBulkCheckSub**](UnibeeApiBeanUserChurnBulkCheckSub.md) |  | [optional] 
**LookupEmail** | Pointer to **string** | Email used to find the user, after light alias (trailing dots) | [optional] 
**Note** | Pointer to **string** | open_churn_history&#x3D;... when there is an open episode | [optional] 
**OpenChurnStatuses** | Pointer to **[]string** | pending/confirmed episodes that are not reactivated | [optional] 
**UserAccountSubscriptionStatuses** | Pointer to **[]string** | userId:status from user_account.subscription_status | [optional] 
**UserIds** | Pointer to **[]int64** | Matched user ids | [optional] 

## Methods

### NewUnibeeApiBeanUserChurnBulkCheckItem

`func NewUnibeeApiBeanUserChurnBulkCheckItem() *UnibeeApiBeanUserChurnBulkCheckItem`

NewUnibeeApiBeanUserChurnBulkCheckItem instantiates a new UnibeeApiBeanUserChurnBulkCheckItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserChurnBulkCheckItemWithDefaults

`func NewUnibeeApiBeanUserChurnBulkCheckItemWithDefaults() *UnibeeApiBeanUserChurnBulkCheckItem`

NewUnibeeApiBeanUserChurnBulkCheckItemWithDefaults instantiates a new UnibeeApiBeanUserChurnBulkCheckItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChurnHistoryEpisodeCount

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetChurnHistoryEpisodeCount() int32`

GetChurnHistoryEpisodeCount returns the ChurnHistoryEpisodeCount field if non-nil, zero value otherwise.

### GetChurnHistoryEpisodeCountOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetChurnHistoryEpisodeCountOk() (*int32, bool)`

GetChurnHistoryEpisodeCountOk returns a tuple with the ChurnHistoryEpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnHistoryEpisodeCount

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetChurnHistoryEpisodeCount(v int32)`

SetChurnHistoryEpisodeCount sets ChurnHistoryEpisodeCount field to given value.

### HasChurnHistoryEpisodeCount

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasChurnHistoryEpisodeCount() bool`

HasChurnHistoryEpisodeCount returns a boolean if a field has been set.

### GetCurrentlyChurned

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetCurrentlyChurned() bool`

GetCurrentlyChurned returns the CurrentlyChurned field if non-nil, zero value otherwise.

### GetCurrentlyChurnedOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetCurrentlyChurnedOk() (*bool, bool)`

GetCurrentlyChurnedOk returns a tuple with the CurrentlyChurned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyChurned

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetCurrentlyChurned(v bool)`

SetCurrentlyChurned sets CurrentlyChurned field to given value.

### HasCurrentlyChurned

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasCurrentlyChurned() bool`

HasCurrentlyChurned returns a boolean if a field has been set.

### GetEffectiveSubscriptions

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetEffectiveSubscriptions() []UnibeeApiBeanUserChurnBulkCheckSub`

GetEffectiveSubscriptions returns the EffectiveSubscriptions field if non-nil, zero value otherwise.

### GetEffectiveSubscriptionsOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetEffectiveSubscriptionsOk() (*[]UnibeeApiBeanUserChurnBulkCheckSub, bool)`

GetEffectiveSubscriptionsOk returns a tuple with the EffectiveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSubscriptions

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetEffectiveSubscriptions(v []UnibeeApiBeanUserChurnBulkCheckSub)`

SetEffectiveSubscriptions sets EffectiveSubscriptions field to given value.

### HasEffectiveSubscriptions

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasEffectiveSubscriptions() bool`

HasEffectiveSubscriptions returns a boolean if a field has been set.

### GetExternalUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetExternalUserIds() []string`

GetExternalUserIds returns the ExternalUserIds field if non-nil, zero value otherwise.

### GetExternalUserIdsOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetExternalUserIdsOk() (*[]string, bool)`

GetExternalUserIdsOk returns a tuple with the ExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetExternalUserIds(v []string)`

SetExternalUserIds sets ExternalUserIds field to given value.

### HasExternalUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasExternalUserIds() bool`

HasExternalUserIds returns a boolean if a field has been set.

### GetFound

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetInputEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetInputEmail() string`

GetInputEmail returns the InputEmail field if non-nil, zero value otherwise.

### GetInputEmailOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetInputEmailOk() (*string, bool)`

GetInputEmailOk returns a tuple with the InputEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetInputEmail(v string)`

SetInputEmail sets InputEmail field to given value.

### HasInputEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasInputEmail() bool`

HasInputEmail returns a boolean if a field has been set.

### GetLatestChurn

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLatestChurn() UnibeeApiBeanUserChurnBulkCheckEpisode`

GetLatestChurn returns the LatestChurn field if non-nil, zero value otherwise.

### GetLatestChurnOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLatestChurnOk() (*UnibeeApiBeanUserChurnBulkCheckEpisode, bool)`

GetLatestChurnOk returns a tuple with the LatestChurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestChurn

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetLatestChurn(v UnibeeApiBeanUserChurnBulkCheckEpisode)`

SetLatestChurn sets LatestChurn field to given value.

### HasLatestChurn

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasLatestChurn() bool`

HasLatestChurn returns a boolean if a field has been set.

### GetLatestSubscription

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLatestSubscription() UnibeeApiBeanUserChurnBulkCheckSub`

GetLatestSubscription returns the LatestSubscription field if non-nil, zero value otherwise.

### GetLatestSubscriptionOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLatestSubscriptionOk() (*UnibeeApiBeanUserChurnBulkCheckSub, bool)`

GetLatestSubscriptionOk returns a tuple with the LatestSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSubscription

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetLatestSubscription(v UnibeeApiBeanUserChurnBulkCheckSub)`

SetLatestSubscription sets LatestSubscription field to given value.

### HasLatestSubscription

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasLatestSubscription() bool`

HasLatestSubscription returns a boolean if a field has been set.

### GetLookupEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLookupEmail() string`

GetLookupEmail returns the LookupEmail field if non-nil, zero value otherwise.

### GetLookupEmailOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetLookupEmailOk() (*string, bool)`

GetLookupEmailOk returns a tuple with the LookupEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetLookupEmail(v string)`

SetLookupEmail sets LookupEmail field to given value.

### HasLookupEmail

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasLookupEmail() bool`

HasLookupEmail returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOpenChurnStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetOpenChurnStatuses() []string`

GetOpenChurnStatuses returns the OpenChurnStatuses field if non-nil, zero value otherwise.

### GetOpenChurnStatusesOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetOpenChurnStatusesOk() (*[]string, bool)`

GetOpenChurnStatusesOk returns a tuple with the OpenChurnStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenChurnStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetOpenChurnStatuses(v []string)`

SetOpenChurnStatuses sets OpenChurnStatuses field to given value.

### HasOpenChurnStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasOpenChurnStatuses() bool`

HasOpenChurnStatuses returns a boolean if a field has been set.

### GetUserAccountSubscriptionStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetUserAccountSubscriptionStatuses() []string`

GetUserAccountSubscriptionStatuses returns the UserAccountSubscriptionStatuses field if non-nil, zero value otherwise.

### GetUserAccountSubscriptionStatusesOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetUserAccountSubscriptionStatusesOk() (*[]string, bool)`

GetUserAccountSubscriptionStatusesOk returns a tuple with the UserAccountSubscriptionStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountSubscriptionStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetUserAccountSubscriptionStatuses(v []string)`

SetUserAccountSubscriptionStatuses sets UserAccountSubscriptionStatuses field to given value.

### HasUserAccountSubscriptionStatuses

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasUserAccountSubscriptionStatuses() bool`

HasUserAccountSubscriptionStatuses returns a boolean if a field has been set.

### GetUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UnibeeApiBeanUserChurnBulkCheckItem) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


