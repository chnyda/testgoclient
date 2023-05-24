# SshUserCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SshPublicKey** | **string** |  | 

## Methods

### NewSshUserCreateDto

`func NewSshUserCreateDto(name string, sshPublicKey string, ) *SshUserCreateDto`

NewSshUserCreateDto instantiates a new SshUserCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserCreateDtoWithDefaults

`func NewSshUserCreateDtoWithDefaults() *SshUserCreateDto`

NewSshUserCreateDtoWithDefaults instantiates a new SshUserCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshUserCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshUserCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshUserCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetSshPublicKey

`func (o *SshUserCreateDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *SshUserCreateDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *SshUserCreateDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


