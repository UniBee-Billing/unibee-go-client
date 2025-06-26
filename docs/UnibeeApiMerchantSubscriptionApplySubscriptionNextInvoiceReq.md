# UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCreditAmount** | Pointer to **int64** | apply promo credit amount | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed if subscriptionId not specified | [optional] 
**ProductId** | Pointer to **int64** | default product will use if productId not specified and subscriptionId is blank | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq

`func NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq() *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq`

NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq instantiates a new UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReqWithDefaults() *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq`

NewUnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetApplyPromoCreditAmount() int64`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetApplyPromoCreditAmountOk() (*int64, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetApplyPromoCreditAmount(v int64)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionApplySubscriptionNextInvoiceReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


