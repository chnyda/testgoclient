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

// checks if the UserResourceChartDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResourceChartDto{}

// UserResourceChartDto struct for UserResourceChartDto
type UserResourceChartDto struct {
	ProjectName NullableString `json:"projectName,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	DiskSize *int64 `json:"diskSize,omitempty"`
	Ram *int64 `json:"ram,omitempty"`
	Cpu *int64 `json:"cpu,omitempty"`
	MaxRam *int64 `json:"maxRam,omitempty"`
	MaxCpu *int64 `json:"maxCpu,omitempty"`
	MaxDiskSize *int64 `json:"maxDiskSize,omitempty"`
}

// NewUserResourceChartDto instantiates a new UserResourceChartDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResourceChartDto() *UserResourceChartDto {
	this := UserResourceChartDto{}
	return &this
}

// NewUserResourceChartDtoWithDefaults instantiates a new UserResourceChartDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResourceChartDtoWithDefaults() *UserResourceChartDto {
	this := UserResourceChartDto{}
	return &this
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResourceChartDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResourceChartDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *UserResourceChartDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *UserResourceChartDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *UserResourceChartDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *UserResourceChartDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetDiskSize() int64 {
	if o == nil || IsNil(o.DiskSize) {
		var ret int64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int64 and assigns it to the DiskSize field.
func (o *UserResourceChartDto) SetDiskSize(v int64) {
	o.DiskSize = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *UserResourceChartDto) SetRam(v int64) {
	o.Ram = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetCpu() int64 {
	if o == nil || IsNil(o.Cpu) {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *UserResourceChartDto) SetCpu(v int64) {
	o.Cpu = &v
}

// GetMaxRam returns the MaxRam field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetMaxRam() int64 {
	if o == nil || IsNil(o.MaxRam) {
		var ret int64
		return ret
	}
	return *o.MaxRam
}

// GetMaxRamOk returns a tuple with the MaxRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetMaxRamOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRam) {
		return nil, false
	}
	return o.MaxRam, true
}

// HasMaxRam returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasMaxRam() bool {
	if o != nil && !IsNil(o.MaxRam) {
		return true
	}

	return false
}

// SetMaxRam gets a reference to the given int64 and assigns it to the MaxRam field.
func (o *UserResourceChartDto) SetMaxRam(v int64) {
	o.MaxRam = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetMaxCpu() int64 {
	if o == nil || IsNil(o.MaxCpu) {
		var ret int64
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetMaxCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxCpu) {
		return nil, false
	}
	return o.MaxCpu, true
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasMaxCpu() bool {
	if o != nil && !IsNil(o.MaxCpu) {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given int64 and assigns it to the MaxCpu field.
func (o *UserResourceChartDto) SetMaxCpu(v int64) {
	o.MaxCpu = &v
}

// GetMaxDiskSize returns the MaxDiskSize field value if set, zero value otherwise.
func (o *UserResourceChartDto) GetMaxDiskSize() int64 {
	if o == nil || IsNil(o.MaxDiskSize) {
		var ret int64
		return ret
	}
	return *o.MaxDiskSize
}

// GetMaxDiskSizeOk returns a tuple with the MaxDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceChartDto) GetMaxDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxDiskSize) {
		return nil, false
	}
	return o.MaxDiskSize, true
}

// HasMaxDiskSize returns a boolean if a field has been set.
func (o *UserResourceChartDto) HasMaxDiskSize() bool {
	if o != nil && !IsNil(o.MaxDiskSize) {
		return true
	}

	return false
}

// SetMaxDiskSize gets a reference to the given int64 and assigns it to the MaxDiskSize field.
func (o *UserResourceChartDto) SetMaxDiskSize(v int64) {
	o.MaxDiskSize = &v
}

func (o UserResourceChartDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResourceChartDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.DiskSize) {
		toSerialize["diskSize"] = o.DiskSize
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.MaxRam) {
		toSerialize["maxRam"] = o.MaxRam
	}
	if !IsNil(o.MaxCpu) {
		toSerialize["maxCpu"] = o.MaxCpu
	}
	if !IsNil(o.MaxDiskSize) {
		toSerialize["maxDiskSize"] = o.MaxDiskSize
	}
	return toSerialize, nil
}

type NullableUserResourceChartDto struct {
	value *UserResourceChartDto
	isSet bool
}

func (v NullableUserResourceChartDto) Get() *UserResourceChartDto {
	return v.value
}

func (v *NullableUserResourceChartDto) Set(val *UserResourceChartDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResourceChartDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResourceChartDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResourceChartDto(val *UserResourceChartDto) *NullableUserResourceChartDto {
	return &NullableUserResourceChartDto{value: val, isSet: true}
}

func (v NullableUserResourceChartDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResourceChartDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

