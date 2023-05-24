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

// checks if the DescribeJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeJobCommand{}

// DescribeJobCommand struct for DescribeJobCommand
type DescribeJobCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeJobCommand instantiates a new DescribeJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeJobCommand(projectId int32, name string, namespace string) *DescribeJobCommand {
	this := DescribeJobCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeJobCommandWithDefaults instantiates a new DescribeJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeJobCommandWithDefaults() *DescribeJobCommand {
	this := DescribeJobCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeJobCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeJobCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeJobCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeJobCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeJobCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeJobCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeJobCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeJobCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeJobCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeJobCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeJobCommand struct {
	value *DescribeJobCommand
	isSet bool
}

func (v NullableDescribeJobCommand) Get() *DescribeJobCommand {
	return v.value
}

func (v *NullableDescribeJobCommand) Set(val *DescribeJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeJobCommand(val *DescribeJobCommand) *NullableDescribeJobCommand {
	return &NullableDescribeJobCommand{value: val, isSet: true}
}

func (v NullableDescribeJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

