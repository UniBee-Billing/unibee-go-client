# UnibeeApiMerchantSubscriptionOnetimeAddonListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**SubscriptionId** | **string** | SubscriptionId, id of subscription which one-time addon purchase history attached | 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonListReq

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonListReq(subscriptionId string, ) *UnibeeApiMerchantSubscriptionOnetimeAddonListReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonListReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonListReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonListReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonListReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


