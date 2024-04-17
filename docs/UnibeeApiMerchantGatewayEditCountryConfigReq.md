# UnibeeApiMerchantGatewayEditCountryConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryConfig** | Pointer to **map[string]bool** | The country config of payment gateway, a map with countryCode as key, and value for enable or disable | [optional] 
**GatewayId** | **int64** | The id of payment gateway | 

## Methods

### NewUnibeeApiMerchantGatewayEditCountryConfigReq

`func NewUnibeeApiMerchantGatewayEditCountryConfigReq(gatewayId int64, ) *UnibeeApiMerchantGatewayEditCountryConfigReq`

NewUnibeeApiMerchantGatewayEditCountryConfigReq instantiates a new UnibeeApiMerchantGatewayEditCountryConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewayEditCountryConfigReqWithDefaults

`func NewUnibeeApiMerchantGatewayEditCountryConfigReqWithDefaults() *UnibeeApiMerchantGatewayEditCountryConfigReq`

NewUnibeeApiMerchantGatewayEditCountryConfigReqWithDefaults instantiates a new UnibeeApiMerchantGatewayEditCountryConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryConfig

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetCountryConfig() map[string]bool`

GetCountryConfig returns the CountryConfig field if non-nil, zero value otherwise.

### GetCountryConfigOk

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetCountryConfigOk() (*map[string]bool, bool)`

GetCountryConfigOk returns a tuple with the CountryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryConfig

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) SetCountryConfig(v map[string]bool)`

SetCountryConfig sets CountryConfig field to given value.

### HasCountryConfig

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) HasCountryConfig() bool`

HasCountryConfig returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantGatewayEditCountryConfigReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


