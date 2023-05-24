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

// checks if the DescribeDeploymentCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeDeploymentCommand{}

// DescribeDeploymentCommand struct for DescribeDeploymentCommand
type DescribeDeploymentCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeDeploymentCommand instantiates a new DescribeDeploymentCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDeploymentCommand(projectId int32, name string, namespace string) *DescribeDeploymentCommand {
	this := DescribeDeploymentCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeDeploymentCommandWithDefaults instantiates a new DescribeDeploymentCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDeploymentCommandWithDefaults() *DescribeDeploymentCommand {
	this := DescribeDeploymentCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeDeploymentCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeDeploymentCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeDeploymentCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeDeploymentCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeDeploymentCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeDeploymentCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeDeploymentCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeDeploymentCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeDeploymentCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeDeploymentCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDeploymentCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeDeploymentCommand struct {
	value *DescribeDeploymentCommand
	isSet bool
}

func (v NullableDescribeDeploymentCommand) Get() *DescribeDeploymentCommand {
	return v.value
}

func (v *NullableDescribeDeploymentCommand) Set(val *DescribeDeploymentCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDeploymentCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDeploymentCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDeploymentCommand(val *DescribeDeploymentCommand) *NullableDescribeDeploymentCommand {
	return &NullableDescribeDeploymentCommand{value: val, isSet: true}
}

func (v NullableDescribeDeploymentCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDeploymentCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


