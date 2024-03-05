# UnibeeInternalLogicGatewayRoInvoiceDetailSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceName** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]UnibeeInternalLogicGatewayRoInvoiceItemDetailRo**](UnibeeInternalLogicGatewayRoInvoiceItemDetailRo.md) |  | [optional] 
**PeriodEnd** | Pointer to **int64** |  | [optional] 
**PeriodStart** | Pointer to **int64** |  | [optional] 
**ProrationDate** | Pointer to **int64** |  | [optional] 
**ProrationScale** | Pointer to **int64** |  | [optional] 
**SubscriptionAmount** | Pointer to **int64** |  | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**TaxAmount** | Pointer to **int64** |  | [optional] 
**TaxScale** | Pointer to **int64** | Tax Scale，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplify

`func NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplify() *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify`

NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplify instantiates a new UnibeeInternalLogicGatewayRoInvoiceDetailSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplifyWithDefaults

`func NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify`

NewUnibeeInternalLogicGatewayRoInvoiceDetailSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoInvoiceDetailSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetLines() []UnibeeInternalLogicGatewayRoInvoiceItemDetailRo`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetLinesOk() (*[]UnibeeInternalLogicGatewayRoInvoiceItemDetailRo, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetLines(v []UnibeeInternalLogicGatewayRoInvoiceItemDetailRo)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetProrationDate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetProrationScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetProrationScale() int64`

GetProrationScale returns the ProrationScale field if non-nil, zero value otherwise.

### GetProrationScaleOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetProrationScaleOk() (*int64, bool)`

GetProrationScaleOk returns a tuple with the ProrationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetProrationScale(v int64)`

SetProrationScale sets ProrationScale field to given value.

### HasProrationScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasProrationScale() bool`

HasProrationScale returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailSimplify) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


