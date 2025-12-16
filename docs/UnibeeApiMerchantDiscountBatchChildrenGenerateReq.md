# UnibeeApiMerchantDiscountBatchChildrenGenerateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLength** | Pointer to **int32** | Optional. Total length of generated child code (prefix + random part). Must be &gt; len(templateCode) and &lt;&#x3D; 190. The random part length must be sufficient to support the template&#39;s target quantity. | [optional] 
**IsAsync** | Pointer to **bool** | Force async execution. Default is false (synchronous). If quantity &gt; 1000, this must be set to true. | [optional] 
**Quantity** | Pointer to **int32** | Number of child codes to generate. If not provided, generates up to the template&#39;s quantity limit minus existing count. Must not exceed the remaining available quantity. | [optional] 
**RandomStyle** | Pointer to **string** | Optional. Random part style: auto|numeric|lower|upper. Default is auto, which infers style from templateCode: all digits -&gt; numeric, all lowercase (with optional digits) -&gt; lower, otherwise upper. | [optional] 
**TemplateCode** | Pointer to **string** | The template&#39;s code prefix (alternative to templateId) | [optional] 
**TemplateId** | Pointer to **int32** | The template&#39;s Id (alternative to templateCode) | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountBatchChildrenGenerateReq

`func NewUnibeeApiMerchantDiscountBatchChildrenGenerateReq() *UnibeeApiMerchantDiscountBatchChildrenGenerateReq`

NewUnibeeApiMerchantDiscountBatchChildrenGenerateReq instantiates a new UnibeeApiMerchantDiscountBatchChildrenGenerateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountBatchChildrenGenerateReqWithDefaults

`func NewUnibeeApiMerchantDiscountBatchChildrenGenerateReqWithDefaults() *UnibeeApiMerchantDiscountBatchChildrenGenerateReq`

NewUnibeeApiMerchantDiscountBatchChildrenGenerateReqWithDefaults instantiates a new UnibeeApiMerchantDiscountBatchChildrenGenerateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeLength

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.

### HasCodeLength

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasCodeLength() bool`

HasCodeLength returns a boolean if a field has been set.

### GetIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetIsAsync() bool`

GetIsAsync returns the IsAsync field if non-nil, zero value otherwise.

### GetIsAsyncOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetIsAsyncOk() (*bool, bool)`

GetIsAsyncOk returns a tuple with the IsAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetIsAsync(v bool)`

SetIsAsync sets IsAsync field to given value.

### HasIsAsync

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasIsAsync() bool`

HasIsAsync returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRandomStyle

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetRandomStyle() string`

GetRandomStyle returns the RandomStyle field if non-nil, zero value otherwise.

### GetRandomStyleOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetRandomStyleOk() (*string, bool)`

GetRandomStyleOk returns a tuple with the RandomStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomStyle

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetRandomStyle(v string)`

SetRandomStyle sets RandomStyle field to given value.

### HasRandomStyle

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasRandomStyle() bool`

HasRandomStyle returns a boolean if a field has been set.

### GetTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UnibeeApiMerchantDiscountBatchChildrenGenerateReq) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


