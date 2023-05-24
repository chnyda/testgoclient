# DescribeNetworkPolicyCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewDescribeNetworkPolicyCommand

`func NewDescribeNetworkPolicyCommand(projectId int32, name string, namespace string, ) *DescribeNetworkPolicyCommand`

NewDescribeNetworkPolicyCommand instantiates a new DescribeNetworkPolicyCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNetworkPolicyCommandWithDefaults

`func NewDescribeNetworkPolicyCommandWithDefaults() *DescribeNetworkPolicyCommand`

NewDescribeNetworkPolicyCommandWithDefaults instantiates a new DescribeNetworkPolicyCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DescribeNetworkPolicyCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeNetworkPolicyCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeNetworkPolicyCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DescribeNetworkPolicyCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeNetworkPolicyCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeNetworkPolicyCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DescribeNetworkPolicyCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DescribeNetworkPolicyCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DescribeNetworkPolicyCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


