# UnibeeApiMerchantSubscriptionOnetimeAddonNewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**UnibeeApiBeanInvoiceSimplify**](UnibeeApiBeanInvoiceSimplify.md) |  | [optional] 
**Link** | Pointer to **string** | if automatic payment is false, Gateway Link will provided that manual payment needed | [optional] 
**Paid** | Pointer to **bool** | true|false,automatic payment is default behavior for one-time addon purchased, payment will create attach to the purchase, when payment is success, return false, otherwise false | [optional] 
**SubscriptionOnetimeAddon** | Pointer to [**UnibeeApiBeanSubscriptionOnetimeAddonSimplify**](UnibeeApiBeanSubscriptionOnetimeAddonSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonNewRes

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewRes() *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes`

NewUnibeeApiMerchantSubscriptionOnetimeAddonNewRes instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonNewResWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewResWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes`

NewUnibeeApiMerchantSubscriptionOnetimeAddonNewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetInvoice() UnibeeApiBeanInvoiceSimplify`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetInvoice(v UnibeeApiBeanInvoiceSimplify)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetSubscriptionOnetimeAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetSubscriptionOnetimeAddon() UnibeeApiBeanSubscriptionOnetimeAddonSimplify`

GetSubscriptionOnetimeAddon returns the SubscriptionOnetimeAddon field if non-nil, zero value otherwise.

### GetSubscriptionOnetimeAddonOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) GetSubscriptionOnetimeAddonOk() (*UnibeeApiBeanSubscriptionOnetimeAddonSimplify, bool)`

GetSubscriptionOnetimeAddonOk returns a tuple with the SubscriptionOnetimeAddon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOnetimeAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) SetSubscriptionOnetimeAddon(v UnibeeApiBeanSubscriptionOnetimeAddonSimplify)`

SetSubscriptionOnetimeAddon sets SubscriptionOnetimeAddon field to given value.

### HasSubscriptionOnetimeAddon

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewRes) HasSubscriptionOnetimeAddon() bool`

HasSubscriptionOnetimeAddon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


