/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantUserSearchGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUserSearchGet200Response{}

// MerchantUserSearchGet200Response struct for MerchantUserSearchGet200Response
type MerchantUserSearchGet200Response struct {
	Code *int32 `json:"code,omitempty"`
	Data *MerchantUserSearchGet200ResponseData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Redirect *string `json:"redirect,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}

// NewMerchantUserSearchGet200Response instantiates a new MerchantUserSearchGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUserSearchGet200Response() *MerchantUserSearchGet200Response {
	this := MerchantUserSearchGet200Response{}
	return &this
}

// NewMerchantUserSearchGet200ResponseWithDefaults instantiates a new MerchantUserSearchGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUserSearchGet200ResponseWithDefaults() *MerchantUserSearchGet200Response {
	this := MerchantUserSearchGet200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200Response) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200Response) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *MerchantUserSearchGet200Response) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200Response) GetData() MerchantUserSearchGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret MerchantUserSearchGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200Response) GetDataOk() (*MerchantUserSearchGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MerchantUserSearchGet200ResponseData and assigns it to the Data field.
func (o *MerchantUserSearchGet200Response) SetData(v MerchantUserSearchGet200ResponseData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MerchantUserSearchGet200Response) SetMessage(v string) {
	o.Message = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200Response) GetRedirect() string {
	if o == nil || IsNil(o.Redirect) {
		var ret string
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200Response) GetRedirectOk() (*string, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200Response) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given string and assigns it to the Redirect field.
func (o *MerchantUserSearchGet200Response) SetRedirect(v string) {
	o.Redirect = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *MerchantUserSearchGet200Response) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUserSearchGet200Response) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *MerchantUserSearchGet200Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *MerchantUserSearchGet200Response) SetRequestId(v string) {
	o.RequestId = &v
}

func (o MerchantUserSearchGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUserSearchGet200Response) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantUserSearchGet200Response struct {
	value *MerchantUserSearchGet200Response
	isSet bool
}

func (v NullableMerchantUserSearchGet200Response) Get() *MerchantUserSearchGet200Response {
	return v.value
}

func (v *NullableMerchantUserSearchGet200Response) Set(val *MerchantUserSearchGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUserSearchGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUserSearchGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUserSearchGet200Response(val *MerchantUserSearchGet200Response) *NullableMerchantUserSearchGet200Response {
	return &NullableMerchantUserSearchGet200Response{value: val, isSet: true}
}

func (v NullableMerchantUserSearchGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUserSearchGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


