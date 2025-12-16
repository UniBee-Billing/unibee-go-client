# UnibeeApiMerchantVatUSVATUploadHistoryListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**GatewayName** | Pointer to **string** | GatewayName, em. TaxJar | [optional] 
**InvoiceId** | Pointer to **string** | Invoice Id | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**UploadType** | Pointer to **[]string** | UploadType, PaymentUpload or RefundUpload, default both | [optional] 

## Methods

### NewUnibeeApiMerchantVatUSVATUploadHistoryListReq

`func NewUnibeeApiMerchantVatUSVATUploadHistoryListReq() *UnibeeApiMerchantVatUSVATUploadHistoryListReq`

NewUnibeeApiMerchantVatUSVATUploadHistoryListReq instantiates a new UnibeeApiMerchantVatUSVATUploadHistoryListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantVatUSVATUploadHistoryListReqWithDefaults

`func NewUnibeeApiMerchantVatUSVATUploadHistoryListReqWithDefaults() *UnibeeApiMerchantVatUSVATUploadHistoryListReq`

NewUnibeeApiMerchantVatUSVATUploadHistoryListReqWithDefaults instantiates a new UnibeeApiMerchantVatUSVATUploadHistoryListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetUploadType

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetUploadType() []string`

GetUploadType returns the UploadType field if non-nil, zero value otherwise.

### GetUploadTypeOk

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) GetUploadTypeOk() (*[]string, bool)`

GetUploadTypeOk returns a tuple with the UploadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadType

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) SetUploadType(v []string)`

SetUploadType sets UploadType field to given value.

### HasUploadType

`func (o *UnibeeApiMerchantVatUSVATUploadHistoryListReq) HasUploadType() bool`

HasUploadType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


