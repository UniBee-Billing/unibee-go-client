# UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Billing address | [optional] 
**City** | Pointer to **string** | Billing city | [optional] 
**CompanyName** | Pointer to **string** | Billing company name | [optional] 
**CountryCode** | Pointer to **string** | Billing country code | [optional] 
**InvoiceId** | **string** | The unique id of invoice | 
**Reason** | Pointer to **string** | Operation reason for audit log | [optional] 
**RegistrationNumber** | Pointer to **string** | Billing registration number | [optional] 
**SendEmail** | Pointer to **string** | Invoice recipient email | [optional] 
**VatNumber** | Pointer to **string** | Billing VAT number | [optional] 
**ZipCode** | Pointer to **string** | Billing zip code | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq

`func NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq(invoiceId string, ) *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq`

NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq instantiates a new UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReqWithDefaults

`func NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReqWithDefaults() *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq`

NewUnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetReason

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UnibeeApiMerchantInvoiceEditPaidInvoiceBillingDetailsReq) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


