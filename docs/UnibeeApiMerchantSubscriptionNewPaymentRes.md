# UnibeeApiMerchantSubscriptionNewPaymentRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalPaymentId** | Pointer to **string** | The external unique id of payment | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**PaymentId** | Pointer to **string** | The unique id of payment | [optional] 
**Status** | Pointer to **int32** | Status, 10-Created|20-Success|30-Failed|40-Cancelled | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionNewPaymentRes

`func NewUnibeeApiMerchantSubscriptionNewPaymentRes() *UnibeeApiMerchantSubscriptionNewPaymentRes`

NewUnibeeApiMerchantSubscriptionNewPaymentRes instantiates a new UnibeeApiMerchantSubscriptionNewPaymentRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionNewPaymentResWithDefaults

`func NewUnibeeApiMerchantSubscriptionNewPaymentResWithDefaults() *UnibeeApiMerchantSubscriptionNewPaymentRes`

NewUnibeeApiMerchantSubscriptionNewPaymentResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionNewPaymentRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


