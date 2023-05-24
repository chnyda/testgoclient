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

// checks if the KeycloakCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeycloakCreateCommand{}

// KeycloakCreateCommand struct for KeycloakCreateCommand
type KeycloakCreateCommand struct {
	Name string `json:"name"`
	Url string `json:"url"`
	RealmsName string `json:"realmsName"`
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// NewKeycloakCreateCommand instantiates a new KeycloakCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeycloakCreateCommand(name string, url string, realmsName string, clientId string, clientSecret string) *KeycloakCreateCommand {
	this := KeycloakCreateCommand{}
	this.Name = name
	this.Url = url
	this.RealmsName = realmsName
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewKeycloakCreateCommandWithDefaults instantiates a new KeycloakCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeycloakCreateCommandWithDefaults() *KeycloakCreateCommand {
	this := KeycloakCreateCommand{}
	return &this
}

// GetName returns the Name field value
func (o *KeycloakCreateCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeycloakCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeycloakCreateCommand) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *KeycloakCreateCommand) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *KeycloakCreateCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *KeycloakCreateCommand) SetUrl(v string) {
	o.Url = v
}

// GetRealmsName returns the RealmsName field value
func (o *KeycloakCreateCommand) GetRealmsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealmsName
}

// GetRealmsNameOk returns a tuple with the RealmsName field value
// and a boolean to check if the value has been set.
func (o *KeycloakCreateCommand) GetRealmsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealmsName, true
}

// SetRealmsName sets field value
func (o *KeycloakCreateCommand) SetRealmsName(v string) {
	o.RealmsName = v
}

// GetClientId returns the ClientId field value
func (o *KeycloakCreateCommand) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *KeycloakCreateCommand) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *KeycloakCreateCommand) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *KeycloakCreateCommand) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *KeycloakCreateCommand) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *KeycloakCreateCommand) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o KeycloakCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeycloakCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	toSerialize["realmsName"] = o.RealmsName
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	return toSerialize, nil
}

type NullableKeycloakCreateCommand struct {
	value *KeycloakCreateCommand
	isSet bool
}

func (v NullableKeycloakCreateCommand) Get() *KeycloakCreateCommand {
	return v.value
}

func (v *NullableKeycloakCreateCommand) Set(val *KeycloakCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKeycloakCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKeycloakCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeycloakCreateCommand(val *KeycloakCreateCommand) *NullableKeycloakCreateCommand {
	return &NullableKeycloakCreateCommand{value: val, isSet: true}
}

func (v NullableKeycloakCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeycloakCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

