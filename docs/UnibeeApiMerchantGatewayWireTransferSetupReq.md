# UnibeeApiMerchantGatewayWireTransferSetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**UnibeeApiBeanGatewayBank**](UnibeeApiBeanGatewayBank.md) |  | [optional] 
**Currency** | **string** | The currency of wire transfer  | 
**MinimumAmount** | **int64** | The minimum amount of wire transfer, cents | 

## Methods

### NewUnibeeApiMerchantGatewayWireTransferSetupReq

`func NewUnibeeApiMerchantGatewayWireTransferSetupReq(currency string, minimumAmount int64, ) *UnibeeApiMerchantGatewayWireTransferSetupReq`

NewUnibeeApiMerchantGatewayWireTransferSetupReq instantiates a new UnibeeApiMerchantGatewayWireTransferSetupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewayWireTransferSetupReqWithDefaults

`func NewUnibeeApiMerchantGatewayWireTransferSetupReqWithDefaults() *UnibeeApiMerchantGatewayWireTransferSetupReq`

NewUnibeeApiMerchantGatewayWireTransferSetupReqWithDefaults instantiates a new UnibeeApiMerchantGatewayWireTransferSetupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBank() UnibeeApiBeanGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBankOk() (*UnibeeApiBeanGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetBank(v UnibeeApiBeanGatewayBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMinimumAmount

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetMinimumAmount() int64`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetMinimumAmountOk() (*int64, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetMinimumAmount(v int64)`

SetMinimumAmount sets MinimumAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


