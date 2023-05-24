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

// checks if the ContactUsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactUsCommand{}

// ContactUsCommand struct for ContactUsCommand
type ContactUsCommand struct {
	Name string `json:"name"`
	BusinessEmail string `json:"businessEmail"`
	CompanyName string `json:"companyName"`
	Comment NullableString `json:"comment,omitempty"`
}

// NewContactUsCommand instantiates a new ContactUsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactUsCommand(name string, businessEmail string, companyName string) *ContactUsCommand {
	this := ContactUsCommand{}
	this.Name = name
	this.BusinessEmail = businessEmail
	this.CompanyName = companyName
	return &this
}

// NewContactUsCommandWithDefaults instantiates a new ContactUsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactUsCommandWithDefaults() *ContactUsCommand {
	this := ContactUsCommand{}
	return &this
}

// GetName returns the Name field value
func (o *ContactUsCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContactUsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContactUsCommand) SetName(v string) {
	o.Name = v
}

// GetBusinessEmail returns the BusinessEmail field value
func (o *ContactUsCommand) GetBusinessEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessEmail
}

// GetBusinessEmailOk returns a tuple with the BusinessEmail field value
// and a boolean to check if the value has been set.
func (o *ContactUsCommand) GetBusinessEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessEmail, true
}

// SetBusinessEmail sets field value
func (o *ContactUsCommand) SetBusinessEmail(v string) {
	o.BusinessEmail = v
}

// GetCompanyName returns the CompanyName field value
func (o *ContactUsCommand) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *ContactUsCommand) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *ContactUsCommand) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactUsCommand) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactUsCommand) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ContactUsCommand) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ContactUsCommand) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ContactUsCommand) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ContactUsCommand) UnsetComment() {
	o.Comment.Unset()
}

func (o ContactUsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactUsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["businessEmail"] = o.BusinessEmail
	toSerialize["companyName"] = o.CompanyName
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableContactUsCommand struct {
	value *ContactUsCommand
	isSet bool
}

func (v NullableContactUsCommand) Get() *ContactUsCommand {
	return v.value
}

func (v *NullableContactUsCommand) Set(val *ContactUsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableContactUsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableContactUsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactUsCommand(val *ContactUsCommand) *NullableContactUsCommand {
	return &NullableContactUsCommand{value: val, isSet: true}
}

func (v NullableContactUsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactUsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


