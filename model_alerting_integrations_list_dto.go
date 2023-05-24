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

// checks if the AlertingIntegrationsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingIntegrationsListDto{}

// AlertingIntegrationsListDto struct for AlertingIntegrationsListDto
type AlertingIntegrationsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Token NullableString `json:"token,omitempty"`
	AlertingIntegrationType NullableString `json:"alertingIntegrationType,omitempty"`
	AlertingProfileName NullableString `json:"alertingProfileName,omitempty"`
}

// NewAlertingIntegrationsListDto instantiates a new AlertingIntegrationsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingIntegrationsListDto() *AlertingIntegrationsListDto {
	this := AlertingIntegrationsListDto{}
	return &this
}

// NewAlertingIntegrationsListDtoWithDefaults instantiates a new AlertingIntegrationsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingIntegrationsListDtoWithDefaults() *AlertingIntegrationsListDto {
	this := AlertingIntegrationsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertingIntegrationsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingIntegrationsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertingIntegrationsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertingIntegrationsListDto) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingIntegrationsListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingIntegrationsListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AlertingIntegrationsListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AlertingIntegrationsListDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AlertingIntegrationsListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AlertingIntegrationsListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingIntegrationsListDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingIntegrationsListDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *AlertingIntegrationsListDto) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *AlertingIntegrationsListDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *AlertingIntegrationsListDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *AlertingIntegrationsListDto) UnsetToken() {
	o.Token.Unset()
}

// GetAlertingIntegrationType returns the AlertingIntegrationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingIntegrationsListDto) GetAlertingIntegrationType() string {
	if o == nil || IsNil(o.AlertingIntegrationType.Get()) {
		var ret string
		return ret
	}
	return *o.AlertingIntegrationType.Get()
}

// GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingIntegrationsListDto) GetAlertingIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertingIntegrationType.Get(), o.AlertingIntegrationType.IsSet()
}

// HasAlertingIntegrationType returns a boolean if a field has been set.
func (o *AlertingIntegrationsListDto) HasAlertingIntegrationType() bool {
	if o != nil && o.AlertingIntegrationType.IsSet() {
		return true
	}

	return false
}

// SetAlertingIntegrationType gets a reference to the given NullableString and assigns it to the AlertingIntegrationType field.
func (o *AlertingIntegrationsListDto) SetAlertingIntegrationType(v string) {
	o.AlertingIntegrationType.Set(&v)
}
// SetAlertingIntegrationTypeNil sets the value for AlertingIntegrationType to be an explicit nil
func (o *AlertingIntegrationsListDto) SetAlertingIntegrationTypeNil() {
	o.AlertingIntegrationType.Set(nil)
}

// UnsetAlertingIntegrationType ensures that no value is present for AlertingIntegrationType, not even an explicit nil
func (o *AlertingIntegrationsListDto) UnsetAlertingIntegrationType() {
	o.AlertingIntegrationType.Unset()
}

// GetAlertingProfileName returns the AlertingProfileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingIntegrationsListDto) GetAlertingProfileName() string {
	if o == nil || IsNil(o.AlertingProfileName.Get()) {
		var ret string
		return ret
	}
	return *o.AlertingProfileName.Get()
}

// GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingIntegrationsListDto) GetAlertingProfileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertingProfileName.Get(), o.AlertingProfileName.IsSet()
}

// HasAlertingProfileName returns a boolean if a field has been set.
func (o *AlertingIntegrationsListDto) HasAlertingProfileName() bool {
	if o != nil && o.AlertingProfileName.IsSet() {
		return true
	}

	return false
}

// SetAlertingProfileName gets a reference to the given NullableString and assigns it to the AlertingProfileName field.
func (o *AlertingIntegrationsListDto) SetAlertingProfileName(v string) {
	o.AlertingProfileName.Set(&v)
}
// SetAlertingProfileNameNil sets the value for AlertingProfileName to be an explicit nil
func (o *AlertingIntegrationsListDto) SetAlertingProfileNameNil() {
	o.AlertingProfileName.Set(nil)
}

// UnsetAlertingProfileName ensures that no value is present for AlertingProfileName, not even an explicit nil
func (o *AlertingIntegrationsListDto) UnsetAlertingProfileName() {
	o.AlertingProfileName.Unset()
}

func (o AlertingIntegrationsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingIntegrationsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.AlertingIntegrationType.IsSet() {
		toSerialize["alertingIntegrationType"] = o.AlertingIntegrationType.Get()
	}
	if o.AlertingProfileName.IsSet() {
		toSerialize["alertingProfileName"] = o.AlertingProfileName.Get()
	}
	return toSerialize, nil
}

type NullableAlertingIntegrationsListDto struct {
	value *AlertingIntegrationsListDto
	isSet bool
}

func (v NullableAlertingIntegrationsListDto) Get() *AlertingIntegrationsListDto {
	return v.value
}

func (v *NullableAlertingIntegrationsListDto) Set(val *AlertingIntegrationsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingIntegrationsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingIntegrationsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingIntegrationsListDto(val *AlertingIntegrationsListDto) *NullableAlertingIntegrationsListDto {
	return &NullableAlertingIntegrationsListDto{value: val, isSet: true}
}

func (v NullableAlertingIntegrationsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingIntegrationsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


