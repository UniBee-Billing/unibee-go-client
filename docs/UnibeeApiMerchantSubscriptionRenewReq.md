# UnibeeApiMerchantSubscriptionRenewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode, override subscription discount | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId, use subscription&#39;s gateway if not provide | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, override subscription taxPercentage if provide | [optional] 
**UserId** | **int64** | UserId | 

## Methods

### NewUnibeeApiMerchantSubscriptionRenewReq

`func NewUnibeeApiMerchantSubscriptionRenewReq(subscriptionId string, userId int64, ) *UnibeeApiMerchantSubscriptionRenewReq`

NewUnibeeApiMerchantSubscriptionRenewReq instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults() *UnibeeApiMerchantSubscriptionRenewReq`

NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscount() UnibeeApiBeanExternalDiscountParam`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


