# UnibeeApiMerchantGatewaySetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayKey** | Pointer to **string** | GatewayKey | [optional] 
**GatewayName** | **string** | GatewayName, stripe|paypal|changelly | 
**GatewaySecret** | Pointer to **string** | GatewaySecret | [optional] 

## Methods

### NewUnibeeApiMerchantGatewaySetupReq

`func NewUnibeeApiMerchantGatewaySetupReq(gatewayName string, ) *UnibeeApiMerchantGatewaySetupReq`

NewUnibeeApiMerchantGatewaySetupReq instantiates a new UnibeeApiMerchantGatewaySetupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewaySetupReqWithDefaults

`func NewUnibeeApiMerchantGatewaySetupReqWithDefaults() *UnibeeApiMerchantGatewaySetupReq`

NewUnibeeApiMerchantGatewaySetupReqWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.


### GetGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecret() string`

GetGatewaySecret returns the GatewaySecret field if non-nil, zero value otherwise.

### GetGatewaySecretOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecretOk() (*string, bool)`

GetGatewaySecretOk returns a tuple with the GatewaySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewaySecret(v string)`

SetGatewaySecret sets GatewaySecret field to given value.

### HasGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewaySecret() bool`

HasGatewaySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


