# UnibeeApiBeanDetailPaymentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**UnibeeApiBeanGateway**](UnibeeApiBeanGateway.md) |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) |  | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailPaymentDetail

`func NewUnibeeApiBeanDetailPaymentDetail() *UnibeeApiBeanDetailPaymentDetail`

NewUnibeeApiBeanDetailPaymentDetail instantiates a new UnibeeApiBeanDetailPaymentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailPaymentDetailWithDefaults

`func NewUnibeeApiBeanDetailPaymentDetailWithDefaults() *UnibeeApiBeanDetailPaymentDetail`

NewUnibeeApiBeanDetailPaymentDetailWithDefaults instantiates a new UnibeeApiBeanDetailPaymentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *UnibeeApiBeanDetailPaymentDetail) GetGateway() UnibeeApiBeanGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanDetailPaymentDetail) GetGatewayOk() (*UnibeeApiBeanGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanDetailPaymentDetail) SetGateway(v UnibeeApiBeanGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanDetailPaymentDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiBeanDetailPaymentDetail) GetInvoice() UnibeeApiBeanDetailInvoiceDetail`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiBeanDetailPaymentDetail) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiBeanDetailPaymentDetail) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiBeanDetailPaymentDetail) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanDetailPaymentDetail) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanDetailPaymentDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanDetailPaymentDetail) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanDetailPaymentDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailPaymentDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailPaymentDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailPaymentDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailPaymentDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


