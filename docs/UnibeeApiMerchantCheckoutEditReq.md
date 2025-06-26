# UnibeeApiMerchantCheckoutEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutId** | **int64** | checkout id | 
**Data** | Pointer to **map[string]map[string]interface{}** | checkout config data | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**StagingData** | Pointer to **map[string]map[string]interface{}** | checkout staging config data | [optional] 

## Methods

### NewUnibeeApiMerchantCheckoutEditReq

`func NewUnibeeApiMerchantCheckoutEditReq(checkoutId int64, ) *UnibeeApiMerchantCheckoutEditReq`

NewUnibeeApiMerchantCheckoutEditReq instantiates a new UnibeeApiMerchantCheckoutEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCheckoutEditReqWithDefaults

`func NewUnibeeApiMerchantCheckoutEditReqWithDefaults() *UnibeeApiMerchantCheckoutEditReq`

NewUnibeeApiMerchantCheckoutEditReqWithDefaults instantiates a new UnibeeApiMerchantCheckoutEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutId

`func (o *UnibeeApiMerchantCheckoutEditReq) GetCheckoutId() int64`

GetCheckoutId returns the CheckoutId field if non-nil, zero value otherwise.

### GetCheckoutIdOk

`func (o *UnibeeApiMerchantCheckoutEditReq) GetCheckoutIdOk() (*int64, bool)`

GetCheckoutIdOk returns a tuple with the CheckoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutId

`func (o *UnibeeApiMerchantCheckoutEditReq) SetCheckoutId(v int64)`

SetCheckoutId sets CheckoutId field to given value.


### GetData

`func (o *UnibeeApiMerchantCheckoutEditReq) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiMerchantCheckoutEditReq) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiMerchantCheckoutEditReq) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeApiMerchantCheckoutEditReq) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantCheckoutEditReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCheckoutEditReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCheckoutEditReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCheckoutEditReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCheckoutEditReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCheckoutEditReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCheckoutEditReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCheckoutEditReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStagingData

`func (o *UnibeeApiMerchantCheckoutEditReq) GetStagingData() map[string]map[string]interface{}`

GetStagingData returns the StagingData field if non-nil, zero value otherwise.

### GetStagingDataOk

`func (o *UnibeeApiMerchantCheckoutEditReq) GetStagingDataOk() (*map[string]map[string]interface{}, bool)`

GetStagingDataOk returns a tuple with the StagingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingData

`func (o *UnibeeApiMerchantCheckoutEditReq) SetStagingData(v map[string]map[string]interface{})`

SetStagingData sets StagingData field to given value.

### HasStagingData

`func (o *UnibeeApiMerchantCheckoutEditReq) HasStagingData() bool`

HasStagingData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


