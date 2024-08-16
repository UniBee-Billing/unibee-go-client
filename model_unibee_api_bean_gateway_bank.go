/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiBeanGatewayBank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanGatewayBank{}

// UnibeeApiBeanGatewayBank struct for UnibeeApiBeanGatewayBank
type UnibeeApiBeanGatewayBank struct {
	// The AccountHolder of wire transfer 
	AccountHolder string `json:"accountHolder"`
	// The address of wire transfer 
	Address string `json:"address"`
	// The BIC of wire transfer 
	Bic string `json:"bic"`
	// The IBAN of wire transfer 
	Iban string `json:"iban"`
}

type _UnibeeApiBeanGatewayBank UnibeeApiBeanGatewayBank

// NewUnibeeApiBeanGatewayBank instantiates a new UnibeeApiBeanGatewayBank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanGatewayBank(accountHolder string, address string, bic string, iban string) *UnibeeApiBeanGatewayBank {
	this := UnibeeApiBeanGatewayBank{}
	this.AccountHolder = accountHolder
	this.Address = address
	this.Bic = bic
	this.Iban = iban
	return &this
}

// NewUnibeeApiBeanGatewayBankWithDefaults instantiates a new UnibeeApiBeanGatewayBank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanGatewayBankWithDefaults() *UnibeeApiBeanGatewayBank {
	this := UnibeeApiBeanGatewayBank{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value
func (o *UnibeeApiBeanGatewayBank) GetAccountHolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGatewayBank) GetAccountHolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolder, true
}

// SetAccountHolder sets field value
func (o *UnibeeApiBeanGatewayBank) SetAccountHolder(v string) {
	o.AccountHolder = v
}

// GetAddress returns the Address field value
func (o *UnibeeApiBeanGatewayBank) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGatewayBank) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UnibeeApiBeanGatewayBank) SetAddress(v string) {
	o.Address = v
}

// GetBic returns the Bic field value
func (o *UnibeeApiBeanGatewayBank) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGatewayBank) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *UnibeeApiBeanGatewayBank) SetBic(v string) {
	o.Bic = v
}

// GetIban returns the Iban field value
func (o *UnibeeApiBeanGatewayBank) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGatewayBank) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *UnibeeApiBeanGatewayBank) SetIban(v string) {
	o.Iban = v
}

func (o UnibeeApiBeanGatewayBank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanGatewayBank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolder"] = o.AccountHolder
	toSerialize["address"] = o.Address
	toSerialize["bic"] = o.Bic
	toSerialize["iban"] = o.Iban
	return toSerialize, nil
}

func (o *UnibeeApiBeanGatewayBank) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountHolder",
		"address",
		"bic",
		"iban",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiBeanGatewayBank := _UnibeeApiBeanGatewayBank{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiBeanGatewayBank)

	if err != nil {
		return err
	}

	*o = UnibeeApiBeanGatewayBank(varUnibeeApiBeanGatewayBank)

	return err
}

type NullableUnibeeApiBeanGatewayBank struct {
	value *UnibeeApiBeanGatewayBank
	isSet bool
}

func (v NullableUnibeeApiBeanGatewayBank) Get() *UnibeeApiBeanGatewayBank {
	return v.value
}

func (v *NullableUnibeeApiBeanGatewayBank) Set(val *UnibeeApiBeanGatewayBank) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanGatewayBank) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanGatewayBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanGatewayBank(val *UnibeeApiBeanGatewayBank) *NullableUnibeeApiBeanGatewayBank {
	return &NullableUnibeeApiBeanGatewayBank{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanGatewayBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanGatewayBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


