# UnibeeApiMerchantEmailGatewaySetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | IsDefault, default is true | [optional] [default to true]
**Data** | **string** | data | 
**GatewayName** | **string** | GatewayName, e.m. sendgrid | 

## Methods

### NewUnibeeApiMerchantEmailGatewaySetupReq

`func NewUnibeeApiMerchantEmailGatewaySetupReq(data string, gatewayName string, ) *UnibeeApiMerchantEmailGatewaySetupReq`

NewUnibeeApiMerchantEmailGatewaySetupReq instantiates a new UnibeeApiMerchantEmailGatewaySetupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailGatewaySetupReqWithDefaults

`func NewUnibeeApiMerchantEmailGatewaySetupReqWithDefaults() *UnibeeApiMerchantEmailGatewaySetupReq`

NewUnibeeApiMerchantEmailGatewaySetupReqWithDefaults instantiates a new UnibeeApiMerchantEmailGatewaySetupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetData

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetData(v string)`

SetData sets Data field to given value.


### GetGatewayName

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantEmailGatewaySetupReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


