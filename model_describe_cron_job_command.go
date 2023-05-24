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

// checks if the DescribeCronJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeCronJobCommand{}

// DescribeCronJobCommand struct for DescribeCronJobCommand
type DescribeCronJobCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeCronJobCommand instantiates a new DescribeCronJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCronJobCommand(projectId int32, name string, namespace string) *DescribeCronJobCommand {
	this := DescribeCronJobCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeCronJobCommandWithDefaults instantiates a new DescribeCronJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCronJobCommandWithDefaults() *DescribeCronJobCommand {
	this := DescribeCronJobCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeCronJobCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeCronJobCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeCronJobCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeCronJobCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeCronJobCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeCronJobCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeCronJobCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeCronJobCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeCronJobCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeCronJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCronJobCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeCronJobCommand struct {
	value *DescribeCronJobCommand
	isSet bool
}

func (v NullableDescribeCronJobCommand) Get() *DescribeCronJobCommand {
	return v.value
}

func (v *NullableDescribeCronJobCommand) Set(val *DescribeCronJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCronJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCronJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCronJobCommand(val *DescribeCronJobCommand) *NullableDescribeCronJobCommand {
	return &NullableDescribeCronJobCommand{value: val, isSet: true}
}

func (v NullableDescribeCronJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCronJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

