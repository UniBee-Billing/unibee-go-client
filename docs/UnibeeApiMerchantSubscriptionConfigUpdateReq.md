# UnibeeApiMerchantSubscriptionConfigUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowngradeEffectImmediately** | Pointer to **bool** | DowngradeEffectImmediately, Immediate Downgrade (by default, the downgrades takes effect at the end of the period ） | [optional] 
**GatewayVATRule** | Pointer to [**[]UnibeeApiBeanMerchantVatRule**](UnibeeApiBeanMerchantVatRule.md) |  | [optional] 
**IncompleteExpireTime** | Pointer to **int32** | IncompleteExpireTime, seconds, Incomplete Status Duration(The period during which subscription remains in “incomplete”) | [optional] 
**InvoiceEmail** | Pointer to **bool** | InvoiceEmail, Enable Invoice Email (Toggle to send invoice email to customers) | [optional] 
**ShowZeroInvoice** | Pointer to **bool** | ShowZeroInvoice, Display Invoices With Zero Amount (Invoice With Zero Amount will hidden in list by default) | [optional] 
**TryAutomaticPaymentBeforePeriodEnd** | Pointer to **int32** | TryAutomaticPaymentBeforePeriodEnd, Auto-charge Start Before Period End （Time Difference for Auto-Payment Activation Before Period End） | [optional] 
**UpgradeProration** | Pointer to **bool** | UpgradeProration, Prorated Upgrade Invoices(Upgrades will generate prorated invoice by default) | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionConfigUpdateReq

`func NewUnibeeApiMerchantSubscriptionConfigUpdateReq() *UnibeeApiMerchantSubscriptionConfigUpdateReq`

NewUnibeeApiMerchantSubscriptionConfigUpdateReq instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionConfigUpdateReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionConfigUpdateReqWithDefaults() *UnibeeApiMerchantSubscriptionConfigUpdateReq`

NewUnibeeApiMerchantSubscriptionConfigUpdateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDowngradeEffectImmediately

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetDowngradeEffectImmediately() bool`

GetDowngradeEffectImmediately returns the DowngradeEffectImmediately field if non-nil, zero value otherwise.

### GetDowngradeEffectImmediatelyOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetDowngradeEffectImmediatelyOk() (*bool, bool)`

GetDowngradeEffectImmediatelyOk returns a tuple with the DowngradeEffectImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradeEffectImmediately

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetDowngradeEffectImmediately(v bool)`

SetDowngradeEffectImmediately sets DowngradeEffectImmediately field to given value.

### HasDowngradeEffectImmediately

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasDowngradeEffectImmediately() bool`

HasDowngradeEffectImmediately returns a boolean if a field has been set.

### GetGatewayVATRule

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetGatewayVATRule() []UnibeeApiBeanMerchantVatRule`

GetGatewayVATRule returns the GatewayVATRule field if non-nil, zero value otherwise.

### GetGatewayVATRuleOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetGatewayVATRuleOk() (*[]UnibeeApiBeanMerchantVatRule, bool)`

GetGatewayVATRuleOk returns a tuple with the GatewayVATRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVATRule

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetGatewayVATRule(v []UnibeeApiBeanMerchantVatRule)`

SetGatewayVATRule sets GatewayVATRule field to given value.

### HasGatewayVATRule

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasGatewayVATRule() bool`

HasGatewayVATRule returns a boolean if a field has been set.

### GetIncompleteExpireTime

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetIncompleteExpireTime() int32`

GetIncompleteExpireTime returns the IncompleteExpireTime field if non-nil, zero value otherwise.

### GetIncompleteExpireTimeOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetIncompleteExpireTimeOk() (*int32, bool)`

GetIncompleteExpireTimeOk returns a tuple with the IncompleteExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteExpireTime

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetIncompleteExpireTime(v int32)`

SetIncompleteExpireTime sets IncompleteExpireTime field to given value.

### HasIncompleteExpireTime

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasIncompleteExpireTime() bool`

HasIncompleteExpireTime returns a boolean if a field has been set.

### GetInvoiceEmail

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetInvoiceEmail() bool`

GetInvoiceEmail returns the InvoiceEmail field if non-nil, zero value otherwise.

### GetInvoiceEmailOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetInvoiceEmailOk() (*bool, bool)`

GetInvoiceEmailOk returns a tuple with the InvoiceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceEmail

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetInvoiceEmail(v bool)`

SetInvoiceEmail sets InvoiceEmail field to given value.

### HasInvoiceEmail

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasInvoiceEmail() bool`

HasInvoiceEmail returns a boolean if a field has been set.

### GetShowZeroInvoice

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetShowZeroInvoice() bool`

GetShowZeroInvoice returns the ShowZeroInvoice field if non-nil, zero value otherwise.

### GetShowZeroInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetShowZeroInvoiceOk() (*bool, bool)`

GetShowZeroInvoiceOk returns a tuple with the ShowZeroInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowZeroInvoice

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetShowZeroInvoice(v bool)`

SetShowZeroInvoice sets ShowZeroInvoice field to given value.

### HasShowZeroInvoice

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasShowZeroInvoice() bool`

HasShowZeroInvoice returns a boolean if a field has been set.

### GetTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetTryAutomaticPaymentBeforePeriodEnd() int32`

GetTryAutomaticPaymentBeforePeriodEnd returns the TryAutomaticPaymentBeforePeriodEnd field if non-nil, zero value otherwise.

### GetTryAutomaticPaymentBeforePeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetTryAutomaticPaymentBeforePeriodEndOk() (*int32, bool)`

GetTryAutomaticPaymentBeforePeriodEndOk returns a tuple with the TryAutomaticPaymentBeforePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetTryAutomaticPaymentBeforePeriodEnd(v int32)`

SetTryAutomaticPaymentBeforePeriodEnd sets TryAutomaticPaymentBeforePeriodEnd field to given value.

### HasTryAutomaticPaymentBeforePeriodEnd

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasTryAutomaticPaymentBeforePeriodEnd() bool`

HasTryAutomaticPaymentBeforePeriodEnd returns a boolean if a field has been set.

### GetUpgradeProration

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetUpgradeProration() bool`

GetUpgradeProration returns the UpgradeProration field if non-nil, zero value otherwise.

### GetUpgradeProrationOk

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetUpgradeProrationOk() (*bool, bool)`

GetUpgradeProrationOk returns a tuple with the UpgradeProration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeProration

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetUpgradeProration(v bool)`

SetUpgradeProration sets UpgradeProration field to given value.

### HasUpgradeProration

`func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasUpgradeProration() bool`

HasUpgradeProration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


