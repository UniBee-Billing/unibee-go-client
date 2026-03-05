# UnibeeApiMerchantSubscriptionRenewPreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCredit** | Pointer to **bool** | Whether promo credit is effectively applied in this renew preview. | [optional] 
**Currency** | Pointer to **string** | Currency used for the renew preview. | [optional] 
**DiscountAmount** | Pointer to **int64** | Total discount amount applied to the renew invoice, including promo credit if applicable. | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** | Original invoice amount before discounts and promo credits for the renew, in minor units. | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**TaxAmount** | Pointer to **int64** | Total tax amount applied to the renew invoice, in minor units. | [optional] 
**TotalAmount** | Pointer to **int64** | Final payable amount for the renew invoice, including tax and after discounts, in minor units. | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionRenewPreviewRes

`func NewUnibeeApiMerchantSubscriptionRenewPreviewRes() *UnibeeApiMerchantSubscriptionRenewPreviewRes`

NewUnibeeApiMerchantSubscriptionRenewPreviewRes instantiates a new UnibeeApiMerchantSubscriptionRenewPreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionRenewPreviewResWithDefaults

`func NewUnibeeApiMerchantSubscriptionRenewPreviewResWithDefaults() *UnibeeApiMerchantSubscriptionRenewPreviewRes`

NewUnibeeApiMerchantSubscriptionRenewPreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewPreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


