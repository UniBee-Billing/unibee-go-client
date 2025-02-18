# UnibeeApiMerchantMetricEventListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]UnibeeApiBeanDetailMerchantMetricEventDetail**](UnibeeApiBeanDetailMerchantMetricEventDetail.md) | User Metric Event List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventListRes

`func NewUnibeeApiMerchantMetricEventListRes() *UnibeeApiMerchantMetricEventListRes`

NewUnibeeApiMerchantMetricEventListRes instantiates a new UnibeeApiMerchantMetricEventListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventListResWithDefaults

`func NewUnibeeApiMerchantMetricEventListResWithDefaults() *UnibeeApiMerchantMetricEventListRes`

NewUnibeeApiMerchantMetricEventListResWithDefaults instantiates a new UnibeeApiMerchantMetricEventListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *UnibeeApiMerchantMetricEventListRes) GetEvents() []UnibeeApiBeanDetailMerchantMetricEventDetail`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UnibeeApiMerchantMetricEventListRes) GetEventsOk() (*[]UnibeeApiBeanDetailMerchantMetricEventDetail, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UnibeeApiMerchantMetricEventListRes) SetEvents(v []UnibeeApiBeanDetailMerchantMetricEventDetail)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UnibeeApiMerchantMetricEventListRes) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantMetricEventListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantMetricEventListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantMetricEventListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantMetricEventListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


