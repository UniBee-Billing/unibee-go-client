# UnibeeApiMerchantMetricEventCurrentValueReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email， UserId, ExternalUserId, or Email provides one of three options | [optional] [default to "account@unibee.dev"]
**ExternalUserId** | Pointer to **string** | ExternalUserId， UserId, ExternalUserId, or Email provides one of three options | [optional] 
**MetricCode** | **string** | MetricCode | 
**ProductId** | Pointer to **int64** | Id of product. Default product will use if productId not specified and subscriptionId is blank | [optional] 
**UserId** | Pointer to **int64** | UserId， UserId, ExternalUserId, or Email provides one of three options | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventCurrentValueReq

`func NewUnibeeApiMerchantMetricEventCurrentValueReq(metricCode string, ) *UnibeeApiMerchantMetricEventCurrentValueReq`

NewUnibeeApiMerchantMetricEventCurrentValueReq instantiates a new UnibeeApiMerchantMetricEventCurrentValueReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventCurrentValueReqWithDefaults

`func NewUnibeeApiMerchantMetricEventCurrentValueReqWithDefaults() *UnibeeApiMerchantMetricEventCurrentValueReq`

NewUnibeeApiMerchantMetricEventCurrentValueReqWithDefaults instantiates a new UnibeeApiMerchantMetricEventCurrentValueReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.


### GetProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


