# UnibeeApiMerchantVatSetupUSVATGatewayReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | IsDefault, default is true | [optional] [default to true]
**ApiKeys** | Pointer to [**UnibeeApiBeanUSVATGatewayConnectionAPIKeys**](UnibeeApiBeanUSVATGatewayConnectionAPIKeys.md) |  | [optional] 
**GatewayName** | **string** | GatewayName, em. TaxJar | 

## Methods

### NewUnibeeApiMerchantVatSetupUSVATGatewayReq

`func NewUnibeeApiMerchantVatSetupUSVATGatewayReq(gatewayName string, ) *UnibeeApiMerchantVatSetupUSVATGatewayReq`

NewUnibeeApiMerchantVatSetupUSVATGatewayReq instantiates a new UnibeeApiMerchantVatSetupUSVATGatewayReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantVatSetupUSVATGatewayReqWithDefaults

`func NewUnibeeApiMerchantVatSetupUSVATGatewayReqWithDefaults() *UnibeeApiMerchantVatSetupUSVATGatewayReq`

NewUnibeeApiMerchantVatSetupUSVATGatewayReqWithDefaults instantiates a new UnibeeApiMerchantVatSetupUSVATGatewayReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetApiKeys

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetApiKeys() UnibeeApiBeanUSVATGatewayConnectionAPIKeys`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetApiKeysOk() (*UnibeeApiBeanUSVATGatewayConnectionAPIKeys, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) SetApiKeys(v UnibeeApiBeanUSVATGatewayConnectionAPIKeys)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantVatSetupUSVATGatewayReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


