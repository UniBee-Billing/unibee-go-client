# UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | Emails; provide exactly one of userIds, emails, or externalUserIds | [optional] 
**ExternalUserIds** | Pointer to **[]string** | ExternalUserIds; provide exactly one of userIds, emails, or externalUserIds | [optional] 
**MetricCode** | **string** | MetricCode | 
**ProductId** | Pointer to **int64** | Id of product. Default product will use if productId not specified | [optional] 
**UserIds** | Pointer to **[]int64** | UserIds; provide exactly one of userIds, emails, or externalUserIds | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReq

`func NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReq(metricCode string, ) *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq`

NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReq instantiates a new UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReqWithDefaults

`func NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReqWithDefaults() *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq`

NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchReqWithDefaults instantiates a new UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetExternalUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetExternalUserIds() []string`

GetExternalUserIds returns the ExternalUserIds field if non-nil, zero value otherwise.

### GetExternalUserIdsOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetExternalUserIdsOk() (*[]string, bool)`

GetExternalUserIdsOk returns a tuple with the ExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) SetExternalUserIds(v []string)`

SetExternalUserIds sets ExternalUserIds field to given value.

### HasExternalUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) HasExternalUserIds() bool`

HasExternalUserIds returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.


### GetProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchReq) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


