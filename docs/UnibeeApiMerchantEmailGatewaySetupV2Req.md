# UnibeeApiMerchantEmailGatewaySetupV2Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | Whether setup the gateway as default or not, default is false | [optional] 
**ApiCredential** | Pointer to [**UnibeeApiBeanEmailGatewayConnectionAPIKeysUpdate**](UnibeeApiBeanEmailGatewayConnectionAPIKeysUpdate.md) |  | [optional] 
**GatewayName** | **string** | The name of email gateway, available for sendgrid|smtp | 

## Methods

### NewUnibeeApiMerchantEmailGatewaySetupV2Req

`func NewUnibeeApiMerchantEmailGatewaySetupV2Req(gatewayName string, ) *UnibeeApiMerchantEmailGatewaySetupV2Req`

NewUnibeeApiMerchantEmailGatewaySetupV2Req instantiates a new UnibeeApiMerchantEmailGatewaySetupV2Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailGatewaySetupV2ReqWithDefaults

`func NewUnibeeApiMerchantEmailGatewaySetupV2ReqWithDefaults() *UnibeeApiMerchantEmailGatewaySetupV2Req`

NewUnibeeApiMerchantEmailGatewaySetupV2ReqWithDefaults instantiates a new UnibeeApiMerchantEmailGatewaySetupV2Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetApiCredential

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetApiCredential() UnibeeApiBeanEmailGatewayConnectionAPIKeysUpdate`

GetApiCredential returns the ApiCredential field if non-nil, zero value otherwise.

### GetApiCredentialOk

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetApiCredentialOk() (*UnibeeApiBeanEmailGatewayConnectionAPIKeysUpdate, bool)`

GetApiCredentialOk returns a tuple with the ApiCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredential

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) SetApiCredential(v UnibeeApiBeanEmailGatewayConnectionAPIKeysUpdate)`

SetApiCredential sets ApiCredential field to given value.

### HasApiCredential

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) HasApiCredential() bool`

HasApiCredential returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantEmailGatewaySetupV2Req) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


