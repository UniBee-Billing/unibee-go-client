# UnibeeApiMerchantTagEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]int64** | Plan ids bound to this tag. Omit &#x3D; unchanged; [] &#x3D; clear; non-empty &#x3D; full replace. | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**TagId** | **int64** |  | 

## Methods

### NewUnibeeApiMerchantTagEditReq

`func NewUnibeeApiMerchantTagEditReq(tagId int64, ) *UnibeeApiMerchantTagEditReq`

NewUnibeeApiMerchantTagEditReq instantiates a new UnibeeApiMerchantTagEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTagEditReqWithDefaults

`func NewUnibeeApiMerchantTagEditReqWithDefaults() *UnibeeApiMerchantTagEditReq`

NewUnibeeApiMerchantTagEditReqWithDefaults instantiates a new UnibeeApiMerchantTagEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *UnibeeApiMerchantTagEditReq) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UnibeeApiMerchantTagEditReq) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UnibeeApiMerchantTagEditReq) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UnibeeApiMerchantTagEditReq) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantTagEditReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantTagEditReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantTagEditReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantTagEditReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantTagEditReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantTagEditReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantTagEditReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantTagEditReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantTagEditReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantTagEditReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantTagEditReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantTagEditReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetSortOrder

`func (o *UnibeeApiMerchantTagEditReq) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UnibeeApiMerchantTagEditReq) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UnibeeApiMerchantTagEditReq) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UnibeeApiMerchantTagEditReq) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTagId

`func (o *UnibeeApiMerchantTagEditReq) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UnibeeApiMerchantTagEditReq) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UnibeeApiMerchantTagEditReq) SetTagId(v int64)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


