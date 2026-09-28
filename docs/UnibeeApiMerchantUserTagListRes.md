# UnibeeApiMerchantUserTagListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationStatus** | Pointer to **string** | pending|building|ready|failed | [optional] 
**CalculationVersion** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]UnibeeApiBeanUserSystemTag**](UnibeeApiBeanUserSystemTag.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantUserTagListRes

`func NewUnibeeApiMerchantUserTagListRes() *UnibeeApiMerchantUserTagListRes`

NewUnibeeApiMerchantUserTagListRes instantiates a new UnibeeApiMerchantUserTagListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserTagListResWithDefaults

`func NewUnibeeApiMerchantUserTagListResWithDefaults() *UnibeeApiMerchantUserTagListRes`

NewUnibeeApiMerchantUserTagListResWithDefaults instantiates a new UnibeeApiMerchantUserTagListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationStatus

`func (o *UnibeeApiMerchantUserTagListRes) GetCalculationStatus() string`

GetCalculationStatus returns the CalculationStatus field if non-nil, zero value otherwise.

### GetCalculationStatusOk

`func (o *UnibeeApiMerchantUserTagListRes) GetCalculationStatusOk() (*string, bool)`

GetCalculationStatusOk returns a tuple with the CalculationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationStatus

`func (o *UnibeeApiMerchantUserTagListRes) SetCalculationStatus(v string)`

SetCalculationStatus sets CalculationStatus field to given value.

### HasCalculationStatus

`func (o *UnibeeApiMerchantUserTagListRes) HasCalculationStatus() bool`

HasCalculationStatus returns a boolean if a field has been set.

### GetCalculationVersion

`func (o *UnibeeApiMerchantUserTagListRes) GetCalculationVersion() int32`

GetCalculationVersion returns the CalculationVersion field if non-nil, zero value otherwise.

### GetCalculationVersionOk

`func (o *UnibeeApiMerchantUserTagListRes) GetCalculationVersionOk() (*int32, bool)`

GetCalculationVersionOk returns a tuple with the CalculationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationVersion

`func (o *UnibeeApiMerchantUserTagListRes) SetCalculationVersion(v int32)`

SetCalculationVersion sets CalculationVersion field to given value.

### HasCalculationVersion

`func (o *UnibeeApiMerchantUserTagListRes) HasCalculationVersion() bool`

HasCalculationVersion returns a boolean if a field has been set.

### GetTags

`func (o *UnibeeApiMerchantUserTagListRes) GetTags() []UnibeeApiBeanUserSystemTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UnibeeApiMerchantUserTagListRes) GetTagsOk() (*[]UnibeeApiBeanUserSystemTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UnibeeApiMerchantUserTagListRes) SetTags(v []UnibeeApiBeanUserSystemTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UnibeeApiMerchantUserTagListRes) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


