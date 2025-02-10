# UnibeeApiMerchantCreditNewCreditRechargeRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAccount** | Pointer to [**UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) |  | [optional] 
**CreditRecharge** | Pointer to [**UnibeeApiBeanCreditRecharge**](UnibeeApiBeanCreditRecharge.md) |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchant**](UnibeeApiBeanMerchant.md) |  | [optional] 
**Paid** | Pointer to **bool** | Paid，true|false | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantCreditNewCreditRechargeRes

`func NewUnibeeApiMerchantCreditNewCreditRechargeRes() *UnibeeApiMerchantCreditNewCreditRechargeRes`

NewUnibeeApiMerchantCreditNewCreditRechargeRes instantiates a new UnibeeApiMerchantCreditNewCreditRechargeRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditNewCreditRechargeResWithDefaults

`func NewUnibeeApiMerchantCreditNewCreditRechargeResWithDefaults() *UnibeeApiMerchantCreditNewCreditRechargeRes`

NewUnibeeApiMerchantCreditNewCreditRechargeResWithDefaults instantiates a new UnibeeApiMerchantCreditNewCreditRechargeRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAccount

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditAccount() UnibeeApiBeanCreditAccount`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetCreditAccount(v UnibeeApiBeanCreditAccount)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetCreditRecharge

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditRecharge() UnibeeApiBeanCreditRecharge`

GetCreditRecharge returns the CreditRecharge field if non-nil, zero value otherwise.

### GetCreditRechargeOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetCreditRechargeOk() (*UnibeeApiBeanCreditRecharge, bool)`

GetCreditRechargeOk returns a tuple with the CreditRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRecharge

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetCreditRecharge(v UnibeeApiBeanCreditRecharge)`

SetCreditRecharge sets CreditRecharge field to given value.

### HasCreditRecharge

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasCreditRecharge() bool`

HasCreditRecharge returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantCreditNewCreditRechargeRes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


