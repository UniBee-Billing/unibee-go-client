# UnibeeApiMerchantGatewayWireTransferSetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**UnibeeApiBeanDetailGatewayBank**](UnibeeApiBeanDetailGatewayBank.md) |  | [optional] 
**Currency** | **string** | The currency of wire transfer  | 
**DisplayName** | Pointer to **string** | The displayName of payment gateway | [optional] 
**GatewayIcons** | Pointer to **[][]string** | The icons of payment gateway | [optional] 
**MinimumAmount** | **int64** | The minimum amount of wire transfer, cents | 
**Sort** | Pointer to **int32** | The sort value of payment gateway, The bigger, the closer to the front | [optional] 

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

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBank() UnibeeApiBeanDetailGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetBankOk() (*UnibeeApiBeanDetailGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetBank(v UnibeeApiBeanDetailGatewayBank)`

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


### GetDisplayName

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayIcons

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetGatewayIcons() [][]string`

GetGatewayIcons returns the GatewayIcons field if non-nil, zero value otherwise.

### GetGatewayIconsOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetGatewayIconsOk() (*[][]string, bool)`

GetGatewayIconsOk returns a tuple with the GatewayIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIcons

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetGatewayIcons(v [][]string)`

SetGatewayIcons sets GatewayIcons field to given value.

### HasGatewayIcons

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) HasGatewayIcons() bool`

HasGatewayIcons returns a boolean if a field has been set.

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


### GetSort

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UnibeeApiMerchantGatewayWireTransferSetupReq) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


