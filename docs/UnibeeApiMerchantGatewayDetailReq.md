# UnibeeApiMerchantGatewayDetailReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **int32** | The id of payment gateway, either gatewayId or gatewayName | [optional] 
**GatewayName** | Pointer to **string** | The name of payment gateway, , either gatewayId or gatewayName, stripe|paypal|changelly|unitpay|payssion|cryptadium | [optional] 

## Methods

### NewUnibeeApiMerchantGatewayDetailReq

`func NewUnibeeApiMerchantGatewayDetailReq() *UnibeeApiMerchantGatewayDetailReq`

NewUnibeeApiMerchantGatewayDetailReq instantiates a new UnibeeApiMerchantGatewayDetailReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewayDetailReqWithDefaults

`func NewUnibeeApiMerchantGatewayDetailReqWithDefaults() *UnibeeApiMerchantGatewayDetailReq`

NewUnibeeApiMerchantGatewayDetailReqWithDefaults instantiates a new UnibeeApiMerchantGatewayDetailReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantGatewayDetailReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantGatewayDetailReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantGatewayDetailReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiMerchantGatewayDetailReq) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


