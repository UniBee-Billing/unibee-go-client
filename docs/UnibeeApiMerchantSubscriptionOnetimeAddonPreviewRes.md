# UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**TaxAmount** | Pointer to **int64** |  | [optional] 
**TaxPercentage** | Pointer to **int64** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes() *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewResWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewResWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetAddon() UnibeeApiBeanPlan`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetAddonOk() (*UnibeeApiBeanPlan, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetAddon(v UnibeeApiBeanPlan)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewRes) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


