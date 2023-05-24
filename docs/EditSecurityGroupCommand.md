# EditSecurityGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Protocol** | Pointer to [**SecurityGroupProtocol**](SecurityGroupProtocol.md) |  | [optional] 
**PortMinRange** | Pointer to **int32** |  | [optional] 
**PortMaxRange** | Pointer to **int32** |  | [optional] 
**RemoteIpPrefix** | **string** |  | 

## Methods

### NewEditSecurityGroupCommand

`func NewEditSecurityGroupCommand(id int32, name string, remoteIpPrefix string, ) *EditSecurityGroupCommand`

NewEditSecurityGroupCommand instantiates a new EditSecurityGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSecurityGroupCommandWithDefaults

`func NewEditSecurityGroupCommandWithDefaults() *EditSecurityGroupCommand`

NewEditSecurityGroupCommandWithDefaults instantiates a new EditSecurityGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditSecurityGroupCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditSecurityGroupCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditSecurityGroupCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *EditSecurityGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditSecurityGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditSecurityGroupCommand) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *EditSecurityGroupCommand) GetProtocol() SecurityGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *EditSecurityGroupCommand) GetProtocolOk() (*SecurityGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *EditSecurityGroupCommand) SetProtocol(v SecurityGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *EditSecurityGroupCommand) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPortMinRange

`func (o *EditSecurityGroupCommand) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *EditSecurityGroupCommand) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *EditSecurityGroupCommand) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.

### HasPortMinRange

`func (o *EditSecurityGroupCommand) HasPortMinRange() bool`

HasPortMinRange returns a boolean if a field has been set.

### GetPortMaxRange

`func (o *EditSecurityGroupCommand) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *EditSecurityGroupCommand) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *EditSecurityGroupCommand) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.

### HasPortMaxRange

`func (o *EditSecurityGroupCommand) HasPortMaxRange() bool`

HasPortMaxRange returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *EditSecurityGroupCommand) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *EditSecurityGroupCommand) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *EditSecurityGroupCommand) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


