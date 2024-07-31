# MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**Link** | Pointer to **string** | if automatic payment is false, Gateway Link will provided that manual payment needed | [optional] 
**Paid** | Pointer to **bool** | true|false,automatic payment is default behavior for one-time addon purchased, payment will create attach to the purchase, when payment is success, return false, otherwise false | [optional] 
**SubscriptionOnetimeAddon** | Pointer to [**UnibeeApiBeanSubscriptionOnetimeAddon**](UnibeeApiBeanSubscriptionOnetimeAddon.md) |  | [optional] 

## Methods

### NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData

`func NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData() *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData`

NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData instantiates a new MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseDataWithDefaults

`func NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseDataWithDefaults() *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData`

NewMerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetLink

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaid

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetSubscriptionOnetimeAddon

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetSubscriptionOnetimeAddon() UnibeeApiBeanSubscriptionOnetimeAddon`

GetSubscriptionOnetimeAddon returns the SubscriptionOnetimeAddon field if non-nil, zero value otherwise.

### GetSubscriptionOnetimeAddonOk

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) GetSubscriptionOnetimeAddonOk() (*UnibeeApiBeanSubscriptionOnetimeAddon, bool)`

GetSubscriptionOnetimeAddonOk returns a tuple with the SubscriptionOnetimeAddon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOnetimeAddon

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) SetSubscriptionOnetimeAddon(v UnibeeApiBeanSubscriptionOnetimeAddon)`

SetSubscriptionOnetimeAddon sets SubscriptionOnetimeAddon field to given value.

### HasSubscriptionOnetimeAddon

`func (o *MerchantSubscriptionNewOnetimeAddonPaymentPost200ResponseData) HasSubscriptionOnetimeAddon() bool`

HasSubscriptionOnetimeAddon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


