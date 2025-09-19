# MerchantInvoiceCreditNoteListPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNotes** | Pointer to [**[]UnibeeApiBeanDetailCreditNoteDetail**](UnibeeApiBeanDetailCreditNoteDetail.md) | CreditNote Detail Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantInvoiceCreditNoteListPost200ResponseData

`func NewMerchantInvoiceCreditNoteListPost200ResponseData() *MerchantInvoiceCreditNoteListPost200ResponseData`

NewMerchantInvoiceCreditNoteListPost200ResponseData instantiates a new MerchantInvoiceCreditNoteListPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantInvoiceCreditNoteListPost200ResponseDataWithDefaults

`func NewMerchantInvoiceCreditNoteListPost200ResponseDataWithDefaults() *MerchantInvoiceCreditNoteListPost200ResponseData`

NewMerchantInvoiceCreditNoteListPost200ResponseDataWithDefaults instantiates a new MerchantInvoiceCreditNoteListPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNotes

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) GetCreditNotes() []UnibeeApiBeanDetailCreditNoteDetail`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) GetCreditNotesOk() (*[]UnibeeApiBeanDetailCreditNoteDetail, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) SetCreditNotes(v []UnibeeApiBeanDetailCreditNoteDetail)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantInvoiceCreditNoteListPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


