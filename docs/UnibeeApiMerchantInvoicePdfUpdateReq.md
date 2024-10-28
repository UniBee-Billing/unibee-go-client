# UnibeeApiMerchantInvoicePdfUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The unique id of invoice | 
**IssueAddress** | Pointer to **string** | IssueAddress | [optional] 
**IssueCompanyName** | Pointer to **string** | IssueCompanyName | [optional] 
**IssueRegNumber** | Pointer to **string** | IssueRegNumber | [optional] 
**IssueVatNumber** | Pointer to **string** | IssueVatNumber | [optional] 
**LocalizedCurrency** | Pointer to **string** | LocalizedCurrency, To display localized currency amount | [optional] 
**LocalizedExchangeRate** | Pointer to **float32** | LocalizedExchangeRate, exchange rate must set while LocalizedCurrency enabled | [optional] 
**LocalizedExchangeRateDescription** | Pointer to **float32** | LocalizedExchangeRateDescription | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**SendUserEmail** | Pointer to **bool** | Whether sen invoice email to user or not，default false | [optional] [default to false]
**ShowDetailItem** | Pointer to **bool** | ShowDetailItem, whether to display detail item information in pdf generate, unitAmount, quantity, etc. | [optional] [default to false]
**Template** | Pointer to **string** | Template | [optional] 

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

### GetLocalizedCurrency

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedCurrency() string`

GetLocalizedCurrency returns the LocalizedCurrency field if non-nil, zero value otherwise.

### GetLocalizedCurrencyOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedCurrencyOk() (*string, bool)`

GetLocalizedCurrencyOk returns a tuple with the LocalizedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedCurrency

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetLocalizedCurrency(v string)`

SetLocalizedCurrency sets LocalizedCurrency field to given value.

### HasLocalizedCurrency

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasLocalizedCurrency() bool`

HasLocalizedCurrency returns a boolean if a field has been set.

### GetLocalizedExchangeRate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRate() float32`

GetLocalizedExchangeRate returns the LocalizedExchangeRate field if non-nil, zero value otherwise.

### GetLocalizedExchangeRateOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRateOk() (*float32, bool)`

GetLocalizedExchangeRateOk returns a tuple with the LocalizedExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedExchangeRate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetLocalizedExchangeRate(v float32)`

SetLocalizedExchangeRate sets LocalizedExchangeRate field to given value.

### HasLocalizedExchangeRate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasLocalizedExchangeRate() bool`

HasLocalizedExchangeRate returns a boolean if a field has been set.

### GetLocalizedExchangeRateDescription

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRateDescription() float32`

GetLocalizedExchangeRateDescription returns the LocalizedExchangeRateDescription field if non-nil, zero value otherwise.

### GetLocalizedExchangeRateDescriptionOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRateDescriptionOk() (*float32, bool)`

GetLocalizedExchangeRateDescriptionOk returns a tuple with the LocalizedExchangeRateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedExchangeRateDescription

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetLocalizedExchangeRateDescription(v float32)`

SetLocalizedExchangeRateDescription sets LocalizedExchangeRateDescription field to given value.

### HasLocalizedExchangeRateDescription

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasLocalizedExchangeRateDescription() bool`

HasLocalizedExchangeRateDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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

### GetShowDetailItem

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetShowDetailItem() bool`

GetShowDetailItem returns the ShowDetailItem field if non-nil, zero value otherwise.

### GetShowDetailItemOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetShowDetailItemOk() (*bool, bool)`

GetShowDetailItemOk returns a tuple with the ShowDetailItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDetailItem

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetShowDetailItem(v bool)`

SetShowDetailItem sets ShowDetailItem field to given value.

### HasShowDetailItem

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasShowDetailItem() bool`

HasShowDetailItem returns a boolean if a field has been set.

### GetTemplate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


