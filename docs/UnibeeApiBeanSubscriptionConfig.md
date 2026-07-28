# UnibeeApiBeanSubscriptionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoChargeMaxRetryCount** | Pointer to **int64** | AutoChargeMaxRetryCount, max automatic-charge attempts for one invoice including the first automatic attempt; 0 &#x3D; unlimited (current behavior) | [optional] 
**AutoChargeRetryInterval** | Pointer to **int64** | AutoChargeRetryInterval, seconds to wait before retrying a failed auto-charge within the window, default 12h (43200) | [optional] 
**DiscountNewUserIgnoreFreeZeroInvoice** | Pointer to **bool** | DiscountNewUserIgnoreFreeZeroInvoice, when true ignore only true free-plan zero invoices (no discount code / discount amount / promo credit) for new-user discount scope; default false (legacy: any paid subscription invoice counts as not new) | [optional] 
**DowngradeEffectImmediately** | Pointer to **bool** | DowngradeEffectImmediately, whether subscription update should effect immediately or at period end, default at period end | [optional] 
**DowngradeNonImmediatelyEffectBeforePeriodEnd** | Pointer to **int64** | DowngradeNonImmediatelyEffectBeforePeriodEnd, for non-immediate downgrade, can take effect before period end (in seconds), default 30 minutes (1800) | [optional] 
**FiatExchangeApiKey** | Pointer to **string** |  | [optional] 
**GatewayVATRule** | Pointer to **string** |  | [optional] 
**IncompleteExpireTime** | Pointer to **int64** | IncompleteExpireTime, em.. default 1day for plan of month type | [optional] 
**InvoiceEmail** | Pointer to **bool** | InvoiceEmail, whether to send invoice email to user, default yes | [optional] 
**InvoicePdfGenerate** | Pointer to **bool** | InvoicePdfGenerate, whether to generate invoice pdf to user, default yes | [optional] 
**ShowZeroInvoice** | Pointer to **bool** | ShowZeroInvoice, show zero invoice or not, default no | [optional] 
**TryAutomaticPaymentBeforePeriodEnd** | Pointer to **int64** | TryAutomaticPaymentBeforePeriodEnd, default 30 min | [optional] 
**UpgradeProration** | Pointer to **bool** | UpgradeProration, whether subscription update generation proration invoice or not, default yes | [optional] 
**UpgradeResetBillingCycle** | Pointer to **bool** | UpgradeResetBillingCycle, for same-interval upgrades only: when true (default), credit unused old plan, charge full new plan cycle from upgrade time and reset billing anchor; when false, prorate new plan to current period end (legacy fixed-end). No effect when billing interval changes. | [optional] 

## Methods

### NewUnibeeApiBeanSubscriptionConfig

`func NewUnibeeApiBeanSubscriptionConfig() *UnibeeApiBeanSubscriptionConfig`

NewUnibeeApiBeanSubscriptionConfig instantiates a new UnibeeApiBeanSubscriptionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionConfigWithDefaults

`func NewUnibeeApiBeanSubscriptionConfigWithDefaults() *UnibeeApiBeanSubscriptionConfig`

NewUnibeeApiBeanSubscriptionConfigWithDefaults instantiates a new UnibeeApiBeanSubscriptionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanSubscriptionConfig) GetAutoChargeMaxRetryCount() int64`

GetAutoChargeMaxRetryCount returns the AutoChargeMaxRetryCount field if non-nil, zero value otherwise.

### GetAutoChargeMaxRetryCountOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetAutoChargeMaxRetryCountOk() (*int64, bool)`

GetAutoChargeMaxRetryCountOk returns a tuple with the AutoChargeMaxRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanSubscriptionConfig) SetAutoChargeMaxRetryCount(v int64)`

SetAutoChargeMaxRetryCount sets AutoChargeMaxRetryCount field to given value.

### HasAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanSubscriptionConfig) HasAutoChargeMaxRetryCount() bool`

HasAutoChargeMaxRetryCount returns a boolean if a field has been set.

### GetAutoChargeRetryInterval

`func (o *UnibeeApiBeanSubscriptionConfig) GetAutoChargeRetryInterval() int64`

GetAutoChargeRetryInterval returns the AutoChargeRetryInterval field if non-nil, zero value otherwise.

### GetAutoChargeRetryIntervalOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetAutoChargeRetryIntervalOk() (*int64, bool)`

GetAutoChargeRetryIntervalOk returns a tuple with the AutoChargeRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeRetryInterval

`func (o *UnibeeApiBeanSubscriptionConfig) SetAutoChargeRetryInterval(v int64)`

SetAutoChargeRetryInterval sets AutoChargeRetryInterval field to given value.

### HasAutoChargeRetryInterval

`func (o *UnibeeApiBeanSubscriptionConfig) HasAutoChargeRetryInterval() bool`

HasAutoChargeRetryInterval returns a boolean if a field has been set.

### GetDiscountNewUserIgnoreFreeZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) GetDiscountNewUserIgnoreFreeZeroInvoice() bool`

GetDiscountNewUserIgnoreFreeZeroInvoice returns the DiscountNewUserIgnoreFreeZeroInvoice field if non-nil, zero value otherwise.

### GetDiscountNewUserIgnoreFreeZeroInvoiceOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetDiscountNewUserIgnoreFreeZeroInvoiceOk() (*bool, bool)`

GetDiscountNewUserIgnoreFreeZeroInvoiceOk returns a tuple with the DiscountNewUserIgnoreFreeZeroInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountNewUserIgnoreFreeZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) SetDiscountNewUserIgnoreFreeZeroInvoice(v bool)`

SetDiscountNewUserIgnoreFreeZeroInvoice sets DiscountNewUserIgnoreFreeZeroInvoice field to given value.

### HasDiscountNewUserIgnoreFreeZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) HasDiscountNewUserIgnoreFreeZeroInvoice() bool`

HasDiscountNewUserIgnoreFreeZeroInvoice returns a boolean if a field has been set.

### GetDowngradeEffectImmediately

`func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeEffectImmediately() bool`

GetDowngradeEffectImmediately returns the DowngradeEffectImmediately field if non-nil, zero value otherwise.

### GetDowngradeEffectImmediatelyOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeEffectImmediatelyOk() (*bool, bool)`

GetDowngradeEffectImmediatelyOk returns a tuple with the DowngradeEffectImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradeEffectImmediately

`func (o *UnibeeApiBeanSubscriptionConfig) SetDowngradeEffectImmediately(v bool)`

SetDowngradeEffectImmediately sets DowngradeEffectImmediately field to given value.

### HasDowngradeEffectImmediately

`func (o *UnibeeApiBeanSubscriptionConfig) HasDowngradeEffectImmediately() bool`

HasDowngradeEffectImmediately returns a boolean if a field has been set.

### GetDowngradeNonImmediatelyEffectBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeNonImmediatelyEffectBeforePeriodEnd() int64`

GetDowngradeNonImmediatelyEffectBeforePeriodEnd returns the DowngradeNonImmediatelyEffectBeforePeriodEnd field if non-nil, zero value otherwise.

### GetDowngradeNonImmediatelyEffectBeforePeriodEndOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeNonImmediatelyEffectBeforePeriodEndOk() (*int64, bool)`

GetDowngradeNonImmediatelyEffectBeforePeriodEndOk returns a tuple with the DowngradeNonImmediatelyEffectBeforePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradeNonImmediatelyEffectBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) SetDowngradeNonImmediatelyEffectBeforePeriodEnd(v int64)`

SetDowngradeNonImmediatelyEffectBeforePeriodEnd sets DowngradeNonImmediatelyEffectBeforePeriodEnd field to given value.

### HasDowngradeNonImmediatelyEffectBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) HasDowngradeNonImmediatelyEffectBeforePeriodEnd() bool`

HasDowngradeNonImmediatelyEffectBeforePeriodEnd returns a boolean if a field has been set.

### GetFiatExchangeApiKey

`func (o *UnibeeApiBeanSubscriptionConfig) GetFiatExchangeApiKey() string`

GetFiatExchangeApiKey returns the FiatExchangeApiKey field if non-nil, zero value otherwise.

### GetFiatExchangeApiKeyOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetFiatExchangeApiKeyOk() (*string, bool)`

GetFiatExchangeApiKeyOk returns a tuple with the FiatExchangeApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatExchangeApiKey

`func (o *UnibeeApiBeanSubscriptionConfig) SetFiatExchangeApiKey(v string)`

SetFiatExchangeApiKey sets FiatExchangeApiKey field to given value.

### HasFiatExchangeApiKey

`func (o *UnibeeApiBeanSubscriptionConfig) HasFiatExchangeApiKey() bool`

HasFiatExchangeApiKey returns a boolean if a field has been set.

### GetGatewayVATRule

`func (o *UnibeeApiBeanSubscriptionConfig) GetGatewayVATRule() string`

GetGatewayVATRule returns the GatewayVATRule field if non-nil, zero value otherwise.

### GetGatewayVATRuleOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetGatewayVATRuleOk() (*string, bool)`

GetGatewayVATRuleOk returns a tuple with the GatewayVATRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVATRule

`func (o *UnibeeApiBeanSubscriptionConfig) SetGatewayVATRule(v string)`

SetGatewayVATRule sets GatewayVATRule field to given value.

### HasGatewayVATRule

`func (o *UnibeeApiBeanSubscriptionConfig) HasGatewayVATRule() bool`

HasGatewayVATRule returns a boolean if a field has been set.

### GetIncompleteExpireTime

`func (o *UnibeeApiBeanSubscriptionConfig) GetIncompleteExpireTime() int64`

GetIncompleteExpireTime returns the IncompleteExpireTime field if non-nil, zero value otherwise.

### GetIncompleteExpireTimeOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetIncompleteExpireTimeOk() (*int64, bool)`

GetIncompleteExpireTimeOk returns a tuple with the IncompleteExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteExpireTime

`func (o *UnibeeApiBeanSubscriptionConfig) SetIncompleteExpireTime(v int64)`

SetIncompleteExpireTime sets IncompleteExpireTime field to given value.

### HasIncompleteExpireTime

`func (o *UnibeeApiBeanSubscriptionConfig) HasIncompleteExpireTime() bool`

HasIncompleteExpireTime returns a boolean if a field has been set.

### GetInvoiceEmail

`func (o *UnibeeApiBeanSubscriptionConfig) GetInvoiceEmail() bool`

GetInvoiceEmail returns the InvoiceEmail field if non-nil, zero value otherwise.

### GetInvoiceEmailOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetInvoiceEmailOk() (*bool, bool)`

GetInvoiceEmailOk returns a tuple with the InvoiceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceEmail

`func (o *UnibeeApiBeanSubscriptionConfig) SetInvoiceEmail(v bool)`

SetInvoiceEmail sets InvoiceEmail field to given value.

### HasInvoiceEmail

`func (o *UnibeeApiBeanSubscriptionConfig) HasInvoiceEmail() bool`

HasInvoiceEmail returns a boolean if a field has been set.

### GetInvoicePdfGenerate

`func (o *UnibeeApiBeanSubscriptionConfig) GetInvoicePdfGenerate() bool`

GetInvoicePdfGenerate returns the InvoicePdfGenerate field if non-nil, zero value otherwise.

### GetInvoicePdfGenerateOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetInvoicePdfGenerateOk() (*bool, bool)`

GetInvoicePdfGenerateOk returns a tuple with the InvoicePdfGenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfGenerate

`func (o *UnibeeApiBeanSubscriptionConfig) SetInvoicePdfGenerate(v bool)`

SetInvoicePdfGenerate sets InvoicePdfGenerate field to given value.

### HasInvoicePdfGenerate

`func (o *UnibeeApiBeanSubscriptionConfig) HasInvoicePdfGenerate() bool`

HasInvoicePdfGenerate returns a boolean if a field has been set.

### GetShowZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) GetShowZeroInvoice() bool`

GetShowZeroInvoice returns the ShowZeroInvoice field if non-nil, zero value otherwise.

### GetShowZeroInvoiceOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetShowZeroInvoiceOk() (*bool, bool)`

GetShowZeroInvoiceOk returns a tuple with the ShowZeroInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) SetShowZeroInvoice(v bool)`

SetShowZeroInvoice sets ShowZeroInvoice field to given value.

### HasShowZeroInvoice

`func (o *UnibeeApiBeanSubscriptionConfig) HasShowZeroInvoice() bool`

HasShowZeroInvoice returns a boolean if a field has been set.

### GetTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) GetTryAutomaticPaymentBeforePeriodEnd() int64`

GetTryAutomaticPaymentBeforePeriodEnd returns the TryAutomaticPaymentBeforePeriodEnd field if non-nil, zero value otherwise.

### GetTryAutomaticPaymentBeforePeriodEndOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetTryAutomaticPaymentBeforePeriodEndOk() (*int64, bool)`

GetTryAutomaticPaymentBeforePeriodEndOk returns a tuple with the TryAutomaticPaymentBeforePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) SetTryAutomaticPaymentBeforePeriodEnd(v int64)`

SetTryAutomaticPaymentBeforePeriodEnd sets TryAutomaticPaymentBeforePeriodEnd field to given value.

### HasTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiBeanSubscriptionConfig) HasTryAutomaticPaymentBeforePeriodEnd() bool`

HasTryAutomaticPaymentBeforePeriodEnd returns a boolean if a field has been set.

### GetUpgradeProration

`func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeProration() bool`

GetUpgradeProration returns the UpgradeProration field if non-nil, zero value otherwise.

### GetUpgradeProrationOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeProrationOk() (*bool, bool)`

GetUpgradeProrationOk returns a tuple with the UpgradeProration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeProration

`func (o *UnibeeApiBeanSubscriptionConfig) SetUpgradeProration(v bool)`

SetUpgradeProration sets UpgradeProration field to given value.

### HasUpgradeProration

`func (o *UnibeeApiBeanSubscriptionConfig) HasUpgradeProration() bool`

HasUpgradeProration returns a boolean if a field has been set.

### GetUpgradeResetBillingCycle

`func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeResetBillingCycle() bool`

GetUpgradeResetBillingCycle returns the UpgradeResetBillingCycle field if non-nil, zero value otherwise.

### GetUpgradeResetBillingCycleOk

`func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeResetBillingCycleOk() (*bool, bool)`

GetUpgradeResetBillingCycleOk returns a tuple with the UpgradeResetBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeResetBillingCycle

`func (o *UnibeeApiBeanSubscriptionConfig) SetUpgradeResetBillingCycle(v bool)`

SetUpgradeResetBillingCycle sets UpgradeResetBillingCycle field to given value.

### HasUpgradeResetBillingCycle

`func (o *UnibeeApiBeanSubscriptionConfig) HasUpgradeResetBillingCycle() bool`

HasUpgradeResetBillingCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


