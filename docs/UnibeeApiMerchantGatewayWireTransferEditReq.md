# UnibeeApiMerchantGatewayWireTransferEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**UnibeeApiBeanGatewayBank**](UnibeeApiBeanGatewayBank.md) |  | [optional] 
**Currency** | **string** | The currency of wire transfer  | 
**GatewayId** | **int64** | The id of payment gateway | 
**MinimumAmount** | **int64** | The minimum amount of wire transfer, cents | 

## Methods

### NewUnibeeApiMerchantGatewayWireTransferEditReq

`func NewUnibeeApiMerchantGatewayWireTransferEditReq(currency string, gatewayId int64, minimumAmount int64, ) *UnibeeApiMerchantGatewayWireTransferEditReq`

NewUnibeeApiMerchantGatewayWireTransferEditReq instantiates a new UnibeeApiMerchantGatewayWireTransferEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewayWireTransferEditReqWithDefaults

`func NewUnibeeApiMerchantGatewayWireTransferEditReqWithDefaults() *UnibeeApiMerchantGatewayWireTransferEditReq`

NewUnibeeApiMerchantGatewayWireTransferEditReqWithDefaults instantiates a new UnibeeApiMerchantGatewayWireTransferEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetBank() UnibeeApiBeanGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetBankOk() (*UnibeeApiBeanGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetBank(v UnibeeApiBeanGatewayBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetMinimumAmount

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetMinimumAmount() int64`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) GetMinimumAmountOk() (*int64, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnibeeApiMerchantGatewayWireTransferEditReq) SetMinimumAmount(v int64)`

SetMinimumAmount sets MinimumAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


