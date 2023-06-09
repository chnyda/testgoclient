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

// checks if the DuplicateNameCheckerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicateNameCheckerCommand{}

// DuplicateNameCheckerCommand struct for DuplicateNameCheckerCommand
type DuplicateNameCheckerCommand struct {
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// NewDuplicateNameCheckerCommand instantiates a new DuplicateNameCheckerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicateNameCheckerCommand(type_ string, name string) *DuplicateNameCheckerCommand {
	this := DuplicateNameCheckerCommand{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewDuplicateNameCheckerCommandWithDefaults instantiates a new DuplicateNameCheckerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicateNameCheckerCommandWithDefaults() *DuplicateNameCheckerCommand {
	this := DuplicateNameCheckerCommand{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateNameCheckerCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateNameCheckerCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *DuplicateNameCheckerCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *DuplicateNameCheckerCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *DuplicateNameCheckerCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *DuplicateNameCheckerCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetType returns the Type field value
func (o *DuplicateNameCheckerCommand) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DuplicateNameCheckerCommand) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DuplicateNameCheckerCommand) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *DuplicateNameCheckerCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DuplicateNameCheckerCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DuplicateNameCheckerCommand) SetName(v string) {
	o.Name = v
}

func (o DuplicateNameCheckerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicateNameCheckerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableDuplicateNameCheckerCommand struct {
	value *DuplicateNameCheckerCommand
	isSet bool
}

func (v NullableDuplicateNameCheckerCommand) Get() *DuplicateNameCheckerCommand {
	return v.value
}

func (v *NullableDuplicateNameCheckerCommand) Set(val *DuplicateNameCheckerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicateNameCheckerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicateNameCheckerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicateNameCheckerCommand(val *DuplicateNameCheckerCommand) *NullableDuplicateNameCheckerCommand {
	return &NullableDuplicateNameCheckerCommand{value: val, isSet: true}
}

func (v NullableDuplicateNameCheckerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicateNameCheckerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


