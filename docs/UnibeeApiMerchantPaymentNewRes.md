# UnibeeApiMerchantPaymentNewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalPaymentId** | Pointer to **string** | The external unique id of payment | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**PaymentId** | Pointer to **string** | The unique id of payment | [optional] 
**Status** | Pointer to **int32** | Status, 10-Created|20-Success|30-Failed|40-Cancelled | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentNewRes

`func NewUnibeeApiMerchantPaymentNewRes() *UnibeeApiMerchantPaymentNewRes`

NewUnibeeApiMerchantPaymentNewRes instantiates a new UnibeeApiMerchantPaymentNewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentNewResWithDefaults

`func NewUnibeeApiMerchantPaymentNewResWithDefaults() *UnibeeApiMerchantPaymentNewRes`

NewUnibeeApiMerchantPaymentNewResWithDefaults instantiates a new UnibeeApiMerchantPaymentNewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UnibeeApiMerchantPaymentNewRes) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UnibeeApiMerchantPaymentNewRes) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UnibeeApiMerchantPaymentNewRes) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UnibeeApiMerchantPaymentNewRes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeApiMerchantPaymentNewRes) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiMerchantPaymentNewRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiMerchantPaymentNewRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiMerchantPaymentNewRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiMerchantPaymentNewRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentNewRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiMerchantPaymentNewRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantPaymentNewRes) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantPaymentNewRes) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantPaymentNewRes) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantPaymentNewRes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


