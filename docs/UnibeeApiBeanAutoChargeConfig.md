# UnibeeApiBeanAutoChargeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoChargeLeadTime** | Pointer to **int32** | plan-level auto-charge lead time before period end (seconds); nil &#x3D; fall back to merchant TryAutomaticPaymentBeforePeriodEnd; value must be &gt;&#x3D; 300. | [optional] 
**AutoChargeMaxRetryCount** | Pointer to **int32** | plan-level max auto-charge attempts for one invoice, including the first automatic attempt; nil &#x3D; fall back; 0 &#x3D; unlimited. | [optional] 
**AutoChargeRetryInterval** | Pointer to **int32** | plan-level retry interval after a failed auto-charge (seconds); nil &#x3D; fall back to merchant/system default; value must be &gt;&#x3D; 300. | [optional] 
**IncompleteExpireTime** | Pointer to **int32** | plan-level incomplete/expire grace time after period/trial end (seconds); nil &#x3D; fall back to merchant/system default; 0 &#x3D; no grace period, expire immediately after period/trial end. | [optional] 

## Methods

### NewUnibeeApiBeanAutoChargeConfig

`func NewUnibeeApiBeanAutoChargeConfig() *UnibeeApiBeanAutoChargeConfig`

NewUnibeeApiBeanAutoChargeConfig instantiates a new UnibeeApiBeanAutoChargeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanAutoChargeConfigWithDefaults

`func NewUnibeeApiBeanAutoChargeConfigWithDefaults() *UnibeeApiBeanAutoChargeConfig`

NewUnibeeApiBeanAutoChargeConfigWithDefaults instantiates a new UnibeeApiBeanAutoChargeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoChargeLeadTime

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeLeadTime() int32`

GetAutoChargeLeadTime returns the AutoChargeLeadTime field if non-nil, zero value otherwise.

### GetAutoChargeLeadTimeOk

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeLeadTimeOk() (*int32, bool)`

GetAutoChargeLeadTimeOk returns a tuple with the AutoChargeLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeLeadTime

`func (o *UnibeeApiBeanAutoChargeConfig) SetAutoChargeLeadTime(v int32)`

SetAutoChargeLeadTime sets AutoChargeLeadTime field to given value.

### HasAutoChargeLeadTime

`func (o *UnibeeApiBeanAutoChargeConfig) HasAutoChargeLeadTime() bool`

HasAutoChargeLeadTime returns a boolean if a field has been set.

### GetAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeMaxRetryCount() int32`

GetAutoChargeMaxRetryCount returns the AutoChargeMaxRetryCount field if non-nil, zero value otherwise.

### GetAutoChargeMaxRetryCountOk

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeMaxRetryCountOk() (*int32, bool)`

GetAutoChargeMaxRetryCountOk returns a tuple with the AutoChargeMaxRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanAutoChargeConfig) SetAutoChargeMaxRetryCount(v int32)`

SetAutoChargeMaxRetryCount sets AutoChargeMaxRetryCount field to given value.

### HasAutoChargeMaxRetryCount

`func (o *UnibeeApiBeanAutoChargeConfig) HasAutoChargeMaxRetryCount() bool`

HasAutoChargeMaxRetryCount returns a boolean if a field has been set.

### GetAutoChargeRetryInterval

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeRetryInterval() int32`

GetAutoChargeRetryInterval returns the AutoChargeRetryInterval field if non-nil, zero value otherwise.

### GetAutoChargeRetryIntervalOk

`func (o *UnibeeApiBeanAutoChargeConfig) GetAutoChargeRetryIntervalOk() (*int32, bool)`

GetAutoChargeRetryIntervalOk returns a tuple with the AutoChargeRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeRetryInterval

`func (o *UnibeeApiBeanAutoChargeConfig) SetAutoChargeRetryInterval(v int32)`

SetAutoChargeRetryInterval sets AutoChargeRetryInterval field to given value.

### HasAutoChargeRetryInterval

`func (o *UnibeeApiBeanAutoChargeConfig) HasAutoChargeRetryInterval() bool`

HasAutoChargeRetryInterval returns a boolean if a field has been set.

### GetIncompleteExpireTime

`func (o *UnibeeApiBeanAutoChargeConfig) GetIncompleteExpireTime() int32`

GetIncompleteExpireTime returns the IncompleteExpireTime field if non-nil, zero value otherwise.

### GetIncompleteExpireTimeOk

`func (o *UnibeeApiBeanAutoChargeConfig) GetIncompleteExpireTimeOk() (*int32, bool)`

GetIncompleteExpireTimeOk returns a tuple with the IncompleteExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteExpireTime

`func (o *UnibeeApiBeanAutoChargeConfig) SetIncompleteExpireTime(v int32)`

SetIncompleteExpireTime sets IncompleteExpireTime field to given value.

### HasIncompleteExpireTime

`func (o *UnibeeApiBeanAutoChargeConfig) HasIncompleteExpireTime() bool`

HasIncompleteExpireTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


