# UnibeeApiMerchantInvoiceDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNotes** | Pointer to [**[]UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) | CreditNotes Object List Link To Invoice | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceDetailRes

`func NewUnibeeApiMerchantInvoiceDetailRes() *UnibeeApiMerchantInvoiceDetailRes`

NewUnibeeApiMerchantInvoiceDetailRes instantiates a new UnibeeApiMerchantInvoiceDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceDetailResWithDefaults

`func NewUnibeeApiMerchantInvoiceDetailResWithDefaults() *UnibeeApiMerchantInvoiceDetailRes`

NewUnibeeApiMerchantInvoiceDetailResWithDefaults instantiates a new UnibeeApiMerchantInvoiceDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNotes

`func (o *UnibeeApiMerchantInvoiceDetailRes) GetCreditNotes() []UnibeeApiBeanDetailInvoiceDetail`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *UnibeeApiMerchantInvoiceDetailRes) GetCreditNotesOk() (*[]UnibeeApiBeanDetailInvoiceDetail, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *UnibeeApiMerchantInvoiceDetailRes) SetCreditNotes(v []UnibeeApiBeanDetailInvoiceDetail)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *UnibeeApiMerchantInvoiceDetailRes) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantInvoiceDetailRes) GetInvoice() UnibeeApiBeanDetailInvoiceDetail`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantInvoiceDetailRes) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantInvoiceDetailRes) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantInvoiceDetailRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


