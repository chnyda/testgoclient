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

// checks if the CreateAwsCloudCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAwsCloudCommand{}

// CreateAwsCloudCommand struct for CreateAwsCloudCommand
type CreateAwsCloudCommand struct {
	Name string `json:"name"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey"`
	AwsAccessKeyId string `json:"awsAccessKeyId"`
	AzCount *int32 `json:"azCount,omitempty"`
	AwsRegion string `json:"awsRegion"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewCreateAwsCloudCommand instantiates a new CreateAwsCloudCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAwsCloudCommand(name string, awsSecretAccessKey string, awsAccessKeyId string, awsRegion string) *CreateAwsCloudCommand {
	this := CreateAwsCloudCommand{}
	this.Name = name
	this.AwsSecretAccessKey = awsSecretAccessKey
	this.AwsAccessKeyId = awsAccessKeyId
	this.AwsRegion = awsRegion
	return &this
}

// NewCreateAwsCloudCommandWithDefaults instantiates a new CreateAwsCloudCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAwsCloudCommandWithDefaults() *CreateAwsCloudCommand {
	this := CreateAwsCloudCommand{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAwsCloudCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAwsCloudCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAwsCloudCommand) SetName(v string) {
	o.Name = v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value
func (o *CreateAwsCloudCommand) GetAwsSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *CreateAwsCloudCommand) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsSecretAccessKey, true
}

// SetAwsSecretAccessKey sets field value
func (o *CreateAwsCloudCommand) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = v
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value
func (o *CreateAwsCloudCommand) GetAwsAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value
// and a boolean to check if the value has been set.
func (o *CreateAwsCloudCommand) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccessKeyId, true
}

// SetAwsAccessKeyId sets field value
func (o *CreateAwsCloudCommand) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = v
}

// GetAzCount returns the AzCount field value if set, zero value otherwise.
func (o *CreateAwsCloudCommand) GetAzCount() int32 {
	if o == nil || IsNil(o.AzCount) {
		var ret int32
		return ret
	}
	return *o.AzCount
}

// GetAzCountOk returns a tuple with the AzCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAwsCloudCommand) GetAzCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AzCount) {
		return nil, false
	}
	return o.AzCount, true
}

// HasAzCount returns a boolean if a field has been set.
func (o *CreateAwsCloudCommand) HasAzCount() bool {
	if o != nil && !IsNil(o.AzCount) {
		return true
	}

	return false
}

// SetAzCount gets a reference to the given int32 and assigns it to the AzCount field.
func (o *CreateAwsCloudCommand) SetAzCount(v int32) {
	o.AzCount = &v
}

// GetAwsRegion returns the AwsRegion field value
func (o *CreateAwsCloudCommand) GetAwsRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
func (o *CreateAwsCloudCommand) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRegion, true
}

// SetAwsRegion sets field value
func (o *CreateAwsCloudCommand) SetAwsRegion(v string) {
	o.AwsRegion = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAwsCloudCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAwsCloudCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateAwsCloudCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateAwsCloudCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateAwsCloudCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateAwsCloudCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o CreateAwsCloudCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAwsCloudCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	toSerialize["awsAccessKeyId"] = o.AwsAccessKeyId
	if !IsNil(o.AzCount) {
		toSerialize["azCount"] = o.AzCount
	}
	toSerialize["awsRegion"] = o.AwsRegion
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableCreateAwsCloudCommand struct {
	value *CreateAwsCloudCommand
	isSet bool
}

func (v NullableCreateAwsCloudCommand) Get() *CreateAwsCloudCommand {
	return v.value
}

func (v *NullableCreateAwsCloudCommand) Set(val *CreateAwsCloudCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAwsCloudCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAwsCloudCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAwsCloudCommand(val *CreateAwsCloudCommand) *NullableCreateAwsCloudCommand {
	return &NullableCreateAwsCloudCommand{value: val, isSet: true}
}

func (v NullableCreateAwsCloudCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAwsCloudCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

