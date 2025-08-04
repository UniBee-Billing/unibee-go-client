# UnibeeApiMerchantInvoiceListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountEnd** | Pointer to **int32** | The filter end amount of invoice | [optional] 
**AmountStart** | Pointer to **int32** | The filter start amount of invoice | [optional] 
**Count** | Pointer to **int32** | Count By Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart | [optional] 
**Currency** | Pointer to **string** | The currency of invoice | [optional] 
**DeleteInclude** | Pointer to **bool** | Deleted Involved，Need Admin Permission | [optional] 
**FirstName** | Pointer to **string** | The firstName of invoice | [optional] 
**GatewayIds** | Pointer to **[]int64** | GatewayIds, Search Filter GatewayIds | [optional] 
**LastName** | Pointer to **string** | The lastName of invoice | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**ReportTimeEnd** | Pointer to **int64** | ReportTimeEnd | [optional] 
**ReportTimeStart** | Pointer to **int64** | ReportTimeStart | [optional] 
**SendEmail** | Pointer to **string** | The filter email of invoice | [optional] 
**SortField** | Pointer to **string** | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**Type** | Pointer to **int32** | invoice Type, 0-payment, 1-refund | [optional] 
**UserId** | Pointer to **int64** | The filter userid of invoice | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceListReq

`func NewUnibeeApiMerchantInvoiceListReq() *UnibeeApiMerchantInvoiceListReq`

NewUnibeeApiMerchantInvoiceListReq instantiates a new UnibeeApiMerchantInvoiceListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceListReqWithDefaults

`func NewUnibeeApiMerchantInvoiceListReqWithDefaults() *UnibeeApiMerchantInvoiceListReq`

NewUnibeeApiMerchantInvoiceListReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountEnd

`func (o *UnibeeApiMerchantInvoiceListReq) GetAmountEnd() int32`

GetAmountEnd returns the AmountEnd field if non-nil, zero value otherwise.

### GetAmountEndOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetAmountEndOk() (*int32, bool)`

GetAmountEndOk returns a tuple with the AmountEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountEnd

`func (o *UnibeeApiMerchantInvoiceListReq) SetAmountEnd(v int32)`

SetAmountEnd sets AmountEnd field to given value.

### HasAmountEnd

`func (o *UnibeeApiMerchantInvoiceListReq) HasAmountEnd() bool`

HasAmountEnd returns a boolean if a field has been set.

### GetAmountStart

`func (o *UnibeeApiMerchantInvoiceListReq) GetAmountStart() int32`

GetAmountStart returns the AmountStart field if non-nil, zero value otherwise.

### GetAmountStartOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetAmountStartOk() (*int32, bool)`

GetAmountStartOk returns a tuple with the AmountStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountStart

`func (o *UnibeeApiMerchantInvoiceListReq) SetAmountStart(v int32)`

SetAmountStart sets AmountStart field to given value.

### HasAmountStart

`func (o *UnibeeApiMerchantInvoiceListReq) HasAmountStart() bool`

HasAmountStart returns a boolean if a field has been set.

### GetCount

`func (o *UnibeeApiMerchantInvoiceListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantInvoiceListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantInvoiceListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantInvoiceListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeleteInclude

`func (o *UnibeeApiMerchantInvoiceListReq) GetDeleteInclude() bool`

GetDeleteInclude returns the DeleteInclude field if non-nil, zero value otherwise.

### GetDeleteIncludeOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetDeleteIncludeOk() (*bool, bool)`

GetDeleteIncludeOk returns a tuple with the DeleteInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInclude

`func (o *UnibeeApiMerchantInvoiceListReq) SetDeleteInclude(v bool)`

SetDeleteInclude sets DeleteInclude field to given value.

### HasDeleteInclude

`func (o *UnibeeApiMerchantInvoiceListReq) HasDeleteInclude() bool`

HasDeleteInclude returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiMerchantInvoiceListReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantInvoiceListReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantInvoiceListReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayIds

`func (o *UnibeeApiMerchantInvoiceListReq) GetGatewayIds() []int64`

GetGatewayIds returns the GatewayIds field if non-nil, zero value otherwise.

### GetGatewayIdsOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetGatewayIdsOk() (*[]int64, bool)`

GetGatewayIdsOk returns a tuple with the GatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIds

`func (o *UnibeeApiMerchantInvoiceListReq) SetGatewayIds(v []int64)`

SetGatewayIds sets GatewayIds field to given value.

### HasGatewayIds

`func (o *UnibeeApiMerchantInvoiceListReq) HasGatewayIds() bool`

HasGatewayIds returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantInvoiceListReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantInvoiceListReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantInvoiceListReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantInvoiceListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantInvoiceListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantInvoiceListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetReportTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeEnd() int64`

GetReportTimeEnd returns the ReportTimeEnd field if non-nil, zero value otherwise.

### GetReportTimeEndOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeEndOk() (*int64, bool)`

GetReportTimeEndOk returns a tuple with the ReportTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) SetReportTimeEnd(v int64)`

SetReportTimeEnd sets ReportTimeEnd field to given value.

### HasReportTimeEnd

`func (o *UnibeeApiMerchantInvoiceListReq) HasReportTimeEnd() bool`

HasReportTimeEnd returns a boolean if a field has been set.

### GetReportTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeStart() int64`

GetReportTimeStart returns the ReportTimeStart field if non-nil, zero value otherwise.

### GetReportTimeStartOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeStartOk() (*int64, bool)`

GetReportTimeStartOk returns a tuple with the ReportTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) SetReportTimeStart(v int64)`

SetReportTimeStart sets ReportTimeStart field to given value.

### HasReportTimeStart

`func (o *UnibeeApiMerchantInvoiceListReq) HasReportTimeStart() bool`

HasReportTimeStart returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeApiMerchantInvoiceListReq) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeApiMerchantInvoiceListReq) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeApiMerchantInvoiceListReq) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantInvoiceListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantInvoiceListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantInvoiceListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantInvoiceListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantInvoiceListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantInvoiceListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantInvoiceListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantInvoiceListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantInvoiceListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantInvoiceListReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantInvoiceListReq) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantInvoiceListReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantInvoiceListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantInvoiceListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantInvoiceListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantInvoiceListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


