# UnibeeApiBeanUserMetricChargeInvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsedValue** | Pointer to **int64** | CurrentUsedValue | [optional] 
**ChargePricing** | Pointer to [**UnibeeApiBeanPlanMetricMeteredChargeParam**](UnibeeApiBeanPlanMetricMeteredChargeParam.md) |  | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanMetricPlanChargeLine**](UnibeeApiBeanMetricPlanChargeLine.md) | Lines | [optional] 
**MaxEventId** | Pointer to **int64** |  | [optional] 
**MetricId** | **int64** | MetricId | 
**MinEventId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**TotalChargeAmount** | Pointer to **int64** | TotalChargeAmount | [optional] 

## Methods

### NewUnibeeApiBeanUserMetricChargeInvoiceItem

`func NewUnibeeApiBeanUserMetricChargeInvoiceItem(metricId int64, ) *UnibeeApiBeanUserMetricChargeInvoiceItem`

NewUnibeeApiBeanUserMetricChargeInvoiceItem instantiates a new UnibeeApiBeanUserMetricChargeInvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUserMetricChargeInvoiceItemWithDefaults

`func NewUnibeeApiBeanUserMetricChargeInvoiceItemWithDefaults() *UnibeeApiBeanUserMetricChargeInvoiceItem`

NewUnibeeApiBeanUserMetricChargeInvoiceItemWithDefaults instantiates a new UnibeeApiBeanUserMetricChargeInvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsedValue

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetCurrentUsedValue() int64`

GetCurrentUsedValue returns the CurrentUsedValue field if non-nil, zero value otherwise.

### GetCurrentUsedValueOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetCurrentUsedValueOk() (*int64, bool)`

GetCurrentUsedValueOk returns a tuple with the CurrentUsedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsedValue

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetCurrentUsedValue(v int64)`

SetCurrentUsedValue sets CurrentUsedValue field to given value.

### HasCurrentUsedValue

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasCurrentUsedValue() bool`

HasCurrentUsedValue returns a boolean if a field has been set.

### GetChargePricing

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetChargePricing() UnibeeApiBeanPlanMetricMeteredChargeParam`

GetChargePricing returns the ChargePricing field if non-nil, zero value otherwise.

### GetChargePricingOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetChargePricingOk() (*UnibeeApiBeanPlanMetricMeteredChargeParam, bool)`

GetChargePricingOk returns a tuple with the ChargePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargePricing

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetChargePricing(v UnibeeApiBeanPlanMetricMeteredChargeParam)`

SetChargePricing sets ChargePricing field to given value.

### HasChargePricing

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasChargePricing() bool`

HasChargePricing returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetLines() []UnibeeApiBeanMetricPlanChargeLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetLinesOk() (*[]UnibeeApiBeanMetricPlanChargeLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetLines(v []UnibeeApiBeanMetricPlanChargeLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMaxEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMaxEventId() int64`

GetMaxEventId returns the MaxEventId field if non-nil, zero value otherwise.

### GetMaxEventIdOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMaxEventIdOk() (*int64, bool)`

GetMaxEventIdOk returns a tuple with the MaxEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMaxEventId(v int64)`

SetMaxEventId sets MaxEventId field to given value.

### HasMaxEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasMaxEventId() bool`

HasMaxEventId returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.


### GetMinEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMinEventId() int64`

GetMinEventId returns the MinEventId field if non-nil, zero value otherwise.

### GetMinEventIdOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMinEventIdOk() (*int64, bool)`

GetMinEventIdOk returns a tuple with the MinEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMinEventId(v int64)`

SetMinEventId sets MinEventId field to given value.

### HasMinEventId

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasMinEventId() bool`

HasMinEventId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalChargeAmount

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetTotalChargeAmount() int64`

GetTotalChargeAmount returns the TotalChargeAmount field if non-nil, zero value otherwise.

### GetTotalChargeAmountOk

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetTotalChargeAmountOk() (*int64, bool)`

GetTotalChargeAmountOk returns a tuple with the TotalChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChargeAmount

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetTotalChargeAmount(v int64)`

SetTotalChargeAmount sets TotalChargeAmount field to given value.

### HasTotalChargeAmount

`func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasTotalChargeAmount() bool`

HasTotalChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


