# UnibeeApiMerchantInvoiceCreditNoteListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNotes** | Pointer to [**[]UnibeeApiBeanDetailCreditNoteDetail**](UnibeeApiBeanDetailCreditNoteDetail.md) | CreditNote Detail Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceCreditNoteListRes

`func NewUnibeeApiMerchantInvoiceCreditNoteListRes() *UnibeeApiMerchantInvoiceCreditNoteListRes`

NewUnibeeApiMerchantInvoiceCreditNoteListRes instantiates a new UnibeeApiMerchantInvoiceCreditNoteListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceCreditNoteListResWithDefaults

`func NewUnibeeApiMerchantInvoiceCreditNoteListResWithDefaults() *UnibeeApiMerchantInvoiceCreditNoteListRes`

NewUnibeeApiMerchantInvoiceCreditNoteListResWithDefaults instantiates a new UnibeeApiMerchantInvoiceCreditNoteListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNotes

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) GetCreditNotes() []UnibeeApiBeanDetailCreditNoteDetail`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) GetCreditNotesOk() (*[]UnibeeApiBeanDetailCreditNoteDetail, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) SetCreditNotes(v []UnibeeApiBeanDetailCreditNoteDetail)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantInvoiceCreditNoteListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


