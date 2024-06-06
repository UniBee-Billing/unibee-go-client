# UnibeeApiMerchantInvoicePdfUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The unique id of invoice | 
**IssueAddress** | Pointer to **string** | IssueAddress | [optional] 
**IssueCompanyName** | Pointer to **string** | IssueCompanyName | [optional] 
**IssueRegNumber** | Pointer to **string** | IssueRegNumber | [optional] 
**IssueVatNumber** | Pointer to **string** | IssueVatNumber | [optional] 
**SendUserEmail** | Pointer to **bool** | Whether sen invoice email to user or not，default false | [optional] [default to false]

## Methods

### NewUnibeeApiMerchantInvoicePdfUpdateReq

`func NewUnibeeApiMerchantInvoicePdfUpdateReq(invoiceId string, ) *UnibeeApiMerchantInvoicePdfUpdateReq`

NewUnibeeApiMerchantInvoicePdfUpdateReq instantiates a new UnibeeApiMerchantInvoicePdfUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoicePdfUpdateReqWithDefaults

`func NewUnibeeApiMerchantInvoicePdfUpdateReqWithDefaults() *UnibeeApiMerchantInvoicePdfUpdateReq`

NewUnibeeApiMerchantInvoicePdfUpdateReqWithDefaults instantiates a new UnibeeApiMerchantInvoicePdfUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetIssueAddress

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueAddress() string`

GetIssueAddress returns the IssueAddress field if non-nil, zero value otherwise.

### GetIssueAddressOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueAddressOk() (*string, bool)`

GetIssueAddressOk returns a tuple with the IssueAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAddress

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueAddress(v string)`

SetIssueAddress sets IssueAddress field to given value.

### HasIssueAddress

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueAddress() bool`

HasIssueAddress returns a boolean if a field has been set.

### GetIssueCompanyName

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueCompanyName() string`

GetIssueCompanyName returns the IssueCompanyName field if non-nil, zero value otherwise.

### GetIssueCompanyNameOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueCompanyNameOk() (*string, bool)`

GetIssueCompanyNameOk returns a tuple with the IssueCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCompanyName

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueCompanyName(v string)`

SetIssueCompanyName sets IssueCompanyName field to given value.

### HasIssueCompanyName

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueCompanyName() bool`

HasIssueCompanyName returns a boolean if a field has been set.

### GetIssueRegNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueRegNumber() string`

GetIssueRegNumber returns the IssueRegNumber field if non-nil, zero value otherwise.

### GetIssueRegNumberOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueRegNumberOk() (*string, bool)`

GetIssueRegNumberOk returns a tuple with the IssueRegNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRegNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueRegNumber(v string)`

SetIssueRegNumber sets IssueRegNumber field to given value.

### HasIssueRegNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueRegNumber() bool`

HasIssueRegNumber returns a boolean if a field has been set.

### GetIssueVatNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueVatNumber() string`

GetIssueVatNumber returns the IssueVatNumber field if non-nil, zero value otherwise.

### GetIssueVatNumberOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueVatNumberOk() (*string, bool)`

GetIssueVatNumberOk returns a tuple with the IssueVatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueVatNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueVatNumber(v string)`

SetIssueVatNumber sets IssueVatNumber field to given value.

### HasIssueVatNumber

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueVatNumber() bool`

HasIssueVatNumber returns a boolean if a field has been set.

### GetSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetSendUserEmail() bool`

GetSendUserEmail returns the SendUserEmail field if non-nil, zero value otherwise.

### GetSendUserEmailOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetSendUserEmailOk() (*bool, bool)`

GetSendUserEmailOk returns a tuple with the SendUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetSendUserEmail(v bool)`

SetSendUserEmail sets SendUserEmail field to given value.

### HasSendUserEmail

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasSendUserEmail() bool`

HasSendUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


