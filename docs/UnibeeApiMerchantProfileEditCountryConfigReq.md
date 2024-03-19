# UnibeeApiMerchantProfileEditCountryConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | CountryCode | 
**Name** | Pointer to **string** | name | [optional] 
**VatEnable** | Pointer to **bool** | VatEnable, Default true | [optional] 

## Methods

### NewUnibeeApiMerchantProfileEditCountryConfigReq

`func NewUnibeeApiMerchantProfileEditCountryConfigReq(countryCode string, ) *UnibeeApiMerchantProfileEditCountryConfigReq`

NewUnibeeApiMerchantProfileEditCountryConfigReq instantiates a new UnibeeApiMerchantProfileEditCountryConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProfileEditCountryConfigReqWithDefaults

`func NewUnibeeApiMerchantProfileEditCountryConfigReqWithDefaults() *UnibeeApiMerchantProfileEditCountryConfigReq`

NewUnibeeApiMerchantProfileEditCountryConfigReqWithDefaults instantiates a new UnibeeApiMerchantProfileEditCountryConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVatEnable

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetVatEnable() bool`

GetVatEnable returns the VatEnable field if non-nil, zero value otherwise.

### GetVatEnableOk

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) GetVatEnableOk() (*bool, bool)`

GetVatEnableOk returns a tuple with the VatEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatEnable

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) SetVatEnable(v bool)`

SetVatEnable sets VatEnable field to given value.

### HasVatEnable

`func (o *UnibeeApiMerchantProfileEditCountryConfigReq) HasVatEnable() bool`

HasVatEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


