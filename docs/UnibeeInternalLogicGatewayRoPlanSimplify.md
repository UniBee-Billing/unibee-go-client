# UnibeeInternalLogicGatewayRoPlanSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | amount, cent, without tax | [optional] 
**BindingAddonIds** | Pointer to **string** | binded addon planIds，split with , | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**ExtraMetricData** | Pointer to **string** |  | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**IntervalCount** | Pointer to **int32** | period unit count | [optional] 
**IntervalUnit** | Pointer to **string** | period unit,day|month|year|week | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**ProductDescription** | Pointer to **string** | product description | [optional] 
**ProductName** | Pointer to **string** | product name | [optional] 
**PublishStatus** | Pointer to **int32** | 1-UnPublish,2-Publish, Use For Display Plan At UserPortal | [optional] 
**Status** | Pointer to **int32** | status，1-editing，2-active，3-inactive，4-expired | [optional] 
**TaxScale** | Pointer to **int32** | tax scale 1000 &#x3D; 10% | [optional] 
**Type** | Pointer to **int32** | type，1-main plan，2-addon plan | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoPlanSimplify

`func NewUnibeeInternalLogicGatewayRoPlanSimplify() *UnibeeInternalLogicGatewayRoPlanSimplify`

NewUnibeeInternalLogicGatewayRoPlanSimplify instantiates a new UnibeeInternalLogicGatewayRoPlanSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoPlanSimplifyWithDefaults

`func NewUnibeeInternalLogicGatewayRoPlanSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoPlanSimplify`

NewUnibeeInternalLogicGatewayRoPlanSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoPlanSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBindingAddonIds

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetBindingAddonIds() string`

GetBindingAddonIds returns the BindingAddonIds field if non-nil, zero value otherwise.

### GetBindingAddonIdsOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetBindingAddonIdsOk() (*string, bool)`

GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingAddonIds

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetBindingAddonIds(v string)`

SetBindingAddonIds sets BindingAddonIds field to given value.

### HasBindingAddonIds

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasBindingAddonIds() bool`

HasBindingAddonIds returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraMetricData

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetExtraMetricData() string`

GetExtraMetricData returns the ExtraMetricData field if non-nil, zero value otherwise.

### GetExtraMetricDataOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetExtraMetricDataOk() (*string, bool)`

GetExtraMetricDataOk returns a tuple with the ExtraMetricData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMetricData

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetExtraMetricData(v string)`

SetExtraMetricData sets ExtraMetricData field to given value.

### HasExtraMetricData

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasExtraMetricData() bool`

HasExtraMetricData returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetPublishStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetPublishStatus() int32`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetPublishStatusOk() (*int32, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetPublishStatus(v int32)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetTaxScale() int32`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetTaxScaleOk() (*int32, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetTaxScale(v int32)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalLogicGatewayRoPlanSimplify) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


