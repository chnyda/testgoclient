# PatchPdbCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Yaml** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewPatchPdbCommand

`func NewPatchPdbCommand(projectId int32, yaml string, name string, namespace string, ) *PatchPdbCommand`

NewPatchPdbCommand instantiates a new PatchPdbCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPdbCommandWithDefaults

`func NewPatchPdbCommandWithDefaults() *PatchPdbCommand`

NewPatchPdbCommandWithDefaults instantiates a new PatchPdbCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PatchPdbCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PatchPdbCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PatchPdbCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetYaml

`func (o *PatchPdbCommand) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *PatchPdbCommand) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *PatchPdbCommand) SetYaml(v string)`

SetYaml sets Yaml field to given value.


### GetName

`func (o *PatchPdbCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPdbCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPdbCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PatchPdbCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchPdbCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchPdbCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


