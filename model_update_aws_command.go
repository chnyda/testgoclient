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

// checks if the UpdateAwsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAwsCommand{}

// UpdateAwsCommand struct for UpdateAwsCommand
type UpdateAwsCommand struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey"`
	AwsAccessKeyId string `json:"awsAccessKeyId"`
}

// NewUpdateAwsCommand instantiates a new UpdateAwsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAwsCommand(id int32, name string, awsSecretAccessKey string, awsAccessKeyId string) *UpdateAwsCommand {
	this := UpdateAwsCommand{}
	this.Id = id
	this.Name = name
	this.AwsSecretAccessKey = awsSecretAccessKey
	this.AwsAccessKeyId = awsAccessKeyId
	return &this
}

// NewUpdateAwsCommandWithDefaults instantiates a new UpdateAwsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAwsCommandWithDefaults() *UpdateAwsCommand {
	this := UpdateAwsCommand{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateAwsCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAwsCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAwsCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UpdateAwsCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAwsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAwsCommand) SetName(v string) {
	o.Name = v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value
func (o *UpdateAwsCommand) GetAwsSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *UpdateAwsCommand) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsSecretAccessKey, true
}

// SetAwsSecretAccessKey sets field value
func (o *UpdateAwsCommand) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = v
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value
func (o *UpdateAwsCommand) GetAwsAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value
// and a boolean to check if the value has been set.
func (o *UpdateAwsCommand) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccessKeyId, true
}

// SetAwsAccessKeyId sets field value
func (o *UpdateAwsCommand) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = v
}

func (o UpdateAwsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAwsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	toSerialize["awsAccessKeyId"] = o.AwsAccessKeyId
	return toSerialize, nil
}

type NullableUpdateAwsCommand struct {
	value *UpdateAwsCommand
	isSet bool
}

func (v NullableUpdateAwsCommand) Get() *UpdateAwsCommand {
	return v.value
}

func (v *NullableUpdateAwsCommand) Set(val *UpdateAwsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAwsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAwsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAwsCommand(val *UpdateAwsCommand) *NullableUpdateAwsCommand {
	return &NullableUpdateAwsCommand{value: val, isSet: true}
}

func (v NullableUpdateAwsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAwsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


