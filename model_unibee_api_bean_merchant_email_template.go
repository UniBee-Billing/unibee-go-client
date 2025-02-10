/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantEmailTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantEmailTemplate{}

// UnibeeApiBeanMerchantEmailTemplate struct for UnibeeApiBeanMerchantEmailTemplate
type UnibeeApiBeanMerchantEmailTemplate struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	GatewayTemplateId *string `json:"gatewayTemplateId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LanguageData *string `json:"languageData,omitempty"`
	MerchantId *int64 `json:"merchantId,omitempty"`
	Status *string `json:"status,omitempty"`
	TemplateAttachName *string `json:"templateAttachName,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	TemplateTitle *string `json:"templateTitle,omitempty"`
	// update utc time
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewUnibeeApiBeanMerchantEmailTemplate instantiates a new UnibeeApiBeanMerchantEmailTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantEmailTemplate() *UnibeeApiBeanMerchantEmailTemplate {
	this := UnibeeApiBeanMerchantEmailTemplate{}
	return &this
}

// NewUnibeeApiBeanMerchantEmailTemplateWithDefaults instantiates a new UnibeeApiBeanMerchantEmailTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantEmailTemplateWithDefaults() *UnibeeApiBeanMerchantEmailTemplate {
	this := UnibeeApiBeanMerchantEmailTemplate{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetGatewayTemplateId returns the GatewayTemplateId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetGatewayTemplateId() string {
	if o == nil || IsNil(o.GatewayTemplateId) {
		var ret string
		return ret
	}
	return *o.GatewayTemplateId
}

// GetGatewayTemplateIdOk returns a tuple with the GatewayTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetGatewayTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayTemplateId) {
		return nil, false
	}
	return o.GatewayTemplateId, true
}

// HasGatewayTemplateId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasGatewayTemplateId() bool {
	if o != nil && !IsNil(o.GatewayTemplateId) {
		return true
	}

	return false
}

// SetGatewayTemplateId gets a reference to the given string and assigns it to the GatewayTemplateId field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetGatewayTemplateId(v string) {
	o.GatewayTemplateId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetId(v int64) {
	o.Id = &v
}

// GetLanguageData returns the LanguageData field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetLanguageData() string {
	if o == nil || IsNil(o.LanguageData) {
		var ret string
		return ret
	}
	return *o.LanguageData
}

// GetLanguageDataOk returns a tuple with the LanguageData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetLanguageDataOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageData) {
		return nil, false
	}
	return o.LanguageData, true
}

// HasLanguageData returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasLanguageData() bool {
	if o != nil && !IsNil(o.LanguageData) {
		return true
	}

	return false
}

// SetLanguageData gets a reference to the given string and assigns it to the LanguageData field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetLanguageData(v string) {
	o.LanguageData = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetStatus(v string) {
	o.Status = &v
}

// GetTemplateAttachName returns the TemplateAttachName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateAttachName() string {
	if o == nil || IsNil(o.TemplateAttachName) {
		var ret string
		return ret
	}
	return *o.TemplateAttachName
}

// GetTemplateAttachNameOk returns a tuple with the TemplateAttachName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateAttachNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateAttachName) {
		return nil, false
	}
	return o.TemplateAttachName, true
}

// HasTemplateAttachName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasTemplateAttachName() bool {
	if o != nil && !IsNil(o.TemplateAttachName) {
		return true
	}

	return false
}

// SetTemplateAttachName gets a reference to the given string and assigns it to the TemplateAttachName field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetTemplateAttachName(v string) {
	o.TemplateAttachName = &v
}

// GetTemplateContent returns the TemplateContent field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateContent() string {
	if o == nil || IsNil(o.TemplateContent) {
		var ret string
		return ret
	}
	return *o.TemplateContent
}

// GetTemplateContentOk returns a tuple with the TemplateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateContentOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateContent) {
		return nil, false
	}
	return o.TemplateContent, true
}

// HasTemplateContent returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasTemplateContent() bool {
	if o != nil && !IsNil(o.TemplateContent) {
		return true
	}

	return false
}

// SetTemplateContent gets a reference to the given string and assigns it to the TemplateContent field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetTemplateContent(v string) {
	o.TemplateContent = &v
}

// GetTemplateDescription returns the TemplateDescription field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateDescription() string {
	if o == nil || IsNil(o.TemplateDescription) {
		var ret string
		return ret
	}
	return *o.TemplateDescription
}

// GetTemplateDescriptionOk returns a tuple with the TemplateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateDescription) {
		return nil, false
	}
	return o.TemplateDescription, true
}

// HasTemplateDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasTemplateDescription() bool {
	if o != nil && !IsNil(o.TemplateDescription) {
		return true
	}

	return false
}

// SetTemplateDescription gets a reference to the given string and assigns it to the TemplateDescription field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetTemplateDescription(v string) {
	o.TemplateDescription = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTemplateTitle returns the TemplateTitle field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateTitle() string {
	if o == nil || IsNil(o.TemplateTitle) {
		var ret string
		return ret
	}
	return *o.TemplateTitle
}

// GetTemplateTitleOk returns a tuple with the TemplateTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetTemplateTitleOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateTitle) {
		return nil, false
	}
	return o.TemplateTitle, true
}

// HasTemplateTitle returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasTemplateTitle() bool {
	if o != nil && !IsNil(o.TemplateTitle) {
		return true
	}

	return false
}

// SetTemplateTitle gets a reference to the given string and assigns it to the TemplateTitle field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetTemplateTitle(v string) {
	o.TemplateTitle = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplate) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UnibeeApiBeanMerchantEmailTemplate) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o UnibeeApiBeanMerchantEmailTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantEmailTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.GatewayTemplateId) {
		toSerialize["gatewayTemplateId"] = o.GatewayTemplateId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LanguageData) {
		toSerialize["languageData"] = o.LanguageData
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TemplateAttachName) {
		toSerialize["templateAttachName"] = o.TemplateAttachName
	}
	if !IsNil(o.TemplateContent) {
		toSerialize["templateContent"] = o.TemplateContent
	}
	if !IsNil(o.TemplateDescription) {
		toSerialize["templateDescription"] = o.TemplateDescription
	}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	if !IsNil(o.TemplateTitle) {
		toSerialize["templateTitle"] = o.TemplateTitle
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantEmailTemplate struct {
	value *UnibeeApiBeanMerchantEmailTemplate
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantEmailTemplate) Get() *UnibeeApiBeanMerchantEmailTemplate {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplate) Set(val *UnibeeApiBeanMerchantEmailTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantEmailTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantEmailTemplate(val *UnibeeApiBeanMerchantEmailTemplate) *NullableUnibeeApiBeanMerchantEmailTemplate {
	return &NullableUnibeeApiBeanMerchantEmailTemplate{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantEmailTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


