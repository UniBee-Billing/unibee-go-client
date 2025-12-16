# UnibeeApiSystemInvoiceCreateNextSplitPaymentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount for this split payment (in cents) | 
**GatewayId** | Pointer to **int64** | The gateway id for this split payment (optional, defaults to invoice gateway) | [optional] 
**GatewayPaymentType** | Pointer to **string** | The gateway payment type (optional, required if gateway has multiple payment types) | [optional] 
**InvoiceId** | **string** | The invoice id for split payment | 
**St** | **string** | security token | 

## Methods

### NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReq

`func NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReq(amount int64, invoiceId string, st string, ) *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq`

NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReq instantiates a new UnibeeApiSystemInvoiceCreateNextSplitPaymentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReqWithDefaults

`func NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReqWithDefaults() *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq`

NewUnibeeApiSystemInvoiceCreateNextSplitPaymentReqWithDefaults instantiates a new UnibeeApiSystemInvoiceCreateNextSplitPaymentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetGatewayId

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSt

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetSt() string`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) GetStOk() (*string, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *UnibeeApiSystemInvoiceCreateNextSplitPaymentReq) SetSt(v string)`

SetSt sets St field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


