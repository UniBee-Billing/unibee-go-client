# UnibeeApiBeanMerchantZapierSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | subscription created at | [optional] 
**EndpointId** | Pointer to **int64** | linked endpoint id | [optional] 
**Events** | Pointer to **[]string** | subscription events | [optional] 
**HookUrl** | Pointer to **string** | zapier hook url | [optional] 
**Secret** | Pointer to **string** | zapier secret | [optional] 
**Status** | Pointer to **string** | subscription status | [optional] 
**UpdatedAt** | Pointer to **int64** | subscription updated at | [optional] 
**ZapId** | Pointer to **string** | zapier zap id | [optional] 

## Methods

### NewUnibeeApiBeanMerchantZapierSubscription

`func NewUnibeeApiBeanMerchantZapierSubscription() *UnibeeApiBeanMerchantZapierSubscription`

NewUnibeeApiBeanMerchantZapierSubscription instantiates a new UnibeeApiBeanMerchantZapierSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantZapierSubscriptionWithDefaults

`func NewUnibeeApiBeanMerchantZapierSubscriptionWithDefaults() *UnibeeApiBeanMerchantZapierSubscription`

NewUnibeeApiBeanMerchantZapierSubscriptionWithDefaults instantiates a new UnibeeApiBeanMerchantZapierSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndpointId

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetEndpointId() int64`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetEndpointIdOk() (*int64, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetEndpointId(v int64)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEvents

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHookUrl

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetHookUrl() string`

GetHookUrl returns the HookUrl field if non-nil, zero value otherwise.

### GetHookUrlOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetHookUrlOk() (*string, bool)`

GetHookUrlOk returns a tuple with the HookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookUrl

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetHookUrl(v string)`

SetHookUrl sets HookUrl field to given value.

### HasHookUrl

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasHookUrl() bool`

HasHookUrl returns a boolean if a field has been set.

### GetSecret

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetZapId

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetZapId() string`

GetZapId returns the ZapId field if non-nil, zero value otherwise.

### GetZapIdOk

`func (o *UnibeeApiBeanMerchantZapierSubscription) GetZapIdOk() (*string, bool)`

GetZapIdOk returns a tuple with the ZapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZapId

`func (o *UnibeeApiBeanMerchantZapierSubscription) SetZapId(v string)`

SetZapId sets ZapId field to given value.

### HasZapId

`func (o *UnibeeApiBeanMerchantZapierSubscription) HasZapId() bool`

HasZapId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


