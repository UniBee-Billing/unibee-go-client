# UnibeeApiBeanMerchantVatNumberVerifyHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAddress** | Pointer to **string** | company_address | [optional] 
**CompanyName** | Pointer to **string** | company_name | [optional] 
**CountryCode** | Pointer to **string** | country_code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**ManualValidate** | Pointer to **bool** | manual_validate | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Status** | Pointer to **int64** | status, 0-Invalid，1-Valid | [optional] 
**ValidateGateway** | Pointer to **string** | validate_gateway | [optional] 
**ValidateMessage** | Pointer to **string** | validate_message | [optional] 
**VatNumber** | Pointer to **string** | vat_number | [optional] 

## Methods

### NewUnibeeApiBeanMerchantVatNumberVerifyHistory

`func NewUnibeeApiBeanMerchantVatNumberVerifyHistory() *UnibeeApiBeanMerchantVatNumberVerifyHistory`

NewUnibeeApiBeanMerchantVatNumberVerifyHistory instantiates a new UnibeeApiBeanMerchantVatNumberVerifyHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantVatNumberVerifyHistoryWithDefaults

`func NewUnibeeApiBeanMerchantVatNumberVerifyHistoryWithDefaults() *UnibeeApiBeanMerchantVatNumberVerifyHistory`

NewUnibeeApiBeanMerchantVatNumberVerifyHistoryWithDefaults instantiates a new UnibeeApiBeanMerchantVatNumberVerifyHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAddress

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### GetCompanyName

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManualValidate

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetManualValidate() bool`

GetManualValidate returns the ManualValidate field if non-nil, zero value otherwise.

### GetManualValidateOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetManualValidateOk() (*bool, bool)`

GetManualValidateOk returns a tuple with the ManualValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualValidate

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetManualValidate(v bool)`

SetManualValidate sets ManualValidate field to given value.

### HasManualValidate

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasManualValidate() bool`

HasManualValidate returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidateGateway

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetValidateGateway() string`

GetValidateGateway returns the ValidateGateway field if non-nil, zero value otherwise.

### GetValidateGatewayOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetValidateGatewayOk() (*string, bool)`

GetValidateGatewayOk returns a tuple with the ValidateGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateGateway

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetValidateGateway(v string)`

SetValidateGateway sets ValidateGateway field to given value.

### HasValidateGateway

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasValidateGateway() bool`

HasValidateGateway returns a boolean if a field has been set.

### GetValidateMessage

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetValidateMessage() string`

GetValidateMessage returns the ValidateMessage field if non-nil, zero value otherwise.

### GetValidateMessageOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetValidateMessageOk() (*string, bool)`

GetValidateMessageOk returns a tuple with the ValidateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateMessage

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetValidateMessage(v string)`

SetValidateMessage sets ValidateMessage field to given value.

### HasValidateMessage

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasValidateMessage() bool`

HasValidateMessage returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanMerchantVatNumberVerifyHistory) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


