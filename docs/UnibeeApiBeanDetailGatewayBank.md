# UnibeeApiBeanDetailGatewayBank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ABARoutingNumber** | Pointer to **string** | The ABARoutingNumber of wire transfer  | [optional] 
**CNAPS** | Pointer to **string** | The CNAPS of wire transfer  | [optional] 
**Remarks** | Pointer to **string** | The Remarks additional content  | [optional] 
**AccountHolder** | **string** | The AccountHolder of wire transfer  | 
**AccountNumber** | Pointer to **string** | The Account Number | [optional] 
**Address** | **string** | The address of wire transfer  | 
**BankName** | Pointer to **string** | The Bank Name | [optional] 
**Bic** | Pointer to **string** | The BIC of wire transfer  | [optional] 
**BsbCode** | Pointer to **string** | The BSB Code | [optional] 
**Iban** | Pointer to **string** | The IBAN of wire transfer  | [optional] 
**SwiftCode** | Pointer to **string** | The Swift Code | [optional] 

## Methods

### NewUnibeeApiBeanDetailGatewayBank

`func NewUnibeeApiBeanDetailGatewayBank(accountHolder string, address string, ) *UnibeeApiBeanDetailGatewayBank`

NewUnibeeApiBeanDetailGatewayBank instantiates a new UnibeeApiBeanDetailGatewayBank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailGatewayBankWithDefaults

`func NewUnibeeApiBeanDetailGatewayBankWithDefaults() *UnibeeApiBeanDetailGatewayBank`

NewUnibeeApiBeanDetailGatewayBankWithDefaults instantiates a new UnibeeApiBeanDetailGatewayBank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetABARoutingNumber

`func (o *UnibeeApiBeanDetailGatewayBank) GetABARoutingNumber() string`

GetABARoutingNumber returns the ABARoutingNumber field if non-nil, zero value otherwise.

### GetABARoutingNumberOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetABARoutingNumberOk() (*string, bool)`

GetABARoutingNumberOk returns a tuple with the ABARoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetABARoutingNumber

`func (o *UnibeeApiBeanDetailGatewayBank) SetABARoutingNumber(v string)`

SetABARoutingNumber sets ABARoutingNumber field to given value.

### HasABARoutingNumber

`func (o *UnibeeApiBeanDetailGatewayBank) HasABARoutingNumber() bool`

HasABARoutingNumber returns a boolean if a field has been set.

### GetCNAPS

`func (o *UnibeeApiBeanDetailGatewayBank) GetCNAPS() string`

GetCNAPS returns the CNAPS field if non-nil, zero value otherwise.

### GetCNAPSOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetCNAPSOk() (*string, bool)`

GetCNAPSOk returns a tuple with the CNAPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCNAPS

`func (o *UnibeeApiBeanDetailGatewayBank) SetCNAPS(v string)`

SetCNAPS sets CNAPS field to given value.

### HasCNAPS

`func (o *UnibeeApiBeanDetailGatewayBank) HasCNAPS() bool`

HasCNAPS returns a boolean if a field has been set.

### GetRemarks

`func (o *UnibeeApiBeanDetailGatewayBank) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *UnibeeApiBeanDetailGatewayBank) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *UnibeeApiBeanDetailGatewayBank) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetAccountHolder

`func (o *UnibeeApiBeanDetailGatewayBank) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *UnibeeApiBeanDetailGatewayBank) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.


### GetAccountNumber

`func (o *UnibeeApiBeanDetailGatewayBank) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UnibeeApiBeanDetailGatewayBank) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UnibeeApiBeanDetailGatewayBank) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAddress

`func (o *UnibeeApiBeanDetailGatewayBank) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiBeanDetailGatewayBank) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBankName

`func (o *UnibeeApiBeanDetailGatewayBank) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *UnibeeApiBeanDetailGatewayBank) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *UnibeeApiBeanDetailGatewayBank) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBic

`func (o *UnibeeApiBeanDetailGatewayBank) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *UnibeeApiBeanDetailGatewayBank) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *UnibeeApiBeanDetailGatewayBank) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetBsbCode

`func (o *UnibeeApiBeanDetailGatewayBank) GetBsbCode() string`

GetBsbCode returns the BsbCode field if non-nil, zero value otherwise.

### GetBsbCodeOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetBsbCodeOk() (*string, bool)`

GetBsbCodeOk returns a tuple with the BsbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsbCode

`func (o *UnibeeApiBeanDetailGatewayBank) SetBsbCode(v string)`

SetBsbCode sets BsbCode field to given value.

### HasBsbCode

`func (o *UnibeeApiBeanDetailGatewayBank) HasBsbCode() bool`

HasBsbCode returns a boolean if a field has been set.

### GetIban

`func (o *UnibeeApiBeanDetailGatewayBank) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *UnibeeApiBeanDetailGatewayBank) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *UnibeeApiBeanDetailGatewayBank) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwiftCode

`func (o *UnibeeApiBeanDetailGatewayBank) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *UnibeeApiBeanDetailGatewayBank) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *UnibeeApiBeanDetailGatewayBank) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *UnibeeApiBeanDetailGatewayBank) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


