# UnibeeApiSystemMcpdocSearchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Result limit, default 20, max 100 | [optional] 
**PathPrefix** | Pointer to **string** | Filter by path prefix | [optional] 
**Query** | Pointer to **string** | Keyword for path/summary/description | [optional] 
**Tag** | Pointer to **string** | Filter by exact tag | [optional] 

## Methods

### NewUnibeeApiSystemMcpdocSearchReq

`func NewUnibeeApiSystemMcpdocSearchReq() *UnibeeApiSystemMcpdocSearchReq`

NewUnibeeApiSystemMcpdocSearchReq instantiates a new UnibeeApiSystemMcpdocSearchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemMcpdocSearchReqWithDefaults

`func NewUnibeeApiSystemMcpdocSearchReqWithDefaults() *UnibeeApiSystemMcpdocSearchReq`

NewUnibeeApiSystemMcpdocSearchReqWithDefaults instantiates a new UnibeeApiSystemMcpdocSearchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *UnibeeApiSystemMcpdocSearchReq) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UnibeeApiSystemMcpdocSearchReq) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UnibeeApiSystemMcpdocSearchReq) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UnibeeApiSystemMcpdocSearchReq) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPathPrefix

`func (o *UnibeeApiSystemMcpdocSearchReq) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *UnibeeApiSystemMcpdocSearchReq) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *UnibeeApiSystemMcpdocSearchReq) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *UnibeeApiSystemMcpdocSearchReq) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetQuery

`func (o *UnibeeApiSystemMcpdocSearchReq) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UnibeeApiSystemMcpdocSearchReq) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UnibeeApiSystemMcpdocSearchReq) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UnibeeApiSystemMcpdocSearchReq) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTag

`func (o *UnibeeApiSystemMcpdocSearchReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UnibeeApiSystemMcpdocSearchReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UnibeeApiSystemMcpdocSearchReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UnibeeApiSystemMcpdocSearchReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


