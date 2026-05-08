# UnibeeApiSystemMcpdocEndpointRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description | [optional] 
**Error** | Pointer to **string** | Error Message | [optional] 
**Found** | Pointer to **bool** | Endpoint Found Or Not | [optional] 
**Method** | Pointer to **string** | HTTP Method | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Path** | Pointer to **string** | Endpoint Path | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**Responses** | Pointer to **map[string]interface{}** |  | [optional] 
**Security** | Pointer to **map[string]interface{}** |  | [optional] 
**SpecUrl** | Pointer to **string** | OpenAPI Spec URL | [optional] 
**Summary** | Pointer to **string** | Summary | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 

## Methods

### NewUnibeeApiSystemMcpdocEndpointRes

`func NewUnibeeApiSystemMcpdocEndpointRes() *UnibeeApiSystemMcpdocEndpointRes`

NewUnibeeApiSystemMcpdocEndpointRes instantiates a new UnibeeApiSystemMcpdocEndpointRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemMcpdocEndpointResWithDefaults

`func NewUnibeeApiSystemMcpdocEndpointResWithDefaults() *UnibeeApiSystemMcpdocEndpointRes`

NewUnibeeApiSystemMcpdocEndpointResWithDefaults instantiates a new UnibeeApiSystemMcpdocEndpointRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetError

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFound

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetMethod

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParameters

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRequestBody

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponses

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetResponses() map[string]interface{}`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetResponsesOk() (*map[string]interface{}, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetResponses(v map[string]interface{})`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSecurity

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSecurity() map[string]interface{}`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSecurityOk() (*map[string]interface{}, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetSecurity(v map[string]interface{})`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSpecUrl

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSpecUrl() string`

GetSpecUrl returns the SpecUrl field if non-nil, zero value otherwise.

### GetSpecUrlOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSpecUrlOk() (*string, bool)`

GetSpecUrlOk returns a tuple with the SpecUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecUrl

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetSpecUrl(v string)`

SetSpecUrl sets SpecUrl field to given value.

### HasSpecUrl

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasSpecUrl() bool`

HasSpecUrl returns a boolean if a field has been set.

### GetSummary

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UnibeeApiSystemMcpdocEndpointRes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UnibeeApiSystemMcpdocEndpointRes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UnibeeApiSystemMcpdocEndpointRes) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


