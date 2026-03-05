# MerchantSubscriptionRenewPreviewPost200ResponseData

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

### NewMerchantSubscriptionRenewPreviewPost200ResponseData

`func NewMerchantSubscriptionRenewPreviewPost200ResponseData() *MerchantSubscriptionRenewPreviewPost200ResponseData`

NewMerchantSubscriptionRenewPreviewPost200ResponseData instantiates a new MerchantSubscriptionRenewPreviewPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionRenewPreviewPost200ResponseDataWithDefaults

`func NewMerchantSubscriptionRenewPreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionRenewPreviewPost200ResponseData`

NewMerchantSubscriptionRenewPreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionRenewPreviewPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCredit

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetInvoice

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetSubscription

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTaxAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MerchantSubscriptionRenewPreviewPost200ResponseData) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


