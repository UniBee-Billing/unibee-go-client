# UnibeeApiBeanUSVATGlobalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AvailableGatewayNames** | Pointer to **[]string** |  | [optional] 
**CheckAddressViaGateway** | Pointer to **bool** |  | [optional] 
**DefaultGatewayName** | Pointer to **string** |  | [optional] 
**FromAddress** | Pointer to [**UnibeeApiBeanUSAddress**](UnibeeApiBeanUSAddress.md) |  | [optional] 
**NexusAddresses** | Pointer to [**[]UnibeeApiBeanUSAddress**](UnibeeApiBeanUSAddress.md) |  | [optional] 
**TaxCode** | Pointer to **string** |  | [optional] 
**ToAddress** | Pointer to [**UnibeeApiBeanUSAddress**](UnibeeApiBeanUSAddress.md) |  | [optional] 
**UploadInvoiceToGateway** | Pointer to **bool** |  | [optional] 

## Methods

### NewUnibeeApiBeanUSVATGlobalConfig

`func NewUnibeeApiBeanUSVATGlobalConfig() *UnibeeApiBeanUSVATGlobalConfig`

NewUnibeeApiBeanUSVATGlobalConfig instantiates a new UnibeeApiBeanUSVATGlobalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanUSVATGlobalConfigWithDefaults

`func NewUnibeeApiBeanUSVATGlobalConfigWithDefaults() *UnibeeApiBeanUSVATGlobalConfig`

NewUnibeeApiBeanUSVATGlobalConfigWithDefaults instantiates a new UnibeeApiBeanUSVATGlobalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvailableGatewayNames

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetAvailableGatewayNames() []string`

GetAvailableGatewayNames returns the AvailableGatewayNames field if non-nil, zero value otherwise.

### GetAvailableGatewayNamesOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetAvailableGatewayNamesOk() (*[]string, bool)`

GetAvailableGatewayNamesOk returns a tuple with the AvailableGatewayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableGatewayNames

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetAvailableGatewayNames(v []string)`

SetAvailableGatewayNames sets AvailableGatewayNames field to given value.

### HasAvailableGatewayNames

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasAvailableGatewayNames() bool`

HasAvailableGatewayNames returns a boolean if a field has been set.

### GetCheckAddressViaGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetCheckAddressViaGateway() bool`

GetCheckAddressViaGateway returns the CheckAddressViaGateway field if non-nil, zero value otherwise.

### GetCheckAddressViaGatewayOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetCheckAddressViaGatewayOk() (*bool, bool)`

GetCheckAddressViaGatewayOk returns a tuple with the CheckAddressViaGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAddressViaGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetCheckAddressViaGateway(v bool)`

SetCheckAddressViaGateway sets CheckAddressViaGateway field to given value.

### HasCheckAddressViaGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasCheckAddressViaGateway() bool`

HasCheckAddressViaGateway returns a boolean if a field has been set.

### GetDefaultGatewayName

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetDefaultGatewayName() string`

GetDefaultGatewayName returns the DefaultGatewayName field if non-nil, zero value otherwise.

### GetDefaultGatewayNameOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetDefaultGatewayNameOk() (*string, bool)`

GetDefaultGatewayNameOk returns a tuple with the DefaultGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayName

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetDefaultGatewayName(v string)`

SetDefaultGatewayName sets DefaultGatewayName field to given value.

### HasDefaultGatewayName

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasDefaultGatewayName() bool`

HasDefaultGatewayName returns a boolean if a field has been set.

### GetFromAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetFromAddress() UnibeeApiBeanUSAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetFromAddressOk() (*UnibeeApiBeanUSAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetFromAddress(v UnibeeApiBeanUSAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetNexusAddresses

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetNexusAddresses() []UnibeeApiBeanUSAddress`

GetNexusAddresses returns the NexusAddresses field if non-nil, zero value otherwise.

### GetNexusAddressesOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetNexusAddressesOk() (*[]UnibeeApiBeanUSAddress, bool)`

GetNexusAddressesOk returns a tuple with the NexusAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusAddresses

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetNexusAddresses(v []UnibeeApiBeanUSAddress)`

SetNexusAddresses sets NexusAddresses field to given value.

### HasNexusAddresses

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasNexusAddresses() bool`

HasNexusAddresses returns a boolean if a field has been set.

### GetTaxCode

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetToAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetToAddress() UnibeeApiBeanUSAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetToAddressOk() (*UnibeeApiBeanUSAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetToAddress(v UnibeeApiBeanUSAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetUploadInvoiceToGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetUploadInvoiceToGateway() bool`

GetUploadInvoiceToGateway returns the UploadInvoiceToGateway field if non-nil, zero value otherwise.

### GetUploadInvoiceToGatewayOk

`func (o *UnibeeApiBeanUSVATGlobalConfig) GetUploadInvoiceToGatewayOk() (*bool, bool)`

GetUploadInvoiceToGatewayOk returns a tuple with the UploadInvoiceToGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadInvoiceToGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) SetUploadInvoiceToGateway(v bool)`

SetUploadInvoiceToGateway sets UploadInvoiceToGateway field to given value.

### HasUploadInvoiceToGateway

`func (o *UnibeeApiBeanUSVATGlobalConfig) HasUploadInvoiceToGateway() bool`

HasUploadInvoiceToGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


