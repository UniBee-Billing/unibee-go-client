# UnibeeApiBeanSplitPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | Payment creation time (UTC timestamp) | [optional] 
**Currency** | Pointer to **string** | Payment currency | [optional] 
**GatewayId** | Pointer to **int64** | The gateway id | [optional] 
**GatewayName** | Pointer to **string** | The gateway name (e.g., stripe, paypal) | [optional] 
**Link** | Pointer to **string** | Payment link for checkout | [optional] 
**PaidTime** | Pointer to **int64** | Payment success time (UTC timestamp, 0 if not paid) | [optional] 
**PaymentId** | Pointer to **string** | The payment id | [optional] 
**SplitSequence** | Pointer to **int32** | Split payment sequence number (1, 2, 3...) | [optional] 
**Status** | Pointer to **int32** | Payment status (10-pending, 20-success, 30-failure, 40-cancel) | [optional] 
**TotalAmount** | Pointer to **int64** | Payment amount in cents | [optional] 

## Methods

### NewUnibeeApiBeanSplitPaymentInfo

`func NewUnibeeApiBeanSplitPaymentInfo() *UnibeeApiBeanSplitPaymentInfo`

NewUnibeeApiBeanSplitPaymentInfo instantiates a new UnibeeApiBeanSplitPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSplitPaymentInfoWithDefaults

`func NewUnibeeApiBeanSplitPaymentInfoWithDefaults() *UnibeeApiBeanSplitPaymentInfo`

NewUnibeeApiBeanSplitPaymentInfoWithDefaults instantiates a new UnibeeApiBeanSplitPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanSplitPaymentInfo) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanSplitPaymentInfo) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanSplitPaymentInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanSplitPaymentInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanSplitPaymentInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanSplitPaymentInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanSplitPaymentInfo) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanSplitPaymentInfo) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanSplitPaymentInfo) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanSplitPaymentInfo) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanSplitPaymentInfo) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanSplitPaymentInfo) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanSplitPaymentInfo) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanSplitPaymentInfo) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanSplitPaymentInfo) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiBeanSplitPaymentInfo) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiBeanSplitPaymentInfo) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiBeanSplitPaymentInfo) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanSplitPaymentInfo) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanSplitPaymentInfo) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanSplitPaymentInfo) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetSplitSequence

`func (o *UnibeeApiBeanSplitPaymentInfo) GetSplitSequence() int32`

GetSplitSequence returns the SplitSequence field if non-nil, zero value otherwise.

### GetSplitSequenceOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetSplitSequenceOk() (*int32, bool)`

GetSplitSequenceOk returns a tuple with the SplitSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitSequence

`func (o *UnibeeApiBeanSplitPaymentInfo) SetSplitSequence(v int32)`

SetSplitSequence sets SplitSequence field to given value.

### HasSplitSequence

`func (o *UnibeeApiBeanSplitPaymentInfo) HasSplitSequence() bool`

HasSplitSequence returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanSplitPaymentInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanSplitPaymentInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanSplitPaymentInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanSplitPaymentInfo) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanSplitPaymentInfo) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanSplitPaymentInfo) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanSplitPaymentInfo) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


