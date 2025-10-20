# UnibeeApiMerchantVatNumberValidateHistoryReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Per Page | [optional] 
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**SearchKey** | Pointer to **string** | Search Key, vatNumber, validateGateway, company, company address, message | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | status, 0-Invalid，1-Valid | [optional] 
**ValidateGateway** | Pointer to **string** | Filter Validate Gateway, vatsense | [optional] 
**VatNumber** | Pointer to **string** | Filter Vat Number | [optional] 

## Methods

### NewUnibeeApiMerchantVatNumberValidateHistoryReq

`func NewUnibeeApiMerchantVatNumberValidateHistoryReq() *UnibeeApiMerchantVatNumberValidateHistoryReq`

NewUnibeeApiMerchantVatNumberValidateHistoryReq instantiates a new UnibeeApiMerchantVatNumberValidateHistoryReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantVatNumberValidateHistoryReqWithDefaults

`func NewUnibeeApiMerchantVatNumberValidateHistoryReqWithDefaults() *UnibeeApiMerchantVatNumberValidateHistoryReq`

NewUnibeeApiMerchantVatNumberValidateHistoryReqWithDefaults instantiates a new UnibeeApiMerchantVatNumberValidateHistoryReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSearchKey

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidateGateway

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetValidateGateway() string`

GetValidateGateway returns the ValidateGateway field if non-nil, zero value otherwise.

### GetValidateGatewayOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetValidateGatewayOk() (*string, bool)`

GetValidateGatewayOk returns a tuple with the ValidateGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateGateway

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetValidateGateway(v string)`

SetValidateGateway sets ValidateGateway field to given value.

### HasValidateGateway

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasValidateGateway() bool`

HasValidateGateway returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantVatNumberValidateHistoryReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


