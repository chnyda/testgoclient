/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateAlertingIntegrationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertingIntegrationCommand{}

// CreateAlertingIntegrationCommand struct for CreateAlertingIntegrationCommand
type CreateAlertingIntegrationCommand struct {
	Url string `json:"url"`
	Token NullableString `json:"token,omitempty"`
	AlertingIntegrationType AlertingIntegrationType `json:"alertingIntegrationType"`
	AlertingProfileId int32 `json:"alertingProfileId"`
}

// NewCreateAlertingIntegrationCommand instantiates a new CreateAlertingIntegrationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertingIntegrationCommand(url string, alertingIntegrationType AlertingIntegrationType, alertingProfileId int32) *CreateAlertingIntegrationCommand {
	this := CreateAlertingIntegrationCommand{}
	this.Url = url
	this.AlertingIntegrationType = alertingIntegrationType
	this.AlertingProfileId = alertingProfileId
	return &this
}

// NewCreateAlertingIntegrationCommandWithDefaults instantiates a new CreateAlertingIntegrationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertingIntegrationCommandWithDefaults() *CreateAlertingIntegrationCommand {
	this := CreateAlertingIntegrationCommand{}
	return &this
}

// GetUrl returns the Url field value
func (o *CreateAlertingIntegrationCommand) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateAlertingIntegrationCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateAlertingIntegrationCommand) SetUrl(v string) {
	o.Url = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingIntegrationCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingIntegrationCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAlertingIntegrationCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *CreateAlertingIntegrationCommand) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *CreateAlertingIntegrationCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *CreateAlertingIntegrationCommand) UnsetToken() {
	o.Token.Unset()
}

// GetAlertingIntegrationType returns the AlertingIntegrationType field value
func (o *CreateAlertingIntegrationCommand) GetAlertingIntegrationType() AlertingIntegrationType {
	if o == nil {
		var ret AlertingIntegrationType
		return ret
	}

	return o.AlertingIntegrationType
}

// GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field value
// and a boolean to check if the value has been set.
func (o *CreateAlertingIntegrationCommand) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertingIntegrationType, true
}

// SetAlertingIntegrationType sets field value
func (o *CreateAlertingIntegrationCommand) SetAlertingIntegrationType(v AlertingIntegrationType) {
	o.AlertingIntegrationType = v
}

// GetAlertingProfileId returns the AlertingProfileId field value
func (o *CreateAlertingIntegrationCommand) GetAlertingProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertingProfileId
}

// GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field value
// and a boolean to check if the value has been set.
func (o *CreateAlertingIntegrationCommand) GetAlertingProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertingProfileId, true
}

// SetAlertingProfileId sets field value
func (o *CreateAlertingIntegrationCommand) SetAlertingProfileId(v int32) {
	o.AlertingProfileId = v
}

func (o CreateAlertingIntegrationCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlertingIntegrationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	toSerialize["alertingIntegrationType"] = o.AlertingIntegrationType
	toSerialize["alertingProfileId"] = o.AlertingProfileId
	return toSerialize, nil
}

type NullableCreateAlertingIntegrationCommand struct {
	value *CreateAlertingIntegrationCommand
	isSet bool
}

func (v NullableCreateAlertingIntegrationCommand) Get() *CreateAlertingIntegrationCommand {
	return v.value
}

func (v *NullableCreateAlertingIntegrationCommand) Set(val *CreateAlertingIntegrationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertingIntegrationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertingIntegrationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertingIntegrationCommand(val *CreateAlertingIntegrationCommand) *NullableCreateAlertingIntegrationCommand {
	return &NullableCreateAlertingIntegrationCommand{value: val, isSet: true}
}

func (v NullableCreateAlertingIntegrationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertingIntegrationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

