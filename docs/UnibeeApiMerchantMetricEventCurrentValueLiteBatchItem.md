# UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 0 success; same business code as single lite on failure (e.g. 51) | [optional] 
**CurrentValue** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExternalUserId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**TotalLimit** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItem

`func NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItem() *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem`

NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItem instantiates a new UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItemWithDefaults

`func NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItemWithDefaults() *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem`

NewUnibeeApiMerchantMetricEventCurrentValueLiteBatchItemWithDefaults instantiates a new UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetCurrentValue() int64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetCurrentValueOk() (*int64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetCurrentValue(v int64)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetMessage

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetTotalLimit() int64`

GetTotalLimit returns the TotalLimit field if non-nil, zero value otherwise.

### GetTotalLimitOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetTotalLimitOk() (*int64, bool)`

GetTotalLimitOk returns a tuple with the TotalLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetTotalLimit(v int64)`

SetTotalLimit sets TotalLimit field to given value.

### HasTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasTotalLimit() bool`

HasTotalLimit returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricEventCurrentValueLiteBatchItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


