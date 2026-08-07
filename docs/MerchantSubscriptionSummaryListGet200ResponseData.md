# MerchantSubscriptionSummaryListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Subscriptions** | Pointer to [**[]UnibeeApiMerchantSubscriptionSubscriptionSummary**](UnibeeApiMerchantSubscriptionSubscriptionSummary.md) |  | [optional] 

## Methods

### NewMerchantSubscriptionSummaryListGet200ResponseData

`func NewMerchantSubscriptionSummaryListGet200ResponseData() *MerchantSubscriptionSummaryListGet200ResponseData`

NewMerchantSubscriptionSummaryListGet200ResponseData instantiates a new MerchantSubscriptionSummaryListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionSummaryListGet200ResponseDataWithDefaults

`func NewMerchantSubscriptionSummaryListGet200ResponseDataWithDefaults() *MerchantSubscriptionSummaryListGet200ResponseData`

NewMerchantSubscriptionSummaryListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionSummaryListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetSubscriptions

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) GetSubscriptions() []UnibeeApiMerchantSubscriptionSubscriptionSummary`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) GetSubscriptionsOk() (*[]UnibeeApiMerchantSubscriptionSubscriptionSummary, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) SetSubscriptions(v []UnibeeApiMerchantSubscriptionSubscriptionSummary)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *MerchantSubscriptionSummaryListGet200ResponseData) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


