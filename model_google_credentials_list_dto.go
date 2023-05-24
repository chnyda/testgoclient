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

// checks if the GoogleCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleCredentialsListDto{}

// GoogleCredentialsListDto struct for GoogleCredentialsListDto
type GoogleCredentialsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	PartnerLogo NullableString `json:"partnerLogo,omitempty"`
	PartnerName NullableString `json:"partnerName,omitempty"`
	FolderId NullableString `json:"folderId,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	BillingAccountId NullableString `json:"billingAccountId,omitempty"`
	Zones []string `json:"zones,omitempty"`
	Region NullableString `json:"region,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	BillingAccountName NullableString `json:"billingAccountName,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	ContinentName NullableString `json:"continentName,omitempty"`
}

// NewGoogleCredentialsListDto instantiates a new GoogleCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCredentialsListDto() *GoogleCredentialsListDto {
	this := GoogleCredentialsListDto{}
	return &this
}

// NewGoogleCredentialsListDtoWithDefaults instantiates a new GoogleCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCredentialsListDtoWithDefaults() *GoogleCredentialsListDto {
	this := GoogleCredentialsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GoogleCredentialsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GoogleCredentialsListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GoogleCredentialsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetName() {
	o.Name.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *GoogleCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *GoogleCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *GoogleCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *GoogleCredentialsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetPartnerLogo returns the PartnerLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetPartnerLogo() string {
	if o == nil || IsNil(o.PartnerLogo.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerLogo.Get()
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerLogo.Get(), o.PartnerLogo.IsSet()
}

// HasPartnerLogo returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasPartnerLogo() bool {
	if o != nil && o.PartnerLogo.IsSet() {
		return true
	}

	return false
}

// SetPartnerLogo gets a reference to the given NullableString and assigns it to the PartnerLogo field.
func (o *GoogleCredentialsListDto) SetPartnerLogo(v string) {
	o.PartnerLogo.Set(&v)
}
// SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil
func (o *GoogleCredentialsListDto) SetPartnerLogoNil() {
	o.PartnerLogo.Set(nil)
}

// UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetPartnerLogo() {
	o.PartnerLogo.Unset()
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerName.Get()
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetPartnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerName.Get(), o.PartnerName.IsSet()
}

// HasPartnerName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasPartnerName() bool {
	if o != nil && o.PartnerName.IsSet() {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given NullableString and assigns it to the PartnerName field.
func (o *GoogleCredentialsListDto) SetPartnerName(v string) {
	o.PartnerName.Set(&v)
}
// SetPartnerNameNil sets the value for PartnerName to be an explicit nil
func (o *GoogleCredentialsListDto) SetPartnerNameNil() {
	o.PartnerName.Set(nil)
}

// UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetPartnerName() {
	o.PartnerName.Unset()
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *GoogleCredentialsListDto) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *GoogleCredentialsListDto) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *GoogleCredentialsListDto) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *GoogleCredentialsListDto) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetBillingAccountId() string {
	if o == nil || IsNil(o.BillingAccountId.Get()) {
		var ret string
		return ret
	}
	return *o.BillingAccountId.Get()
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetBillingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAccountId.Get(), o.BillingAccountId.IsSet()
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasBillingAccountId() bool {
	if o != nil && o.BillingAccountId.IsSet() {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given NullableString and assigns it to the BillingAccountId field.
func (o *GoogleCredentialsListDto) SetBillingAccountId(v string) {
	o.BillingAccountId.Set(&v)
}
// SetBillingAccountIdNil sets the value for BillingAccountId to be an explicit nil
func (o *GoogleCredentialsListDto) SetBillingAccountIdNil() {
	o.BillingAccountId.Set(nil)
}

// UnsetBillingAccountId ensures that no value is present for BillingAccountId, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetBillingAccountId() {
	o.BillingAccountId.Unset()
}

// GetZones returns the Zones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetZones() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasZones() bool {
	if o != nil && IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *GoogleCredentialsListDto) SetZones(v []string) {
	o.Zones = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *GoogleCredentialsListDto) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *GoogleCredentialsListDto) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetRegion() {
	o.Region.Unset()
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *GoogleCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *GoogleCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetBillingAccountName returns the BillingAccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetBillingAccountName() string {
	if o == nil || IsNil(o.BillingAccountName.Get()) {
		var ret string
		return ret
	}
	return *o.BillingAccountName.Get()
}

// GetBillingAccountNameOk returns a tuple with the BillingAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetBillingAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAccountName.Get(), o.BillingAccountName.IsSet()
}

// HasBillingAccountName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasBillingAccountName() bool {
	if o != nil && o.BillingAccountName.IsSet() {
		return true
	}

	return false
}

// SetBillingAccountName gets a reference to the given NullableString and assigns it to the BillingAccountName field.
func (o *GoogleCredentialsListDto) SetBillingAccountName(v string) {
	o.BillingAccountName.Set(&v)
}
// SetBillingAccountNameNil sets the value for BillingAccountName to be an explicit nil
func (o *GoogleCredentialsListDto) SetBillingAccountNameNil() {
	o.BillingAccountName.Set(nil)
}

// UnsetBillingAccountName ensures that no value is present for BillingAccountName, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetBillingAccountName() {
	o.BillingAccountName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *GoogleCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GoogleCredentialsListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleCredentialsListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleCredentialsListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *GoogleCredentialsListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}
// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *GoogleCredentialsListDto) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *GoogleCredentialsListDto) UnsetContinentName() {
	o.ContinentName.Unset()
}

func (o GoogleCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.PartnerLogo.IsSet() {
		toSerialize["partnerLogo"] = o.PartnerLogo.Get()
	}
	if o.PartnerName.IsSet() {
		toSerialize["partnerName"] = o.PartnerName.Get()
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.BillingAccountId.IsSet() {
		toSerialize["billingAccountId"] = o.BillingAccountId.Get()
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.BillingAccountName.IsSet() {
		toSerialize["billingAccountName"] = o.BillingAccountName.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.ContinentName.IsSet() {
		toSerialize["continentName"] = o.ContinentName.Get()
	}
	return toSerialize, nil
}

type NullableGoogleCredentialsListDto struct {
	value *GoogleCredentialsListDto
	isSet bool
}

func (v NullableGoogleCredentialsListDto) Get() *GoogleCredentialsListDto {
	return v.value
}

func (v *NullableGoogleCredentialsListDto) Set(val *GoogleCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCredentialsListDto(val *GoogleCredentialsListDto) *NullableGoogleCredentialsListDto {
	return &NullableGoogleCredentialsListDto{value: val, isSet: true}
}

func (v NullableGoogleCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


