# UnibeeApiMerchantPaymentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | the item total amount,cent | [optional] 
**AmountExcludingTax** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Tax** | Pointer to **int64** |  | [optional] 
**TaxScale** | Pointer to **int64** | Tax Scale，1000 &#x3D; 10% | [optional] 
**UnitAmountExcludingTax** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentItem

`func NewUnibeeApiMerchantPaymentItem() *UnibeeApiMerchantPaymentItem`

NewUnibeeApiMerchantPaymentItem instantiates a new UnibeeApiMerchantPaymentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentItemWithDefaults

`func NewUnibeeApiMerchantPaymentItemWithDefaults() *UnibeeApiMerchantPaymentItem`

NewUnibeeApiMerchantPaymentItemWithDefaults instantiates a new UnibeeApiMerchantPaymentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantPaymentItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantPaymentItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantPaymentItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantPaymentItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) GetAmountExcludingTax() int64`

GetAmountExcludingTax returns the AmountExcludingTax field if non-nil, zero value otherwise.

### GetAmountExcludingTaxOk

`func (o *UnibeeApiMerchantPaymentItem) GetAmountExcludingTaxOk() (*int64, bool)`

GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) SetAmountExcludingTax(v int64)`

SetAmountExcludingTax sets AmountExcludingTax field to given value.

### HasAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) HasAmountExcludingTax() bool`

HasAmountExcludingTax returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPaymentItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPaymentItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantPaymentItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantPaymentItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantPaymentItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantPaymentItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantPaymentItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantPaymentItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantPaymentItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantPaymentItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTax

`func (o *UnibeeApiMerchantPaymentItem) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *UnibeeApiMerchantPaymentItem) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *UnibeeApiMerchantPaymentItem) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *UnibeeApiMerchantPaymentItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeApiMerchantPaymentItem) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeApiMerchantPaymentItem) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeApiMerchantPaymentItem) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeApiMerchantPaymentItem) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetUnitAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) GetUnitAmountExcludingTax() int64`

GetUnitAmountExcludingTax returns the UnitAmountExcludingTax field if non-nil, zero value otherwise.

### GetUnitAmountExcludingTaxOk

`func (o *UnibeeApiMerchantPaymentItem) GetUnitAmountExcludingTaxOk() (*int64, bool)`

GetUnitAmountExcludingTaxOk returns a tuple with the UnitAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) SetUnitAmountExcludingTax(v int64)`

SetUnitAmountExcludingTax sets UnitAmountExcludingTax field to given value.

### HasUnitAmountExcludingTax

`func (o *UnibeeApiMerchantPaymentItem) HasUnitAmountExcludingTax() bool`

HasUnitAmountExcludingTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


