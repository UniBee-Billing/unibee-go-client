# UnibeeApiMerchantInvoicePdfGenerateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice ID | 
**SendUserEmail** | Pointer to **bool** | Whether Send Invoice Email To User Or Not，Default Not Send | [optional] [default to false]

## Methods

### NewUnibeeApiMerchantInvoicePdfGenerateReq

`func NewUnibeeApiMerchantInvoicePdfGenerateReq(invoiceId string, ) *UnibeeApiMerchantInvoicePdfGenerateReq`

NewUnibeeApiMerchantInvoicePdfGenerateReq instantiates a new UnibeeApiMerchantInvoicePdfGenerateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoicePdfGenerateReqWithDefaults

`func NewUnibeeApiMerchantInvoicePdfGenerateReqWithDefaults() *UnibeeApiMerchantInvoicePdfGenerateReq`

NewUnibeeApiMerchantInvoicePdfGenerateReqWithDefaults instantiates a new UnibeeApiMerchantInvoicePdfGenerateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetSendUserEmail() bool`

GetSendUserEmail returns the SendUserEmail field if non-nil, zero value otherwise.

### GetSendUserEmailOk

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) GetSendUserEmailOk() (*bool, bool)`

GetSendUserEmailOk returns a tuple with the SendUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) SetSendUserEmail(v bool)`

SetSendUserEmail sets SendUserEmail field to given value.

### HasSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfGenerateReq) HasSendUserEmail() bool`

HasSendUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


