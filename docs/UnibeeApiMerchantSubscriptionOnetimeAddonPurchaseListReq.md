# UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | Filter AddonIds, Default All | [optional] 
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**PeriodEnd** | Pointer to **int64** | PeriodEnd，addon effective period end，UTC timestamp，seconds | [optional] 
**PeriodStart** | Pointer to **int64** | PeriodStart，addon effective period start，UTC timestamp，seconds | [optional] 
**Status** | Pointer to **[]int32** | Filter Status, Default All，1-Create｜2-Paid｜3-Cancel｜4-Expired | [optional] 
**SubscriptionIds** | Pointer to **[]string** | Filter SubscriptionIds, Default All | [optional] 
**UserId** | **int64** | UserId | 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq(userId int64, ) *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPurchaseListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


