# UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Subscription recurring amount, minor units. | [optional] 
**Currency** | Pointer to **string** | Subscription currency. | [optional] 
**PlanId** | Pointer to **int64** | Plan id of the subscription. | [optional] 
**Status** | Pointer to **int32** | Subscription status, 1-Pending|2-Active|3-PendingInActive|4-Cancel|5-Expire|6-Suspend|7-Incomplete|8-Processing|9-Failed. | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription id. | [optional] 
**UserId** | Pointer to **int64** | Owner user id. | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscription

`func NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscription() *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription`

NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscription instantiates a new UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscriptionWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscriptionWithDefaults() *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription`

NewUnibeeApiCheckoutSubscriptionCheckoutCreateSubscriptionWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiCheckoutSubscriptionCheckoutCreateSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


