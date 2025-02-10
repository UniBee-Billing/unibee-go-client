# UnibeeApiMerchantSubscriptionUpdatePreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**DiscountMessage** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**NextPeriodInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**ProrationDate** | Pointer to **int64** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionUpdatePreviewRes

`func NewUnibeeApiMerchantSubscriptionUpdatePreviewRes() *UnibeeApiMerchantSubscriptionUpdatePreviewRes`

NewUnibeeApiMerchantSubscriptionUpdatePreviewRes instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionUpdatePreviewResWithDefaults

`func NewUnibeeApiMerchantSubscriptionUpdatePreviewResWithDefaults() *UnibeeApiMerchantSubscriptionUpdatePreviewRes`

NewUnibeeApiMerchantSubscriptionUpdatePreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountMessage() string`

GetDiscountMessage returns the DiscountMessage field if non-nil, zero value otherwise.

### GetDiscountMessageOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetDiscountMessageOk() (*string, bool)`

GetDiscountMessageOk returns a tuple with the DiscountMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetDiscountMessage(v string)`

SetDiscountMessage sets DiscountMessage field to given value.

### HasDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasDiscountMessage() bool`

HasDiscountMessage returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetNextPeriodInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoice() UnibeeApiBeanInvoice`

GetNextPeriodInvoice returns the NextPeriodInvoice field if non-nil, zero value otherwise.

### GetNextPeriodInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetNextPeriodInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetNextPeriodInvoiceOk returns a tuple with the NextPeriodInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPeriodInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetNextPeriodInvoice(v UnibeeApiBeanInvoice)`

SetNextPeriodInvoice sets NextPeriodInvoice field to given value.

### HasNextPeriodInvoice

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasNextPeriodInvoice() bool`

HasNextPeriodInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


