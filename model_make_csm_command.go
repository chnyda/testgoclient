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

// checks if the MakeCsmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MakeCsmCommand{}

// MakeCsmCommand struct for MakeCsmCommand
type MakeCsmCommand struct {
	UserId string `json:"userId"`
	Mode string `json:"mode"`
}

// NewMakeCsmCommand instantiates a new MakeCsmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMakeCsmCommand(userId string, mode string) *MakeCsmCommand {
	this := MakeCsmCommand{}
	this.UserId = userId
	this.Mode = mode
	return &this
}

// NewMakeCsmCommandWithDefaults instantiates a new MakeCsmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMakeCsmCommandWithDefaults() *MakeCsmCommand {
	this := MakeCsmCommand{}
	return &this
}

// GetUserId returns the UserId field value
func (o *MakeCsmCommand) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *MakeCsmCommand) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *MakeCsmCommand) SetUserId(v string) {
	o.UserId = v
}

// GetMode returns the Mode field value
func (o *MakeCsmCommand) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *MakeCsmCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *MakeCsmCommand) SetMode(v string) {
	o.Mode = v
}

func (o MakeCsmCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MakeCsmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["mode"] = o.Mode
	return toSerialize, nil
}

type NullableMakeCsmCommand struct {
	value *MakeCsmCommand
	isSet bool
}

func (v NullableMakeCsmCommand) Get() *MakeCsmCommand {
	return v.value
}

func (v *NullableMakeCsmCommand) Set(val *MakeCsmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMakeCsmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMakeCsmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMakeCsmCommand(val *MakeCsmCommand) *NullableMakeCsmCommand {
	return &NullableMakeCsmCommand{value: val, isSet: true}
}

func (v NullableMakeCsmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMakeCsmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


