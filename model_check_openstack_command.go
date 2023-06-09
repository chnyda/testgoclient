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

// checks if the CheckOpenstackCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckOpenstackCommand{}

// CheckOpenstackCommand struct for CheckOpenstackCommand
type CheckOpenstackCommand struct {
	OpenStackUser string `json:"openStackUser"`
	OpenStackPassword string `json:"openStackPassword"`
	OpenStackUrl string `json:"openStackUrl"`
	OpenStackDomain string `json:"openStackDomain"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
}

// NewCheckOpenstackCommand instantiates a new CheckOpenstackCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckOpenstackCommand(openStackUser string, openStackPassword string, openStackUrl string, openStackDomain string) *CheckOpenstackCommand {
	this := CheckOpenstackCommand{}
	this.OpenStackUser = openStackUser
	this.OpenStackPassword = openStackPassword
	this.OpenStackUrl = openStackUrl
	this.OpenStackDomain = openStackDomain
	return &this
}

// NewCheckOpenstackCommandWithDefaults instantiates a new CheckOpenstackCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckOpenstackCommandWithDefaults() *CheckOpenstackCommand {
	this := CheckOpenstackCommand{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value
func (o *CheckOpenstackCommand) GetOpenStackUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenStackUser
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value
// and a boolean to check if the value has been set.
func (o *CheckOpenstackCommand) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenStackUser, true
}

// SetOpenStackUser sets field value
func (o *CheckOpenstackCommand) SetOpenStackUser(v string) {
	o.OpenStackUser = v
}

// GetOpenStackPassword returns the OpenStackPassword field value
func (o *CheckOpenstackCommand) GetOpenStackPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenStackPassword
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value
// and a boolean to check if the value has been set.
func (o *CheckOpenstackCommand) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenStackPassword, true
}

// SetOpenStackPassword sets field value
func (o *CheckOpenstackCommand) SetOpenStackPassword(v string) {
	o.OpenStackPassword = v
}

// GetOpenStackUrl returns the OpenStackUrl field value
func (o *CheckOpenstackCommand) GetOpenStackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenStackUrl
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value
// and a boolean to check if the value has been set.
func (o *CheckOpenstackCommand) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenStackUrl, true
}

// SetOpenStackUrl sets field value
func (o *CheckOpenstackCommand) SetOpenStackUrl(v string) {
	o.OpenStackUrl = v
}

// GetOpenStackDomain returns the OpenStackDomain field value
func (o *CheckOpenstackCommand) GetOpenStackDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenStackDomain
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value
// and a boolean to check if the value has been set.
func (o *CheckOpenstackCommand) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenStackDomain, true
}

// SetOpenStackDomain sets field value
func (o *CheckOpenstackCommand) SetOpenStackDomain(v string) {
	o.OpenStackDomain = v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *CheckOpenstackCommand) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckOpenstackCommand) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *CheckOpenstackCommand) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *CheckOpenstackCommand) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

func (o CheckOpenstackCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckOpenstackCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["openStackUser"] = o.OpenStackUser
	toSerialize["openStackPassword"] = o.OpenStackPassword
	toSerialize["openStackUrl"] = o.OpenStackUrl
	toSerialize["openStackDomain"] = o.OpenStackDomain
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	return toSerialize, nil
}

type NullableCheckOpenstackCommand struct {
	value *CheckOpenstackCommand
	isSet bool
}

func (v NullableCheckOpenstackCommand) Get() *CheckOpenstackCommand {
	return v.value
}

func (v *NullableCheckOpenstackCommand) Set(val *CheckOpenstackCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckOpenstackCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckOpenstackCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckOpenstackCommand(val *CheckOpenstackCommand) *NullableCheckOpenstackCommand {
	return &NullableCheckOpenstackCommand{value: val, isSet: true}
}

func (v NullableCheckOpenstackCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckOpenstackCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


