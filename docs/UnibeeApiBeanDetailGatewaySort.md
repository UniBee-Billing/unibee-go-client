# UnibeeApiBeanDetailGatewaySort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **int64** | The gateway id | [optional] 
**GatewayName** | Pointer to **string** | Required, The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium | [optional] 
**Sort** | Pointer to **int64** | Required, The sort value of payment gateway, should greater than 0, The higher the value, the lower the ranking | [optional] 

## Methods

### NewUnibeeApiBeanDetailGatewaySort

`func NewUnibeeApiBeanDetailGatewaySort() *UnibeeApiBeanDetailGatewaySort`

NewUnibeeApiBeanDetailGatewaySort instantiates a new UnibeeApiBeanDetailGatewaySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailGatewaySortWithDefaults

`func NewUnibeeApiBeanDetailGatewaySortWithDefaults() *UnibeeApiBeanDetailGatewaySort`

NewUnibeeApiBeanDetailGatewaySortWithDefaults instantiates a new UnibeeApiBeanDetailGatewaySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailGatewaySort) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailGatewaySort) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanDetailGatewaySort) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanDetailGatewaySort) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanDetailGatewaySort) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetSort

`func (o *UnibeeApiBeanDetailGatewaySort) GetSort() int64`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UnibeeApiBeanDetailGatewaySort) GetSortOk() (*int64, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UnibeeApiBeanDetailGatewaySort) SetSort(v int64)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UnibeeApiBeanDetailGatewaySort) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


