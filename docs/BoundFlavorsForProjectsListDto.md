# BoundFlavorsForProjectsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**IsAzure** | Pointer to **bool** |  | [optional] 
**IsAws** | Pointer to **bool** |  | [optional] 
**IsOpenstack** | Pointer to **bool** |  | [optional] 
**IsGoogle** | Pointer to **bool** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**MaxDataDiskCount** | Pointer to **NullableInt32** |  | [optional] 
**HasLinuxSpotPrice2** | Pointer to **bool** |  | [optional] 
**HasWindowsSpotPrice2** | Pointer to **bool** |  | [optional] 
**LinuxSpotPrice** | Pointer to **NullableString** |  | [optional] 
**LinuxPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsSpotPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsPrice** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBoundFlavorsForProjectsListDto

`func NewBoundFlavorsForProjectsListDto() *BoundFlavorsForProjectsListDto`

NewBoundFlavorsForProjectsListDto instantiates a new BoundFlavorsForProjectsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundFlavorsForProjectsListDtoWithDefaults

`func NewBoundFlavorsForProjectsListDtoWithDefaults() *BoundFlavorsForProjectsListDto`

NewBoundFlavorsForProjectsListDtoWithDefaults instantiates a new BoundFlavorsForProjectsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoundFlavorsForProjectsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoundFlavorsForProjectsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoundFlavorsForProjectsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoundFlavorsForProjectsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoundFlavorsForProjectsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundFlavorsForProjectsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundFlavorsForProjectsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoundFlavorsForProjectsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BoundFlavorsForProjectsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BoundFlavorsForProjectsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *BoundFlavorsForProjectsListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BoundFlavorsForProjectsListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BoundFlavorsForProjectsListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *BoundFlavorsForProjectsListDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *BoundFlavorsForProjectsListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *BoundFlavorsForProjectsListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *BoundFlavorsForProjectsListDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *BoundFlavorsForProjectsListDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetProjectId

`func (o *BoundFlavorsForProjectsListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BoundFlavorsForProjectsListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BoundFlavorsForProjectsListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BoundFlavorsForProjectsListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BoundFlavorsForProjectsListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BoundFlavorsForProjectsListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIsAzure

`func (o *BoundFlavorsForProjectsListDto) GetIsAzure() bool`

GetIsAzure returns the IsAzure field if non-nil, zero value otherwise.

### GetIsAzureOk

`func (o *BoundFlavorsForProjectsListDto) GetIsAzureOk() (*bool, bool)`

GetIsAzureOk returns a tuple with the IsAzure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzure

`func (o *BoundFlavorsForProjectsListDto) SetIsAzure(v bool)`

SetIsAzure sets IsAzure field to given value.

### HasIsAzure

`func (o *BoundFlavorsForProjectsListDto) HasIsAzure() bool`

HasIsAzure returns a boolean if a field has been set.

### GetIsAws

`func (o *BoundFlavorsForProjectsListDto) GetIsAws() bool`

GetIsAws returns the IsAws field if non-nil, zero value otherwise.

### GetIsAwsOk

`func (o *BoundFlavorsForProjectsListDto) GetIsAwsOk() (*bool, bool)`

GetIsAwsOk returns a tuple with the IsAws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAws

`func (o *BoundFlavorsForProjectsListDto) SetIsAws(v bool)`

SetIsAws sets IsAws field to given value.

### HasIsAws

`func (o *BoundFlavorsForProjectsListDto) HasIsAws() bool`

HasIsAws returns a boolean if a field has been set.

### GetIsOpenstack

`func (o *BoundFlavorsForProjectsListDto) GetIsOpenstack() bool`

GetIsOpenstack returns the IsOpenstack field if non-nil, zero value otherwise.

### GetIsOpenstackOk

`func (o *BoundFlavorsForProjectsListDto) GetIsOpenstackOk() (*bool, bool)`

GetIsOpenstackOk returns a tuple with the IsOpenstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenstack

`func (o *BoundFlavorsForProjectsListDto) SetIsOpenstack(v bool)`

SetIsOpenstack sets IsOpenstack field to given value.

### HasIsOpenstack

`func (o *BoundFlavorsForProjectsListDto) HasIsOpenstack() bool`

HasIsOpenstack returns a boolean if a field has been set.

### GetIsGoogle

`func (o *BoundFlavorsForProjectsListDto) GetIsGoogle() bool`

GetIsGoogle returns the IsGoogle field if non-nil, zero value otherwise.

### GetIsGoogleOk

`func (o *BoundFlavorsForProjectsListDto) GetIsGoogleOk() (*bool, bool)`

GetIsGoogleOk returns a tuple with the IsGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGoogle

`func (o *BoundFlavorsForProjectsListDto) SetIsGoogle(v bool)`

SetIsGoogle sets IsGoogle field to given value.

### HasIsGoogle

`func (o *BoundFlavorsForProjectsListDto) HasIsGoogle() bool`

HasIsGoogle returns a boolean if a field has been set.

### GetProjectName

`func (o *BoundFlavorsForProjectsListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BoundFlavorsForProjectsListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BoundFlavorsForProjectsListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BoundFlavorsForProjectsListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BoundFlavorsForProjectsListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BoundFlavorsForProjectsListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetMaxDataDiskCount

`func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCount() int32`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCountOk() (*int32, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCount(v int32)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.

### HasMaxDataDiskCount

`func (o *BoundFlavorsForProjectsListDto) HasMaxDataDiskCount() bool`

HasMaxDataDiskCount returns a boolean if a field has been set.

### SetMaxDataDiskCountNil

`func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *BoundFlavorsForProjectsListDto) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
### GetHasLinuxSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) GetHasLinuxSpotPrice2() bool`

GetHasLinuxSpotPrice2 returns the HasLinuxSpotPrice2 field if non-nil, zero value otherwise.

### GetHasLinuxSpotPrice2Ok

`func (o *BoundFlavorsForProjectsListDto) GetHasLinuxSpotPrice2Ok() (*bool, bool)`

GetHasLinuxSpotPrice2Ok returns a tuple with the HasLinuxSpotPrice2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinuxSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) SetHasLinuxSpotPrice2(v bool)`

SetHasLinuxSpotPrice2 sets HasLinuxSpotPrice2 field to given value.

### HasHasLinuxSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) HasHasLinuxSpotPrice2() bool`

HasHasLinuxSpotPrice2 returns a boolean if a field has been set.

### GetHasWindowsSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) GetHasWindowsSpotPrice2() bool`

GetHasWindowsSpotPrice2 returns the HasWindowsSpotPrice2 field if non-nil, zero value otherwise.

### GetHasWindowsSpotPrice2Ok

`func (o *BoundFlavorsForProjectsListDto) GetHasWindowsSpotPrice2Ok() (*bool, bool)`

GetHasWindowsSpotPrice2Ok returns a tuple with the HasWindowsSpotPrice2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWindowsSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) SetHasWindowsSpotPrice2(v bool)`

SetHasWindowsSpotPrice2 sets HasWindowsSpotPrice2 field to given value.

### HasHasWindowsSpotPrice2

`func (o *BoundFlavorsForProjectsListDto) HasHasWindowsSpotPrice2() bool`

HasHasWindowsSpotPrice2 returns a boolean if a field has been set.

### GetLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPrice() string`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPriceOk() (*string, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPrice(v string)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.

### HasLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) HasLinuxSpotPrice() bool`

HasLinuxSpotPrice returns a boolean if a field has been set.

### SetLinuxSpotPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
### GetLinuxPrice

`func (o *BoundFlavorsForProjectsListDto) GetLinuxPrice() string`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetLinuxPriceOk() (*string, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *BoundFlavorsForProjectsListDto) SetLinuxPrice(v string)`

SetLinuxPrice sets LinuxPrice field to given value.

### HasLinuxPrice

`func (o *BoundFlavorsForProjectsListDto) HasLinuxPrice() bool`

HasLinuxPrice returns a boolean if a field has been set.

### SetLinuxPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPrice() string`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPriceOk() (*string, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPrice(v string)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.

### HasWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) HasWindowsSpotPrice() bool`

HasWindowsSpotPrice returns a boolean if a field has been set.

### SetWindowsSpotPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
### GetWindowsPrice

`func (o *BoundFlavorsForProjectsListDto) GetWindowsPrice() string`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetWindowsPriceOk() (*string, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *BoundFlavorsForProjectsListDto) SetWindowsPrice(v string)`

SetWindowsPrice sets WindowsPrice field to given value.

### HasWindowsPrice

`func (o *BoundFlavorsForProjectsListDto) HasWindowsPrice() bool`

HasWindowsPrice returns a boolean if a field has been set.

### SetWindowsPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


