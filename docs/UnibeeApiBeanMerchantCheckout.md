# UnibeeApiBeanMerchantCheckout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** | checkout config data | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Id** | Pointer to **int64** | ID | [optional] 
**IsDefault** | Pointer to **bool** | is default | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**StagingData** | Pointer to **map[string]map[string]interface{}** | checkout staging config data | [optional] 
**UpdateTime** | Pointer to **int64** | update utc time | [optional] 

## Methods

### NewUnibeeApiBeanMerchantCheckout

`func NewUnibeeApiBeanMerchantCheckout() *UnibeeApiBeanMerchantCheckout`

NewUnibeeApiBeanMerchantCheckout instantiates a new UnibeeApiBeanMerchantCheckout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantCheckoutWithDefaults

`func NewUnibeeApiBeanMerchantCheckoutWithDefaults() *UnibeeApiBeanMerchantCheckout`

NewUnibeeApiBeanMerchantCheckoutWithDefaults instantiates a new UnibeeApiBeanMerchantCheckout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantCheckout) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantCheckout) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantCheckout) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantCheckout) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetData

`func (o *UnibeeApiBeanMerchantCheckout) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiBeanMerchantCheckout) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiBeanMerchantCheckout) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeApiBeanMerchantCheckout) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanMerchantCheckout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanMerchantCheckout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanMerchantCheckout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanMerchantCheckout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantCheckout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantCheckout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantCheckout) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantCheckout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *UnibeeApiBeanMerchantCheckout) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiBeanMerchantCheckout) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiBeanMerchantCheckout) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiBeanMerchantCheckout) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantCheckout) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantCheckout) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantCheckout) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantCheckout) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantCheckout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantCheckout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantCheckout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantCheckout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStagingData

`func (o *UnibeeApiBeanMerchantCheckout) GetStagingData() map[string]map[string]interface{}`

GetStagingData returns the StagingData field if non-nil, zero value otherwise.

### GetStagingDataOk

`func (o *UnibeeApiBeanMerchantCheckout) GetStagingDataOk() (*map[string]map[string]interface{}, bool)`

GetStagingDataOk returns a tuple with the StagingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingData

`func (o *UnibeeApiBeanMerchantCheckout) SetStagingData(v map[string]map[string]interface{})`

SetStagingData sets StagingData field to given value.

### HasStagingData

`func (o *UnibeeApiBeanMerchantCheckout) HasStagingData() bool`

HasStagingData returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UnibeeApiBeanMerchantCheckout) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UnibeeApiBeanMerchantCheckout) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UnibeeApiBeanMerchantCheckout) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UnibeeApiBeanMerchantCheckout) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


