# UnibeeApiMerchantTagNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Optional UI color | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Name** | **string** | Display name | 
**PlanIds** | Pointer to **[]int64** | Plan ids to bind to this tag after creation. Omit or [] &#x3D; no bindings. | [optional] 
**Scope** | **int32** | 0-plan, 1-user, 2-discount_code; P1 only 0 | 
**SortOrder** | Pointer to **int32** | Sort order in list UI | [optional] 

## Methods

### NewUnibeeApiMerchantTagNewReq

`func NewUnibeeApiMerchantTagNewReq(name string, scope int32, ) *UnibeeApiMerchantTagNewReq`

NewUnibeeApiMerchantTagNewReq instantiates a new UnibeeApiMerchantTagNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTagNewReqWithDefaults

`func NewUnibeeApiMerchantTagNewReqWithDefaults() *UnibeeApiMerchantTagNewReq`

NewUnibeeApiMerchantTagNewReqWithDefaults instantiates a new UnibeeApiMerchantTagNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *UnibeeApiMerchantTagNewReq) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UnibeeApiMerchantTagNewReq) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UnibeeApiMerchantTagNewReq) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UnibeeApiMerchantTagNewReq) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantTagNewReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantTagNewReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantTagNewReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantTagNewReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantTagNewReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantTagNewReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantTagNewReq) SetName(v string)`

SetName sets Name field to given value.


### GetPlanIds

`func (o *UnibeeApiMerchantTagNewReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantTagNewReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantTagNewReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantTagNewReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetScope

`func (o *UnibeeApiMerchantTagNewReq) GetScope() int32`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UnibeeApiMerchantTagNewReq) GetScopeOk() (*int32, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UnibeeApiMerchantTagNewReq) SetScope(v int32)`

SetScope sets Scope field to given value.


### GetSortOrder

`func (o *UnibeeApiMerchantTagNewReq) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UnibeeApiMerchantTagNewReq) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UnibeeApiMerchantTagNewReq) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UnibeeApiMerchantTagNewReq) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


