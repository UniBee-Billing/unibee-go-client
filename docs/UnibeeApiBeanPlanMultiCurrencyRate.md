# UnibeeApiBeanPlanMultiCurrencyRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Source currency (plan default currency) | [optional] 
**Rate** | Pointer to **float32** | Exchange rate from &#39;from&#39; to &#39;to&#39; at the time of billing | [optional] 
**To** | Pointer to **string** | Target charge currency (usually invoice currency) | [optional] 

## Methods

### NewUnibeeApiBeanPlanMultiCurrencyRate

`func NewUnibeeApiBeanPlanMultiCurrencyRate() *UnibeeApiBeanPlanMultiCurrencyRate`

NewUnibeeApiBeanPlanMultiCurrencyRate instantiates a new UnibeeApiBeanPlanMultiCurrencyRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanMultiCurrencyRateWithDefaults

`func NewUnibeeApiBeanPlanMultiCurrencyRateWithDefaults() *UnibeeApiBeanPlanMultiCurrencyRate`

NewUnibeeApiBeanPlanMultiCurrencyRateWithDefaults instantiates a new UnibeeApiBeanPlanMultiCurrencyRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRate

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTo

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *UnibeeApiBeanPlanMultiCurrencyRate) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


