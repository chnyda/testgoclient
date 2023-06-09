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

// checks if the CreateSecurityGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSecurityGroupCommand{}

// CreateSecurityGroupCommand struct for CreateSecurityGroupCommand
type CreateSecurityGroupCommand struct {
	Name string `json:"name"`
	Protocol *SecurityGroupProtocol `json:"protocol,omitempty"`
	PortMinRange *int32 `json:"portMinRange,omitempty"`
	PortMaxRange *int32 `json:"portMaxRange,omitempty"`
	RemoteIpPrefix string `json:"remoteIpPrefix"`
	StandAloneProfileId int32 `json:"standAloneProfileId"`
}

// NewCreateSecurityGroupCommand instantiates a new CreateSecurityGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityGroupCommand(name string, remoteIpPrefix string, standAloneProfileId int32) *CreateSecurityGroupCommand {
	this := CreateSecurityGroupCommand{}
	this.Name = name
	this.RemoteIpPrefix = remoteIpPrefix
	this.StandAloneProfileId = standAloneProfileId
	return &this
}

// NewCreateSecurityGroupCommandWithDefaults instantiates a new CreateSecurityGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityGroupCommandWithDefaults() *CreateSecurityGroupCommand {
	this := CreateSecurityGroupCommand{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSecurityGroupCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSecurityGroupCommand) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CreateSecurityGroupCommand) GetProtocol() SecurityGroupProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret SecurityGroupProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetProtocolOk() (*SecurityGroupProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *CreateSecurityGroupCommand) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given SecurityGroupProtocol and assigns it to the Protocol field.
func (o *CreateSecurityGroupCommand) SetProtocol(v SecurityGroupProtocol) {
	o.Protocol = &v
}

// GetPortMinRange returns the PortMinRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupCommand) GetPortMinRange() int32 {
	if o == nil || IsNil(o.PortMinRange) {
		var ret int32
		return ret
	}
	return *o.PortMinRange
}

// GetPortMinRangeOk returns a tuple with the PortMinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetPortMinRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMinRange) {
		return nil, false
	}
	return o.PortMinRange, true
}

// HasPortMinRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupCommand) HasPortMinRange() bool {
	if o != nil && !IsNil(o.PortMinRange) {
		return true
	}

	return false
}

// SetPortMinRange gets a reference to the given int32 and assigns it to the PortMinRange field.
func (o *CreateSecurityGroupCommand) SetPortMinRange(v int32) {
	o.PortMinRange = &v
}

// GetPortMaxRange returns the PortMaxRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupCommand) GetPortMaxRange() int32 {
	if o == nil || IsNil(o.PortMaxRange) {
		var ret int32
		return ret
	}
	return *o.PortMaxRange
}

// GetPortMaxRangeOk returns a tuple with the PortMaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetPortMaxRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMaxRange) {
		return nil, false
	}
	return o.PortMaxRange, true
}

// HasPortMaxRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupCommand) HasPortMaxRange() bool {
	if o != nil && !IsNil(o.PortMaxRange) {
		return true
	}

	return false
}

// SetPortMaxRange gets a reference to the given int32 and assigns it to the PortMaxRange field.
func (o *CreateSecurityGroupCommand) SetPortMaxRange(v int32) {
	o.PortMaxRange = &v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value
func (o *CreateSecurityGroupCommand) GetRemoteIpPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIpPrefix
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIpPrefix, true
}

// SetRemoteIpPrefix sets field value
func (o *CreateSecurityGroupCommand) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix = v
}

// GetStandAloneProfileId returns the StandAloneProfileId field value
func (o *CreateSecurityGroupCommand) GetStandAloneProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StandAloneProfileId
}

// GetStandAloneProfileIdOk returns a tuple with the StandAloneProfileId field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupCommand) GetStandAloneProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandAloneProfileId, true
}

// SetStandAloneProfileId sets field value
func (o *CreateSecurityGroupCommand) SetStandAloneProfileId(v int32) {
	o.StandAloneProfileId = v
}

func (o CreateSecurityGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSecurityGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.PortMinRange) {
		toSerialize["portMinRange"] = o.PortMinRange
	}
	if !IsNil(o.PortMaxRange) {
		toSerialize["portMaxRange"] = o.PortMaxRange
	}
	toSerialize["remoteIpPrefix"] = o.RemoteIpPrefix
	toSerialize["standAloneProfileId"] = o.StandAloneProfileId
	return toSerialize, nil
}

type NullableCreateSecurityGroupCommand struct {
	value *CreateSecurityGroupCommand
	isSet bool
}

func (v NullableCreateSecurityGroupCommand) Get() *CreateSecurityGroupCommand {
	return v.value
}

func (v *NullableCreateSecurityGroupCommand) Set(val *CreateSecurityGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityGroupCommand(val *CreateSecurityGroupCommand) *NullableCreateSecurityGroupCommand {
	return &NullableCreateSecurityGroupCommand{value: val, isSet: true}
}

func (v NullableCreateSecurityGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


