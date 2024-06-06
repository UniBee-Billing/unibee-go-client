# UnibeeApiMerchantSubscriptionListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]UnibeeApiBeanDetailSubscriptionDetail**](UnibeeApiBeanDetailSubscriptionDetail.md) | Subscriptions | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionListRes

`func NewUnibeeApiMerchantSubscriptionListRes() *UnibeeApiMerchantSubscriptionListRes`

NewUnibeeApiMerchantSubscriptionListRes instantiates a new UnibeeApiMerchantSubscriptionListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionListResWithDefaults

`func NewUnibeeApiMerchantSubscriptionListResWithDefaults() *UnibeeApiMerchantSubscriptionListRes`

NewUnibeeApiMerchantSubscriptionListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *UnibeeApiMerchantSubscriptionListRes) GetSubscriptions() []UnibeeApiBeanDetailSubscriptionDetail`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UnibeeApiMerchantSubscriptionListRes) GetSubscriptionsOk() (*[]UnibeeApiBeanDetailSubscriptionDetail, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UnibeeApiMerchantSubscriptionListRes) SetSubscriptions(v []UnibeeApiBeanDetailSubscriptionDetail)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UnibeeApiMerchantSubscriptionListRes) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantSubscriptionListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantSubscriptionListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantSubscriptionListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantSubscriptionListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


