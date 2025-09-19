# UnibeeApiMerchantInvoiceCreditNoteListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count By Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Currency** | Pointer to **string** | The currency of invoice | [optional] 
**Emails** | Pointer to **string** | The email list of invoice user, split by commas or semicolons | [optional] 
**File** | Pointer to [***os.File**](*os.File.md) | Email CSV File To Search | [optional] 
**GatewayIds** | Pointer to **[]int64** | GatewayIds, Search Filter GatewayIds | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**PlanIds** | Pointer to **[]int64** | PlanIds, Search Filter PlanIds | [optional] 
**SearchKey** | Pointer to **string** | The search key of invoice | [optional] 
**SortField** | Pointer to **string** | Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | The status of invoice, 2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceCreditNoteListReq

`func NewUnibeeApiMerchantInvoiceCreditNoteListReq() *UnibeeApiMerchantInvoiceCreditNoteListReq`

NewUnibeeApiMerchantInvoiceCreditNoteListReq instantiates a new UnibeeApiMerchantInvoiceCreditNoteListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceCreditNoteListReqWithDefaults

`func NewUnibeeApiMerchantInvoiceCreditNoteListReqWithDefaults() *UnibeeApiMerchantInvoiceCreditNoteListReq`

NewUnibeeApiMerchantInvoiceCreditNoteListReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceCreditNoteListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmails

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetEmails() string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetEmailsOk() (*string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetEmails(v string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFile

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetGatewayIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetGatewayIds() []int64`

GetGatewayIds returns the GatewayIds field if non-nil, zero value otherwise.

### GetGatewayIdsOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetGatewayIdsOk() (*[]int64, bool)`

GetGatewayIdsOk returns a tuple with the GatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetGatewayIds(v []int64)`

SetGatewayIds sets GatewayIds field to given value.

### HasGatewayIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasGatewayIds() bool`

HasGatewayIds returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantInvoiceCreditNoteListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


