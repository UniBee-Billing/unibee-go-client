# UnibeeApiSystemInvoiceCreatePaymentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **int64** | The gateway id (optional when not switching gateway) | [optional] 
**GatewayPaymentType** | Pointer to **string** | The gateway payment type (optional, required when gateway has multiple payment types) | [optional] 
**InvoiceId** | **string** | The invoice id | 
**St** | **string** | security token | 

## Methods

### NewUnibeeApiSystemInvoiceCreatePaymentReq

`func NewUnibeeApiSystemInvoiceCreatePaymentReq(invoiceId string, st string, ) *UnibeeApiSystemInvoiceCreatePaymentReq`

NewUnibeeApiSystemInvoiceCreatePaymentReq instantiates a new UnibeeApiSystemInvoiceCreatePaymentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceCreatePaymentReqWithDefaults

`func NewUnibeeApiSystemInvoiceCreatePaymentReqWithDefaults() *UnibeeApiSystemInvoiceCreatePaymentReq`

NewUnibeeApiSystemInvoiceCreatePaymentReqWithDefaults instantiates a new UnibeeApiSystemInvoiceCreatePaymentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSt

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetSt() string`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) GetStOk() (*string, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *UnibeeApiSystemInvoiceCreatePaymentReq) SetSt(v string)`

SetSt sets St field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


