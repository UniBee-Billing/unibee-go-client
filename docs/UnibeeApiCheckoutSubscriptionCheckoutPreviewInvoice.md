# UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Invoice currency. | [optional] 
**DiscountAmount** | Pointer to **int64** | Actual discount deduction applied to this invoice, minor units. | [optional] 
**OriginAmount** | Pointer to **int64** | Original invoice amount before discounts, minor units. | [optional] 
**PromoCreditDiscountAmount** | Pointer to **int64** | Promo credit deduction applied to this invoice, minor units. | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** | Subscription amount excluding tax, minor units. | [optional] 
**TaxAmount** | Pointer to **int64** | Tax amount, minor units. | [optional] 
**TaxPercentage** | Pointer to **int64** | Tax percentage, 1000 &#x3D; 10%. | [optional] 
**TotalAmount** | Pointer to **int64** | Final payable amount including tax and after discounts, minor units. | [optional] 
**UsVATCalculateGatewayName** | Pointer to **string** | Display name of the US VAT calculation gateway, when applicable. | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice

`func NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice() *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice`

NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice instantiates a new UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoiceWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoiceWithDefaults() *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice`

NewUnibeeApiCheckoutSubscriptionCheckoutPreviewInvoiceWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetPromoCreditDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetPromoCreditDiscountAmount() int64`

GetPromoCreditDiscountAmount returns the PromoCreditDiscountAmount field if non-nil, zero value otherwise.

### GetPromoCreditDiscountAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetPromoCreditDiscountAmountOk() (*int64, bool)`

GetPromoCreditDiscountAmountOk returns a tuple with the PromoCreditDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetPromoCreditDiscountAmount(v int64)`

SetPromoCreditDiscountAmount sets PromoCreditDiscountAmount field to given value.

### HasPromoCreditDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasPromoCreditDiscountAmount() bool`

HasPromoCreditDiscountAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUsVATCalculateGatewayName

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetUsVATCalculateGatewayName() string`

GetUsVATCalculateGatewayName returns the UsVATCalculateGatewayName field if non-nil, zero value otherwise.

### GetUsVATCalculateGatewayNameOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) GetUsVATCalculateGatewayNameOk() (*string, bool)`

GetUsVATCalculateGatewayNameOk returns a tuple with the UsVATCalculateGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVATCalculateGatewayName

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) SetUsVATCalculateGatewayName(v string)`

SetUsVATCalculateGatewayName sets UsVATCalculateGatewayName field to given value.

### HasUsVATCalculateGatewayName

`func (o *UnibeeApiCheckoutSubscriptionCheckoutPreviewInvoice) HasUsVATCalculateGatewayName() bool`

HasUsVATCalculateGatewayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


