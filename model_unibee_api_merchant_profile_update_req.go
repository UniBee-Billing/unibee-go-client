/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProfileUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileUpdateReq{}

// UnibeeApiMerchantProfileUpdateReq struct for UnibeeApiMerchantProfileUpdateReq
type UnibeeApiMerchantProfileUpdateReq struct {
	// address
	Address *string `json:"address,omitempty"`
	// company_logo
	CompanyLogo *string `json:"companyLogo,omitempty"`
	// company_name
	CompanyName *string `json:"companyName,omitempty"`
	// email
	Email *string `json:"email,omitempty"`
	// User Portal Host
	Host *string `json:"host,omitempty"`
	// phone
	Phone *string `json:"phone,omitempty"`
	// User TimeZone
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewUnibeeApiMerchantProfileUpdateReq instantiates a new UnibeeApiMerchantProfileUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileUpdateReq() *UnibeeApiMerchantProfileUpdateReq {
	this := UnibeeApiMerchantProfileUpdateReq{}
	return &this
}

// NewUnibeeApiMerchantProfileUpdateReqWithDefaults instantiates a new UnibeeApiMerchantProfileUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileUpdateReqWithDefaults() *UnibeeApiMerchantProfileUpdateReq {
	this := UnibeeApiMerchantProfileUpdateReq{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetAddress(v string) {
	o.Address = &v
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetCompanyLogo() string {
	if o == nil || IsNil(o.CompanyLogo) {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetCompanyLogoOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyLogo) {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasCompanyLogo() bool {
	if o != nil && !IsNil(o.CompanyLogo) {
		return true
	}

	return false
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetEmail(v string) {
	o.Email = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetHost(v string) {
	o.Host = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetPhone(v string) {
	o.Phone = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileUpdateReq) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileUpdateReq) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *UnibeeApiMerchantProfileUpdateReq) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o UnibeeApiMerchantProfileUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.CompanyLogo) {
		toSerialize["companyLogo"] = o.CompanyLogo
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProfileUpdateReq struct {
	value *UnibeeApiMerchantProfileUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileUpdateReq) Get() *UnibeeApiMerchantProfileUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileUpdateReq) Set(val *UnibeeApiMerchantProfileUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileUpdateReq(val *UnibeeApiMerchantProfileUpdateReq) *NullableUnibeeApiMerchantProfileUpdateReq {
	return &NullableUnibeeApiMerchantProfileUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


