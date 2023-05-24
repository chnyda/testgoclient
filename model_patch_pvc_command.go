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

// checks if the PatchPvcCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPvcCommand{}

// PatchPvcCommand struct for PatchPvcCommand
type PatchPvcCommand struct {
	ProjectId int32 `json:"projectId"`
	Yaml string `json:"yaml"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewPatchPvcCommand instantiates a new PatchPvcCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPvcCommand(projectId int32, yaml string, name string, namespace string) *PatchPvcCommand {
	this := PatchPvcCommand{}
	this.ProjectId = projectId
	this.Yaml = yaml
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewPatchPvcCommandWithDefaults instantiates a new PatchPvcCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPvcCommandWithDefaults() *PatchPvcCommand {
	this := PatchPvcCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchPvcCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchPvcCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchPvcCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetYaml returns the Yaml field value
func (o *PatchPvcCommand) GetYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
func (o *PatchPvcCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaml, true
}

// SetYaml sets field value
func (o *PatchPvcCommand) SetYaml(v string) {
	o.Yaml = v
}

// GetName returns the Name field value
func (o *PatchPvcCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchPvcCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchPvcCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *PatchPvcCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PatchPvcCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PatchPvcCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o PatchPvcCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPvcCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["yaml"] = o.Yaml
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullablePatchPvcCommand struct {
	value *PatchPvcCommand
	isSet bool
}

func (v NullablePatchPvcCommand) Get() *PatchPvcCommand {
	return v.value
}

func (v *NullablePatchPvcCommand) Set(val *PatchPvcCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPvcCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPvcCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPvcCommand(val *PatchPvcCommand) *NullablePatchPvcCommand {
	return &NullablePatchPvcCommand{value: val, isSet: true}
}

func (v NullablePatchPvcCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPvcCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

