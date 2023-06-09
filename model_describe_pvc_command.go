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

// checks if the DescribePvcCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePvcCommand{}

// DescribePvcCommand struct for DescribePvcCommand
type DescribePvcCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribePvcCommand instantiates a new DescribePvcCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePvcCommand(projectId int32, name string, namespace string) *DescribePvcCommand {
	this := DescribePvcCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribePvcCommandWithDefaults instantiates a new DescribePvcCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePvcCommandWithDefaults() *DescribePvcCommand {
	this := DescribePvcCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribePvcCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribePvcCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribePvcCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribePvcCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribePvcCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribePvcCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribePvcCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribePvcCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribePvcCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribePvcCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePvcCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribePvcCommand struct {
	value *DescribePvcCommand
	isSet bool
}

func (v NullableDescribePvcCommand) Get() *DescribePvcCommand {
	return v.value
}

func (v *NullableDescribePvcCommand) Set(val *DescribePvcCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePvcCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePvcCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePvcCommand(val *DescribePvcCommand) *NullableDescribePvcCommand {
	return &NullableDescribePvcCommand{value: val, isSet: true}
}

func (v NullableDescribePvcCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePvcCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


