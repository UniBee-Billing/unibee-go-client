# MerchantInvoiceDetailGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNotes** | Pointer to [**[]UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) | CreditNotes Object List Link To Invoice | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) |  | [optional] 

## Methods

### NewMerchantInvoiceDetailGet200ResponseData

`func NewMerchantInvoiceDetailGet200ResponseData() *MerchantInvoiceDetailGet200ResponseData`

NewMerchantInvoiceDetailGet200ResponseData instantiates a new MerchantInvoiceDetailGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantInvoiceDetailGet200ResponseDataWithDefaults

`func NewMerchantInvoiceDetailGet200ResponseDataWithDefaults() *MerchantInvoiceDetailGet200ResponseData`

NewMerchantInvoiceDetailGet200ResponseDataWithDefaults instantiates a new MerchantInvoiceDetailGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNotes

`func (o *MerchantInvoiceDetailGet200ResponseData) GetCreditNotes() []UnibeeApiBeanDetailInvoiceDetail`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *MerchantInvoiceDetailGet200ResponseData) GetCreditNotesOk() (*[]UnibeeApiBeanDetailInvoiceDetail, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *MerchantInvoiceDetailGet200ResponseData) SetCreditNotes(v []UnibeeApiBeanDetailInvoiceDetail)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *MerchantInvoiceDetailGet200ResponseData) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetInvoice

`func (o *MerchantInvoiceDetailGet200ResponseData) GetInvoice() UnibeeApiBeanDetailInvoiceDetail`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MerchantInvoiceDetailGet200ResponseData) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MerchantInvoiceDetailGet200ResponseData) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MerchantInvoiceDetailGet200ResponseData) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


