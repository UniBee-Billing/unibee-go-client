# UnibeeApiCheckoutSubscriptionCreateRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **map[string]interface{}** |  | [optional] 
**InvoiceId** | Pointer to **string** | The unique id of invoice | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**OtherPendingCryptoSubscription** | Pointer to [**UnibeeApiBeanDetailSubscriptionDetail**](UnibeeApiBeanDetailSubscriptionDetail.md) |  | [optional] 
**Paid** | Pointer to **bool** |  | [optional] 
**PaymentId** | Pointer to **string** | The unique id of payment | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**Token** | Pointer to **string** | token | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCreateRes

`func NewUnibeeApiCheckoutSubscriptionCreateRes() *UnibeeApiCheckoutSubscriptionCreateRes`

NewUnibeeApiCheckoutSubscriptionCreateRes instantiates a new UnibeeApiCheckoutSubscriptionCreateRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCreateResWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCreateResWithDefaults() *UnibeeApiCheckoutSubscriptionCreateRes`

NewUnibeeApiCheckoutSubscriptionCreateResWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCreateRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetOtherPendingCryptoSubscription() UnibeeApiBeanDetailSubscriptionDetail`

GetOtherPendingCryptoSubscription returns the OtherPendingCryptoSubscription field if non-nil, zero value otherwise.

### GetOtherPendingCryptoSubscriptionOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetOtherPendingCryptoSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool)`

GetOtherPendingCryptoSubscriptionOk returns a tuple with the OtherPendingCryptoSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetOtherPendingCryptoSubscription(v UnibeeApiBeanDetailSubscriptionDetail)`

SetOtherPendingCryptoSubscription sets OtherPendingCryptoSubscription field to given value.

### HasOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasOtherPendingCryptoSubscription() bool`

HasOtherPendingCryptoSubscription returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiCheckoutSubscriptionCreateRes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


