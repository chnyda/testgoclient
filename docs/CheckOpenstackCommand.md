# CheckOpenstackCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | **string** |  | 
**OpenStackPassword** | **string** |  | 
**OpenStackUrl** | **string** |  | 
**OpenStackDomain** | **string** |  | 
**IsAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheckOpenstackCommand

`func NewCheckOpenstackCommand(openStackUser string, openStackPassword string, openStackUrl string, openStackDomain string, ) *CheckOpenstackCommand`

NewCheckOpenstackCommand instantiates a new CheckOpenstackCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckOpenstackCommandWithDefaults

`func NewCheckOpenstackCommandWithDefaults() *CheckOpenstackCommand`

NewCheckOpenstackCommandWithDefaults instantiates a new CheckOpenstackCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *CheckOpenstackCommand) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *CheckOpenstackCommand) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *CheckOpenstackCommand) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### GetOpenStackPassword

`func (o *CheckOpenstackCommand) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *CheckOpenstackCommand) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *CheckOpenstackCommand) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### GetOpenStackUrl

`func (o *CheckOpenstackCommand) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *CheckOpenstackCommand) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *CheckOpenstackCommand) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### GetOpenStackDomain

`func (o *CheckOpenstackCommand) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *CheckOpenstackCommand) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *CheckOpenstackCommand) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.


### GetIsAdmin

`func (o *CheckOpenstackCommand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CheckOpenstackCommand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CheckOpenstackCommand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CheckOpenstackCommand) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


