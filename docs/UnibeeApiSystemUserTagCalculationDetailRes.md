# UnibeeApiSystemUserTagCalculationDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationStatus** | Pointer to **string** |  | [optional] 
**CalculationVersion** | Pointer to **int32** |  | [optional] 
**Evidences** | Pointer to [**[]UnibeeApiBeanUserTagEvidence**](UnibeeApiBeanUserTagEvidence.md) |  | [optional] 
**LastCalculatedAt** | Pointer to **int64** |  | [optional] 
**LastErrorCode** | Pointer to **string** |  | [optional] 
**LeaseExpiresAt** | Pointer to **int64** |  | [optional] 
**LeaseToken** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]UnibeeApiBeanUserSystemTag**](UnibeeApiBeanUserSystemTag.md) |  | [optional] 

## Methods

### NewUnibeeApiSystemUserTagCalculationDetailRes

`func NewUnibeeApiSystemUserTagCalculationDetailRes() *UnibeeApiSystemUserTagCalculationDetailRes`

NewUnibeeApiSystemUserTagCalculationDetailRes instantiates a new UnibeeApiSystemUserTagCalculationDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemUserTagCalculationDetailResWithDefaults

`func NewUnibeeApiSystemUserTagCalculationDetailResWithDefaults() *UnibeeApiSystemUserTagCalculationDetailRes`

NewUnibeeApiSystemUserTagCalculationDetailResWithDefaults instantiates a new UnibeeApiSystemUserTagCalculationDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationStatus

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetCalculationStatus() string`

GetCalculationStatus returns the CalculationStatus field if non-nil, zero value otherwise.

### GetCalculationStatusOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetCalculationStatusOk() (*string, bool)`

GetCalculationStatusOk returns a tuple with the CalculationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationStatus

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetCalculationStatus(v string)`

SetCalculationStatus sets CalculationStatus field to given value.

### HasCalculationStatus

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasCalculationStatus() bool`

HasCalculationStatus returns a boolean if a field has been set.

### GetCalculationVersion

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetCalculationVersion() int32`

GetCalculationVersion returns the CalculationVersion field if non-nil, zero value otherwise.

### GetCalculationVersionOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetCalculationVersionOk() (*int32, bool)`

GetCalculationVersionOk returns a tuple with the CalculationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationVersion

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetCalculationVersion(v int32)`

SetCalculationVersion sets CalculationVersion field to given value.

### HasCalculationVersion

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasCalculationVersion() bool`

HasCalculationVersion returns a boolean if a field has been set.

### GetEvidences

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetEvidences() []UnibeeApiBeanUserTagEvidence`

GetEvidences returns the Evidences field if non-nil, zero value otherwise.

### GetEvidencesOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetEvidencesOk() (*[]UnibeeApiBeanUserTagEvidence, bool)`

GetEvidencesOk returns a tuple with the Evidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidences

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetEvidences(v []UnibeeApiBeanUserTagEvidence)`

SetEvidences sets Evidences field to given value.

### HasEvidences

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasEvidences() bool`

HasEvidences returns a boolean if a field has been set.

### GetLastCalculatedAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLastCalculatedAt() int64`

GetLastCalculatedAt returns the LastCalculatedAt field if non-nil, zero value otherwise.

### GetLastCalculatedAtOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLastCalculatedAtOk() (*int64, bool)`

GetLastCalculatedAtOk returns a tuple with the LastCalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculatedAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetLastCalculatedAt(v int64)`

SetLastCalculatedAt sets LastCalculatedAt field to given value.

### HasLastCalculatedAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasLastCalculatedAt() bool`

HasLastCalculatedAt returns a boolean if a field has been set.

### GetLastErrorCode

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLastErrorCode() string`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLastErrorCodeOk() (*string, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetLastErrorCode(v string)`

SetLastErrorCode sets LastErrorCode field to given value.

### HasLastErrorCode

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasLastErrorCode() bool`

HasLastErrorCode returns a boolean if a field has been set.

### GetLeaseExpiresAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLeaseExpiresAt() int64`

GetLeaseExpiresAt returns the LeaseExpiresAt field if non-nil, zero value otherwise.

### GetLeaseExpiresAtOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLeaseExpiresAtOk() (*int64, bool)`

GetLeaseExpiresAtOk returns a tuple with the LeaseExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiresAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetLeaseExpiresAt(v int64)`

SetLeaseExpiresAt sets LeaseExpiresAt field to given value.

### HasLeaseExpiresAt

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasLeaseExpiresAt() bool`

HasLeaseExpiresAt returns a boolean if a field has been set.

### GetLeaseToken

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLeaseToken() string`

GetLeaseToken returns the LeaseToken field if non-nil, zero value otherwise.

### GetLeaseTokenOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetLeaseTokenOk() (*string, bool)`

GetLeaseTokenOk returns a tuple with the LeaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseToken

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetLeaseToken(v string)`

SetLeaseToken sets LeaseToken field to given value.

### HasLeaseToken

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasLeaseToken() bool`

HasLeaseToken returns a boolean if a field has been set.

### GetTags

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetTags() []UnibeeApiBeanUserSystemTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) GetTagsOk() (*[]UnibeeApiBeanUserSystemTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) SetTags(v []UnibeeApiBeanUserSystemTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UnibeeApiSystemUserTagCalculationDetailRes) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


