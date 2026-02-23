# UnibeeApiMerchantMetricListMetricQueryableInvoicesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count per page | [optional] 
**Page** | Pointer to **int32** | Page, start 0 | [optional] 
**SubscriptionId** | Pointer to **string** | Optional subscription id to filter by subscription | [optional] 
**UserId** | **int64** | UserId | 

## Methods

### NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReq

`func NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReq(userId int64, ) *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq`

NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReq instantiates a new UnibeeApiMerchantMetricListMetricQueryableInvoicesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReqWithDefaults

`func NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReqWithDefaults() *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq`

NewUnibeeApiMerchantMetricListMetricQueryableInvoicesReqWithDefaults instantiates a new UnibeeApiMerchantMetricListMetricQueryableInvoicesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricListMetricQueryableInvoicesReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


