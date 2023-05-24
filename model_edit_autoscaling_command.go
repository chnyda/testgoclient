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

// checks if the EditAutoscalingCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAutoscalingCommand{}

// EditAutoscalingCommand struct for EditAutoscalingCommand
type EditAutoscalingCommand struct {
	ProjectId int32 `json:"projectId"`
	MinSize *int32 `json:"minSize,omitempty"`
	MaxSize *int32 `json:"maxSize,omitempty"`
}

// NewEditAutoscalingCommand instantiates a new EditAutoscalingCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAutoscalingCommand(projectId int32) *EditAutoscalingCommand {
	this := EditAutoscalingCommand{}
	this.ProjectId = projectId
	return &this
}

// NewEditAutoscalingCommandWithDefaults instantiates a new EditAutoscalingCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAutoscalingCommandWithDefaults() *EditAutoscalingCommand {
	this := EditAutoscalingCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *EditAutoscalingCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *EditAutoscalingCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *EditAutoscalingCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetMinSize returns the MinSize field value if set, zero value otherwise.
func (o *EditAutoscalingCommand) GetMinSize() int32 {
	if o == nil || IsNil(o.MinSize) {
		var ret int32
		return ret
	}
	return *o.MinSize
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAutoscalingCommand) GetMinSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSize) {
		return nil, false
	}
	return o.MinSize, true
}

// HasMinSize returns a boolean if a field has been set.
func (o *EditAutoscalingCommand) HasMinSize() bool {
	if o != nil && !IsNil(o.MinSize) {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given int32 and assigns it to the MinSize field.
func (o *EditAutoscalingCommand) SetMinSize(v int32) {
	o.MinSize = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *EditAutoscalingCommand) GetMaxSize() int32 {
	if o == nil || IsNil(o.MaxSize) {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAutoscalingCommand) GetMaxSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSize) {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *EditAutoscalingCommand) HasMaxSize() bool {
	if o != nil && !IsNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *EditAutoscalingCommand) SetMaxSize(v int32) {
	o.MaxSize = &v
}

func (o EditAutoscalingCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAutoscalingCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.MinSize) {
		toSerialize["minSize"] = o.MinSize
	}
	if !IsNil(o.MaxSize) {
		toSerialize["maxSize"] = o.MaxSize
	}
	return toSerialize, nil
}

type NullableEditAutoscalingCommand struct {
	value *EditAutoscalingCommand
	isSet bool
}

func (v NullableEditAutoscalingCommand) Get() *EditAutoscalingCommand {
	return v.value
}

func (v *NullableEditAutoscalingCommand) Set(val *EditAutoscalingCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAutoscalingCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAutoscalingCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAutoscalingCommand(val *EditAutoscalingCommand) *NullableEditAutoscalingCommand {
	return &NullableEditAutoscalingCommand{value: val, isSet: true}
}

func (v NullableEditAutoscalingCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAutoscalingCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


