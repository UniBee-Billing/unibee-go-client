# UnibeeApiMerchantInvoiceSummaryListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Page size, maximum 100 | [optional] [default to 20]
**IncludeMetadata** | Pointer to **bool** | Whether to include invoice metadata | [optional] [default to false]
**InvoiceId** | Pointer to **string** | Filter by exact invoice ID | [optional] 
**Page** | Pointer to **int32** | Page number, starting from 0 | [optional] [default to 0]
**SortField** | Pointer to **string** | Sort field: gmt_create or id | [optional] [default to "gmt_create"]
**SortType** | Pointer to **string** | Sort direction: asc or desc | [optional] [default to "desc"]
**Status** | Pointer to **[]int32** | Invoice status: 0-init | 1-pending | 2-processing | 3-paid | 4-failed | 5-cancelled | [optional] 
**SubscriptionId** | Pointer to **string** | Filter by exact subscription ID | [optional] 
**UserId** | Pointer to **int64** | Filter by user ID | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceSummaryListReq

`func NewUnibeeApiMerchantInvoiceSummaryListReq() *UnibeeApiMerchantInvoiceSummaryListReq`

NewUnibeeApiMerchantInvoiceSummaryListReq instantiates a new UnibeeApiMerchantInvoiceSummaryListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceSummaryListReqWithDefaults

`func NewUnibeeApiMerchantInvoiceSummaryListReqWithDefaults() *UnibeeApiMerchantInvoiceSummaryListReq`

NewUnibeeApiMerchantInvoiceSummaryListReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceSummaryListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIncludeMetadata

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetIncludeMetadata() bool`

GetIncludeMetadata returns the IncludeMetadata field if non-nil, zero value otherwise.

### GetIncludeMetadataOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetIncludeMetadataOk() (*bool, bool)`

GetIncludeMetadataOk returns a tuple with the IncludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadata

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetIncludeMetadata(v bool)`

SetIncludeMetadata sets IncludeMetadata field to given value.

### HasIncludeMetadata

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasIncludeMetadata() bool`

HasIncludeMetadata returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantInvoiceSummaryListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


