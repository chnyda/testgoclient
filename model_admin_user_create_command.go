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

// checks if the AdminUserCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminUserCreateCommand{}

// AdminUserCreateCommand struct for AdminUserCreateCommand
type AdminUserCreateCommand struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password interface{} `json:"password"`
	Role *UserRole `json:"role,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewAdminUserCreateCommand instantiates a new AdminUserCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUserCreateCommand(email string, username string, password interface{}) *AdminUserCreateCommand {
	this := AdminUserCreateCommand{}
	this.Email = email
	this.Username = username
	this.Password = password
	return &this
}

// NewAdminUserCreateCommandWithDefaults instantiates a new AdminUserCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUserCreateCommandWithDefaults() *AdminUserCreateCommand {
	this := AdminUserCreateCommand{}
	return &this
}

// GetEmail returns the Email field value
func (o *AdminUserCreateCommand) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AdminUserCreateCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AdminUserCreateCommand) SetEmail(v string) {
	o.Email = v
}

// GetUsername returns the Username field value
func (o *AdminUserCreateCommand) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AdminUserCreateCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AdminUserCreateCommand) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AdminUserCreateCommand) GetPassword() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AdminUserCreateCommand) GetPasswordOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AdminUserCreateCommand) SetPassword(v interface{}) {
	o.Password = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AdminUserCreateCommand) GetRole() UserRole {
	if o == nil || IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserCreateCommand) GetRoleOk() (*UserRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AdminUserCreateCommand) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *AdminUserCreateCommand) SetRole(v UserRole) {
	o.Role = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminUserCreateCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUserCreateCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AdminUserCreateCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *AdminUserCreateCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *AdminUserCreateCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *AdminUserCreateCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o AdminUserCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminUserCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableAdminUserCreateCommand struct {
	value *AdminUserCreateCommand
	isSet bool
}

func (v NullableAdminUserCreateCommand) Get() *AdminUserCreateCommand {
	return v.value
}

func (v *NullableAdminUserCreateCommand) Set(val *AdminUserCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUserCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUserCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUserCreateCommand(val *AdminUserCreateCommand) *NullableAdminUserCreateCommand {
	return &NullableAdminUserCreateCommand{value: val, isSet: true}
}

func (v NullableAdminUserCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUserCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


