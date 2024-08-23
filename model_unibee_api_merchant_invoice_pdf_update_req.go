/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoicePdfUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoicePdfUpdateReq{}

// UnibeeApiMerchantInvoicePdfUpdateReq struct for UnibeeApiMerchantInvoicePdfUpdateReq
type UnibeeApiMerchantInvoicePdfUpdateReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// IssueAddress
	IssueAddress *string `json:"issueAddress,omitempty"`
	// IssueCompanyName
	IssueCompanyName *string `json:"issueCompanyName,omitempty"`
	// IssueRegNumber
	IssueRegNumber *string `json:"issueRegNumber,omitempty"`
	// IssueVatNumber
	IssueVatNumber *string `json:"issueVatNumber,omitempty"`
	// LocalizedCurrency, To display localized currency amount
	LocalizedCurrency *string `json:"localizedCurrency,omitempty"`
	// LocalizedExchangeRate, exchange rate must set while LocalizedCurrency enabled
	LocalizedExchangeRate *float32 `json:"localizedExchangeRate,omitempty"`
	// Whether sen invoice email to user or not，default false
	SendUserEmail *bool `json:"sendUserEmail,omitempty"`
	// ShowDetailItem, whether to display detail item information in pdf generate, unitAmount, quantity, etc.
	ShowDetailItem *bool `json:"showDetailItem,omitempty"`
	// Template
	Template *string `json:"template,omitempty"`
}

type _UnibeeApiMerchantInvoicePdfUpdateReq UnibeeApiMerchantInvoicePdfUpdateReq

// NewUnibeeApiMerchantInvoicePdfUpdateReq instantiates a new UnibeeApiMerchantInvoicePdfUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoicePdfUpdateReq(invoiceId string) *UnibeeApiMerchantInvoicePdfUpdateReq {
	this := UnibeeApiMerchantInvoicePdfUpdateReq{}
	this.InvoiceId = invoiceId
	var sendUserEmail bool = false
	this.SendUserEmail = &sendUserEmail
	var showDetailItem bool = false
	this.ShowDetailItem = &showDetailItem
	return &this
}

// NewUnibeeApiMerchantInvoicePdfUpdateReqWithDefaults instantiates a new UnibeeApiMerchantInvoicePdfUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoicePdfUpdateReqWithDefaults() *UnibeeApiMerchantInvoicePdfUpdateReq {
	this := UnibeeApiMerchantInvoicePdfUpdateReq{}
	var sendUserEmail bool = false
	this.SendUserEmail = &sendUserEmail
	var showDetailItem bool = false
	this.ShowDetailItem = &showDetailItem
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetIssueAddress returns the IssueAddress field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueAddress() string {
	if o == nil || IsNil(o.IssueAddress) {
		var ret string
		return ret
	}
	return *o.IssueAddress
}

// GetIssueAddressOk returns a tuple with the IssueAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IssueAddress) {
		return nil, false
	}
	return o.IssueAddress, true
}

// HasIssueAddress returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueAddress() bool {
	if o != nil && !IsNil(o.IssueAddress) {
		return true
	}

	return false
}

// SetIssueAddress gets a reference to the given string and assigns it to the IssueAddress field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueAddress(v string) {
	o.IssueAddress = &v
}

// GetIssueCompanyName returns the IssueCompanyName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueCompanyName() string {
	if o == nil || IsNil(o.IssueCompanyName) {
		var ret string
		return ret
	}
	return *o.IssueCompanyName
}

// GetIssueCompanyNameOk returns a tuple with the IssueCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.IssueCompanyName) {
		return nil, false
	}
	return o.IssueCompanyName, true
}

// HasIssueCompanyName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueCompanyName() bool {
	if o != nil && !IsNil(o.IssueCompanyName) {
		return true
	}

	return false
}

// SetIssueCompanyName gets a reference to the given string and assigns it to the IssueCompanyName field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueCompanyName(v string) {
	o.IssueCompanyName = &v
}

// GetIssueRegNumber returns the IssueRegNumber field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueRegNumber() string {
	if o == nil || IsNil(o.IssueRegNumber) {
		var ret string
		return ret
	}
	return *o.IssueRegNumber
}

// GetIssueRegNumberOk returns a tuple with the IssueRegNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueRegNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IssueRegNumber) {
		return nil, false
	}
	return o.IssueRegNumber, true
}

// HasIssueRegNumber returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueRegNumber() bool {
	if o != nil && !IsNil(o.IssueRegNumber) {
		return true
	}

	return false
}

// SetIssueRegNumber gets a reference to the given string and assigns it to the IssueRegNumber field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueRegNumber(v string) {
	o.IssueRegNumber = &v
}

// GetIssueVatNumber returns the IssueVatNumber field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueVatNumber() string {
	if o == nil || IsNil(o.IssueVatNumber) {
		var ret string
		return ret
	}
	return *o.IssueVatNumber
}

// GetIssueVatNumberOk returns a tuple with the IssueVatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetIssueVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IssueVatNumber) {
		return nil, false
	}
	return o.IssueVatNumber, true
}

// HasIssueVatNumber returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasIssueVatNumber() bool {
	if o != nil && !IsNil(o.IssueVatNumber) {
		return true
	}

	return false
}

// SetIssueVatNumber gets a reference to the given string and assigns it to the IssueVatNumber field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetIssueVatNumber(v string) {
	o.IssueVatNumber = &v
}

// GetLocalizedCurrency returns the LocalizedCurrency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedCurrency() string {
	if o == nil || IsNil(o.LocalizedCurrency) {
		var ret string
		return ret
	}
	return *o.LocalizedCurrency
}

// GetLocalizedCurrencyOk returns a tuple with the LocalizedCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.LocalizedCurrency) {
		return nil, false
	}
	return o.LocalizedCurrency, true
}

// HasLocalizedCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasLocalizedCurrency() bool {
	if o != nil && !IsNil(o.LocalizedCurrency) {
		return true
	}

	return false
}

// SetLocalizedCurrency gets a reference to the given string and assigns it to the LocalizedCurrency field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetLocalizedCurrency(v string) {
	o.LocalizedCurrency = &v
}

// GetLocalizedExchangeRate returns the LocalizedExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRate() float32 {
	if o == nil || IsNil(o.LocalizedExchangeRate) {
		var ret float32
		return ret
	}
	return *o.LocalizedExchangeRate
}

// GetLocalizedExchangeRateOk returns a tuple with the LocalizedExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetLocalizedExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.LocalizedExchangeRate) {
		return nil, false
	}
	return o.LocalizedExchangeRate, true
}

// HasLocalizedExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasLocalizedExchangeRate() bool {
	if o != nil && !IsNil(o.LocalizedExchangeRate) {
		return true
	}

	return false
}

// SetLocalizedExchangeRate gets a reference to the given float32 and assigns it to the LocalizedExchangeRate field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetLocalizedExchangeRate(v float32) {
	o.LocalizedExchangeRate = &v
}

// GetSendUserEmail returns the SendUserEmail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetSendUserEmail() bool {
	if o == nil || IsNil(o.SendUserEmail) {
		var ret bool
		return ret
	}
	return *o.SendUserEmail
}

// GetSendUserEmailOk returns a tuple with the SendUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetSendUserEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendUserEmail) {
		return nil, false
	}
	return o.SendUserEmail, true
}

// HasSendUserEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasSendUserEmail() bool {
	if o != nil && !IsNil(o.SendUserEmail) {
		return true
	}

	return false
}

// SetSendUserEmail gets a reference to the given bool and assigns it to the SendUserEmail field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetSendUserEmail(v bool) {
	o.SendUserEmail = &v
}

// GetShowDetailItem returns the ShowDetailItem field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetShowDetailItem() bool {
	if o == nil || IsNil(o.ShowDetailItem) {
		var ret bool
		return ret
	}
	return *o.ShowDetailItem
}

// GetShowDetailItemOk returns a tuple with the ShowDetailItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetShowDetailItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowDetailItem) {
		return nil, false
	}
	return o.ShowDetailItem, true
}

// HasShowDetailItem returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasShowDetailItem() bool {
	if o != nil && !IsNil(o.ShowDetailItem) {
		return true
	}

	return false
}

// SetShowDetailItem gets a reference to the given bool and assigns it to the ShowDetailItem field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetShowDetailItem(v bool) {
	o.ShowDetailItem = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *UnibeeApiMerchantInvoicePdfUpdateReq) SetTemplate(v string) {
	o.Template = &v
}

func (o UnibeeApiMerchantInvoicePdfUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoicePdfUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	if !IsNil(o.IssueAddress) {
		toSerialize["issueAddress"] = o.IssueAddress
	}
	if !IsNil(o.IssueCompanyName) {
		toSerialize["issueCompanyName"] = o.IssueCompanyName
	}
	if !IsNil(o.IssueRegNumber) {
		toSerialize["issueRegNumber"] = o.IssueRegNumber
	}
	if !IsNil(o.IssueVatNumber) {
		toSerialize["issueVatNumber"] = o.IssueVatNumber
	}
	if !IsNil(o.LocalizedCurrency) {
		toSerialize["localizedCurrency"] = o.LocalizedCurrency
	}
	if !IsNil(o.LocalizedExchangeRate) {
		toSerialize["localizedExchangeRate"] = o.LocalizedExchangeRate
	}
	if !IsNil(o.SendUserEmail) {
		toSerialize["sendUserEmail"] = o.SendUserEmail
	}
	if !IsNil(o.ShowDetailItem) {
		toSerialize["showDetailItem"] = o.ShowDetailItem
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoicePdfUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
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

	varUnibeeApiMerchantInvoicePdfUpdateReq := _UnibeeApiMerchantInvoicePdfUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoicePdfUpdateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoicePdfUpdateReq(varUnibeeApiMerchantInvoicePdfUpdateReq)

	return err
}

type NullableUnibeeApiMerchantInvoicePdfUpdateReq struct {
	value *UnibeeApiMerchantInvoicePdfUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoicePdfUpdateReq) Get() *UnibeeApiMerchantInvoicePdfUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoicePdfUpdateReq) Set(val *UnibeeApiMerchantInvoicePdfUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoicePdfUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoicePdfUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoicePdfUpdateReq(val *UnibeeApiMerchantInvoicePdfUpdateReq) *NullableUnibeeApiMerchantInvoicePdfUpdateReq {
	return &NullableUnibeeApiMerchantInvoicePdfUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoicePdfUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoicePdfUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


