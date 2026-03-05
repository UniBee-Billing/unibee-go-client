# UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Required when lines are provided, line amount in cents (including tax). | [optional] 
**Description** | Pointer to **string** | Optional, line item description. | [optional] 
**DiscountAmount** | Pointer to **int64** | Optional, discount amount in cents for this line. | [optional] 
**Name** | Pointer to **string** | Optional, line item name. If empty, system will fall back to productName or invoiceName. | [optional] 
**Quantity** | Pointer to **int64** | Optional, quantity for this line, default 1 when 0 or negative. | [optional] 
**TaxAmount** | Pointer to **int64** | Optional, tax amount in cents for this line. If omitted, system may derive distribution from invoice-level taxAmount. | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem

`func NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem() *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem`

NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem instantiates a new UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItemWithDefaults

`func NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItemWithDefaults() *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem`

NewUnibeeApiMerchantInvoiceHistoryInvoiceImportLineItemWithDefaults instantiates a new UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


