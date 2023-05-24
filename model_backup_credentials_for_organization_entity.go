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

// checks if the BackupCredentialsForOrganizationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupCredentialsForOrganizationEntity{}

// BackupCredentialsForOrganizationEntity struct for BackupCredentialsForOrganizationEntity
type BackupCredentialsForOrganizationEntity struct {
	BackupCredentialId *int32 `json:"backupCredentialId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewBackupCredentialsForOrganizationEntity instantiates a new BackupCredentialsForOrganizationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupCredentialsForOrganizationEntity() *BackupCredentialsForOrganizationEntity {
	this := BackupCredentialsForOrganizationEntity{}
	return &this
}

// NewBackupCredentialsForOrganizationEntityWithDefaults instantiates a new BackupCredentialsForOrganizationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupCredentialsForOrganizationEntityWithDefaults() *BackupCredentialsForOrganizationEntity {
	this := BackupCredentialsForOrganizationEntity{}
	return &this
}

// GetBackupCredentialId returns the BackupCredentialId field value if set, zero value otherwise.
func (o *BackupCredentialsForOrganizationEntity) GetBackupCredentialId() int32 {
	if o == nil || IsNil(o.BackupCredentialId) {
		var ret int32
		return ret
	}
	return *o.BackupCredentialId
}

// GetBackupCredentialIdOk returns a tuple with the BackupCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsForOrganizationEntity) GetBackupCredentialIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BackupCredentialId) {
		return nil, false
	}
	return o.BackupCredentialId, true
}

// HasBackupCredentialId returns a boolean if a field has been set.
func (o *BackupCredentialsForOrganizationEntity) HasBackupCredentialId() bool {
	if o != nil && !IsNil(o.BackupCredentialId) {
		return true
	}

	return false
}

// SetBackupCredentialId gets a reference to the given int32 and assigns it to the BackupCredentialId field.
func (o *BackupCredentialsForOrganizationEntity) SetBackupCredentialId(v int32) {
	o.BackupCredentialId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsForOrganizationEntity) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsForOrganizationEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BackupCredentialsForOrganizationEntity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BackupCredentialsForOrganizationEntity) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BackupCredentialsForOrganizationEntity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BackupCredentialsForOrganizationEntity) UnsetName() {
	o.Name.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *BackupCredentialsForOrganizationEntity) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsForOrganizationEntity) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *BackupCredentialsForOrganizationEntity) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *BackupCredentialsForOrganizationEntity) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o BackupCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupCredentialsForOrganizationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupCredentialId) {
		toSerialize["backupCredentialId"] = o.BackupCredentialId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableBackupCredentialsForOrganizationEntity struct {
	value *BackupCredentialsForOrganizationEntity
	isSet bool
}

func (v NullableBackupCredentialsForOrganizationEntity) Get() *BackupCredentialsForOrganizationEntity {
	return v.value
}

func (v *NullableBackupCredentialsForOrganizationEntity) Set(val *BackupCredentialsForOrganizationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupCredentialsForOrganizationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupCredentialsForOrganizationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupCredentialsForOrganizationEntity(val *BackupCredentialsForOrganizationEntity) *NullableBackupCredentialsForOrganizationEntity {
	return &NullableBackupCredentialsForOrganizationEntity{value: val, isSet: true}
}

func (v NullableBackupCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupCredentialsForOrganizationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


