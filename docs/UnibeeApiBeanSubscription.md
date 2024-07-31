# UnibeeApiBeanSubscription

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
**DefaultPaymentMethodId** | Pointer to **string** |  | [optional] 
**DunningTime** | Pointer to **int64** | dunning_time, utc time | [optional] 
**ExternalSubscriptionId** | Pointer to **string** | external_subscription_id | [optional] 
**Features** | Pointer to **string** | features | [optional] 
**FirstPaidTime** | Pointer to **int64** | first success payment time | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayStatus** | Pointer to **string** | gateway status，Stripe：https://stripe.com/docs/billing/subscriptions/webhooks  Paypal：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTime** | Pointer to **int64** |  | [optional] 
**LatestInvoiceId** | Pointer to **string** | latest_invoice_id | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PendingUpdateId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **int64** | plan id | [optional] 
**ProductId** | Pointer to **int64** | product id | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | status，1-Pending｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9- Failed | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TaskTime** | Pointer to **string** | task_time | [optional] 
**TaxPercentage** | Pointer to **int64** | TaxPercentage，1000 &#x3D; 10% | [optional] 
**TestClock** | Pointer to **int64** | test_clock, simulator clock for subscription, if set , sub will out of cronjob controll | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**Type** | Pointer to **int32** | sub type, 0-gateway sub, 1-unibee sub | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiBeanSubscription

`func NewUnibeeApiBeanSubscription() *UnibeeApiBeanSubscription`

NewUnibeeApiBeanSubscription instantiates a new UnibeeApiBeanSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionWithDefaults

`func NewUnibeeApiBeanSubscriptionWithDefaults() *UnibeeApiBeanSubscription`

NewUnibeeApiBeanSubscriptionWithDefaults instantiates a new UnibeeApiBeanSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeApiBeanSubscription) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeApiBeanSubscription) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeApiBeanSubscription) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeApiBeanSubscription) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeApiBeanSubscription) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanSubscription) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanSubscription) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanSubscription) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *UnibeeApiBeanSubscription) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeApiBeanSubscription) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeApiBeanSubscription) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeApiBeanSubscription) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *UnibeeApiBeanSubscription) GetCancelAtPeriodEnd() int32`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *UnibeeApiBeanSubscription) GetCancelAtPeriodEndOk() (*int32, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *UnibeeApiBeanSubscription) SetCancelAtPeriodEnd(v int32)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *UnibeeApiBeanSubscription) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCancelReason

`func (o *UnibeeApiBeanSubscription) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *UnibeeApiBeanSubscription) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *UnibeeApiBeanSubscription) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *UnibeeApiBeanSubscription) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanSubscription) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanSubscription) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanSubscription) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanSubscription) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanSubscription) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanSubscription) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanSubscription) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanSubscription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanSubscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanSubscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanSubscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanSubscription) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeApiBeanSubscription) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeApiBeanSubscription) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeApiBeanSubscription) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeApiBeanSubscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeApiBeanSubscription) GetCurrentPeriodStart() int64`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeApiBeanSubscription) GetCurrentPeriodStartOk() (*int64, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeApiBeanSubscription) SetCurrentPeriodStart(v int64)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeApiBeanSubscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetDefaultPaymentMethodId

`func (o *UnibeeApiBeanSubscription) GetDefaultPaymentMethodId() string`

GetDefaultPaymentMethodId returns the DefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodIdOk

`func (o *UnibeeApiBeanSubscription) GetDefaultPaymentMethodIdOk() (*string, bool)`

GetDefaultPaymentMethodIdOk returns a tuple with the DefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethodId

`func (o *UnibeeApiBeanSubscription) SetDefaultPaymentMethodId(v string)`

SetDefaultPaymentMethodId sets DefaultPaymentMethodId field to given value.

### HasDefaultPaymentMethodId

`func (o *UnibeeApiBeanSubscription) HasDefaultPaymentMethodId() bool`

HasDefaultPaymentMethodId returns a boolean if a field has been set.

### GetDunningTime

`func (o *UnibeeApiBeanSubscription) GetDunningTime() int64`

GetDunningTime returns the DunningTime field if non-nil, zero value otherwise.

### GetDunningTimeOk

`func (o *UnibeeApiBeanSubscription) GetDunningTimeOk() (*int64, bool)`

GetDunningTimeOk returns a tuple with the DunningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningTime

`func (o *UnibeeApiBeanSubscription) SetDunningTime(v int64)`

SetDunningTime sets DunningTime field to given value.

### HasDunningTime

`func (o *UnibeeApiBeanSubscription) HasDunningTime() bool`

HasDunningTime returns a boolean if a field has been set.

### GetExternalSubscriptionId

`func (o *UnibeeApiBeanSubscription) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *UnibeeApiBeanSubscription) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *UnibeeApiBeanSubscription) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *UnibeeApiBeanSubscription) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### GetFeatures

`func (o *UnibeeApiBeanSubscription) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UnibeeApiBeanSubscription) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UnibeeApiBeanSubscription) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UnibeeApiBeanSubscription) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFirstPaidTime

`func (o *UnibeeApiBeanSubscription) GetFirstPaidTime() int64`

GetFirstPaidTime returns the FirstPaidTime field if non-nil, zero value otherwise.

### GetFirstPaidTimeOk

`func (o *UnibeeApiBeanSubscription) GetFirstPaidTimeOk() (*int64, bool)`

GetFirstPaidTimeOk returns a tuple with the FirstPaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaidTime

`func (o *UnibeeApiBeanSubscription) SetFirstPaidTime(v int64)`

SetFirstPaidTime sets FirstPaidTime field to given value.

### HasFirstPaidTime

`func (o *UnibeeApiBeanSubscription) HasFirstPaidTime() bool`

HasFirstPaidTime returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiBeanSubscription) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiBeanSubscription) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiBeanSubscription) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiBeanSubscription) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanSubscription) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanSubscription) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanSubscription) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanSubscription) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeApiBeanSubscription) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeApiBeanSubscription) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeApiBeanSubscription) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeApiBeanSubscription) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanSubscription) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanSubscription) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanSubscription) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeApiBeanSubscription) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeApiBeanSubscription) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeApiBeanSubscription) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeApiBeanSubscription) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLatestInvoiceId

`func (o *UnibeeApiBeanSubscription) GetLatestInvoiceId() string`

GetLatestInvoiceId returns the LatestInvoiceId field if non-nil, zero value otherwise.

### GetLatestInvoiceIdOk

`func (o *UnibeeApiBeanSubscription) GetLatestInvoiceIdOk() (*string, bool)`

GetLatestInvoiceIdOk returns a tuple with the LatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoiceId

`func (o *UnibeeApiBeanSubscription) SetLatestInvoiceId(v string)`

SetLatestInvoiceId sets LatestInvoiceId field to given value.

### HasLatestInvoiceId

`func (o *UnibeeApiBeanSubscription) HasLatestInvoiceId() bool`

HasLatestInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanSubscription) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanSubscription) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanSubscription) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanSubscription) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanSubscription) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanSubscription) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanSubscription) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanSubscription) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanSubscription) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanSubscription) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanSubscription) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanSubscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPendingUpdateId

`func (o *UnibeeApiBeanSubscription) GetPendingUpdateId() string`

GetPendingUpdateId returns the PendingUpdateId field if non-nil, zero value otherwise.

### GetPendingUpdateIdOk

`func (o *UnibeeApiBeanSubscription) GetPendingUpdateIdOk() (*string, bool)`

GetPendingUpdateIdOk returns a tuple with the PendingUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateId

`func (o *UnibeeApiBeanSubscription) SetPendingUpdateId(v string)`

SetPendingUpdateId sets PendingUpdateId field to given value.

### HasPendingUpdateId

`func (o *UnibeeApiBeanSubscription) HasPendingUpdateId() bool`

HasPendingUpdateId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanSubscription) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanSubscription) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanSubscription) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiBeanSubscription) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiBeanSubscription) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiBeanSubscription) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiBeanSubscription) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanSubscription) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanSubscription) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanSubscription) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanSubscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiBeanSubscription) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiBeanSubscription) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiBeanSubscription) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiBeanSubscription) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanSubscription) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanSubscription) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanSubscription) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaskTime

`func (o *UnibeeApiBeanSubscription) GetTaskTime() string`

GetTaskTime returns the TaskTime field if non-nil, zero value otherwise.

### GetTaskTimeOk

`func (o *UnibeeApiBeanSubscription) GetTaskTimeOk() (*string, bool)`

GetTaskTimeOk returns a tuple with the TaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTime

`func (o *UnibeeApiBeanSubscription) SetTaskTime(v string)`

SetTaskTime sets TaskTime field to given value.

### HasTaskTime

`func (o *UnibeeApiBeanSubscription) HasTaskTime() bool`

HasTaskTime returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanSubscription) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanSubscription) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanSubscription) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanSubscription) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTestClock

`func (o *UnibeeApiBeanSubscription) GetTestClock() int64`

GetTestClock returns the TestClock field if non-nil, zero value otherwise.

### GetTestClockOk

`func (o *UnibeeApiBeanSubscription) GetTestClockOk() (*int64, bool)`

GetTestClockOk returns a tuple with the TestClock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestClock

`func (o *UnibeeApiBeanSubscription) SetTestClock(v int64)`

SetTestClock sets TestClock field to given value.

### HasTestClock

`func (o *UnibeeApiBeanSubscription) HasTestClock() bool`

HasTestClock returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiBeanSubscription) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiBeanSubscription) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiBeanSubscription) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiBeanSubscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanSubscription) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanSubscription) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanSubscription) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanSubscription) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanSubscription) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanSubscription) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanSubscription) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanSubscription) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanSubscription) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanSubscription) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


