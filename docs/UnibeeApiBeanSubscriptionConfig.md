# UnibeeApiBeanSubscriptionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowngradeEffectImmediately** | Pointer to **bool** | DowngradeEffectImmediately, whether subscription update should effect immediately or at period end, default at period end | [optional] 
**GatewayVATRule** | Pointer to **string** |  | [optional] 
**IncompleteExpireTime** | Pointer to **int64** | IncompleteExpireTime, em.. default 1day for plan of month type | [optional] 
**InvoiceEmail** | Pointer to **bool** | InvoiceEmail, whether to send invoice email to user, default yes | [optional] 
**InvoicePdfGenerate** | Pointer to **bool** | InvoicePdfGenerate, whether to generate invoice pdf to user, default yes | [optional] 
**ShowZeroInvoice** | Pointer to **bool** | ShowZeroInvoice, show zero invoice or not, default no | [optional] 
**TryAutomaticPaymentBeforePeriodEnd** | Pointer to **int64** | TryAutomaticPaymentBeforePeriodEnd, default 30 min | [optional] 
**UpgradeProration** | Pointer to **bool** | UpgradeProration, whether subscription update generation proration invoice or not, default yes | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


