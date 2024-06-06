# UnibeeApiMerchantGatewaySetupWebhookReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **int64** | The id of payment gateway | 
**WebhookSecret** | Pointer to **string** | The secret of gateway webhook | [optional] 

## Methods

### NewUnibeeApiMerchantGatewaySetupWebhookReq

`func NewUnibeeApiMerchantGatewaySetupWebhookReq(gatewayId int64, ) *UnibeeApiMerchantGatewaySetupWebhookReq`

NewUnibeeApiMerchantGatewaySetupWebhookReq instantiates a new UnibeeApiMerchantGatewaySetupWebhookReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewaySetupWebhookReqWithDefaults

`func NewUnibeeApiMerchantGatewaySetupWebhookReqWithDefaults() *UnibeeApiMerchantGatewaySetupWebhookReq`

NewUnibeeApiMerchantGatewaySetupWebhookReqWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupWebhookReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetWebhookSecret

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *UnibeeApiMerchantGatewaySetupWebhookReq) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


