# UnibeeInternalModelEntityOverseaPayPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | amount, cent, without tax | [optional] 
**BindingAddonIds** | Pointer to **string** | binded addon planIds，split with , | [optional] 
**CompanyId** | Pointer to **int64** | company id | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**ExtraMetricData** | Pointer to **string** |  | [optional] 
**GatewayProductDescription** | Pointer to **string** | gateway product description | [optional] 
**GatewayProductName** | Pointer to **string** | gateway product name | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**IntervalCount** | Pointer to **int32** | period unit count | [optional] 
**IntervalUnit** | Pointer to **string** | period unit,day|month|year|week | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MetaData** | Pointer to **string** | meta_data(json) | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**PublishStatus** | Pointer to **int32** | 1-UnPublish,2-Publish, Use For Display Plan At UserPortal | [optional] 
**Status** | Pointer to **int32** | status，1-editing，2-active，3-inactive，4-expired | [optional] 
**TaxInclusive** | Pointer to **int32** | deperated | [optional] 
**TaxScale** | Pointer to **int32** | tax scale 1000 &#x3D; 10% | [optional] 
**Type** | Pointer to **int32** | type，1-main plan，2-addon plan | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayPlan

`func NewUnibeeInternalModelEntityOverseaPayPlan() *UnibeeInternalModelEntityOverseaPayPlan`

NewUnibeeInternalModelEntityOverseaPayPlan instantiates a new UnibeeInternalModelEntityOverseaPayPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayPlanWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayPlanWithDefaults() *UnibeeInternalModelEntityOverseaPayPlan`

NewUnibeeInternalModelEntityOverseaPayPlanWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBindingAddonIds

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetBindingAddonIds() string`

GetBindingAddonIds returns the BindingAddonIds field if non-nil, zero value otherwise.

### GetBindingAddonIdsOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetBindingAddonIdsOk() (*string, bool)`

GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingAddonIds

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetBindingAddonIds(v string)`

SetBindingAddonIds sets BindingAddonIds field to given value.

### HasBindingAddonIds

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasBindingAddonIds() bool`

HasBindingAddonIds returns a boolean if a field has been set.

### GetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraMetricData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetExtraMetricData() string`

GetExtraMetricData returns the ExtraMetricData field if non-nil, zero value otherwise.

### GetExtraMetricDataOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetExtraMetricDataOk() (*string, bool)`

GetExtraMetricDataOk returns a tuple with the ExtraMetricData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMetricData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetExtraMetricData(v string)`

SetExtraMetricData sets ExtraMetricData field to given value.

### HasExtraMetricData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasExtraMetricData() bool`

HasExtraMetricData returns a boolean if a field has been set.

### GetGatewayProductDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductDescription() string`

GetGatewayProductDescription returns the GatewayProductDescription field if non-nil, zero value otherwise.

### GetGatewayProductDescriptionOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductDescriptionOk() (*string, bool)`

GetGatewayProductDescriptionOk returns a tuple with the GatewayProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayProductDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGatewayProductDescription(v string)`

SetGatewayProductDescription sets GatewayProductDescription field to given value.

### HasGatewayProductDescription

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGatewayProductDescription() bool`

HasGatewayProductDescription returns a boolean if a field has been set.

### GetGatewayProductName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductName() string`

GetGatewayProductName returns the GatewayProductName field if non-nil, zero value otherwise.

### GetGatewayProductNameOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductNameOk() (*string, bool)`

GetGatewayProductNameOk returns a tuple with the GatewayProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayProductName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGatewayProductName(v string)`

SetGatewayProductName sets GatewayProductName field to given value.

### HasGatewayProductName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGatewayProductName() bool`

HasGatewayProductName returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPublishStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPublishStatus() int32`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPublishStatusOk() (*int32, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetPublishStatus(v int32)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxInclusive

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxInclusive() int32`

GetTaxInclusive returns the TaxInclusive field if non-nil, zero value otherwise.

### GetTaxInclusiveOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxInclusiveOk() (*int32, bool)`

GetTaxInclusiveOk returns a tuple with the TaxInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxInclusive

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetTaxInclusive(v int32)`

SetTaxInclusive sets TaxInclusive field to given value.

### HasTaxInclusive

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasTaxInclusive() bool`

HasTaxInclusive returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxScale() int32`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxScaleOk() (*int32, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetTaxScale(v int32)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalModelEntityOverseaPayPlan) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalModelEntityOverseaPayPlan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


