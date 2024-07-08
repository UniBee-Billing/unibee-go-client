/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantWebhookResendPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantWebhookResendPost200Response{}

// MerchantWebhookResendPost200Response struct for MerchantWebhookResendPost200Response
type MerchantWebhookResendPost200Response struct {
	Code *int32 `json:"code,omitempty"`
	Data *MerchantWebhookResendPost200ResponseData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Redirect *string `json:"redirect,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}

// NewMerchantWebhookResendPost200Response instantiates a new MerchantWebhookResendPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantWebhookResendPost200Response() *MerchantWebhookResendPost200Response {
	this := MerchantWebhookResendPost200Response{}
	return &this
}

// NewMerchantWebhookResendPost200ResponseWithDefaults instantiates a new MerchantWebhookResendPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWebhookResendPost200ResponseWithDefaults() *MerchantWebhookResendPost200Response {
	this := MerchantWebhookResendPost200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200Response) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200Response) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *MerchantWebhookResendPost200Response) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200Response) GetData() MerchantWebhookResendPost200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret MerchantWebhookResendPost200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200Response) GetDataOk() (*MerchantWebhookResendPost200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MerchantWebhookResendPost200ResponseData and assigns it to the Data field.
func (o *MerchantWebhookResendPost200Response) SetData(v MerchantWebhookResendPost200ResponseData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MerchantWebhookResendPost200Response) SetMessage(v string) {
	o.Message = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200Response) GetRedirect() string {
	if o == nil || IsNil(o.Redirect) {
		var ret string
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200Response) GetRedirectOk() (*string, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200Response) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given string and assigns it to the Redirect field.
func (o *MerchantWebhookResendPost200Response) SetRedirect(v string) {
	o.Redirect = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200Response) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200Response) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *MerchantWebhookResendPost200Response) SetRequestId(v string) {
	o.RequestId = &v
}

func (o MerchantWebhookResendPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantWebhookResendPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Redirect) {
		toSerialize["redirect"] = o.Redirect
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableMerchantWebhookResendPost200Response struct {
	value *MerchantWebhookResendPost200Response
	isSet bool
}

func (v NullableMerchantWebhookResendPost200Response) Get() *MerchantWebhookResendPost200Response {
	return v.value
}

func (v *NullableMerchantWebhookResendPost200Response) Set(val *MerchantWebhookResendPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantWebhookResendPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantWebhookResendPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantWebhookResendPost200Response(val *MerchantWebhookResendPost200Response) *NullableMerchantWebhookResendPost200Response {
	return &NullableMerchantWebhookResendPost200Response{value: val, isSet: true}
}

func (v NullableMerchantWebhookResendPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantWebhookResendPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


