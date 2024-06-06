# UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The unique id of invoice | 
**Reason** | Pointer to **string** | The reason of mark action | [optional] 
**TransferNumber** | **string** | The transfer number of invoice | 

## Methods

### NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq

`func NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq(invoiceId string, transferNumber string, ) *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq`

NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReq instantiates a new UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReqWithDefaults

`func NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReqWithDefaults() *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq`

NewUnibeeApiMerchantInvoiceMarkWireTransferSuccessReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetReason

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTransferNumber

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetTransferNumber() string`

GetTransferNumber returns the TransferNumber field if non-nil, zero value otherwise.

### GetTransferNumberOk

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) GetTransferNumberOk() (*string, bool)`

GetTransferNumberOk returns a tuple with the TransferNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferNumber

`func (o *UnibeeApiMerchantInvoiceMarkWireTransferSuccessReq) SetTransferNumber(v string)`

SetTransferNumber sets TransferNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


