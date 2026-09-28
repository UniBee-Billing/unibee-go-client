# UnibeeApiBeanExpiringCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Card brand | [optional] 
**DaysUntilExpire** | Pointer to **int32** | Days from UTC today to expireDate | [optional] 
**ExpMonth** | Pointer to **int32** | Expiry month | [optional] 
**ExpYear** | Pointer to **int32** | Expiry year | [optional] 
**ExpireDate** | Pointer to **string** | Last calendar day of expiry month, UTC YYYY-MM-DD | [optional] 
**Id** | Pointer to **string** | Gateway payment method id | [optional] 
**Last4** | Pointer to **string** | Card last 4 | [optional] 
**LastSyncTime** | Pointer to **int64** | Last Stripe retrieve UTC seconds | [optional] 
**SubscriptionIds** | Pointer to **[]string** | Active subscriptions auto-charging this card | [optional] 

## Methods

### NewUnibeeApiBeanExpiringCard

`func NewUnibeeApiBeanExpiringCard() *UnibeeApiBeanExpiringCard`

NewUnibeeApiBeanExpiringCard instantiates a new UnibeeApiBeanExpiringCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanExpiringCardWithDefaults

`func NewUnibeeApiBeanExpiringCardWithDefaults() *UnibeeApiBeanExpiringCard`

NewUnibeeApiBeanExpiringCardWithDefaults instantiates a new UnibeeApiBeanExpiringCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *UnibeeApiBeanExpiringCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UnibeeApiBeanExpiringCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UnibeeApiBeanExpiringCard) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UnibeeApiBeanExpiringCard) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDaysUntilExpire

`func (o *UnibeeApiBeanExpiringCard) GetDaysUntilExpire() int32`

GetDaysUntilExpire returns the DaysUntilExpire field if non-nil, zero value otherwise.

### GetDaysUntilExpireOk

`func (o *UnibeeApiBeanExpiringCard) GetDaysUntilExpireOk() (*int32, bool)`

GetDaysUntilExpireOk returns a tuple with the DaysUntilExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilExpire

`func (o *UnibeeApiBeanExpiringCard) SetDaysUntilExpire(v int32)`

SetDaysUntilExpire sets DaysUntilExpire field to given value.

### HasDaysUntilExpire

`func (o *UnibeeApiBeanExpiringCard) HasDaysUntilExpire() bool`

HasDaysUntilExpire returns a boolean if a field has been set.

### GetExpMonth

`func (o *UnibeeApiBeanExpiringCard) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *UnibeeApiBeanExpiringCard) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *UnibeeApiBeanExpiringCard) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *UnibeeApiBeanExpiringCard) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *UnibeeApiBeanExpiringCard) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *UnibeeApiBeanExpiringCard) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *UnibeeApiBeanExpiringCard) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *UnibeeApiBeanExpiringCard) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetExpireDate

`func (o *UnibeeApiBeanExpiringCard) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *UnibeeApiBeanExpiringCard) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *UnibeeApiBeanExpiringCard) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *UnibeeApiBeanExpiringCard) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanExpiringCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanExpiringCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanExpiringCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanExpiringCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLast4

`func (o *UnibeeApiBeanExpiringCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *UnibeeApiBeanExpiringCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *UnibeeApiBeanExpiringCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *UnibeeApiBeanExpiringCard) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *UnibeeApiBeanExpiringCard) GetLastSyncTime() int64`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *UnibeeApiBeanExpiringCard) GetLastSyncTimeOk() (*int64, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *UnibeeApiBeanExpiringCard) SetLastSyncTime(v int64)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *UnibeeApiBeanExpiringCard) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeApiBeanExpiringCard) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeApiBeanExpiringCard) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeApiBeanExpiringCard) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeApiBeanExpiringCard) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


