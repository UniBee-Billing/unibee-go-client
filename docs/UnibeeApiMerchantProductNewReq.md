# UnibeeApiMerchantProductNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**ProductName** | Pointer to **string** | ProductName | [optional] 
**Status** | Pointer to **int32** | status，1-active，2-inactive, default active | [optional] 
**UsVATConfig** | Pointer to [**UnibeeApiBeanUSVATConfig**](UnibeeApiBeanUSVATConfig.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantProductNewReq

`func NewUnibeeApiMerchantProductNewReq() *UnibeeApiMerchantProductNewReq`

NewUnibeeApiMerchantProductNewReq instantiates a new UnibeeApiMerchantProductNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProductNewReqWithDefaults

`func NewUnibeeApiMerchantProductNewReqWithDefaults() *UnibeeApiMerchantProductNewReq`

NewUnibeeApiMerchantProductNewReqWithDefaults instantiates a new UnibeeApiMerchantProductNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnibeeApiMerchantProductNewReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantProductNewReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantProductNewReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantProductNewReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiMerchantProductNewReq) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiMerchantProductNewReq) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiMerchantProductNewReq) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiMerchantProductNewReq) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiMerchantProductNewReq) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiMerchantProductNewReq) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiMerchantProductNewReq) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiMerchantProductNewReq) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantProductNewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantProductNewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantProductNewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantProductNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiMerchantProductNewReq) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiMerchantProductNewReq) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiMerchantProductNewReq) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiMerchantProductNewReq) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantProductNewReq) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantProductNewReq) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantProductNewReq) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantProductNewReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsVATConfig

`func (o *UnibeeApiMerchantProductNewReq) GetUsVATConfig() UnibeeApiBeanUSVATConfig`

GetUsVATConfig returns the UsVATConfig field if non-nil, zero value otherwise.

### GetUsVATConfigOk

`func (o *UnibeeApiMerchantProductNewReq) GetUsVATConfigOk() (*UnibeeApiBeanUSVATConfig, bool)`

GetUsVATConfigOk returns a tuple with the UsVATConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVATConfig

`func (o *UnibeeApiMerchantProductNewReq) SetUsVATConfig(v UnibeeApiBeanUSVATConfig)`

SetUsVATConfig sets UsVATConfig field to given value.

### HasUsVATConfig

`func (o *UnibeeApiMerchantProductNewReq) HasUsVATConfig() bool`

HasUsVATConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


