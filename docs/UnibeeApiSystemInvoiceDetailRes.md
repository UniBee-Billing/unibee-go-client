# UnibeeApiSystemInvoiceDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** |  | [optional] 
**CreditNotes** | Pointer to [**[]UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) | CreditNotes Object List Link To Invoice | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) |  | [optional] 
**Language** | Pointer to **string** | User language preference for display | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceDetailRes

`func NewUnibeeApiSystemInvoiceDetailRes() *UnibeeApiSystemInvoiceDetailRes`

NewUnibeeApiSystemInvoiceDetailRes instantiates a new UnibeeApiSystemInvoiceDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceDetailResWithDefaults

`func NewUnibeeApiSystemInvoiceDetailResWithDefaults() *UnibeeApiSystemInvoiceDetailRes`

NewUnibeeApiSystemInvoiceDetailResWithDefaults instantiates a new UnibeeApiSystemInvoiceDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiSystemInvoiceDetailRes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCreditNotes

`func (o *UnibeeApiSystemInvoiceDetailRes) GetCreditNotes() []UnibeeApiBeanDetailInvoiceDetail`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *UnibeeApiSystemInvoiceDetailRes) GetCreditNotesOk() (*[]UnibeeApiBeanDetailInvoiceDetail, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *UnibeeApiSystemInvoiceDetailRes) SetCreditNotes(v []UnibeeApiBeanDetailInvoiceDetail)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *UnibeeApiSystemInvoiceDetailRes) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiSystemInvoiceDetailRes) GetInvoice() UnibeeApiBeanDetailInvoiceDetail`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiSystemInvoiceDetailRes) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiSystemInvoiceDetailRes) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiSystemInvoiceDetailRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetLanguage

`func (o *UnibeeApiSystemInvoiceDetailRes) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UnibeeApiSystemInvoiceDetailRes) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UnibeeApiSystemInvoiceDetailRes) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UnibeeApiSystemInvoiceDetailRes) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiSystemInvoiceDetailRes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiSystemInvoiceDetailRes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


