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

// checks if the BecomePartnerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BecomePartnerCommand{}

// BecomePartnerCommand struct for BecomePartnerCommand
type BecomePartnerCommand struct {
	FullName string `json:"fullName"`
	Email string `json:"email"`
}

// NewBecomePartnerCommand instantiates a new BecomePartnerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBecomePartnerCommand(fullName string, email string) *BecomePartnerCommand {
	this := BecomePartnerCommand{}
	this.FullName = fullName
	this.Email = email
	return &this
}

// NewBecomePartnerCommandWithDefaults instantiates a new BecomePartnerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBecomePartnerCommandWithDefaults() *BecomePartnerCommand {
	this := BecomePartnerCommand{}
	return &this
}

// GetFullName returns the FullName field value
func (o *BecomePartnerCommand) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *BecomePartnerCommand) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *BecomePartnerCommand) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *BecomePartnerCommand) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BecomePartnerCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BecomePartnerCommand) SetEmail(v string) {
	o.Email = v
}

func (o BecomePartnerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BecomePartnerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullName"] = o.FullName
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableBecomePartnerCommand struct {
	value *BecomePartnerCommand
	isSet bool
}

func (v NullableBecomePartnerCommand) Get() *BecomePartnerCommand {
	return v.value
}

func (v *NullableBecomePartnerCommand) Set(val *BecomePartnerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBecomePartnerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBecomePartnerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBecomePartnerCommand(val *BecomePartnerCommand) *NullableBecomePartnerCommand {
	return &NullableBecomePartnerCommand{value: val, isSet: true}
}

func (v NullableBecomePartnerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBecomePartnerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


