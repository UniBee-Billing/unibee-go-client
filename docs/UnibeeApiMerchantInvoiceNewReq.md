# UnibeeApiMerchantInvoiceNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency | 
**Finish** | Pointer to **bool** |  | [optional] 
**GatewayId** | **int64** | Gateway Id | 
**Lines** | Pointer to [**[]UnibeeApiMerchantInvoiceNewInvoiceItemParam**](UnibeeApiMerchantInvoiceNewInvoiceItemParam.md) |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**TaxPercentage** | **int64** | TaxPercentage，1000 represent 10% | 
**UserId** | **int64** | UserId | 

## Methods

### NewUnibeeApiMerchantInvoiceNewReq

`func NewUnibeeApiMerchantInvoiceNewReq(currency string, gatewayId int64, taxPercentage int64, userId int64, ) *UnibeeApiMerchantInvoiceNewReq`

NewUnibeeApiMerchantInvoiceNewReq instantiates a new UnibeeApiMerchantInvoiceNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceNewReqWithDefaults

`func NewUnibeeApiMerchantInvoiceNewReqWithDefaults() *UnibeeApiMerchantInvoiceNewReq`

NewUnibeeApiMerchantInvoiceNewReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFinish

`func (o *UnibeeApiMerchantInvoiceNewReq) GetFinish() bool`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetFinishOk() (*bool, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *UnibeeApiMerchantInvoiceNewReq) SetFinish(v bool)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *UnibeeApiMerchantInvoiceNewReq) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceNewReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceNewReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetLines

`func (o *UnibeeApiMerchantInvoiceNewReq) GetLines() []UnibeeApiMerchantInvoiceNewInvoiceItemParam`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetLinesOk() (*[]UnibeeApiMerchantInvoiceNewInvoiceItemParam, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiMerchantInvoiceNewReq) SetLines(v []UnibeeApiMerchantInvoiceNewInvoiceItemParam)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiMerchantInvoiceNewReq) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantInvoiceNewReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantInvoiceNewReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantInvoiceNewReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantInvoiceNewReq) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantInvoiceNewReq) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.


### GetUserId

`func (o *UnibeeApiMerchantInvoiceNewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantInvoiceNewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantInvoiceNewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


