# UnibeeApiBeanInvoiceItemSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**AmountExcludingTax** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**PeriodEnd** | Pointer to **int64** |  | [optional] 
**PeriodStart** | Pointer to **int64** |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Proration** | Pointer to **bool** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Tax** | Pointer to **int64** |  | [optional] 
**TaxPercentage** | Pointer to **int64** | Tax Percentage，1000 &#x3D; 10% | [optional] 
**UnitAmountExcludingTax** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiBeanInvoiceItemSimplify

`func NewUnibeeApiBeanInvoiceItemSimplify() *UnibeeApiBeanInvoiceItemSimplify`

NewUnibeeApiBeanInvoiceItemSimplify instantiates a new UnibeeApiBeanInvoiceItemSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanInvoiceItemSimplifyWithDefaults

`func NewUnibeeApiBeanInvoiceItemSimplifyWithDefaults() *UnibeeApiBeanInvoiceItemSimplify`

NewUnibeeApiBeanInvoiceItemSimplifyWithDefaults instantiates a new UnibeeApiBeanInvoiceItemSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountExcludingTax() int64`

GetAmountExcludingTax returns the AmountExcludingTax field if non-nil, zero value otherwise.

### GetAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetAmountExcludingTaxOk() (*int64, bool)`

GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetAmountExcludingTax(v int64)`

SetAmountExcludingTax sets AmountExcludingTax field to given value.

### HasAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasAmountExcludingTax() bool`

HasAmountExcludingTax returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProration

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetProration() bool`

GetProration returns the Proration field if non-nil, zero value otherwise.

### GetProrationOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetProrationOk() (*bool, bool)`

GetProrationOk returns a tuple with the Proration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProration

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetProration(v bool)`

SetProration sets Proration field to given value.

### HasProration

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasProration() bool`

HasProration returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUnitAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetUnitAmountExcludingTax() int64`

GetUnitAmountExcludingTax returns the UnitAmountExcludingTax field if non-nil, zero value otherwise.

### GetUnitAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceItemSimplify) GetUnitAmountExcludingTaxOk() (*int64, bool)`

GetUnitAmountExcludingTaxOk returns a tuple with the UnitAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) SetUnitAmountExcludingTax(v int64)`

SetUnitAmountExcludingTax sets UnitAmountExcludingTax field to given value.

### HasUnitAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceItemSimplify) HasUnitAmountExcludingTax() bool`

HasUnitAmountExcludingTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


