# UnibeeApiMerchantMetricUserMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addon | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**IsPaid** | Pointer to **bool** | IsPaid | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Product** | Pointer to [**UnibeeApiBeanProduct**](UnibeeApiBeanProduct.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 
**UserMerchantMetricStats** | Pointer to [**[]UnibeeApiBeanUserMerchantMetricStat**](UnibeeApiBeanUserMerchantMetricStat.md) | UserMerchantMetricStats | [optional] 

## Methods

### NewUnibeeApiMerchantMetricUserMetric

`func NewUnibeeApiMerchantMetricUserMetric() *UnibeeApiMerchantMetricUserMetric`

NewUnibeeApiMerchantMetricUserMetric instantiates a new UnibeeApiMerchantMetricUserMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricUserMetricWithDefaults

`func NewUnibeeApiMerchantMetricUserMetricWithDefaults() *UnibeeApiMerchantMetricUserMetric`

NewUnibeeApiMerchantMetricUserMetricWithDefaults instantiates a new UnibeeApiMerchantMetricUserMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiMerchantMetricUserMetric) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiMerchantMetricUserMetric) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiMerchantMetricUserMetric) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantMetricUserMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantMetricUserMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantMetricUserMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiMerchantMetricUserMetric) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiMerchantMetricUserMetric) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiMerchantMetricUserMetric) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProduct

`func (o *UnibeeApiMerchantMetricUserMetric) GetProduct() UnibeeApiBeanProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetProductOk() (*UnibeeApiBeanProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UnibeeApiMerchantMetricUserMetric) SetProduct(v UnibeeApiBeanProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UnibeeApiMerchantMetricUserMetric) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantMetricUserMetric) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantMetricUserMetric) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantMetricUserMetric) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStats() []UnibeeApiBeanUserMerchantMetricStat`

GetUserMerchantMetricStats returns the UserMerchantMetricStats field if non-nil, zero value otherwise.

### GetUserMerchantMetricStatsOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStatsOk() (*[]UnibeeApiBeanUserMerchantMetricStat, bool)`

GetUserMerchantMetricStatsOk returns a tuple with the UserMerchantMetricStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) SetUserMerchantMetricStats(v []UnibeeApiBeanUserMerchantMetricStat)`

SetUserMerchantMetricStats sets UserMerchantMetricStats field to given value.

### HasUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) HasUserMerchantMetricStats() bool`

HasUserMerchantMetricStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


