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

// checks if the PatchPdbCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPdbCommand{}

// PatchPdbCommand struct for PatchPdbCommand
type PatchPdbCommand struct {
	ProjectId int32 `json:"projectId"`
	Yaml string `json:"yaml"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewPatchPdbCommand instantiates a new PatchPdbCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPdbCommand(projectId int32, yaml string, name string, namespace string) *PatchPdbCommand {
	this := PatchPdbCommand{}
	this.ProjectId = projectId
	this.Yaml = yaml
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewPatchPdbCommandWithDefaults instantiates a new PatchPdbCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPdbCommandWithDefaults() *PatchPdbCommand {
	this := PatchPdbCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchPdbCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchPdbCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchPdbCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetYaml returns the Yaml field value
func (o *PatchPdbCommand) GetYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
func (o *PatchPdbCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaml, true
}

// SetYaml sets field value
func (o *PatchPdbCommand) SetYaml(v string) {
	o.Yaml = v
}

// GetName returns the Name field value
func (o *PatchPdbCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchPdbCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchPdbCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *PatchPdbCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PatchPdbCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PatchPdbCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o PatchPdbCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPdbCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["yaml"] = o.Yaml
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullablePatchPdbCommand struct {
	value *PatchPdbCommand
	isSet bool
}

func (v NullablePatchPdbCommand) Get() *PatchPdbCommand {
	return v.value
}

func (v *NullablePatchPdbCommand) Set(val *PatchPdbCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPdbCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPdbCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPdbCommand(val *PatchPdbCommand) *NullablePatchPdbCommand {
	return &NullablePatchPdbCommand{value: val, isSet: true}
}

func (v NullablePatchPdbCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPdbCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


