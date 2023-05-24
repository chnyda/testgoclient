# CheckTanzuCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Url** | **string** |  | 
**Password** | **string** |  | 
**Namespace** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**VolumeType** | **string** |  | 

## Methods

### NewCheckTanzuCommand

`func NewCheckTanzuCommand(username string, url string, password string, namespace string, volumeType string, ) *CheckTanzuCommand`

NewCheckTanzuCommand instantiates a new CheckTanzuCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTanzuCommandWithDefaults

`func NewCheckTanzuCommandWithDefaults() *CheckTanzuCommand`

NewCheckTanzuCommandWithDefaults instantiates a new CheckTanzuCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CheckTanzuCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CheckTanzuCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CheckTanzuCommand) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUrl

`func (o *CheckTanzuCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckTanzuCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckTanzuCommand) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPassword

`func (o *CheckTanzuCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckTanzuCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CheckTanzuCommand) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetNamespace

`func (o *CheckTanzuCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CheckTanzuCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CheckTanzuCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPort

`func (o *CheckTanzuCommand) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CheckTanzuCommand) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CheckTanzuCommand) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CheckTanzuCommand) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CheckTanzuCommand) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CheckTanzuCommand) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetVolumeType

`func (o *CheckTanzuCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CheckTanzuCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CheckTanzuCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


