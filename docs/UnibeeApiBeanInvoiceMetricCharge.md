# UnibeeApiBeanInvoiceMetricCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **int32** | Charge type: 0-standard pricing, 1-graduated pricing | [optional] 
**MetricCode** | Pointer to **string** | Metric code | [optional] 
**MetricId** | Pointer to **int64** | Metric ID | [optional] 
**MetricName** | Pointer to **string** | Metric name | [optional] 
**MetricType** | Pointer to **int32** | Metric type: 2-charge_metered, 3-charge_recurring | [optional] 
**StandardAmount** | Pointer to **int64** | Standard amount per unit (cent), used when chargeType&#x3D;0 | [optional] 
**StandardStartValue** | Pointer to **int64** | Start value for standard pricing | [optional] 

## Methods

### NewUnibeeApiBeanInvoiceMetricCharge

`func NewUnibeeApiBeanInvoiceMetricCharge() *UnibeeApiBeanInvoiceMetricCharge`

NewUnibeeApiBeanInvoiceMetricCharge instantiates a new UnibeeApiBeanInvoiceMetricCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanInvoiceMetricChargeWithDefaults

`func NewUnibeeApiBeanInvoiceMetricChargeWithDefaults() *UnibeeApiBeanInvoiceMetricCharge`

NewUnibeeApiBeanInvoiceMetricChargeWithDefaults instantiates a new UnibeeApiBeanInvoiceMetricCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetChargeType() int32`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetChargeTypeOk() (*int32, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetChargeType(v int32)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.

### HasMetricCode

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasMetricCode() bool`

HasMetricCode returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricType

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricType() int32`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetMetricTypeOk() (*int32, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetMetricType(v int32)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetStandardAmount

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetStandardAmount() int64`

GetStandardAmount returns the StandardAmount field if non-nil, zero value otherwise.

### GetStandardAmountOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetStandardAmountOk() (*int64, bool)`

GetStandardAmountOk returns a tuple with the StandardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardAmount

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetStandardAmount(v int64)`

SetStandardAmount sets StandardAmount field to given value.

### HasStandardAmount

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasStandardAmount() bool`

HasStandardAmount returns a boolean if a field has been set.

### GetStandardStartValue

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetStandardStartValue() int64`

GetStandardStartValue returns the StandardStartValue field if non-nil, zero value otherwise.

### GetStandardStartValueOk

`func (o *UnibeeApiBeanInvoiceMetricCharge) GetStandardStartValueOk() (*int64, bool)`

GetStandardStartValueOk returns a tuple with the StandardStartValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardStartValue

`func (o *UnibeeApiBeanInvoiceMetricCharge) SetStandardStartValue(v int64)`

SetStandardStartValue sets StandardStartValue field to given value.

### HasStandardStartValue

`func (o *UnibeeApiBeanInvoiceMetricCharge) HasStandardStartValue() bool`

HasStandardStartValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


