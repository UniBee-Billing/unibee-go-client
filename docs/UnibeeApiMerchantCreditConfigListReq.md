# UnibeeApiMerchantCreditConfigListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | currency | [optional] 
**Types** | Pointer to **[]int32** | type list of credit account, 1-main account, 2-promo credit account | [optional] 

## Methods

### NewUnibeeApiMerchantCreditConfigListReq

`func NewUnibeeApiMerchantCreditConfigListReq() *UnibeeApiMerchantCreditConfigListReq`

NewUnibeeApiMerchantCreditConfigListReq instantiates a new UnibeeApiMerchantCreditConfigListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditConfigListReqWithDefaults

`func NewUnibeeApiMerchantCreditConfigListReqWithDefaults() *UnibeeApiMerchantCreditConfigListReq`

NewUnibeeApiMerchantCreditConfigListReqWithDefaults instantiates a new UnibeeApiMerchantCreditConfigListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantCreditConfigListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditConfigListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditConfigListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantCreditConfigListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTypes

`func (o *UnibeeApiMerchantCreditConfigListReq) GetTypes() []int32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *UnibeeApiMerchantCreditConfigListReq) GetTypesOk() (*[]int32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *UnibeeApiMerchantCreditConfigListReq) SetTypes(v []int32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *UnibeeApiMerchantCreditConfigListReq) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


