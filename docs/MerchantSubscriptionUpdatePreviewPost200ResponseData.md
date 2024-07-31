# MerchantSubscriptionUpdatePreviewPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**NextPeriodInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**ProrationDate** | Pointer to **int64** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 

## Methods

### NewMerchantSubscriptionUpdatePreviewPost200ResponseData

`func NewMerchantSubscriptionUpdatePreviewPost200ResponseData() *MerchantSubscriptionUpdatePreviewPost200ResponseData`

NewMerchantSubscriptionUpdatePreviewPost200ResponseData instantiates a new MerchantSubscriptionUpdatePreviewPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionUpdatePreviewPost200ResponseDataWithDefaults

`func NewMerchantSubscriptionUpdatePreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionUpdatePreviewPost200ResponseData`

NewMerchantSubscriptionUpdatePreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUpdatePreviewPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetNextPeriodInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetNextPeriodInvoice() UnibeeApiBeanInvoice`

GetNextPeriodInvoice returns the NextPeriodInvoice field if non-nil, zero value otherwise.

### GetNextPeriodInvoiceOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetNextPeriodInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetNextPeriodInvoiceOk returns a tuple with the NextPeriodInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPeriodInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetNextPeriodInvoice(v UnibeeApiBeanInvoice)`

SetNextPeriodInvoice sets NextPeriodInvoice field to given value.

### HasNextPeriodInvoice

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasNextPeriodInvoice() bool`

HasNextPeriodInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetProrationDate

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetTotalAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MerchantSubscriptionUpdatePreviewPost200ResponseData) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


