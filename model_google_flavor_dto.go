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

// checks if the GoogleFlavorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleFlavorDto{}

// GoogleFlavorDto struct for GoogleFlavorDto
type GoogleFlavorDto struct {
	Name NullableString `json:"name,omitempty"`
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Ram NullableInt64 `json:"ram,omitempty"`
	Description interface{} `json:"description,omitempty"`
	LinuxPrice NullableString `json:"linuxPrice,omitempty"`
	WindowsPrice NullableString `json:"windowsPrice,omitempty"`
	LinuxSpotPrice NullableString `json:"linuxSpotPrice,omitempty"`
	WindowsSpotPrice NullableString `json:"windowsSpotPrice,omitempty"`
}

// NewGoogleFlavorDto instantiates a new GoogleFlavorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleFlavorDto() *GoogleFlavorDto {
	this := GoogleFlavorDto{}
	return &this
}

// NewGoogleFlavorDtoWithDefaults instantiates a new GoogleFlavorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleFlavorDtoWithDefaults() *GoogleFlavorDto {
	this := GoogleFlavorDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GoogleFlavorDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GoogleFlavorDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GoogleFlavorDto) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *GoogleFlavorDto) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *GoogleFlavorDto) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *GoogleFlavorDto) UnsetCpu() {
	o.Cpu.Unset()
}

// GetRam returns the Ram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram.Get()) {
		var ret int64
		return ret
	}
	return *o.Ram.Get()
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetRamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ram.Get(), o.Ram.IsSet()
}

// HasRam returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasRam() bool {
	if o != nil && o.Ram.IsSet() {
		return true
	}

	return false
}

// SetRam gets a reference to the given NullableInt64 and assigns it to the Ram field.
func (o *GoogleFlavorDto) SetRam(v int64) {
	o.Ram.Set(&v)
}
// SetRamNil sets the value for Ram to be an explicit nil
func (o *GoogleFlavorDto) SetRamNil() {
	o.Ram.Set(nil)
}

// UnsetRam ensures that no value is present for Ram, not even an explicit nil
func (o *GoogleFlavorDto) UnsetRam() {
	o.Ram.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *GoogleFlavorDto) SetDescription(v interface{}) {
	o.Description = v
}

// GetLinuxPrice returns the LinuxPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetLinuxPrice() string {
	if o == nil || IsNil(o.LinuxPrice.Get()) {
		var ret string
		return ret
	}
	return *o.LinuxPrice.Get()
}

// GetLinuxPriceOk returns a tuple with the LinuxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetLinuxPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxPrice.Get(), o.LinuxPrice.IsSet()
}

// HasLinuxPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasLinuxPrice() bool {
	if o != nil && o.LinuxPrice.IsSet() {
		return true
	}

	return false
}

// SetLinuxPrice gets a reference to the given NullableString and assigns it to the LinuxPrice field.
func (o *GoogleFlavorDto) SetLinuxPrice(v string) {
	o.LinuxPrice.Set(&v)
}
// SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil
func (o *GoogleFlavorDto) SetLinuxPriceNil() {
	o.LinuxPrice.Set(nil)
}

// UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
func (o *GoogleFlavorDto) UnsetLinuxPrice() {
	o.LinuxPrice.Unset()
}

// GetWindowsPrice returns the WindowsPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetWindowsPrice() string {
	if o == nil || IsNil(o.WindowsPrice.Get()) {
		var ret string
		return ret
	}
	return *o.WindowsPrice.Get()
}

// GetWindowsPriceOk returns a tuple with the WindowsPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetWindowsPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsPrice.Get(), o.WindowsPrice.IsSet()
}

// HasWindowsPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasWindowsPrice() bool {
	if o != nil && o.WindowsPrice.IsSet() {
		return true
	}

	return false
}

// SetWindowsPrice gets a reference to the given NullableString and assigns it to the WindowsPrice field.
func (o *GoogleFlavorDto) SetWindowsPrice(v string) {
	o.WindowsPrice.Set(&v)
}
// SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil
func (o *GoogleFlavorDto) SetWindowsPriceNil() {
	o.WindowsPrice.Set(nil)
}

// UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
func (o *GoogleFlavorDto) UnsetWindowsPrice() {
	o.WindowsPrice.Unset()
}

// GetLinuxSpotPrice returns the LinuxSpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetLinuxSpotPrice() string {
	if o == nil || IsNil(o.LinuxSpotPrice.Get()) {
		var ret string
		return ret
	}
	return *o.LinuxSpotPrice.Get()
}

// GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetLinuxSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxSpotPrice.Get(), o.LinuxSpotPrice.IsSet()
}

// HasLinuxSpotPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasLinuxSpotPrice() bool {
	if o != nil && o.LinuxSpotPrice.IsSet() {
		return true
	}

	return false
}

// SetLinuxSpotPrice gets a reference to the given NullableString and assigns it to the LinuxSpotPrice field.
func (o *GoogleFlavorDto) SetLinuxSpotPrice(v string) {
	o.LinuxSpotPrice.Set(&v)
}
// SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil
func (o *GoogleFlavorDto) SetLinuxSpotPriceNil() {
	o.LinuxSpotPrice.Set(nil)
}

// UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
func (o *GoogleFlavorDto) UnsetLinuxSpotPrice() {
	o.LinuxSpotPrice.Unset()
}

// GetWindowsSpotPrice returns the WindowsSpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorDto) GetWindowsSpotPrice() string {
	if o == nil || IsNil(o.WindowsSpotPrice.Get()) {
		var ret string
		return ret
	}
	return *o.WindowsSpotPrice.Get()
}

// GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorDto) GetWindowsSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsSpotPrice.Get(), o.WindowsSpotPrice.IsSet()
}

// HasWindowsSpotPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasWindowsSpotPrice() bool {
	if o != nil && o.WindowsSpotPrice.IsSet() {
		return true
	}

	return false
}

// SetWindowsSpotPrice gets a reference to the given NullableString and assigns it to the WindowsSpotPrice field.
func (o *GoogleFlavorDto) SetWindowsSpotPrice(v string) {
	o.WindowsSpotPrice.Set(&v)
}
// SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil
func (o *GoogleFlavorDto) SetWindowsSpotPriceNil() {
	o.WindowsSpotPrice.Set(nil)
}

// UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
func (o *GoogleFlavorDto) UnsetWindowsSpotPrice() {
	o.WindowsSpotPrice.Unset()
}

func (o GoogleFlavorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleFlavorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Ram.IsSet() {
		toSerialize["ram"] = o.Ram.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LinuxPrice.IsSet() {
		toSerialize["linuxPrice"] = o.LinuxPrice.Get()
	}
	if o.WindowsPrice.IsSet() {
		toSerialize["windowsPrice"] = o.WindowsPrice.Get()
	}
	if o.LinuxSpotPrice.IsSet() {
		toSerialize["linuxSpotPrice"] = o.LinuxSpotPrice.Get()
	}
	if o.WindowsSpotPrice.IsSet() {
		toSerialize["windowsSpotPrice"] = o.WindowsSpotPrice.Get()
	}
	return toSerialize, nil
}

type NullableGoogleFlavorDto struct {
	value *GoogleFlavorDto
	isSet bool
}

func (v NullableGoogleFlavorDto) Get() *GoogleFlavorDto {
	return v.value
}

func (v *NullableGoogleFlavorDto) Set(val *GoogleFlavorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleFlavorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleFlavorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleFlavorDto(val *GoogleFlavorDto) *NullableGoogleFlavorDto {
	return &NullableGoogleFlavorDto{value: val, isSet: true}
}

func (v NullableGoogleFlavorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleFlavorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


