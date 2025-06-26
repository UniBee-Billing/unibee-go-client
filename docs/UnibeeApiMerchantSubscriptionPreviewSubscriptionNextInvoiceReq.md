# UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed if subscriptionId not specified | [optional] 
**ProductId** | Pointer to **int64** | default product will use if productId not specified and subscriptionId is blank | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq

`func NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq() *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq`

NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq instantiates a new UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReqWithDefaults() *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq`

NewUnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionPreviewSubscriptionNextInvoiceReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


