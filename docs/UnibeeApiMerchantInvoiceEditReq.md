# UnibeeApiMerchantInvoiceEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of invoice | [optional] 
**Finish** | Pointer to **bool** |  | [optional] 
**GatewayId** | Pointer to **int64** | The gateway id of invoice | [optional] 
**InvoiceId** | **string** | The unique id of invoice | 
**Lines** | Pointer to [**[]UnibeeApiMerchantInvoiceNewInvoiceItemParam**](UnibeeApiMerchantInvoiceNewInvoiceItemParam.md) |  | [optional] 
**Name** | Pointer to **string** | The name of invoice | [optional] 
**TaxPercentage** | Pointer to **int64** | The tax percentage of invoice，1000&#x3D;10% | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceEditReq

`func NewUnibeeApiMerchantInvoiceEditReq(invoiceId string, ) *UnibeeApiMerchantInvoiceEditReq`

NewUnibeeApiMerchantInvoiceEditReq instantiates a new UnibeeApiMerchantInvoiceEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceEditReqWithDefaults

`func NewUnibeeApiMerchantInvoiceEditReqWithDefaults() *UnibeeApiMerchantInvoiceEditReq`

NewUnibeeApiMerchantInvoiceEditReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceEditReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceEditReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantInvoiceEditReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFinish

`func (o *UnibeeApiMerchantInvoiceEditReq) GetFinish() bool`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetFinishOk() (*bool, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *UnibeeApiMerchantInvoiceEditReq) SetFinish(v bool)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *UnibeeApiMerchantInvoiceEditReq) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceEditReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceEditReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantInvoiceEditReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceEditReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceEditReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetLines

`func (o *UnibeeApiMerchantInvoiceEditReq) GetLines() []UnibeeApiMerchantInvoiceNewInvoiceItemParam`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetLinesOk() (*[]UnibeeApiMerchantInvoiceNewInvoiceItemParam, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiMerchantInvoiceEditReq) SetLines(v []UnibeeApiMerchantInvoiceNewInvoiceItemParam)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiMerchantInvoiceEditReq) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantInvoiceEditReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantInvoiceEditReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantInvoiceEditReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantInvoiceEditReq) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantInvoiceEditReq) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantInvoiceEditReq) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantInvoiceEditReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


