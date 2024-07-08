# UnibeeApiBeanVatCountryRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** | TaxId | [optional] 
**IsEU** | Pointer to **bool** |  | [optional] 
**StandardTaxPercentage** | Pointer to **int64** | Tax税率，万分位，1000 表示 10% | [optional] 
**VatSupport** | Pointer to **bool** | vat support,true or false | [optional] 

## Methods

### NewUnibeeApiBeanVatCountryRate

`func NewUnibeeApiBeanVatCountryRate() *UnibeeApiBeanVatCountryRate`

NewUnibeeApiBeanVatCountryRate instantiates a new UnibeeApiBeanVatCountryRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanVatCountryRateWithDefaults

`func NewUnibeeApiBeanVatCountryRateWithDefaults() *UnibeeApiBeanVatCountryRate`

NewUnibeeApiBeanVatCountryRateWithDefaults instantiates a new UnibeeApiBeanVatCountryRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiBeanVatCountryRate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanVatCountryRate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanVatCountryRate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanVatCountryRate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *UnibeeApiBeanVatCountryRate) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UnibeeApiBeanVatCountryRate) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UnibeeApiBeanVatCountryRate) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UnibeeApiBeanVatCountryRate) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanVatCountryRate) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanVatCountryRate) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanVatCountryRate) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanVatCountryRate) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanVatCountryRate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanVatCountryRate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanVatCountryRate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanVatCountryRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEU

`func (o *UnibeeApiBeanVatCountryRate) GetIsEU() bool`

GetIsEU returns the IsEU field if non-nil, zero value otherwise.

### GetIsEUOk

`func (o *UnibeeApiBeanVatCountryRate) GetIsEUOk() (*bool, bool)`

GetIsEUOk returns a tuple with the IsEU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEU

`func (o *UnibeeApiBeanVatCountryRate) SetIsEU(v bool)`

SetIsEU sets IsEU field to given value.

### HasIsEU

`func (o *UnibeeApiBeanVatCountryRate) HasIsEU() bool`

HasIsEU returns a boolean if a field has been set.

### GetStandardTaxPercentage

`func (o *UnibeeApiBeanVatCountryRate) GetStandardTaxPercentage() int64`

GetStandardTaxPercentage returns the StandardTaxPercentage field if non-nil, zero value otherwise.

### GetStandardTaxPercentageOk

`func (o *UnibeeApiBeanVatCountryRate) GetStandardTaxPercentageOk() (*int64, bool)`

GetStandardTaxPercentageOk returns a tuple with the StandardTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardTaxPercentage

`func (o *UnibeeApiBeanVatCountryRate) SetStandardTaxPercentage(v int64)`

SetStandardTaxPercentage sets StandardTaxPercentage field to given value.

### HasStandardTaxPercentage

`func (o *UnibeeApiBeanVatCountryRate) HasStandardTaxPercentage() bool`

HasStandardTaxPercentage returns a boolean if a field has been set.

### GetVatSupport

`func (o *UnibeeApiBeanVatCountryRate) GetVatSupport() bool`

GetVatSupport returns the VatSupport field if non-nil, zero value otherwise.

### GetVatSupportOk

`func (o *UnibeeApiBeanVatCountryRate) GetVatSupportOk() (*bool, bool)`

GetVatSupportOk returns a tuple with the VatSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatSupport

`func (o *UnibeeApiBeanVatCountryRate) SetVatSupport(v bool)`

SetVatSupport sets VatSupport field to given value.

### HasVatSupport

`func (o *UnibeeApiBeanVatCountryRate) HasVatSupport() bool`

HasVatSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


