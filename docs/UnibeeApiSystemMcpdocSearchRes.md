# UnibeeApiSystemMcpdocSearchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UnibeeApiSystemMcpdocEndpointItem**](UnibeeApiSystemMcpdocEndpointItem.md) | Matched Endpoints | [optional] 
**SpecUrl** | Pointer to **string** | OpenAPI Spec URL | [optional] 
**Total** | Pointer to **int32** | Matched Item Count | [optional] 

## Methods

### NewUnibeeApiSystemMcpdocSearchRes

`func NewUnibeeApiSystemMcpdocSearchRes() *UnibeeApiSystemMcpdocSearchRes`

NewUnibeeApiSystemMcpdocSearchRes instantiates a new UnibeeApiSystemMcpdocSearchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemMcpdocSearchResWithDefaults

`func NewUnibeeApiSystemMcpdocSearchResWithDefaults() *UnibeeApiSystemMcpdocSearchRes`

NewUnibeeApiSystemMcpdocSearchResWithDefaults instantiates a new UnibeeApiSystemMcpdocSearchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UnibeeApiSystemMcpdocSearchRes) GetItems() []UnibeeApiSystemMcpdocEndpointItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiSystemMcpdocSearchRes) GetItemsOk() (*[]UnibeeApiSystemMcpdocEndpointItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiSystemMcpdocSearchRes) SetItems(v []UnibeeApiSystemMcpdocEndpointItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UnibeeApiSystemMcpdocSearchRes) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSpecUrl

`func (o *UnibeeApiSystemMcpdocSearchRes) GetSpecUrl() string`

GetSpecUrl returns the SpecUrl field if non-nil, zero value otherwise.

### GetSpecUrlOk

`func (o *UnibeeApiSystemMcpdocSearchRes) GetSpecUrlOk() (*string, bool)`

GetSpecUrlOk returns a tuple with the SpecUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecUrl

`func (o *UnibeeApiSystemMcpdocSearchRes) SetSpecUrl(v string)`

SetSpecUrl sets SpecUrl field to given value.

### HasSpecUrl

`func (o *UnibeeApiSystemMcpdocSearchRes) HasSpecUrl() bool`

HasSpecUrl returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiSystemMcpdocSearchRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiSystemMcpdocSearchRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiSystemMcpdocSearchRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiSystemMcpdocSearchRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


