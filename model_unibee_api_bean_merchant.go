/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchant{}

// UnibeeApiBeanMerchant struct for UnibeeApiBeanMerchant
type UnibeeApiBeanMerchant struct {
	// address
	Address *string `json:"address,omitempty"`
	// business_num
	BusinessNum *string `json:"businessNum,omitempty"`
	// company_logo
	CompanyLogo *string `json:"companyLogo,omitempty"`
	// company_name
	CompanyName *string `json:"companyName,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// email
	Email *string `json:"email,omitempty"`
	HomeUrl *string `json:"homeUrl,omitempty"`
	// merchant user portal host
	Host *string `json:"host,omitempty"`
	// merchant_id
	Id *int64 `json:"id,omitempty"`
	// idcard
	Idcard *string `json:"idcard,omitempty"`
	// location
	Location *string `json:"location,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// phone
	Phone *string `json:"phone,omitempty"`
	// merchant default time zone
	TimeZone *string `json:"timeZone,omitempty"`
	// type
	Type *int32 `json:"type,omitempty"`
	// create_user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanMerchant instantiates a new UnibeeApiBeanMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchant() *UnibeeApiBeanMerchant {
	this := UnibeeApiBeanMerchant{}
	return &this
}

// NewUnibeeApiBeanMerchantWithDefaults instantiates a new UnibeeApiBeanMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantWithDefaults() *UnibeeApiBeanMerchant {
	this := UnibeeApiBeanMerchant{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UnibeeApiBeanMerchant) SetAddress(v string) {
	o.Address = &v
}

// GetBusinessNum returns the BusinessNum field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetBusinessNum() string {
	if o == nil || IsNil(o.BusinessNum) {
		var ret string
		return ret
	}
	return *o.BusinessNum
}

// GetBusinessNumOk returns a tuple with the BusinessNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetBusinessNumOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessNum) {
		return nil, false
	}
	return o.BusinessNum, true
}

// HasBusinessNum returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasBusinessNum() bool {
	if o != nil && !IsNil(o.BusinessNum) {
		return true
	}

	return false
}

// SetBusinessNum gets a reference to the given string and assigns it to the BusinessNum field.
func (o *UnibeeApiBeanMerchant) SetBusinessNum(v string) {
	o.BusinessNum = &v
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetCompanyLogo() string {
	if o == nil || IsNil(o.CompanyLogo) {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetCompanyLogoOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyLogo) {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasCompanyLogo() bool {
	if o != nil && !IsNil(o.CompanyLogo) {
		return true
	}

	return false
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *UnibeeApiBeanMerchant) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UnibeeApiBeanMerchant) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchant) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiBeanMerchant) SetEmail(v string) {
	o.Email = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiBeanMerchant) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UnibeeApiBeanMerchant) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchant) SetId(v int64) {
	o.Id = &v
}

// GetIdcard returns the Idcard field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetIdcard() string {
	if o == nil || IsNil(o.Idcard) {
		var ret string
		return ret
	}
	return *o.Idcard
}

// GetIdcardOk returns a tuple with the Idcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetIdcardOk() (*string, bool) {
	if o == nil || IsNil(o.Idcard) {
		return nil, false
	}
	return o.Idcard, true
}

// HasIdcard returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasIdcard() bool {
	if o != nil && !IsNil(o.Idcard) {
		return true
	}

	return false
}

// SetIdcard gets a reference to the given string and assigns it to the Idcard field.
func (o *UnibeeApiBeanMerchant) SetIdcard(v string) {
	o.Idcard = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UnibeeApiBeanMerchant) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanMerchant) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeApiBeanMerchant) SetPhone(v string) {
	o.Phone = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *UnibeeApiBeanMerchant) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanMerchant) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchant) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchant) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchant) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanMerchant) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanMerchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BusinessNum) {
		toSerialize["businessNum"] = o.BusinessNum
	}
	if !IsNil(o.CompanyLogo) {
		toSerialize["companyLogo"] = o.CompanyLogo
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Idcard) {
		toSerialize["idcard"] = o.Idcard
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchant struct {
	value *UnibeeApiBeanMerchant
	isSet bool
}

func (v NullableUnibeeApiBeanMerchant) Get() *UnibeeApiBeanMerchant {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchant) Set(val *UnibeeApiBeanMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchant(val *UnibeeApiBeanMerchant) *NullableUnibeeApiBeanMerchant {
	return &NullableUnibeeApiBeanMerchant{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


