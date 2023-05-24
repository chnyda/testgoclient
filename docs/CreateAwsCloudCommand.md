# CreateAwsCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AwsSecretAccessKey** | **string** |  | 
**AwsAccessKeyId** | **string** |  | 
**AzCount** | Pointer to **int32** |  | [optional] 
**AwsRegion** | **string** |  | 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateAwsCloudCommand

`func NewCreateAwsCloudCommand(name string, awsSecretAccessKey string, awsAccessKeyId string, awsRegion string, ) *CreateAwsCloudCommand`

NewCreateAwsCloudCommand instantiates a new CreateAwsCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAwsCloudCommandWithDefaults

`func NewCreateAwsCloudCommandWithDefaults() *CreateAwsCloudCommand`

NewCreateAwsCloudCommandWithDefaults instantiates a new CreateAwsCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAwsCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAwsCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAwsCloudCommand) SetName(v string)`

SetName sets Name field to given value.


### GetAwsSecretAccessKey

`func (o *CreateAwsCloudCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *CreateAwsCloudCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *CreateAwsCloudCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### GetAwsAccessKeyId

`func (o *CreateAwsCloudCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *CreateAwsCloudCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *CreateAwsCloudCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### GetAzCount

`func (o *CreateAwsCloudCommand) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *CreateAwsCloudCommand) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *CreateAwsCloudCommand) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *CreateAwsCloudCommand) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### GetAwsRegion

`func (o *CreateAwsCloudCommand) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *CreateAwsCloudCommand) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *CreateAwsCloudCommand) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetOrganizationId

`func (o *CreateAwsCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAwsCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAwsCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAwsCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAwsCloudCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAwsCloudCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


