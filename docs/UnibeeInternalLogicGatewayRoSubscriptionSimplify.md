# UnibeeInternalLogicGatewayRoSubscriptionSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonData** | Pointer to **string** | plan addon json data | [optional] 
**Amount** | Pointer to **int64** | amount, cent | [optional] 
**BillingCycleAnchor** | Pointer to **int64** | billing_cycle_anchor | [optional] 
**CancelAtPeriodEnd** | Pointer to **int32** | whether cancel at period end，0-false | 1-true | [optional] 
**CancelReason** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**CurrentPeriodEnd** | Pointer to **int64** | current_period_end, utc time | [optional] 
**CurrentPeriodStart** | Pointer to **int64** | current_period_start, utc time | [optional] 
**DunningTime** | Pointer to **int64** | dunning_time, utc time | [optional] 
**FirstPaidTime** | Pointer to **int64** | first success payment time | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayItemData** | Pointer to **string** | gateway_item_data | [optional] 
**GatewayStatus** | Pointer to **string** | gateway status，Stripe：https://stripe.com/docs/billing/subscriptions/webhooks  Paypal：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTime** | Pointer to **int64** |  | [optional] 
**LatestInvoiceId** | Pointer to **string** | latest_invoice_id | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PendingUpdateId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **int64** | plan id | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | status，0-Init | 1-Create｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TaskTime** | Pointer to **string** | task_time | [optional] 
**TaxScale** | Pointer to **int64** | Tax Scale，1000 &#x3D; 10% | [optional] 
**TestClock** | Pointer to **int64** | test_clock, simulator clock for subscription, if set , sub will out of cronjob controll | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**Type** | Pointer to **int32** | sub type, 0-gateway sub, 1-unibee sub | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoSubscriptionSimplify

`func NewUnibeeInternalLogicGatewayRoSubscriptionSimplify() *UnibeeInternalLogicGatewayRoSubscriptionSimplify`

NewUnibeeInternalLogicGatewayRoSubscriptionSimplify instantiates a new UnibeeInternalLogicGatewayRoSubscriptionSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoSubscriptionSimplifyWithDefaults

`func NewUnibeeInternalLogicGatewayRoSubscriptionSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoSubscriptionSimplify`

NewUnibeeInternalLogicGatewayRoSubscriptionSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoSubscriptionSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCancelAtPeriodEnd() int32`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCancelAtPeriodEndOk() (*int32, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCancelAtPeriodEnd(v int32)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCancelReason

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrentPeriodStart() int64`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetCurrentPeriodStartOk() (*int64, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetCurrentPeriodStart(v int64)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetDunningTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetDunningTime() int64`

GetDunningTime returns the DunningTime field if non-nil, zero value otherwise.

### GetDunningTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetDunningTimeOk() (*int64, bool)`

GetDunningTimeOk returns a tuple with the DunningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetDunningTime(v int64)`

SetDunningTime sets DunningTime field to given value.

### HasDunningTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasDunningTime() bool`

HasDunningTime returns a boolean if a field has been set.

### GetFirstPaidTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetFirstPaidTime() int64`

GetFirstPaidTime returns the FirstPaidTime field if non-nil, zero value otherwise.

### GetFirstPaidTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetFirstPaidTimeOk() (*int64, bool)`

GetFirstPaidTimeOk returns a tuple with the FirstPaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaidTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetFirstPaidTime(v int64)`

SetFirstPaidTime sets FirstPaidTime field to given value.

### HasFirstPaidTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasFirstPaidTime() bool`

HasFirstPaidTime returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayItemData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayItemData() string`

GetGatewayItemData returns the GatewayItemData field if non-nil, zero value otherwise.

### GetGatewayItemDataOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayItemDataOk() (*string, bool)`

GetGatewayItemDataOk returns a tuple with the GatewayItemData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayItemData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetGatewayItemData(v string)`

SetGatewayItemData sets GatewayItemData field to given value.

### HasGatewayItemData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasGatewayItemData() bool`

HasGatewayItemData returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLatestInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLatestInvoiceId() string`

GetLatestInvoiceId returns the LatestInvoiceId field if non-nil, zero value otherwise.

### GetLatestInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLatestInvoiceIdOk() (*string, bool)`

GetLatestInvoiceIdOk returns a tuple with the LatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetLatestInvoiceId(v string)`

SetLatestInvoiceId sets LatestInvoiceId field to given value.

### HasLatestInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasLatestInvoiceId() bool`

HasLatestInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPendingUpdateId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetPendingUpdateId() string`

GetPendingUpdateId returns the PendingUpdateId field if non-nil, zero value otherwise.

### GetPendingUpdateIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetPendingUpdateIdOk() (*string, bool)`

GetPendingUpdateIdOk returns a tuple with the PendingUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetPendingUpdateId(v string)`

SetPendingUpdateId sets PendingUpdateId field to given value.

### HasPendingUpdateId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasPendingUpdateId() bool`

HasPendingUpdateId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaskTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTaskTime() string`

GetTaskTime returns the TaskTime field if non-nil, zero value otherwise.

### GetTaskTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTaskTimeOk() (*string, bool)`

GetTaskTimeOk returns a tuple with the TaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetTaskTime(v string)`

SetTaskTime sets TaskTime field to given value.

### HasTaskTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasTaskTime() bool`

HasTaskTime returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTestClock

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTestClock() int64`

GetTestClock returns the TestClock field if non-nil, zero value otherwise.

### GetTestClockOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTestClockOk() (*int64, bool)`

GetTestClockOk returns a tuple with the TestClock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestClock

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetTestClock(v int64)`

SetTestClock sets TestClock field to given value.

### HasTestClock

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasTestClock() bool`

HasTestClock returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeInternalLogicGatewayRoSubscriptionSimplify) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


