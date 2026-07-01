# UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPaidAmount** | Pointer to **int64** | Required. Historical paid amount in minor units (cent), used directly as the upgrade proration credit base, must be greater than 0 | [optional] 
**SkipImportCheck** | Pointer to **bool** | Optional, default false. When true, skips the import-state check (subscription Data must start with ImportBy), allowing the proration base to be set up even on non-imported subscriptions or ones that already entered the next period. The first-invoice zero-amount/paid/no-discount checks still apply. | [optional] 
**StartTime** | Pointer to **int64** | Optional. Historical period start time (utc seconds). When greater than 0, overrides the proration window start; must be earlier than the subscription current period end | [optional] 
**SubscriptionId** | Pointer to **string** | Required. The id of subscription to backfill the proration base for | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReq

`func NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReq() *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq`

NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReq instantiates a new UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReqWithDefaults() *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq`

NewUnibeeApiMerchantSubscriptionImportProrationBaseSetupReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPaidAmount

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetLastPaidAmount() int64`

GetLastPaidAmount returns the LastPaidAmount field if non-nil, zero value otherwise.

### GetLastPaidAmountOk

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetLastPaidAmountOk() (*int64, bool)`

GetLastPaidAmountOk returns a tuple with the LastPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaidAmount

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) SetLastPaidAmount(v int64)`

SetLastPaidAmount sets LastPaidAmount field to given value.

### HasLastPaidAmount

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) HasLastPaidAmount() bool`

HasLastPaidAmount returns a boolean if a field has been set.

### GetSkipImportCheck

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetSkipImportCheck() bool`

GetSkipImportCheck returns the SkipImportCheck field if non-nil, zero value otherwise.

### GetSkipImportCheckOk

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetSkipImportCheckOk() (*bool, bool)`

GetSkipImportCheckOk returns a tuple with the SkipImportCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipImportCheck

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) SetSkipImportCheck(v bool)`

SetSkipImportCheck sets SkipImportCheck field to given value.

### HasSkipImportCheck

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) HasSkipImportCheck() bool`

HasSkipImportCheck returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionImportProrationBaseSetupReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


